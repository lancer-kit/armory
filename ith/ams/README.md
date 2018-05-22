# ams
--
    import "gitlab.inn4science.com/vcg/go-common/ith/ams"

ITH integration. ACCOUNT MANAGEMENT SERVICES

## Usage

```go
var ErrAccountStatusInvalid = errors.New("AccountStatus is invalid")
```

```go
var ErrAccountTypeInvalid = errors.New("AccountType is invalid")
```

```go
var ErrActionConfirmationInvalid = errors.New("ActionConfirmation is invalid")
```

```go
var ErrAddressTypeInvalid = errors.New("AddressType is invalid")
```

```go
var ErrPhoneTypeInvalid = errors.New("PhoneType is invalid")
```

```go
var ErrWebResourceInvalid = errors.New("WebResource is invalid")
```

#### type Account

```go
type Account struct {
	Uid                       string             `json:"uid"`                       //Account UID in ITH platform
	Country                   *Country           `json:"country"`                   //Account country object
	Language                  *Language          `json:"language"`                  //Account language object
	CommunicationLanguage     *Language          `json:"communicationLanguage"`     //Account communication language object
	Type                      AccountType        `json:"type"`                      //Account type: see AccountType
	Status                    AccountStatus      `json:"status"`                    //Account status: see AccountStatus
	AccountPhones             AccountPhones      `json:"accountPhones"`             //List of account phones, []*AccountPhone
	AccountSettings           AccountSettings    `json:"accountSettings"`           //List of account settings
	AccountEmails             AccountEmails      `json:"accountEmails"`             //List of account emails
	Addresses                 Addresses          `json:"addresses"`                 //List of Address objects
	Person                    *Person            `json:"person"`                    //Person object
	Company                   *Company           `json:"company,omitempty"`         //Company object (for merchant only)
	AffiliateId               string             `json:"affiliateId"`               //Affiliate ID, String(50)
	CampaignId                string             `json:"campaignId"`                //Campaign ID, String(50)
	BannerId                  string             `json:"bannerId"`                  //Banner ID, String(50)
	CustomParameters          string             `json:"customParameters"`          //Custom parameters, String(255)
	AccountSecret             string             `json:"accountSecret,omitempty"`   //Account secret (for merchant only), String(20)
	MerchantUid               string             `json:"merchantUid,omitempty"`     //Merchant UID (for merchant only), String(36)
	Timezone                  int                `json:"timezone"`                  //Account time zone ID
	WeekStartsOn              string             `json:"weekStartsOn"`              //Start day of the week, String(2)
	CurrencyConversion        bool               `json:"currencyConversion"`        //Currency conversion is enabled
	AlwaysRefundEWallet       bool               `json:"alwaysRefundEWallet"`       //Refunds are transferred to EWallet
	ConfirmOutTransaction     bool               `json:"confirmOutTransaction"`     //2 step verification for outgoing transactions
	ConfirmLogin              bool               `json:"confirmLogin"`              //2 step verification for login
	ActionConfirmationEnabled bool               `json:"actionConfirmationEnabled"` //2 step verification enabled
	ActionConfirmationType    ActionConfirmation `json:"actionConfirmationType"`
	Test                      bool               `json:"test"` //Account is test
}
```


#### type AccountEmail

```go
type AccountEmail struct {
	Uid       string `json:"uid,omitempty"`
	Email     string `json:"email"`
	Confirmed bool   `json:"confirmed,omitempty"`
	Primary   bool   `json:"primary,omitempty"`
	Type      string `json:"type,omitempty"`
}
```

Item of AccountEmails list

#### type AccountEmails

```go
type AccountEmails []*AccountEmail
```

AccountEmails list of AccountEmail

     "accountEmails": [
    		{
    			"uid": "5340c5dd139528019a01821ba8bc7f09",
    			"email": "john@enauda.com",
    			"confirmed": false,
    			"primary": true
    		}
    	],

