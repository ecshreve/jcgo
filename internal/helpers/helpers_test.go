package helpers_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ecshreve/jcgo/internal/helpers"
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
				"data_organization_groups_dispatchRoute_auditLogs_afterState_arrivedAt",
				"data_organization_groups_dispatchRoute_auditLogs_afterState_departureTimeMs",
				"data_organization_groups_dispatchRoute_auditLogs_afterState_destinationName",
				"data_organization_groups_dispatchRoute_auditLogs_afterState_id",
				"data_organization_groups_dispatchRoute_auditLogs_afterState_jobState",
				"data_organization_groups_dispatchRoute_auditLogs_beforeState_arrivedAt",
				"data_organization_groups_dispatchRoute_auditLogs_beforeState_departureTimeMs",
				"data_organization_groups_dispatchRoute_auditLogs_beforeState_destinationName",
				"data_organization_groups_dispatchRoute_auditLogs_beforeState_id",
				"data_organization_groups_dispatchRoute_auditLogs_beforeState_jobState",
				"data_organization_groups_dispatchRoute_auditLogs_changedAtMs",
				"data_organization_groups_dispatchRoute_auditLogs_events_eventAt",
				"data_organization_groups_dispatchRoute_auditLogs_events_eventType",
				"data_organization_groups_dispatchRoute_auditLogs_id",
			},
			expected: []string{
				"afterState_arrivedAt",
				"afterState_departureTimeMs",
				"afterState_destinationName",
				"afterState_id",
				"afterState_jobState",
				"beforeState_arrivedAt",
				"beforeState_departureTimeMs",
				"beforeState_destinationName",
				"beforeState_id",
				"beforeState_jobState",
				"changedAtMs",
				"events_eventAt",
				"events_eventType",
				"id",
			},
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.description, func(t *testing.T) {
			actual := helpers.TruncateColumnHeaders(testcase.input)
			assert.Equal(t, testcase.expected, actual)
		})
	}
}
