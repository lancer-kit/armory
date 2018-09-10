// generated by goplater enum --type WorkerState; DO NOT EDIT
package routines

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

func init() {
	// stub usage of json for situation when
	// (Un)MarshalJSON methods will be omitted
	_ = json.Delim('s')

	// stub usage of sql/driver for situation when
	// Scan/Value methods will be omitted
	_ = driver.Bool
	_ = sql.LevelDefault
}

var ErrWorkerStateInvalid = errors.New("WorkerState is invalid")

var defWorkerStateNameToValue = map[string]WorkerState{
	"WorkerWrongStateChange": WorkerWrongStateChange,
	"WorkerNull":             WorkerNull,
	"WorkerDisabled":         WorkerDisabled,
	"WorkerPresent":          WorkerPresent,
	"WorkerEnabled":          WorkerEnabled,
	"WorkerInitialized":      WorkerInitialized,
	"WorkerRun":              WorkerRun,
	"WorkerStopped":          WorkerStopped,
	"WorkerFailed":           WorkerFailed,
}

var defWorkerStateValueToName = map[WorkerState]string{
	WorkerWrongStateChange: "WorkerWrongStateChange",
	WorkerNull:             "WorkerNull",
	WorkerDisabled:         "WorkerDisabled",
	WorkerPresent:          "WorkerPresent",
	WorkerEnabled:          "WorkerEnabled",
	WorkerInitialized:      "WorkerInitialized",
	WorkerRun:              "WorkerRun",
	WorkerStopped:          "WorkerStopped",
	WorkerFailed:           "WorkerFailed",
}

// String is generated so WorkerState satisfies fmt.Stringer.
func (r WorkerState) String() string {
	s, ok := defWorkerStateValueToName[r]
	if !ok {
		return fmt.Sprintf("WorkerState(%d)", r)
	}
	return s
}

// Validate verifies that value is predefined for WorkerState.
func (r WorkerState) Validate() error {
	_, ok := defWorkerStateValueToName[r]
	if !ok {
		return ErrWorkerStateInvalid
	}
	return nil
}

// MarshalJSON is generated so WorkerState satisfies json.Marshaler.
func (r WorkerState) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(r).(fmt.Stringer); ok {
		return json.Marshal(s.String())
	}
	s, ok := defWorkerStateValueToName[r]
	if !ok {
		return nil, fmt.Errorf("WorkerState(%d) is invalid value", r)
	}
	return json.Marshal(s)
}

// UnmarshalJSON is generated so WorkerState satisfies json.Unmarshaler.
func (r *WorkerState) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("WorkerState: should be a string, got %s", string(data))
	}
	v, ok := defWorkerStateNameToValue[s]
	if !ok {
		return fmt.Errorf("WorkerState(%q) is invalid value", s)
	}
	*r = v
	return nil
}

// Value is generated so WorkerState satisfies db row driver.Valuer.
func (r WorkerState) Value() (driver.Value, error) {
	s, ok := defWorkerStateValueToName[r]
	if !ok {
		return nil, nil
	}
	return s, nil
}

// Value is generated so WorkerState satisfies db row driver.Scanner.
func (r *WorkerState) Scan(src interface{}) error {
	switch v := src.(type) {
	case string:
		val, _ := defWorkerStateNameToValue[v]
		*r = val
		return nil
	case []byte:
		var i WorkerState
		err := json.Unmarshal(v, &i)
		if err != nil {
			return errors.New("WorkerState: can't unmarshal column data")
		}

		*r = i
		return nil
	case int, int8, int32, int64, uint, uint8, uint32, uint64:
		ni := sql.NullInt64{}
		err := ni.Scan(v)
		if err != nil {
			return errors.New("WorkerState: can't scan column data into int64")
		}

		*r = WorkerState(ni.Int64)
		return nil
	}
	return errors.New("WorkerState: invalid type")
}
