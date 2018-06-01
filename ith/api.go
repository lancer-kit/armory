package ith

import "net/url"

type ErrorData struct {
	ErrorCode    int    `json:"errorCode"`
	ErrorMessage string `json:"errorMessage"`
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
	Credentials struct {
		AccessToken  string
		RefreshToken string
		SetAt        int64
		ExpiresIn    int64
	}
}
