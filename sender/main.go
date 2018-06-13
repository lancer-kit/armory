package sender

type LetterType int

const (
	LetterUniversal LetterType = 1 + iota
	LetterAdminSignUp
	LetterUserEmailVerify
	LetterUserRecovery
	LetterUserNewDevice
)

const NATSTopic = "sender.letters"

type Message struct {
	Type LetterType `json:"type"`
	Data MsgData    `json:"data"`
}

type MsgData struct {
	Base      Base      `json:"singUp,omitempty"`
	Device    Device    `json:"device,omitempty"`
	Universal Universal `json:"universal,omitempty"`
}

type Base struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Link     string `json:"link"`
}

type Device struct {
	Device   string `json:"device"`
	Location string `json:"location"`
	Ip       string `json:"ip"`
}

type Universal struct {
	Email   string `json:"email"`
	Subject string `json:"subject"`
	Text    string `json:"text"`
	HTML    string `json:"html"`
}
