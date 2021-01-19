package contract

import (
	"fmt"
	"github.com/Parker-Yang/cc-braveTroops/internal/contract/base"
	"github.com/Parker-Yang/cc-braveTroops/internal/contract/transfer"
	"github.com/Parker-Yang/def-braveTroops/proto/epkg"
	"github.com/Parker-Yang/def-braveTroops/proto/prefix"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

type Contract struct{}

func New() *Contract {
	return &Contract{}
}

var Logger = shim.NewLogger(prefix.BraveTroops)

func (c *Contract) Init(stub shim.ChaincodeStubInterface) peer.Response {
	Logger.Infof("contract init")
	return shim.Success(nil)
}

func (c *Contract) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	fn, args := stub.GetFunctionAndParameters()
	Logger.Infof("contract invoke: fn = %s", fn)
	var (
		payload []byte
		err     error
	)
	switch fn {
	case prefix.FnPutState:
		payload, err = base.PutState(stub, args)
	case prefix.FnGetState:
		payload, err = base.GetState(stub, args)
	case prefix.FnTransfer:
		payload, err = transfer.Transfer(stub, args)
	case prefix.FnGetHistoryState:
		return base.GetHistoryState(stub, args)
	default:
		err = fmt.Errorf("unsupported fn")
	}
	if err != nil {
		return shim.Error(epkg.WrapFailV2(err))
	}
	return shim.Success(payload)
}
