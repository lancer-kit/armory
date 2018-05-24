package httpx

import (
	"encoding/json"
	"net/http"

	"bytes"

	"github.com/pkg/errors"
)

var httpClient = http.DefaultClient

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

func PostJSON(url string, data interface{}, headers map[string]string) (*http.Response, error) {
	rawData, err := json.Marshal(data)
	if err != nil {
		return nil, errors.Wrap(err, "unable to marshal body")
	}

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(rawData))
	req.Header.Set("Content-Type", "application/json")
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	return httpClient.Do(req)
}
