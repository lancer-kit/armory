package txst

import (
	"testing"

	"gitlab.inn4science.com/vcg/go-common/types/currency"
	"github.com/stretchr/testify/assert"
)

func TestCommissionClearing_SplitBalance(t *testing.T) {
	var testCases = []CommissionClearing{
		{
			CommissionBalance: currency.Coin(42),
			NullablePart:      currency.Coin(14),
			MerchantPart:      currency.Coin(28),
		},
		{
			CommissionBalance: currency.Coin(100),
			NullablePart:      currency.Coin(33),
			MerchantPart:      currency.Coin(67),
		},
	}

	for _, op := range testCases {
		testOp := CommissionClearing{
			CommissionBalance: op.CommissionBalance,
		}

		testOp.SplitBalance()
		assert.Equal(t, op.MerchantPart, testOp.MerchantPart)
		assert.Equal(t, op.NullablePart, testOp.NullablePart)
	}
}
