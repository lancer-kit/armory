package ams

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

//AddressType:
//	* B – Business;
//	* H – Home;
//	* O – Other;
//	* C – Communication
type AddressType int

const (
	AddressTypeBusiness      AddressType = iota + 1 //B – Business;
	AddressTypeHome                                 //H – Home;
	AddressTypeOther                                //O – Other;
	AddressTypeCommunication                        //C – Communication

)

var ErrAddressTypeInvalid = errors.New("AddressType is invalid")

var defAddressTypeNameToValue = map[string]AddressType{
	"B": AddressTypeBusiness,
	"H": AddressTypeHome,
	"O": AddressTypeOther,
	"C": AddressTypeCommunication,
}

var defAddressTypeValueToName = map[AddressType]string{
	AddressTypeBusiness:      "B",
	AddressTypeHome:          "H",
	AddressTypeOther:         "O",
	AddressTypeCommunication: "C",
}

// String is generated so AddressType satisfies fmt.Stringer.
func (r AddressType) String() string {
	s, ok := defAddressTypeValueToName[r]
	if !ok {
		return fmt.Sprintf("AddressType(%d)", r)
	}
	return s
}

// Validate verifies that value is predefined for AddressType.
func (r AddressType) Validate() error {
	_, ok := defAddressTypeValueToName[r]
	if !ok {
		return ErrAddressTypeInvalid
	}
	return nil
}

// MarshalJSON is generated so AddressType satisfies json.Marshaler.
func (r AddressType) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(r).(fmt.Stringer); ok {
		return json.Marshal(s.String())
	}
	s, ok := defAddressTypeValueToName[r]
	if !ok {
		return nil, fmt.Errorf("AddressType(%d) is invalid value", r)
	}
	return json.Marshal(s)
}

// UnmarshalJSON is generated so AddressType satisfies json.Unmarshaler.
func (r *AddressType) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("AddressType: should be a string, got %s", string(data))
	}
	v, ok := defAddressTypeNameToValue[s]
	if !ok {
		return fmt.Errorf("AddressType(%q) is invalid value", s)
	}
	*r = v
	return nil
}

// Value is generated so AddressType satisfies db row driver.Valuer.
func (r AddressType) Value() (driver.Value, error) {
	j, err := json.Marshal(r)
	return j, err
}

// Value is generated so AddressType satisfies db row driver.Scanner.
func (r *AddressType) Scan(src interface{}) error {
	switch src.(type) {
	case string:
		val, ok := defAddressTypeNameToValue[src.(string)]
		if !ok {
			return errors.New("AddressType: can't unmarshal column data")
		}
		*r = val
		return nil
	case []byte:
		source := src.([]byte)
		var i AddressType
		err := json.Unmarshal(source, &i)
		if err != nil {
			return errors.New("AddressType: can't unmarshal column data")
		}

		*r = i
		return nil
	case int, int8, int32, int64, uint, uint8, uint32, uint64:
		ni := sql.NullInt64{}
		err := ni.Scan(src)
		if err != nil {
			return errors.New("AddressType: can't scan column data into int64")
		}

		*r = AddressType(ni.Int64)
		return nil
	}
	return errors.New("AddressType: invalid type")
}
