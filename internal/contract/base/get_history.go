package base

import (
	"encoding/json"
	"fmt"

	"github.com/Parker-Yang/def-braveTroops/proto"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

func GetHistoryState(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {
		return shim.Error(fmt.Sprintf("failed to get history, incorrect number of arguments: %v", len(args)))
	}
	req := &proto.GetHistoryRequest{}
	err := json.Unmarshal([]byte(args[0]), req)
	if err != nil {
		return shim.Error(fmt.Sprintf("failed to get history, failed to unmarshal request: %v", err))
	}
	iter, err := stub.GetHistoryForKey(req.Key)
	if err != nil {
		return shim.Error(fmt.Sprintf("failed to get history, failed to call stub.GetHistoryForKey: %v", err))
	}
	defer iter.Close()
	resp := &proto.GetHistoryResponse{}
	for iter.HasNext() {
		km, err := iter.Next()
		if err != nil {
			return shim.Error(fmt.Sprintf("failed to get history, failed to iterate next: %v", err))
		}
		keyModification := &proto.KeyModification{
			TxId:      km.TxId,
			Value:     string(km.Value),
			Timestamp: km.Timestamp.GetSeconds(),
			IsDelete:  km.IsDelete,
		}
		resp.KeyModifications = append(resp.KeyModifications, keyModification)
	}
	payload, err := json.Marshal(resp)
	if err != nil {
		return shim.Error(fmt.Sprintf("failed to get history, failed to marshal response: %v", err))
	}
	return shim.Success(payload)
}
