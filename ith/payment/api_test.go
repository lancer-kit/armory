package payment

import (
	"log"
	"net/url"
	"testing"
	"time"

	"encoding/json"

	"fmt"

	"gitlab.inn4science.com/vcg/go-common/types/currency"
)

func TestAPI_GetAuthToken(t *testing.T) {
	api := API{}
	api.Config.BaseURL, _ = url.Parse("http://demo-api.enauda.com/")
	result, err := api.GetAuthToken(&AuthRequest{
		Username: "RdHMQMZCGZKRyXHf",
		Password: "UFByqq07dMxe7m0",
	})
	log.Print(err)
	log.Print(result)

	result, err = api.RefreshAuthToken(result.RefreshToken)
	log.Print(err)
	log.Print(result.AccessToken)

	order := &Order{
		Seller: &Account{
			Email: "vipcoin-merchant@coinfide.com",
		},
		Buyer: &Account{
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
	rawOrder, _ := json.Marshal(order)
	fmt.Println(string(rawOrder))

	list, err := api.GetOrderList(
		time.Now().Add(-1*30*24*time.Hour).Unix(),
		time.Now().Unix())
	log.Print(err)
	log.Print(list)
	api.GetOrderDetails("GetOrderDetails", "8ed61188-a43d-44f2-85b5-563df4bf92b8")
}
