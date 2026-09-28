package githubapi

import (
	"context"
	"io"
	"net/http"
)

// Client is the transport-verb seam over the GitHub REST API, exposing the
// union of the verbs both CLIs use against go-gh's *api.RESTClient. It is
// deliberately NOT a per-operation domain interface — domain shaping lives in
// the service layer, keeping this interface narrow.
//
// The concrete implementation is go-gh's *api.RESTClient (which satisfies every
// method structurally; NewClient / RequireAuthClient return one). Tests use the
// in-memory fake in cli/shared/githubtest.
type Client interface {
	// Get issues a GET and decodes the JSON body into resp (resp may be nil
	// for existence-only checks).
	Get(path string, resp interface{}) error
	// Post issues a POST with body and decodes the JSON response into resp
	// (resp may be nil).
	Post(path string, body io.Reader, resp interface{}) error
	// Patch issues a PATCH with body and decodes the JSON response into resp
	// (resp may be nil).
	Patch(path string, body io.Reader, resp interface{}) error
	// Request issues an arbitrary-method request and returns the raw response,
	// so callers can read headers (e.g., Link for pagination) or status codes
	// the decode-and-discard verbs hide.
	Request(method string, path string, body io.Reader) (*http.Response, error)
	// RequestWithContext is Request bound to a context, so callers can
	// cancel / deadline an in-flight request (go-gh's default client has no
	// HTTP timeout, so deadline-bounded enumerations need an external one).
	RequestWithContext(ctx context.Context, method string, path string, body io.Reader) (*http.Response, error)
}
