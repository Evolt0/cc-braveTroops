package transfer

import (
	"github.com/Parker-Yang/cc-braveTroops/pkg"
	"github.com/Parker-Yang/def-braveTroops/proto"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func Transfer(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	req := &proto.AmountsReq{}
	err := pkg.Decode(args, req)
	if err != nil {
		return nil, err
	}
	_, err = pkg.Check(stub, req, &req.BodyData)
	if err != nil {
		return nil, err
	}
	_, err = pkg.GetUser(stub, req.RID)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
