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
		ExternalOrderID: "VCG-ORDER-wqefmolseemopptuymo",
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

	extPurchaseOrder := &Order{
		Seller: &auth.Account{
			Email: "vipcoin-merchant@coinfide.com",
		},
		Buyer: &auth.Account{
			Email: "vipcoin-standard@coinfide.com",
		},
		CurrencyCode:    "EUR",
		ExternalOrderID: "VCG-ORDER-wqefmolseemopptuymo",
		OrderItems: []OrderItem{
			{
				Type:        OrderITypeItem,
				Name:        "Item1",
				Description: "First item description",
				PriceUnit:   currency.Fiat(10 * currency.One),
				Quantity:    currency.Fiat(1 * currency.One),
			},
		},
		ExternalPayout: &ExternalPayout{
			Method: 1,
		},
	}

	_, err = json.Marshal(order)
	assert.NoError(t, err)

	createdOrder, err := api.CreateOrder(order)
	assert.NoError(t, err)
	assert.NotEmpty(t, createdOrder)

	list, err := api.GetOrderList(
		time.Now().Add(-1*30*24*time.Hour).Unix(),
		time.Now().Unix())
	assert.NoError(t, err)
	assert.NotEmpty(t, list)

	orderDetails, err := api.GetOrderDetails("", "VCG-ORDER-wqefmolseemopptuymo")
	assert.NoError(t, err)
	assert.NotEmpty(t, orderDetails)

	refundItem := RefundRequest{"19c4321a-28f0-4903-a379-d55650580c11", 4, "comment"}
	refundOrder, err := api.Refund(refundItem)
	assert.NoError(t, err)
	assert.NotEmpty(t, refundOrder)

	updateStatusReq := UpdateOrderStatusRequest{"", "MP"}
	newStatus, err := api.SetOrderNewStatus(updateStatusReq)
	assert.NoError(t, err)
	assert.NotEmpty(t, newStatus)

	draft, err := api.CreateOrderDraft(order)
	assert.NoError(t, err)
	assert.NotEmpty(t, draft)

	updatedDraft, err := api.UpdateOrderDraft(order)
	assert.NoError(t, err)
	assert.NotEmpty(t, updatedDraft)

	err1 := api.DeleteOrderDraft("c1547085-411b-4d77-9344-8c74cd946aa5")
	assert.NoError(t, err1)

	//sentDraft, err := api.SendOrderDraft("5182c7bf-6c54-4bdf-a59c-6e147f4082ee")
	//assert.NoError(t, err)
	//assert.NotEmpty(t, sentDraft)

	createdExternalPurchaseOrder, err := api.CreateExternalPurchaseOrder(extPurchaseOrder)
	assert.NoError(t, err)
	assert.NotEmpty(t, createdExternalPurchaseOrder)

	orderTariff, err := api.GetOrderTariff("fe1380f2-a148-4683-95a4-56af5627678f")
	assert.NoError(t, err)
	assert.NotEmpty(t, orderTariff)
}
