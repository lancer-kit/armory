package ith

import "net/url"

type Config struct {
	BaseURL url.URL
	Path    string
}

const (
	APIAuthToken        = "/paymentapi/auth/token"
	APIAuthTokenRefresh = "/paymentapi/auth/token/refresh"
	APIOrderCreate      = "/paymentapi/order/create"
	APIOrderData        = "/paymentapi/order/details"
	APIOrdersList       = "/paymentapi/order/list"
	APIOrderRefund      = "/paymentapi/order/refund"
)