#### type AccountPhone

```go
type AccountPhone struct {
	Uid               string    `json:"uid,omitempty"`         //Phone UID
	CountryCode       string    `json:"countryCode,omitempty"` //Phone country code
	Number            string    `json:"number,omitempty"`      //Phone number
	Type              PhoneType `json:"type,omitempty"`        //Phone type (see: PhoneType)
	ContactPreference bool      `json:"contactPreference"`     //Phone is preferred for communication
	Primary           bool      `json:"primary,omitempty"`     //Phone is primary
	Confirmed         bool      `json:"confirmed,omitempty"`   //Phone is confirmed by account holder
}
```

Item of AccountPhones

#### type AccountPhones

```go
type AccountPhones []*AccountPhone
```

AccountPhones, list of AccountPhone

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

#### type AccountResponse

```go
type AccountResponse struct {
	Account *Account `json:"account"`
}
```

Standard account response (from doc.example)

#### type AccountSetting

```go
type AccountSetting struct {
	Name     string `json:"name"`               //Setting name
	Value    string `json:"value,omitempty"`    //Setting value
	Category string `json:"category,omitempty"` //Setting category
}
```

Item of AccountSettings list

#### type AccountSettings

```go
type AccountSettings []*AccountSetting
```

List of account settings

     "accountSettings": [
    		{
    			"name": "externalAccountUid",
    			"value": "EX-ACC-UID-1234"
    			"category": "ACC"
    		}
    	],

#### type AccountStatus

```go
type AccountStatus int
```

AccountStatus:

    * SA – Standard: Automatically Registered
    * SR – Standard: Registration Requested
    * SC – Standard: Registration Confirmed
    * SF – Standard: Customer Identified
    * SB – Standard: Blocked
    * SD – Standard: Closed
    * BR – Business: Registration Requested
    * BC – Business: Registration Confirmed (Read Only)
    * BF – Business: Registration Finished (Agreement Signed)
    * BM – Business: Requires Moderation
    * BS – Business: Suspended (Blocked)
    * BD – Business: Closed
    * MR – Merchant: Registration Requested
    * MC – Merchant: Registration Confirmed (Read Only)
    * MF – Merchant: Registration Finished (Agreement Signed)
    * MM – Merchant: Requires Moderation
    * MS – Merchant: Suspended (Blocked)
    * MD – Merchant: Closed

```go
const (
	StStandardAutomaticallyRegistered AccountStatus = iota //SA – Standard: Automatically Registered
	StStandardRegistrationRequested                        //SR – Standard: Registration Requested
	StStandardRegistrationConfirmed                        //SC – Standard: Registration Confirmed
	StStandardCustomerIdentified                           //SF – Standard: Customer Identified
	StStandardBlocked                                      //SB – Standard: Blocked
	StStandardClosed                                       //SD – Standard: Closed
	StBusinessRegistrationRequested                        //BR – Business: Registration Requested
	StBusinessRegistrationConfirmed                        //BC – Business: Registration Confirmed (Read Only)
	StBusinessRegistrationFinished                         //BF – Business: Registration Finished (Agreement Signed)
	StBusinessRequiresModeration                           //BM – Business: Requires Moderation
	StBusinessSuspended                                    //BS – Business: Suspended (Blocked)
	StBusinessClosed                                       //BD – Business: Closed
	StMerchantRegistrationRequested                        //MR – Merchant: Registration Requested
	StMerchantRegistrationConfirmed                        //MC – Merchant: Registration Confirmed (Read Only)
	StMerchantRegistrationFinished                         //MF – Merchant: Registration Finished (Agreement Signed)
	StMerchantRequiresModeration                           //MM – Merchant: Requires Moderation
	StMerchantSuspended                                    //MS – Merchant: Suspended (Blocked)
	StMerchantClosed                                       //MD – Merchant: Closed
)
```

#### func (AccountStatus) MarshalJSON

