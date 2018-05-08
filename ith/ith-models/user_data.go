package ith_models
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
	 */
	Address struct {
		CountryCode       string `json:"countryCode"`       //ISO2 country code, required
		City              string `json:"city"`              //City, required
		FirstAddressLine  string `json:"firstAddressLine"`  //First address line, required
		SecondAddressLine string `json:"secondAddressLine"` //Second address line, optional
		State             string `json:"state"`             //State, optional
		PostalCode        string `json:"postalCode"`        //Postal code, required
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
		Email                        string   `json:"email"`                        //Email, required
		Phone                        *Phone   `json:"phone"`                        //Phone
		ExternalUid                  string   `json:"externalUid"`                  //Account UID in external system
		Name                         string   `json:"name"`                         //Name, required
		SureName                     string   `json:"surename"`                     //Surname, required
		BirthDate                    string   `json:"birthDate"`                    //Date of birth, Format - yyyyMMddHHmmss
		Language                     string   `json:"language"`                     //ISO2 language code
		Address                      *Address `json:"address"`                      //Address
		WebSite                      string   `json:"website"`                      //Website
		TaxpayerIdentificationNumber string   `json:"taxpayerIdentificationNumber"` //Taxpayer Identification Number
		AdditionalInfo               string   `json:"additionalInfo"`               //Additional information
	}
)