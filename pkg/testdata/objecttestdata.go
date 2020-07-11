package testdata

type ObjectTestData struct {
	SimpleMapInput            map[string]interface{}
	SimpleNestedMapInput      map[string]interface{}
	SimpleMixedNestedMapInput map[string]interface{}
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

	return &ObjectTestData{
		SimpleMapInput:            simpleMapInput,
		SimpleNestedMapInput:      simpleNestedMapInput,
		SimpleMixedNestedMapInput: simpleMixedNestedMapInput,
	}
}