```go
func (r AccountStatus) MarshalJSON() ([]byte, error)
```
MarshalJSON is generated so AccountStatus satisfies json.Marshaler.

#### func (*AccountStatus) Scan

```go
func (r *AccountStatus) Scan(src interface{}) error
```
Value is generated so AccountStatus satisfies db row driver.Scanner.

#### func (AccountStatus) String

```go
func (r AccountStatus) String() string
```
String is generated so AccountStatus satisfies fmt.Stringer.

#### func (*AccountStatus) UnmarshalJSON

```go
func (r *AccountStatus) UnmarshalJSON(data []byte) error
```
UnmarshalJSON is generated so AccountStatus satisfies json.Unmarshaler.

#### func (AccountStatus) Validate

```go
func (r AccountStatus) Validate() error
```
Validate verifies that value is predefined for AccountStatus.

#### func (AccountStatus) Value

```go
func (r AccountStatus) Value() (driver.Value, error)
```
Value is generated so AccountStatus satisfies db row driver.Valuer.

#### type AccountType

```go
type AccountType int
```

AccountType:

    * S - Standard
    * M - Merchant
    * B - Business

```go
const (
	AccountTypeStandard AccountType = iota //Standard
	AccountTypeMerchant                    //Merchant
	AccountTypeBusiness                    //Business
)
```

#### func (AccountType) MarshalJSON

```go
func (r AccountType) MarshalJSON() ([]byte, error)
```
MarshalJSON is generated so AccountType satisfies json.Marshaler.

#### func (*AccountType) Scan

```go
func (r *AccountType) Scan(src interface{}) error
```
Value is generated so AccountType satisfies db row driver.Scanner.

#### func (AccountType) String

```go
func (r AccountType) String() string
```
String is generated so AccountType satisfies fmt.Stringer.

#### func (*AccountType) UnmarshalJSON

```go
func (r *AccountType) UnmarshalJSON(data []byte) error
```
UnmarshalJSON is generated so AccountType satisfies json.Unmarshaler.

#### func (AccountType) Validate

```go
func (r AccountType) Validate() error
```
Validate verifies that value is predefined for AccountType.

#### func (AccountType) Value

```go
func (r AccountType) Value() (driver.Value, error)
```
Value is generated so AccountType satisfies db row driver.Valuer.

#### type ActionConfirmation

```go
type ActionConfirmation int
```

Field type for Account.ActionConfirmationType

    * "EMAIL" – via email;
    * "SMS" – via phone
    * "GAUTH" – via Google Authenticator

```go
const (
	ActionConfirmationEmail ActionConfirmation = iota //EMAIL – via email;
	ActionConfirmationSms                             //SMS – via phone
	ActionConfirmationGAuth                           //GAUTH – via Google Authenticator
)
```

#### func (ActionConfirmation) MarshalJSON

```go
func (r ActionConfirmation) MarshalJSON() ([]byte, error)
```
MarshalJSON is generated so ActionConfirmation satisfies json.Marshaler.

#### func (*ActionConfirmation) Scan

```go
func (r *ActionConfirmation) Scan(src interface{}) error
```
Value is generated so ActionConfirmation satisfies db row driver.Scanner.

#### func (ActionConfirmation) String

```go
func (r ActionConfirmation) String() string
```
String is generated so ActionConfirmation satisfies fmt.Stringer.

#### func (*ActionConfirmation) UnmarshalJSON

```go
func (r *ActionConfirmation) UnmarshalJSON(data []byte) error
```
UnmarshalJSON is generated so ActionConfirmation satisfies json.Unmarshaler.

#### func (ActionConfirmation) Validate

```go
func (r ActionConfirmation) Validate() error
```
Validate verifies that value is predefined for ActionConfirmation.

#### func (ActionConfirmation) Value

```go
func (r ActionConfirmation) Value() (driver.Value, error)
```
Value is generated so ActionConfirmation satisfies db row driver.Valuer.

