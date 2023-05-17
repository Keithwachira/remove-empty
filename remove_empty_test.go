package empty

import (
	"testing"
)

func TestIsZero(t *testing.T) {
	tests := []struct {
		value    interface{}
		expected bool
		name     string
	}{
		{value: nil, expected: true, name: "Test nil value"},
		{value: 0.0, expected: true, name: "Test nil float"},
		{value: 0.1, expected: false, name: "Test nil float"},
		{value: -0.1, expected: false, name: "Test negative int"},

		{value: 0, expected: true, name: "Test integer value"},

		{value: false, expected: true, name: "Test false bool"},

		{value: []int{}, expected: true, name: "Test empty slice"},

		{value: map[string]string{}, expected: true, name: "Test empty map"},

		{value: "", expected: true, name: "Test empty string"},

		{value: 1, expected: false, name: "Test non zero integer"},

		{value: true, expected: false, name: "Test true bool"},

		{value: []int{1, 2, 3}, expected: false, name: "Test non empty slice"},

		{value: map[string]string{"key": "value"}, expected: false, name: "Test no nil map"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isEmpty(tt.value)
			if result != tt.expected {
				t.Errorf("isEmpty(%v) returned %v, expected %v", tt.value, result, tt.expected)
			}
		})
	}
}
