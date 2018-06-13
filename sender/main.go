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

// Message is the data for some template.
// The type field indicates which template will be sent.
type Message struct {
	Type LetterType `json:"type"`
	Data MsgData    `json:"data"`
}

// MsgData data for letter templates.
type MsgData struct {
	Base      Base      `json:"singUp,omitempty"`
	Device    Device    `json:"device,omitempty"`
	Universal Universal `json:"universal,omitempty"`
}

// Base is a structure for the base letter template.
type Base struct {
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