#### type Address

```go
type Address struct {
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
```

Address type, item of Addresses list

#### type AddressRequest

```go
type AddressRequest struct {
	CountryCode string `json:"countryCode"` //Required, ISO-2 country code, String(2)
	City        string `json:"city"`        //Required, String(50)
}
```

AddressRequest - used in "Create Standard Account" method of Account Management
Services

Example:

     {
    		"city": "Riga",
    		"countryCode": "LV",
    		"firstAddressLine": "Duntes 4",
    		"secondAddressLine": "Office 403",
    		"postalCode": "1234",
    		"state": "Rigas rajons"
    	}

#### func (*AddressRequest) Validate

```go
func (r *AddressRequest) Validate() error
```

#### type AddressType

```go
type AddressType int
```

AddressType:

    * B – Business;
    * H – Home;
    * O – Other;
    * C – Communication

```go
const (
	AddressTypeBusiness      AddressType = iota //B – Business;
	AddressTypeHome                             //H – Home;
	AddressTypeOther                            //O – Other;
	AddressTypeCommunication                    //C – Communication

)
```

#### func (AddressType) MarshalJSON

```go
func (r AddressType) MarshalJSON() ([]byte, error)
```
MarshalJSON is generated so AddressType satisfies json.Marshaler.

#### func (*AddressType) Scan

```go
func (r *AddressType) Scan(src interface{}) error
```
Value is generated so AddressType satisfies db row driver.Scanner.

#### func (AddressType) String

```go
func (r AddressType) String() string
```
String is generated so AddressType satisfies fmt.Stringer.

#### func (*AddressType) UnmarshalJSON

```go
func (r *AddressType) UnmarshalJSON(data []byte) error
```
UnmarshalJSON is generated so AddressType satisfies json.Unmarshaler.

#### func (AddressType) Validate

```go
func (r AddressType) Validate() error
```
Validate verifies that value is predefined for AddressType.

#### func (AddressType) Value

```go
func (r AddressType) Value() (driver.Value, error)
```
Value is generated so AddressType satisfies db row driver.Valuer.

#### type Addresses

```go
type Addresses []*Address
```

Address type and address list

Example:

     {
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
    	]
     }

#### type AffiliateInfo

```go
type AffiliateInfo struct {
	AffiliateId      string `json:"affiliateId"`      //Optional, String(50), Affiliate ID
	CampaignId       string `json:"campaignId"`       //Optional, String(50), Campaign ID
	BannerId         string `json:"bannerId"`         //Optional, String(50), Banner ID
	CustomParameters string `json:"customParameters"` //Optional, String(255), Custom parameters
}
```

Additional field affiliateInfo

    "affiliateInfo":
    	{
    		"affiliateId": "AF4",
    		"campaignId": "C539",
    		"bannerId": "BRT13",
    		"customParameters": "tr=24&hd=3"
    	}

#### type AmsDate

```go
type AmsDate time.Time
```


#### func (AmsDate) MarshalJSON

```go
func (r AmsDate) MarshalJSON() ([]byte, error)
```
MarshalJSON AmsDate satisfies json.Marshaler.

#### func (AmsDate) String

```go
func (r AmsDate) String() string
```
String is generated so AddressType satisfies fmt.Stringer.

#### func (AmsDate) Time

```go
func (r AmsDate) Time() time.Time
```

#### func (*AmsDate) UnmarshalJSON

```go
func (r *AmsDate) UnmarshalJSON(data []byte) error
```
UnmarshalJSON AmsDate satisfies json.Unmarshaler.

#### func (AmsDate) Validate

```go
func (r AmsDate) Validate() error
```
Validate verifies that value is predefined for AddressType.

#### type Company

