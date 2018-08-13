package httpx

import (
	"testing"

	"net/http"

	"bytes"

	"github.com/stretchr/testify/assert"
	"gitlab.inn4science.com/gophers/service-kit/crypto"
)

func TestXClient_Auth(t *testing.T) {
	client := XClient{}

	client.auth = false
	assert.Equal(t, false, client.Auth(), "XClient.Auth() must return actual value")
	assert.Equal(t, client.auth, client.Auth(), "XClient.Auth() must return actual value")

	client.auth = true
	assert.Equal(t, true, client.Auth(), "XClient.Auth() must return actual value")
	assert.Equal(t, client.auth, client.Auth(), "XClient.Auth() must return actual value")

}

func TestXClient_OffAuth(t *testing.T) {
	client := NewXClient()
	assert.False(t, client.Auth(), "XClient.Auth() disabled by default")

	client.auth = true
	client.OffAuth()
	assert.False(t, client.Auth(), "XClient.Auth() must be disabled")
}

func TestXClient_OnAuth(t *testing.T) {
	client := NewXClient()
	assert.False(t, client.Auth(), "XClient.Auth() disabled by default")

	client.OnAuth()
	assert.True(t, client.Auth(), "XClient.Auth() must be enabled")
}

func TestXClient_SetAuth(t *testing.T) {
	client := NewXClient()
	const service = "test_service"
	kp := crypto.RandomKP()
	client.SetAuth(service, kp)

	assert.True(t, client.Auth(), "XClient.Auth() must be enabled after XClient.SetAuth()")
	assert.Equal(t, kp, client.kp)
	assert.Equal(t, kp.Public, client.PublicKey())
}

func TestXClient_SignRequest(t *testing.T) {
	client := NewXClient()
	const service = "test_service"
	kp := crypto.RandomKP()
	client.SetAuth(service, kp)

	req, err := http.NewRequest(http.MethodGet, "http://example.com/test?user=foo", nil)
	_ = err

	req, err = client.SignRequest(req, nil)
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

	req, err = client.SignRequest(req, nil)
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
