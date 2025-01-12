package tests

import (
	"testing"

	utils "github.com/abroudoux/branch/internal/utils"
)

func TestCleanString(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "String with leading asterisk and spaces",
			input:    "* Hello World ",
			expected: "Hello World",
		},
		{
			name:     "String with only leading asterisk",
			input:    "*Hello World",
			expected: "Hello World",
		},
		{
			name:     "String with leading and trailing spaces",
			input:    "  Hello World  ",
			expected: "Hello World",
		},
		{
			name:     "String with no asterisk or spaces",
			input:    "Hello World",
			expected: "Hello World",
		},
		{
			name:     "Empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "String with only spaces",
			input:    "    ",
			expected: "",
		},
		{
			name:     "String with multiple asterisks",
			input:    "** Hello World",
			expected: "* Hello World",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := utils.CleanString(tc.input)
			if result != tc.expected {
				t.Errorf("Expected '%s', but got '%s'", tc.expected, result)
			}
		})
	}
}
