package main

import (
	"fmt"
	"encoding/json"
)

type Msg struct{
	Status 	bool	`json:"Status"`
	Code 	int		`json:"Code"`
	Message string	`json:"Message"`
}

type Currency struct{
	TokenName 		string	`json:"TokenName"`
	TokenSymbol 	string	`json:"TokenSymbol"`
	TotalSupply 	float32	`json:"TotalSupply"`
}

type Token struct {
	Lock		bool	`json:"Lock"`
	Currency	map[string]Currency	`json:"Currency"`
}

func (token *Token) transfer (_from *Account, _to *Account, _currency string, _value float32) []byte{

	var rev []byte
	if (token.Lock){
		msg := &Msg{Status: false, Code: 0, Message: "锁仓状态，停止一切转账活动"}
		rev, _ = json.Marshal(msg)
		return rev
	}
	if(_from.Frozen ) {
		msg := &Msg{Status: false, Code: 0, Message: "From 账号冻结"}
		rev, _ = json.Marshal(msg)
		return rev
	}
	if( _to.Frozen) {
		msg := &Msg{Status: false, Code: 0, Message: "To 账号冻结"}
		rev, _ = json.Marshal(msg)
		return rev
	}
	if(!token.isCurrency(_currency)){
		msg := &Msg{Status: false, Code: 0, Message: "货币符号不存在"}
		rev, _ = json.Marshal(msg)
		return rev
	}
	if(_from.BalanceOf[_currency] >= _value){
		_from.BalanceOf[_currency] -= _value;
		_to.BalanceOf[_currency] += _value;

		msg := &Msg{Status: true, Code: 0, Message: "转账成功"}
		rev, _ = json.Marshal(msg)
		return rev
	}else{
		msg := &Msg{Status: false, Code: 0, Message: "余额不足"}
		rev, _ = json.Marshal(msg)
		return rev
	}
	
}
func (token *Token) initialSupply(_name string, _symbol string, _supply float32, _account *Account) []byte{

	token.Currency[_symbol] = Currency{TokenName: _name, TokenSymbol: _symbol, TotalSupply: _supply};
	if _account.BalanceOf[_symbol] > 0 {
		msg := &Msg{Status: false, Code: 0, Message: "账号中存在代币"}
		rev, _ := json.Marshal(msg)
		return rev
	}else{
		_account.BalanceOf[_symbol] = _supply
		msg := &Msg{Status: true, Code: 0, Message: "代币初始化成功"}
		rev, _ := json.Marshal(msg)
		return rev
	}
	
}

func (token *Token) mint(_currency string, _amount float32, _account *Account) []byte{
	if(!token.isCurrency(_currency)){
		msg := &Msg{Status: false, Code: 0, Message: "货币符号不存在"}
		rev, _ := json.Marshal(msg)
		return rev
	}
	cur := token.Currency[_currency]
	cur.TotalSupply += _amount;
	token.Currency[_currency] = cur
	_account.BalanceOf[_currency] = _amount;

	msg := &Msg{Status: true, Code: 0, Message: "代币增发成功"}
	rev, _ := json.Marshal(msg)
	return rev
	
}
func (token *Token) burn(_currency string, _amount float32, _account *Account) []byte{
	if(!token.isCurrency(_currency)){
		msg := &Msg{Status: false, Code: 0, Message: "货币符号不存在"}
		rev, _ := json.Marshal(msg)
		return rev
	}
	if(token.Currency[_currency].TotalSupply >= _amount){
		cur := token.Currency[_currency]
		cur.TotalSupply -= _amount;
		token.Currency[_currency] = cur
		_account.BalanceOf[_currency] -= _amount;

		msg := &Msg{Status: false, Code: 0, Message: "代币回收成功"}
		rev, _ := json.Marshal(msg)
		return rev
	}else{
		msg := &Msg{Status: false, Code: 0, Message: "代币回收失败，回收额度不足"}
		rev, _ := json.Marshal(msg)
		return rev
	}
	
}
func (token *Token) isCurrency(_currency string) bool {
	if _, ok := token.Currency[_currency]; ok {
		return true
	}else{
		return false
	}
}
func (token *Token) setLock(_look bool) bool {
	token.Lock = _look
	return token.Lock
}
type Account struct {
	Name			string	`json:"Name"`
	Frozen			bool	`json:"Frozen"`
	BalanceOf		map[string]float32	`json:"BalanceOf"`
}
func (account *Account) balance (_currency string) map[string]float32{
	bal	:= map[string]float32{_currency:account.BalanceOf[_currency]}
	return bal
}

func (account *Account) balanceAll() map[string]float32{
	return account.BalanceOf
}

func main(){

	// token := &token.Token{Currency: map[string]Currency{}}
	token()
	transfer()
	frozen()
	balance()
	lock()
}

func token(){
	token := &Token{Currency: map[string]Currency{}}

	coinbase := &Account{
		Name: "Coinbase",
		Frozen: false,
		BalanceOf: map[string]float32{}}

	token.initialSupply("水果币","Apple",10000, coinbase)
	token.initialSupply("积分币","PPC",10000, coinbase)

	token1, _ := json.Marshal(token)
	fmt.Println(string(token1))

	token.mint("Apple", 10000, coinbase) 

	token2, _ := json.Marshal(token)
	fmt.Println(string(token2))

	token.burn("Apple", 500, coinbase)

	tokenJson, _ := json.Marshal(token)
	fmt.Println(string(tokenJson))

	// fmt.Println(strconv.Itoa(int(coinbase.balance("Apple"))))
}

func balance(){
	fmt.Println("balance -----")
	account := &Account{
		Name: "Tom",
		Frozen: false,
		BalanceOf: map[string]float32{"RMB":1000, "USD":100, "CNY": 5000}}

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
		BalanceOf: map[string]float32{}}

	account2 := &Account{
		Name: "Tom",
		Frozen: false,
		BalanceOf: map[string]float32{"RMB":1000}}

	from, _ := json.Marshal(account1)
	fmt.Println(string(from))

	to, _ := json.Marshal(account2)
	fmt.Println(string(to))

	token := &Token{Currency: map[string]Currency{}}
	msg := token.initialSupply("积分币","RMB",10000, account1)
	fmt.Println(string(msg))
	rev := token.transfer(account1,account2,"RMB", 500.99)
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
		BalanceOf: map[string]float32{"RMB":1000}}

	account2 := &Account{
		Name: "Tom",
		Frozen: false,
		BalanceOf: map[string]float32{"RMB":1000}}

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

func lock(){
	fmt.Println("Lock -----")
	account1 := &Account{
		Name: "Jerry",
		Frozen: false,
		BalanceOf: map[string]float32{"RMB":1000.05}}

	account2 := &Account{
		Name: "Tom",
		Frozen: false,
		BalanceOf: map[string]float32{"RMB":1000.009}}

	from, _ := json.Marshal(account1)
	fmt.Println(string(from))

	to, _ := json.Marshal(account2)
	fmt.Println(string(to))

	

	token := &Token{Currency: map[string]Currency{}}
	msg := token.initialSupply("积分币","RMB",10000, account1)

	rev1 := token.transfer(account1,account2,"RMB", 2000)
	fmt.Println(string(rev1))

	token.setLock(true)
	fmt.Println(string(msg))
	rev := token.transfer(account1,account2,"RMB", 500.05)
	fmt.Println(string(rev))

	from1, _ := json.Marshal(account1)
	fmt.Println(string(from1))

	to1, _ := json.Marshal(account2)
	fmt.Println(string(to1))

	token1, _ := json.Marshal(token)
	fmt.Println(string(token1))
}