package mining

import (
	"github.com/Evolt0/cc-braveTroops/pkg"
	"github.com/Evolt0/def-braveTroops/proto"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func PoW(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	req := &proto.MiningReq{}
	err := pkg.Decode(args, req)
	if err != nil {
		return nil, err
	}
	user, err := pkg.Check(stub, req, &req.BodyData)
	if err != nil {
		return nil, err
	}
	err = execPoW(stub, user, req)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func execPoW(stub shim.ChaincodeStubInterface, user *proto.User, req *proto.MiningReq) error {
	return nil
}
