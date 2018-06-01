package kyc

import (
	"net/url"
	"testing"

	"io/ioutil"

	"github.com/stretchr/testify/assert"
	"gitlab.inn4science.com/vcg/go-common/crypto"
	"gitlab.inn4science.com/vcg/go-common/ith/auth"
)

func TestAPI_GetAuthToken(t *testing.T) {
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

	api.Config.BaseURL, _ = url.Parse("http://demo-commonapi.enauda.com/")

	doc := new(Document)
	doc.Type = DocumentTypeSLF
	doc.SubType = DocumentSubTypeSLF
	doc.FileName1 = "slf.jpg"
	slfFile, err := ioutil.ReadFile("slf.jpg")
	doc.B64File1 = crypto.Base64Encode(slfFile)
	documentUploadResult, err := api.CreateDocument(doc)
	assert.NoError(t, err)
	assert.Empty(t, documentUploadResult.ErrorData)
	documentsResult, err := api.GetDocumentList()
	assert.NoError(t, err)
	assert.Empty(t, documentsResult.ErrorData)
}
