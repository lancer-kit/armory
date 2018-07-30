package txst

import (
	"fmt"

	"gitlab.inn4science.com/vcg/go-common/crypto"
	"gitlab.inn4science.com/vcg/go-common/types/currency"
	"gitlab.inn4science.com/vcg/go-common/types/sharet"
)

// Repayment
type Repayment struct {
	ID             int64                  `json:"id"`
	RID            string                 `json:"uid"`
	Wallet         sharet.RepaymentWallet `json:"wallet"`
	Amount         currency.Coin          `json:"amount"`
	CreatedAt      int64                  `json:"createdAt"`
	MerchantWallet string                 `json:"merchantWallet"`
}

func (repayment *Repayment) ToOperations() OperationSet {
	result := OperationSet{
		&Operation{
			OperationID: crypto.HashStrings(
				repayment.RID,
				repayment.Amount.String(),
				fmt.Sprintf("%d", repayment.CreatedAt)),
			Counterparty: repayment.MerchantWallet,
			Amount:       repayment.Amount,
			Type:         OpTypeDecreaseBalance,
			Reference:    repayment.RID,
			TxType:       TxTypeRepayment,
			CreatedAt:    repayment.CreatedAt,
		},
		&Operation{
			OperationID: crypto.HashStrings(
				repayment.RID,
				repayment.Wallet.WalletId,
				repayment.Amount.String(),
				fmt.Sprintf("%d", repayment.CreatedAt)),
			Counterparty: repayment.Wallet.WalletId,
			Amount:       repayment.Amount,
			Type:         OpTypeSystemTransfer,
			Reference:    repayment.RID,
			TxType:       TxTypeRepayment,
			CreatedAt:    repayment.CreatedAt,
		},
	}

	return result
}

func (repayment *Repayment) ToOpSource() OperationSource {
	return OperationSource{Repayment: repayment}
}

func (repayment *Repayment) TxType() TxType {
	return TxTypeRepayment
}

func (repayment *Repayment) UID() string {
	return repayment.RID
}
