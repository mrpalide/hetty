// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package api

import (
	"fmt"
	"io"
	"net/url"
	"strconv"
	"time"

	"github.com/oklog/ulid"
)

type ClearHTTPRequestLogResult struct {
	Success bool `json:"success"`
}

type CloseProjectResult struct {
	Success bool `json:"success"`
}

type DeleteProjectResult struct {
	Success bool `json:"success"`
}

type DeleteSenderRequestsResult struct {
	Success bool `json:"success"`
}

type HTTPHeader struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type HTTPHeaderInput struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type HTTPRequestLog struct {
	ID        ulid.ULID        `json:"id"`
	URL       string           `json:"url"`
	Method    HTTPMethod       `json:"method"`
	Proto     string           `json:"proto"`
	Headers   []HTTPHeader     `json:"headers"`
	Body      *string          `json:"body"`
	Timestamp time.Time        `json:"timestamp"`
	Response  *HTTPResponseLog `json:"response"`
}

type HTTPRequestLogFilter struct {
	OnlyInScope      bool    `json:"onlyInScope"`
	SearchExpression *string `json:"searchExpression"`
}

type HTTPRequestLogFilterInput struct {
	OnlyInScope      *bool   `json:"onlyInScope"`
	SearchExpression *string `json:"searchExpression"`
}

type HTTPResponseLog struct {
	// Will be the same ID as its related request ID.
	ID           ulid.ULID    `json:"id"`
	Proto        HTTPProtocol `json:"proto"`
	StatusCode   int          `json:"statusCode"`
	StatusReason string       `json:"statusReason"`
	Body         *string      `json:"body"`
	Headers      []HTTPHeader `json:"headers"`
}

type Project struct {
	ID       ulid.ULID `json:"id"`
	Name     string    `json:"name"`
	IsActive bool      `json:"isActive"`
}

type ScopeHeader struct {
	Key   *string `json:"key"`
	Value *string `json:"value"`
}

type ScopeHeaderInput struct {
	Key   *string `json:"key"`
	Value *string `json:"value"`
}

type ScopeRule struct {
	URL    *string      `json:"url"`
	Header *ScopeHeader `json:"header"`
	Body   *string      `json:"body"`
}

type ScopeRuleInput struct {
	URL    *string           `json:"url"`
	Header *ScopeHeaderInput `json:"header"`
	Body   *string           `json:"body"`
}

type SenderRequest struct {
	ID                 ulid.ULID        `json:"id"`
	SourceRequestLogID *ulid.ULID       `json:"sourceRequestLogID"`
	URL                *url.URL         `json:"url"`
	Method             HTTPMethod       `json:"method"`
	Proto              HTTPProtocol     `json:"proto"`
	Headers            []HTTPHeader     `json:"headers"`
	Body               *string          `json:"body"`
	Timestamp          time.Time        `json:"timestamp"`
	Response           *HTTPResponseLog `json:"response"`
}

type SenderRequestFilter struct {
	OnlyInScope      bool    `json:"onlyInScope"`
	SearchExpression *string `json:"searchExpression"`
}

type SenderRequestFilterInput struct {
	OnlyInScope      *bool   `json:"onlyInScope"`
	SearchExpression *string `json:"searchExpression"`
}

type SenderRequestInput struct {
	ID      *ulid.ULID        `json:"id"`
	URL     *url.URL          `json:"url"`
	Method  *HTTPMethod       `json:"method"`
	Proto   *HTTPProtocol     `json:"proto"`
	Headers []HTTPHeaderInput `json:"headers"`
	Body    *string           `json:"body"`
}

type HTTPMethod string

const (
	HTTPMethodGet     HTTPMethod = "GET"
	HTTPMethodHead    HTTPMethod = "HEAD"
	HTTPMethodPost    HTTPMethod = "POST"
	HTTPMethodPut     HTTPMethod = "PUT"
	HTTPMethodDelete  HTTPMethod = "DELETE"
	HTTPMethodConnect HTTPMethod = "CONNECT"
	HTTPMethodOptions HTTPMethod = "OPTIONS"
	HTTPMethodTrace   HTTPMethod = "TRACE"
	HTTPMethodPatch   HTTPMethod = "PATCH"
)

var AllHTTPMethod = []HTTPMethod{
	HTTPMethodGet,
	HTTPMethodHead,
	HTTPMethodPost,
	HTTPMethodPut,
	HTTPMethodDelete,
	HTTPMethodConnect,
	HTTPMethodOptions,
	HTTPMethodTrace,
	HTTPMethodPatch,
}

func (e HTTPMethod) IsValid() bool {
	switch e {
	case HTTPMethodGet, HTTPMethodHead, HTTPMethodPost, HTTPMethodPut, HTTPMethodDelete, HTTPMethodConnect, HTTPMethodOptions, HTTPMethodTrace, HTTPMethodPatch:
		return true
	}
	return false
}

func (e HTTPMethod) String() string {
	return string(e)
}

func (e *HTTPMethod) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = HTTPMethod(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid HttpMethod", str)
	}
	return nil
}

func (e HTTPMethod) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type HTTPProtocol string

const (
	HTTPProtocolHTTP10 HTTPProtocol = "HTTP10"
	HTTPProtocolHTTP11 HTTPProtocol = "HTTP11"
	HTTPProtocolHTTP20 HTTPProtocol = "HTTP20"
)

var AllHTTPProtocol = []HTTPProtocol{
	HTTPProtocolHTTP10,
	HTTPProtocolHTTP11,
	HTTPProtocolHTTP20,
}

func (e HTTPProtocol) IsValid() bool {
	switch e {
	case HTTPProtocolHTTP10, HTTPProtocolHTTP11, HTTPProtocolHTTP20:
		return true
	}
	return false
}

func (e HTTPProtocol) String() string {
	return string(e)
}

func (e *HTTPProtocol) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = HTTPProtocol(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid HttpProtocol", str)
	}
	return nil
}

func (e HTTPProtocol) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
