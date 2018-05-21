package ams

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

// AccountType:
//  * S - Standard
//  * M - Merchant
//  * B - Business
type AccountType int

const (
	AccountTypeStandard AccountType = iota //Standard
	AccountTypeMerchant                    //Merchant
	AccountTypeBusiness                    //Business
)
var ErrAccountTypeInvalid = errors.New("AccountType is invalid")

var defAccountTypeNameToValue = map[string]AccountType{
	"S": AccountTypeStandard,
	"M": AccountTypeMerchant,
	"B": AccountTypeBusiness,
}

var defAccountTypeValueToName = map[AccountType]string{
	AccountTypeStandard: "S",
	AccountTypeMerchant: "M",
	AccountTypeBusiness: "B",
}

// String is generated so AccountType satisfies fmt.Stringer.
func (r AccountType) String() string {
	s, ok := defAccountTypeValueToName[r]
	if !ok {
		return fmt.Sprintf("AccountType(%d)", r)
	}
	return s
}

// Validate verifies that value is predefined for AccountType.
func (r AccountType) Validate() error {
	_, ok := defAccountTypeValueToName[r]
	if !ok {
		return ErrAccountTypeInvalid
	}
	return nil
}

// MarshalJSON is generated so AccountType satisfies json.Marshaler.
func (r AccountType) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(r).(fmt.Stringer); ok {
		return json.Marshal(s.String())
	}
	s, ok := defAccountTypeValueToName[r]
	if !ok {
		return nil, fmt.Errorf("AccountType(%d) is invalid value", r)
	}
	return json.Marshal(s)
}

// UnmarshalJSON is generated so AccountType satisfies json.Unmarshaler.
func (r *AccountType) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("AccountType: should be a string, got %s", string(data))
	}
	v, ok := defAccountTypeNameToValue[s]
	if !ok {
		return fmt.Errorf("AccountType(%q) is invalid value", s)
	}
	*r = v
	return nil
}

// Value is generated so AccountType satisfies db row driver.Valuer.
func (r AccountType) Value() (driver.Value, error) {
	j, err := json.Marshal(r)
	return j, err
}

// Value is generated so AccountType satisfies db row driver.Scanner.
func (r *AccountType) Scan(src interface{}) error {
	switch src.(type) {
	case string:
		val, ok := defAccountTypeNameToValue[src.(string)]
		if !ok {
			return errors.New("AccountType: can't unmarshal column data")
		}
		*r = val
		return nil
	case []byte:
		source := src.([]byte)
		var i AccountType
		err := json.Unmarshal(source, &i)
		if err != nil {
			return errors.New("AccountType: can't unmarshal column data")
		}

		*r = i
		return nil
	case int, int8, int32, int64, uint, uint8, uint32, uint64:
		ni := sql.NullInt64{}
		err := ni.Scan(src)
		if err != nil {
			return errors.New("AccountType: can't scan column data into int64")
		}

		*r = AccountType(ni.Int64)
		return nil
	}
	return errors.New("AccountType: invalid type")
}
