// ITH integration.
// ACCOUNT MANAGEMENT SERVICES
package ams

type (
	//Use (countryCode and number) OR fullNumber
	//
	//Example:
	//	...
	//  "phone": {
	//	  "countryCode ": "371",
	//	  "number": "20000000",
	//	  "fullNumber": "37120000000",
	//   }
	//  ...
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
	Addresses []*Address

	//Address type, item of Addresses list
	Address struct {
		Id        int64 `json:"id,omitempty" db:"id"`                //user-integration data fields
		UserId    int64 `json:"userId,omitempty" db:"user_id"`       //user-integration data fields
		CountryId int64 `json:"countryId,omitempty" db:"country_id"` //user-integration data fields
		//Ams data structure
		Uid               string      `json:"uid,omitempty" db:"uid"`                     //Address UID, Optional for @address
		Country           *Country    `json:"country,omitempty" db:"-"`                   //Country object, Optional for @address
		City              string      `json:"city" db:"city"`                             //City, required
		FirstAddressLine  string      `json:"firstAddressLine" db:"first_address_line"`   //First address line, required
		SecondAddressLine string      `json:"secondAddressLine" db:"second_address_line"` //Second address line, optional
		State             string      `json:"state" db:"state"`                           //State, optional
		PostalCode        string      `json:"postalCode" db:"postal_code"`                //Postal code, required
		Type              AddressType `json:"type,omitempty" db:"type"`                   //Optional for @address (see AddressType)
		Primary           bool        `json:"primary,omitempty" db:"primary"`             //Optional for @address
	}

	//Additional field `country`
	//
	//   "country": {
	//		"code": "LV",
	//		"name": "Latvia",
	//		"brandedCardsAvailable": true
	//	},
	Country struct {
		Id int64 `json:"internal_id,omitempty" db:"id"` //Internal for user-integration
		//ITH.AMS data structure
		AmsId                 int64  `json:"id,omitempty" db:"ams_id"`
		Code                  string `json:"code" db:"code"`
		Name                  string `json:"name" db:"name"`
		BrandedCardsAvailable bool   `json:"brandedCardsAvailable,omitempty" db:"branded_cards_available"`
		RegistrationAllowed   bool   `json:"registrationAllowed,omitempty" db:"registration_allowed"`
	}

	//Additional field `language`
	//
	//  "language": {
	//		"code": "ru",
	//		"name": "Russian"
	//	},
	//
	Language struct {
		Uid     string `json:"uid,omitempty" db:"uid"`
		Code    string `json:"code" db:"code"` //ISO2 language code
		Name    string `json:"name,omitempty" db:"name"`
		Type    string `json:"type,omitempty" db:"type"`
		Primary bool   `json:"primary,omitempty" db:"primary"`
	}

	//AccountPhones, list of AccountPhone
	//
	//  "accountPhones": [
	//		{
	//			"uid": "3e22eb32b25e7e3b6c7261e3d2d2654c",
	//			"countryCode": "371",
	//			"number": "20000000",
	//			"type": "M",
	//			"primary": true,
	//			"confirmed": true
	//		}
	//	],
	AccountPhones []*AccountPhone

	//Item of AccountPhones
	AccountPhone struct {
		Id     int64 `json:"id,omitempty" db:"id"`          //Internal for user-integration
		UserId int64 `json:"userId,omitempty" db:"user_id"` //Internal for user-integration
		//ITH.AMS data structure
		Uid               string    `json:"uid,omitempty" db:"uid"`                    //Phone UID
		CountryCode       string    `json:"countryCode,omitempty" db:"country_code"`   //Phone country code
		Number            string    `json:"number,omitempty" db:"number"`              //Phone number
		Type              PhoneType `json:"type,omitempty" db:"type"`                  //Phone type (see: PhoneType)
		ContactPreference bool      `json:"contactPreference" db:"contact_preference"` //Phone is preferred for communication
		Primary           bool      `json:"primary,omitempty" db:"primary"`            //Phone is primary
		Confirmed         bool      `json:"confirmed,omitempty" db:"confirmed"`        //Phone is confirmed by account holder
	}

	//List of account settings
	//
	//  "accountSettings": [
	//		{
	//			"name": "externalAccountUid",
	//			"value": "EX-ACC-UID-1234"
	//			"category": "ACC"
	//		}
	//	],
	AccountSettings []*AccountSetting

	//Item of AccountSettings list
	AccountSetting struct {
		Id     int64 `json:"id,omitempty" db:"id"`          //Internal for user-integration
		UserId int64 `json:"userId,omitempty" db:"user_id"` //Internal for user-integration
		//ITH.AMS data structure
		Name     string `json:"name" db:"name"`                   //Setting name
		Value    string `json:"value,omitempty" db:"value"`       //Setting value
		Category string `json:"category,omitempty" db:"category"` //Setting category
	}

	//AccountEmails list of AccountEmail
	//
	//  "accountEmails": [
	//		{
	//			"uid": "5340c5dd139528019a01821ba8bc7f09",
	//			"email": "john@enauda.com",
	//			"confirmed": false,
	//			"primary": true
	//		}
	//	],
	AccountEmails []*AccountEmail

	//Item of AccountEmails list
	AccountEmail struct {
		Id     int64 `json:"id,omitempty" db:"id"`          //Internal for user-integration
		UserId int64 `json:"userId,omitempty" db:"user_id"` //Internal for user-integration
		//ITH.AMS data structure
		Uid       string `json:"uid,omitempty" db:"uid"`
		Email     string `json:"email" db:"email"`
		Confirmed bool   `json:"confirmed,omitempty" db:"confirmed"`
		Primary   bool   `json:"primary,omitempty" db:"primary"`
		Type      string `json:"type,omitempty" db:"type"`
	}

	//Person object
	//
	//	   "person": {
	//		   "name": "John",
	//		   "surname": "Doe",
	//		   "birthDate": "19810509000000",
	//           "pep": false
	//	   },
	Person struct {
		Name      string  `json:"name"`      //Name
		Surname   string  `json:"surname"`   //Surname
		BirthDate AmsDate `json:"birthDate"` //Date of birth,Format – yyyyMMddHHmmss
		Pep       bool    `json:"pep"`       //Person in politically exposed person (PEP)
	}

	//ErrorData - any response
	ErrorData struct {
		ErrorCode    int         `json:"errorCode"`    //Error code
		ErrorMessage string      `json:"errorMessage"` //Localized error message. Supported languages are English, Russian, and Latvian. English is used	when no customer locale is available
		RequestUid   string      `json:"requestUid"`   //Request UID, used for investigation of exceptional cases
		Parameters   interface{} `json:"parameters"`   //Error extended parameters
	}
)
