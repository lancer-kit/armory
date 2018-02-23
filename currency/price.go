package currency

import (
	"math/big"
)

type Price float64

// String returns an "price string" with price precision.
func (price Price) String() string {
	var res big.Rat
	res.SetFloat64(float64(price))
	return res.FloatString(PricePrecision)
}

// MarshalJSON implementation of the `json.Marshaller` interface.
func (price Price) MarshalJSON() ([]byte, error) {
	str := price.String()
	return []byte(str), nil
}

func (price *Price) Round() Price {
	f := BankRound(float64(*price), PricePrecision)
	*price = Price(f)
	return *price
}

func (price *Price) Float() float64 {
	return float64(*price)
}
