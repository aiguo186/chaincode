package main

import "fmt"
import "encoding/json"

type Gasoline struct {
	Number 			string	`json:"Number"`
	Balance			float64	`json:"Balance"`
	Status 			string	`json:"Status"`
	Message 		map[string]string	`json:"Message"`
}

func (gasoline *Gasoline) initial(_number string, _balance float64,_msg string){
	gasoline.Number 	= _number
	gasoline.Balance 	= _balance
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

func main(){
	
	gasoline := &Gasoline{}
	
	gasoline.initial("1000000",100.00,"XX员工初始化了这张卡")
	gasoline.activate("小张激活了这张卡")
	gasoline.recharge("陈某某充值了这张卡")
	gasoline.discard("已充值废弃这张卡")
	gasolineJson, _ := json.Marshal(gasoline)
	fmt.Println(string(gasolineJson))

	fmt.Println("-----")

	gasoline.initial("2000000",150.00,"小赵初始化了这张卡")
	gasoline.recharge("陈某某充值了这张卡")
	gasoline.discard("废弃这张卡")

	gasolineJson, _ = json.Marshal(gasoline)
	fmt.Println(string(gasolineJson))

	fmt.Println("-----")

	gasoline.initial("3000000",150.00,"小赵初始化了这张卡")
	
	gasoline.discard("废弃这张卡")
	gasoline.recharge("陈某某充值了这张卡")

	gasolineJson, _ = json.Marshal(gasoline)
	fmt.Println(string(gasolineJson))

}