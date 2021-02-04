package base

import (
	"encoding/json"

	"github.com/Evolt0/def-braveTroops/consts/status"
	"github.com/Evolt0/def-braveTroops/proto"
	"github.com/Evolt0/def-braveTroops/proto/epkg"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func PutState(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	if len(args) != 1 {
		return nil, epkg.Wrapf(status.BadRequest, "failed to put State: %v", len(args))
	}
	req := &proto.PutState{}
	err := json.Unmarshal([]byte(args[0]), req)
	if err != nil {
		return nil, epkg.Wrapf(status.InternalServerError, "failed to unmarshal request: %v", err)
	}
	marshal, err := json.Marshal(req)
	if err != nil {
		return nil, epkg.Wrapf(status.InternalServerError, "failed to marshal request: %v", err)
	}
	err = stub.PutState(req.Key, marshal)
	if err != nil {
		return nil, epkg.Wrapf(status.InternalServerError, "failed to unmarshal request: %v", err)
	}
	result, err := epkg.WrapSucc(nil)
	if err != nil {
		return nil, epkg.Wrapf(status.InternalServerError, "failed to marshal request: %v", err)
	}
	return result, nil
}
