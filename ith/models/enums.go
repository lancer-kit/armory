package models

import "errors"

//todo: maybe redesign as iota enum
type OrderStatus string

const (
	OrderDraft                 OrderStatus = "DR"
	OrderDeleted               OrderStatus = "DE"
	OrderSent                  OrderStatus = "SE"
	OrderCancelled             OrderStatus = "CA"
	OrderPaid                  OrderStatus = "PA"
	OrderRefunded              OrderStatus = "RF"
	OrderPartiallyRefunded     OrderStatus = "RP"
	OrderMarkedPaid            OrderStatus = "MP"
	OrderMarkedRefunded        OrderStatus = "MR"
	OrderChargebacked          OrderStatus = "CF"
	OrderPartiallyChargebacked OrderStatus = "CP"
	OrderExpired               OrderStatus = "EX"
)

var (
	ErrOrderStatusInvalid = errors.New("invalid orders status")

	validOrderStatuses = map[OrderStatus]struct{}{
		OrderDraft:                 {},
		OrderDeleted:               {},
		OrderSent:                  {},
		OrderCancelled:             {},
		OrderPaid:                  {},
		OrderRefunded:              {},
		OrderPartiallyRefunded:     {},
		OrderMarkedPaid:            {},
		OrderMarkedRefunded:        {},
		OrderChargebacked:          {},
		OrderPartiallyChargebacked: {},
		OrderExpired:               {},
	}
)

func (os OrderStatus) Validate() error {
	if _, ok := validOrderStatuses[os]; !ok {
		return ErrOrderStatusInvalid
	}
	return nil
}
