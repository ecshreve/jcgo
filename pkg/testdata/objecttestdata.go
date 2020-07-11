package testdata

type ObjectTestData struct {
	SimpleMapInput            map[string]interface{}
	SimpleNestedMapInput      map[string]interface{}
	SimpleMixedNestedMapInput map[string]interface{}
	SimpleArrayMapInput       interface{}
	SimpleMapArrayInput       interface{}
	ComplexInput1             interface{}
	ComplexInput2             interface{}
}

func NewObjectTestData() *ObjectTestData {
	simpleMapInput := map[string]interface{}{
		"key1": "value1",
		"key2": "value2",
		"key3": "value3",
	}

	simpleNestedMapInput := map[string]interface{}{
		"outer": map[string]interface{}{
			"inner": "inner val",
		},
	}

	simpleMixedNestedMapInput := map[string]interface{}{
		"key": "val",
		"outer": map[string]interface{}{
			"inner": "inner val",
		},
	}

	simpleArrayMapInput := []interface{}{
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
		map[string]interface{}{
			"key1": "val7",
			"key2": "val8",
			"key3": "val9",
		},
	}

	simpleMapArrayInput := map[string]interface{}{
		"outer1": []interface{}{
			"val1",
			"val2",
			"val3",
		},
		"outer2": []interface{}{
			"val4",
			"val5",
			"val6",
		},
	}

	complexInput1 := map[string]interface{}{
		"outer1": []interface{}{
			map[string]interface{}{
				"key1": "val1",
				"key2": "val2",
				"key3": []interface{}{"nval1", "nval2"},
			},
		},
		"outer2": []interface{}{
			map[string]interface{}{
				"key1": "val3",
				"key2": "val4",
				"key3": []interface{}{"nval3", "nval4"},
			},
		},
	}

	complexInput2 := map[string]interface{}{
		"data": []interface{}{
			map[string]interface{}{
				"key1": "val1",
				"nestedmap": map[string]interface{}{
					"nested1": "nestedval1",
					"nested2": "nestedval2",
				},
				"nestedslice": []interface{}{
					map[string]interface{}{
						"before": []interface{}{
							map[string]interface{}{
								"nestedslicemap1": "nestedslicemapval1",
								"nestedslicemap2": "nestedslicemapval2",
							},
							map[string]interface{}{
								"nestedslicemap1": "nestedslicemapval3",
								"nestedslicemap2": "nestedslicemapval4",
							},
						},
						"after": []interface{}{
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
					map[string]interface{}{
						"before": []interface{}{
							map[string]interface{}{
								"nestedslicemap1": "nestedslicemapval11",
								"nestedslicemap2": "nestedslicemapval12",
							},
							map[string]interface{}{
								"nestedslicemap1": "nestedslicemapval13",
								"nestedslicemap2": "nestedslicemapval14",
							},
						},
						"after": []interface{}{
							map[string]interface{}{
								"nestedslicemap1": "nestedslicemapval15",
								"nestedslicemap2": "nestedslicemapval16",
							},
							map[string]interface{}{
								"nestedslicemap1": "nestedslicemapval17",
								"nestedslicemap2": "nestedslicemapval18",
							},
						},
					},
				},
			},
		},
	}

	return &ObjectTestData{
		SimpleMapInput:            simpleMapInput,
		SimpleNestedMapInput:      simpleNestedMapInput,
		SimpleMixedNestedMapInput: simpleMixedNestedMapInput,
		SimpleArrayMapInput:       simpleArrayMapInput,
		SimpleMapArrayInput:       simpleMapArrayInput,
		ComplexInput1:             complexInput1,
		ComplexInput2:             complexInput2,
	}
}
