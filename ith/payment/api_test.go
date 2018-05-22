package payment

import (
	"log"
	"net/url"
	"testing"
)

func TestAPI_GetAuthToken(t *testing.T) {
	api := API{}
	api.Config.BaseURL, _ = url.Parse("http://demo-api.enauda.com/")
	result, err := api.GetAuthToken(&AuthRequest{
		Username: "RdHMQMZCGZKRyXHf",
		Password: "UFByqq07dMxe7m0",
	})
	log.Print(err)
	log.Print(result.AccessToken)
	//
	//result, err = api.RefreshAuthToken(result.RefreshToken)
	//log.Print(err)
	//log.Print(result.AccessToken)

	list, err := api.GetOrderList()
	log.Print(err)
	log.Print(list)

}
