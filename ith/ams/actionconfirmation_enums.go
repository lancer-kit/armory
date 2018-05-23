package ams

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

//Field type for Account.ActionConfirmationType
//
//  * "EMAIL" – via email;
//  * "SMS" – via phone
//  * "GAUTH" – via Google Authenticator
type ActionConfirmation int

var ErrActionConfirmationInvalid = errors.New("ActionConfirmation is invalid")

const (
	ActionConfirmationEmail ActionConfirmation = iota + 1 //EMAIL – via email;
	ActionConfirmationSms                                 //SMS – via phone
	ActionConfirmationGAuth                               //GAUTH – via Google Authenticator
)

var defActionConfirmationNameToValue = map[string]ActionConfirmation{
	"EMAIL": ActionConfirmationEmail,
	"SMS":   ActionConfirmationSms,
	"GAUTH": ActionConfirmationGAuth,
}

var defActionConfirmationValueToName = map[ActionConfirmation]string{
	ActionConfirmationEmail: "EMAIL",
	ActionConfirmationSms:   "SMS",
	ActionConfirmationGAuth: "GAUTH",
}

// String is generated so ActionConfirmation satisfies fmt.Stringer.
func (r ActionConfirmation) String() string {
	s, ok := defActionConfirmationValueToName[r]
	if !ok {
		return fmt.Sprintf("ActionConfirmation(%d)", r)
	}
	return s
}

// Validate verifies that value is predefined for ActionConfirmation.
func (r ActionConfirmation) Validate() error {
	_, ok := defActionConfirmationValueToName[r]
	if !ok {
		return ErrActionConfirmationInvalid
	}
	return nil
}

// MarshalJSON is generated so ActionConfirmation satisfies json.Marshaler.
func (r ActionConfirmation) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(r).(fmt.Stringer); ok {
		return json.Marshal(s.String())
	}
	s, ok := defActionConfirmationValueToName[r]
	if !ok {
		return nil, fmt.Errorf("ActionConfirmation(%d) is invalid value", r)
	}
	return json.Marshal(s)
}

// UnmarshalJSON is generated so ActionConfirmation satisfies json.Unmarshaler.
func (r *ActionConfirmation) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("ActionConfirmation: should be a string, got %s", string(data))
	}
	v, ok := defActionConfirmationNameToValue[s]
	if !ok {
		return fmt.Errorf("ActionConfirmation(%q) is invalid value", s)
	}
	*r = v
	return nil
}

// Value is generated so ActionConfirmation satisfies db row driver.Valuer.
func (r ActionConfirmation) Value() (driver.Value, error) {
	j, err := json.Marshal(r)
	return j, err
}

// Value is generated so ActionConfirmation satisfies db row driver.Scanner.
func (r *ActionConfirmation) Scan(src interface{}) error {
	switch src.(type) {
	case string:
		val, ok := defActionConfirmationNameToValue[src.(string)]
		if !ok {
			return errors.New("ActionConfirmation: can't unmarshal column data")
		}
		*r = val
		return nil
	case []byte:
		source := src.([]byte)
		var i ActionConfirmation
		err := json.Unmarshal(source, &i)
		if err != nil {
			return errors.New("ActionConfirmation: can't unmarshal column data")
		}

		*r = i
		return nil
	case int, int8, int32, int64, uint, uint8, uint32, uint64:
		ni := sql.NullInt64{}
		err := ni.Scan(src)
		if err != nil {
			return errors.New("ActionConfirmation: can't scan column data into int64")
		}

		*r = ActionConfirmation(ni.Int64)
		return nil
	}
	return errors.New("ActionConfirmation: invalid type")
}
