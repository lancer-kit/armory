package sharet

import "gitlab.inn4science.com/vcg/go-common/types/currency"

type RepaymentWallet struct {
	WalletId    string        `db:"walletId" json:"walletId"`
	SharesCount int64         `db:"count" json:"sharesCount"`
	Amount      currency.Coin `db:"amount" json:"amount"`
}
