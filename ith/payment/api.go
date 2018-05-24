package payment

import (
	"net/http"

	"gitlab.inn4science.com/vcg/go-common/types/currency"
)

const (
	APIAuthToken        = "/paymentapi/auth/token"
	APIAuthTokenRefresh = "/paymentapi/auth/token/refresh"
	APIOrderCreate      = "/paymentapi/order/create"
	APIOrderData        = "/paymentapi/order/details"
	APIOrdersList       = "/paymentapi/order/list"
	APIOrderRefund      = "/paymentapi/order/refund"
)

type ErrorData struct {
	ErrorCode    int    `json:"errorCode"`
	ErrorMessage string `json:"errorMessage"`
}

type AuthRequest struct {
	Username string `json:"username"` // String(150); Merchant’s username
	Password string `json:"password"` // String(50); Merchant’s password
}

type AuthResponse struct {
	ErrorData    *ErrorData `json:"errorData,omitempty"`    // Is null for successful operation
	AccessToken  string     `json:"accessToken,omitempty"`  // String(50); Access token for integration services
	RefreshToken string     `json:"refreshToken,omitempty"` // String(50); Refresh token for access token renewal
	ExpiresIn    int64      `json:"expiresIn,omitempty"`    // Expiration time for access token (seconds). 3600s by default (1h)
}

type RefreshRequest struct {
	RefreshToken string `json:"refreshToken"` // String(50); Refresh token for access token renewal
}

func SetAuth(r *http.Request, accessToken string) *http.Request {
	r.Header.Set("Authorization", "Bearer "+accessToken)
	return r
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

const AuthHeader = "Authorization"

var AuthHeaderVal = func(accessToken string) (value string) {
	return "Bearer " + accessToken
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
