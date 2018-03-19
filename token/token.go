package main 

/* 
--------------------------------------------------
Author: netkiller <netkiller@msn.com>
Home: http://www.netkiller.cn
Data: 2018-03-19
--------------------------------------------------
*/

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)



type Token struct {
	Owner		string	`json:"Owner"`
	TotalSupply int		`json:"TotalSupply"`
	TokenName 	string	`json:"TokenName"`
	TokenSymbol string	`json:"TokenSymbol"`
	BalanceOf	map[string]int	`json:"BalanceOf"`
	FrozenAccount	map[string]int	`json:"FrozenAccount"`
	Lock		bool	`json:"Lock"`
}

func (token *Token) initialSupply(){
	if(token.TotalSupply == 0){
		token.BalanceOf[token.Owner] = token.TotalSupply;
	}
}

func (token *Token) transfer (_from string, _to string, _value int){
	if(token) return
	if(token.BalanceOf[_from] >= _value){
		token.BalanceOf[_from] -= _value;
		token.BalanceOf[_to] += _value;
	}
}

func (token *Token) balance (_from string) int{
	return token.BalanceOf[_from]
}

func (token *Token) burn(_value int) {
	if(token) return
	if(token.BalanceOf[token.Owner] >= _value){
		token.BalanceOf[token.Owner] -= _value;
		token.TotalSupply -= _value;
	}
}

func (token *Token) burnFrom(_from string, _value int) {
	if(token) return
	if(token.BalanceOf[_from] >= _value){
		token.BalanceOf[_from] -= _value;
		token.TotalSupply -= _value;
	}
}

func (token *Token) mint(_value int) {
	if(token) return
	token.BalanceOf[token.Owner] += _value;
	token.TotalSupply += _value;
	
}

func (token *Token) setLock(_lock bool) {
	token.Lock = _look;	
}

func (token *Token) frozen(_account string, _status bool) {
	token.FrozenAccount[_account] = _status;	
}

type Account struct {
	Owner		string	`json:"Owner"`
	TokenName 	string	`json:"TokenName"`
	TokenSymbol string	`json:"TokenSymbol"`
	Balance		int		`json:"BalanceOf"`
	Frozen		bool	`json:"FrozenAccount"`
}

// Define the Smart Contract structure
type SmartContract struct {
}

func (s *SmartContract) Init(stub shim.ChaincodeStubInterface) pb.Response {

	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	symbol:= args[0]
	name  := args[1]
	supply,_:= strconv.Atoi(args[2])

	token := &Token{
		Owner: "coinbase",
		TotalSupply: supply,
		TokenName: name,
		TokenSymbol: symbol,
		BalanceOf: map[string]int{},
		FrozenAccount: map[string]int{},
		Lock: false}

	token.initialSupply()

	tokenAsBytes, _ := json.Marshal(token)
	err := stub.PutState(symbol, tokenAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}
	fmt.Printf("Init %s \n", string(tokenAsBytes))

	return shim.Success(nil)
}

func (s *SmartContract) transferToken(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. Expecting 4")
	}
	_from 	:= args[1]
	_to	:= args[2]
	_amount,_	:= strconv.Atoi(args[3])
	if(_amount <= 0){
		return shim.Error("Incorrect number of amount")
	}

	tokenAsBytes,err := stub.GetState(args[0])
	if err != nil {
		return shim.Error(err.Error())
	}
	fmt.Printf("transferToken - begin %s \n", string(tokenAsBytes))

	token := Token{}

	json.Unmarshal(tokenAsBytes, &token)
	token.transfer(_from, _to, _amount)

	tokenAsBytes, err = json.Marshal(token)
	if err != nil {
		return shim.Error(err.Error())
	}
	err = stub.PutState(args[0], tokenAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}
	fmt.Printf("transferToken - end %s \n", string(tokenAsBytes))

	return shim.Success(nil)
}
func (s *SmartContract) mintToken(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	
	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	tokenAsBytes,err := stub.GetState(args[0])
	if err != nil {
		return shim.Error(err.Error())
	}
	// fmt.Printf("setLock - begin %s \n", string(tokenAsBytes))

	token := Token{}

	json.Unmarshal(tokenAsBytes, &token)
	_amount,_	:= strconv.Atoi(args[1])
	token.mint(_amount)

	tokenAsBytes, err = json.Marshal(token)
	if err != nil {
		return shim.Error(err.Error())
	}
	err = stub.PutState(args[0], tokenAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}
	fmt.Printf("mintToken - end %s \n", string(tokenAsBytes))

	return shim.Success(nil)
}

func (s *SmartContract) setLock(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	
	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	tokenAsBytes,err := stub.GetState(args[0])
	if err != nil {
		return shim.Error(err.Error())
	}
	// fmt.Printf("setLock - begin %s \n", string(tokenAsBytes))

	token := Token{}

	json.Unmarshal(tokenAsBytes, &token)
	token.setLock(args[1])

	tokenAsBytes, err = json.Marshal(token)
	if err != nil {
		return shim.Error(err.Error())
	}
	err = stub.PutState(args[0], tokenAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}
	fmt.Printf("setLock - end %s \n", string(tokenAsBytes))

	return shim.Success(nil)
}
func (s *SmartContract) frozenAccount(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	
	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	tokenAsBytes,err := stub.GetState(args[0])
	if err != nil {
		return shim.Error(err.Error())
	}
	// fmt.Printf("setLock - begin %s \n", string(tokenAsBytes))

	token := Token{}

	json.Unmarshal(tokenAsBytes, &token)

	var status bool
	if(args[1] == "true"){
		status = true;
	}else{
		status = false
	}

	token.frozen(args[1],status)

	tokenAsBytes, err = json.Marshal(token)
	if err != nil {
		return shim.Error(err.Error())
	}
	err = stub.PutState(args[0], tokenAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}
	fmt.Printf("frozenAccount - end %s \n", string(tokenAsBytes))

	return shim.Success(nil)
}

func (s *SmartContract) balanceToken(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	tokenAsBytes,err := stub.GetState(args[0])
	if err != nil {
		return shim.Error(err.Error())
	}
	token := Token{}
	json.Unmarshal(tokenAsBytes, &token)
	amount := token.balance(args[1])
	status := token.FrozenAccount[args[1]]

	account := Account{
		Owner: token.Owner,
		TokenName: token.TokenName,
		TokenSymbol: token.TokenSymbol
		Balance: amount,
		Frozen: status}

	// value := strconv.Itoa(amount)
	
	tokenAsBytes, _ := json.Marshal(account)
	fmt.Printf("%s balance is %s \n", args[1], value)	

	return shim.Success(tokenAsBytes)
}
func (s *SmartContract) Query(stub shim.ChaincodeStubInterface) pb.Response {
	function, args := stub.GetFunctionAndParameters()
	if function == "balanceToken" {
		return s.balanceToken(stub, args)
	} else {

	}
}

func (s *SmartContract) Invoke(stub shim.ChaincodeStubInterface) pb.Response {

	// Retrieve the requested Smart Contract function and arguments
	function, args := stub.GetFunctionAndParameters()
	// Route to the appropriate handler function to interact with the ledger appropriately
	if function == "setLock" {
		return s.setLock(stub, args)
	} else if function == "transferToken" {
		return s.transferToken(stub, args)
	} else if function == "frozenAccount" {
		return s.frozenAccount(stub, args)
	} else if function == "mintToken" {
		return s.mintToken(stub, args)
	}

	return shim.Error("Invalid Smart Contract function name.")
}

// The main function is only relevant in unit test mode. Only included here for completeness.
func main() {

	// Create a new Smart Contract
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}
