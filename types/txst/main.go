package txst

type Transactional interface {
	UID() string
	ToOpSource() OperationSource
	ToOperations() OperationSet
	TxType() TxType
}

type BaseRow struct {
	ID       int64 `db:"id" json:"id"`
	RowCount int64 `db:"row_count" json:"-"`
}
