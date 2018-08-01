package ams

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

type AmsDate time.Time

func (r *AmsDate) Empty() bool {
	if r == nil {
		return true
	}
	return r.Time().IsZero()
}

func DateFromInt(i int64) AmsDate {
	return AmsDate(time.Unix(i, 0))
}

// String is generated so AddressType satisfies fmt.Stringer.
func (r *AmsDate) String() string {
	if r.Empty() {
		return ""
	}
	y, m, d := time.Time(*r).Date()
	return fmt.Sprintf("%04d%02d%02d", y, m, d) + "000000"
}

func (r AmsDate) Time() time.Time {
	return time.Time(r)
}

// Validate verifies that value is predefined for AddressType.
func (r *AmsDate) Validate() error {
	if len(r.String()) != 14 {
		return fmt.Errorf("invalid AmsDate type: %s", r.String())
	}
	return nil
}

// MarshalJSON  AmsDate satisfies json.Marshaler.
func (r *AmsDate) MarshalJSON() ([]byte, error) {
	if r.Empty() {
		s := ""
		return json.Marshal(&s)
	}
	s := r.String()
	println("AMS Date:", s)
	a, e := json.Marshal(&s)
	println("AMS Date (marshalled:", string(a))
	return a, e
}

// UnmarshalJSON  AmsDate satisfies json.Unmarshaler.
func (r *AmsDate) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("AmsDate: should be a string, got %s", string(data))
	}
	if len(s) != 14 {
		return fmt.Errorf("AmsDate: format error")
	}
	y, err := strconv.Atoi(string(s[0:4]))
	if err != nil {
		return fmt.Errorf("AmsDate: year error, %s", err.Error())
	}
	m, err := strconv.Atoi(string(s[4:6]))
	if err != nil {
		return fmt.Errorf("AmsDate: month error, %s", err.Error())
	}
	d, err := strconv.Atoi(string(s[6:8]))
	if err != nil {
		return fmt.Errorf("AmsDate: date error, %s", err.Error())
	}
	t := time.Date(y, time.Month(m), d, 0, 0, 0, 0, time.Local)
	*r = AmsDate(t)
	return nil
}
