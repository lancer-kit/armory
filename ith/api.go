package ith

import (
	"net/url"
	"time"
)

type ErrorData struct {
	ErrorCode    int    `json:"errorCode"`
	ErrorMessage string `json:"errorMessage"`
	RequestUid   string `json:"requestUid,omitempty"` //Request UID, used for investigation of exceptional cases
}

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
	Credentials Credentials
}

type Credentials struct {
	AccessToken  string `json:"accessToken"`  // String(50); Access token for integration services.
	RefreshToken string `json:"refreshToken"` // String(50); Refresh token for access token renewal.
	SetAt        int64  `json:"setAt"`        // Time of the last token renewal (unix timestamp).
	ExpiresIn    int64  `json:"expiresIn"`    // Expiration time for access token (seconds).
}

// Expired returns true if the token expires after 1 minute.
func (c Credentials) Expired() bool {
	const delta int64 = 60

	return time.Now().UTC().Unix() > c.SetAt+c.ExpiresIn+delta
}
