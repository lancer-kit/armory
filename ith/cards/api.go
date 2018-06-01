// Cards Service is used to get linked cards for specified account.
package cards

import "gitlab.inn4science.com/vcg/go-common/ith"

const (
	APIGetLinkedCards = "/commonapi/account/creditcard"
	APIVerifyCard     = "/ccommonapi/creditcard/verify"
)

type CardListResult struct {
	ErrorData ith.ErrorData `json:"errorData,omitempty"`
	CardList  []Card        `json:"cardList,omitempty"`
}

type VerifyCardRequest struct {
	CardUID          string `json:"cardUid"`          // String(36); Linked card UID
	VerificationCode string `json:"verificationCode"` // String(8); Card verification code from bank statement
}

type Card struct {
	UID          string     `json:"uid"`          // String(36); Card UID
	CardHolder   string     `json:"cardHolder"`   // String(50); Card holder name
	MaskedNumber string     `json:"maskedNumber"` // String(19); Card masked number
	ExpMonth     int        `json:"expMonth"`     // Card expiration month
	ExpYear      int        `json:"expYear"`      // Card expiration year
	Type         CardType   `json:"type"`         // String(1); Card type
	Status       CardStatus `json:"status"`       // String(1); Card status
	Primary      bool       `json:"primary"`      // Card is set as primary
}

//go:generate goplater -type=CardType,CardStatus -transform=none -tprefix=false

// Card type:
// A – American Express
// M – MasterCard
// V – Visa
type CardType int

const (
	CardTypeAmExpress CardType = iota + 1
	CardTypeMasterCard
	CardTypeVisa
)

// Card status:
// E – Expired
// N – Verification Required
// V – Verified
type CardStatus int

const (
	CardStatusExpired CardStatus = iota + 1
	CardStatusNotVerified
	CardStatusVerified
)
