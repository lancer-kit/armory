package txtypes

import (
	"fmt"

	"gitlab.inn4science.com/vcg/go-common/crypto"
	"gitlab.inn4science.com/vcg/go-common/types/currency"
)

// Deposit is a TODO: fill.
type Deposit struct {
	BaseRow

	DepositID      string         `json:"depositId"`
	ToWalletID     string         `json:"toWallet"`
	Amount         currency.Coin  `json:"amount"`
	Price          currency.Price `json:"price"`
	Currency       string         `json:"currency"`
	CreatedAt      int64          `json:"createdAt"`
	EmissionWallet string         `json:"emissionWallet"`
}

func (deposit *Deposit) ToOperations() OperationSet {
	result := OperationSet{
		&Operation{
			OperationID: crypto.HashStrings(
				deposit.DepositID,
				deposit.ToWalletID,
				deposit.Amount.String(),
				fmt.Sprintf("%d", deposit.CreatedAt)),
			Counterparty: deposit.ToWalletID,
			Amount:       deposit.Amount,
			Type:         OpTypeIncreaseBalance,
			Reference:    deposit.DepositID,
			TxType:       TxTypeDeposit,
			CreatedAt:    deposit.CreatedAt,
		},

		&Operation{
			OperationID: crypto.HashStrings(
				deposit.DepositID,
				deposit.EmissionWallet,
				deposit.Amount.String(),
				fmt.Sprintf("%d", deposit.CreatedAt)),
			Counterparty: deposit.EmissionWallet,
			Amount:       deposit.Amount,
			Type:         OpTypeEmitCoins,
			Reference:    deposit.DepositID,
			TxType:       TxTypeDeposit,
			CreatedAt:    deposit.CreatedAt,
		},
	}
	return result
}

func (deposit *Deposit) ToOpSource() OperationSource {
	return OperationSource{Deposit: deposit}
}

func (deposit *Deposit) TxType() TxType {
	return TxTypeDeposit
}

func (deposit *Deposit) UID() string {
	return deposit.DepositID
}
