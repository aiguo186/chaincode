package main 

/* 
--------------------------------------------------
Author: netkiller <netkiller@msn.com>
Home: http://www.netkiller.cn
Data: 2018-03-20 11:00 PM
--------------------------------------------------

CORE_PEER_ADDRESS=peer:7051 CORE_CHAINCODE_ID_NAME=token3:1.0 chaincode/token/token3

peer chaincode install -n token3 -v 1.0 -p chaincodedev/chaincode/token
peer chaincode instantiate -C myc -n token3 -v 1.0 -c '{"Args":[""]}' -P "OR ('Org1MSP.member','Org2MSP.member')"


peer chaincode invoke -C myc -n token3 -c '{"function":"createAccount","Args":["coinbase"]}'
peer chaincode invoke -C myc -n token3 -v 1.0 -c '{"function":"showAccount","Args":["coinbase"]}'
peer chaincode invoke -C myc -n token3 -c '{"function":"balanceAll","Args":["coinbase"]}'

peer chaincode invoke -C myc -n token3 -c '{"function":"initCurrency","Args":["Netkiller Token","NKC","1000000","coinbase"]}'
peer chaincode invoke -C myc -n token3 -c '{"function":"initCurrency","Args":["NEO Token","NEC","1000000","coinbase"]}'

peer chaincode invoke -C myc -n token3 -c '{"function":"setLock","Args":["true"]}'
peer chaincode invoke -C myc -n token3 -c '{"function":"setLock","Args":["false"]}'

peer chaincode invoke -C myc -n token3 -c '{"function":"mintToken","Args":["NKC","5000","coinbase"]}'

peer chaincode invoke -C myc -n token3 -c '{"function":"createAccount","Args":["netkiller"]}'
peer chaincode invoke -C myc -n token3 -c '{"function":"transferToken","Args":["coinbase","netkiller","NKC","100"]}'		
peer chaincode invoke -C myc -n token3 -c '{"function":"balance","Args":["netkiller","NKC"]}'

peer chaincode invoke -C myc -n token3 -c '{"function":"frozenAccount","Args":["netkiller","true"]}'

--------------------------------------------------

	+----------+    +-----------+    +-----------+    
	| New      | -> | Activated | -> | Recharged |
	+----------+    +-----------+    +-----------+    
		 |                |
		 V                |
	+----------+          |
    | Discard  | <--------+
	+----------+
         |
         V
	+----------+ 
    | Delete   |
	+----------+ 
*/

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)
// ----------

type Gasoline struct {
	Number 			string	`json:"Number"`
	Amount			float64	`json:"Amount"`
	Status 			string	`json:"Status"`
	Message 		map[string]string	`json:"Message"`
}

func (gasoline *Gasoline) initial(_number string, _amount float64,_msg string){
	gasoline.Number 	= _number
	gasoline.Amount 	= _amount
	gasoline.Status		= "New"
	gasoline.Message	= map[string]string{gasoline.Status : _msg}
}

func (gasoline *Gasoline) activate (_msg string) bool{
	if(gasoline.Status == "New"){
		gasoline.Status = "Activated"
		gasoline.Message[gasoline.Status] = _msg
		return true
	}
	return false
}

func (gasoline *Gasoline) recharge (_msg string) bool{
	if(gasoline.Status == "Activated"){
		gasoline.Status = "Recharged"
		gasoline.Message[gasoline.Status] = _msg
		return true
	}
	return false
}

func (gasoline *Gasoline) discard (_msg string) bool{
	if gasoline.Status != "Recharged" {
		gasoline.Status = "Discard"
		gasoline.Message[gasoline.Status] = _msg
		return true
	}
	return false
}

func (gasoline *Gasoline) delete() bool{
	if gasoline.Status != "Discard" {
		gasoline.Status = "Deleted"
		return true
	}
	return false
}
// -----------


// Define the Smart Contract structure
type SmartContract struct {}

func (s *SmartContract) Init(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success(nil)
}

func (s *SmartContract) Query(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Error("Invalid Smart Contract function name.")
}

func (s *SmartContract) Invoke(stub shim.ChaincodeStubInterface) pb.Response {

	// Retrieve the requested Smart Contract function and arguments
	function, args := stub.GetFunctionAndParameters()
	// Route to the appropriate handler function to interact with the ledger appropriately
	if function == "initial" {
		return s.initialGasoline(stub, args)
	} else if function == "activate" {
		return s.activateGasoline(stub, args)
	} else if function == "recharge" {
		return s.rechargeGasoline(stub, args)
	} else if function == "discard" {
		return s.discardGasoline(stub, args)
	} else {
		return s.showGasoline(stub, args)
	}

	return shim.Error("Invalid Smart Contract function name.")
}

