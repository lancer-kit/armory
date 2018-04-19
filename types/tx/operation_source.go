package tx

import (
	"database/sql/driver"
	"encoding/json"

	"github.com/pkg/errors"
)

type OperationSource struct {
	Deposit            *Deposit            `json:"deposit,omitempty"`
	Payment            *Payment            `json:"payment,omitempty"`
	CommissionClearing *CommissionClearing `json:"commissionClearing,,omitempty"`
}

func (p OperationSource) Value() (driver.Value, error) {
	j, err := json.Marshal(p)
	return j, err
}

func (p *OperationSource) Scan(src interface{}) error {
	source, ok := src.([]byte)
	if !ok {
		return errors.New("Type assertion .([]byte) failed.")
	}

	var i OperationSource
	err := json.Unmarshal(source, &i)
	if err != nil {
		return errors.Wrap(err, "OperationSource: can't unmarshal column data")
	}

	*p = i
	return nil
}
