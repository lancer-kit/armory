package uis

import (
	"errors"
	"fmt"
	"time"

	"github.com/go-ozzo/ozzo-validation"
	"gitlab.inn4science.com/vcg/go-common/ith/ams"
)

type (
	//swagger:model
	User struct {
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
		//Integrations      IntegrationMap `json:"integrations,omitempty" db:"-"`             //list of integration-specific data
	}
	//alias to ams.AddressRequest
	Address = ams.AddressRequest

	//Structure to process register and update
	//swagger:model
	UserRequest struct {
		User    *User    `json:"user"`
		Address *Address `json:"address"`
	}

	//IntegrationMap map[string]interface{}
	UserStatus    int
	UserBirthDate int64

	//swagger:model
	UpdateResult struct {
		Request    *ams.UserUpdateRequest //User request with fields filled to be updated
		Token      string                 //User token to process request from ams.API.UpdateProfile()
		NeedUpdate bool                   //True, when any filed in Request not empty
	}
)

const (
	UserStatusBlocked        UserStatus = -10
	UserStatusNew            UserStatus = 0
	UserStatusPhoneVerified  UserStatus = 20
	UserStatusAuthorized     UserStatus = 30
	UserStatusOnVerification UserStatus = 40
	UserStatusVerified       UserStatus = 50
)

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

func (t *User) Validate() error {
	return validation.ValidateStruct(t,
		validation.Field(&t.Id, validation.Required),
		validation.Field(&t.Phone, validation.Required, validation.Length(6, 20)),
		validation.Field(&t.Email, validation.Required),
		validation.Field(&t.Status, validation.Required),
		validation.Field(&t.FirstName, validation.Required),
		validation.Field(&t.LastName, validation.Required),
		validation.Field(&t.LanguageMarker, validation.Required),
		validation.Field(&t.CountryMarker, validation.Required),
		validation.Field(&t.PreferredCurrency, validation.Required),
		validation.Field(&t.UserKey, validation.Required),
	)
}

func (t *User) GetStatusName() (string, error) {
	if v, ok := StatusName[t.Status]; ok {
		return v, nil
	}
	return "", ErrorInvalidStatus
}

func (t UserStatus) String() string {
	if v, ok := StatusName[t]; ok {
		return v
	}
	return ""
}

func (t *UserBirthDate) FromSQLDate(s string) (d UserBirthDate, err error) {
	tmp, ce := time.Parse("2006-01-02", s)
	if ce != nil {
		err = errors.New(ErrorInvalidDate + ce.Error())
		return
	}
	d = UserBirthDate(tmp.UTC().Unix())
	if t != nil {
		*t = d
	}
	return
}

//UserBirthDate.String - Stringer interface
func (t UserBirthDate) String() string {
	tmp := time.Unix(int64(t), 0)
	return tmp.Format("2006-01-02")
}

//ToIthAmsString - convert to ITH AMS string presentation
func (t *UserBirthDate) ToIthAmsString() string {
	r := time.Unix(int64(*t), 0)
	y, m, d := time.Time(r).Date()
	return fmt.Sprintf("%04d%02d%02d", y, m, d) + "000000"
}

//ToAmsDate - convert UserBirthDate to ams.AmsDate
func (t UserBirthDate) ToAmsDate() ams.AmsDate {
	return ams.AmsDate(time.Unix(int64(t), 0))
}
