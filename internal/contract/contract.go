package contract

import (
	"fmt"

	"github.com/Evolt0/cc-braveTroops/internal/contract/base"
	"github.com/Evolt0/cc-braveTroops/internal/contract/mining"
	"github.com/Evolt0/cc-braveTroops/internal/contract/transfer"
	"github.com/Evolt0/cc-braveTroops/internal/contract/user"
	"github.com/Evolt0/cc-braveTroops/pkg"
	"github.com/Evolt0/def-braveTroops/proto/epkg"
	"github.com/Evolt0/def-braveTroops/proto/prefix"
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
	err := pkg.InitTarget(stub)
	if err != nil {
		return shim.Error(epkg.WrapFailV2(err))
	}
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
	case prefix.FnListLedger:
		payload, err = transfer.List(stub, args)
	case prefix.FnPoW:
		payload, err = mining.PoW(stub, args)
	case prefix.FnGetTarget:
		payload, err = mining.GetTarget(stub, args)
	case prefix.FnListMining:
		payload, err = mining.List(stub, args)
	case prefix.FnCreateUser:
		payload, err = user.Create(stub, args)
	case prefix.FnListUser:
		payload, err = user.List(stub, args)
	case prefix.FnListLedgerByID:
		payload, err = user.ListLedgerByID(stub, args)
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
