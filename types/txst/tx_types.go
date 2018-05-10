package txst

import (
	"database/sql/driver"
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

type TxType int

const (
	TxTypeDeposit TxType = iota
	TxTypePayment
	TxTypeCommissionClearing
	TxTypeRepayment
)

var (
	txTypeStrings = map[TxType]string{
		TxTypeDeposit:            "deposit",
		TxTypePayment:            "payment",
		TxTypeCommissionClearing: "commission_clearing",
		TxTypeRepayment:          "repayment",
	}

	txTypeValueFromString = map[string]TxType{
		"deposit":             TxTypeDeposit,
		"payment":             TxTypePayment,
		"commission_clearing": TxTypeCommissionClearing,
		"repayment":           TxTypeRepayment,
	}
)

var TxOperationsDetails = map[TxType]struct {
	Fixed bool
	Count int
	Types map[OperationType]int
}{
	TxTypeDeposit: {
		Count: 2,
		Fixed: true,
		Types: map[OperationType]int{
			OpTypeIncreaseBalance: 1, OpTypeEmitCoins: 1,
		},
	},

	TxTypePayment: {
		Count: 4,
		Fixed: true,
		Types: map[OperationType]int{
			OpTypeIncreaseBalance: 1, OpTypeDecreaseBalance: 1, OpTypeSystemTransfer: 2,
		},
	},

	TxTypeCommissionClearing: {
		Count: 3,
		Fixed: true,
		Types: map[OperationType]int{
			OpTypeDecreaseBalance: 1, OpTypeSystemTransfer: 1, OpTypeNullification: 1,
		},
	},
	TxTypeRepayment: {
		Count: 2,
		Fixed: true,
		Types: map[OperationType]int{
			OpTypeDecreaseBalance: 1, OpTypeSystemTransfer: 1,
		},
	},
}

func (tt TxType) String() string {
	str, ok := txTypeStrings[tt]
	if !ok {
		return fmt.Sprintf("TxType(%d)", tt)
	}
	return str
}

func (tt *TxType) UnmarshalJSON(data []byte) error {
	str := strings.Trim(string(data), "\"")

	var ok bool
	*tt, ok = txTypeValueFromString[str]
	if !ok {
		return fmt.Errorf("unknown TxType: %s", str)
	}

	return nil
}

func (tt TxType) MarshalJSON() ([]byte, error) {
	str := tt.String()
	return []byte(fmt.Sprintf(`"%s"`, str)), nil
}

func (tt TxType) Value() (driver.Value, error) {
	return tt.String(), nil
}

func (tt *TxType) Scan(src interface{}) error {
	source, ok := src.(string)
	if !ok {
		return errors.New("Type assertion .(string) failed.")
	}
	source = strings.Trim(source, "\"")

	*tt, ok = txTypeValueFromString[source]
	if !ok {
		return fmt.Errorf("unknown TxType: %s", source)
	}

	return nil
}
