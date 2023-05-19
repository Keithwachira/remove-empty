package empty

import (
	"encoding/json"
	"errors"
	"reflect"
)

func RemoveEmptyValuesFromStruct(data interface{}) (interface{}, error) {
	if reflect.ValueOf(data).Kind() != reflect.Struct {
		return nil, errors.New("the value you have passed is not a struct")
	}

	b, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	return RemoveEmptyValueFromJSON(string(b))
}

func RemoveEmptyValueFromJSON(jsonString string) (interface{}, error) {
	if !json.Valid([]byte(jsonString)) {
		return nil, errors.New("this is not a valid json string")
	}

	var data interface{}

	if err := json.Unmarshal([]byte(jsonString), &data); err != nil {
		return nil, err
	}

	return removeEmptyValues(data), nil
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

// isEmpty returns true if the passed value is the zero object.
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
