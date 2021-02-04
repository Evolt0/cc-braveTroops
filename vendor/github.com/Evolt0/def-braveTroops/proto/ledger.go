package proto

// 账目
type Amounts struct {
	// 账本id
	ID string `json:"id"`
	//对象类型定义
	ObjectType string `json:"objectType"`
	// 发起人Sender
	SID string `json:"sID"`
	// 收款人receiver
	RID string `json:"rID"`
	// 金额
	Change float64 `json:"change"`
	// 产生时间
	CreateAt int64 `json:"createAt"`
}

// 账本
type Ledger struct {
	List []Amounts `json:"list"`
}

type AmountsReq struct {
	// 收款人receiver
	RID string `json:"rID" binding:"required"`
	// 金额
	Change float64 `json:"change" binding:"required"`
	BodyData
}
