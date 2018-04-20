package txst

import (
	"database/sql/driver"
	"encoding/json"

	"github.com/pkg/errors"
)

type OperationSet []*Operation

func (p OperationSet) Value() (driver.Value, error) {
	j, err := json.Marshal(p)
	return j, err
}

func (p *OperationSet) Scan(src interface{}) error {
	source, ok := src.([]byte)
	if !ok {
		return errors.New("Type assertion .([]byte) failed.")
	}

	i := make(OperationSet, 0)
	err := json.Unmarshal(source, &i)
	if err != nil {
		return errors.Wrap(err, "OperationSet: can't unmarshal column data")
	}

	*p = i
	return nil
}
