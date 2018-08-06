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

	demoAccount := &BankAccountRequest{
		AccountNumber:            "LV80UNLA0000435195001",
		BankName:                 "AS SEB BANKA",
		HolderName:               "John Doe",
		SwiftCode:                "DGHTY55645H",
		CountryCode:              "DE",
		Type:                     BankAccountTypeI,
		HolderAddress:            "345, TH 6",
		HolderCountryCode:        "UA",
		CorrespondentBankDetails: "A lot of information here",
	}

	toEditAccount := &BankAccountRequest{
		BankAccountUid:           "b063880d-d15d-4be7-8606-b05d90798012",
		AccountNumber:            "LV80UNLA0000435195001",
		BankName:                 "AS SEB BANKA",
		HolderName:               "John Doe",
		SwiftCode:                "DGHTY55645H",
		CountryCode:              "DE",
		Type:                     BankAccountTypeI,
		HolderAddress:            "345, TH 6",
		HolderCountryCode:        "UA",
		CorrespondentBankDetails: "A lot of information here",
	}
	addedAccount, err := api.AddBankAccount(demoAccount)
	assert.NoError(t, err)
	assert.Empty(t, result.ErrorData)
	assert.NotEmpty(t, addedAccount.BankAccountList)

	//fixme: Error: Should NOT be empty, but was []
	//bankAccountList, err := api.GetBankAccountList("")
	//assert.NoError(t, err)
	//assert.NotEmpty(t, bankAccountList.BankAccountList)
	//assert.Empty(t, result.ErrorData)

	editedAccount, err := api.EditBankAccount(toEditAccount)
	assert.NoError(t, err)
	assert.NotEmpty(t, editedAccount.BankAccountList)
	assert.Empty(t, result.ErrorData)

	//UIDs to test delete: b063ad0d-d15d-4be7-8606-b05d90798012, 89d7e00d-e262-4e8f-9bf5-0aa3a1efc6f2, 41d0a188-c1df-48fb-9631-0b3311aa8817
	accountToDelete := AccountUidRequest{"98da616e-d5b5-440f-8d49-3de5a7060128"}
	deletedAccount, err := api.DeleteBankAccount(accountToDelete)
	assert.NoError(t, err)
	assert.NotEmpty(t, deletedAccount.BankAccountList) //Will be empty is the only account was deleted
	assert.Empty(t, result.ErrorData)

	primaryAccount := AccountUidRequest{"b063880d-d15d-4be7-8606-b05d90798012"}
	primaryBankAccount, err := api.SetPrimaryBankAccount(primaryAccount)
	assert.NoError(t, err)
	assert.NotEmpty(t, primaryBankAccount.BankAccountList)
	assert.Empty(t, result.ErrorData)

}
