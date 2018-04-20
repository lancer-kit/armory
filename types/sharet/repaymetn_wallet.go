package sharet

import "gitlab.inn4science.com/vcg/go-common/types/currency"

type RepaymentWallet struct {
	WalletId    string `db:"walletId"`
	SharesCount int64  `db:"count"`
	Amount      currency.Coin
}
