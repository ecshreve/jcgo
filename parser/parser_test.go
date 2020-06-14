package parser_test

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/kr/pretty"
	"github.com/stretchr/testify/assert"

	"github.com/ecshreve/jcgo/parser"
)

func TestTransform(t *testing.T) {
	spew.Config.DisablePointerAddresses = true
	spew.Config.DisableCapacities = true
	spew.Config.Indent = "  "

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
		{
			description: "map with basic slice",
			input: map[string]interface{}{
				"data": []interface{}{
					map[string]interface{}{
						"key1": "val1",
						"key2": "val2",
						"key3": "val3",
					},
					map[string]interface{}{
						"key1": "val4",
						"key2": "val5",
						"key3": "val6",
					},
				},
			},
			expected: &parser.MapObj{
				Prefix: "",
				Val: []parser.Object{
					&parser.SliceObj{
						Prefix: "data",
						Val: []parser.Object{
							&parser.MapObj{
								Prefix: "data",
								Val: []parser.Object{
									&parser.StringObj{
										Prefix: "data_key1",
										Val:    "val1",
									},
									&parser.StringObj{
										Prefix: "data_key2",
										Val:    "val2",
									},
									&parser.StringObj{
										Prefix: "data_key3",
										Val:    "val3",
									},
								},
							},
							&parser.MapObj{
								Prefix: "data",
								Val: []parser.Object{
									&parser.StringObj{
										Prefix: "data_key1",
										Val:    "val4",
									},
									&parser.StringObj{
										Prefix: "data_key2",
										Val:    "val5",
									},
									&parser.StringObj{
										Prefix: "data_key3",
										Val:    "val6",
									},
								},
							},
						},
					},
				},
			},
		},
		{
			description: "map with slice with nested map",
			input: map[string]interface{}{
				"data": []interface{}{
					map[string]interface{}{
						"key1": "val1",
						"nestedmap": map[string]interface{}{
							"nested1": "nestedval1",
							"nested2": "nestedval2",
						},
					},
					map[string]interface{}{
						"key1": "val4",
						"nestedmap": map[string]interface{}{
							"nested1": "nestedval3",
							"nested2": "nestedval4",
						},
					},
				},
			},
			expected: &parser.MapObj{
				Prefix: "",
				Val: []parser.Object{
					&parser.SliceObj{
						Prefix: "data",
						Val: []parser.Object{
							&parser.MapObj{
								Prefix: "data",
								Val: []parser.Object{
									&parser.StringObj{
										Prefix: "data_key1",
										Val:    "val1",
									},
									&parser.MapObj{
										Prefix: "data_nestedmap",
										Val: []parser.Object{
											&parser.StringObj{
												Prefix: "data_nestedmap_nested1",
												Val:    "nestedval1",
											},
											&parser.StringObj{
												Prefix: "data_nestedmap_nested2",
												Val:    "nestedval2",
											},
										},
									},
								},
							},
							&parser.MapObj{
								Prefix: "data",
								Val: []parser.Object{
									&parser.StringObj{
										Prefix: "data_key1",
										Val:    "val4",
									},
									&parser.MapObj{
										Prefix: "data_nestedmap",
										Val: []parser.Object{
											&parser.StringObj{
												Prefix: "data_nestedmap_nested1",
												Val:    "nestedval3",
											},
											&parser.StringObj{
												Prefix: "data_nestedmap_nested2",
												Val:    "nestedval4",
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			description: "complex map with nested maps and slices",
			input: map[string]interface{}{
				"data": []interface{}{
					map[string]interface{}{
						"key1": "val1",
						"nestedmap": map[string]interface{}{
							"nested1": "nestedval1",
							"nested2": "nestedval2",
						},
						"nestedslice": []interface{}{
							map[string]interface{}{
								"nestedslicemap1": "nestedslicemapval1",
								"nestedslicemap2": "nestedslicemapval2",
							},
							map[string]interface{}{
								"nestedslicemap1": "nestedslicemapval3",
								"nestedslicemap2": "nestedslicemapval4",
							},
						},
					},
					map[string]interface{}{
						"key1": "val4",
						"nestedmap": map[string]interface{}{
							"nested1": "nestedval3",
							"nested2": "nestedval4",
						},
						"nestedslice": []interface{}{
							map[string]interface{}{
								"nestedslicemap1": "nestedslicemapval5",
								"nestedslicemap2": "nestedslicemapval6",
							},
							map[string]interface{}{
								"nestedslicemap1": "nestedslicemapval7",
								"nestedslicemap2": "nestedslicemapval8",
							},
						},
					},
				},
			},
			expected: &parser.MapObj{
				Prefix: "",
				Val: []parser.Object{
					&parser.SliceObj{
						Prefix: "data",
						Val: []parser.Object{
							&parser.MapObj{
								Prefix: "data",
								Val: []parser.Object{
									&parser.StringObj{
										Prefix: "data_key1",
										Val:    "val1",
									},
									&parser.MapObj{
										Prefix: "data_nestedmap",
										Val: []parser.Object{
											&parser.StringObj{
												Prefix: "data_nestedmap_nested1",
												Val:    "nestedval1",
											},
											&parser.StringObj{
												Prefix: "data_nestedmap_nested2",
												Val:    "nestedval2",
											},
										},
									},
									&parser.SliceObj{
										Prefix: "data_nestedslice",
										Val: []parser.Object{
											&parser.MapObj{
												Prefix: "data_nestedslice",
												Val: []parser.Object{
													&parser.StringObj{
														Prefix: "data_nestedslice_nestedslicemap1",
														Val:    "nestedslicemapval1",
													},
													&parser.StringObj{
														Prefix: "data_nestedslice_nestedslicemap2",
														Val:    "nestedslicemapval2",
													},
												},
											},
											&parser.MapObj{
												Prefix: "data_nestedslice",
												Val: []parser.Object{
													&parser.StringObj{
														Prefix: "data_nestedslice_nestedslicemap1",
														Val:    "nestedslicemapval3",
													},
													&parser.StringObj{
														Prefix: "data_nestedslice_nestedslicemap2",
														Val:    "nestedslicemapval4",
													},
												},
											},
										},
									},
								},
							},
							&parser.MapObj{
								Prefix: "data",
								Val: []parser.Object{
									&parser.StringObj{
										Prefix: "data_key1",
										Val:    "val4",
									},
									&parser.MapObj{
										Prefix: "data_nestedmap",
										Val: []parser.Object{
											&parser.StringObj{
												Prefix: "data_nestedmap_nested1",
												Val:    "nestedval3",
											},
											&parser.StringObj{
												Prefix: "data_nestedmap_nested2",
												Val:    "nestedval4",
											},
										},
									},
									&parser.SliceObj{
										Prefix: "data_nestedslice",
										Val: []parser.Object{
											&parser.MapObj{
												Prefix: "data_nestedslice",
												Val: []parser.Object{
													&parser.StringObj{
														Prefix: "data_nestedslice_nestedslicemap1",
														Val:    "nestedslicemapval5",
													},
													&parser.StringObj{
														Prefix: "data_nestedslice_nestedslicemap2",
														Val:    "nestedslicemapval6",
													},
												},
											},
											&parser.MapObj{
												Prefix: "data_nestedslice",
												Val: []parser.Object{
													&parser.StringObj{
														Prefix: "data_nestedslice_nestedslicemap1",
														Val:    "nestedslicemapval7",
													},
													&parser.StringObj{
														Prefix: "data_nestedslice_nestedslicemap2",
														Val:    "nestedslicemapval8",
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.description, func(t *testing.T) {
			actual := parser.Transform(testcase.input)
			assert.Equal(t, testcase.expected, actual)
			spew.Dump(actual)
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
