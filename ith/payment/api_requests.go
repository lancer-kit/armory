package payment

import (
	"net/url"

	"time"

	"github.com/pkg/errors"
	"gitlab.inn4science.com/vcg/go-common/api/httpx"
)

type Config struct {
	BaseURL  *url.URL
	Path     string
	Username string
	Password string
}

func (cfg *Config) GetURL(path string) *url.URL {
	ur := *cfg.BaseURL
	ur.Path = path
	return &ur
}

type API struct {
	Config      Config
	Credentials struct {
		AccessToken  string
		RefreshToken string
		SetAt        int64
		ExpiresIn    int64
	}
}

func (api *API) GetAuthToken(request *AuthRequest) (response *AuthResponse, _ error) {
	if request == nil {
		request = &AuthRequest{
			Username: api.Config.Username,
			Password: api.Config.Password,
		}
	}

	u := api.Config.GetURL(APIAuthToken)
	httpResp, err := httpx.PostJSON(u.String(), &request, nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get auth token")
	}
	response = new(AuthResponse)
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

func (api *API) RefreshAuthToken(refreshToken string) (response *AuthResponse, _ error) {
	u := api.Config.GetURL(APIAuthTokenRefresh)
	request := RefreshRequest{RefreshToken: refreshToken}

	httpResp, err := httpx.PostJSON(u.String(), &request, nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to refresh auth token")
	}

	response = new(AuthResponse)
	err = httpx.ParseJSONResult(httpResp, response)
	if err != nil {

	}
	return response, err
}

func (api *API) GetOrderList() ([]OrderShort, error) {
	u := api.Config.GetURL(APIOrdersList)
	request := struct {
		DateFrom string `json:"dateFrom"`
		DateTo   string `json:"dateTo"`
	}{DateFrom: "20160101120000", DateTo: "20180520120000"}

	httpResp, err := httpx.PostJSON(
		u.String(),
		&request,
		map[string]string{
			AuthHeader: AuthHeaderVal(api.Credentials.AccessToken),
		})
	if err != nil {
		return nil, errors.Wrap(err, "failed to refresh auth token")
	}
	print(httpResp.Status)

	response := make([]OrderShort, 0)
	data := make([]byte, 0)
	httpResp.Body.Read(data)
	print(string(data))
	//err = httpx.ParseJSONResult(httpResp, response)
	//if err != nil {
	//
	//}
	return response, err
}
