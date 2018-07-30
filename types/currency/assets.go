package currency

import "fmt"

// PriceList is [quote asset]price model.
type PriceList map[string]Price

type Asset struct {
	Code   string
	IsCoin bool
}

const (
	CodeVipCoin = "VC"
	CodeEur     = "EUR"
)

var (
	VCCoin = Asset{
		Code:   CodeVipCoin,
		IsCoin: true,
	}
	EUR = Asset{
		Code:   CodeEur,
		IsCoin: false,
	}

	Assets = map[string]Asset{
		CodeVipCoin: VCCoin,
		CodeEur:     EUR,
	}

	coinPrices = map[string]PriceList{
		CodeVipCoin: {
			CodeEur: 0.3464,
		},
	}
)

// GetPrice returns actual price for passed pair.
func GetPrice(coin, currency string) Price {
	prices := coinPrices[coin]
	return prices[currency]
}

// UpdateCoinPrices updates actual price for passed pair.
func UpdateCoinPrices(coin, currency string, price Price) {
	prices := coinPrices[coin]
	prices[currency] = price
	coinPrices[coin] = prices
}

// Convert converts some `amount` of asset to the `destination asset`.
func (asset *Asset) Convert(amount Amount, destAsset string) (ConversionResult, error) {
	if asset.IsCoin {
		prices, ok := coinPrices[asset.Code]
		if !ok {
			return ConversionResult{},
				fmt.Errorf("prices for coin %s not found", asset.Code)
		}
		price, ok := prices[destAsset]
		if !ok {
			return ConversionResult{},
				fmt.Errorf("coin prices for asset %s not found", destAsset)
		}
		am := Coin(amount)
		return am.Convert(price), nil
	}

	prices, ok := coinPrices[destAsset]
	if !ok {
		return ConversionResult{},
			fmt.Errorf("prices for coin %s not found", destAsset)
	}

	price, ok := prices[asset.Code]
	if !ok {
		return ConversionResult{},
			fmt.Errorf("coin prices for asset %s not found", asset.Code)
	}

	am := Fiat(amount)
	return am.Convert(price), nil
}