func (s *SmartContract) initialGasoline(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. Expecting 4")
	}

	_key  	:= args[0]
	_number := args[1]
	_amount,_:= strconv.ParseFloat(args[2], 64)
	_message:= args[3]

	if(_amount <= 0){
		return shim.Error("Incorrect number of amount")
	}

	existAsBytes,err := stub.GetState(_key)
	fmt.Printf("GetState(%s) %s \n", _key, string(existAsBytes))
	if string(existAsBytes) != "" {
		fmt.Println("Failed to create account, Duplicate key.")
		return shim.Error("Failed to create account, Duplicate key.")
	}

	gasoline := &Gasoline{}
	gasoline.initial(_number, _amount, _message)

	gasolineAsBytes, _ := json.Marshal(gasoline)
	err = stub.PutState(_key, gasolineAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}
	fmt.Printf("initialGasoline %s \n", string(gasolineAsBytes))

	return shim.Success(gasolineAsBytes)
}

func (s *SmartContract) showGasoline(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}
	_key 	:= args[0]

	gasolineAsBytes,err := stub.GetState(_key)
	if err != nil {
		return shim.Error(err.Error())
	}else{
		fmt.Printf("showGasoline %s \n", string(gasolineAsBytes))
	}
	return shim.Success(gasolineAsBytes)
}

func (s *SmartContract) activateGasoline(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}
	_key 	:= args[0]
	_msg 	:= args[1]

	gasolineAsBytes,err := stub.GetState(_key)
	if err != nil {
		return shim.Error(err.Error())
	}else{
		fmt.Printf("activateGasoline %s \n", string(gasolineAsBytes))
	}

	gasoline := Gasoline{}
	json.Unmarshal(gasolineAsBytes, &gasoline)
	gasoline.activate(_msg)
	
	gasolineAsBytes, err = json.Marshal(gasoline)
	err = stub.PutState(_key, gasolineAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}
	fmt.Printf("activateGasoline - end %s \n", string(gasolineAsBytes))
	
	return shim.Success(gasolineAsBytes)
}

func (s *SmartContract) rechargeGasoline(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}
	_key 	:= args[0]
	_msg 	:= args[1]

	gasolineAsBytes,err := stub.GetState(_key)
	if err != nil {
		return shim.Error(err.Error())
	}else{
		fmt.Printf("rechargeGasoline %s \n", string(gasolineAsBytes))
	}

	gasoline := Gasoline{}
	json.Unmarshal(gasolineAsBytes, &gasoline)
	gasoline.recharge(_msg)
	
	gasolineAsBytes, err = json.Marshal(gasoline)
	err = stub.PutState(_key, gasolineAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}
	fmt.Printf("rechargeGasoline - end %s \n", string(gasolineAsBytes))
	
	return shim.Success(gasolineAsBytes)
}

func (s *SmartContract) discardGasoline(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}
	_key 	:= args[0]
	_msg 	:= args[1]

	gasolineAsBytes,err := stub.GetState(_key)
	if err != nil {
		return shim.Error(err.Error())
	}else{
		fmt.Printf("discardGasoline %s \n", string(gasolineAsBytes))
	}

	gasoline := Gasoline{}
	json.Unmarshal(gasolineAsBytes, &gasoline)
	gasoline.discard(_msg)
	
	gasolineAsBytes, err = json.Marshal(gasoline)
	err = stub.PutState(_key, gasolineAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}
	fmt.Printf("discardGasoline - end %s \n", string(gasolineAsBytes))
	
	return shim.Success(gasolineAsBytes)
}

func (s *SmartContract) deleteGasoline(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}
	_key 	:= args[0]

	gasolineAsBytes,err := stub.GetState(_key)
	if err != nil {
		return shim.Error(err.Error())
	}else{
		fmt.Printf("discardGasoline %s \n", string(gasolineAsBytes))
	}

	gasoline := Gasoline{}
	json.Unmarshal(gasolineAsBytes, &gasoline)
	
	if(gasoline.Status == "Discard"){
		gasoline.delete()
		err = stub.DelState(_key)
		if err != nil {
			return shim.Error(err.Error())
		}
	}

	return shim.Success(gasolineAsBytes)
}

// The main function is only relevant in unit test mode. Only included here for completeness.
func main() {

	// Create a new Smart Contract
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}
