package vcg_go_common

import (
	"fmt"
	"testing"

	"github.com/inn4sc/vcg-transaction/common/currency"
	"github.com/stretchr/testify/assert"
)

func TestVerifySignature(t *testing.T) {
	privKey, pubKey := GenKeyPair()
	fmt.Println("Private Key: ", privKey)
	fmt.Println("Public  Key: ", pubKey)

	message := fmt.Sprintf("%s:%s", currency.Amount(4212340000).String(), "test 42")
	fmt.Println(message)
	sig, err := SignMessage(privKey, message)
	assert.Equal(t, nil, err)

	fmt.Println("Signature: ", sig)

	ok, err := VerifySignature(pubKey, message, sig)
	assert.Equal(t, nil, err)
	assert.Equal(t, true, ok)
}
