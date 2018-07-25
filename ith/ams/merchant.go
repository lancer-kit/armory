package ams

import "github.com/go-ozzo/ozzo-validation"

// swagger:model
type MerchantRequest struct {

	// OAuth client ID
	//
	// required: true
	// example: "vipcoin"
	// min length: 3
	ClientId string `json:"clientId"`

	// OAuth client secret
	//
	// required: true
	// example: "vipcoinpass"
	// min length: 9
	ClientSecret string `json:"clientSecret"`

	// ITH account ID
	//
	// required: true
	// example: "100-020-425-40"
	AccountUid string `json:"accountUid"`

	// ITH access token, required when create
	//
	// required: false
	// example: bdad264b7f8b9896d73436b234e4bddd
	AccessToken string `json:"accessToken,omitempty"`

	// Password for vipcoin (plain). Required on registration request
	//
	// required: false
	// min length: 8
	// example: dad26-8be4!
	Password string `json:"password,omitempty"`

	// ITH Account model
	//
	// required: true
	Account *Account `json:"account"`

	// Internal, for validation
	isCreateRequest bool `json:"-"`
}

// swagger:model
type MerchantResponse struct {

	// Vipcoin user ID.Empty when error.
	//
	// required: false
	// example: MER-123
	ExternalAccountUid string `json:"externalAccountUid"`

	// Error data. Null if OK (code 200)
	//
	// required: false
	ErrorData *ErrorData `json:"errorData,omitempty"`
}

func (t *MerchantRequest) Validate() (err error) {
	rules := []*validation.FieldRules{
		validation.Field(&t.Account, validation.Required),
		validation.Field(&t.AccountUid, validation.Required),
		validation.Field(&t.ClientId, validation.Required),
		validation.Field(&t.ClientSecret, validation.Required),
	}
	if t.isCreateRequest {
		rules = append(rules,
			validation.Field(&t.Password, validation.Required, validation.Length(8, 128)),
			validation.Field(&t.AccessToken, validation.Required))
	}

	return validation.ValidateStruct(t, rules...)
}
