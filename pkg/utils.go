package pkg

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/Evolt0/def-braveTroops/consts/status"
	"github.com/Evolt0/def-braveTroops/proto"
	"github.com/Evolt0/def-braveTroops/proto/epkg"
	"github.com/Evolt0/def-braveTroops/proto/pkg/order"
	"github.com/Evolt0/def-braveTroops/proto/prefix"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

var (
	Logger = shim.NewLogger(prefix.BraveTroops)
	CST    = time.FixedZone("CST", 8*60*60)
)

// 获取区块链时间戳
func GetTimestamp(stub shim.ChaincodeStubInterface) int64 {
	t, err := stub.GetTxTimestamp()
	if err != nil {
		return 0
	}
	return t.Seconds
}

// 生成随机id
func NewUUID(stub shim.ChaincodeStubInterface, salt string) string {
	data := stub.GetTxID() + salt
	hash := sha256.Sum256([]byte(data))
	return fmt.Sprintf("%x", hash)
}

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

func PutState(stub shim.ChaincodeStubInterface, key string, data interface{}) ([]byte, error) {
	Logger.Infof("Put state")
	state, err := stub.GetState(key)
	if err != nil {
		return nil, err
	}
	if len(state) != 0 {
		return nil, epkg.Wrapf(status.BadRequest, key+" data already exist")
	}
	marshal, err := json.Marshal(data)
	if err != nil {
		return nil, epkg.Wrapf(status.InternalServerError, "failed to marshal "+err.Error())
	}
	err = stub.PutState(key, marshal)
	if err != nil {
		return nil, epkg.Wrapf(status.InternalServerError, "failed to put "+err.Error())
	}
	return marshal, nil
}

// 更新数据
func UpdateState(stub shim.ChaincodeStubInterface, key string, data interface{}) ([]byte, error) {
	Logger.Infof("Update state")
	state, err := stub.GetState(key)
	if err != nil {
		return nil, err
	}
	if len(state) == 0 {
		return nil, epkg.Wrapf(status.BadRequest, key+" data not exist")
	}
	marshal, err := json.Marshal(data)
	if err != nil {
		return nil, epkg.Wrapf(status.InternalServerError, "failed to marshal "+err.Error())
	}
	err = stub.PutState(key, marshal)
	if err != nil {
		return nil, epkg.Wrapf(status.InternalServerError, "failed to put "+err.Error())
	}
	return marshal, nil
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
	/*err = RSA.Verify(data.Sign, reqUser.PubKey, ordered)
	if err != nil {
		return nil, epkg.Wrapf(status.BadRequest, err.Error())
	}
	Logger.Infof("Verify Sign")*/
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

// 初始化工作量证明target
func InitTarget(stub shim.ChaincodeStubInterface) error {
	key, err := NewCompositeKey(stub, prefix.Target, "")
	if err != nil {
		return err
	}
	target := &proto.Target{
		Prefix:       prefix.BraveTroops + strconv.FormatInt(GetTimestamp(stub), 10),
		NumZero:      1 << 2,
		LastUpdateAt: GetTimestamp(stub),
	}
	marshal, err := json.Marshal(target)
	if err != nil {
		return epkg.Wrapf(status.InternalServerError, "failed to marshal : %v", err)
	}
	err = stub.PutState(key, marshal)
	if err != nil {
		return epkg.Wrapf(status.InternalServerError, "failed to put: %v", err)
	}
	return nil
}
