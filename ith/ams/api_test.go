package ams

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	testRequest = UserRegistrationRequest{
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
	reqToken = AuthCodeRequest{AccessToken: "8392745c3c3055e68075b0aed98be118"}
)

func TestAPI_CreateProfile(t *testing.T) {
	api := NewAPI("http://demo-api.enauda.com/", "", "vipcoin", "vipcoinpass")
	resp, err, st := api.CreateProfile(&testRequest)
	if !assert.NoError(t, err) {
		return
	}
	assert.NotEmpty(t, resp)
	assert.Equal(t, RequestStatusOk, st)
}

func TestAPI_GetToken(t *testing.T) {
	api := NewAPI("http://demo-api.enauda.com/", "http://demo-commonapi.enauda.com", "vipcoin", "vipcoinpass")
	r, e := api.GetToken(&reqToken)
	assert.NoError(t, e)
	if assert.NotEmpty(t, r) {
		println("AccessToken:", r.AccessToken)
		println("RefreshToken:", r.RefreshToken)
	}
	if reqToken.AccessToken == r.AccessToken {
		api.log.Warning("Access token is identical to received")
	}

}

func TestAPI_UpdateProfile(t *testing.T) {
	api := NewAPI("http://demo-api.enauda.com/", "http://demo-commonapi.enauda.com", "vipcoin", "vipcoinpass")
	resp, err, st := api.UpdateProfile(&UserUpdateRequest{
		ExternalAccountUid: "STD-6",
		FirstName:          "Eugene",
	}, reqToken.AccessToken)
	if !assert.NoError(t, err) {
		return
	}
	assert.NotEmpty(t, api)
	assert.NotEmpty(t, resp)
	assert.Equal(t, RequestStatusOk, st)
	api.log.Info(resp)

}

func TestAPI_GetFullProfile(t *testing.T) {
	api := NewAPI("http://demo-api.enauda.com/", "http://demo-commonapi.enauda.com", "vipcoin", "vipcoinpass")
	resp, err, st := api.GetFullProfile("b575be6b56ec7506f15e0429dc92b436")
	if !assert.NoError(t, err) {
		assert.Fail(t, "unable to get profile")
		return
	}
	assert.NotEmpty(t, resp)
	assert.Equal(t, RequestStatusOk, st)
	assert.Equal(t, "100-001-319-33", resp.Uid)
}
