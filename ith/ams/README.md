# ams
--
    import "gitlab.inn4science.com/vcg/go-common/ith/ams"

ITH integration. ACCOUNT MANAGEMENT SERVICES

## Usage

```go
const (
	APICreate = "/partnerapi/account/register"
	APIupdate = "/partnerapi/account/update"
	APICode   = "/commonapi/auth/" + param + "/authorization_code"
	APIToken  = "/partnerapi/token/code"
)
```

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
var ErrRequestStatusInvalid = errors.New("RequestStatus is invalid")
```

```go
var ErrWebResourceInvalid = errors.New("WebResource is invalid")
```

#### type API

```go
type API struct {
	Config Config

	Auth ith.API
}
```


#### func  NewAPI

```go
func NewAPI(baseUrl, commonUrl, client, secret string) *API
```

#### func (*API) CreateProfile

```go
func (api *API) CreateProfile(req *UserRegistrationRequest) (usr *UserRegistrationResponse, err error, status RequestStatus)
```
CreateProfile - request partner API to create the new standard user profile

#### func (*API) GetCode

```go
func (api *API) GetCode(req *AuthCodeRequest) (code *AuthCodeResponse, err error)
```
Get Authorization Code Send request to ITH.Authorization service. Service is
used to receive one-time authorization code. This single use code could be used
to transfer user session from one module to another.

#### func (*API) GetToken

```go
func (api *API) GetToken(req *AuthCodeRequest) (token *AuthTokenResponse, err error)
```
GetToken - Get Authorization Token Send request to ITH.Authorization service.
Service is used to receive customer access token and refresh token using
one-time authorization code. Received access token should be used in other
services calls.

#### func (*API) SetLogger

```go
func (api *API) SetLogger(entry log.Entry)
```
Set new logger on ams.API

#### func (*API) UpdateProfile

```go
func (api *API) UpdateProfile(req *UserUpdateRequest, token string) (usr *UserRegistrationResponse, err error, status RequestStatus)
```
CreateProfile - request partner API to update the standard user profile

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

Example:

    {
    	"uid": "100-020-425-40",
    	"country": {/*object*/},
    	"language": {/*object*/},
    	"communicationLanguage": {/*object*/},
    	"type": "S",
    	"status": "SC",
    	"accountPhones": [/*list*/],
    	"accountSettings": [/*list*/],
    	"accountEmails": [/*list*/],
    	"addresses": [/*list*/],
    	"person": {/*object*/},
    	"affiliateId": "AF4",
    	"campaignId": "C539",
    	"bannerId": "BRT13",
    	"customParameters": "tr=24&hd=3",
    	"timezone": 16,
    	"weekStartsOn": "MO",
    	"currencyConversion": true,
    	"alwaysRefundEWallet": false,
    	"confirmOutTransaction": false,
    	"confirmLogin": false,
    	"actionConfirmationEnabled": false,
    	"test": false
    }

#### type AccountEmail

```go
type AccountEmail struct {
	Id     int64 `json:"id,omitempty" db:"id"`          //Internal for user-integration
	UserId int64 `json:"userId,omitempty" db:"user_id"` //Internal for user-integration
	//ITH.AMS data structure
	Uid       string `json:"uid,omitempty" db:"uid"`
	Email     string `json:"email" db:"email"`
	Confirmed bool   `json:"confirmed,omitempty" db:"confirmed"`
	Primary   bool   `json:"primary,omitempty" db:"primary"`
	Type      string `json:"type,omitempty" db:"type"`
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
	Id     int64 `json:"id,omitempty" db:"id"`          //Internal for user-integration
	UserId int64 `json:"userId,omitempty" db:"user_id"` //Internal for user-integration
	//ITH.AMS data structure
	Name     string `json:"name" db:"name"`                   //Setting name
	Value    string `json:"value,omitempty" db:"value"`       //Setting value
	Category string `json:"category,omitempty" db:"category"` //Setting category
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
	StStandardAutomaticallyRegistered AccountStatus = iota + 1 //SA – Standard: Automatically Registered
	StStandardRegistrationRequested                            //SR – Standard: Registration Requested
	StStandardRegistrationConfirmed                            //SC – Standard: Registration Confirmed
	StStandardCustomerIdentified                               //SF – Standard: Customer Identified
	StStandardBlocked                                          //SB – Standard: Blocked
	StStandardClosed                                           //SD – Standard: Closed
	StBusinessRegistrationRequested                            //BR – Business: Registration Requested
	StBusinessRegistrationConfirmed                            //BC – Business: Registration Confirmed (Read Only)
	StBusinessRegistrationFinished                             //BF – Business: Registration Finished (Agreement Signed)
	StBusinessRequiresModeration                               //BM – Business: Requires Moderation
	StBusinessSuspended                                        //BS – Business: Suspended (Blocked)
	StBusinessClosed                                           //BD – Business: Closed
	StMerchantRegistrationRequested                            //MR – Merchant: Registration Requested
	StMerchantRegistrationConfirmed                            //MC – Merchant: Registration Confirmed (Read Only)
	StMerchantRegistrationFinished                             //MF – Merchant: Registration Finished (Agreement Signed)
	StMerchantRequiresModeration                               //MM – Merchant: Requires Moderation
	StMerchantSuspended                                        //MS – Merchant: Suspended (Blocked)
	StMerchantClosed                                           //MD – Merchant: Closed
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
	AccountTypeStandard AccountType = iota + 1 //Standard
	AccountTypeMerchant                        //Merchant
	AccountTypeBusiness                        //Business
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
	ActionConfirmationNone ActionConfirmation = iota //EMAIL – via email;
	ActionConfirmationEmail
	ActionConfirmationSms   //SMS – via phone
	ActionConfirmationGAuth //GAUTH – via Google Authenticator
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
```

Address type, item of Addresses list

#### type AddressRequest

```go
type AddressRequest struct {
	CountryCode       string `json:"countryCode"`       //Required, String(2), ISO-2 country code
	City              string `json:"city"`              //Required, String(50), City
	FirstAddressLine  string `json:"firstAddressLine"`  //Required, String(60), First address line
	SecondAddressLine string `json:"secondAddressLine"` //Optional, String(60), Second address line
	PostalCode        string `json:"postalCode"`        //Required, String(10), Postal code
	State             string `json:"state"`             //Optional,String(50), State
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
func (r *AddressRequest) Validate() (err error)
```
Validate verifies that value is predefined for AddressRequest.

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
	AddressTypeBusiness      AddressType = iota + 1 //B – Business;
	AddressTypeHome                                 //H – Home;
	AddressTypeOther                                //O – Other;
	AddressTypeCommunication                        //C – Communication

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

#### type AddressUpdate

```go
type AddressUpdate struct {
	CountryCode       string `json:"countryCode,omitempty"`       //Optional/fill if updated, String(2), ISO-2 country code
	City              string `json:"city,omitempty"`              //Optional/fill if updated, String(50), City
	FirstAddressLine  string `json:"firstAddressLine,omitempty"`  //Optional/fill if updated, String(60), First address line
	SecondAddressLine string `json:"secondAddressLine,omitempty"` //Optional/fill if updated, String(60), Second address line
	PostalCode        string `json:"postalCode,omitempty"`        //Optional/fill if updated, String(10), Postal code
	State             string `json:"state,omitempty"`             //Optional/fill if updated,String(50), State
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

Additional field affiliateInfo - used in "Create Standard Account" method of
Account Management Services

    "affiliateInfo":
    	{
    		"affiliateId": "AF4",
    		"campaignId": "C539",
    		"bannerId": "BRT13",
    		"customParameters": "tr=24&hd=3"
    	}

#### func (*AffiliateInfo) Validate

```go
func (r *AffiliateInfo) Validate() (err error)
```
Validate verifies that value is predefined for AffiliateInfo.

#### type AmsDate

```go
type AmsDate time.Time
```


#### func  AmsDateFromInt

```go
func AmsDateFromInt(i int64) AmsDate
```

#### func (AmsDate) Empty

```go
func (r AmsDate) Empty() bool
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

#### type AuthCodeRequest

```go
type AuthCodeRequest struct {
	AccessToken string `json:"accessToken,omitempty"` //user access token
	Username    string `json:"username,omitempty"`    //Account username (email)
	Password    string `json:"password,omitempty"`    //Account password (plain password)
}
```


#### func (*AuthCodeRequest) Validate

```go
func (r *AuthCodeRequest) Validate() error
```

#### type AuthCodeResponse

```go
type AuthCodeResponse struct {
	//ErrorData *ErrorData `json:"errorData,omitempty"` //Not returned if operation is successful
	Code string `json:"code"` //One-time authorization code
}
```


#### type AuthTokenRequest

```go
type AuthTokenRequest struct {
	ClientId     string `json:"clientId"`     //OAuth client ID
	ClientSecret string `json:"clientSecret"` //OAuth client secret
	Code         string `json:"code"`         //One-time authorization code
}
```


#### type AuthTokenResponse

```go
type AuthTokenResponse struct {
	ErrorData    *ErrorData `json:"errorData,omitempty"` //Not returned if operation is successful
	AccessToken  string     `json:"accessToken"`         //Access token for integration services
	RefreshToken string     `json:"refreshToken"`        //Refresh token for access token renewal
	ExpiresIn    int64      `json:"expiresIn"`           //Expiration time for access token (seconds)
}
```


#### type Company

```go
type Company struct {
	Id     int64  `json:"id,omitempty" db:"id"`          //Internal for user-integration
	UserId int64  `json:"userId,omitempty" db:"user_id"` //Internal for user-integration
	Uid    string `json:"uid,omitempty" db:"uid"`        //Internal for user-integration Company UID

	//ITH.AMS data structure
	BusinessName                    string  `json:"businessName" db:"business_name"`                                         //Company name, String(255)
	CategoryId                      int     `json:"categoryId" db:"category_id"`                                             //Category ID, Integer
	BusinessTypeId                  int     `json:"businessTypeId" db:"business_type_id"`                                    //Business type ID, Integer
	CardStatementName               string  `json:"cardStatementName" db:"card_statement_name"`                              //Card statement name, String(50)
	CardStatementNameExt            string  `json:"cardStatementNameExt" db:"card_statement_name_ext"`                       //Extended card statement name, String(50)
	CallbackUrl                     string  `json:"callbackUrl" db:"callback_url"`                                           //URL for callbacks
	RollingReservePrc               float64 `json:"rollingReservePrc" db:"rolling_reserve_prc"`                              //Rolling reserve rate (in %), Number
	RollingReserveHoldDays          int     `json:"rollingReserveHoldDays" db:"rolling_reserve_hold_days"`                   //Rolling reserve hold days
	SendCallback                    bool    `json:"sendCallback" db:"send_callback"`                                         //Send callbacks for merchant
	AcceptUndefinedProvisionChannel bool    `json:"acceptUndefinedProvisionChannel" db:"accept_undefined_provision_channel"` //Accept undefined provision channels
	AllowDuplicateOrderExternalId   bool    `json:"allowDuplicateOrderExternalId" db:"allow_duplicate_order_external_id"`    //Allow duplicate order external ID
	AllowNotificationsForSeller     bool    `json:"allowNotificationsForSeller" db:"allow_notifications_for_seller"`         //Send notifications for seller
	AllowNotificationsForBuyer      bool    `json:"allowNotificationsForBuyer" db:"allow_notifications_for_buyer"`           //Send notifications for buyer
	AllowPartialPayments            bool    `json:"allowPartialPayments" db:"allow_partial_payments"`                        //Allow partial payments
}
```

Company Company object (for merchant only) swagger:model

#### type Config

```go
type Config struct {
	BaseURL   string
	CommonURL string
	Client    string //partner API client uid
	Secret    string //partner API secret
}
```


#### type Country

```go
type Country struct {
	Id int64 `json:"internal_id,omitempty" db:"id"` //Internal for user-integration
	//ITH.AMS data structure
	AmsId                 int64  `json:"id,omitempty" db:"ams_id"`
	Code                  string `json:"code" db:"code"`
	Name                  string `json:"name" db:"name"`
	BrandedCardsAvailable bool   `json:"brandedCardsAvailable,omitempty" db:"branded_cards_available"`
	RegistrationAllowed   bool   `json:"registrationAllowed,omitempty" db:"registration_allowed"`
}
```

Additional field `country`

      "country": {
    		"code": "LV",
    		"name": "Latvia",
    		"brandedCardsAvailable": true
    	},

#### type ErrorData

```go
type ErrorData struct {
	ErrorCode    int         `json:"errorCode"`    //Error code
	ErrorMessage string      `json:"errorMessage"` //Localized error message. Supported languages are English, Russian, and Latvian. English is used	when no customer locale is available
	RequestUid   string      `json:"requestUid"`   //Request UID, used for investigation of exceptional cases
	Parameters   interface{} `json:"parameters"`   //Error extended parameters
}
```

ErrorData - any response

#### type Language

```go
type Language struct {
	Uid     string `json:"uid,omitempty" db:"uid"`
	Code    string `json:"code" db:"code"` //ISO2 language code
	Name    string `json:"name,omitempty" db:"name"`
	Type    string `json:"type,omitempty" db:"type"`
	Primary bool   `json:"primary,omitempty" db:"primary"`
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
	PhoneTypeMobile PhoneType = iota + 1 //Mobile
	PhoneTypeHome                        //Home
	PhoneTypeWork                        //Work
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

#### type RequestStatus

```go
type RequestStatus int
```


```go
const (
	RequestStatusNone RequestStatus = iota
	RequestStatusValidationError
	RequestStatusDbError
	RequestStatusNetworkError
	RequestStatusPartnerError
	RequestStatusOk
	UpdateValidationError
	UpdateDbError
	UpdateNetworkError
	UpdatePartnerError
)
```

#### func (RequestStatus) MarshalJSON

```go
func (r RequestStatus) MarshalJSON() ([]byte, error)
```
MarshalJSON is generated so RequestStatus satisfies json.Marshaler.

#### func (*RequestStatus) Scan

```go
func (r *RequestStatus) Scan(src interface{}) error
```
Value is generated so RequestStatus satisfies db row driver.Scanner.

#### func (RequestStatus) String

```go
func (r RequestStatus) String() string
```
String is generated so RequestStatus satisfies fmt.Stringer.

#### func (*RequestStatus) UnmarshalJSON

```go
func (r *RequestStatus) UnmarshalJSON(data []byte) error
```
UnmarshalJSON is generated so RequestStatus satisfies json.Unmarshaler.

#### func (RequestStatus) Validate

```go
func (r RequestStatus) Validate() error
```
Validate verifies that value is predefined for RequestStatus.

#### func (RequestStatus) Value

```go
func (r RequestStatus) Value() (driver.Value, error)
```
Value is generated so RequestStatus satisfies db row driver.Valuer.

#### type UserRegistrationRequest

```go
type UserRegistrationRequest struct {
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

#### func (*UserRegistrationRequest) Validate

```go
func (r *UserRegistrationRequest) Validate() (err error)
```
Validate verifies that value is predefined for AffiliateInfo.

#### type UserRegistrationResponse

```go
type UserRegistrationResponse struct {
	ErrorData          *ErrorData `json:"errorData,omitempty"` //null if OK
	AccountUid         string     `json:"accountUid"`
	ExternalAccountUid string     `json:"externalAccountUid"`
	AccessToken        string     `json:"accessToken"`
	Account            *Account   `json:"account"`
}
```

UserRegistrationResponse response from ITH Account Management Services (AMS)

     {
    	"accountUid": "100-014-275-55",
    	"externalAccountUid": "EX-ACC-UID-1234",
    	"accessToken": "bdad264b7f8b9896d73436b234e4bddd",
    	"account": {....}
     }

#### type UserUpdateRequest

```go
type UserUpdateRequest struct {
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

#### func (UserUpdateRequest) Validate

```go
func (r UserUpdateRequest) Validate() (err error)
```

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
	WebResourceWeb      WebResource = iota + 1 // W – Web;
	WebResourceFacebook                        //F – Facebook;
	WebResourceTwitter                         //T – Twitter
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
