package proto

// 用户模型
type User struct {
	// 用户id（根据用户公钥Hash获得）
	ID string `json:"id"`
	// 用户名
	Name string `json:"name"`
	//对象类型定义
	ObjectType string `json:"objectType"`
	// 用户公钥
	PubKey string `json:"pubKey"`
	// 账户余额
	Balance float64 `json:"balance"`
	// 产生时间
	CreateAt int64 `json:"createAt"`
}

type UserReq struct {
	// 用户id（根据用户公钥Hash获得）
	ID string `json:"id" binding:"required"`
	// 用户名
	Name string `json:"name" binding:"required"`
	// 用户公钥
	PubKey string `json:"pubKey" binding:"required"`
}

type ListUser struct {
	List []User `json:"list"`
}

type KeyPair struct {
	Pub string `json:"pub"`
	Pri string `json:"pri"`
}