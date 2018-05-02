package sharet

type RepaymentWallet struct {
	WalletId    string `db:"wallet_id" json:"walletId"`
	UserId      int64  `db:"user_id" json:"userId"`
	SharesCount int64  `db:"count" json:"sharesCount"`
}
