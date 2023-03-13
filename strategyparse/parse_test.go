package strategyparse

import (
	"reflect"
	"strings"
	"testing"
)

func TestParseArgs(t *testing.T) {
	testCases := []struct {
		name          string
		input         []string
		expected      GameStrategy
		expectedError string
	}{
		//story 1
		{
			name:          "Invalid number of arguments",
			input:         []string{"10"},
			expected:      GameStrategy{},
			expectedError: "invalid number of arguments",
		},
		{
			name:          "First argument valid second argument invalid",
			input:         []string{"20", "foo"},
			expected:      GameStrategy{},
			expectedError: "invalid syntax",
		},
		{
			name:          "First argument invalid second argument nvalid",
			input:         []string{"bar", "20"},
			expected:      GameStrategy{},
			expectedError: "invalid syntax",
		},
		// empty string
		{
			name:          "Both Invalid argument",
			input:         []string{"foo", "bar"},
			expected:      GameStrategy{},
			expectedError: "invalid syntax",
		},
		{
			name:          "Valid argument",
			input:         []string{"20", "30"},
			expected:      GameStrategy{Player1: 20, Player2: 30, Story: 1},
			expectedError: "",
		},
		{
			name:          "Out of range negative input for both arguments",
			input:         []string{"-20", "-50"},
			expected:      GameStrategy{},
			expectedError: "must specify a number greater than 0",
		},
		// if we have "-" anywhere in the args it satisfies one of the stories
		// even if it is not intended. e.g. "20-"
		{
			name:          "Input containing '-' after the number",
			input:         []string{"20-", "50-"},
			expected:      GameStrategy{},
			expectedError: "invalid syntax",
		},
		{
			name:          "string containing just '-' ",
			input:         []string{"--", "--"},
			expected:      GameStrategy{},
			expectedError: "invalid syntax",
		},
		//story 2
		{
			name:          "Valid arguments for story 2",
			input:         []string{"10", "1-100"},
			expected:      GameStrategy{Player1: 10, Story: 2},
			expectedError: "",
		},
		{
			name:          "InValid arguments for story 2",
			input:         []string{"10", "-"},
			expected:      GameStrategy{},
			expectedError: "invalid syntax",
		},

		{
			name:          "InValid arguments that doesn't satisfy any if condition",
			input:         []string{"1-100", "15"},
			expected:      GameStrategy{},
			expectedError: "invalid syntax",
		},
		//story 3
		{
			name:          "valid arguments for story 3",
			input:         []string{"1-100", "1-100"},
			expected:      GameStrategy{Story: 3},
			expectedError: "",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := ParseAndValidateArgs(tc.input)

			if err != nil {
				if !strings.Contains(err.Error(), tc.expectedError) {
					t.Errorf("Expected %v Got %v", tc.expectedError, err)
				}
			}
			if !reflect.DeepEqual(got, tc.expected) {
				t.Errorf("Expected %v Got %v", tc.expected, got)
			}
		})
	}

}
