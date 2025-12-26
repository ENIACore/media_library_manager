package extractor

import (
	//"log/slog"
	"testing"
)

func TestExtractTitle(t *testing.T) {
	//log := slog.Default()
	tests := []struct {
		name		string
		input		[]string
		expected	string
	}{
		{
			name: 		"format <title>.<year (optional)>.<resolution, codec, source, or audio>",
			input:		[]string{
				"MOVIE",
				"TITLE",
				"2020",
				"1080P",
			},
			expected:	"MOVIE.TITLE",
		},
		{
			name: 		"format <title>.<year (optional)>.<season or ep>",
			input:		[]string{
				"MOVIE",
				"TITLE",
				"2020",
				"S4",
			},
			expected:	"MOVIE.TITLE",
		},
		{
			name: 		"format <title>.<year (optional)>.<file ext>",
			input:		[]string{
				"MOVIE",
				"TITLE",
				"2020",
				".MP4",
			},
			expected:	"MOVIE.TITLE",
		},
		{
			name: 		"format <title>.<year (optional)>",
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
			/*
			title := extractTitle(test.input, log)	
			if title != test.expected {
				t.Errorf("extractTitle = %v, want %v", title, test.expected)
			}
			*/

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
			name:		"valid resolution",
			input:		[]string{
				"2160I",
			},
			expected:	"4K",
		},
		{
			name:		"invalid resolution",
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
			name:		"codec without seperators",
			input:		[]string{
				"AOV1",
			},
			expected:	"AV1",
		},
		{
			name:		"codec with seperators",
			input:		[]string{
				"SVT",
				"AV1",
			},
			expected:	"AV1",
		},
		{
			name:		"invalid codec with seperators",
			input:		[]string{
				"INCORRECT",
				"SVT",
				"AV1",
			},
			expected:	"",
		},
		{
			name:		"invalid codec without seperators",
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
			name:		"source without seperators",
			input:		[]string{
				"BLURAY",
			},
			expected:	"BluRay",
		},
		{
			name:		"source with seperators",
			input:		[]string{
				"BD",
				"RIP",
			},
			expected:	"BluRay",
		},
		{
			name:		"invalid source with seperators",
			input:		[]string{
				"INCORRECT",
				"BD",
				"RIP",
			},
			expected:	"",
		},
		{
			name:		"invalid source without seperators",
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
			name:		"audio without seperators",
			input:		[]string{
				"DTSX",
			},
			expected:	"DTS-X",
		},
		{
			name:		"audio with seperators",
			input:		[]string{
				"DTS",
				"X",
			},
			expected:	"DTS-X",
		},
		{
			name:		"invalid audio without seperators",
			input:		[]string{
				"INCORRECTDTSX",
			},
			expected:	"",
		},
		{
			name:		"invalid audio with seperators",
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
	}{
		{
			name:			"valid year",		
			input:			"2000",
			expectedValue:	2000,
		},
		{
			name:			"too early year",			
			input: 			"1900",
			expectedValue: 	-1,
		},
		{
			name:			"too late year",		
			input: 			"3000",
			expectedValue: 	-1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			year := parseYear(test.input)
			if year != test.expectedValue {
				t.Errorf("parseYear = %v, want %v", year, test.expectedValue)
			}
		})
	}
}

func TestParseSeason(t *testing.T) {
	tests := []struct {
		name			string
		input			[]string
		expected		int
	}{
		{
			name:		"valid season with number",
			input:		[]string{
				"S04",
				"1080P",
				"X265",
				"MP4",
			},
			expected: 4,
		},
		{
			name:		"valid season without number",
			input:		[]string{
				"SEASON",
				"1080P",
				"X265",
				"MP4",
			},
			expected: 0,
		},
		{
			name:		"invalid season",
			input:		[]string{
				"INCORRECTS04",
				"1080P",
				"X265",
				"MP4",
			},
			expected: -1,
		},
		{
			name:		"season at second segment",
			input:		[]string{
				"INCORRECT",
				"S04",
				"1080P",
				"X265",
				"MP4",
			},
			expected: -1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			season := parseSeason(test.input)
			if season != test.expected {
				t.Errorf("parseSeason = %v, want %v", season, test.expected)
			}
		})
	}
}

