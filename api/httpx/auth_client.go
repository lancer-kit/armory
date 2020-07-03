package httpx

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/lancer-kit/armory/crypto"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type SecuredClient interface {
	Client

	CloneWAuth() SecuredClient

	// Auth returns current state of authentication flag.
	Auth() bool
	// OnAuth disables request authentication.
	OffAuth() SecuredClient
	// OnAuth enables request authentication.
	OnAuth() SecuredClient
	// PublicKey returns client public key.
	PublicKey() crypto.Key
	// Service returns auth service name.
	Service() string
	// SetAuth sets the auth credentials.
	SetAuth(service string, kp crypto.KP) SecuredClient
	// SignRequest takes body hash, some headers and full URL path,
	// sings this request details using the `client.privateKey` and adds the auth headers.
	SignRequest(req *http.Request, body []byte, headers map[string]string) (*http.Request, error)
	// VerifyBody checks the request body match with it hash.
	VerifyBody(r *http.Request, body []byte) (bool, error)
	// VerifyRequest checks the request auth headers.
	VerifyRequest(r *http.Request, publicKey string) (bool, error)

	PostSignedWithHeaders(url string, data interface{}, headers map[string]string) (*http.Response, error)
	GetSignedWithHeaders(url string, headers map[string]string) (*http.Response, error)
}

type SXClient struct {
	XClient

	auth    bool
	kp      crypto.KP
	service string
}

func NewSXClient() *SXClient {
	xClient := NewXClient()
	return &SXClient{
		XClient: *xClient,
	}
}

// WithAuth returns default client with set auth data.
func WithAuth(service string, kp crypto.KP) SecuredClient {
	return NewSXClient().SetAuth(service, kp)
}

// Auth returns current state of authentication flag.
func (client *SXClient) Auth() bool {
	return client.auth
}

func (client *SXClient) OffAuth() SecuredClient {
	client.auth = false
	return client
}

// OnAuth enables request authentication.
func (client *SXClient) OnAuth() SecuredClient {
	client.auth = true
	return client
}

// Service returns auth service name.
func (client *SXClient) Service() string {
	return client.service
}

// PublicKey returns client public key.
func (client *SXClient) PublicKey() crypto.Key {
	return client.kp.Public
}

// SetAuth sets the auth credentials.
func (client *SXClient) SetAuth(service string, kp crypto.KP) SecuredClient {
	newClient := client.clone()

	newClient.kp = kp
	newClient.auth = true
	newClient.service = service

	return newClient
}

// PostJSON, sets passed `headers` and `body` and executes RequestJSON with POST method.
func (client *SXClient) PostJSON(url string, bodyStruct interface{}, headers Headers) (*http.Response, error) {
	return client.RequestJSON(http.MethodPost, url, bodyStruct, headers)
}

// PutJSON, sets passed `headers` and `body` and executes RequestJSON with PUT method.
func (client *SXClient) PutJSON(url string, bodyStruct interface{}, headers Headers) (*http.Response, error) {
	return client.RequestJSON(http.MethodPut, url, bodyStruct, headers)
}

// PatchJSON, sets passed `headers` and `body` and executes RequestJSON with PATCH method.
func (client *SXClient) PatchJSON(url string, bodyStruct interface{}, headers Headers) (*http.Response, error) {
	return client.RequestJSON(http.MethodPatch, url, bodyStruct, headers)
}

// GetJSON, sets passed `headers` and executes RequestJSON with GET method.
func (client *SXClient) GetJSON(url string, headers Headers) (*http.Response, error) {
	return client.RequestJSON(http.MethodGet, url, nil, headers)
}

// DeleteJSON, sets passed `headers` and executes RequestJSON with DELETE method.
func (client *SXClient) DeleteJSON(url string, headers Headers) (*http.Response, error) {
	return client.RequestJSON(http.MethodDelete, url, nil, headers)
}

