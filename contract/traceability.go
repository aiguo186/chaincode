package main

import "fmt"
import "encoding/json"

const (
	Origin = iota 	// 0
	Factory        	// 1
	QA       		// 2
	Shipping		// 3
	Received		// 4
	Pending			// 5
	Supermarket		// 6
)

type structElement struct {
	Name string `json:"name"`
	Company string `json:"company"`
	Description string `json:"description"`
}

type structLogistics struct {
    Stations string `json:"stations"`	// 中转站
	Date	string  `json:"date"`  // 转运日期
	Status	uint8	`json:"status"`  // 状态
    Message    string `json:"message"` // 留言信息
}

type Traceability struct {
	Name string	`json:"name"`
	Address string	`json:"address"`
	Attribute	map[string]string 	`json:"attribute"`
	Element		[]structElement		`json:"element"`
	Logistics	map[string]structLogistics	`json:"logistics"`
}

func (traceability *Traceability) setName(_name string) {
    traceability.Name = _name
}

func (traceability *Traceability) getName() string {
    return traceability.Name
}

func (traceability *Traceability) putAttribute(_key string, _value string) {
    traceability.Attribute[_key] = _value
}

func (traceability *Traceability) putLogistics(_key string, _value structLogistics) {
    traceability.Logistics[_key] = _value
}

func main(){
	
	traceability := &Traceability{
		Name: "牦牛肉干",
		Address: "内蒙古呼和浩特",
		Attribute: map[string]string{},
		Element: []structElement{structElement{Name:"塑料袋",Company: "XXX塑料制品有限公司", Description: "外包装"},structElement{Name:"辣椒粉",Company: "XXX调味品有限公司", Description: "采摘年份2016-10-10"},structElement{Name:"调和油",Company: "XXX调味品有限公司", Description: "生产日期2016-10-10"}},
		Logistics: map[string]structLogistics{}}
	
	traceability.putAttribute("Color","Red")
	traceability.putAttribute("Size","10")
	traceability.putAttribute("Weight","100kg")
	
	traceability.putLogistics("1", structLogistics{"呼和浩特","2016-10-15", Origin, "牦牛收购"})
	traceability.putLogistics("2", structLogistics{"呼和浩特","2016-10-18", Factory, "牦牛宰杀"})
	traceability.putLogistics("3", structLogistics{"呼和浩特","2016-10-15", QA, "经过质检"})
	traceability.putLogistics("4", structLogistics{"北京市","2016-10-15", Shipping, "运输中"})
	traceability.putLogistics("5", structLogistics{"杭州市","2016-10-15", Shipping, "XXX冷库"})
	traceability.putLogistics("5", structLogistics{"深圳市","2016-10-15", Supermarket, "XXX超市"})
	traceability.putLogistics("5", structLogistics{"龙华区","2016-10-15", Received, "用户签收"})
	

	traceabilityJson, _ := json.Marshal(traceability)
	fmt.Println(string(traceabilityJson))

}