package display

import (
	"testing"

	"github.com/fatih/color"
	"github.com/stretchr/testify/assert"
)

func TestColorizeMorseString(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "only dots",
			input:    "...",
			expected: color.CyanString(".") + color.CyanString(".") + color.CyanString("."),
		},
		{
			name:     "only dashes",
			input:    "---",
			expected: color.RedString("-") + color.RedString("-") + color.RedString("-"),
		},
		{
			name:     "mixed",
			input:    ".-.",
			expected: color.CyanString(".") + color.RedString("-") + color.CyanString("."),
		},
		{
			name:     "with spaces",
			input:    ". - .",
			expected: color.CyanString(".") + " " + color.RedString("-") + " " + color.CyanString("."),
		},
		{
			name:     "empty string",
			input:    "",
			expected: "",
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := ColorizeMorseString(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}
