package ams

import (
	"testing"
	"encoding/json"
	"github.com/stretchr/testify/assert"
)

var responseExample = `
{
		"accountUid": "100-014-275-55",
		"externalAccountUid": "EX-ACC-UID-1234",
		"accessToken": "bdad264b7f8b9896d73436b234e4bddd",
		"account": {
			"uid": "100-014-275-55",
			"country": {
				"code": "LV",
				"name": "Latvia",
				"brandedCardsAvailable": true
			},
			"language": {
				"code": "ru",
				"name": "Russian"
			},
			"type": "S",
			"status": "SR",
			"accountPhones": [
				{
					"uid": "3e22eb32b25e7e3b6c7261e3d2d2654c",
					"countryCode": "371",
					"number": "20000000",
					"type": "M",
					"primary": true,
					"confirmed": true
				}
			],
			"accountSettings": [
				{
					"name": "externalAccountUid",
					"value": "EX-ACC-UID-1234"
				}
			],
			"accountEmails": [
				{
					"uid": "5340c5dd139528019a01821ba8bc7f09",
					"email": "john@enauda.com",
					"confirmed": false,
					"primary": true
				}
			],
			"addresses": [
				{
					"uid": "ce50bd5486c0eca0dff2e7d8f45132fd",
					"country": {
						"id": 123,
						"code": "LV",
						"name": "Latvia",
						"brandedCardsAvailable": true
					},
					"city": "Riga",
					"firstAddressLine": "Duntes 4",
					"secondAddressLine": "Office 403",
					"state": "Rigas rajons",
					"postalCode": "1234",
					"type": "H",
					"primary": true
				}
			],
			"person": {
				"name": "John",
				"surname": "Doe",
				"birthDate": "19810509000000"
			},
			"affiliateId": "AF4",
			"campaignId": "C539",
			"bannerId": "BRT13",
			"customParameters": "tr=24&hd=3",
			"currencyConversion": false,
			"alwaysRefundEWallet": false,
			"confirmOutTransaction": false,
			"test": false
		}
	}
`

func TestMarshallAddress(t *testing.T) {
	a := Address{
		//	CountryCode:       "LV",
		Country: &Country{
			Code:                  "LV",
			Name:                  "Latvia",
			BrandedCardsAvailable: true,
			RegistrationAllowed:   true,
		},
		City:              "Riga",
		FirstAddressLine:  "Duntes 4",
		SecondAddressLine: "Office 403",
		State:             "Rigas rajons",
		PostalCode:        "1234",
		Type:              "H",
		Primary:           true,
	}
	b, e := json.Marshal(&a)
	assert.Equal(t, nil, e)
	println(string(b))
	assert.Contains(t, string(b), `"country":`)
}

func TestParseResponse(t *testing.T) {
	var a UserRegistrationResponse
	e := json.Unmarshal([]byte(responseExample), &a)
	assert.Equal(t, nil, e)
	assert.Equal(t, 1, len(a.Account.AccountEmails))
	assert.Equal(t, 1, len(a.Account.AccountPhones))
	assert.Equal(t, 1, len(a.Account.Addresses))
	assert.Equal(t, "Doe", a.Account.Person.Surname)
}
