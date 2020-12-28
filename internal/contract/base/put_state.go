package base

import (
	"encoding/json"
	"fmt"

	"github.com/Parker-Yang/def-braveTroops/proto"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

func PutState(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {
		return shim.Error(fmt.Sprintf("failed to get state, incorrect number of arguments: %v", len(args)))
	}
	req := &proto.PutState{}
	err := json.Unmarshal([]byte(args[0]), req)
	if err != nil {
		return shim.Error(fmt.Sprintf("failed to get state, failed to unmarshal request: %v", err))
	}
	err = stub.PutState(req.Key, []byte(req.Value))
	if err != nil {
		return shim.Error(fmt.Sprintf("failed to get value, failed to call stub.PutState: %v", err))
	}
	return shim.Success(nil)
}
