package txst

import "gitlab.inn4science.com/vcg/go-common/api/render"

type Transactional interface {
	UID() string
	ToOpSource() OperationSource
	ToOperations() OperationSet
	TxType() TxType
}

// TODO: Update TX, check another and remove
type BaseRow struct {
	render.BaseRow
}
