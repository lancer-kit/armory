package ams

import (
	"testing"
	"encoding/json"
	"github.com/stretchr/testify/assert"
)

var responseExample = `
{
  "accountUid": "100-020-425-40",
  "externalAccountUid": "EX-ACC-UID-1234",
  "account": {
    "uid": "100-020-425-40",
    "country": {
      "code": "LV",
      "name": "Latvia",
      "brandedCardsAvailable": true,
      "registrationAllowed": true
    },
    "language": {
      "code": "ru"
    },
    "communicationLanguage": {
      "code": "en"
    },
    "type": "S",
    "status": "SC",
    "accountPhones": [
      {
        "uid": "bd527410d2b3b9146a2ed72a30ba06e3",
        "countryCode": "371",
        "number": "20000010",
        "type": "M",
        "contactPreference": false,
        "primary": true,
        "confirmed": true
      }
    ],
    "accountSettings": [
      {
        "name": "externalAccountUid",
        "value": "EX-ACC-UID-1234",
        "category": "default_category"
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
        "uid": "ac825e3bf86d71aa8d6848832031704a",
        "country": {
          "id": 123,
          "code": "LV",
          "name": "Latvia",
          "brandedCardsAvailable": true,
          "registrationAllowed": true
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
      "birthDate": "19810509000000",
      "pep": false
    },
    "affiliateId": "AF4",
    "campaignId": "C539",
    "bannerId": "BRT13",
    "customParameters": "tr=24&hd=3",
    "timezone": 16,
    "weekStartsOn": "MO",
    "currencyConversion": true,
    "alwaysRefundEWallet": false,
    "confirmOutTransaction": false,
    "confirmLogin": false,
    "actionConfirmationEnabled": false,
    "test": false
  }
}
`

func TestParseResponse(t *testing.T) {
	var a UserRegistrationResponse
	e := json.Unmarshal([]byte(responseExample), &a)
	assert.Equal(t, nil, e)
	assert.Equal(t, 1, len(a.Account.AccountEmails))
	assert.Equal(t, 1, len(a.Account.AccountPhones))
	assert.Equal(t, 1, len(a.Account.Addresses))
	assert.Equal(t, "Doe", a.Account.Person.Surname)
	testCountry(t, a.Account.Country)
	testLanguage(t, a.Account.Language)
	testAccountEmails(t, a.Account.AccountEmails)
}
