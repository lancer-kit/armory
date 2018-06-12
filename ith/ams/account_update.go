package ams

import "errors"

type (
	//AddressRequest - used in "Create Standard Account" method of Account Management Services
	//
	//Example:
	//  {
	//		"city": "Riga",
	//		"countryCode": "LV",
	//		"firstAddressLine": "Duntes 4",
	//		"secondAddressLine": "Office 403",
	//		"postalCode": "1234",
	//		"state": "Rigas rajons"
	//	}
	AddressUpdate struct {
		CountryCode       string `json:"countryCode,omitempty"`       //Optional/fill if updated, String(2), ISO-2 country code
		City              string `json:"city,omitempty"`              //Optional/fill if updated, String(50), City
		FirstAddressLine  string `json:"firstAddressLine,omitempty"`  //Optional/fill if updated, String(60), First address line
		SecondAddressLine string `json:"secondAddressLine,omitempty"` //Optional/fill if updated, String(60), Second address line
		PostalCode        string `json:"postalCode,omitempty"`        //Optional/fill if updated, String(10), Postal code
		State             string `json:"state,omitempty"`             //Optional/fill if updated,String(50), State
	}

	// Request body for POST https://<host>:<port>/partnerapi/account/register
	//
	// Example:
	//
	//  {
	//		"clientId": "#partner-code#",
	//		"clientSecret": "#partner-pass#",
	//		"externalAccountUid": "EX-ACC-UID-1234",
	//		"phone": "37120000000",
	//		"email": "john@enauda.com",
	//		"password": "pAsSw0rD",
	//		"firstName": "John",
	//		"lastName": "Doe",
	//		"birthDate": "19810509000000",
	//		"country": "LV",
	//		"language": "ru",
	//		"address": {
	//			"city": "Riga",
	//			"countryCode": "LV",
	//			"firstAddressLine": "Duntes 4",
	//			"secondAddressLine": "Office 403",
	//			"postalCode": "1234",
	//			"state": "Rigas rajons"
	//		},
	//		"affiliateInfo": {
	//			"affiliateId": "AF4",
	//			"campaignId": "C539",
	//			"bannerId": "BRT13",
	//			"customParameters": "tr=24&hd=3"
	//		}
	//  }
	UserUpdateRequest struct {
		ClientId           string         `json:"clientId"`                 //Required, String(50), OAuth client ID
		ClientSecret       string         `json:"clientSecret"`             //Required, String(32), OAuth client secret
		ExternalAccountUid string         `json:"externalAccountUid"`       //Required, String(50), User API user ID
		Phone              string         `json:"phone,omitempty"`          //Optional/fill if updated, String(30), Full phone number. Min length 5
		Email              string         `json:"email,omitempty"`          //Optional/fill if updated, String(150), Email
		Password           string         `json:"password,omitempty"`       //Optional/fill if updated, String(50), User account password (plain ?)
		FirstName          string         `json:"firstName,omitempty"`      //Optional/fill if updated, String(50), Name
		LastName           string         `json:"lastName,omitempty"`       //Optional/fill if updated, String(50), Surname
		BirthDate          AmsDate        `json:"birthDate,omitempty"`      //Optional/fill if updated, Date, Format - yyyyMMddHHmmss
		Country            string         `json:"country,omitempty"`        //Optional/fill if updated, String(2), ISO2 country code
		Language           string         `json:"language,omitempty"`       //Optional/fill if updated, String(2), ISO2 language code
		Address            *AddressUpdate `json:"address,omitempty"`        //Optional/fill if updated, User account address
		AffiliateInfo      *AffiliateInfo `json:"affiliateInfo,ommitempty"` //Optional, Affiliate information
	}
)

func (r UserUpdateRequest) Validate() (err error) {
	msg := ""
	msg += isRequired(r.ClientId, "clientId")
	msg += maxLen(r.ClientId, "clientId", 50)
	msg += isRequired(r.ClientSecret, "clientSecret")
	msg += maxLen(r.ClientSecret, "clientSecret", 32)
	msg += isRequired(r.ExternalAccountUid, "externalAccountUid")
	msg += maxLen(r.ExternalAccountUid, "externalAccountUid", 50)
	if msg != "" {
		err = errors.New(msg)
	}
	return
}
