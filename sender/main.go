// sender is a client package for sender service.
// It allow to make and send emails.
// `Message` can be sent through the HTTP or NATS.
package sender

// LetterType is an enum of predefined letter templates.
type LetterType int

const (
	LetterUniversal LetterType = 1 + iota
	LetterAdminSignUp
	LetterUserEmailVerify
	LetterUserRecovery
	LetterUserNewDevice
)

// NATSTopic is a topic in NATS, through which the sender receives new messages.
const NATSTopic = "sender.letters"

// HTTPURL is URL path in which the sender receives new messages.
const HTTPURL = "/v1/email"

// Message is the data for some template.
// The type field indicates which template will be sent.
type Message struct {
	// Type indicates which template will be used.
	Type LetterType `json:"type"`
	// Data to fill in the template, depends on the `Type`.
	Data MsgData `json:"data"`
}

// MsgData data for letter templates.
type MsgData struct {
	Base      Base      `json:"base,omitempty"`
	Device    Device    `json:"device,omitempty"`
	Universal Universal `json:"universal,omitempty"`
}

// Base is a structure for the base letter template.
type Base struct {
	// Email is a addressee of the letter.
	Email    string `json:"email"`
	Username string `json:"username"`
	Link     string `json:"link"`
}

// Device is a data extension for the `new device` letter.
type Device struct {
	Device   string `json:"device"`
	Location string `json:"location"`
	Ip       string `json:"ip"`
}

// Universal is a message that does not have a template
// in the sender and must be sent as is. Data from the another fields
// of the MsgData will be ignored.
type Universal struct {
	Email   string `json:"email"`
	Subject string `json:"subject"`
	Text    string `json:"text"`
	HTML    string `json:"html"`
}