```go
type Company struct {
	BusinessName                    string  `json:"businessName"`                    //Company name, String(255)
	CategoryId                      int     `json:"categoryId"`                      //Category ID, Integer
	BusinessTypeId                  int     `json:"businessTypeId"`                  //Business type ID, Integer
	CardStatementName               string  `json:"cardStatementName"`               //Card statement name, String(50)
	CardStatementNameExt            string  `json:"cardStatementNameExt"`            //Extended card statement name, String(50)
	CallbackUrl                     string  `json:"callbackUrl"`                     //URL for callbacks
	RollingReservePrc               float64 `json:"rollingReservePrc"`               //Rolling reserve rate (in %), Number
	RollingReserveHoldDays          int     `json:"rollingReserveHoldDays"`          //Rolling reserve hold days
	SendCallback                    bool    `json:"sendCallback"`                    //Send callbacks for merchant
	AcceptUndefinedProvisionChannel bool    `json:"acceptUndefinedProvisionChannel"` //Accept undefined provision channels
	AllowDuplicateOrderExternalId   bool    `json:"allowDuplicateOrderExternalId"`   //Allow duplicate order external ID
	AllowNotificationsForSeller     bool    `json:"allowNotificationsForSeller"`     //Send notifications for seller
	AllowNotificationsForBuyer      bool    `json:"allowNotificationsForBuyer"`      //Send notifications for buyer
	AllowPartialPayments            bool    `json:"allowPartialPayments"`            //Allow partial payments
}
```

Company Company object (for merchant only)

#### type Country

```go
type Country struct {
	Id                    int64  `json:"id,omitempty"`
	Code                  string `json:"code"`
	Name                  string `json:"name"`
	BrandedCardsAvailable bool   `json:"brandedCardsAvailable,omitempty"`
	RegistrationAllowed   bool   `json:"registrationAllowed,omitempty"`
}
```

Additional field `country`

      "country": {
    		"code": "LV",
    		"name": "Latvia",
    		"brandedCardsAvailable": true
    	},

#### type Language

```go
type Language struct {
	Uid     string `json:"uid,omitempty"`
	Code    string `json:"code"` //ISO2 language code
	Name    string `json:"name,omitempty"`
	Type    string `json:"type,omitempty"`
	Primary bool   `json:"primary,omitempty"`
}
```

Additional field `language`

     "language": {
    		"code": "ru",
    		"name": "Russian"
    	},

#### type Person

```go
type Person struct {
	Name      string  `json:"name"`      //Name
	Surname   string  `json:"surname"`   //Surname
	BirthDate AmsDate `json:"birthDate"` //Date of birth,Format – yyyyMMddHHmmss
	Pep       bool    `json:"pep"`       //Person in politically exposed person (PEP)
}
```

Person object

    	   "person": {
    		   "name": "John",
    		   "surname": "Doe",
    		   "birthDate": "19810509000000",
              "pep": false
    	   },

#### type Phone

```go
type Phone struct {
	CountryCode *string `json:"countryCode,omitempty"` //Phone country code (optional, when fullNumber filled)
	Number      *string `json:"number,omitempty"`      //Phone number (optional, when fullNumber filled)
	FullNumber  *string `json:"fullNumber,omitempty"`  //Full phone number. Min length 5 (optional, when countryCode+number filled)
}
```

Use (countryCode and number) OR fullNumber

Example:

    	...
     "phone": {
    	  "countryCode ": "371",
    	  "number": "20000000",
    	  "fullNumber": "37120000000",
      }
     ...

#### type PhoneType

```go
type PhoneType int
```

PhoneType:

* M - Mobile

* H - Home

* W - Work

```go
const (
	PhoneTypeMobile PhoneType = iota //Mobile
	PhoneTypeHome                    //Home
	PhoneTypeWork                    //Work
)
```

#### func (PhoneType) MarshalJSON

```go
func (r PhoneType) MarshalJSON() ([]byte, error)
```
MarshalJSON is generated so PhoneType satisfies json.Marshaler.

#### func (*PhoneType) Scan

