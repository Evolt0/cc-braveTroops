package transfer

import (
	"github.com/Evolt0/cc-braveTroops/pkg"
	"github.com/Evolt0/def-braveTroops/consts/status"
	"github.com/Evolt0/def-braveTroops/proto"
	"github.com/Evolt0/def-braveTroops/proto/epkg"
	"github.com/Evolt0/def-braveTroops/proto/prefix"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func Transfer(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	req := &proto.AmountsReq{}
	err := pkg.Decode(args, req)
	if err != nil {
		return nil, err
	}
	user, err := pkg.Check(stub, req, &req.BodyData)
	if err != nil {
		return nil, err
	}
	receiver, err := pkg.GetUser(stub, req.RID)
	if err != nil {
		return nil, err
	}
	err = execTransfer(stub, user, receiver, req)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func execTransfer(stub shim.ChaincodeStubInterface, user, receiver *proto.User, req *proto.AmountsReq) error {
	if user.Balance < req.Change {
		return epkg.Wrapf(status.Conflict, "balance not enough")
	}
	user.Balance = user.Balance - req.Change
	receiver.Balance = receiver.Balance + req.Change
	userKey, err := pkg.NewCompositeKey(stub, prefix.User, user.ID)
	if err != nil {
		return err
	}
	_, err = pkg.UpdateState(stub, userKey, user)
	if err != nil {
		return err
	}
	receiverKey, err := pkg.NewCompositeKey(stub, prefix.User, receiver.ID)
	if err != nil {
		return err
	}
	_, err = pkg.UpdateState(stub, receiverKey, receiver)
	if err != nil {
		return err
	}
	amounts := newAmounts(stub, user, receiver, req)
	amountsKey, err := pkg.NewCompositeKey(stub, prefix.Amounts, amounts.ID)
	if err != nil {
		return err
	}
	_, err = pkg.PutState(stub, amountsKey, amounts)
	if err != nil {
		return err
	}
	return nil
}
