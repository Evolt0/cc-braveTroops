package proto

// 挖矿模型
type Mining struct {
	ID string `json:"id"`
	//对象类型定义
	ObjectType string `json:"objectType"`
	// 用户id
	UID string `json:"uID"`
	// 用户名
	UName string `json:"uName"`
	// 计算结果hash
	ResultHash string `json:"resultHash"`
	// 计算结果
	Result string `json:"result"`
	// 产生时间
	CreateAt int64 `json:"createAt"`
	// 目标前缀
	TargetPrefix string `json:"targetPrefix"`
	// 目标零数
	TargetNumZero int64 `json:"targetNumZero"`
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

// 挖矿目标
type Target struct {
	// 前缀
	Prefix string `json:"prefix"`
	// 零数
	NumZero int64 `json:"numZero"`
	// 更新时间
	LastUpdateAt int64 `json:"lastUpdateAt"`
}
