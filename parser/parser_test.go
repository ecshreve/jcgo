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

func TestParse(t *testing.T) {
	testcases := []struct {
		description string
		input       parser.Object
		expected    [][]string
	}{
		{
			description: "simple",
			input: &parser.StringObj{
				Prefix: "key1",
				Val:    "val1",
			},
			expected: [][]string{
				[]string{"key1"},
				[]string{"val1"},
			},
		},
		{
			description: "simple map",
			input: &parser.MapObj{
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
			expected: [][]string{
				[]string{"key1", "key2", "key3"},
				[]string{"val1", "val2", "val3"},
			},
		},
		{
			description: "simple nested",
			input: &parser.MapObj{
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
			expected: [][]string{
				[]string{"outer1_inner1", "outer1_inner2", "outer2"},
				[]string{"innerval1", "innerval2", "outerval2"},
			},
		},
		{
			description: "double nested",
			input: &parser.MapObj{
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
			expected: [][]string{
				[]string{"outer1_inner1", "outer1_nestedmap_nested1", "outer1_nestedmap_nested2", "outer2"},
				[]string{"innerval1", "nestedval1", "nestedval2", "outerval2"},
			},
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.description, func(t *testing.T) {
			actual := testcase.input.Parse()
			assert.Equal(t, testcase.expected, actual)
			pretty.Print(actual)
		})
	}
}
