package tx

import (
	"fmt"
	"strings"
)

type OperationType int

const (
	OpTypeSystemTransfer OperationType = iota
	OpTypeIncreaseBalance
	OpTypeDecreaseBalance
	OpTypeEmitCoins
	OpTypeNullification
)

var operationTypeStrings = map[OperationType]string{
	OpTypeSystemTransfer:  "system_transfer",
	OpTypeIncreaseBalance: "increase_balance",
	OpTypeDecreaseBalance: "decrease_balance",
	OpTypeEmitCoins:       "emit_coins",
	OpTypeNullification:   "nullification",
}

var operationTypeValues = map[string]OperationType{
	"system_transfer":  OpTypeSystemTransfer,
	"increase_balance": OpTypeIncreaseBalance,
	"decrease_balance": OpTypeDecreaseBalance,
	"emit_coins":       OpTypeEmitCoins,
	"nullification":    OpTypeNullification,
}

// String implementation of the `fmt.Stringer` interface.
func (ot OperationType) String() string {
	return operationTypeStrings[ot]
}

// MarshalJSON implementation of the `json.Marshaller` interface.
func (ot OperationType) MarshalJSON() ([]byte, error) {
	opStr := operationTypeStrings[ot]
	return []byte(fmt.Sprintf(`"%s"`, opStr)), nil
}

// MarshalJSON implementation of the `json.Marshaller` interface.
func (ot *OperationType) UnmarshalJSON(data []byte) error {
	str := strings.Trim(string(data), "\"")
	opType, ok := operationTypeValues[str]
	if !ok {
		return fmt.Errorf("invalid OperationType: %v", string(data))
	}
	*ot = opType
	return nil
}
