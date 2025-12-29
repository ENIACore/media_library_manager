package extractor

import (
	"testing"
)

func TestParseVideoExt(t *testing.T) {
	tests := []struct {
		name			string
		input			[]string
		expected		string
	}{
		{
			name:		"valid ext",
			input:		[]string{
				"MP4",
			},
			expected: "MP4",
		},
		{
			name:		"invalid ext",
			input:		[]string{
				"MP6",
			},
			expected: "",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ext := parseVideoExt(test.input)
			if ext != test.expected {
				t.Errorf("parseVideoExt = %v, want %v", ext, test.expected)
			}
		})
	}
}

func TestParseSubtitleExt(t *testing.T) {
	tests := []struct {
		name			string
		input			[]string
		expected		string
	}{
		{
			name:		"valid ext",
			input:		[]string{
				"SRT",
			},
			expected: "SRT",
		},
		{
			name:		"invalid ext",
			input:		[]string{
				"SRTT",
			},
			expected: "",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ext := parseSubtitleExt(test.input)
			if ext != test.expected {
				t.Errorf("parseSubtitleExt = %v, want %v", ext, test.expected)
			}
		})
	}
}


func TestParseAudioExt(t *testing.T) {
	tests := []struct {
		name			string
		input			[]string
		expected		string
	}{
		{
			name:		"valid ext",
			input:		[]string{
				"MP3",
			},
			expected: "MP3",
		},
		{
			name:		"invalid ext",
			input:		[]string{
				"MP6",
			},
			expected: "",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ext := parseAudioExt(test.input)
			if ext != test.expected {
				t.Errorf("parseAudioExt = %v, want %v", ext, test.expected)
			}
		})
	}
}
