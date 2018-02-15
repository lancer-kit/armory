package currency

import (
	"encoding/json"
	"fmt"
	"math/big"
)

const (
	One      = 100000000
	CoinPrec = 8
	FiatPrec = 2
)

// Amount is a type for coins or fiats values.
type Amount int64

// Parse returns `Amount` parsed from string or parse error.
func Parse(str string) (Amount, error) {
	var f, one, result big.Float

	_, ok := f.SetString(str)
	if !ok {
		return Amount(0), fmt.Errorf("cannot parse amount: %s", str)
	}

	one.SetInt64(One)
	result.Mul(&f, &one)

	i, _ := result.Int64()
	return Amount(i), nil
}

// FromFloat converts float64 value to the `Amount`.
func FromFloat(val float64) Amount {
	var one, mult, bigVal big.Rat
	one.SetInt64(One)
	bigVal.SetFloat64(val)
	mult.Mul(&bigVal, &one)

	res, _ := mult.Float64()
	return Amount(int64(res))
}

// StringFromInt64 returns an "amount string" from the provided raw int64 value `v`.
func StringFromInt64(v int64, prec int) string {
	var f, one, r big.Rat
	f.SetInt64(v)
	one.SetInt64(One)
	r.Quo(&f, &one)
	return r.FloatString(prec)
}

// String returns an "amount string" with coin prec.
func (a Amount) String() string {
	return StringFromInt64(int64(a), CoinPrec)
}

// CurrencyString returns an "amount string" with currency prec.
func (a Amount) CurrencyString() string {
	return StringFromInt64(int64(a), FiatPrec)
}

// UnmarshalJSON implementation of the `json.Unmarshaller` interface.
func (a *Amount) UnmarshalJSON(data []byte) error {
	var fVal float64
	err := json.Unmarshal(data, &fVal)
	if err == nil {
		*a = FromFloat(fVal)
		return nil
	}
	var sVal string
	err = json.Unmarshal(data, &sVal)
	if err == nil {
		*a, err = Parse(sVal)
		return err
	}

	return fmt.Errorf("invalid amount value: %s", string(data))
}

// MarshalJSON implementation of the `json.Marshaller` interface.
func (a Amount) MarshalJSON() ([]byte, error) {
	str := a.String()
	return []byte(str), nil
}

// GetPercent calculates the percentage value from the sum and rounds it up.
func (a Amount) GetPercent(percent int64) Amount {
	var amountRat, percentRat, base, result big.Float
	amountRat.SetInt64(int64(a))
	percentRat.SetInt64(percent)
	base.SetInt64(100)

	result.Quo(&amountRat, &base)
	result.Mul(&result, &percentRat)
	res, acc := result.Int64()
	if acc == big.Below {
		res += 1
	}
	return Amount(res)
}
