package ith

import (
	"fmt"
	"time"
)

const TimeLayout = "20060102150405"

type Time string

func (t Time) FromTimestamp(tm int64) Time {
	return Time(time.Unix(tm, 0).Format(TimeLayout))
}

func (t Time) FromTime(tm *time.Time) Time {
	return Time(tm.Format(TimeLayout))
}

func (t Time) ToTime() *time.Time {
	if !t.Valid() {
		return new(time.Time)
	}
	tm, _ := time.Parse(TimeLayout, string(t))
	return &tm
}

func (t Time) Validate() error {
	_, err := time.Parse(TimeLayout, string(t))
	if err != nil {
		return fmt.Errorf("Time(%s): invalid value", t)
	}
	return nil
}

func (t Time) Valid() bool {
	return t.Validate() == nil
}
