package models

type (
	/**
	 Use (countryCode and number) OR fullNumber
	 Example:
	```json
		...
	"phone": {
		"countryCode ": "371",
		"number": "20000000",
		"fullNumber": "37120000000",
	}
	...
	```
	*/
	Phone struct {
		CountryCode *string `json:"countryCode,omitempty"` //Phone country code (optional, when fullNumber filled)
		Number      *string `json:"number,omitempty"`      //Phone number (optional, when fullNumber filled)
		FullNumber  *string `json:"fullNumber,omitempty"`  //Full phone number. Min length 5 (optional, when countryCode+number filled)
	}

	/**
	Example:
	```json
	"address": {
		"countryCode": "LV",
		"city": "Riga",
		"firstAddressLine": "Duntes 4",
		"secondAddressLine": "",
		"state": "Riga",
		"postalCode": "LV-1010"
	}
	```
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
	 */
	Address struct {
		Uid               string   `json:"uid,omitempty"`         //Optional for @address
		CountryCode       string   `json:"countryCode,omitempty"` //ISO2 country code, required
		Country           *Country `json:"country,omitempty"`     //Optional for @address
		City              string   `json:"city"`                  //City, required
		FirstAddressLine  string   `json:"firstAddressLine"`      //First address line, required
		SecondAddressLine string   `json:"secondAddressLine"`     //Second address line, optional
		State             string   `json:"state"`                 //State, optional
		PostalCode        string   `json:"postalCode"`            //Postal code, required
		Type              string   `json:"type,omitempty"`        //Optional for @address
		Primary           bool     `json:"primary,omitempty"`     //Optional for @address
	}

	/**
	"account": {
		"email": "john@doe.com",
		"phone": {
			"countryCode": "371",
			"number": "20000000"
		},
		"externalUid": "1234567890",
		"name": "John",
		"surname": "Doe",
		"birthDate": "19951015000000",
		"language": "en",
		"address": {
			"countryCode": "LV",
			"city": "Riga",
			"firstAddressLine": "Duntes 4",
			"secondAddressLine": "Office 403",
			"state": "Rigas",
			"postalCode": "LV-2020"
		},
		"website": "http://wwww.doe.com",
		"taxpayerIdentificationNumber": "TAX83642",
		"additionalInfo": "Additional information for buyer"
	}
	 */
	Account struct {
		Email                        string   `json:"email,omitempty"`                        //Email, required
		Phone                        *Phone   `json:"phone,omitempty"`                        //Phone
		ExternalUid                  string   `json:"externalUid,omitempty"`                  //Account UID in external system
		Name                         string   `json:"name"`                                   //Name, required
		SureName                     string   `json:"surename"`                               //Surname, required
		BirthDate                    string   `json:"birthDate"`                              //Date of birth, Format - yyyyMMddHHmmss
		Language                     string   `json:"language,omitempty"`                     //ISO2 language code
		Address                      *Address `json:"address,omitempty"`                      //Address
		WebSite                      string   `json:"website,omitempty"`                      //Website
		TaxpayerIdentificationNumber string   `json:"taxpayerIdentificationNumber,omitempty"` //Taxpayer Identification Number
		AdditionalInfo               string   `json:"additionalInfo,omitempty"`               //Additional information
	}

	/*
	-----------------------------------------------------------------------
	Request body for POST https://<host>:<port>/partnerapi/account/register
	-----------------------------------------------------------------------
	{
		"clientId": "#partner-code#",
		"clientSecret": "#partner-pass#",
		"externalAccountUid": "EX-ACC-UID-1234",
		"phone": "37120000000",
		"email": "john@enauda.com",
		"password": "pAsSw0rD",
		"firstName": "John",
		"lastName": "Doe",
		"birthDate": "19810509000000",
		"country": "LV",
		"language": "ru",
		"address": {
			"city": "Riga",
			"countryCode": "LV",
			"firstAddressLine": "Duntes 4",
			"secondAddressLine": "Office 403",
			"postalCode": "1234",
			"state": "Rigas rajons"
		},
		"affiliateInfo": {
			"affiliateId": "AF4",
			"campaignId": "C539",
			"bannerId": "BRT13",
			"customParameters": "tr=24&hd=3"
		}
	}
	 */
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

	/**
	Additional field `affiliateInfo`
	"affiliateInfo":
		{
			"affiliateId": "AF4",
			"campaignId": "C539",
			"bannerId": "BRT13",
			"customParameters": "tr=24&hd=3"
		}
	 */
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
		Code    string `json:"code"`
		Name    string `json:"name"`
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
		Uid         string `json:"uid,omitempty"`
		CountryCode string `json:"countryCode,omitempty"`
		Number      string `json:"number,omitempty"`
		Type        string `json:"type,omitempty"`
		Primary     bool   `json:"primary,omitempty"`
		Confirmed   bool   `json:"confirmed,omitempty"`
	}

	/**
	"accountSettings": [
			{
				"name": "externalAccountUid",
				"value": "EX-ACC-UID-1234"
			}
		],
	 */
	AccountSetting struct {
		Name  string `json:"name"`
		Value string `json:"value,omitempty"`
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

	/**Standard account
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

	 */

	Person struct {
		Name      string `json:"name"`
		Surname   string `json:"surname"`
		BirthDate string `json:"birthDate"`
	}
	StandardAccount struct {
		Uid                   string            `json:"uid"`
		Country               *Country          `json:"country,omitempty"`
		Language              *Language         `json:"language,omitempty"`
		Type                  string            `json:"type,omitempty"`
		Status                string            `json:"status,omitempty"`
		AccountPhones         []*AccountPhone   `json:"accountPhones"`
		AccountSettings       []*AccountSetting `json:"accountSettings"`
		AccountEmails         []*AccountEmail   `json:"accountEmails"`
		Addresses             []*Address        `json:"addresses"`
		Person                *Person           `json:"person"`
		AffiliateId           string            `json:"affiliateId"`
		CampaignId            string            `json:"campaignId"`
		BannerId              string            `json:"bannerId"`
		CustomParameters      string            `json:"customParameters"`
		CurrencyConversion    bool              `json:"currencyConversion"`
		AlwaysRefundEWallet   bool              `json:"alwaysRefundEWallet"`
		ConfirmOutTransaction bool              `json:"confirmOutTransaction"`
		Test                  bool              `json:"test"`
	}

	/**
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

	 */
	UserRegistrationResponse struct {
		AccountUid         string           `json:"accountUid"`
		ExternalAccountUid string           `json:"externalAccountUid"`
		AccessToken        string           `json:"accessToken"`
		Account            *StandardAccount `json:"account"`
	}
)
