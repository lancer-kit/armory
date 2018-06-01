package kyc

import (
	"fmt"

	"github.com/pkg/errors"
	"gitlab.inn4science.com/vcg/go-common/api/httpx"
	"gitlab.inn4science.com/vcg/go-common/ith/auth"
)

type API struct {
	auth.API
}

func (api *API) CreateDocument(document *Document) (*CreateDocumentRequest, error) {
	u := api.Config.GetURL(APIDocumentUpload)

	httpResp, err := httpx.PostJSON(
		u.String(),
		document,
		map[string]string{
			auth.Header: auth.HeaderVal(api.Credentials.AccessToken),
		})
	if err != nil {
		return nil, errors.Wrap(err, "failed to refresh auth token")
	}

	if httpResp.StatusCode != 200 {
		return nil, fmt.Errorf("request failed with status - %d", httpResp.StatusCode)
	}

	response := new(CreateDocumentRequest)
	err = httpx.ParseJSONResult(httpResp, response)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse response")
	}
	return response, err
}

func (api *API) GetDocumentList() (*ListResp, error) {
	u := api.Config.GetURL(APIDocumentList)
	req := &struct {
	}{}

	httpResp, err := httpx.PostJSON(
		u.String(),
		req,
		map[string]string{
			auth.Header: auth.HeaderVal(api.Credentials.AccessToken),
		})
	if err != nil {
		return nil, errors.Wrap(err, "failed to get documents list")
	}

	if httpResp.StatusCode != 200 {
		return nil, fmt.Errorf("request failed with status - %d", httpResp.StatusCode)
	}

	response := new(ListResp)
	err = httpx.ParseJSONResult(httpResp, response)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse response")
	}
	return response, err
}
