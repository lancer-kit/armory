package tx

import (
	"fmt"

	"gitlab.inn4science.com/vcg/go-common/crypto"
	"gitlab.inn4science.com/vcg/go-common/types/currency"
)

const (
	FeeMerchantPart             = 50
	FeeReferrerPart             = 25
	FeeNullablePart             = 25
	CommissionClearingThreshold = currency.Coin(0.1 * currency.One)
)

type CommissionClearing struct {
	CommissionWallet     string `json:"commissionWallet"`
	MerchantBucketWallet string `json:"merchantBucketWallet"`
	NullWallet           string `json:"nullWallet"`

	CommissionBalance currency.Coin `json:"commissionBalance"`
	LockSum           currency.Coin `json:"lockSum"`
	MerchantPart      currency.Coin `json:"merchantPart"`
	NullablePart      currency.Coin `json:"nullablePart"`

	Hash      string `json:"hash,omitempty"`
	CreatedAt int64  `json:"createdAt"`
}

func (op *CommissionClearing) SplitBalance() {
	totalPart := FeeReferrerPart + FeeMerchantPart + FeeNullablePart

	balanceWithRefererPart := int64(op.CommissionBalance) * int64(totalPart) / int64(FeeMerchantPart+FeeNullablePart)
	fullBalance := currency.Coin(balanceWithRefererPart)
	op.MerchantPart = fullBalance.GetPercent(FeeMerchantPart)
	op.NullablePart = fullBalance.GetPercent(FeeNullablePart)

	if op.MerchantPart+op.NullablePart > op.CommissionBalance {
		op.NullablePart = op.CommissionBalance - op.MerchantPart
	}

	op.LockSum = op.MerchantPart + op.NullablePart
}

func (op *CommissionClearing) ToOperations() OperationSet {
	balance := op.CommissionBalance

	now := op.CreatedAt
	reference := crypto.HashStrings(
		fmt.Sprintf("commission_wallet_clearing::balance:%s;parts:%s|%s;t:%d",
			balance.String(), op.MerchantPart, op.NullablePart, now),
	)

	result := OperationSet{
		&Operation{
			OperationID: crypto.HashStrings(
				reference,
				op.MerchantBucketWallet,
				op.MerchantPart.String(),
				fmt.Sprintf("%d", now),
			),
			Type:         OpTypeSystemTransfer,
			TxType:       TxTypeCommissionClearing,
			Counterparty: op.MerchantBucketWallet,
			Amount:       op.MerchantPart,
			Reference:    reference,
			CreatedAt:    now,
		},

		&Operation{
			OperationID: crypto.HashStrings(
				reference,
				op.NullWallet,
				op.NullablePart.String(),
				fmt.Sprintf("%d", now),
			),
			Type:         OpTypeNullification,
			TxType:       TxTypeCommissionClearing,
			Counterparty: op.NullWallet,
			Amount:       op.NullablePart,
			Reference:    reference,
			CreatedAt:    now,
		},

		&Operation{
			OperationID: crypto.HashStrings(
				reference,
				op.CommissionWallet,
				op.LockSum.String(),
				fmt.Sprintf("%d", now),
			),
			Type:         OpTypeDecreaseBalance,
			TxType:       TxTypeCommissionClearing,
			Counterparty: op.CommissionWallet,
			Amount:       op.LockSum,
			Reference:    reference,
			CreatedAt:    now,
		},
	}

	return result
}

func (op *CommissionClearing) ToOpSource() OperationSource {
	return OperationSource{CommissionClearing: op}
}
func (op *CommissionClearing) TxType() TxType {
	return TxTypeCommissionClearing
}

func (op *CommissionClearing) UID() string {
	if op.Hash == "" {
		op.Hash, _ = crypto.HashData(op)
	}

	return op.Hash
}
