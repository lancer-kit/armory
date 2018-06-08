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
