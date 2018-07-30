package ams

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

// WebResource:
//
//  * W – Web;
//  * F – Facebook;
//  * T – Twitter
type WebResource int

const (
	WebResourceWeb      WebResource = iota + 1 // W – Web;
	WebResourceFacebook                        //F – Facebook;
	WebResourceTwitter                         //T – Twitter
)

var ErrWebResourceInvalid = errors.New("WebResource is invalid")

var defWebResourceNameToValue = map[string]WebResource{
	"W": WebResourceWeb,
	"F": WebResourceFacebook,
	"T": WebResourceTwitter,
}

var defWebResourceValueToName = map[WebResource]string{
	WebResourceWeb:      "W",
	WebResourceFacebook: "F",
	WebResourceTwitter:  "T",
}

// String is generated so WebResource satisfies fmt.Stringer.
func (r WebResource) String() string {
	s, ok := defWebResourceValueToName[r]
	if !ok {
		return fmt.Sprintf("WebResource(%d)", r)
	}
	return s
}

// Validate verifies that value is predefined for WebResource.
func (r WebResource) Validate() error {
	_, ok := defWebResourceValueToName[r]
	if !ok {
		return ErrWebResourceInvalid
	}
	return nil
}

// MarshalJSON is generated so WebResource satisfies json.Marshaler.
func (r WebResource) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(r).(fmt.Stringer); ok {
		return json.Marshal(s.String())
	}
	s, ok := defWebResourceValueToName[r]
	if !ok {
		return nil, fmt.Errorf("WebResource(%d) is invalid value", r)
	}
	return json.Marshal(s)
}

// UnmarshalJSON is generated so WebResource satisfies json.Unmarshaler.
func (r *WebResource) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("WebResource: should be a string, got %s", string(data))
	}
	v, ok := defWebResourceNameToValue[s]
	if !ok {
		return fmt.Errorf("WebResource(%q) is invalid value", s)
	}
	*r = v
	return nil
}

// Value is generated so WebResource satisfies db row driver.Valuer.
func (r WebResource) Value() (driver.Value, error) {
	j, err := json.Marshal(r)
	return j, err
}

// Value is generated so WebResource satisfies db row driver.Scanner.
func (r *WebResource) Scan(src interface{}) error {
	switch src.(type) {
	case string:
		val, ok := defWebResourceNameToValue[src.(string)]
		if !ok {
			return errors.New("WebResource: can't unmarshal column data")
		}
		*r = val
		return nil
	case []byte:
		source := src.([]byte)
		var i WebResource
		err := json.Unmarshal(source, &i)
		if err != nil {
			return errors.New("WebResource: can't unmarshal column data")
		}

		*r = i
		return nil
	case int, int8, int32, int64, uint, uint8, uint32, uint64:
		ni := sql.NullInt64{}
		err := ni.Scan(src)
		if err != nil {
			return errors.New("WebResource: can't scan column data into int64")
		}

		*r = WebResource(ni.Int64)
		return nil
	}
	return errors.New("WebResource: invalid type")
}
