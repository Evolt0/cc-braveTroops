package mining

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"github.com/Evolt0/cc-braveTroops/pkg"
	"github.com/Evolt0/def-braveTroops/consts/status"
	"github.com/Evolt0/def-braveTroops/proto"
	"github.com/Evolt0/def-braveTroops/proto/epkg"
	"github.com/Evolt0/def-braveTroops/proto/prefix"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"strings"
)

func PoW(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	req := &proto.MiningReq{}
	err := pkg.Decode(args, req)
	if err != nil {
		return nil, err
	}
	user, err := pkg.Check(stub, req, &req.BodyData)
	if err != nil {
		return nil, err
	}
	result, err := execPoW(stub, user, req)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func execPoW(stub shim.ChaincodeStubInterface, user *proto.User, req *proto.MiningReq) ([]byte, error) {
	result, err := verifyHash(stub, req)
	if err != nil {
		return nil, err
	}
	user.Balance += 50
	userKey, err := pkg.NewCompositeKey(stub, prefix.User, user.ID)
	if err != nil {
		return nil, err
	}
	_, err = pkg.UpdateState(stub, userKey, user)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func verifyHash(stub shim.ChaincodeStubInterface, data *proto.MiningReq) ([]byte, error) {
	key, err := pkg.NewCompositeKey(stub, prefix.Target, "")
	if err != nil {
		return nil, err
	}
	state, err := stub.GetState(key)
	if err != nil {
		return nil, epkg.Wrapf(status.InternalServerError, "failed to put: %v", err)
	}
	target := &proto.Target{}
	err = json.Unmarshal(state, target)
	if err != nil {
		return nil, epkg.Wrapf(status.InternalServerError, "failed to unmarshal: %v", err)
	}
	if !(data.ResultHash[:target.NumZero] == strings.Repeat("0", int(target.NumZero))) {
		return nil, epkg.Wrapf(status.BadRequest, "failed to mining: bad target")
	}
	hash := sha256.Sum256([]byte(data.Result))
	result := hex.EncodeToString(hash[:])
	if data.ResultHash != result {
		return nil, epkg.Wrapf(status.BadRequest, "failed to mining: bad hash")
	}
	target.LastUpdateAt = pkg.GetTimestamp(stub)
	marshal, err := pkg.UpdateState(stub, key, target)
	if err != nil {
		return nil, err
	}
	return marshal, nil
}
