package payment

import (
	"fmt"

	"strconv"

	"github.com/pkg/errors"
	"gitlab.inn4science.com/vcg/go-common/api/httpx"
	"gitlab.inn4science.com/vcg/go-common/ith/auth"
	_ "gitlab.inn4science.com/vcg/go-common/ith/auth"
)

type API struct {
	auth.API
}

func (api *API) CreateOrder(order *Order) (*CreateOrderRequest, error) {
	u := api.Config.GetURL(APIOrderCreate)

	httpResp, err := httpx.PostJSON(
		u.String(),
		&CreateOrderRequest{Order: order},
		api.AuthHeader(),
	)
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
		api.AuthHeader(),
	)
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
	timeSince := strconv.Itoa(int(since))
	timeFrom := strconv.Itoa(int(until))

	request := struct {
		DateFrom string `json:"dateFrom"`
		DateTo   string `json:"dateTo"`
	}{

		DateFrom: timeSince,
		DateTo:   timeFrom,
	}

	httpResp, err := httpx.PostJSON(
		u.String(),
		&request,
		api.AuthHeader(),
	)
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

func (api *API) Refund(refund RefundRequest) (*RefundRequest, error) {
	u := api.Config.GetURL(APIOrderRefund)

	httpResp, err := httpx.PostJSON(
		u.String(),
		refund,
		api.AuthHeader(),
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to refresh auth token")
	}

	if httpResp.StatusCode != 200 {
		return nil, fmt.Errorf("request failed with status - %d", httpResp.StatusCode)
	}

	response := new(RefundRequest)
	err = httpx.ParseJSONResult(httpResp, response)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse response")
	}
	return response, err
}

func (api *API) SetOrderNewStatus(updateItem UpdateOrderStatusRequest) (*CreateOrderRequest, error) {
	u := api.Config.GetURL(APIOrderNewStatus)

	httpResp, err := httpx.PostJSON(
		u.String(),
		updateItem,
		api.AuthHeader(),
	)
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

func (api *API) CreateOrderDraft(order *Order) (*CreateOrderRequest, error) {
	u := api.Config.GetURL(APIOrderDraft)

	httpResp, err := httpx.PostJSON(
		u.String(),
		&CreateOrderRequest{Order: order},
		api.AuthHeader(),
	)
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

func (api *API) UpdateOrderDraft(order *Order) (*CreateOrderRequest, error) {
	u := api.Config.GetURL(APIOrderDraftUpdate)

	httpResp, err := httpx.PostJSON(
		u.String(),
		&CreateOrderRequest{Order: order},
		api.AuthHeader(),
	)
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

func (api *API) DeleteOrderDraft(uid string) error {
	u := api.Config.GetURL(APIOrderDraftDelete)

	httpResp, err := httpx.PostJSON(
		u.String(),
		&DeleteOrderUID{UID: uid},
		api.AuthHeader(),
	)
	if err != nil {
		return errors.Wrap(err, "failed to refresh auth token")
	}

	if httpResp.StatusCode != 200 {
		return fmt.Errorf("request failed with status - %d", httpResp.StatusCode)
	}
	return err
}

func (api *API) SendOrderDraft(uid string) (*CreateOrderRequest, error) {
	u := api.Config.GetURL(APIOrderDraftSend)

	httpResp, err := httpx.PostJSON(
		u.String(),
		&DeleteOrderUID{UID: uid},
		api.AuthHeader(),
	)

	if err != nil {
		return nil, errors.Wrap(err, "failed to refresh auth token")
	}

	if httpResp.StatusCode != 200 {
		return nil, fmt.Errorf("request failed with status - %d", httpResp.StatusCode)
	}

	response := new(CreateOrderRequest)
	return response, err
}

func (api *API) CreateExternalPurchaseOrder(order *Order) (*CreateOrderRequest, error) {
	u := api.Config.GetURL(APIOrderCreate)
	if order.ExternalPayout == nil {
		return nil, fmt.Errorf("External Payout is a required parameter")
	}
	httpResp, err := httpx.PostJSON(
		u.String(),
		&CreateOrderRequest{Order: order},
		api.AuthHeader(),
	)
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

func (api *API) GetOrderTariff(orderUid string) (*GetOrderTariffResponse, error) {
	u := api.Config.GetURL(APIGetOrderTariff + orderUid)
	httpResp, err := httpx.PostJSON(
		u.String(),
		orderUid,
		api.AuthHeader(),
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to refresh auth token")
	}

	if httpResp.StatusCode != 200 {
		return nil, fmt.Errorf("request failed with status - %d", httpResp.StatusCode)
	}

	response := new(GetOrderTariffResponse)
	err = httpx.ParseJSONResult(httpResp, response)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse response")
	}
	return response, err
}
