package status

import "net/http"

//go:generate  Ytools -t=Status
type Status int

const (
	// Ok 正确响应
	Ok Status = http.StatusOK
)

const (
	// BadRequest 请求参数错误
	BadRequest Status = http.StatusBadRequest
)

const (
	// NotFound 资源未找到
	NotFound Status = http.StatusNotFound
)

const (
	// InternalServerError 内部处理错误
	InternalServerError Status = http.StatusInternalServerError
)

const (
	// Conflict 状态冲突
	Conflict Status = http.StatusConflict
)

const (
	// Forbidden 禁止访问
	Forbidden Status = http.StatusForbidden
)

const (
	// Unauthorized 未授权
	Unauthorized Status = http.StatusUnauthorized
)
