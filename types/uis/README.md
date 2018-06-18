# uis
--
    import "gitlab.inn4science.com/vcg/go-common/types/uis"


## Usage

```go
const SqlTimeLayout = "2006-01-02 15:04:05"
```

```go
var (
	StatusName = map[UserStatus]string{
		UserStatusBlocked:        "user blocked",
		UserStatusNew:            "new",
		UserStatusPhoneVerified:  "phone number verified",
		UserStatusAuthorized:     "user authorized",
		UserStatusOnVerification: "user documents on verification",
		UserStatusVerified:       "user phone and documents verified",
	}
	ErrorInvalidStatus = errors.New("invalid user status")
	ErrorInvalidDate   = "invalid date: "
)
```

#### func  GetBoolPtr

```go
func GetBoolPtr(v *bool) bool
```

#### type Address

```go
type Address = ams.AddressRequest
```

alias to ams.AddressRequest

#### type IthUserData

```go
type IthUserData struct {
	Id                        int64                  `json:"id"`                                                //User ID.Same in user-api service
	AccountId                 string                 `json:"accountId" db:"account_id"`                         //External account uid
	AccessToken               string                 `json:"accessToken" db:"access_token"`                     //Refresh token
	AccountType               ams.AccountType        `json:"accountType" db:"account_type"`                     //See ams.AccountType
	AccountStatus             ams.AccountStatus      `json:"accountStatus" db:"account_status"`                 //See ams.AccountStatus
	AffilateId                string                 `json:"affilateId" db:"affiliate_id"`                      //Affilate id (referral)  in ITH AMS
	Language                  string                 `json:"language" db:"language"`                            //String(2). User language ISO2 code
	CommunicationLanguage     string                 `json:"communicationLanguage" db:"communication_language"` //String(2). Communication language
	CampaignId                string                 `json:"campaignId" db:"campaign_id"`                       //Campaign Id (referral)  in ITH AMS
	BannerId                  string                 `json:"bannerId" db:"banner_id"`                           //Banner Id (referral) in ITH AMS
	CustomParameters          string                 `json:"customParameters" db:"custom_parameters"`           //Custom parameters in ITH AMS
	AccountSecret             string                 `json:"accountSecret" db:"account_secret"`                 //Merchant account in ITH AMS
	MerchantUid               string                 `json:"merchantUid" db:"merchant_uid"`                     //Merchant UID in ITH AMS
	AccountTimezone           int                    `json:"accountTimezone" db:"account_timezone"`             //Timezone
	AccountWeekStartsOn       string                 `json:"accountWeekStartsOn" db:"account_week_starts_on"`
	CurrencyConversion        bool                   `json:"currencyConversion" db:"currency_conversion"`
	AlwaysRefundEwallet       bool                   `json:"alwaysRefundEwallet" db:"always_refund_ewallet"`
	ConfirmOutTransaction     bool                   `json:"confirmOutTransaction" db:"confirm_out_transaction"`
	ConfirmLogin              bool                   `json:"confirmLogin" db:"confirm_login"`
	ActionConfirmationEnabled bool                   `json:"actionConfirmationEnabled" db:"action_confirmation_enabled"`
	ActionConfirmationType    ams.ActionConfirmation `json:"actionConfirmationType" db:"action_confirmation_type"`
	Test                      *bool                  `json:"test,omitempty" db:"test"`
}
```


#### type RegistrationStatus

```go
type RegistrationStatus struct {
	UserId       int64             `json:"userId"`
	IsRegistered bool              `json:"isRegistered"`
	Status       ams.AccountStatus `json:"status"`
}
```


#### type RequestState

```go
type RequestState struct {
	Id          int64             `json:"id" db:"id"`
	UserId      int64             `json:"-" db:"user_id"`
	State       ams.RequestStatus `json:"state" db:"state"`
	Data        string            `json:"data,omitempty" db:"data"`
	CreatedAt   *SqlTime          `json:"createdAt,omitempty" db:"created_at"`
	ProcessedAt *SqlTime          `json:"processedAt,omitempty" db:"processed_at"`
}
```


#### type SqlTime

```go
type SqlTime int64
```


#### func (*SqlTime) FromString

