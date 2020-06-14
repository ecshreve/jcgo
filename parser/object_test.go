package parser_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ecshreve/jcgo/parser"
)

func TestObjectCreation(t *testing.T) {
	testcases := []struct {
		description string
		input       interface{}
		expectError bool
	}{
		{
			description: "create StringObj",
			input:       "this is a string",
			expectError: false,
		},
		{
			description: "create BoolObj",
			input:       true,
			expectError: false,
		},
		{
			description: "create FloatObj",
			input:       float64(10.0),
			expectError: false,
		},
		{
			description: "invalid scalar input",
			input:       int64(1),
			expectError: true,
		},
		{
			description: "create MapObj",
			input:       map[string]interface{}{"key": "value"},
			expectError: false,
		},
		{
			description: "invalid type in map",
			input:       map[string]interface{}{"key": int64(1)},
			expectError: true,
		},
		{
			description: "create SliceObj",
			input:       []interface{}{"str1", "str2"},
			expectError: false,
		},
		{
			description: "invalid type in slice",
			input:       []interface{}{"str1", int64(1)},
			expectError: true,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.description, func(t *testing.T) {
			_, err := parser.ObjectFromInterface("", testcase.input)
			assert.Equal(t, testcase.expectError, err != nil)
		})
	}
}
