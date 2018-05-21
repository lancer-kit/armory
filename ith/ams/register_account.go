package ams

type (
	// Request body for POST https://<host>:<port>/partnerapi/account/register
	//
	// Example:
	//
	//  {
	//	"clientId": "#partner-code#",
	//	"clientSecret": "#partner-pass#",
	//	"externalAccountUid": "EX-ACC-UID-1234",
	//	"phone": "37120000000",
	//	"email": "john@enauda.com",
	//	"password": "pAsSw0rD",
	//	"firstName": "John",
	//	"lastName": "Doe",
	//	"birthDate": "19810509000000",
	//	"country": "LV",
	//	"language": "ru",
	//	"address": {
	//		"city": "Riga",
	//		"countryCode": "LV",
	//		"firstAddressLine": "Duntes 4",
	//		"secondAddressLine": "Office 403",
	//		"postalCode": "1234",
	//		"state": "Rigas rajons"
	//	},
	//	"affiliateInfo": {
	//		"affiliateId": "AF4",
	//		"campaignId": "C539",
	//		"bannerId": "BRT13",
	//		"customParameters": "tr=24&hd=3"
	//	}
	//  }
	UserRegistrationRequest struct {
		ClientId           string `json:"clientId"`           //Required, String(50), OAuth client ID
		ClientSecret       string `json:"clientSecret"`       //Required, String(32), OAuth client secret
		ExternalAccountUid string `json:"externalAccountUid"` //Required, String(50), User API user ID
		Phone              string `json:"phone"`              //Required, String(30), Full phone number. Min length 5
		Email              string `json:"email"`              //Optional, String(150), Email
		Password           string `json:"password"`           //Required, String(50), User account password (plain ?)
		FirstName          string `json:"firstName"`          //Required, String(50), Name
		LastName           string `json:"lastName"`           //Required, String(50), Surname
		BirthDate          string `json:"birthDate"`          //Optional, Date, Format - yyyyMMddHHmmss
		Country            string `json:"country"`            //Required, String(2), ISO2 country code
		Language           string `json:"language"`           //Optional, String(2), ISO2 language code
		Address            *Address                           //Required, User account address
		AffiliateInfo      *AffiliateInfo                     //Optional, Affiliate information
	}
)