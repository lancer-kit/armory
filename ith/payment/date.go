package payment

import (
	"fmt"
	"time"
)

const TimeLayout = "20060102150405"

type IthTime string

func (t IthTime) FromTimestamp(tm int64) IthTime {
	return IthTime(time.Unix(tm, 0).Format(TimeLayout))
}

func (t IthTime) FromTime(tm *time.Time) IthTime {
	return IthTime(tm.Format(TimeLayout))
}

func (t IthTime) ToTime() *time.Time {
	if !t.Valid() {
		return new(time.Time)
	}
	tm, _ := time.Parse(TimeLayout, string(t))
	return &tm
}

func (t IthTime) Validate() error {
	_, err := time.Parse(TimeLayout, string(t))
	if err != nil {
		return fmt.Errorf("IthTime(%s): invalid value", t)
	}
	return nil
}

func (t IthTime) Valid() bool {
	return t.Validate() == nil
}
