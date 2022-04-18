package portainer

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidateSetup(t *testing.T) {
	var tests = []struct {
		name     string
		setup    string
		expected bool
	}{
		{
			name:     "Invalid Length",
			setup:    "[{]",
			expected: false,
		},
		{
			name:     "Invalid Characters",
			setup:    "<{}>",
			expected: false,
		},
		{
			name:     "Invalid Setup",
			setup:    "([)]{()}()",
			expected: false,
		},
		{
			name:     "Valid Setup Example 1",
			setup:    "([])",
			expected: true,
		},
		{
			name:     "Valid Setup Example 2",
			setup:    "([]{}())",
			expected: true,
		},
		{
			name:     "Valid Setup Example 3",
			setup:    "(){}()",
			expected: true,
		},
		{
			name:     "Valid Complex Setup",
			setup:    "[{()}{[()()]}{[()]}]",
			expected: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := ValidateSetup(test.setup)
			require.Equal(t, test.expected, result)
		})
	}
}
