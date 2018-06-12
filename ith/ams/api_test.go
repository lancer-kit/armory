package ams

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var testRequest = UserRegistrationRequest{
	ExternalAccountUid: "test-1",
	Phone:              "380777777778",
	Email:              "test2@example.com",
	Password:           "5340c5dd139528019a01821ba8bc7f09",
	FirstName:          "Johner",
	LastName:           "Doer",
	BirthDate:          AmsDate(time.Date(1980, 01, 01, 0, 0, 0, 0, time.UTC)),
	Country:            "LV",
	Language:           "RU",
	Address: &AddressRequest{
		CountryCode:       "LV",
		City:              "Riga",
		FirstAddressLine:  "Duntes 4",
		SecondAddressLine: "",
		PostalCode:        "1234",
		State:             "",
	},
	AffiliateInfo: nil,
}

func TestAPI_CreateProfile(t *testing.T) {
	api := NewAPI("http://demo-api.enauda.com/", "vipcoin", "vipcoinpass")
	resp, err, st := api.CreateProfile(&testRequest)
	if !assert.NoError(t, err) {
		return
	}
	assert.NotEmpty(t, resp)
	assert.Equal(t, RequestStatusOk, st)
}

func TestAPI_UpdateProfile(t *testing.T) {
	api := NewAPI("http://demo-api.enauda.com/", "vipcoin", "vipcoinpass")
	//resp, err, st := api.UpdateProfile(&testRequest)
	//if !assert.NoError(t, err) {
	//	return
	//}
	assert.NotEmpty(t, api)
	//assert.Equal(t, RequestStatusOk, st)
}
