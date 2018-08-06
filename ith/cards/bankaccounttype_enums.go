// generated by goplater -type=BankAccountType -transform=none -tprefix=false; DO NOT EDIT
package cards

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

var ErrBankAccountTypeInvalid = errors.New("BankAccountType is invalid")

var defBankAccountTypeNameToValue = map[string]BankAccountType{
	"I": BankAccountTypeI,
	"E": BankAccountTypeE,
}

var defBankAccountTypeValueToName = map[BankAccountType]string{
	BankAccountTypeI: "I",
	BankAccountTypeE: "E",
}

// String is generated so BankAccountType satisfies fmt.Stringer.
func (r BankAccountType) String() string {
	s, ok := defBankAccountTypeValueToName[r]
	if !ok {
		return fmt.Sprintf("BankAccountType(%d)", r)
	}
	return s
}

// Validate verifies that value is predefined for BankAccountType.
func (r BankAccountType) Validate() error {
	_, ok := defBankAccountTypeValueToName[r]
	if !ok {
		return ErrBankAccountTypeInvalid
	}
	return nil
}

// MarshalJSON is generated so BankAccountType satisfies json.Marshaler.
func (r BankAccountType) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(r).(fmt.Stringer); ok {
		return json.Marshal(s.String())
	}
	s, ok := defBankAccountTypeValueToName[r]
	if !ok {
		return nil, fmt.Errorf("BankAccountType(%d) is invalid value", r)
	}
	return json.Marshal(s)
}

// UnmarshalJSON is generated so BankAccountType satisfies json.Unmarshaler.
func (r *BankAccountType) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("BankAccountType: should be a string, got %s", string(data))
	}
	v, ok := defBankAccountTypeNameToValue[s]
	if !ok {
		return fmt.Errorf("BankAccountType(%q) is invalid value", s)
	}
	*r = v
	return nil
}

// Value is generated so BankAccountType satisfies db row driver.Valuer.
func (r BankAccountType) Value() (driver.Value, error) {
	s, ok := defBankAccountTypeValueToName[r]
	if !ok {
		return "", nil
	}
	return s, nil
}

// Value is generated so BankAccountType satisfies db row driver.Scanner.
func (r *BankAccountType) Scan(src interface{}) error {
	switch src.(type) {
	case string:
		val, _ := defBankAccountTypeNameToValue[src.(string)]
		*r = val
		return nil
	case []byte:
		source := src.([]byte)
		var i BankAccountType
		err := json.Unmarshal(source, &i)
		if err != nil {
			return errors.New("BankAccountType: can't unmarshal column data")
		}

		*r = i
		return nil
	case int, int8, int32, int64, uint, uint8, uint32, uint64:
		ni := sql.NullInt64{}
		err := ni.Scan(src)
		if err != nil {
			return errors.New("BankAccountType: can't scan column data into int64")
		}

		*r = BankAccountType(ni.Int64)
		return nil
	}
	return errors.New("BankAccountType: invalid type")
}
