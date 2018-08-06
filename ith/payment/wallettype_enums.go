// generated by goplater -type=WalletType -transform=snake -tprefix=false; DO NOT EDIT
package payment

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

var ErrWalletTypeInvalid = errors.New("WalletType is invalid")

var defWalletTypeNameToValue = map[string]WalletType{
	"s":  WalletTypeS,
	"r":  WalletTypeR,
	"ac": WalletTypeAC,
}

var defWalletTypeValueToName = map[WalletType]string{
	WalletTypeS:  "s",
	WalletTypeR:  "r",
	WalletTypeAC: "ac",
}

// String is generated so WalletType satisfies fmt.Stringer.
func (r WalletType) String() string {
	s, ok := defWalletTypeValueToName[r]
	if !ok {
		return fmt.Sprintf("WalletType(%d)", r)
	}
	return s
}

// Validate verifies that value is predefined for WalletType.
func (r WalletType) Validate() error {
	_, ok := defWalletTypeValueToName[r]
	if !ok {
		return ErrWalletTypeInvalid
	}
	return nil
}

// MarshalJSON is generated so WalletType satisfies json.Marshaler.
func (r WalletType) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(r).(fmt.Stringer); ok {
		return json.Marshal(s.String())
	}
	s, ok := defWalletTypeValueToName[r]
	if !ok {
		return nil, fmt.Errorf("WalletType(%d) is invalid value", r)
	}
	return json.Marshal(s)
}

// UnmarshalJSON is generated so WalletType satisfies json.Unmarshaler.
func (r *WalletType) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("WalletType: should be a string, got %s", string(data))
	}
	v, ok := defWalletTypeNameToValue[s]
	if !ok {
		return fmt.Errorf("WalletType(%q) is invalid value", s)
	}
	*r = v
	return nil
}

// Value is generated so WalletType satisfies db row driver.Valuer.
func (r WalletType) Value() (driver.Value, error) {
	s, ok := defWalletTypeValueToName[r]
	if !ok {
		return "", nil
	}
	return s, nil
}

// Value is generated so WalletType satisfies db row driver.Scanner.
func (r *WalletType) Scan(src interface{}) error {
	switch src.(type) {
	case string:
		val, _ := defWalletTypeNameToValue[src.(string)]
		*r = val
		return nil
	case []byte:
		source := src.([]byte)
		var i WalletType
		err := json.Unmarshal(source, &i)
		if err != nil {
			return errors.New("WalletType: can't unmarshal column data")
		}

		*r = i
		return nil
	case int, int8, int32, int64, uint, uint8, uint32, uint64:
		ni := sql.NullInt64{}
		err := ni.Scan(src)
		if err != nil {
			return errors.New("WalletType: can't scan column data into int64")
		}

		*r = WalletType(ni.Int64)
		return nil
	}
	return errors.New("WalletType: invalid type")
}
