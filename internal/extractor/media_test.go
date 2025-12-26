package extractor

import (
	"log/slog"
	"testing"
)

func TestExtractTitle(t *testing.T) {
	log := slog.Default()
	tests := []struct {
		name		string
		input		[]string
		expected	string
	}{
		{
			name: 		"parse format <title>.<year (optional)>.<resolution, codec, source, or audio>",
			input:		[]string{
				"MOVIE",
				"TITLE",
				"2020",
				"1080P",
			},
			expected:	"MOVIE.TITLE",
		},
		{
			name: 		"parse format <title>.<year (optional)>.<season or ep>",
			input:		[]string{
				"MOVIE",
				"TITLE",
				"2020",
				"S4",
			},
			expected:	"MOVIE.TITLE",
		},
		{
			name: 		"parse format <title>.<year (optional)>.<file ext>",
			input:		[]string{
				"MOVIE",
				"TITLE",
				"2020",
				".MP4",
			},
			expected:	"MOVIE.TITLE",
		},
		{
			name: 		"parse format <title>.<year (optional)>",
			input:		[]string{
				"MOVIE",
				"TITLE",
				"2020",
			},
			expected:	"MOVIE.TITLE",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			title := extractTitle(test.input, log)	
			if title != test.expected {
				t.Errorf("extractTitle = %v, want %v", title, test.expected)
			}

		})
	}
}

func TestParseResolution(t *testing.T) {
	tests := []struct {
		name		string
		input		[]string
		expected	string
	}{
		{
			name:		"test valid resolution",
			input:		[]string{
				"2160I",
			},
			expected:	"4K",
		},
		{
			name:		"test invalid resolution",
			input:		[]string{
				"2160X",
			},
			expected:	"",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := parseResolution(test.input)
			if	res != test.expected {
				t.Errorf("parseResolution = %v, want %v", res, test.expected)
			}
		})
	}
}
func TestParseCodec(t *testing.T) {
	tests := []struct {
		name		string
		input		[]string
		expected	string
	}{
		{
			name:		"test codec without seperators",
			input:		[]string{
				"AOV1",
			},
			expected:	"AV1",
		},
		{
			name:		"test codec with seperators",
			input:		[]string{
				"SVT",
				"AV1",
			},
			expected:	"AV1",
		},
		{
			name:		"test invalid codec with seperators",
			input:		[]string{
				"INCORRECT",
				"SVT",
				"AV1",
			},
			expected:	"",
		},
		{
			name:		"test invalid codec without seperators",
			input:		[]string{
				"INCORRECTSVT",
			},
			expected:	"",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := parseCodec(test.input)
			if	res != test.expected {
				t.Errorf("parseCodec = %v, want %v", res, test.expected)
			}
		})
	}
}


func TestParseSource(t *testing.T) {
	tests := []struct {
		name		string
		input		[]string
		expected	string
	}{
		{
			name:		"test source without seperators",
			input:		[]string{
				"BLURAY",
			},
			expected:	"BluRay",
		},
		{
			name:		"test source with seperators",
			input:		[]string{
				"BD",
				"RIP",
			},
			expected:	"BluRay",
		},
		{
			name:		"test invalid source with seperators",
			input:		[]string{
				"INCORRECT",
				"BD",
				"RIP",
			},
			expected:	"",
		},
		{
			name:		"test invalid source without seperators",
			input:		[]string{
				"INCORRECTBLURAY",
			},
			expected:	"",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := parseSource(test.input)
			if	res != test.expected {
				t.Errorf("parseCodec = %v, want %v", res, test.expected)
			}
		})
	}
}

func TestParseAudio(t *testing.T) {
	tests := []struct {
		name		string
		input		[]string
		expected	string
	}{
		{
			name:		"test audio without seperators",
			input:		[]string{
				"DTSX",
			},
			expected:	"DTS-X",
		},
		{
			name:		"test audio with seperators",
			input:		[]string{
				"DTS",
				"X",
			},
			expected:	"DTS-X",
		},
		{
			name:		"test invalid audio without seperators",
			input:		[]string{
				"INCORRECTDTSX",
			},
			expected:	"",
		},
		{
			name:		"test invalid audio with seperators",
			input:		[]string{
				"INCORRECT",
				"DTSX",
			},
			expected:	"",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := parseAudio(test.input)
			if	res != test.expected {
				t.Errorf("parseCodec = %v, want %v", res, test.expected)
			}
		})
	}
}


func TestParseYear(t *testing.T) {
	tests := []struct {
		name			string
		input			string
		expectedValue	int
		expectedErr		bool
	}{
		{
			name:			"test valid year",		
			input:			"2000",
			expectedValue:	2000,
			expectedErr: 	false,
		},
		{
			name:			"test too early year",			
			input: 			"1900",
			expectedValue: 	0,
			expectedErr: 	true,
		},
		{
			name:			"test too late year",		
			input: 			"3000",
			expectedValue: 	0,
			expectedErr: 	true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			year, err := parseYear(test.input)
			if year != test.expectedValue {
				t.Errorf("parseYear = %v, want %v", year, test.expectedValue)
			}
			if err != nil && !test.expectedErr {
				t.Errorf("err = %v, no error expected", err)
			}
		})
	}
}
