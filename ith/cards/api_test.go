package cards

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.inn4science.com/vcg/go-common/ith/auth"
)

func TestAPI_GetCardList(t *testing.T) {
	api := API{}
	api.Config.BaseURL, _ = url.Parse("http://demo-api.enauda.com/")
	result, err := api.GetAuthToken(&auth.Request{
		Username: "380939440888",
		Password: "VipCoin12345",
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

	api.Config.BaseURL, _ = url.Parse("http://demo-commonapi.enauda.com/")

	documentsResult, err := api.GetCardList()
	assert.NoError(t, err)
	assert.NotEmpty(t, documentsResult)
	assert.Empty(t, result.ErrorData)
}
