package transfer

import (
	"github.com/Evolt0/cc-braveTroops/pkg"
	"github.com/Evolt0/def-braveTroops/proto"
	"github.com/Evolt0/def-braveTroops/proto/prefix"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func newAmounts(stub shim.ChaincodeStubInterface, sender, receiver *proto.User, req *proto.AmountsReq) *proto.Amounts {
	return &proto.Amounts{
		ID:         pkg.NewUUID(stub, prefix.Amounts),
		ObjectType: prefix.Amounts,
		SID:        req.ID,
		SName:      sender.Name,
		RID:        req.RID,
		RName:      receiver.Name,
		Change:     req.Change,
		CreateAt:   req.Timestamp,
	}
}
