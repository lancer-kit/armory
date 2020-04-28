package httpx

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

const (
	HeaderBodyHash    = "X-Auth-BHash"
	HeaderSignature   = "X-Auth-Signature"
	HeaderSigner      = "X-Auth-Signer"
	HeaderService     = "X-Auth-Service"
	HeaderJWTParsed   = "jwt"
	HeaderHeadersList = "X-Custom-Headers"
)

type XClient struct {
	http.Client

	defaultHeaders Headers
	cookies        []*http.Cookie
	logger         *logrus.Entry
}

func NewXClient() *XClient {
	return &XClient{
		Client:         http.Client{Timeout: defaultTimeout},
		defaultHeaders: map[string]string{},
		cookies:        []*http.Cookie{},
	}
}

// SetHTTP - Set customized instance of http.Client
func (client *XClient) SetHTTP(hc http.Client) Client {
	newClient := client.clone()
	newClient.Client = hc
	return newClient
}

// SetLogger - Set logger to enable log requests
func (client *XClient) SetLogger(logger *logrus.Entry) Client {
	newClient := client.clone()
	newClient.logger = logger
	return newClient
}

// DefaultCookies returns a client's default cookies.
func (client *XClient) DefaultCookies() []*http.Cookie {
	return client.cookies
}

// SetCookies sets a default cookies to the client.
func (client *XClient) SetDefaultCookies(cookies []*http.Cookie) Client {
	newClient := client.clone()
	newClient.cookies = append(newClient.cookies, cookies...)
	return newClient
}

// RemoveDefaultCookies removes a default client's cookies.
func (client *XClient) RemoveDefaultCookies() Client {
	newClient := client.clone()
	newClient.cookies = nil
	return newClient
}

// WithCookies append cookies to the client and return new instance.
func (client *XClient) WithCookies(cookies []*http.Cookie) Client {
	newClient := client.clone()
	newClient.cookies = append(newClient.cookies, cookies...)
	return newClient
}

// DefaultHeaders returns a client's default headers.
func (client *XClient) DefaultHeaders() Headers {
	return client.defaultHeaders
}

// SetDefaultHeaders sets a default headers to the client.
func (client *XClient) SetDefaultHeaders(headers Headers) Client {
	newClient := client.clone()

	if newClient.defaultHeaders == nil {
		newClient.defaultHeaders = map[string]string{}
	}

	for key := range headers {
		newClient.defaultHeaders[key] = headers[key]
	}
	return newClient
}

// SetHeader sets new default header to the client.
func (client *XClient) SetHeader(key, val string) Client {
	newClient := client.clone()

	if newClient.defaultHeaders == nil {
		newClient.defaultHeaders = map[string]string{}
	}
	newClient.defaultHeaders[key] = val
	return newClient
}

// RemoveDefaultHeaders removes a default client's headers.
func (client *XClient) RemoveDefaultHeaders() Client {
	newClient := client.clone()
	newClient.defaultHeaders = map[string]string{}
	return newClient
}

// WithHeaders append headers to the client and return new instance.
func (client *XClient) WithHeaders(headers Headers) Client {
	newClient := client.clone()
	for key := range headers {
		newClient.defaultHeaders[key] = headers[key]
	}
	return newClient
}

// PostJSON, sets passed `headers` and `body` and executes RequestJSON with POST method.
func (client *XClient) PostJSON(url string, body interface{}, headers Headers) (*http.Response, error) {
	return client.RequestJSON(http.MethodPost, url, body, headers)
}

// PutJSON, sets passed `headers` and `body` and executes RequestJSON with PUT method.
func (client *XClient) PutJSON(url string, body interface{}, headers Headers) (*http.Response, error) {
	return client.RequestJSON(http.MethodPut, url, body, headers)
}

// PatchJSON, sets passed `headers` and `body` and executes RequestJSON with PATCH method.
func (client *XClient) PatchJSON(url string, body interface{}, headers Headers) (*http.Response, error) {
	return client.RequestJSON(http.MethodPatch, url, body, headers)
}

// GetJSON, sets passed `headers` and executes RequestJSON with GET method.
func (client *XClient) GetJSON(url string, headers Headers) (*http.Response, error) {
	return client.RequestJSON(http.MethodGet, url, nil, headers)
}

// DeleteJSON, sets passed `headers` and executes RequestJSON with DELETE method.
func (client *XClient) DeleteJSON(url string, headers Headers) (*http.Response, error) {
	return client.RequestJSON(http.MethodDelete, url, nil, headers)
}

// RequestJSON creates and executes new request with JSON content type.
func (client *XClient) RequestJSON(method string, url string, body interface{}, headers Headers) (*http.Response, error) {
	var rawData []byte
	switch v := body.(type) {
	case []byte:
		rawData = v
	default:
		var err error
		rawData, err = json.Marshal(body)
		if err != nil {
			return nil, errors.Wrap(err, "unable to marshal body")
		}
	}

	if client.logger != nil {
		client.logger.WithFields(logrus.Fields{
			"method":  method,
			"url":     url,
			"headers": headers,
			"body":    string(rawData),
		}).Trace("do json request")
	}

	var bodyBuf io.Reader
	bodyBuf = bytes.NewBuffer(rawData)

	req, err := http.NewRequest(method, url, bodyBuf)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	for _, k := range client.cookies {
		req.AddCookie(k)
	}

	for key, value := range client.defaultHeaders {
		req.Header.Set(key, value)
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	return client.Do(req)
}

// ParseJSONBody decodes `json` body from the `http.Request`.
// !> `dest` must be a pointer value.
func (client *XClient) ParseJSONBody(r *http.Request, dest interface{}) error {
	defer r.Body.Close()
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return errors.Wrap(err, "failed to unmarshal request body")
	}

	if client.logger != nil {
		client.logger.WithFields(logrus.Fields{
			"method":      r.Method,
			"remote_url":  r.RemoteAddr,
			"request_url": r.RequestURI,
			"body":        string(b)}).Trace("parse json request")
	}

	err = json.Unmarshal(b, dest)
	if err != nil {
		return errors.Wrap(err, "failed to unmarshal request body")
	}
	return nil
}

// ParseJSONResult decodes `json` body from the `http.Response` body into `dest`
// > `dest` must be a pointer value.
func (client *XClient) ParseJSONResult(httpResp *http.Response, dest interface{}) error {
	defer httpResp.Body.Close()
	b, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		return errors.Wrap(err, "failed to unmarshal request body")
	}

	if client != nil && client.logger != nil {
		client.logger.WithFields(logrus.Fields{
			"method":      httpResp.Request.Method,
			"remote_url":  httpResp.Request.RemoteAddr,
			"request_url": httpResp.Request.RequestURI,
			"status":      httpResp.StatusCode,
			"body":        string(b)}).Trace("parse json response")
	}

	err = json.Unmarshal(b, dest)
	if err != nil {
		return errors.Wrap(err, "failed to unmarshal response body")
	}
	return nil
}

func (client *XClient) Clone() Client {
	return client.clone()
}

func (client *XClient) clone() *XClient {
	var clone = NewXClient()
	if client == nil {
		return clone
	}

	if client.logger != nil {
		clone.logger = client.logger.WithField("", "")
	}

	if len(client.cookies) > 0 {
		copy(clone.cookies, client.cookies)
	}

	if len(client.defaultHeaders) > 0 {
		for name := range client.defaultHeaders {
			clone.defaultHeaders[name] = client.defaultHeaders[name]
		}
	}

	return clone
}
