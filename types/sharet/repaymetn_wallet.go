package sharet

import "gitlab.inn4science.com/vcg/go-common/types/currency"

type RepaymentWallet struct {
	WalletId    string `db:"wallet_id"`
	SharesCount int64  `db:"count"`
	Amount      currency.Coin
}
