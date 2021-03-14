package mining

import (
	"github.com/Evolt0/cc-braveTroops/pkg"
	"github.com/Evolt0/def-braveTroops/proto"
	"github.com/Evolt0/def-braveTroops/proto/prefix"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func newMining(stub shim.ChaincodeStubInterface, user *proto.User, data *proto.Target, req *proto.MiningReq) *proto.Mining {
	return &proto.Mining{
		ID:            pkg.NewUUID(stub, prefix.Mining),
		ObjectType:    prefix.Mining,
		UID:           user.ID,
		UName:         user.Name,
		ResultHash:    req.ResultHash,
		Result:        req.Result,
		CreateAt:      req.Timestamp,
		TargetPrefix:  data.Prefix,
		TargetNumZero: data.NumZero,
	}
}
