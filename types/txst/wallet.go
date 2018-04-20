package txst

import (
	"fmt"

	"gitlab.inn4science.com/vcg/go-common/types/currency"
)

type WalletType int

const (
	WalletTypeNull WalletType = iota
	WalletTypeEmission
	WalletTypeCommission
	WalletTypeGeneral
	WalletTypeSystem
	WalletTypeMerchantBucket

	SystemWalletsUserID = 0
)

var WalletTypeStings = map[WalletType]string{
	WalletTypeNull:           "null",
	WalletTypeEmission:       "emission",
	WalletTypeCommission:     "commission",
	WalletTypeGeneral:        "general",
	WalletTypeSystem:         "system",
	WalletTypeMerchantBucket: "merchant",
}

func (t WalletType) String() string {
	str, ok := WalletTypeStings[t]
	if !ok {
		str = fmt.Sprintf("WalletType(%d)", t)
	}

	return str
}

// Wallet is a representation of the `wallets` table.
type Wallet struct {
	BaseRow

	WalletID  string        `db:"wallet_id" json:"walletId"`
	UserID    int64         `db:"user_id" json:"userId"`
	Balance   currency.Coin `db:"balance" json:"balance"`
	Locked    currency.Coin `db:"locked" json:"locked"`
	Debit     currency.Coin `db:"debit" json:"debit"`
	Credit    currency.Coin `db:"credit" json:"credit"`
	Type      WalletType    `db:"type" json:"type"`
	CreatedAt int64         `db:"created_at" json:"createdAt"`
	UpdatedAt int64         `db:"updated_at" json:"updatedAt"`
}
