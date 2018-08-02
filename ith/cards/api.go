// Cards Service is used to get linked cards for specified account.
package cards

import (
	"gitlab.inn4science.com/vcg/go-common/ith"
)

const (
	APIGetLinkedCards  = "/commonapi/account/creditcard"
	APIVerifyCard      = "/commonapi/creditcard/verify"
	GetBankAccountList = "/commonapi/bankaccount/list"
	AddBankAccount     = "/commonapi/bankaccount/add"
	EditBankAccount    = "/commonapi/bankaccount/edit"
	DeleteBankAccount  = "/commonapi/bankaccount/delete"
	SetPrimaryBankAccount = "/commonapi/bankaccount/primary"
)

type CardListResult struct {
	ErrorData *ith.ErrorData `json:"errorData,omitempty"`
	CardList  []Card         `json:"cardList,omitempty"`
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

type Country struct {
	Code                   string `json:"code"`                   //String(2); ISO country code
	Name                   string `json:"name"`                   //String(255); ISO country code
	BradnsedCardsAvaliable bool   `json:"bradnsedCardsAvaliable"` //Branded cards are supported for this country
	RegistrationAllowed    bool   `json:"registrationAllowed"`    //Registration from this country is supported
}

type BankAccount struct {
	PublicId                 string  `json:"publicId, omitempty"`      // String(36); Bank account UID
	AccountNumber            string  `json:"accountNumber"`            // String(255); Bank account number
	BankName                 string  `json:"bankName"`                 // String(255); Bank name
	HolderName               string  `json:"holderName"`               // String(255); Bank account holder name
	SwiftCode                string  `json:"swiftCode"`                // String(255); Bank SWIFT code
	Country                  Country `json:"country"`                  //Bank country object
	Type                     string  `json:"type"`                     //String(1); Bank account type: I – Internal (shown in customer UI), E – External
	HolderAddress            string  `json:"holderAddress"`            //String(70); Bank account holder address
	HolderCountry            Country `json:"holderCountry"`            //Bank account holder country object
	CorrespondentBankDetails string  `json:"correspondentBankDetails"` //String(500); Correspondent bank details

}
type ErrorData struct {
	ith.ErrorData
}

type BankAccountList struct {
	ErrorData         *ErrorData    `json:"errorData,omitempty"`
	BankAccountList   []BankAccount `json:"bankAccountList,omitempty"`
	NewBankAccountUid string        `json:"NewBankAccountUid,omitempty"` //String(36); UID of added bank account
}

type BankAccountRequest struct {
	BankAccountUid           string `json:"bankAccountUid, omitempty"` // String(36); Bank account UID
	AccountNumber            string `json:"accountNumber"`             // String(255); Bank account number
	BankName                 string `json:"bankName"`                  // String(255); Bank name
	HolderName               string `json:"holderName"`                // String(255); Bank account holder name
	SwiftCode                string `json:"swiftCode"`                 // String(255); Bank SWIFT code
	CountryCode              string `json:"countryCode"`               //String(2); Bank country ISO code
	Type                     string `json:"type"`                      //String(1); Bank account type: I – Internal (shown in customer UI), E – External
	HolderAddress            string `json:"holderAddress"`             //String(70); Bank account holder address
	HolderCountryCode        string `json:"holderCountryCode"`         //String(2); Bank account holder country ISO code
	CorrespondentBankDetails string `json:"correspondentBankDetails"`  //Correspondent bank details
}

type AccountUidRequest struct {
	BankAccountUid string `json:"bankAccountUid"` // String(36); Bank account UID
}

///go:generate goplater -type=CardType,CardStatus -transform=none -tprefix=false

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

//type BankAccountType int
//
//const (
//	I BankAccountType = 1 + iota
//	E
//)