```go
func (r *PhoneType) Scan(src interface{}) error
```
Value is generated so PhoneType satisfies db row driver.Scanner.

#### func (PhoneType) String

```go
func (r PhoneType) String() string
```
String is generated so PhoneType satisfies fmt.Stringer.

#### func (*PhoneType) UnmarshalJSON

```go
func (r *PhoneType) UnmarshalJSON(data []byte) error
```
UnmarshalJSON is generated so PhoneType satisfies json.Unmarshaler.

#### func (PhoneType) Validate

```go
func (r PhoneType) Validate() error
```
Validate verifies that value is predefined for PhoneType.

#### func (PhoneType) Value

```go
func (r PhoneType) Value() (driver.Value, error)
```
Value is generated so PhoneType satisfies db row driver.Valuer.

#### type UserRegistrationRequest

```go
type UserRegistrationRequest struct {
	ClientId           string         `json:"clientId"`           //Required, String(50), OAuth client ID
	ClientSecret       string         `json:"clientSecret"`       //Required, String(32), OAuth client secret
	ExternalAccountUid string         `json:"externalAccountUid"` //Required, String(50), User API user ID
	Phone              string         `json:"phone"`              //Required, String(30), Full phone number. Min length 5
	Email              string         `json:"email"`              //Optional, String(150), Email
	Password           string         `json:"password"`           //Required, String(50), User account password (plain ?)
	FirstName          string         `json:"firstName"`          //Required, String(50), Name
	LastName           string         `json:"lastName"`           //Required, String(50), Surname
	BirthDate          string         `json:"birthDate"`          //Optional, Date, Format - yyyyMMddHHmmss
	Country            string         `json:"country"`            //Required, String(2), ISO2 country code
	Language           string         `json:"language"`           //Optional, String(2), ISO2 language code
	Address            *Address       //Required, User account address
	AffiliateInfo      *AffiliateInfo //Optional, Affiliate information
}
```

Request body for POST https://<host>:<port>/partnerapi/account/register

Example:

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

#### type UserRegistrationResponse

```go
type UserRegistrationResponse struct {
	AccountUid         string   `json:"accountUid"`
	ExternalAccountUid string   `json:"externalAccountUid"`
	AccessToken        string   `json:"accessToken"`
	Account            *Account `json:"account"`
}
```

UserRegistrationResponse response from ITH Account Management Services (AMS)

     {
    	"accountUid": "100-014-275-55",
    	"externalAccountUid": "EX-ACC-UID-1234",
    	"accessToken": "bdad264b7f8b9896d73436b234e4bddd",
    	"account": {....}
     }

#### type WebResource

```go
type WebResource int
```

WebResource:

    * W – Web;
    * F – Facebook;
    * T – Twitter

```go
const (
	WebResourceWeb      WebResource = iota // W – Web;
	WebResourceFacebook                    //F – Facebook;
	WebResourceTwitter                     //T – Twitter
)
```

#### func (WebResource) MarshalJSON

```go
func (r WebResource) MarshalJSON() ([]byte, error)
```
MarshalJSON is generated so WebResource satisfies json.Marshaler.

#### func (*WebResource) Scan

```go
func (r *WebResource) Scan(src interface{}) error
```
Value is generated so WebResource satisfies db row driver.Scanner.

#### func (WebResource) String

```go
func (r WebResource) String() string
```
String is generated so WebResource satisfies fmt.Stringer.

#### func (*WebResource) UnmarshalJSON

```go
func (r *WebResource) UnmarshalJSON(data []byte) error
```
UnmarshalJSON is generated so WebResource satisfies json.Unmarshaler.

#### func (WebResource) Validate

```go
func (r WebResource) Validate() error
```
Validate verifies that value is predefined for WebResource.

#### func (WebResource) Value

```go
func (r WebResource) Value() (driver.Value, error)
```
Value is generated so WebResource satisfies db row driver.Valuer.
