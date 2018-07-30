package payment

import (
	"gitlab.inn4science.com/vcg/go-common/ith"
	"gitlab.inn4science.com/vcg/go-common/types/currency"
)

const (
	APIOrderCreate = "/paymentapi/order/create"
	APIOrderData   = "/paymentapi/order/details"
	APIOrdersList  = "/paymentapi/order/list"
)

type ErrorData struct {
	ith.ErrorData
}

type OrderDetailsRequest struct {
	UID             string `json:"uid"`
	ExternalOrderID string `json:"externalOrderId"`
}

type RefundRequest struct {
	UID     string        `json:"uid"`
	Amount  currency.Fiat `json:"amount"`
	Comment string        `json:"comment"`
}

type OrdersListResp struct {
	ErrorData *ErrorData   `json:"errorData,omitempty"`
	OrderList []OrderShort `json:"orderList,omitempty"`
}
type CreateOrderRequest struct {
	ErrorData   *ErrorData `json:"errorData,omitempty"`
	Order       *Order     `json:"order,omitempty"`
	RedirectUrl string     `json:"redirectUrl,omitempty"`
}
