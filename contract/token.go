package main

import "fmt"
import "encoding/json"

type Token struct {
	Owner			string	`json:"Owner"`
	TotalSupply 	uint	`json:"TotalSupply"`
	TokenName 		string	`json:"TokenName"`
	TokenSymbol 	string	`json:"TokenSymbol"`
	BalanceOf		map[string]uint	`json:"BalanceOf"`
}

func (token *Token) initialSupply(){
	token.BalanceOf[token.Owner] = token.TotalSupply;
}

func (token *Token) transfer (_from string, _to string, _value uint){
	if(token.BalanceOf[_from] >= _value){
		token.BalanceOf[_from] -= _value;
		token.BalanceOf[_to] += _value;
	}
}

func (token *Token) balance (_from string) uint{
	return token.BalanceOf[_from]
}

func (token *Token) burn(_value uint) {
	if(token.BalanceOf[token.Owner] >= _value){
		token.BalanceOf[token.Owner] -= _value;
		token.TotalSupply -= _value;
	}
}

func (token *Token) burnFrom(_from string, _value uint) {
	if(token.BalanceOf[_from] >= _value){
		token.BalanceOf[_from] -= _value;
		token.TotalSupply -= _value;
	}
}

func (token *Token) mint(_value uint) {
	
	token.BalanceOf[token.Owner] += _value;
	token.TotalSupply += _value;
	
}

func main(){
	
	token := &Token{
		Owner: "netkiller",
		TotalSupply: 10000,
		TokenName: "代币通正",
		TokenSymbol: "COIN",
		BalanceOf: map[string]uint{}}
	
	token.initialSupply()
	
	fmt.Println(token.balance("netkiller"))

	token.transfer("netkiller","neo", 100)

	fmt.Println(token.balance("netkiller"))
	fmt.Println(token.balance("neo"))

	tokenJson, _ := json.Marshal(token)
	fmt.Println(string(tokenJson))

}