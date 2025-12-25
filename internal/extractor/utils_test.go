package extractor

import (
	"testing"
	"log/slog"
	"fmt"
)

func TestSanitizeName(t *testing.T) {
	logger := slog.Default()
	tests := []struct{
		name		string
		input		string
		expected	string
	}{
		{
			// Tests removing all special characters, removing spaces, removing quotations, capitalization, removing trailing and preceding ".", combining names by a single "."
			name:		"remove all specicial characters",
			input:		"!@#$%^&*()-_+=\"'. mOvIe .!@#$%^&*()-_+=\"' tItLe !@#$%^&*()-_+=\"'.",
			expected:	"MOVIE.TITLE",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			sanitizedName := sanitizeName(test.input, logger)
			if sanitizedName != test.expected {
				t.Errorf("sanitizeName() = %v, want %v", sanitizedName, test.expected)
			}
		})
	}
}

func TestMatchParts(t *testing.T) {
	tests := []struct{
		name	string
	}{
		{
			name: "test",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

		})
	}
}
