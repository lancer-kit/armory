package cards

import (
	"fmt"

	"github.com/pkg/errors"
	"gitlab.inn4science.com/vcg/go-common/api/httpx"
	"gitlab.inn4science.com/vcg/go-common/ith/auth"
)

type API struct {
	auth.API
}

func (api *API) GetCardList() (*CardListResult, error) {
	err := api.EnsureAuth()
	if err != nil {
		return nil, err
	}

	httpResp, err := httpx.PostJSON(
		api.Config.GetURL(APIGetLinkedCards).String(),
		nil,
		api.AuthHeader(),
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to linked cards")
	}

	if httpResp.StatusCode != 200 {
		return nil, fmt.Errorf("request failed with status - %d", httpResp.StatusCode)
	}

	response := new(CardListResult)
	err = httpx.ParseJSONResult(httpResp, response)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse linked cards response")
	}
	return response, err
}

func (api *API) VerifyCard(request VerifyCardRequest) (*CardListResult, error) {
	err := api.EnsureAuth()
	if err != nil {
		return nil, err
	}

	httpResp, err := httpx.PostJSON(
		api.Config.GetURL(APIVerifyCard).String(),
		&request,
		api.AuthHeader(),
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to linked cards")
	}

	if httpResp.StatusCode != 200 {
		return nil, fmt.Errorf("request failed with status - %d", httpResp.StatusCode)
	}

	response := new(CardListResult)
	err = httpx.ParseJSONResult(httpResp, response)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse linked cards response")
	}
	return response, err
}

func (api *API) GetBankAccountList(request string) (*BankAccountList, error) {
	err := api.EnsureAuth()
	if err != nil {
		return nil, err
	}

	httpResp, err := httpx.PostJSON(
		api.Config.GetURL(APIGetBankAccountList).String(),
		&request,
		api.AuthHeader(),
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get account list")
	}

	if httpResp.StatusCode != 200 {
		return nil, fmt.Errorf("request failed with status - %d", httpResp.StatusCode)
	}

	response := new(BankAccountList)
	err = httpx.ParseJSONResult(httpResp, response)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse account list cards response")
	}
	return response, err
}

func (api *API) AddBankAccount(request *BankAccountRequest) (*BankAccountList, error) {
	err := api.EnsureAuth()
	if err != nil {
		return nil, err
	}

	httpResp, err := httpx.PutJSON(
		api.Config.GetURL(APIAddBankAccount).String(),
		&request,
		api.AuthHeader(),
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to add bank account")
	}

	if httpResp.StatusCode != 200 {
		return nil, fmt.Errorf("request failed with status - %d", httpResp.StatusCode)
	}

	response := new(BankAccountList)
	err = httpx.ParseJSONResult(httpResp, response)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse account list cards response")
	}
	//fmt.Println(response.BankAccountList)
	return response, err
}

func (api *API) EditBankAccount(request *BankAccountRequest) (*BankAccountList, error) {
	err := api.EnsureAuth()
	if err != nil {
		return nil, err
	}

	httpResp, err := httpx.PostJSON(
		api.Config.GetURL(APIEditBankAccount).String(),
		&request,
		api.AuthHeader(),
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to edit bank account")
	}

	if httpResp.StatusCode != 200 {
		return nil, fmt.Errorf("request failed with status - %d", httpResp.StatusCode)
	}

	response := new(BankAccountList)
	err = httpx.ParseJSONResult(httpResp, response)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse account list cards response")
	}
	return response, err
}

func (api *API) DeleteBankAccount(request AccountUidRequest) (*BankAccountList, error) {
	err := api.EnsureAuth()
	if err != nil {
		return nil, err
	}

	httpResp, err := httpx.PostJSON(
		api.Config.GetURL(APIDeleteBankAccount).String(),
		&request,
		api.AuthHeader(),
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to edit bank account")
	}

	if httpResp.StatusCode != 200 {
		return nil, fmt.Errorf("request failed with status - %d", httpResp.StatusCode)
	}

	response := new(BankAccountList)
	err = httpx.ParseJSONResult(httpResp, response)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse account list cards response")
	}
	return response, err
}

func (api *API) SetPrimaryBankAccount(request AccountUidRequest) (*BankAccountList, error) {
	err := api.EnsureAuth()
	if err != nil {
		return nil, err
	}

	httpResp, err := httpx.PostJSON(
		api.Config.GetURL(APISetPrimaryBankAccount).String(),
		&request,
		api.AuthHeader(),
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to edit bank account")
	}

	if httpResp.StatusCode != 200 {
		return nil, fmt.Errorf("request failed with status - %d", httpResp.StatusCode)
	}

	response := new(BankAccountList)
	err = httpx.ParseJSONResult(httpResp, response)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse account list cards response")
	}
	return response, err
}
