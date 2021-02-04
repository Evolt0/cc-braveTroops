package base

import (
	"encoding/json"

	"github.com/Evolt0/def-braveTroops/consts/status"
	"github.com/Evolt0/def-braveTroops/proto"
	"github.com/Evolt0/def-braveTroops/proto/epkg"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func GetState(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	if len(args) != 1 {
		return nil, epkg.Wrapf(status.BadRequest, "failed to get State:%v", len(args))
	}
	req := &proto.GetState{}
	err := json.Unmarshal([]byte(args[0]), req)
	if err != nil {
		return nil, epkg.Wrapf(status.InternalServerError, "failed to unmarshal request: %v", err)
	}
	state, err := stub.GetState(req.Key)
	if err != nil {
		return nil, epkg.Wrapf(status.InternalServerError, "failed to get state: %v", err)
	}
	return state, nil
}
