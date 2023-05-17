package empty

import (
	"reflect"
)

func RemoveEmptyValues(data interface{}) interface{} {
	switch value := data.(type) {
	case map[string]interface{}:
		for key, v := range value {
			if isEmpty(v) {
				delete(value, key)
			} else {
				value[key] = RemoveEmptyValues(v)
			}
		}
		return value
	case []interface{}:
		var result []interface{}
		for _, v := range value {
			if !isEmpty(v) {
				result = append(result, RemoveEmptyValues(v))
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
