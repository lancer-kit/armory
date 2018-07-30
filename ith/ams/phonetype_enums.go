package ams

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

// PhoneType:
//
//* M - Mobile
//
//* H - Home
//
//* W - Work
type PhoneType int

const (
	PhoneTypeMobile PhoneType = iota + 1 //Mobile
	PhoneTypeHome                        //Home
	PhoneTypeWork                        //Work
)

var ErrPhoneTypeInvalid = errors.New("PhoneType is invalid")

var defPhoneTypeNameToValue = map[string]PhoneType{
	"M": PhoneTypeMobile,
	"H": PhoneTypeHome,
	"W": PhoneTypeWork,
}

var defPhoneTypeValueToName = map[PhoneType]string{
	PhoneTypeMobile: "M",
	PhoneTypeHome:   "H",
	PhoneTypeWork:   "W",
}

// String is generated so PhoneType satisfies fmt.Stringer.
func (r PhoneType) String() string {
	s, ok := defPhoneTypeValueToName[r]
	if !ok {
		return fmt.Sprintf("PhoneType(%d)", r)
	}
	return s
}

// Validate verifies that value is predefined for PhoneType.
func (r PhoneType) Validate() error {
	_, ok := defPhoneTypeValueToName[r]
	if !ok {
		return ErrPhoneTypeInvalid
	}
	return nil
}

// MarshalJSON is generated so PhoneType satisfies json.Marshaler.
func (r PhoneType) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(r).(fmt.Stringer); ok {
		return json.Marshal(s.String())
	}
	s, ok := defPhoneTypeValueToName[r]
	if !ok {
		return nil, fmt.Errorf("PhoneType(%d) is invalid value", r)
	}
	return json.Marshal(s)
}

// UnmarshalJSON is generated so PhoneType satisfies json.Unmarshaler.
func (r *PhoneType) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("PhoneType: should be a string, got %s", string(data))
	}
	v, ok := defPhoneTypeNameToValue[s]
	if !ok {
		return fmt.Errorf("PhoneType(%q) is invalid value", s)
	}
	*r = v
	return nil
}

// Value is generated so PhoneType satisfies db row driver.Valuer.
func (r PhoneType) Value() (driver.Value, error) {
	j, err := json.Marshal(r)
	return j, err
}

// Value is generated so PhoneType satisfies db row driver.Scanner.
func (r *PhoneType) Scan(src interface{}) error {
	switch src.(type) {
	case string:
		val, ok := defPhoneTypeNameToValue[src.(string)]
		if !ok {
			return errors.New("PhoneType: can't unmarshal column data")
		}
		*r = val
		return nil
	case []byte:
		source := src.([]byte)
		var i PhoneType
		err := json.Unmarshal(source, &i)
		if err != nil {
			return errors.New("PhoneType: can't unmarshal column data")
		}

		*r = i
		return nil
	case int, int8, int32, int64, uint, uint8, uint32, uint64:
		ni := sql.NullInt64{}
		err := ni.Scan(src)
		if err != nil {
			return errors.New("PhoneType: can't scan column data into int64")
		}

		*r = PhoneType(ni.Int64)
		return nil
	}
	return errors.New("PhoneType: invalid type")
}
