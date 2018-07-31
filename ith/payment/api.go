package payment

import (
	"gitlab.inn4science.com/vcg/go-common/ith"
	"gitlab.inn4science.com/vcg/go-common/types/currency"
)

const (
	APIOrderCreate      = "/paymentapi/order/create"
	APIOrderData        = "/paymentapi/order/details"
	APIOrdersList       = "/paymentapi/order/list"
	APIOrderRefund      = "paymentapi/order/refund"
	APIOrderNewStatus   = "/paymentapi/order/status"
	APIOrderDraft       = "/paymentapi/order/draft/create"
	APIOrderDraftUpdate = "/paymentapi/order/draft/update"
	APIOrderDraftDelete = "/paymentapi/order/draft/delete"
	APIOrderDraftSend   = "/paymentapi/order/draft/send"
	APIGetOrderTariff   = "/paymentapi/tariff/order/"
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
	UID string `json:"uid"`
}

type UpdateOrderStatusRequest struct {
	UID    string `json:"uid"`
	Status string `json:"status"`
}

type ExternalPayout struct {
	Method         PaymentMethod `json:"paymentMethod"`
	BankCarsUid    string        `json:"bankCardUid"`
	BankAccountUid string        `json:"bankAccountUid"`
	WalletUid      string        `json:"walletUid"`
}

type PaymentMethodTariff struct {
	Code                       string        `json:"code"`
	Name                       string        `json:"name"`
	Method                     PaymentMethod `json:"paymentMethod"`
	Commission                 float64       `json:"commission"`
	CommisionPercent           float64       `json:"commissionPercent"`
	CommissionAmountAdditional float64       `json:"commissionAmountAdditional"`
	AmountSent                 float64       `json:"amountSent"`
	AmountReceived             float64       `json:"amountReceived"`
}

type Currency struct {
	Code        string `json:"code"`
	Symbol      string `json:"symbol"`
	Description string `json:"description"`
}

type Wallet struct {
	Uid      string     `json:"uid"`
	Type     WalletType `json:"type"`
	Currency Currency   `json:"type"`
	Balance  float64    `json:"balance"`
	Primary  bool       `json:"primary"`
}

type GetOrderTariffResponse struct {
	ErrorData           *ErrorData          `json:"errorData,omitempty"`
	OrderUid            string              `json:"orderUid,omitempty"`
	OriginalOrderAmount float64             `json:"originalOrderAmount,omitempty"`
	PaymentMethods      PaymentMethodTariff `json:"paymentMethods,omitempty"`
}
