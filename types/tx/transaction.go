package tx

import (
	"fmt"
	"time"

	"gitlab.inn4science.com/vcg/go-common/crypto"
	"github.com/pkg/errors"
)

type TxState string

const (
	TxStatePending   TxState = "pending"
	TxStateSubmitted TxState = "submitted"
	TxStateApplied   TxState = "applied"
	TxStateFailed    TxState = "failed"
)

// Transaction is a representation of the `transactions` table.
type Transaction struct {
	ID           int64           `db:"id" json:"id"`
	TxID         string          `db:"tx_id" json:"txId"`
	Type         TxType          `db:"type" json:"type"`
	State        TxState         `db:"state" json:"state"`
	OpCount      int             `db:"op_count" json:"opCount"`
	Operations   OperationSet      `db:"operations" json:"operations"`
	Source       OperationSource `db:"source" json:"source,omitempty"`
	Signature    string          `db:"signature" json:"signature"`
	Hash         string          `db:"hash" json:"hash"`
	PreviousHash string          `db:"prev_hash" json:"prevHash"`
	CreatedAt    int64           `db:"created_at" json:"createdAt"`
	UpdatedAt    int64           `db:"updated_at" json:"updatedAt"`
}

// HashableTransaction is a representation of the `transactions` for hashing.
type HashableTransaction struct {
	TxID         string              `json:"tx_id"`
	OpCount      int                 `json:"op_count"`
	Operations   []HashableOperation `json:"operations"`
	Type         TxType              `json:"type"`
	PreviousHash string              `json:"prev_hash"`
	CreatedAt    int64               `json:"created_at"`
}

// ToHashable returns representation of the `operations` for hashing.
func (tx *Transaction) ToHashable() (HashableTransaction, error) {
	data := HashableTransaction{
		TxID:         tx.TxID,
		Type:         tx.Type,
		OpCount:      tx.OpCount,
		PreviousHash: tx.PreviousHash,
		CreatedAt:    tx.CreatedAt,
	}

	data.Operations = make([]HashableOperation, len(tx.Operations))
	for i, op := range tx.Operations {
		if op.Signature == "" {
			return data, fmt.Errorf("operation must be signed, op_id: %s", op.OperationID)
		}

		data.Operations[i] = op.ToHashable()
	}

	return data, nil
}

// NewTransaction create new `transaction` with `txType` from `operations`.
func NewTransaction(txType TxType, operations OperationSet) (*Transaction, error) {
	txDetails, ok := TxOperationsDetails[txType]
	if !ok {
		return nil, fmt.Errorf("unknown tx type: %s", txType)
	}

	if txDetails.Fixed && len(operations) != txDetails.Count {
		return nil, fmt.Errorf(
			"number of ops(%d) does not match to expected(%d) for %s transaction",
			len(operations), txDetails.Count, txType)
	}

	permittedOps := txDetails.Types
	ops := make(OperationSet, 0, len(operations))

	for _, op := range operations {
		available, ok := permittedOps[op.Type]
		if !ok {
			return nil, fmt.Errorf("%s operations does not permitted for %s tx", op.Type, txType)
		}

		available--
		if available < 0 {
			return nil, fmt.Errorf("too many %s operations for %s tx", op.Type, txType)
		}

		ops = append(ops, op)
	}

	return &Transaction{
		TxID:       crypto.RandomString(32),
		Type:       txType,
		OpCount:    txDetails.Count,
		Operations: ops,
		CreatedAt:  time.Now().UTC().Unix(),
	}, nil
}

// GetHash creates transaction hash.
func (tx *Transaction) GetHash() (string, error) {
	if tx.Hash != "" {
		return tx.Hash, nil
	}

	if tx.OpCount != len(tx.Operations) {
		return "", fmt.Errorf("op_count field does not match the actual number of operations")
	}

	data, err := tx.ToHashable()
	if err != nil {
		return "", err
	}

	tx.Hash, err = crypto.HashData(data)
	return tx.Hash, err
}

// Sign signs the transaction hash using the private key.
func (tx *Transaction) Sign(privateKey string) (string, error) {
	_, err := tx.GetHash()
	if err != nil {
		return "", errors.Wrap(err, "unable to get hash for tx: "+tx.TxID)
	}
	tx.Signature, err = crypto.SignMessage(privateKey, tx.Hash)
	return tx.Signature, err
}
