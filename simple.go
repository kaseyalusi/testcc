package main

import (
	"github.com/google/uuid"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

type SimpleChaincode struct {
}

var A, B string
var Aval, Bval, X int


func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	initUUID := uuid.New()
	key := initUUID.String()
	err := stub.PutState(key, []byte("0"))
	if err != nil {
		return nil, err
	}
	return []byte(key), nil
}


func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	invokeUUID := uuid.New()
	key := invokeUUID.String()
	err := stub.PutState(key, []byte("0"))
	if err != nil {
		return nil, err
	}
	return []byte(key), nil
}