package httpx

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"github.com/go-chi/chi"
	"github.com/lancer-kit/armory/crypto"
	"github.com/lancer-kit/armory/log"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSXClient_Auth(t *testing.T) {
	client := SXClient{}

	client.auth = false
	assert.Equal(t, false, client.Auth(), "XClient.Auth() must return actual value")
	assert.Equal(t, client.auth, client.Auth(), "XClient.Auth() must return actual value")

	client.auth = true
	assert.Equal(t, true, client.Auth(), "XClient.Auth() must return actual value")
	assert.Equal(t, client.auth, client.Auth(), "XClient.Auth() must return actual value")

}

func TestSXClient_OffAuth(t *testing.T) {
	client := NewSXClient()
	assert.False(t, client.Auth(), "XClient.Auth() disabled by default")

	client.auth = true
	client.OffAuth()
	assert.False(t, client.Auth(), "XClient.Auth() must be disabled")
}

func TestSXClient_OnAuth(t *testing.T) {
	client := NewSXClient()
	assert.False(t, client.Auth(), "XClient.Auth() disabled by default")

	client.OnAuth()
	assert.True(t, client.Auth(), "XClient.Auth() must be enabled")
}

func TestSXClient_SetAuth(t *testing.T) {
	client := NewSXClient()
	const service = "test_service"
	kp := crypto.RandomKP()
	client = client.SetAuth(service, kp).(*SXClient)

	assert.True(t, client.Auth(), "XClient.Auth() must be enabled after XClient.SetAuth()")
	assert.Equal(t, kp, client.kp)
	assert.Equal(t, kp.Public, client.PublicKey())
}

func TestSXClient_SignRequest(t *testing.T) {
	const service = "test_service"
	var kp = crypto.RandomKP()
	client := NewSXClient().SetAuth(service, kp)

	req, err := http.NewRequest(http.MethodGet, "http://example.com/test?user=foo", nil)
	assert.Nil(t, err)

	req, err = client.SignRequest(req, nil, nil)
	assert.Nil(t, err)

	assert.NotEmpty(t, req.Header.Get(HeaderBodyHash))
	assert.NotEmpty(t, req.Header.Get(HeaderSignature))
	assert.NotEmpty(t, req.Header.Get(HeaderService))
	assert.Equal(t, service, req.Header.Get(HeaderService))
	assert.NotEmpty(t, req.Header.Get(HeaderSigner))
	assert.Equal(t, kp.Public.String(), req.Header.Get(HeaderSigner))

	ok, err := client.VerifyRequest(req, kp.Public.String())
	assert.Nil(t, err)
	assert.True(t, ok)

	req, err = http.NewRequest(http.MethodPost, "http://example.com/test?user=foo", bytes.NewBuffer([]byte("{}")))
	_ = err

	req, err = client.SignRequest(req, nil, nil)
	assert.Nil(t, err)

	assert.NotEmpty(t, req.Header.Get(HeaderBodyHash))
	assert.NotEmpty(t, req.Header.Get(HeaderSignature))
	assert.NotEmpty(t, req.Header.Get(HeaderService))
	assert.Equal(t, service, req.Header.Get(HeaderService))
	assert.NotEmpty(t, req.Header.Get(HeaderSigner))
	assert.Equal(t, kp.Public.String(), req.Header.Get(HeaderSigner))

	ok, err = client.VerifyRequest(req, kp.Public.String())
	assert.Nil(t, err)
	assert.True(t, ok)
}

func TestSXClient(t *testing.T) {
	type fakeData struct {
		dat string
	}

	kp := crypto.RandomKP()

	server1, client1 := createFakeService(t, "test server 1", kp)
	server2, _ := createFakeService(t, "test server 2", kp)

	go func() {
		log.Get().Info("Starting test server 1")
		if err := http.ListenAndServe(":3030", server1); err != nil {
			log.Get().WithError(err).Error("Unable to start test server 1")
		}
	}()

	go func() {
		log.Get().Info("Starting test server 2")
		if err := http.ListenAndServe(":4040", server2); err != nil {
			log.Get().WithError(err).Error("Unable to start test server 1")
		}
	}()

	time.Sleep(5 * time.Second)

	data := fakeData{"Fake data"}
	sendCorrectRequests(t, client1, 4040, data)
	sendCorrectRequests(t, client1, 4040, nil)

	sendBadRequests(t, client1, 4040, data)

}

