package status

type Error interface {
	HttpErr() *HttpError
	Error() string
}

func (code Status) HttpErr() *HttpError {
	return &HttpError{
		Code:    int(code),
		Message: GetMsg(code),
	}
}

func (code Status) Error() string {
	return code.HttpErr().Error()
}
