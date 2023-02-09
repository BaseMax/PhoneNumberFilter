package main

import (
	"reflect"
	"testing"
)

func TestSortPhoneNumbers(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]string
		expected map[string]string
	}{
		{
			name:     "empty input",
			input:    map[string]string{},
			expected: map[string]string{},
		},
		{
			name: "one item",
			input: map[string]string{
				"12345678901": "John",
			},
			expected: map[string]string{
				"12345678901": "John",
			},
		},
		{
			name: "multiple items",
			input: map[string]string{
				"12345678901": "John",
				"98765432109": "Jane",
				"55555555555": "Jim",
			},
			expected: map[string]string{
				"55555555555": "Jim",
				"12345678901": "John",
				"98765432109": "Jane",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := sortPhoneNumbers(test.input)
			if !reflect.DeepEqual(actual, test.expected) {
				t.Errorf("expected: %v, actual: %v", test.expected, actual)
			}
		})
	}
}

func main() {
	testSortPhoneNumbers()
}
