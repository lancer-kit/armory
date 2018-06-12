package ams

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type RequestStatus int

const (
	RequestStatusNone RequestStatus = iota
	RequestStatusValidationError
	RequestStatusDbError
	RequestStatusNetworkError
	RequestStatusPartnerError
	RequestStatusOk
)

var ErrRequestStatusInvalid = errors.New("RequestStatus is invalid")

var defRequestStatusNameToValue = map[string]RequestStatus{
	"":    RequestStatusNone,
	"CVE": RequestStatusValidationError,
	"CDE": RequestStatusDbError,
	"CNE": RequestStatusNetworkError,
	"CPE": RequestStatusPartnerError,
	"COK": RequestStatusOk,
}

var defRequestStatusValueToName = map[RequestStatus]string{
	RequestStatusNone:            "",
	RequestStatusValidationError: "CVE",
	RequestStatusDbError:         "CDE",
	RequestStatusNetworkError:    "CNE",
	RequestStatusPartnerError:    "CPE",
	RequestStatusOk:              "COK",
}

// String is generated so RequestStatus satisfies fmt.Stringer.
func (r RequestStatus) String() string {
	s, ok := defRequestStatusValueToName[r]
	if !ok {
		return fmt.Sprintf("RequestStatus(%d)", r)
	}
	return s
}

// Validate verifies that value is predefined for RequestStatus.
func (r RequestStatus) Validate() error {
	_, ok := defRequestStatusValueToName[r]
	if !ok {
		return ErrRequestStatusInvalid
	}
	return nil
}

// MarshalJSON is generated so RequestStatus satisfies json.Marshaler.
func (r RequestStatus) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(r).(fmt.Stringer); ok {
		return json.Marshal(s.String())
	}
	s, ok := defRequestStatusValueToName[r]
	if !ok {
		return nil, fmt.Errorf("RequestStatus(%d) is invalid value", r)
	}
	return json.Marshal(s)
}

// UnmarshalJSON is generated so RequestStatus satisfies json.Unmarshaler.
func (r *RequestStatus) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("RequestStatus: should be a string, got %s", string(data))
	}
	v, ok := defRequestStatusNameToValue[s]
	if !ok {
		return fmt.Errorf("RequestStatus(%q) is invalid value", s)
	}
	*r = v
	return nil
}

// Value is generated so RequestStatus satisfies db row driver.Valuer.
func (r RequestStatus) Value() (driver.Value, error) {
	j, err := json.Marshal(r)
	return j, err
}

// Value is generated so RequestStatus satisfies db row driver.Scanner.
func (r *RequestStatus) Scan(src interface{}) error {
	switch src.(type) {
	case string:
		val, ok := defRequestStatusNameToValue[src.(string)]
		if !ok {
			return errors.New("RequestStatus: can't unmarshal column data")
		}
		*r = val
		return nil
	case []byte:
		source := src.([]byte)
		var i RequestStatus
		err := json.Unmarshal(source, &i)
		if err != nil {
			return errors.New("RequestStatus: can't unmarshal column data")
		}

		*r = i
		return nil
	case int, int8, int32, int64, uint, uint8, uint32, uint64:
		ni := sql.NullInt64{}
		err := ni.Scan(src)
		if err != nil {
			return errors.New("RequestStatus: can't scan column data into int64")
		}

		*r = RequestStatus(ni.Int64)
		return nil
	}
	return errors.New("RequestStatus: invalid type")
}