func createFakeService(t *testing.T, name string, kp crypto.KP) (*chi.Mux, SecuredClient) {
	assertions := require.New(t)

	r := chi.NewRouter()
	client := NewSXClient().SetAuth(name, kp)
	r.Route("/test", func(r chi.Router) {

		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			ok, err := client.VerifyRequest(r, kp.Public.String())
			assertions.NoErrorf(err, "Wrong auth headers in GET request")

			if !ok {
				w.WriteHeader(http.StatusOK)
				_, _ = w.Write([]byte("GET: Wrong request headers"))
				log.Get().Infof("Get request to: \"%s\" has failed", name)
				return
			}

			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte("GET was successful"))
			log.Get().Infof("Get request to: \"%s\" was successful", name)
		})

		r.Get("/bad", func(w http.ResponseWriter, r *http.Request) {
			_, err := client.VerifyRequest(r, kp.Public.String())
			assertions.Errorf(err, "Error with bad header OK ")

			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte("Error evoked, success"))
			log.Get().Infof("Get request to: \"%s\", was successful", name)
		})

		r.Post("/", func(w http.ResponseWriter, r *http.Request) {
			ok, err := client.VerifyRequest(r, kp.Public.String())
			assertions.NoErrorf(err, "Wrong auth headers in POST request")

			if !ok {
				w.WriteHeader(http.StatusOK)
				_, _ = w.Write([]byte("POST: Wrong request headers"))
				log.Get().Infof("POST request to: \"%s\" has failed", name)
				return
			}

			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte("POST was successful"))

			log.Get().Infof("POST request to: \"%s\" was successful", name)
		})

	})

	return r, client
}

func sendCorrectRequests(t *testing.T, client SecuredClient, port int, data interface{}) {
	assertions := require.New(t)

	url := fmt.Sprintf("http://127.0.0.1:%d/test", port)

	if data != nil {
		log.Get().WithField("Testing client requests with data", data).Info("Happy flow")
	} else {
		log.Get().Infof("%s Testing client requests without data", "Happy flow")
	}

	res, err := client.GetJSON(url, nil)
	assertions.NoErrorf(err, "Error when trying to send GET request")
	resBody, _ := ioutil.ReadAll(res.Body)
	log.Get().WithField("GET response: ", string(resBody)).Info("Happy flow")

	res, err = client.PostJSON(url, data, nil)
	assertions.NoErrorf(err, "Error when trying to send POST request")
	resBody, _ = ioutil.ReadAll(res.Body)
	log.Get().WithField("POST response: ", string(resBody)).Info("Happy flow")

	res, err = client.WithCookies([]*http.Cookie{}).PostJSON(url, data, nil)
	assertions.NoErrorf(err, "Error when trying to send POST request")
	resBody, _ = ioutil.ReadAll(res.Body)
	log.Get().WithField("POST response: ", string(resBody)).Info("Happy flow")
}

func sendBadRequests(t *testing.T, iClient SecuredClient, port int, data interface{}) {
	client := iClient.(*SXClient)
	assertions := require.New(t)
	var body io.Reader = nil
	var err error
	var rawData []byte

	if data != nil {
		log.Get().WithField("Testing client requests with data", data).Info("Bad flow")
	} else {
		log.Get().Infof("%s Testing client requests without data", "Bad flow")
	}

	if data != nil {
		rawData, err = json.Marshal(data)
		if err != nil {
			return
		}
		body = bytes.NewBuffer(rawData)
	}

	url := fmt.Sprintf("http://127.0.0.1:%d/test", port)

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("http://127.0.0.1:%d/test/bad", port), body)
	assert.Nil(t, err)

	req, err = client.SignRequest(req, nil, nil)
	assertions.NoErrorf(err, "Error when trying to sign GET request")
	req.Header.Set(HeaderSignature, "bad sign")

	res, err := client.Do(req)
	assertions.NoErrorf(err, "Error when trying to send GET request")
	resBody, _ := ioutil.ReadAll(res.Body)
	log.Get().WithField("GET response: ", string(resBody)).Info("Bad flow")

	req, err = http.NewRequest(http.MethodPost, url, body)
	assertions.NoErrorf(err, "Error when trying to create POST request")

	req, err = client.SignRequest(req, rawData, nil)
	assertions.NoErrorf(err, "Error when trying to sign POST request")
	req.Header.Set(HeaderBodyHash, "bad body hash")

	res, err = client.Do(req)
	assertions.NoErrorf(err, "Error when trying to send POST request")
	resBody, _ = ioutil.ReadAll(res.Body)
	log.Get().WithField("GET response: ", string(resBody)).Info("Bad flow")
}
