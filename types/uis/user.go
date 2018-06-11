package uis

import (
	"errors"
	"fmt"
	"time"

	"gitlab.inn4science.com/vcg/go-common/ith/ams"
)

type (
	User struct {
		Id                int64          `json:"id" db:"id"`
		Phone             string         `json:"phone" db:"phone"`
		Email             string         `json:"email" db:"email"`
		Status            UserStatus     `json:"status" db:"status"`
		FirstName         string         `json:"firstName" db:"first_name"`
		LastName          string         `json:"lastName" db:"last_name"`
		BirthDate         UserBirthDate  `json:"birthDate" db:"birth_date"`
		LanguageMarker    string         `json:"languageMarker" db:"language_marker"`
		CountryMarker     string         `json:"countryMarker" db:"country_marker"`
		PreferredCurrency string         `json:"preferredCurrency" db:"preferred_currency"`
		MailVerified      bool           `json:"mailVerified" db:"mail_verified"`
		UserKey           string         `json:"userKey" db:"user_key"`
		CreatedAt         int64          `json:"createdAt" db:"created_at"`
		UpdatedAt         int64          `json:"updatedAt" db:"updated_at"`
		Integrations      IntegrationMap `json:"integrations,omitempty" db:"-"`
	}

	UserRequest struct {
		User    *User               `json:"user"`
		Address *ams.AddressRequest `json:"address"`
	}

	IntegrationMap map[string]interface{}
	UserStatus     int
	UserBirthDate  int64
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
