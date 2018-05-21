package payment

//go:generate goplater -type=OrderStatus -transform=snake -tprefix=false
type OrderStatus int

const (
	OrderDraft OrderStatus = iota
	OrderDeleted
	OrderSent
	OrderCancelled
	OrderPaid
	OrderRefunded
	OrderPartiallyRefunded
	OrderMarkedPaid
	OrderMarkedRefunded
	OrderChargebacked
	OrderPartiallyChargebacked
	OrderExpired
)

//go:generate goplater -type=TxAction -transform=none -tprefix=false
type TxAction int

const (
	TxActionPurchase TxAction = iota
	TxActionRefundFull
	TxActionRefundPartial
	TxActionChargebackFull
	TxActionChargebackPartial
	TxActionTransferFromRollingReserve
	TxActionTransferToRollingReserve
	TxActionCorrection
	TxActionFee
)

//go:generate goplater -type=TxType -transform=snake -tprefix=false
type TxType int

const (
	TxTypeIn TxType = iota
	TxTypeTransfer
	TxTypeOut
)

//go:generate goplater -type=TxStatus -transform=snake -tprefix=false
type TxStatus int

const (
	TxStatusSuccessful TxStatus = iota
	TxStatusPending
	TxStatusFailed
)

//go:generate goplater -type=OrderIType -transform=snake -tprefix=false
type OrderIType int

const (
	OrderITypeItem OrderIType = iota
	OrderITypeShipping
)
