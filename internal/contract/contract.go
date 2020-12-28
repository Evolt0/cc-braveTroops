package contract

import (
	"github.com/Parker-Yang/cc-braveTroops/internal/contract/base"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

type Contract struct{}

func New() *Contract {
	return &Contract{}
}

const Namespace = "BASE"

var Logger = shim.NewLogger(Namespace)

func (c *Contract) Init(stub shim.ChaincodeStubInterface) peer.Response {
	Logger.Infof("contract init")
	return shim.Success(nil)
}

func (c *Contract) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	fn, args := stub.GetFunctionAndParameters()
	Logger.Infof("contract invoke: fn = %s", fn)
	switch fn {
	case "PutState":
		return base.PutState(stub, args)
	case "GetState":
		return base.GetState(stub, args)
	default:
		return shim.Error("unsupported fn")
	}
}