// RequestJSON creates and executes new request with JSON content type.
func (client *SXClient) RequestJSON(method string, url string, bodyStruct interface{}, headers Headers) (*http.Response, error) {
	var rawData []byte

	switch v := bodyStruct.(type) {
	case []byte:
		rawData = v
	default:
		var err error
		rawData, err = json.Marshal(bodyStruct)
		if err != nil {
			return nil, errors.Wrap(err, "unable to marshal body")
		}
	}

	body := bytes.NewBuffer(rawData)

	if client.logger != nil {
		client.logger.WithFields(logrus.Fields{
			"method":  method,
			"url":     url,
			"headers": headers,
			"body":    string(rawData)}).Debug("do json request")
	}

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	if len(client.cookies) != 0 {
		req = addCookies(req, client.cookies)
	}

	req.Header.Set("Content-Type", "application/json")

	for key, value := range headers {
		client.defaultHeaders[key] = value
	}
	for key, value := range client.defaultHeaders {
		req.Header.Set(key, value)
	}

	if client.auth {
		req, err = client.SignRequest(req, rawData, headers)
		if err != nil {
			return nil, errors.Wrap(err, "unable to sign request")
		}
	}
	return client.Do(req)
}

// SignRequest takes body hash, some headers and full URL path,
// sings this request details using the `client.privateKey` and adds the auth headers.
func (client *SXClient) SignRequest(req *http.Request, body []byte, headers map[string]string) (*http.Request, error) {
	bodyHash, err := crypto.HashData(body)
	if err != nil {
		return nil, errors.Wrap(err, "failed to hash body")
	}

	fullPath := req.URL.Path + req.URL.RawQuery
	signHeaders := headersForSigning(headers)
	msg := messageForSigning(client.service, req.Method, fullPath,
		bodyHash, signHeaders)

	sign, err := client.kp.Sign([]byte(msg))
	if err != nil {
		return nil, errors.Wrap(err, "failed to sign message")
	}

	req.Header.Set(HeaderHeadersList, signHeaders)
	req.Header.Set(HeaderBodyHash, bodyHash)
	req.Header.Set(HeaderSignature, sign)
	req.Header.Set(HeaderService, client.service)
	req.Header.Set(HeaderSigner, client.kp.Public.String())
	return req, nil
}

// VerifyBody checks the request body match with it hash.
func (client *SXClient) VerifyBody(r *http.Request, body []byte) (bool, error) {
	bodyHash, err := crypto.HashData(body)
	if err != nil {
		return false, errors.Wrap(err, "failed to hash body")
	}

	return bodyHash == r.Header.Get(HeaderBodyHash), nil
}

// VerifyRequest checks the request auth headers.
func (client *SXClient) VerifyRequest(r *http.Request, publicKey string) (bool, error) {
	if publicKey != r.Header.Get(HeaderSigner) {
		return false, errors.New("signer mismatch with passed public key")
	}

	bodyHash := r.Header.Get(HeaderBodyHash)
	service := r.Header.Get(HeaderService)
	sign := r.Header.Get(HeaderSignature)
	headers := r.Header.Get(HeaderHeadersList)

	fullPath := r.URL.Path + r.URL.RawQuery
	msg := messageForSigning(service, r.Method, fullPath, bodyHash, headers)

	return crypto.VerifySignature(publicKey, msg, sign)
}

// PostSignedWithHeaders create new POST signed request with headers
func (client *SXClient) PostSignedWithHeaders(
	url string, data interface{}, headers map[string]string) (*http.Response, error) {
	rawData, err := json.Marshal(data)
	if err != nil {
		return nil, errors.Wrap(err, "unable to marshal body")
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(rawData))
	if err != nil {
		return nil, errors.Wrap(err, "failed to create new http request")
	}

	rg, err := client.SignRequest(req, rawData, headers)
	if err != nil {
		return nil, errors.Wrap(err, "unable to create request")
	}

	for key, value := range headers {
		rg.Header.Set(key, value)
	}

	return client.Do(rg)
}

// PostSignedWithHeaders create new signed GET request with headers
func (client *SXClient) GetSignedWithHeaders(url string, headers map[string]string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create new http request")
	}

	rq, err := client.SignRequest(req, nil, headers)
	if err != nil {
		return nil, errors.Wrap(err, "unable to create request")
	}
	for key, value := range headers {
		rq.Header.Set(key, value)
	}

	return client.Do(rq)
}

func (client *SXClient) CloneWAuth() SecuredClient {
	return client.clone()
}

func (client *SXClient) clone() *SXClient {
	var clone = NewSXClient()
	if client == nil {
		return clone
	}
	clone.XClient = *client.XClient.clone()
	clone.auth = client.auth
	clone.kp = client.kp
	clone.service = client.service

	return clone
}
