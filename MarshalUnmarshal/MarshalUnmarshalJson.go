package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	marshalJson()
	unmarshalJson()
}

type person struct {
	First string
	Last  string
	Age   int
}


func marshalJson() {
	james := person{
		First: "James",
		Last:  "Bond",
		Age:   32,
	}
	missMoneypenny := person{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   25,
	}
	people := []person {james, missMoneypenny}
	fmt.Printf("People spies %v\n", people)

	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))
}

type newPerson struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

func unmarshalJson() {
	jsonResponse := `[{"First":"James","Last":"Bond","Age":32},{"First":"Miss","Last":"Moneypenny","Age":25}]`
	jsonBytes := []byte(jsonResponse)
	fmt.Printf("%T\n", jsonResponse)
	fmt.Printf("%T\n", jsonBytes)

	var newPeople []newPerson
	err := json.Unmarshal(jsonBytes, &newPeople)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Unmarshalled data %v\n", newPeople)
}

