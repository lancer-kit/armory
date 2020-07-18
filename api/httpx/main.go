package httpx

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

// Headers is a type for request headers.
type Headers map[string]string

// Client is a interface of extended http.Client.
type Client interface {
	JSONClient
	CookiesSupport
	CustomHeadersSupport
	// Clone returns safe clone of Client.
	Clone() Client
	// SetHTTP - Set customized instance of http.Client
	SetHTTP(hc http.Client) Client
	// SetLogger - Set logger to enable log requests
	SetLogger(logger *logrus.Entry) Client
}

type CookiesSupport interface {
	// DefaultCookies returns a client's default cookies.
	DefaultCookies() []*http.Cookie
	// SetCookies sets a default cookies to a client.
	SetDefaultCookies(cookies []*http.Cookie) Client
	// RemoveDefaultCookies removes a default client's cookies.
	RemoveDefaultCookies() Client
	// WithCookies append cookies to a client and return new instance.
	WithCookies(cookies []*http.Cookie) Client
}

type CustomHeadersSupport interface {
	// DefaultHeaders returns a client's default headers.
	DefaultHeaders() Headers
	// SetDefaultHeaders sets a default headers to a client.
	SetDefaultHeaders(headers Headers) Client
	// SetHeader sets new default header to the client.
	SetHeader(key, val string) Client
	// RemoveDefaultHeaders removes a default client's headers.
	RemoveDefaultHeaders() Client
	// WithHeaders append headers to a client and return new instance.
	WithHeaders(headers Headers) Client
}

type JSONClient interface {
	// PostJSON, sets passed `headers` and `body` and executes RequestJSON with POST method.
	PostJSON(url string, body interface{}, headers Headers) (*http.Response, error)
	// PatchJSON, sets passed `headers` and `body` and executes RequestJSON with PATCH method.
	PatchJSON(url string, body interface{}, headers Headers) (*http.Response, error)
	// PutJSON, sets passed `headers` and `body` and executes RequestJSON with PUT method.
	PutJSON(url string, body interface{}, headers Headers) (*http.Response, error)
	// GetJSON, sets passed `headers` and executes RequestJSON with GET method.
	GetJSON(url string, headers Headers) (*http.Response, error)
	// DeleteJSON, sets passed `headers` and executes RequestJSON with DELETE method.
	DeleteJSON(url string, headers Headers) (*http.Response, error)
	// RequestJSON creates and executes new request with JSON content type.
	RequestJSON(method string, url string, data interface{}, headers Headers) (*http.Response, error)
	// ParseJSONBody decodes `json` body from the `http.Request`.
	ParseJSONBody(r *http.Request, dest interface{}) error
	// ParseJSONResult decodes `json` body from the `http.Response`.
	ParseJSONResult(httpResp *http.Response, dest interface{}) error
}

const defaultTimeout = time.Second * 15

// GetClient returns new Client.
func GetClient() Client {
	return NewXClient()
}

// PostJSON, sets passed `headers` and `body` and executes RequestJSON with POST method.
// Post issues a POST to the specified URL.
//
// Caller should close resp.Body when done reading from it.
//
// If the provided body is an io.Closer, it is closed after the
// request.
//
// Post is a wrapper around DefaultClient.Post.
//
// To set custom headers, use NewRequest and DefaultClient.Do.
//
// See the Client.Do method documentation for details on how redirects
// are handled.
func PostJSON(url string, body interface{}, headers map[string]string) (*http.Response, error) {
	return NewXClient().RequestJSON(http.MethodPost, url, body, headers)
}

// PutJSON, sets passed `headers` and `body` and executes RequestJSON with PUT method.
func PutJSON(url string, body interface{}, headers map[string]string) (*http.Response, error) {
	return NewXClient().RequestJSON(http.MethodPut, url, body, headers)
}

// GetJSON, sets passed `headers` and executes RequestJSON with GET method.
func GetJSON(url string, headers map[string]string) (*http.Response, error) {
	return NewXClient().RequestJSON(http.MethodGet, url, nil, headers)
}

// DeleteJSON, sets passed `headers` and executes RequestJSON with DELETE method.
func DeleteJSON(url string, headers map[string]string) (*http.Response, error) {
	return NewXClient().RequestJSON(http.MethodDelete, url, nil, headers)
}

// RequestJSON creates and executes new request with JSON content type.
func RequestJSON(method string, url string, data interface{}, headers map[string]string) (*http.Response, error) {
	return NewXClient().RequestJSON(method, url, data, headers)
}

// ParseJSONBody decodes `json` body from the `http.Request`.
func ParseJSONBody(r *http.Request, dest interface{}) error {
	return json.NewDecoder(r.Body).Decode(dest)
}

// ParseJSONResult decodes `json` body from the `http.Response`.
func ParseJSONResult(httpResp *http.Response, dest interface{}) error {
	return json.NewDecoder(httpResp.Body).Decode(dest)
}

// WithCookies returns default client with cookies.
func WithCookies(cookies []*http.Cookie) Client {
	return NewXClient().WithCookies(cookies)
}

// WithHeaders append headers to the client and return new instance.
func WithHeaders(headers Headers) Client {
	return NewXClient().WithHeaders(headers)
}
