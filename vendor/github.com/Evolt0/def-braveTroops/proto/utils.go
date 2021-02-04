package proto

type BodyData struct {
	// 发起请求的用户签名
	Sign string `json:"sign" binding:"required"`
	// 发起请求的用户id
	ID string `json:"id" binding:"required"`
	// 操作时间戳
	Timestamp int64 `json:"timestamp" binding:"required"`
}
