// generated by goplater -type=TxStatus -transform=snake -tprefix=false; DO NOT EDIT
package payment

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

var (
	ErrTxStatusInvalid = errors.New("TxStatus is invalid")

	defTxStatusNameToValue = map[string]TxStatus{
		"S": TxStatusSuccessful,
		"P": TxStatusPending,
		"F": TxStatusFailed,
	}

	defTxStatusValueToName = map[TxStatus]string{
		TxStatusSuccessful: "S",
		TxStatusPending:    "P",
		TxStatusFailed:     "F",
	}
)

// String is generated so TxStatus satisfies fmt.Stringer.
func (r TxStatus) String() string {
	s, ok := defTxStatusValueToName[r]
	if !ok {
		return fmt.Sprintf("TxStatus(%d)", r)
	}
	return s
}

// Validate verifies that value is predefined for TxStatus.
func (r TxStatus) Validate() error {
	_, ok := defTxStatusValueToName[r]
	if !ok {
		return ErrTxStatusInvalid
	}
	return nil
}

// MarshalJSON is generated so TxStatus satisfies json.Marshaler.
func (r TxStatus) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(r).(fmt.Stringer); ok {
		return json.Marshal(s.String())
	}
	s, ok := defTxStatusValueToName[r]
	if !ok {
		return nil, fmt.Errorf("TxStatus(%d) is invalid value", r)
	}
	return json.Marshal(s)
}

// UnmarshalJSON is generated so TxStatus satisfies json.Unmarshaler.
func (r *TxStatus) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("TxStatus: should be a string, got %s", string(data))
	}
	v, ok := defTxStatusNameToValue[s]
	if !ok {
		return fmt.Errorf("TxStatus(%q) is invalid value", s)
	}
	*r = v
	return nil
}

// Value is generated so TxStatus satisfies db row driver.Valuer.
func (r TxStatus) Value() (driver.Value, error) {
	j, err := json.Marshal(r)
	return j, err
}

// Value is generated so TxStatus satisfies db row driver.Scanner.
func (r *TxStatus) Scan(src interface{}) error {
	switch src.(type) {
	case string:
		val, ok := defTxStatusNameToValue[src.(string)]
		if !ok {
			return errors.New("TxStatus: can't unmarshal column data")
		}
		*r = val
		return nil
	case []byte:
		source := src.([]byte)
		var i TxStatus
		err := json.Unmarshal(source, &i)
		if err != nil {
			return errors.New("TxStatus: can't unmarshal column data")
		}

		*r = i
		return nil
	case int, int8, int32, int64, uint, uint8, uint32, uint64:
		ni := sql.NullInt64{}
		err := ni.Scan(src)
		if err != nil {
			return errors.New("TxStatus: can't scan column data into int64")
		}

		*r = TxStatus(ni.Int64)
		return nil
	}
	return errors.New("TxStatus: invalid type")
}