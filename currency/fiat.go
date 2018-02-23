package currency

import (
	"math/big"
)

type Fiat Amount

// String returns an "amount string" with amount precision.
func (a Fiat) String() string {
	val := int64(a.Round())
	return StringFromInt64(val, FiatPrecision)
}

// MarshalJSON implementation of the `json.Marshaller` interface.
func (a Fiat) MarshalJSON() ([]byte, error) {
	str := a.String()
	return []byte(str), nil
}

// MarshalJSON implementation of the `json.Marshaller` interface.
func (a Fiat) Convert(price Price) ConversionResult {
	//	fiat / price = coins
	result, _ := new(big.Float).Quo(big.NewFloat(float64(a)), big.NewFloat(float64(price))).Int64()
	coins := Coin(result)

	return ConversionResult{
		Coins: coins,
		Fiat:  a,
		Price: price,
	}
}
func (a Fiat) Round() Fiat {
	f := Amount(a).Float64()
	am := FromFloat(BankRound(f, FiatPrecision))
	return Fiat(am)
}
