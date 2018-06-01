package auth

import (
	"net/http"

	"gitlab.inn4science.com/vcg/go-common/ith"
)

const (
	APIAuthToken        = "/paymentapi/auth/token"
	APIAuthTokenRefresh = "/paymentapi/auth/token/refresh"
)

type ErrorData struct {
	ith.ErrorData
}

type Request struct {
	Username string `json:"username"` // String(150); Merchant’s username
	Password string `json:"password"` // String(50); Merchant’s password
}

type Response struct {
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

const Header = "Authorization"

var HeaderVal = func(accessToken string) (value string) {
	return "Bearer " + accessToken
}
