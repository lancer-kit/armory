package currency

import "math"

type ConversionResult struct {
	Coins Coin  `json:"coins"`
	Fiat  Fiat  `json:"fiat"`
	Price Price `json:"price"`
}

type Convertible interface {
	Convert(Price) ConversionResult
}

func BankRound(val float64, prec int) float64 {
	var round float64
	pow := math.Pow10(prec)
	digit := pow * val
	_, div := math.Modf(digit)

	rem := math.Mod(float64(int64(digit)), 2)
	div = math.Abs(div)

	if div == 0.5 && int64(rem) != 0 {
		round = math.Ceil(digit)
	} else if div > 0.5 {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}

	newVal := round / pow

	return newVal
}
