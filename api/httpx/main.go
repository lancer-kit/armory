package httpx

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

type Headers map[string]string

const defaultTimeout = time.Second * 15

var httpClient = http.Client{Timeout: defaultTimeout}

// SetTimeout updated `httpClient` default timeout (15s).
func SetTimeout(duration time.Duration) {
	httpClient.Timeout = duration
}

// PostJSON, sets passed `headers` and `body` and executes RequestJSON with POST method.
func PostJSON(url string, body interface{}, headers map[string]string) (*http.Response, error) {
	return RequestJSON(http.MethodPost, url, body, headers)
}

// PutJSON, sets passed `headers` and `body` and executes RequestJSON with PUT method.
func PutJSON(url string, body interface{}, headers map[string]string) (*http.Response, error) {
	return RequestJSON(http.MethodPut, url, body, headers)
}

// GetJSON, sets passed `headers` and executes RequestJSON with GET method.
func GetJSON(url string, headers map[string]string) (*http.Response, error) {
	return RequestJSON(http.MethodGet, url, nil, headers)
}

// DeleteJSON, sets passed `headers` and executes RequestJSON with DELETE method.
func DeleteJSON(url string, headers map[string]string) (*http.Response, error) {
	return RequestJSON(http.MethodDelete, url, nil, headers)
}

// RequestJSON creates and executes new request with JSON content type.
func RequestJSON(method string, url string, data interface{}, headers map[string]string) (*http.Response, error) {
	var body io.Reader = nil

	if data != nil {
		rawData, err := json.Marshal(data)
		if err != nil {
			return nil, errors.Wrap(err, "unable to marshal body")
		}
		body = bytes.NewBuffer(rawData)
	}

	req, _ := http.NewRequest(method, url, body)
	req.Header.Set("Content-Type", "application/json")

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	return httpClient.Do(req)
}

func ParseJSONBody(r *http.Request, dest interface{}) error {
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(dest)
	if err != nil {
		return errors.Wrap(err, "failed to unmarshal request body")
	}
	return nil
}

func ParseJSONResult(httpResp *http.Response, dest interface{}) error {
	defer httpResp.Body.Close()
	err := json.NewDecoder(httpResp.Body).Decode(dest)
	if err != nil {
		return errors.Wrap(err, "failed to unmarshal response body")
	}
	return nil
}
