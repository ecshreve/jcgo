package object_test

import (
	"testing"

	"github.com/samsarahq/go/snapshotter"
	"github.com/stretchr/testify/assert"

	oo "github.com/ecshreve/jcgo/pkg/object"
)

func TestObjectCreation(t *testing.T) {
	snap := snapshotter.New(t)
	defer snap.Verify()

	testcases := []struct {
		description string
		input       interface{}
		expectError bool
	}{
		{
			description: "create nil StringObj",
			input:       nil,
			expectError: false,
		},
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
			description: "create NumberObj",
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
	}

	for _, testcase := range testcases {
		t.Run(testcase.description, func(t *testing.T) {
			obj, err := oo.FromInterface("", testcase.input)
			assert.Equal(t, testcase.expectError, err != nil)
			snap.Snapshot(testcase.description, obj)
		})
	}
}

func TestParseObject(t *testing.T) {
	testcases := []struct {
		description string
		input       oo.Object
		expected    [][]string
	}{
		{
			description: "simple string",
			input:       oo.NewStringObj("pref1", "stringVal"),
			expected: [][]string{
				{"pref1"},
				{"stringVal"},
			},
		},
		{
			description: "simple bool",
			input:       oo.NewBoolObj("pref1", true),
			expected: [][]string{
				{"pref1"},
				{"true"},
			},
		},
		{
			description: "simple float number",
			input:       oo.NewNumberObj("pref1", 5.5),
			expected: [][]string{
				{"pref1"},
				{"5.5"},
			},
		},
		{
			description: "simple integer number",
			input:       oo.NewNumberObj("pref1", 5.0),
			expected: [][]string{
				{"pref1"},
				{"5"},
			},
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.description, func(t *testing.T) {
			actual, _ := testcase.input.Parse()
			assert.Equal(t, testcase.expected, actual)
		})
	}
}
