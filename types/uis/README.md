# uis
--
    import "gitlab.inn4science.com/vcg/go-common/types/uis"


## Usage

```
const SqlTimeLayout = "2006-01-02 15:04:05"
```

```
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

```
func GetBoolPtr(v *bool) bool
```

#### type Address

```
type Address = ams.AddressRequest
```

alias to ams.AddressRequest

#### type IthUserData

```
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
	CreatedAt                 int64                  `json:"createdAt" db:"created_at"` //Created @, unix timestamp
	UpdatedAt                 int64                  `json:"updatedAt" db:"updated_at"` //Updated @, unix timestamp
}
```

ITH -specific user data swagger:model

#### type RegistrationStatus

```
type RegistrationStatus struct {
	UserId       int64             `json:"userId"`
	IsRegistered bool              `json:"isRegistered"`
	Status       ams.AccountStatus `json:"status"`
}
```

swagger:model

#### type RequestState

```
type RequestState struct {
	Id          int64             `json:"id" db:"id"`
	UserId      int64             `json:"-" db:"user_id"`
	State       ams.RequestStatus `json:"state" db:"state"`
	Data        string            `json:"data,omitempty" db:"data"`
	CreatedAt   *SqlTime          `json:"createdAt,omitempty" db:"created_at"`
	ProcessedAt *SqlTime          `json:"processedAt,omitempty" db:"processed_at"`
}
```

swagger:model

#### type SqlTime

```
type SqlTime int64
```


#### func (*SqlTime) FromString

```
func (t *SqlTime) FromString(s string) error
```

#### func (*SqlTime) Scan

```
func (t *SqlTime) Scan(src interface{}) error
```

#### func (*SqlTime) String

```
func (t *SqlTime) String() string
```

#### func (*SqlTime) ToTime

```
func (t *SqlTime) ToTime() time.Time
```

#### type UpdateResult

```
type UpdateResult struct {
	Request    *ams.UserUpdateRequest //User request with fields filled to be updated
	Token      string                 //User token to process request from ams.API.UpdateProfile()
	NeedUpdate bool                   //True, when any filed in Request not empty
}
```

swagger:model

#### type User

```
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
	UserKey           string        `json:"userKey" db:"-"`                            //User password (plain, not for save) for create new profile on ITH.AMS
	CreatedAt         int64         `json:"createdAt" db:"created_at"`                 //Create at, unix time stamp
	UpdatedAt         int64         `json:"updatedAt" db:"updated_at"`                 //Updated at, unix time stamp

}
```

swagger:model

#### func (*User) GetStatusName

```
func (t *User) GetStatusName() (string, error)
```

#### func (*User) Validate

```
func (t *User) Validate() error
```

#### type UserAuth

```
type UserAuth struct {
	UserId      int64             `json:"userId"`      //User ID
	AccessToken string            `json:"accessToken"` //Token to process data from/to ITH
	Status      ams.AccountStatus `json:"status"`      //ITH user status (see ams.Account)
}
```

Result of request user-integration service @ /ith/auth Needs to be signed
request (see go-common/auth)

Example of request: `GET` `http:localhost:2094/v1/uis/ith/auth

    Header. jwt:{"jti":"1"}

Response:

     {
    		"userId":1,
    		"accessToken":"2bc23zsd3ffer4g993d"
    		"status": "SC"
    	}

#### type UserBirthDate

```
type UserBirthDate int64
```


#### func (*UserBirthDate) FromSQLDate

```
func (t *UserBirthDate) FromSQLDate(s string) (d UserBirthDate, err error)
```

#### func (UserBirthDate) String

```
func (t UserBirthDate) String() string
```
UserBirthDate.String - Stringer interface

#### func (UserBirthDate) ToAmsDate

```
func (t UserBirthDate) ToAmsDate() ams.AmsDate
```
ToAmsDate - convert UserBirthDate to ams.AmsDate

#### func (UserBirthDate) ToAmsDatePtr

```
func (t UserBirthDate) ToAmsDatePtr() *ams.AmsDate
```
ToAmsDate - convert UserBirthDate to ams.AmsDate

#### func (*UserBirthDate) ToIthAmsString

```
func (t *UserBirthDate) ToIthAmsString() string
```
ToIthAmsString - convert to ITH AMS string presentation

#### type UserRequest

```
type UserRequest struct {
	User    *User    `json:"user"`
	Address *Address `json:"address"`
}
```

Structure to process register and update swagger:model

#### type UserStatus

```
type UserStatus int
```

IntegrationMap map[string]interface{}

```
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

```
func (t UserStatus) String() string
```
