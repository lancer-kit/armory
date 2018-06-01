package cards

import "gitlab.inn4science.com/vcg/go-common/ith/auth"

type API struct {
	auth.API
}

func (api *API) GetCardList() (*CardListResult, error) {
	return nil, nil
}

func (api *API) VerifyCard(request VerifyCardRequest) (*CardListResult, error) {
	return nil, nil
}
