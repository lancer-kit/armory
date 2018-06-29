package txst

import "gitlab.inn4science.com/vcg/go-common/api/render"

type Transactional interface {
	UID() string
	ToOpSource() OperationSource
	ToOperations() OperationSet
	TxType() TxType
}

type BaseRow struct {
	render.BaseRow
}
