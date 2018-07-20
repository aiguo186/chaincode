package main

import (
	"fmt"
	"encoding/json"
	"github.com/netkiller/chaincode/contract/token"
)

func main(){

	// token := &token.Token{Currency: map[string]Currency{}}
	tokenTest()
	// transfer()
	// frozen()
	// balance()
}

func tokenTest(){
	tokenObj := &token.Token{Currency: map[string]token.Currency{}}

	coinbase := &token.Account{
		Name: "Coinbase",
		Frozen: false,
		BalanceOf: map[string]uint{}}

	tokenObj.initialSupply("水果币","Apple",10000, coinbase)
	tokenObj.initialSupply("积分币","PPC",10000, coinbase)

	token1, _ := json.Marshal(tokenObj)
	fmt.Println(string(token1))

	tokenObj.mint("Apple", 10000, coinbase) 

	token2, _ := json.Marshal(tokenObj)
	fmt.Println(string(token2))

	tokenObj.burn("Apple", 500, coinbase)

	tokenJson, _ := json.Marshal(tokenObj)
	fmt.Println(string(tokenJson))

	// fmt.Println(strconv.Itoa(int(coinbase.balance("Apple"))))
}
/*
func balance(){
	fmt.Println("balance -----")
	account := &Account{
		Name: "Tom",
		Frozen: false,
		BalanceOf: map[string]uint{"RMB":1000, "USD":100, "CNY": 5000}}

	value := account.balance("RMB")
	token1, _ := json.Marshal(value)
	fmt.Println(string(token1))

	bal := account.balanceAll()
	token2, _ := json.Marshal(bal)
	fmt.Println(string(token2))
}

func transfer(){
	fmt.Println("transfer -----")
	account1 := &Account{
		Name: "Neo",
		Frozen: false,
		BalanceOf: map[string]uint{}}

	account2 := &Account{
		Name: "Tom",
		Frozen: false,
		BalanceOf: map[string]uint{"RMB":1000}}

	from, _ := json.Marshal(account1)
	fmt.Println(string(from))

	to, _ := json.Marshal(account2)
	fmt.Println(string(to))

	token := &Token{Currency: map[string]Currency{}}
	msg := token.initialSupply("积分币","RMB",10000, account1)
	fmt.Println(string(msg))
	rev := token.transfer(account1,account2,"RMB", 500)
	fmt.Println(string(rev))

	from1, _ := json.Marshal(account1)
	fmt.Println(string(from1))

	to1, _ := json.Marshal(account2)
	fmt.Println(string(to1))

	token1, _ := json.Marshal(token)
	fmt.Println(string(token1))

}

func frozen(){
	fmt.Println("Frozen -----")
	account1 := &Account{
		Name: "Neo",
		Frozen: true,
		BalanceOf: map[string]uint{"RMB":1000}}

	account2 := &Account{
		Name: "Tom",
		Frozen: false,
		BalanceOf: map[string]uint{"RMB":1000}}

	from, _ := json.Marshal(account1)
	fmt.Println(string(from))

	to, _ := json.Marshal(account2)
	fmt.Println(string(to))

	token := &Token{Currency: map[string]Currency{}}
	msg := token.initialSupply("积分币","RMB",10000, account1)
	fmt.Println(string(msg))
	rev := token.transfer(account1,account2,"RMB", 500)
	fmt.Println(string(rev))

	from1, _ := json.Marshal(account1)
	fmt.Println(string(from1))

	to1, _ := json.Marshal(account2)
	fmt.Println(string(to1))

	token1, _ := json.Marshal(token)
	fmt.Println(string(token1))
}
*/