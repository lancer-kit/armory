package natsx

type Message struct {
	EventID string      `json:"eventId"`
	Result  string      `json:"result"`
	Msg     string      `json:"msg"`
	Details interface{} `json:"details"`
}
