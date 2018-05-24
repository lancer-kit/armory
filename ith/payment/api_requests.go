package payment

import (
	"net/url"

	"time"

	"fmt"

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

func (api *API) CreateOrder(order *Order) (*CreateOrderRequest, error) {
	u := api.Config.GetURL(APIOrderCreate)

	httpResp, err := httpx.PostJSON(
		u.String(),
		&CreateOrderRequest{Order: order},
		map[string]string{
			AuthHeader: AuthHeaderVal(api.Credentials.AccessToken),
		})
	if err != nil {
		return nil, errors.Wrap(err, "failed to refresh auth token")
	}

	if httpResp.StatusCode != 200 {
		return nil, fmt.Errorf("request failed with status - %d", httpResp.StatusCode)
	}

	response := new(CreateOrderRequest)
	err = httpx.ParseJSONResult(httpResp, response)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse response")
	}
	return response, err
}

func (api *API) GetOrderDetails(uid, externalOrderId string) (*OrdersListResp, error) {
	u := api.Config.GetURL(APIOrderData)
	req := &struct {
		UID             string `json:"uid,omitempty"`
		ExternalOrderId string `json:"externalOrderId,omitempty"`
	}{
		UID:             uid,
		ExternalOrderId: externalOrderId,
	}

	httpResp, err := httpx.PostJSON(
		u.String(),
		req,
		map[string]string{
			AuthHeader: AuthHeaderVal(api.Credentials.AccessToken),
		})
	if err != nil {
		return nil, errors.Wrap(err, "failed to refresh auth token")
	}

	if httpResp.StatusCode != 200 {
		return nil, fmt.Errorf("request failed with status - %d", httpResp.StatusCode)
	}

	response := new(OrdersListResp)
	err = httpx.ParseJSONResult(httpResp, response)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse response")
	}
	return response, err
}

func (api *API) GetOrderList(since, until int64) (*OrdersListResp, error) {
	u := api.Config.GetURL(APIOrdersList)
	request := struct {
		DateFrom IthTime `json:"dateFrom"`
		DateTo   IthTime `json:"dateTo"`
	}{
		DateFrom: IthTime("").FromTimestamp(since),
		DateTo:   IthTime("").FromTimestamp(until),
	}

	httpResp, err := httpx.PostJSON(
		u.String(),
		&request,
		map[string]string{
			AuthHeader: AuthHeaderVal(api.Credentials.AccessToken),
		})
	if err != nil {
		return nil, errors.Wrap(err, "failed to refresh auth token")
	}

	if httpResp.StatusCode != 200 {
		return nil, fmt.Errorf("request failed with status - %d", httpResp.StatusCode)
	}

	response := new(OrdersListResp)
	err = httpx.ParseJSONResult(httpResp, response)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse response")
	}
	return response, err
}
