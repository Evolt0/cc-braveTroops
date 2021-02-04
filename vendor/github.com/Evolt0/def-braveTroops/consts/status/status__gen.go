// Package status const code comment msg
package status

// noErrorMsg if code is not found, GetMsg will return this
const noErrorMsg = "unknown error"

// messages get msg from const comment
var messages = map[Status]string{

	BadRequest:          "请求参数错误",
	Conflict:            "状态冲突",
	Forbidden:           "禁止访问",
	InternalServerError: "内部处理错误",
	NotFound:            "资源未找到",
	Ok:                  "正确响应",
	Unauthorized:        "未授权",
}

// String return string
func (code Status) String() string {
	return GetMsg(code)
}

// GetMsg get error msg
func GetMsg(code Status) string {
	var (
		msg string
		ok  bool
	)
	if msg, ok = messages[code]; !ok {
		msg = noErrorMsg
	}
	return msg
}

// implement
type Error interface {
	HttpErr() *HttpError
	Error() string
}

// implement method
func (code Status) HttpErr() *HttpError {
	return &HttpError{
		Code:    int(code),
		Message: GetMsg(code),
	}
}

// implement method
func (code Status) Error() string {
	return code.HttpErr().Error()
}
