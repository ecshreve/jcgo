package object_test

import (
	"testing"

	"github.com/samsarahq/go/snapshotter"
	"github.com/stretchr/testify/assert"

	oo "github.com/ecshreve/jcgo/internal/object"
	"github.com/ecshreve/jcgo/internal/testdata"
)

func TestSimpleParseMapObject(t *testing.T) {
	snap := snapshotter.New(t)
	defer snap.Verify()
	data := testdata.NewObjectTestData()

	testcases := []struct {
		description string
		input       map[string]interface{}
		expectError bool
	}{
		{
			description: "simple map",
			input:       data.SimpleMapInput,
			expectError: false,
		},
		{
			description: "simple nested map object",
			input:       data.SimpleNestedMapInput,
			expectError: false,
		},
		{
			description: "simple mixed nested map object",
			input:       data.SimpleMixedNestedMapInput,
			expectError: false,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.description, func(t *testing.T) {
			obj, err := oo.NewMapObj("", testcase.input)
			assert.NoError(t, err)
			snap.Snapshot(testcase.description+" -- MapObj", obj)

			parsed, err := obj.Parse()
			assert.Equal(t, testcase.expectError, err != nil)
			snap.Snapshot(testcase.description+" -- parsed", parsed)
		})
	}
}
