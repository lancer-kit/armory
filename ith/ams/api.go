package ams

import (
	"encoding/json"
	"net/http"

	"fmt"

	"github.com/pkg/errors"
	"gitlab.inn4science.com/vcg/go-common/api/httpx"
	"gitlab.inn4science.com/vcg/go-common/ith"
	"gitlab.inn4science.com/vcg/go-common/log"
)

const (
	APICreate = "/partnerapi/account/register"
	APIupdate = "/partnerapi/account/update"
)

type (
	Config struct {
		BaseURL string
		Client  string //partner API client uid
		Secret  string //partner API secret
	}

	API struct {
		Config Config
		log    log.Entry
		Auth   ith.API
	}

	//ErrorData = ith.ErrorData
)

func NewAPI(baseUrl, client, secret string) *API {

	tmp := &API{
		Config: Config{BaseURL: baseUrl, Client: client, Secret: secret},
		log:    log.Default,
	}

	return tmp
}

//Set new logger on ams.API
func (api *API) SetLogger(entry log.Entry) {
	api.log = entry
}

//CreateProfile - request partner API to create the new standard user profile
func (api *API) CreateProfile(req *UserRegistrationRequest) (usr *UserRegistrationResponse, err error, status RequestStatus) {
	var resp *http.Response
	status = RequestStatusOk
	//Set partner API access codes
	req.ClientId = api.Config.Client
	req.ClientSecret = api.Config.Secret

	err = req.Validate()
	if err != nil {
		err = errors.Wrap(err, "error validating")
		api.log.Warning("api.CreateProfile error validation error")
		status = RequestStatusValidationError
		return
	}

	resp, err = httpx.PostJSON(api.Config.BaseURL+APICreate, req, nil)
	if err != nil {
		err = errors.Wrap(err, "unable to send")
		api.log.WithError(err).Warning("api.CreateProfile error")
		b, _ := json.Marshal(req)
		api.log.Debug("Request:", string(b))
		if resp != nil {
			resp.Body.Read(b)
			api.log.Debug("Response:", string(b))
		}
		status = RequestStatusNetworkError
		return
	}

	if resp == nil {
		err = errors.New("empty http response")
		api.log.WithError(err).Warning("api.CreateProfile error")
		status = RequestStatusNetworkError
		return
	}

	if resp.StatusCode != 200 {
		err = errors.Errorf("response code:%d on request create user profile", resp.StatusCode)
		api.log.
			WithError(err).
			WithField("response", resp).
			WithField("url", resp.Request.URL.String()).
			Warning("api.CreateProfile error")
		status = RequestStatusNetworkError
		return
	}

	usr = new(UserRegistrationResponse)
	err = httpx.ParseJSONResult(resp, usr)
	if err != nil {
		usr = nil
		err = errors.Wrap(err, "unable to unmarshal response")
		api.log.WithError(err).Warning()
		status = RequestStatusPartnerError
		return
	}

	if usr.ErrorData != nil {
		b, _ := json.MarshalIndent(usr.ErrorData, "", "  ")
		err = errors.New("partner response error:" + fmt.Sprint(usr.ErrorData))
		api.log.
			WithField("error-data", string(b)).
			WithError(err).
			WithField("response", resp).
			Warning("api.CreateProfile error")
		b, _ = json.MarshalIndent(req, "", "\t")
		println("Sent:\n", string(b))
		status = RequestStatusPartnerError
	}

	return
}
