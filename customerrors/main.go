package customerrors

type CustomError interface {
	GetErrMessage() string
	GetErrCode() string
	GetErrBody() map[string]string
	GetAll() interface{}

	Init(code, msg string, body map[string]string)

	Error() string
}

type DefaultError struct {
	ErrCode string            `json:"errcode"`
	ErrMsg  string            `json:"errmsg"`
	ErrBody map[string]string `json:"errbody,omitempty"`
}

func (e *DefaultError) GetErrMessage() string {
	return e.ErrMsg
}

func (e *DefaultError) GetErrCode() string {
	return e.ErrCode
}

func (e *DefaultError) GetErrBody() map[string]string {
	return e.ErrBody
}

func (e *DefaultError) Init(code, msg string, body map[string]string) {
	e.ErrCode = code
	e.ErrMsg = msg
	e.ErrBody = body
}

func (e *DefaultError) Error() string {
	return e.ErrMsg
}

func (e *DefaultError) GetAll() interface{} {
	return &e
}