```go
func (t *SqlTime) FromString(s string) error
```

#### func (*SqlTime) Scan

```go
func (t *SqlTime) Scan(src interface{}) error
```

#### func (*SqlTime) String

```go
func (t *SqlTime) String() string
```

#### func (*SqlTime) ToTime

```go
func (t *SqlTime) ToTime() time.Time
```

#### type UpdateResult

```go
type UpdateResult struct {
	Request    *ams.UserUpdateRequest //User request with fields filled to be updated
	Token      string                 //User token to process request from ams.API.UpdateProfile()
	NeedUpdate bool                   //True, when any filed in Request not empty
}
```


#### type User

```go
type User struct {
	Id                int64         `json:"id" db:"id"`                                //user id (same to user-api)
	Phone             string        `json:"phone" db:"phone"`                          //user phone
	Email             string        `json:"email" db:"email"`                          //user email
	Status            UserStatus    `json:"status" db:"status"`                        //status on user-api service
	FirstName         string        `json:"firstName" db:"first_name"`                 //Fist name
	LastName          string        `json:"lastName" db:"last_name"`                   //Last name
	BirthDate         UserBirthDate `json:"birthDate" db:"birth_date"`                 //Bith date
	LanguageMarker    string        `json:"languageMarker" db:"language_marker"`       //ISO 2 user language
	CountryMarker     string        `json:"countryMarker" db:"country_marker"`         //ISO 2 Country
	PreferredCurrency string        `json:"preferredCurrency" db:"preferred_currency"` //ISO 3 User preferred currency
	MailVerified      bool          `json:"mailVerified" db:"mail_verified"`           //Flag - is mail verified on user-api service
	UserKey           string        `json:"userKey" db:"user_key"`                     //User password hash
	CreatedAt         int64         `json:"createdAt" db:"created_at"`                 //Create at, unix time stamp
	UpdatedAt         int64         `json:"updatedAt" db:"updated_at"`                 //Updated at, unix time stamp

}
```


#### func (*User) GetStatusName

```go
func (t *User) GetStatusName() (string, error)
```

#### func (*User) Validate

```go
func (t *User) Validate() error
```

#### type UserAuth

```go
type UserAuth struct {
	UserId      int64             `json:"userId"`      //User ID
	AccessToken string            `json:"accessToken"` //Token to process data from/to ITH
	Status      ams.AccountStatus `json:"status"`      //ITH user status (see ams.Account)
}
```

Result of request user-integration service @ /ith/auth Needs to be signed
request (see go-common/auth) Example of request: `GET`
`http:localhost:2094/v1/uis/ith/auth

    Header. jwt:{"jti":"1"}

Response:

     {
    		"userId":1,
    		"accessToken":"2bc23zsd3ffer4g993d"
    		"status": "SC"
    	}

#### type UserBirthDate

```go
type UserBirthDate int64
```


#### func (*UserBirthDate) FromSQLDate

```go
func (t *UserBirthDate) FromSQLDate(s string) (d UserBirthDate, err error)
```

#### func (UserBirthDate) String

```go
func (t UserBirthDate) String() string
```
UserBirthDate.String - Stringer interface

#### func (UserBirthDate) ToAmsDate

```go
func (t UserBirthDate) ToAmsDate() ams.AmsDate
```
ToAmsDate - convert UserBirthDate to ams.AmsDate

#### func (*UserBirthDate) ToIthAmsString

```go
func (t *UserBirthDate) ToIthAmsString() string
```
ToIthAmsString - convert to ITH AMS string presentation

#### type UserRequest

```go
type UserRequest struct {
	User    *User    `json:"user"`
	Address *Address `json:"address"`
}
```

Structure to process register and update

#### type UserStatus

```go
type UserStatus int
```

IntegrationMap map[string]interface{}

```go
const (
	UserStatusBlocked        UserStatus = -10
	UserStatusNew            UserStatus = 0
	UserStatusPhoneVerified  UserStatus = 20
	UserStatusAuthorized     UserStatus = 30
	UserStatusOnVerification UserStatus = 40
	UserStatusVerified       UserStatus = 50
)
```

#### func (UserStatus) String

```go
func (t UserStatus) String() string
```
