package ams

import "github.com/go-ozzo/ozzo-validation"

// swagger:model
type (
	MerchantRequest struct {

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

		// Callback URL required: true
		// example: https://<host>:<port>/callback
		CallbackUrl string `json:"callbackUrl"`

		// Internal, for validation
		// example: false
		IsCreateRequest bool `json:"-"`
	}
)

func (t *MerchantRequest) Validate() (err error) {
	rules := []*validation.FieldRules{
		validation.Field(&t.AccountUid, validation.Required),
		validation.Field(&t.CallbackUrl, validation.Required),
	}
	if t.AccessToken != "" {
		t.IsCreateRequest = true
	}
	return validation.ValidateStruct(t, rules...)
}
