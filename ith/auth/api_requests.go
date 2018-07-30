package auth

import (
	"time"

	"github.com/pkg/errors"
	"gitlab.inn4science.com/vcg/go-common/api/httpx"
	"gitlab.inn4science.com/vcg/go-common/ith"
)

type API struct {
	ith.API
}

func (api *API) EnsureAuth() error {
	if !api.Credentials.Expired() {
		return nil
	}

	_, err := api.RefreshAuthToken("")
	return err
}

func (api *API) AuthHeader() map[string]string {
	return map[string]string{
		Header: HeaderVal(api.Credentials.AccessToken),
	}
}

func (api *API) GetAuthToken(request *Request) (response *Response, _ error) {
	if request == nil {
		request = &Request{
			Username: api.Config.Username,
			Password: api.Config.Password,
		}
	}

	u := api.Config.GetURL(APIAuthToken)
	httpResp, err := httpx.PostJSON(u.String(), &request, nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get auth token")
	}

	response = new(Response)
	err = httpx.ParseJSONResult(httpResp, response)
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse JSON")
	}

	api.Credentials.AccessToken = response.AccessToken
	api.Credentials.RefreshToken = response.RefreshToken
	api.Credentials.ExpiresIn = response.ExpiresIn
	api.Credentials.SetAt = time.Now().UTC().Unix()
	return response, err
}

func (api *API) RefreshAuthToken(refreshToken string) (response *Response, _ error) {
	if refreshToken == "" {
		refreshToken = api.Credentials.RefreshToken
	}

	u := api.Config.GetURL(APIAuthTokenRefresh)
	request := RefreshRequest{RefreshToken: refreshToken}

	httpResp, err := httpx.PostJSON(u.String(), &request, nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to refresh auth token")
	}

	response = new(Response)
	err = httpx.ParseJSONResult(httpResp, response)
	if err != nil {

	}
	api.Credentials.AccessToken = response.AccessToken
	api.Credentials.RefreshToken = response.RefreshToken
	api.Credentials.ExpiresIn = response.ExpiresIn
	api.Credentials.SetAt = time.Now().UTC().Unix()
	return response, err
}
