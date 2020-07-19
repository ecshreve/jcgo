package parser_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ecshreve/jcgo/pkg/parser"
)

func TestTruncateColumnHeaders(t *testing.T) {
	testcases := []struct {
		description string
		input       []string
		expected    []string
	}{
		{
			description: "empty slice",
			input:       []string{},
			expected:    nil,
		},
		{
			description: "slice with one element",
			input:       []string{"single_header"},
			expected:    []string{"single_header"},
		},
		{
			description: "simple",
			input:       []string{"data_test_one", "data_test_two", "data_three"},
			expected:    []string{"test_one", "test_two", "three"},
		},
		{
			description: "long",
			input: []string{
				"data_one_two_three_four_after_header1",
				"data_one_two_three_four_after_header2",
				"data_one_two_three_four_after_header3",
				"data_one_two_three_four_after_header4",
				"data_one_two_three_four_after_header5",
				"data_one_two_three_four_before_header1",
				"data_one_two_three_four_before_header2",
				"data_one_two_three_four_before_header3",
				"data_one_two_three_four_before_header4",
				"data_one_two_three_four_before_header5",
				"data_one_two_three_four_header6",
				"data_one_two_three_four_events_header7",
				"data_one_two_three_four_events_header8",
				"data_one_two_three_four_header4",
			},
			expected: []string{
				"after_header1",
				"after_header2",
				"after_header3",
				"after_header4",
				"after_header5",
				"before_header1",
				"before_header2",
				"before_header3",
				"before_header4",
				"before_header5",
				"header6",
				"events_header7",
				"events_header8",
				"header4",
			},
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.description, func(t *testing.T) {
			actual := parser.TruncateColumnHeaders(testcase.input)
			assert.Equal(t, testcase.expected, actual)
		})
	}
}
