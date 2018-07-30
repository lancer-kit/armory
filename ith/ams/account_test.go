package ams

import (
	"testing"
	"encoding/json"
	"github.com/stretchr/testify/assert"
)

//	{
//		"id": 123,
//		"code": "LV",
//		"name": "Latvia",
//		"brandedCardsAvailable": true,
//		"registrationAllowed": true
//	}
func testCountry(t *testing.T, c *Country) {
	assert.Equal(t, "LV", c.Code)
	assert.Equal(t, "Latvia", c.Name)
	assert.Equal(t, true, c.BrandedCardsAvailable)
	assert.Equal(t, true, c.RegistrationAllowed)
	if !t.Failed() {
		println("Ok - testCountry", )
	}
}

func testLanguage(t *testing.T, c *Language) {
	assert.Equal(t, "ru", c.Code)
	if !t.Failed() {
		println("Ok - testLanguage", )
	}
}

func testCommunicationLanguage(t *testing.T, c *Language) {
	assert.Equal(t, "en", c.Code)
	if !t.Failed() {
		println("Ok - testCommunicationLanguage", )
	}
}

func testPhone(t *testing.T, p *AccountPhone) {
	assert.Equal(t, "371", p.CountryCode)
	assert.Equal(t, "20000010", p.Number)
	assert.Equal(t, "M", p.Type.String())
	assert.Equal(t, false, p.ContactPreference)
	assert.Equal(t, true, p.Primary)
	assert.Equal(t, true, p.Confirmed)
	if !t.Failed() {
		println("Ok - testPhone", )
	}
}

//"accountSettings": [
//  {
//   "name": "externalAccountUid",
//   "value": "EX-ACC-UID-1234",
//   "category": "default_category"
//  }
//],
func testAccountSettings(t *testing.T, a AccountSettings) {
	if !assert.Equal(t, 1, len(a)) {
		assert.Fail(t, "empty AccountSettings")
		return
	}
	s := a[0]
	assert.Equal(t, "externalAccountUid", s.Name)
	assert.Equal(t, "EX-ACC-UID-1234", s.Value)
	assert.Equal(t, "default_category", s.Category)
	if !t.Failed() {
		println("Ok - testAccountSettings", )
	}
}

//"accountEmails": [
//	{
//	"uid": "5340c5dd139528019a01821ba8bc7f09",
//	"email": "john@enauda.com",
//	"confirmed": false,
//	"primary": true
//	}
//],
func testAccountEmails(t *testing.T, a AccountEmails) {
	if !assert.Equal(t, 1, len(a)) {
		assert.Fail(t, "empty AccountEmails")
		return
	}
	s := a[0]
	assert.Equal(t, "5340c5dd139528019a01821ba8bc7f09", s.Uid)
	assert.Equal(t, "john@enauda.com", s.Email)
	assert.Equal(t, false, s.Confirmed)
	assert.Equal(t, true, s.Primary)
	if !t.Failed() {
		println("Ok - testAccountEmails", )
	}
}

//"addresses": [
//	{
//	"uid": "ac825e3bf86d71aa8d6848832031704a",
//	"country": {
//		"id": 123,
//		"code": "LV",
//		"name": "Latvia",
//		"brandedCardsAvailable": true,
//		"registrationAllowed": true
//	},
//	"city": "Riga",
//	"firstAddressLine": "Duntes 4",
//	"secondAddressLine": "Office 403",
//	"state": "Rigas rajons",
//	"postalCode": "1234",
//	"type": "H",
//	"primary": true
//	}
//],
func testAddresses(t *testing.T, a Addresses) {
	if !assert.Equal(t, 1, len(a)) {
		assert.Fail(t, "empty AccountEmails")
		return
	}
	s := a[0]
	assert.Equal(t, "ac825e3bf86d71aa8d6848832031704a", s.Uid)
	assert.Equal(t, "Riga", s.City)
	assert.Equal(t, "Duntes 4", s.FirstAddressLine)
	assert.Equal(t, "Office 403", s.SecondAddressLine)
	assert.Equal(t, "Rigas rajons", s.State)
	assert.Equal(t, "1234", s.PostalCode)
	assert.Equal(t, "H", s.Type.String())
	assert.Equal(t, AddressTypeHome, s.Type)
	assert.Equal(t, true, s.Primary)
	testCountry(t, s.Country)
	if !t.Failed() {
		println("Ok - testAddresses", )
	}
}

//"person": {
//	"name": "John",
//	"surname": "Doe",
//	"birthDate": "19810509000000",
//	"pep": false
//},
func testPerson(t *testing.T, s *Person) {
	assert.Equal(t, "John", s.Name)
	assert.Equal(t, "Doe", s.Surname)
	assert.Equal(t, "19810509000000", s.BirthDate.String())
	assert.Equal(t, false, s.Pep)
	if !t.Failed() {
		println("Ok - testPerson", )
	}
}

func TestAccountUnMarshalJSON(t *testing.T) {
	var b AccountResponse
	e := json.Unmarshal([]byte(responseExample), &b)
	if !assert.NoError(t, e) {
		assert.Fail(t, "unable to unmarshal Account")
		return
	}
	if !assert.NotEmpty(t, b.Account) {
		assert.Fail(t, "unable to unmarshal AccountResponse")
		return
	}
	a := b.Account
	assert.Equal(t, "100-020-425-40", a.Uid)

	if assert.NotEmpty(t, a.Country) {
		testCountry(t, a.Country)
	}

	if assert.NotEmpty(t, a.Language) {
		testLanguage(t, a.Language)
	}

	if assert.NotEmpty(t, a.CommunicationLanguage) {
		testCommunicationLanguage(t, a.CommunicationLanguage)
	}

	assert.Equal(t, AccountTypeStandard, a.Type)
	assert.Equal(t, "S", a.Type.String())
	assert.NoError(t, a.Type.Validate())
	assert.Equal(t, StStandardRegistrationConfirmed, a.Status)
	assert.NoError(t, a.Status.Validate())
	assert.Equal(t, "SC", a.Status.String())
	if assert.Equal(t, 1, len(a.AccountPhones)) {
		testPhone(t, a.AccountPhones[0])
	}
	testAccountSettings(t, a.AccountSettings)
	testAccountEmails(t, a.AccountEmails)
	testAddresses(t, a.Addresses)
	testPerson(t, a.Person)
}

func TestCountry_UnmarshalJSON(t *testing.T) {
	c := `
{
	"code": "LV",
	"name": "Latvia",
	"brandedCardsAvailable": true,
	"registrationAllowed": true
}`
	var a Country
	e := json.Unmarshal([]byte(c), &a)
	if !assert.Equal(t, nil, e) {
		assert.Fail(t, "unable to unmarshal")
		return
	}
	testCountry(t, &a)
}
