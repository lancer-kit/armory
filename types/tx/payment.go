package tx

import (
	"fmt"

	"gitlab.inn4science.com/vcg/go-common/crypto"
	"gitlab.inn4science.com/vcg/go-common/types/currency"
)

// PaymentDefaultPercent is the default interest rate
// that will be charged to the recipient.
const PaymentDefaultPercent int64 = 1

// Payment is a representation of the `payments` table.
type Payment struct {
	ID           int64         `db:"id" json:"id"`
	PaymentID    string        `db:"payment_id" json:"paymentId"`
	FromWalletID string        `db:"from_wallet" json:"fromWallet"`
	ToWalletID   string        `db:"to_wallet" json:"toWallet"`
	Amount       currency.Coin `db:"amount" json:"amount"`
	Fee          currency.Coin `db:"fee" json:"fee"`
	State        TxState       `db:"state" json:"state"`
	Description  string        `db:"description" json:"description"`
	Referrer     string        `db:"referrer" json:"referrer"`
	CreatedAt    int64         `db:"created_at" json:"createdAt"`
	UpdatedAt    int64         `db:"updated_at" json:"updatedAt"`

	CommissionWallet string `json:"-"`
}

func (payment *Payment) ToOperations() OperationSet {
	referrerPart := payment.Fee.GetPercent(FeeReferrerPart)

	result := OperationSet{
		&Operation{
			OperationID: crypto.HashStrings(
				payment.PaymentID,
				payment.ToWalletID,
				payment.Amount.String(),
				fmt.Sprintf("%d", payment.CreatedAt)),
			Counterparty: payment.ToWalletID,
			Amount:       payment.Amount - payment.Fee,
			Reference:    payment.PaymentID,
			Type:         OpTypeIncreaseBalance,
			TxType:       TxTypePayment,
			CreatedAt:    payment.CreatedAt,
		},

		&Operation{
			OperationID: crypto.HashStrings(
				payment.PaymentID,
				payment.FromWalletID,
				payment.Amount.String(),
				fmt.Sprintf("%d", payment.CreatedAt)),
			Counterparty: payment.FromWalletID,
			Amount:       payment.Amount,
			Reference:    payment.PaymentID,
			Type:         OpTypeDecreaseBalance,
			TxType:       TxTypePayment,
			CreatedAt:    payment.CreatedAt,
		},

		&Operation{
			OperationID: crypto.HashStrings(
				payment.PaymentID,
				payment.CommissionWallet,
				payment.Amount.String(),
				fmt.Sprintf("%d", payment.CreatedAt)),
			Counterparty: payment.CommissionWallet,
			Amount:       payment.Fee - referrerPart,
			Reference:    payment.PaymentID,
			Type:         OpTypeSystemTransfer,
			TxType:       TxTypePayment,
			CreatedAt:    payment.CreatedAt,
		},

		&Operation{
			OperationID: crypto.HashStrings(
				payment.PaymentID,
				payment.Referrer,
				payment.Amount.String(),
				fmt.Sprintf("%d", payment.CreatedAt)),
			Counterparty: payment.Referrer,
			Amount:       referrerPart,
			Reference:    payment.PaymentID,
			Type:         OpTypeSystemTransfer,
			TxType:       TxTypePayment,
			CreatedAt:    payment.CreatedAt,
		},
	}

	return result
}

func (payment *Payment) ToOpSource() OperationSource {
	return OperationSource{Payment: payment}
}

func (payment *Payment) TxType() TxType {
	return TxTypePayment
}

func (payment *Payment) UID() string {
	return payment.PaymentID
}
