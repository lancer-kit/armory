package ams

import (
	"errors"
	"fmt"
)

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
	AddressRequest struct {
		CountryCode       string `json:"countryCode"`       //Required, String(2), ISO-2 country code
		City              string `json:"city"`              //Required, String(50), City
		FirstAddressLine  string `json:"firstAddressLine"`  //Required, String(60), First address line
		SecondAddressLine string `json:"secondAddressLine"` //Optional, String(60), Second address line
		PostalCode        string `json:"postalCode"`        //Required, String(10), Postal code
		State             string `json:"state"`             //Optional,String(50), State
	}

	//Additional field affiliateInfo  - used in "Create Standard Account" method of Account Management Services
	//
	//	"affiliateInfo":
	//		{
	//			"affiliateId": "AF4",
	//			"campaignId": "C539",
	//			"bannerId": "BRT13",
	//			"customParameters": "tr=24&hd=3"
	//		}
	AffiliateInfo struct {
		AffiliateId      string `json:"affiliateId"`      //Optional, String(50), Affiliate ID
		CampaignId       string `json:"campaignId"`       //Optional, String(50), Campaign ID
		BannerId         string `json:"bannerId"`         //Optional, String(50), Banner ID
		CustomParameters string `json:"customParameters"` //Optional, String(255), Custom parameters
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
	UserRegistrationRequest struct {
		ClientId           string          `json:"clientId"`                 //Required, String(50), OAuth client ID
		ClientSecret       string          `json:"clientSecret"`             //Required, String(32), OAuth client secret
		ExternalAccountUid string          `json:"externalAccountUid"`       //Required, String(50), User API user ID
		Phone              string          `json:"phone"`                    //Required, String(30), Full phone number. Min length 5
		Email              string          `json:"email"`                    //Optional, String(150), Email
		Password           string          `json:"password"`                 //Required, String(50), User account password (plain ?)
		FirstName          string          `json:"firstName"`                //Required, String(50), Name
		LastName           string          `json:"lastName"`                 //Required, String(50), Surname
		BirthDate          AmsDate         `json:"birthDate"`                //Optional, Date, Format - yyyyMMddHHmmss
		Country            string          `json:"country"`                  //Required, String(2), ISO2 country code
		Language           string          `json:"language"`                 //Optional, String(2), ISO2 language code
		Address            *AddressRequest `json:"address"`                  //Required, User account address
		AffiliateInfo      *AffiliateInfo  `json:"affiliateInfo,ommitempty"` //Optional, Affiliate information
	}
)

// Validate verifies that value is predefined for AffiliateInfo.
func (r *UserRegistrationRequest) Validate() (err error) {
	msg := ""
	if r.Address == nil {
		msg += "address: required;"
	} else {
		if tmp := r.Address.Validate(); tmp != nil {
			msg += "address: (" + tmp.Error() + ");"
		}
	}
	if r.AffiliateInfo != nil {
		if tmp := r.AffiliateInfo.Validate(); tmp != nil {
			msg += "affiliateInfo: (" + tmp.Error() + ");"
		}
	}
	if !r.BirthDate.Empty() {
		if tmp := r.BirthDate.Validate(); tmp != nil {
			msg += "birthDate: (" + tmp.Error() + ");"
		}
	}
	msg += isRequired(r.ClientId, "clientId")
	msg += maxLen(r.ClientId, "clientId", 50)
	msg += isRequired(r.ClientSecret, "clientSecret")
	msg += maxLen(r.ClientSecret, "clientSecret", 32)
	msg += isRequired(r.ExternalAccountUid, "externalAccountUid")
	msg += maxLen(r.ExternalAccountUid, "externalAccountUid", 50)
	msg += isRequired(r.Phone, "phone")
	msg += maxLen(r.Phone, "phone", 30)
	msg += minLen(r.Phone, "phone", 5)
	msg += isRequired(r.Password, "password")
	msg += maxLen(r.Password, "password", 50)
	msg += minLen(r.Password, "password", 6)
	msg += isRequired(r.FirstName, "firstName")
	msg += maxLen(r.FirstName, "firstName", 50)
	msg += isRequired(r.LastName, "lastName")
	msg += maxLen(r.LastName, "lastName", 50)
	msg += isRequired(r.Country, "country")
	if len(r.Country) != 2 {
		msg += "countryCode: invalid country code;"
	}
	if msg != "" {
		err = errors.New(msg)
	}
	return
}

// Validate verifies that value is predefined for AddressRequest.
func (r *AddressRequest) Validate() (err error) {
	msg := ""

	if len(r.CountryCode) != 2 {
		msg += "countryCode: invalid country code;"
	}

	msg += isRequired(r.City, "city")
	msg += maxLen(r.City, "city", 50)
	msg += minLen(r.City, "city", 3)
	msg += isRequired(r.FirstAddressLine, "firstAddressLine")
	msg += maxLen(r.FirstAddressLine, "firstAddressLine", 60)
	msg += maxLen(r.SecondAddressLine, "secondAddressLine", 60)
	msg += isRequired(r.PostalCode, "postalCode")
	msg += maxLen(r.PostalCode, "postalCode", 10)
	msg += maxLen(r.State, "state", 50)

	if msg != "" {
		err = errors.New(msg)
	}

	return
}

// Validate verifies that value is predefined for AffiliateInfo.
func (r *AffiliateInfo) Validate() (err error) {
	msg := ""
	msg += maxLen(r.AffiliateId, "affiliateId", 50)
	msg += maxLen(r.CampaignId, "campaignId", 50)
	msg += maxLen(r.BannerId, "bannerId", 50)
	msg += maxLen(r.CustomParameters, "bannerId", 255)

	if msg != "" {
		err = errors.New(msg)
	}

	return
}

//Check when data not empty
func isRequired(v string, name string) (s string) {
	if len(v) == 0 {
		s = name + " :required;"
	}
	return
}

//Check data max length
func maxLen(v string, name string, max int) (s string) {
	if len(v) > max {
		s = name + ": max " + fmt.Sprint(max) + " char.;"
	}
	return
}

//Check data min length
func minLen(v string, name string, min int) (s string) {
	if len(v) < min {
		s = name + ": min " + fmt.Sprint(min) + " char.;"
	}
	return
}
