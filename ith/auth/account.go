package auth

type (
	Phone struct {
		CountryCode *string `json:"countryCode,omitempty"` //String(5); Phone country code (optional, when fullNumber filled)
		Number      *string `json:"number,omitempty"`      //String(25); Phone number (optional, when fullNumber filled)
		FullNumber  *string `json:"fullNumber,omitempty"`  //String(30); Full phone number. Min length 5 (optional, when countryCode+number filled)
	}

	Address struct {
		CountryCode       string `json:"countryCode,omitempty"`       //String(2); ISO2 country code, required
		City              string `json:"city,omitempty"`              //String(50); City, required
		FirstAddressLine  string `json:"firstAddressLine,omitempty"`  //String(60); First address line, required
		SecondAddressLine string `json:"secondAddressLine,omitempty"` //String(60); Second address line, optional
		State             string `json:"state,omitempty"`             //String(50); State, optional
		PostalCode        string `json:"postalCode,omitempty"`        //String(10); Postal code, required
	}

	Account struct {
		Email                        string   `json:"email,omitempty,omitempty"`              //String(150); Email, required
		Phone                        *Phone   `json:"phone,omitempty,omitempty"`              //Phone
		ExternalUid                  string   `json:"externalUid,omitempty,omitempty"`        //String(50); Account UID in external system
		Name                         string   `json:"name,omitempty"`                         //String(50); Name, required
		Surname                      string   `json:"surname,omitempty"`                      //String(50); Surname, required
		BirthDate                    string   `json:"birthDate,omitempty"`                    //Date of birth, Format - yyyyMMddHHmmss
		Language                     string   `json:"language,omitempty"`                     //String(2); ISO2 language code
		Address                      *Address `json:"address,omitempty"`                      //Address
		WebSite                      string   `json:"website,omitempty"`                      //String(2048); Website
		TaxpayerIdentificationNumber string   `json:"taxpayerIdentificationNumber,omitempty"` //String(255); Taxpayer Identification Number
		AdditionalInfo               string   `json:"additionalInfo,omitempty"`               //String(5000); Additional information
	}
)
