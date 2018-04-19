package tx

type TxSource struct {
	ID        int64           `db:"id" json:"id"`
	TxID      string          `db:"tx_id" json:"txId"`
	Reference string          `db:"reference" json:"reference"`
	Type      TxType          `db:"type" json:"type"`
	Source    OperationSource `db:"data" json:"source"`
	State     TxState         `db:"state" json:"state"`
	CreatedAt int64           `db:"created_at" json:"createdAt"`
	UpdatedAt int64           `db:"updated_at" json:"updatedAt"`
}
