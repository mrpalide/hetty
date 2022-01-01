package proj

import (
	"context"
	"errors"
	"fmt"
	"log"
	"regexp"
	"sync"
)

type (
	OnProjectOpenFn  func(name string) error
	OnProjectCloseFn func(name string) error
)

// Service is used for managing projects.
type Service interface {
	Open(ctx context.Context, name string) (Project, error)
	Close() error
	Delete(name string) error
	ActiveProject() (Project, error)
	Projects() ([]Project, error)
	OnProjectOpen(fn OnProjectOpenFn)
	OnProjectClose(fn OnProjectCloseFn)
}

type service struct {
	repo              Repository
	activeProject     string
	onProjectOpenFns  []OnProjectOpenFn
	onProjectCloseFns []OnProjectCloseFn
	mu                sync.RWMutex
}

type Project struct {
	Name     string
	IsActive bool
}

var (
	ErrNoProject   = errors.New("proj: no open project")
	ErrNoSettings  = errors.New("proj: settings not found")
	ErrInvalidName = errors.New("proj: invalid name, must be alphanumeric or whitespace chars")
)

var nameRegexp = regexp.MustCompile(`^[\w\d\s]+$`)

// NewService returns a new Service.
func NewService(repo Repository) (Service, error) {
	return &service{
		repo: repo,
	}, nil
}

// Close closes the currently open project database (if there is one).
func (svc *service) Close() error {
	svc.mu.Lock()
	defer svc.mu.Unlock()

	closedProject := svc.activeProject

	if err := svc.repo.Close(); err != nil {
		return fmt.Errorf("proj: could not close project: %w", err)
	}

	svc.activeProject = ""

	svc.emitProjectClosed(closedProject)

	return nil
}

// Delete removes a project database file from disk (if there is one).
func (svc *service) Delete(name string) error {
	if name == "" {
		return errors.New("proj: name cannot be empty")
	}

	if svc.activeProject == name {
		return fmt.Errorf("proj: project (%v) is active", name)
	}

	if err := svc.repo.DeleteProject(name); err != nil {
		return fmt.Errorf("proj: could not delete project: %w", err)
	}

	return nil
}

// Open opens a database identified with `name`. If a database with this
// identifier doesn't exist yet, it will be automatically created.
func (svc *service) Open(ctx context.Context, name string) (Project, error) {
	if !nameRegexp.MatchString(name) {
		return Project{}, ErrInvalidName
	}

	svc.mu.Lock()
	defer svc.mu.Unlock()

	if err := svc.repo.Close(); err != nil {
		return Project{}, fmt.Errorf("proj: could not close previously open database: %w", err)
	}

	if err := svc.repo.OpenProject(name); err != nil {
		return Project{}, fmt.Errorf("proj: could not open database: %w", err)
	}

	svc.activeProject = name
	svc.emitProjectOpened()

	return Project{
		Name:     name,
		IsActive: true,
	}, nil
}

func (svc *service) ActiveProject() (Project, error) {
	activeProject := svc.activeProject
	if activeProject == "" {
		return Project{}, ErrNoProject
	}

	return Project{
		Name: activeProject,
	}, nil
}

func (svc *service) Projects() ([]Project, error) {
	projects, err := svc.repo.Projects()
	if err != nil {
		return nil, fmt.Errorf("proj: could not get projects: %w", err)
	}

	return projects, nil
}

func (svc *service) OnProjectOpen(fn OnProjectOpenFn) {
	svc.mu.Lock()
	defer svc.mu.Unlock()

	svc.onProjectOpenFns = append(svc.onProjectOpenFns, fn)
}

func (svc *service) OnProjectClose(fn OnProjectCloseFn) {
	svc.mu.Lock()
	defer svc.mu.Unlock()

	svc.onProjectCloseFns = append(svc.onProjectCloseFns, fn)
}

func (svc *service) emitProjectOpened() {
	for _, fn := range svc.onProjectOpenFns {
		if err := fn(svc.activeProject); err != nil {
			log.Printf("[ERROR] Could not execute onProjectOpen function: %v", err)
		}
	}
}

func (svc *service) emitProjectClosed(name string) {
	for _, fn := range svc.onProjectCloseFns {
		if err := fn(name); err != nil {
			log.Printf("[ERROR] Could not execute onProjectClose function: %v", err)
		}
	}
}
