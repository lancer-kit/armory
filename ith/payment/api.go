package payment

import (
	"gitlab.inn4science.com/vcg/go-common/ith"
	"gitlab.inn4science.com/vcg/go-common/types/currency"
)

const (
	APIOrderCreate = "/paymentapi/order/create"
	APIOrderData   = "/paymentapi/order/details"
	APIOrdersList  = "/paymentapi/order/list"
	APIOrderRefund = "paymentapi/order/refund"
	APIOrderNewStatus = "/paymentapi/order/status"
	APIOrderDraft = "/paymentapi/order/draft/create"
	APIOrderDraftUpdate = "/paymentapi/order/draft/update"
	APIOrderDraftDelete = "/paymentapi/order/draft/delete"
	APIOrderDraftSend = "/paymentapi/order/draft/send"
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

type DeleteOrderUID struct {
	UID             string `json:"uid"`
}

type UpdateOrderStatusRequest struct {
	UID             string `json:"uid"`
	Status          string `json:"status"`
}