package payment

import (
	"net/url"
	"testing"
	"time"

	"encoding/json"

	"github.com/stretchr/testify/assert"
	"gitlab.inn4science.com/vcg/go-common/ith/auth"
	"gitlab.inn4science.com/vcg/go-common/types/currency"
)

func TestAPI_CreateOrder(t *testing.T) {
	api := API{}
	api.Config.BaseURL, _ = url.Parse("http://demo-api.enauda.com/")
	result, err := api.GetAuthToken(&auth.Request{
		Username: "RdHMQMZCGZKRyXHf",
		Password: "UFByqq07dMxe7m0",
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, result)
	assert.Empty(t, result.ErrorData)
	assert.NotEmpty(t, result.AccessToken)
	assert.NotEmpty(t, result.RefreshToken)
	assert.NotEmpty(t, result.ExpiresIn)

	result, err = api.RefreshAuthToken(result.RefreshToken)
	assert.NoError(t, err)
	assert.NotEmpty(t, result)
	assert.Empty(t, result.ErrorData)
	assert.NotEmpty(t, result.AccessToken)
	assert.NotEmpty(t, result.RefreshToken)
	assert.NotEmpty(t, result.ExpiresIn)

	order := &Order{
		Seller: &auth.Account{
			Email: "vipcoin-merchant@coinfide.com",
		},
		Buyer: &auth.Account{
			Email: "vipcoin-standard@coinfide.com",
		},
		CurrencyCode:    "EUR",
		ExternalOrderID: "VCG-ORDER-wqefmolseemopptuyms",
		OrderItems: []OrderItem{
			{
				Type:        OrderITypeItem,
				Name:        "Item1",
				Description: "First item description",
				PriceUnit:   currency.Fiat(10 * currency.One),
				Quantity:    currency.Fiat(1 * currency.One),
			},
		},
	}
	_, err = json.Marshal(order)
	assert.NoError(t, err)

	list, err := api.GetOrderList(
		time.Now().Add(-1*30*24*time.Hour).Unix(),
		time.Now().Unix())
	assert.NoError(t, err)
	assert.NotEmpty(t, list)
	list, err = api.GetOrderDetails("GetOrderDetails", "8ed61188-a43d-44f2-85b5-563df4bf92b8")
	assert.NoError(t, err)
	assert.NotEmpty(t, list)
}
