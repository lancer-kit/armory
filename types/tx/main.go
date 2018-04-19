package tx

type Transactional interface {
	UID() string
	ToOpSource() OperationSource
	ToOperations() OperationSet
	TxType() TxType
}
