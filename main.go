package main

import (
	"encoding/json"
	"fmt"
	"reflect"
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

	result := removeEmptyValues(data)

	resultJSON, err := json.Marshal(result)
	if err != nil {
		fmt.Println("Failed to marshal JSON:", err)
		return
	}

	fmt.Println(string(resultJSON))
}

func removeEmptyValues(data interface{}) interface{} {
	switch value := data.(type) {
	case map[string]interface{}:
		for key, v := range value {
			if isEmpty(v) {
				delete(value, key)
			} else {
				value[key] = removeEmptyValues(v)
			}
		}
		return value
	case []interface{}:
		var result []interface{}
		for _, v := range value {
			if !isEmpty(v) {
				result = append(result, removeEmptyValues(v))
			}
		}
		return result
	default:
		return value
	}
}

// isEmpty returns true if the passed value is the zero object
func isEmpty(value interface{}) bool {
	if value == nil {
		return true
	}
	v := reflect.ValueOf(value)
	switch v.Kind() {
	case reflect.Slice, reflect.Map:
		return v.Len() == 0
	}

	// compare the types directly with more general coverage
	return reflect.DeepEqual(v.Interface(), reflect.Zero(v.Type()).Interface())
}
