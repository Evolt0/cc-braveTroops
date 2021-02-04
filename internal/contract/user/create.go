package user

import (
	"github.com/Evolt0/cc-braveTroops/pkg"
	"github.com/Evolt0/def-braveTroops/proto"
	"github.com/Evolt0/def-braveTroops/proto/prefix"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func Create(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	req := &proto.UserReq{}
	err := pkg.Decode(args, req)
	if err != nil {
		return nil, err
	}
	err = execCreate(stub, req)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func execCreate(stub shim.ChaincodeStubInterface, req *proto.UserReq) error {
	user := newUser(stub, req)
	userKey, err := pkg.NewCompositeKey(stub, prefix.User, user.ID)
	if err != nil {
		return err
	}
	_, err = pkg.PutState(stub, userKey, user)
	if err != nil {
		return err
	}
	return nil
}
