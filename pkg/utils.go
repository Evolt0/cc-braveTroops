package pkg

import (
	"encoding/json"
	"github.com/Parker-Yang/def-braveTroops/consts/status"
	"github.com/Parker-Yang/def-braveTroops/proto"
	"github.com/Parker-Yang/def-braveTroops/proto/epkg"
	"github.com/Parker-Yang/def-braveTroops/proto/pkg/RSA"
	"github.com/Parker-Yang/def-braveTroops/proto/pkg/order"
	"github.com/Parker-Yang/def-braveTroops/proto/prefix"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"time"
)

var (
	Logger = shim.NewLogger(prefix.BraveTroops)
	CST    = time.FixedZone("CST", 8*60*60)
)

// 获取信息
func GetState(stub shim.ChaincodeStubInterface, key string, result interface{}) error {
	Logger.Infof("Get state")
	state, err := stub.GetState(key)
	if len(state) == 0 {
		return epkg.Wrapf(status.NotFound, key+"state is empty")
	}
	if err != nil {
		return epkg.Wrapf(status.InternalServerError, "failed to get state"+err.Error())
	}
	err = json.Unmarshal(state, result)
	if err != nil {
		return epkg.Wrapf(status.InternalServerError, "failed to unmarshal "+err.Error())
	}
	return nil
}

// 解码数据
func Decode(args []string, req interface{}) error {
	Logger.Infof("Unmarshal request")
	if len(args) != 1 {
		return epkg.Wrapf(status.BadRequest, "failed to pass parameters: incorrect number of arguments: %v", len(args))
	}
	err := json.Unmarshal([]byte(args[0]), req)
	if err != nil {
		return epkg.Wrapf(status.InternalServerError, "failed to unmarshal "+err.Error())
	}
	return nil
}

// 普通用户验证
func Check(stub shim.ChaincodeStubInterface, req interface{}, data *proto.BodyData) (*proto.User, error) {
	Logger.Infof("verify sign and time")
	// 请求排序
	ordered := order.Order(req)
	Logger.Infof("Order request: %v", ordered)
	// 验证时间
	/*err := verify.Time(data.Timestamp, GetTimestamp(stub), prefix.DELTA)
	if err != nil {
		return nil, epkg.Wrapf(errors.Timeout, err.Error())
	}
	Logger.Infof("Verify time")*/
	reqUser, err := GetUser(stub, data.ID)
	if err != nil {
		return nil, err
	}
	// 验证签名
	err = RSA.Verify(data.Sign, reqUser.PubKey, ordered)
	if err != nil {
		return nil, epkg.Wrapf(status.BadRequest, err.Error())
	}
	Logger.Infof("Verify Sign")
	return reqUser, nil
}

// 获取用户
func GetUser(stub shim.ChaincodeStubInterface, id string) (*proto.User, error) {
	Logger.Infof("Get user")
	user := &proto.User{}
	userKey, err := NewCompositeKey(stub, prefix.User, id)
	if err != nil {
		return nil, err
	}
	err = GetState(stub, userKey, user)
	if err != nil {
		return nil, epkg.Wrapf(status.NotFound, err.Error())
	}
	return user, nil
}

// 生成BlockKey
func NewCompositeKey(stub shim.ChaincodeStubInterface, prefixName, id string) (string, error) {
	Logger.Infof("New composite key")
	key := []string{prefixName}
	if id != "" {
		key = append(key, id)
	}
	compositeKey, err := stub.CreateCompositeKey(prefix.BraveTroops, key)
	if err != nil {
		return "", epkg.Wrapf(status.InternalServerError, "failed to new "+prefixName+" key"+err.Error())
	}
	return compositeKey, nil
}
