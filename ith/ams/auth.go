package ams

import "errors"

type (
	AuthCodeRequest struct {
		AccessToken string `json:"accessToken,omitempty"` //user access token
		Username    string `json:"username,omitempty"`    //Account username (email)
		Password    string `json:"password,omitempty"`    //Account password (plain password)
	}

	AuthCodeResponse struct {
		ErrorData *ErrorData `json:"errorData,omitempty"` //Not returned if operation is successful
		Code      string     `json:"code,omitempty"`      //One-time authorization code
	}

	AuthTokenRequest struct {
		ClientId     string `json:"clientId"`     //OAuth client ID
		ClientSecret string `json:"clientSecret"` //OAuth client secret
		Code         string `json:"code"`         //One-time authorization code
	}

	AuthTokenResponse struct {
		ErrorData    *ErrorData `json:"errorData,omitempty"` //Not returned if operation is successful
		AccessToken  string     `json:"accessToken"`         //Access token for integration services
		RefreshToken string     `json:"refreshToken"`        //Refresh token for access token renewal
		ExpiresIn    int64      `json:"expiresIn"`           //Expiration time for access token (seconds)
	}
)

func (r *AuthCodeRequest) Validate() error {
	var err error
	if len(r.AccessToken) > 0 {
		return nil
	}

	if len(r.Username) == 0 || len(r.Password) == 0 {
		err = errors.New("accessToken or username+password is required")
	}

	return err
}
