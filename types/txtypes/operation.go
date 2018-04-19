package txtypes

import (
	"github.com/pkg/errors"
	"gitlab.inn4science.com/vcg/go-common/crypto"
	"gitlab.inn4science.com/vcg/go-common/types/currency"
)

// Operation is a representation of the `operations` table.
type Operation struct {
	BaseRow

	OperationID  string        `db:"operation_id" json:"operationId"`
	Type         OperationType `db:"type" json:"type"`
	TxType       TxType        `db:"tx_type" json:"txType"`
	Counterparty string        `db:"counterparty" json:"counterparty"`
	Amount       currency.Coin `db:"amount" json:"amount"`
	Reference    string        `db:"reference" json:"reference"`
	TxID         *string       `db:"tx_id" json:"txId,omitempty"`
	Hash         string        `db:"hash" json:"hash"`
	Signature    string        `db:"signature" json:"signature"`
	CreatedAt    int64         `db:"created_at" json:"createdAt"`
	UpdatedAt    int64         `db:"updated_at" json:"-"`
}

// HashableOperation is a representation of the `operations` for hashing.
type HashableOperation struct {
	ID           string        `json:"id"`
	Hash         string        `json:"hash,omitempty"`
	Signature    string        `json:"signature,omitempty"`
	Counterparty string        `json:"counterparty"`
	Amount       currency.Coin `json:"amount"`
	Reference    string        `json:"reference"`
	Type         OperationType `json:"type"`
	CreatedAt    int64         `json:"createdAt"`
}

// ToHashable returns representation of the `operations` for hashing.
func (op *Operation) ToHashable() HashableOperation {
	return HashableOperation{
		ID:           op.OperationID,
		Type:         op.Type,
		Counterparty: op.Counterparty,
		Amount:       op.Amount,
		Reference:    op.Reference,
		CreatedAt:    op.CreatedAt,
	}
}

// GetHash create operation hash.
func (op *Operation) GetHash() (string, error) {
	if op.Hash != "" {
		return op.Hash, nil
	}

	var err error
	op.Hash, err = crypto.HashData(op.ToHashable())
	return op.Hash, err
}

// Sign signs the transaction hash using the private key.
func (op *Operation) Sign(privateKey string) (string, error) {
	_, err := op.GetHash()
	if err != nil {
		return "", errors.Wrap(err, "unable to get hash for op: "+op.OperationID)
	}

	op.Signature, err = crypto.SignMessage(privateKey, op.Hash)
	return op.Signature, err
}
