package tx

import (
	"fmt"

	"gitlab.inn4science.com/vcg/go-common/crypto"
	"gitlab.inn4science.com/vcg/go-common/types/currency"
	"gitlab.inn4science.com/vcg/go-common/types/share"
)

// Repayment
type Repayment struct {
	ID             int64                   `json:"id"`
	RID            string                  `json:"uid"`
	ToWallets      []share.RepaymentWallet `json:"toWallets"`
	Amount         currency.Coin           `json:"amount"`
	CreatedAt      int64                   `json:"createdAt"`
	MerchantWallet string                  `json:"merchantWallet"`
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
	}

	for _, w := range repayment.ToWallets {
		result = append(result, &Operation{
			OperationID: crypto.HashStrings(
				repayment.RID,
				w.WalletId,
				w.Amount.String(),
				fmt.Sprintf("%d", repayment.CreatedAt)),
			Counterparty: w.WalletId,
			Amount:       w.Amount,
			Type:         OpTypeSystemTransfer,
			Reference:    repayment.RID,
			TxType:       TxTypeRepayment,
			CreatedAt:    repayment.CreatedAt,
		})
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
