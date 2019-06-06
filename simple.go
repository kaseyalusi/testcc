package main

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type SimpleChaincode struct {
}


func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	initUUID := uuid.New()
	key := initUUID.String()
	err := stub.PutState(key, []byte("0"))
	if err != nil {
		shim.Error(err.Error())
	}
	return shim.Success([]byte(key))
}


func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	invokeUUID := uuid.New()
	key := invokeUUID.String()
	err := stub.PutState(key, []byte("0"))
	if err != nil {
		shim.Error(err.Error())
	}
	return shim.Success([]byte(key))
}

func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}