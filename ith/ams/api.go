package ams

import (
	"encoding/json"
	"net/http"

	"fmt"

	"strings"

	"github.com/pkg/errors"
	"gitlab.inn4science.com/vcg/go-common/api/httpx"
	"gitlab.inn4science.com/vcg/go-common/ith"
	"gitlab.inn4science.com/vcg/go-common/log"
)

const (
	APICreate = "/partnerapi/account/register"
	APIupdate = "/partnerapi/account/update"
	APICode   = "/commonapi/auth/" + param + "/authorization_code"
	APIToken  = "/partnerapi/token/code"
	param     = "{{.client}}"
)

type (
	Config struct {
		BaseURL   string
		CommonURL string
		Client    string //partner API client uid
		Secret    string //partner API secret
	}

	API struct {
		Config Config
		log    log.Entry
		Auth   ith.API
	}

	//ErrorData = ith.ErrorData
)

func NewAPI(baseUrl, commonUrl, client, secret string) *API {

	tmp := &API{
		Config: Config{
			BaseURL:   baseUrl,
			CommonURL: commonUrl,
			Client:    client,
			Secret:    secret,
		},
		log: log.Default,
	}
	if string(tmp.Config.BaseURL[len(tmp.Config.BaseURL)-1:]) == "/" {
		tmp.Config.BaseURL = string(tmp.Config.BaseURL[:len(tmp.Config.BaseURL)-1])
	}
	if string(tmp.Config.CommonURL[len(tmp.Config.CommonURL)-1:]) == "/" {
		tmp.Config.CommonURL = string(tmp.Config.CommonURL[:len(tmp.Config.CommonURL)-1])
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
		api.log.WithError(err).Error()
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

//CreateProfile - request partner API to update the standard user profile
func (api *API) UpdateProfile(req *UserUpdateRequest, token string) (usr *UserRegistrationResponse, err error, status RequestStatus) {
	var resp *http.Response
	status = RequestStatusOk
	//Set partner API access codes
	req.ClientId = api.Config.Client
	req.ClientSecret = api.Config.Secret

	err = req.Validate()
	if err != nil {
		err = errors.Wrap(err, "error validating")
		api.log.Warning("api.UpdateProfile error validation error")
		status = RequestStatusValidationError
		return
	}
	headerMap := make(map[string]string)
	headerMap["Authorization"] = "Bearer " + token
	resp, err = httpx.PostJSON(api.Config.BaseURL+APIupdate, req, headerMap)
	if err != nil {
		err = errors.Wrap(err, "unable to send")
		api.log.WithError(err).Warning("api.UpdateProfile error")
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
		api.log.WithError(err).Warning("api.UpdateProfile error")
		status = RequestStatusNetworkError
		return
	}

	if resp.StatusCode != 200 {
		err = errors.Errorf("response code:%d on request create user profile", resp.StatusCode)
		api.log.
			WithError(err).
			WithField("response", resp).
			WithField("url", resp.Request.URL.String()).
			Warning("api.UpdateProfile error")
		status = RequestStatusNetworkError
		return
	}

	usr = new(UserRegistrationResponse)
	err = httpx.ParseJSONResult(resp, usr)
	if err != nil {
		usr = nil
		err = errors.Wrap(err, "unable to unmarshal response")
		api.log.WithError(err).Error()
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
			Warning("api.UpdateProfile error")
		b, _ = json.MarshalIndent(req, "", "\t")
		println("Sent:\n", string(b))
		status = RequestStatusPartnerError
	}

	return
}

//Get Authorization Code
//Send request to ITH.Authorization service.
//Service is used to receive one-time authorization code. This single use code could be used
//to transfer user session from one module to another.
func (api *API) GetCode(req *AuthCodeRequest) (code *AuthCodeResponse, err error) {
	//--- Get Authorization Code---
	err = req.Validate()
	if err != nil {
		return
	}
	url := strings.Replace(api.Config.CommonURL+APICode, param, api.Config.Client, -1)
	api.log.Debug("URL:", url)
	b, _ := json.MarshalIndent(req, "", "\t")
	api.log.Debug(string(b))
	//return nil, errors.New("STOP")
	api.log.Debug("URL:", url)
	var resp *http.Response
	resp, err = httpx.PostJSON(url, req, nil)
	if err != nil {
		err = errors.Wrap(err, "unable to get one-time authorization code")
		api.log.WithError(err).Error("api.GetCode error")
		return
	}

	if resp == nil {
		err = errors.Wrap(err, "unable to get one-time authorization code. empty response")
		api.log.WithError(err).Error("api.GetCode error")
		return
	}

	if resp.StatusCode != 200 {
		err = errors.New("unable to get one-time authorization code. response code:" + fmt.Sprint(resp.StatusCode))
		api.log.WithError(err).Error("api.GetCode error")
		return
	}
	code = new(AuthCodeResponse)
	err = httpx.ParseJSONResult(resp, code)
	if err != nil {
		errors.Wrap(err, "unable to parse code response json")
		api.log.WithError(err).Error("api.GetCode error")
	}

	return
}

//GetToken - Get Authorization Token
//Send request to ITH.Authorization service.
//Service is used to receive customer access token and refresh token using one-time
//authorization code. Received access token should be used in other services calls.
func (api *API) GetToken(req *AuthCodeRequest) (token *AuthTokenResponse, err error) {
	//--- Get Authorization Code---
	var respCode *AuthCodeResponse
	respCode, err = api.GetCode(req)
	if err != nil {
		return
	}
	//--- Get Token by Authorization Code ---
	url := api.Config.BaseURL + APIToken
	api.log.Debug("URL:", url)
	reqToken := AuthTokenRequest{
		ClientId:     api.Config.Client,
		ClientSecret: api.Config.Secret,
		Code:         respCode.Code,
	}
	var resp *http.Response
	resp, err = httpx.PostJSON(url, reqToken, nil)

	if err != nil {
		err = errors.Wrap(err, "unable to get authorization tokens")
		api.log.WithError(err).Error("api.GetToken error")
		return
	}

	if resp.StatusCode != 200 {
		err = errors.New("unable to get one-time authorization code. response code:" + fmt.Sprint(resp.StatusCode))
		api.log.WithError(err).Error("api.GetToken error")
		return
	}
	token = new(AuthTokenResponse)
	err = httpx.ParseJSONResult(resp, token)
	if err != nil {
		errors.Wrap(err, "unable to parse token response json")
		api.log.WithError(err).Error("api.GetToken error")
		return
	}

	if token.ErrorData != nil {
		err = errors.New("partner response with error:" + token.ErrorData.ErrorMessage)
		api.log.WithError(err).Error("api.GetToken error")
	}

	return
}
