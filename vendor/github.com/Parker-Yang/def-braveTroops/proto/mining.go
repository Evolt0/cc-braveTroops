package proto

// 挖矿模型
type Mining struct {
	ID string `json:"id"`
	// 用户id
	UID string `json:"uID"`
	// 计算结果hash
	ResultHash string `json:"resultHash"`
	// 计算结果
	Result string `json:"result"`
	// 产生时间
	CreateAt int64 `json:"createAt"`
}

// 挖矿请求
type MiningReq struct {
	// 计算结果hash
	ResultHash string `json:"resultHash" binding:"required"`
	// 计算结果
	Result string `json:"result" binding:"required"`
	BodyData
}

// 挖矿列表
type ListMining struct {
	List []Mining `json:"list"`
}

type Target struct {
	Prefix  string `json:"prefix"`
	NumZero int64  `json:"numZero"`
}
