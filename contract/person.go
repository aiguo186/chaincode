package main

import "fmt"
import "encoding/json"

type Person struct {
	No string	`json:"no"`
	Name string	`json:"name"`
	Sex	bool	`json:"sex"`
	Age int		`json:"age"`
    Address string	`json:"address"`
}

func main(){
	
	person := &Person{"430725198001190911","景峯",true,30,"Shenzhen,China"}

	personJson, _ := json.Marshal(person)

	fmt.Println(string(personJson));

	person1 := &Person{
        No: "430725198001190911",
		Name: "Neo Chen",
		Sex: true,
		Age: 35,
		Address: "Shenzhen, China"}

    json2, _ := json.Marshal(person1)
    fmt.Println(string(json2))
}