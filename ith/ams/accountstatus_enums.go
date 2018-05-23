package ams

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

// AccountStatus:
//
//  * SA – Standard: Automatically Registered
//  * SR – Standard: Registration Requested
//  * SC – Standard: Registration Confirmed
//  * SF – Standard: Customer Identified
//  * SB – Standard: Blocked
//  * SD – Standard: Closed
//  * BR – Business: Registration Requested
//  * BC – Business: Registration Confirmed (Read Only)
//  * BF – Business: Registration Finished (Agreement Signed)
//  * BM – Business: Requires Moderation
//  * BS – Business: Suspended (Blocked)
//  * BD – Business: Closed
//  * MR – Merchant: Registration Requested
//  * MC – Merchant: Registration Confirmed (Read Only)
//  * MF – Merchant: Registration Finished (Agreement Signed)
//  * MM – Merchant: Requires Moderation
//  * MS – Merchant: Suspended (Blocked)
//  * MD – Merchant: Closed
type AccountStatus int

const (
	StStandardAutomaticallyRegistered AccountStatus = iota + 1 //SA – Standard: Automatically Registered
	StStandardRegistrationRequested                            //SR – Standard: Registration Requested
	StStandardRegistrationConfirmed                            //SC – Standard: Registration Confirmed
	StStandardCustomerIdentified                               //SF – Standard: Customer Identified
	StStandardBlocked                                          //SB – Standard: Blocked
	StStandardClosed                                           //SD – Standard: Closed
	StBusinessRegistrationRequested                            //BR – Business: Registration Requested
	StBusinessRegistrationConfirmed                            //BC – Business: Registration Confirmed (Read Only)
	StBusinessRegistrationFinished                             //BF – Business: Registration Finished (Agreement Signed)
	StBusinessRequiresModeration                               //BM – Business: Requires Moderation
	StBusinessSuspended                                        //BS – Business: Suspended (Blocked)
	StBusinessClosed                                           //BD – Business: Closed
	StMerchantRegistrationRequested                            //MR – Merchant: Registration Requested
	StMerchantRegistrationConfirmed                            //MC – Merchant: Registration Confirmed (Read Only)
	StMerchantRegistrationFinished                             //MF – Merchant: Registration Finished (Agreement Signed)
	StMerchantRequiresModeration                               //MM – Merchant: Requires Moderation
	StMerchantSuspended                                        //MS – Merchant: Suspended (Blocked)
	StMerchantClosed                                           //MD – Merchant: Closed
)

var ErrAccountStatusInvalid = errors.New("AccountStatus is invalid")

var defAccountStatusNameToValue = map[string]AccountStatus{
	"SA": StStandardAutomaticallyRegistered,
	"SR": StStandardRegistrationRequested,
	"SC": StStandardRegistrationConfirmed,
	"SF": StStandardCustomerIdentified,
	"SB": StStandardBlocked,
	"SD": StStandardClosed,

	"BR": StBusinessRegistrationRequested,
	"BC": StBusinessRegistrationConfirmed,
	"BF": StBusinessRegistrationFinished,
	"BM": StBusinessRequiresModeration,
	"BS": StBusinessSuspended,
	"BD": StBusinessClosed,

	"MR": StMerchantRegistrationRequested,
	"MC": StMerchantRegistrationConfirmed,
	"MF": StMerchantRegistrationFinished,
	"MM": StMerchantRequiresModeration,
	"MS": StMerchantSuspended,
	"MD": StMerchantClosed,
}

var defAccountStatusValueToName = map[AccountStatus]string{
	StStandardAutomaticallyRegistered: "SA",
	StStandardRegistrationRequested:   "SR",
	StStandardRegistrationConfirmed:   "SC",
	StStandardCustomerIdentified:      "SF",
	StStandardBlocked:                 "SB",
	StStandardClosed:                  "SD",

	StBusinessRegistrationRequested: "BR",
	StBusinessRegistrationConfirmed: "BC",
	StBusinessRegistrationFinished:  "BF",
	StBusinessRequiresModeration:    "BM",
	StBusinessSuspended:             "BS",
	StBusinessClosed:                "BD",

	StMerchantRegistrationRequested: "MR",
	StMerchantRegistrationConfirmed: "MC",
	StMerchantRegistrationFinished:  "MF",
	StMerchantRequiresModeration:    "MM",
	StMerchantSuspended:             "MS",
	StMerchantClosed:                "MD",
}

// String is generated so AccountStatus satisfies fmt.Stringer.
func (r AccountStatus) String() string {
	s, ok := defAccountStatusValueToName[r]
	if !ok {
		return fmt.Sprintf("AccountStatus(%d)", r)
	}
	return s
}

// Validate verifies that value is predefined for AccountStatus.
func (r AccountStatus) Validate() error {
	_, ok := defAccountStatusValueToName[r]
	if !ok {
		return ErrAccountStatusInvalid
	}
	return nil
}

// MarshalJSON is generated so AccountStatus satisfies json.Marshaler.
func (r AccountStatus) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(r).(fmt.Stringer); ok {
		return json.Marshal(s.String())
	}
	s, ok := defAccountStatusValueToName[r]
	if !ok {
		return nil, fmt.Errorf("AccountStatus(%d) is invalid value", r)
	}
	return json.Marshal(s)
}

// UnmarshalJSON is generated so AccountStatus satisfies json.Unmarshaler.
func (r *AccountStatus) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("AccountStatus: should be a string, got %s", string(data))
	}
	v, ok := defAccountStatusNameToValue[s]
	if !ok {
		return fmt.Errorf("AccountStatus(%q) is invalid value", s)
	}
	*r = v
	return nil
}

// Value is generated so AccountStatus satisfies db row driver.Valuer.
func (r AccountStatus) Value() (driver.Value, error) {
	j, err := json.Marshal(r)
	return j, err
}

// Value is generated so AccountStatus satisfies db row driver.Scanner.
func (r *AccountStatus) Scan(src interface{}) error {
	switch src.(type) {
	case string:
		val, ok := defAccountStatusNameToValue[src.(string)]
		if !ok {
			return errors.New("AccountStatus: can't unmarshal column data")
		}
		*r = val
		return nil
	case []byte:
		source := src.([]byte)
		var i AccountStatus
		err := json.Unmarshal(source, &i)
		if err != nil {
			return errors.New("AccountStatus: can't unmarshal column data")
		}

		*r = i
		return nil
	case int, int8, int32, int64, uint, uint8, uint32, uint64:
		ni := sql.NullInt64{}
		err := ni.Scan(src)
		if err != nil {
			return errors.New("AccountStatus: can't scan column data into int64")
		}

		*r = AccountStatus(ni.Int64)
		return nil
	}
	return errors.New("AccountStatus: invalid type")
}
