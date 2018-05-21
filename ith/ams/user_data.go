// ITH integration.
// ACCOUNT MANAGEMENT SERVICES
package ams

type (

	// Use (countryCode and number) OR fullNumber
	// Example:
	//```json
	//	...
	//"phone": {
	//	"countryCode ": "371",
	//	"number": "20000000",
	//	"fullNumber": "37120000000",
	//}
	//...
	//```
		Phone struct {
		CountryCode *string `json:"countryCode,omitempty"` //Phone country code (optional, when fullNumber filled)
		Number      *string `json:"number,omitempty"`      //Phone number (optional, when fullNumber filled)
		FullNumber  *string `json:"fullNumber,omitempty"`  //Full phone number. Min length 5 (optional, when countryCode+number filled)
	}

	//Address type and address list
	//
	//Example:
	//
	//  {
	//  "addresses": [
	//		{
	//			"uid": "ce50bd5486c0eca0dff2e7d8f45132fd",
	//			"country": {
	//				"id": 123,
	//				"code": "LV",
	//				"name": "Latvia",
	//				"brandedCardsAvailable": true
	//			},
	//			"city": "Riga",
	//			"firstAddressLine": "Duntes 4",
	//			"secondAddressLine": "Office 403",
	//			"state": "Rigas rajons",
	//			"postalCode": "1234",
	//			"type": "H",
	//			"primary": true
	//		}
	//	]
	//  }
	//
	Address struct {
		Uid               string      `json:"uid,omitempty"`     //Address UID, Optional for @address
		Country           *Country    `json:"country,omitempty"` //Country object, Optional for @address
		City              string      `json:"city"`              //City, required
		FirstAddressLine  string      `json:"firstAddressLine"`  //First address line, required
		SecondAddressLine string      `json:"secondAddressLine"` //Second address line, optional
		State             string      `json:"state"`             //State, optional
		PostalCode        string      `json:"postalCode"`        //Postal code, required
		Type              AddressType `json:"type,omitempty"`    //Optional for @address (see AddressType)
		Primary           bool        `json:"primary,omitempty"` //Optional for @address
	}


	//Additional field affiliateInfo
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

	/**
	Additional field `country`
	"country": {
			"code": "LV",
			"name": "Latvia",
			"brandedCardsAvailable": true
		},
	*/
	Country struct {
		Id                    int64  `json:"id,omitempty"`
		Code                  string `json:"code"`
		Name                  string `json:"name"`
		BrandedCardsAvailable bool   `json:"brandedCardsAvailable,omitempty"`
		RegistrationAllowed   bool   `json:"registrationAllowed,omitempty"`
	}

	/**
	Additional field `language`
	"language": {
			"code": "ru",
			"name": "Russian"
		},
	 */
	Language struct {
		Uid     string `json:"uid,omitempty"`
		Code    string `json:"code"` //ISO2 language code
		Name    string `json:"name,omitempty"`
		Type    string `json:"type,omitempty"`
		Primary bool   `json:"primary,omitempty"`
	}

	/**
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
	 */
	AccountPhone struct {
		Uid               string    `json:"uid,omitempty"`         //Phone UID
		CountryCode       string    `json:"countryCode,omitempty"` //Phone country code
		Number            string    `json:"number,omitempty"`      //Phone number
		Type              PhoneType `json:"type,omitempty"`        //Phone type (see: PhoneType)
		ContactPreference bool      `json:"contactPreference"`     //Phone is preferred for communication
		Primary           bool      `json:"primary,omitempty"`     //Phone is primary
		Confirmed         bool      `json:"confirmed,omitempty"`   //Phone is confirmed by account holder
	}

	//List of account settings
	//  "accountSettings": [
	//		{
	//			"name": "externalAccountUid",
	//			"value": "EX-ACC-UID-1234"
	//			"category": "ACC"
	//		}
	//	],
	AccountSetting struct {
		Name     string `json:"name"`               //Setting name
		Value    string `json:"value,omitempty"`    //Setting value
		Category string `json:"category,omitempty"` //Setting category
	}

	/**
	"accountEmails": [
			{
				"uid": "5340c5dd139528019a01821ba8bc7f09",
				"email": "john@enauda.com",
				"confirmed": false,
				"primary": true
			}
		],
	 */
	AccountEmail struct {
		Uid       string `json:"uid,omitempty"`
		Email     string `json:"email"`
		Confirmed bool   `json:"confirmed,omitempty"`
		Primary   bool   `json:"primary,omitempty"`
		Type      string `json:"type,omitempty"`
	}

	/**
	Person object
		   "person": {
			   "name": "John",
			   "surname": "Doe",
			   "birthDate": "19810509000000",
	           "pep": false
		   },

	*/
	Person struct {
		Name      string `json:"name"`      //Name
		Surname   string `json:"surname"`   //Surname
		BirthDate string `json:"birthDate"` //Date of birth,Format – yyyyMMddHHmmss
		Pep       bool   `json:"pep"`       //Person in politically exposed person (PEP)
	}

	/**
	{
		"accountUid": "100-014-275-55",
		"externalAccountUid": "EX-ACC-UID-1234",
		"accessToken": "bdad264b7f8b9896d73436b234e4bddd",
		"account": {....}
	}
	*/
	UserRegistrationResponse struct {
		AccountUid         string           `json:"accountUid"`
		ExternalAccountUid string           `json:"externalAccountUid"`
		AccessToken        string           `json:"accessToken"`
		Account            *StandardAccount `json:"account"`
	}
)
