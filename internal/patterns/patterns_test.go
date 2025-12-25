package patterns

import (
	"testing"
)

func TestCompilePatternMap(t *testing.T) {
	tests := []struct {
		name		string
		patterns	map[string][]string
		expectedLen	int
	}{
		{
			name: "input 0 patterns",
			patterns: map[string][]string{},
			expectedLen: 0,
		},
		{
			name: "input 1 patterns",
			patterns: map[string][]string{
				"pattern1": { "pattern1" },
			},
			expectedLen: 1,
		},
		{
			name: "input 2 patterns",
			patterns: map[string][]string{
				"pattern1": { "pattern1" },
				"pattern2": { "pattern2" },
			},
			expectedLen: 2,
		},
		{
			name: "input 3 patterns",
			patterns: map[string][]string{
				"pattern1": { "pattern1" },
				"pattern2": { "pattern2" },
				"pattern3": { "pattern3" },
			},
			expectedLen: 3,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			compiledPatterns := compilePatternMap(test.patterns)
			compiledPatternsLen := len(compiledPatterns)

			for key, patterns := range compiledPatterns {
				for _, pattern := range patterns {
					if !pattern.MatchString(key) {
						t.Errorf("MatchString() for %v inside %v false, want true", key, pattern)
					}
				}
			}

			if compiledPatternsLen != test.expectedLen {
				t.Errorf("compilePatternMap() len = %v, want %v", compiledPatternsLen, test.expectedLen)
			}
		})
	}
}


func TestCompilePatternSlice(t *testing.T) {
	tests := []struct {
		name		string
		patterns	[]string
		expectedLen	int
	}{
		{
			name: "input 0 patterns",	
			patterns: []string {
			},
			expectedLen: 0,
		},
		{
			name: "input 1 patterns",	
			patterns: []string {
				"pattern1",
			},
			expectedLen: 1,
		},
		{
			name: "input 2 patterns",	
			patterns: []string {
				"pattern1",
				"pattern2",
			},
			expectedLen: 2,
		},
		{
			name: "input 3 patterns",	
			patterns: []string {
				"pattern1",
				"pattern2",
				"pattern3",
			},
			expectedLen: 3,
		},
	}

	for _, test := range tests {
		compiledPatterns := compilePatternSlice(test.patterns)
		compiledPatternsLen := len(compiledPatterns)

		for i, pattern := range compiledPatterns {
			if !pattern.MatchString(test.patterns[i]) {
				t.Errorf("MatchString() for %v inside %v false, want true", test.patterns[i], pattern)
			}
		}

		if compiledPatternsLen != test.expectedLen {
			t.Errorf("compilePatternMap() len = %v, want %v", compiledPatternsLen, test.expectedLen)
		}
	}

}
