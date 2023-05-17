package main

import (
	"encoding/json"
	"fmt"

	empty "github.com/Keithwachira/remove-empty"
)

func main() {
	// Example JSON data
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

	var data interface{}
	if err := json.Unmarshal([]byte(jsonData), &data); err != nil {
		fmt.Println("Failed to unmarshal JSON:", err)
		return
	}

	result := empty.RemoveEmptyValues(data)

	resultJSON, err := json.Marshal(result)
	if err != nil {
		fmt.Println("Failed to marshal JSON:", err)
		return
	}

	fmt.Println(string(resultJSON))
}
