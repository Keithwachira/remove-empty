package main

import (
	"encoding/json"
	"fmt"
	"log"

	empty "github.com/Keithwachira/remove-empty"
)

func main() {
	// Example JSON data
	/*jsonData := `{
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
		}`*/

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

	/*var data interface{}
	if err := json.Unmarshal([]byte(jsonData), &data); err != nil {
		fmt.Println("Failed to unmarshal JSON:", err)
		return
	}*/

	fromStruct, err := empty.RemoveEmptyValuesFromStruct(data)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(fromStruct)

	/*result, err := empty.RemoveEmptyValueFromJson(jsonData)
	if err != nil {
		fmt.Println("remove empty json failed:", err)
		return
	}*/
	resultJSON, err := json.Marshal(fromStruct)
	if err != nil {
		fmt.Println("Failed to marshal JSON:", err)
		return
	}

	fmt.Println(string(resultJSON))
}
