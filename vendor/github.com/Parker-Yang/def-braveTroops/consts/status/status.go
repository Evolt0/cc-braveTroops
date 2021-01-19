package status

import "net/http"

//go:generate  gen-const-msg -t=Status -m
type Status int

const (
	// 正确响应
	Ok Status = http.StatusOK
)

const (
	// 请求参数错误
	BadRequest Status = http.StatusBadRequest
)

const (
	// 资源未找到
	NotFound Status = http.StatusNotFound
)

const (
	// 内部处理错误
	InternalServerError Status = http.StatusInternalServerError
)

const (
	// 状态冲突
	Conflict Status = http.StatusConflict
)

const (
	// 禁止访问
	Forbidden Status = http.StatusForbidden
)

const (
	// 未授权
	Unauthorized Status = http.StatusUnauthorized
)
