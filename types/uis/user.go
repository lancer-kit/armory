package uis

import (
	"errors"
	"time"
	"fmt"
)

type User struct {
	Id                int64         `json:"id"`
	Phone             string        `json:"phone"`
	Email             string        `json:"email"`
	Status            UserStatus    `json:"status"`
	FirstName         string        `json:"firstName"`
	LastName          string        `json:"lastName"`
	BirthDate         UserBirthDate `json:"birthDate"`
	LanguageMarker    string        `json:"languageMarker"`
	CountryMarker     string        `json:"countryMarker"`
	PreferredCurrency string        `json:"preferredCurrency"`
	MailVerified      bool          `json:"mailVerified"`
	UserKey           string        `json:"-"`
	CreatedAt         int64         `json:"createdAt"`
	UpdatedAt         int64         `json:"updatedAt"`
}

type UserStatus int
type UserBirthDate int64

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

func (t UserBirthDate) String() string {
	tmp := time.Unix(int64(t), 0)
	return tmp.Format("2006-01-02")
}

func (t *UserBirthDate) ToIthAmsString() string {
	r := time.Unix(int64(*t), 0)
	y, m, d := time.Time(r).Date()
	return fmt.Sprintf("%04d%02d%02d", y, m, d) + "000000"
}
