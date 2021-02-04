package user

import (
	"github.com/Evolt0/cc-braveTroops/pkg"
	"github.com/Evolt0/def-braveTroops/proto"
	"github.com/Evolt0/def-braveTroops/proto/prefix"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func newUser(stub shim.ChaincodeStubInterface, req *proto.UserReq) *proto.User {
	return &proto.User{
		ID:         req.ID,
		ObjectType: prefix.User,
		PubKey:     req.PubKey,
		Balance:    0,
		CreateAt:   pkg.GetTimestamp(stub),
	}
}
