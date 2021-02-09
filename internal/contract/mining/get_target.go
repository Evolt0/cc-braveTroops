package mining

import (
	"github.com/Evolt0/cc-braveTroops/pkg"
	"github.com/Evolt0/def-braveTroops/consts/status"
	"github.com/Evolt0/def-braveTroops/proto/epkg"
	"github.com/Evolt0/def-braveTroops/proto/prefix"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func GetTarget(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	key, err := pkg.NewCompositeKey(stub, prefix.Target, "")
	if err != nil {
		return nil, err
	}
	state, err := stub.GetState(key)
	if err != nil {
		return nil, epkg.Wrapf(status.InternalServerError, "failed to put: %v", err)
	}
	return state, nil
}
