package base

import (
	"encoding/json"
	"fmt"

	"github.com/Parker-Yang/def-braveTroops/proto"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

func GetState(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {
		return shim.Error(fmt.Sprintf("failed to get state, incorrect number of arguments: %v", len(args)))
	}
	req := &proto.GetState{}
	err := json.Unmarshal([]byte(args[0]), req)
	if err != nil {
		return shim.Error(fmt.Sprintf("failed to get state, failed to unmarshal request: %v", err))
	}
	state, err := stub.GetState(req.Key)
	if err != nil {
		return shim.Error(fmt.Sprintf("failed to get value, failed to call stub.GetState: %v", err))
	}
	return shim.Success(state)
}
