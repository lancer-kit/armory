package customerrors

func NewErrNotFound(body ...map[string]string) CustomError {
	newError := DefaultError{}
	newError.ErrMsg = "Resource not found"
	newError.ErrCode = "ERR_NOT_FOUND"
	newError.ErrBody = body[0]
	return &newError
}

func NewWrongSessionMark(body ...map[string]string) CustomError {
	newError := DefaultError{}
	newError.ErrMsg = "Wrong Session Mark"
	newError.ErrCode = "WRONG_SESSION_MARK"
	newError.ErrBody = body[0]
	return &newError
}

func NewUnauthorized(body ...map[string]string) CustomError {
	newError := DefaultError{}
	newError.ErrMsg = "Unauthorized"
	newError.ErrCode = "UNAUTHORIZED"
	newError.ErrBody = body[0]
	return &newError
}

func NewForbidden(body ...map[string]string) CustomError {
	newError := DefaultError{}
	newError.ErrMsg = "Access denied"
	newError.ErrCode = "FORBIDDEN"
	newError.ErrBody = body[0]
	return &newError
}

func NewUnprocessableEntity(body ...map[string]string) CustomError {
	newError := DefaultError{}
	newError.ErrMsg = "Bad post parameters"
	newError.ErrCode = "UNPROCESSABLE"
	newError.ErrBody = body[0]
	return &newError
}
