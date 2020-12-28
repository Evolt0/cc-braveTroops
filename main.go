package main

import (
	"log"

	"github.com/Parker-Yang/cc-braveTroops/internal/contract"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func main() {
	if err := shim.Start(contract.New()); err != nil {
		log.Printf("failed to start contract: %v", err)
	}
}
