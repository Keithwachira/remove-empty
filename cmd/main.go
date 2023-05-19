package main

import (
	"fmt"
	"log"

	empty "github.com/Keithwachira/remove-empty"
)

func main() {
	jsonData := `{
			"name": "iatchi",
			"age": 30,
	        "mmmh":"",
			"address": {
				"street": "123 Main St",
				"city": "",
				"zip": 0
			},
			"numbers": [1, 2, 0, 3, 0],
			"emptyArray": [],
			"emptyObject": {}
		}`

	type GodMode struct {
		Name  string
		Last  float64
		Mimi  int
		Keith string
	}

	data := GodMode{
		Name:  "mmmj",
		Last:  4.2,
		Mimi:  0,
		Keith: "",
	}

	fromStruct, err := empty.RemoveEmptyValuesFromStruct(data)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Printf("%+v\n", fromStruct)

	result, err := empty.RemoveEmptyValueFromJSON(jsonData)
	if err != nil {
		fmt.Println("remove empty json failed:", err)
		return
	}

	fmt.Printf("%+v\n", result)
}
