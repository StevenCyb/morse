package translator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTextToMorse(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		input     string
		separator string
		expected  string
	}{
		{
			name:      "simple letters",
			input:     "AB",
			separator: " ",
			expected:  ".- -...",
		},
		{
			name:      "digits",
			input:     "12",
			separator: " ",
			expected:  ".---- ..---",
		},
		{
			name:      "punctuation",
			input:     ".?",
			separator: " ",
			expected:  ".-.-.- ..--..",
		},
		{
			name:      "unknown char",
			input:     "A*B",
			separator: " ",
			expected:  ".- # -...",
		},
		{
			name:      "empty string",
			input:     "",
			separator: " ",
			expected:  "",
		},
		{
			name:      "words",
			input:     "a b",
			separator: " ",
			expected:  ".-  -...",
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := TextToMorse(tt.input, tt.separator)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestMorseToText(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name      string
		input     string
		separator string
		expected  string
	}{
		{
			name:      "simple letters",
			input:     ".- -...",
			separator: " ",
			expected:  "AB",
		},
		{
			name:      "digits",
			input:     ".---- ..---",
			separator: " ",
			expected:  "12",
		},
		{
			name:      "punctuation",
			input:     ".-.-.- ..--..",
			separator: " ",
			expected:  ".?",
		},
		{
			name:      "unknown morse",
			input:     ".- ....-.- -...",
			separator: " ",
			expected:  "A#B",
		},
		{
			name:      "empty string",
			input:     "",
			separator: " ",
			expected:  "",
		},
		{
			name:      "words",
			input:     ".-  -...",
			separator: " ",
			expected:  "A B",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := MorseToText(tt.input, tt.separator)
			assert.Equal(t, tt.expected, result)
		})
	}
}
