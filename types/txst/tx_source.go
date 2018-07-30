package txst

type TxSource struct {
	BaseRow

	TxID      string          `db:"tx_id" json:"txId"`
	Reference string          `db:"reference" json:"reference"`
	Type      TxType          `db:"type" json:"type"`
	Source    OperationSource `db:"data" json:"source"`
	State     TxState         `db:"state" json:"state"`
	CreatedAt int64           `db:"created_at" json:"createdAt"`
	UpdatedAt int64           `db:"updated_at" json:"updatedAt"`
}
