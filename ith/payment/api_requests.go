package payment

import (
	"fmt"

	"github.com/pkg/errors"
	"gitlab.inn4science.com/vcg/go-common/api/httpx"
	"gitlab.inn4science.com/vcg/go-common/ith"
	"gitlab.inn4science.com/vcg/go-common/ith/auth"
)

type API struct {
	auth.API
}

func (api *API) CreateOrder(order *Order) (*CreateOrderRequest, error) {
	u := api.Config.GetURL(APIOrderCreate)

	httpResp, err := httpx.PostJSON(
		u.String(),
		&CreateOrderRequest{Order: order},
		map[string]string{
			auth.Header: auth.HeaderVal(api.Credentials.AccessToken),
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
			auth.Header: auth.HeaderVal(api.Credentials.AccessToken),
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
		DateFrom ith.Time `json:"dateFrom"`
		DateTo   ith.Time `json:"dateTo"`
	}{
		DateFrom: ith.Time("").FromTimestamp(since),
		DateTo:   ith.Time("").FromTimestamp(until),
	}

	httpResp, err := httpx.PostJSON(
		u.String(),
		&request,
		map[string]string{
			auth.Header: auth.HeaderVal(api.Credentials.AccessToken),
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
