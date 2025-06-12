package main

import (
    "fmt"
    "github.com/hyperledger/fabric-chaincode-go/shim"
    "github.com/hyperledger/fabric-protos-go/peer"
)

// SimpleChaincode структура для chaincode
type SimpleChaincode struct{}

// Init вызывается при инициализации chaincode
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) peer.Response {
    fmt.Println("Chaincode initialized")
    return shim.Success(nil)
}

// Invoke точка входа для транзакций
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
    function, args := stub.GetFunctionAndParameters()
    
    switch function {
    case "Set":
        return t.Set(stub, args)
    case "Get":
        return t.Get(stub, args)
    default:
        return shim.Error("Invalid function name")
    }
}

// Set сохраняет пару ключ-значение в блокчейне
func (t *SimpleChaincode) Set(stub shim.ChaincodeStubInterface, args []string) peer.Response {
    if len(args) != 2 {
        return shim.Error("Requires 2 arguments: key and value")
    }
    
    key := args[0]
    value := args[1]
    
    err := stub.PutState(key, []byte(value))
    if err != nil {
        return shim.Error(fmt.Sprintf("Failed to set value: %s", err))
    }
    
    return shim.Success(nil)
}

// Get возвращает значение по ключу
func (t *SimpleChaincode) Get(stub shim.ChaincodeStubInterface, args []string) peer.Response {
    if len(args) != 1 {
        return shim.Error("Requires 1 argument: key")
    }
    
    key := args[0]
    value, err := stub.GetState(key)
    if err != nil {
        return shim.Error(fmt.Sprintf("Failed to get value: %s", err))
    }
    if value == nil {
        return shim.Error("Key not found")
    }
    
    return shim.Success(value)
}

// Точка входа
func main() {
    err := shim.Start(new(SimpleChaincode))
    if err != nil {
        fmt.Printf("Error starting chaincode: %s", err)
    }
}