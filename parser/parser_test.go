package parser_test

import (
	"testing"

	"github.com/kr/pretty"
	"github.com/stretchr/testify/assert"

	"github.com/ecshreve/jcgo/parser"
)

func TestTransform(t *testing.T) {
	testcases := []struct {
		description string
		input       map[string]interface{}
		expected    parser.Object
	}{
		{
			description: "basic map",
			input: map[string]interface{}{
				"key1": "val1",
				"key2": "val2",
				"key3": "val3",
			},
			expected: &parser.MapObj{
				Prefix: "",
				Val: []parser.Object{
					&parser.StringObj{
						Prefix: "key1",
						Val:    "val1",
					},
					&parser.StringObj{
						Prefix: "key2",
						Val:    "val2",
					},
					&parser.StringObj{
						Prefix: "key3",
						Val:    "val3",
					},
				},
			},
		},
		{
			description: "basic nested map",
			input: map[string]interface{}{
				"outer1": map[string]interface{}{
					"inner1": "innerval1",
					"inner2": "innerval2",
				},
				"outer2": "outerval2",
			},
			expected: &parser.MapObj{
				Prefix: "",
				Val: []parser.Object{
					&parser.MapObj{
						Prefix: "outer1",
						Val: []parser.Object{
							&parser.StringObj{
								Prefix: "outer1_inner1",
								Val:    "innerval1",
							},
							&parser.StringObj{
								Prefix: "outer1_inner2",
								Val:    "innerval2",
							},
						},
					},
					&parser.StringObj{
						Prefix: "outer2",
						Val:    "outerval2",
					},
				},
			},
		},
		{
			description: "basic double nested map",
			input: map[string]interface{}{
				"outer1": map[string]interface{}{
					"inner1": "innerval1",
					"nestedmap": map[string]interface{}{
						"nested1": "nestedval1",
						"nested2": "nestedval2",
					},
				},
				"outer2": "outerval2",
			},
			expected: &parser.MapObj{
				Prefix: "",
				Val: []parser.Object{
					&parser.MapObj{
						Prefix: "outer1",
						Val: []parser.Object{
							&parser.StringObj{
								Prefix: "outer1_inner1",
								Val:    "innerval1",
							},
							&parser.MapObj{
								Prefix: "outer1_nestedmap",
								Val: []parser.Object{
									&parser.StringObj{
										Prefix: "outer1_nestedmap_nested1",
										Val:    "nestedval1",
									},
									&parser.StringObj{
										Prefix: "outer1_nestedmap_nested2",
										Val:    "nestedval2",
									},
								},
							},
						},
					},
					&parser.StringObj{
						Prefix: "outer2",
						Val:    "outerval2",
					},
				},
			},
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.description, func(t *testing.T) {
			actual := parser.Transform(testcase.input)
			assert.Equal(t, testcase.expected, actual)
			pretty.Print(actual)
		})
	}
}
