package sharet

type RepaymentWallet struct {
	WalletId    string `db:"walletId" json:"walletId"`
	SharesCount int64  `db:"count" json:"sharesCount"`
}
