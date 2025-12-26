package patterns

import (
	"regexp"
	"testing"
)

func TestCompilePatternGroups(t *testing.T) {
	tests := []struct {
		name		string
		patterns	[]PatternGroup
		expectedLen	int
	}{
		{
			name: "input 0 patterns",
			patterns: []PatternGroup{},
			expectedLen: 0,
		},
		{
			name: "input 1 patterns",
			patterns: []PatternGroup{
				{ Key: "pattern1", Patterns: []Pattern{ "pattern1" }},
			},
			expectedLen: 1,
		},
		{
			name: "input 2 patterns",
			patterns: []PatternGroup{
				{ Key: "pattern1", Patterns: []Pattern{ "pattern1" }},
				{ Key: "pattern2", Patterns: []Pattern{ "pattern2" }},
			},
			expectedLen: 2,
		},
		{
			name: "input 3 patterns",
			patterns: []PatternGroup{
				{ Key: "pattern1", Patterns: []Pattern{ "pattern1" }},
				{ Key: "pattern2", Patterns: []Pattern{ "pattern2" }},
				{ Key: "pattern3", Patterns: []Pattern{ "pattern3" }},
			},
			expectedLen: 3,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			compiledPatternGroups := compilePatternGroups(test.patterns)
			compiledPatternsLen := len(compiledPatternGroups)

			for _, group := range compiledPatternGroups {
				for _, pattern := range group.Patterns {
					if !(*regexp.Regexp)(pattern).MatchString(group.Key) {
						t.Errorf("MatchString() for %v inside %v false, want true", group.Key, pattern)
					}
				}
			}

			if compiledPatternsLen != test.expectedLen {
				t.Errorf("compilePatternMap() len = %v, want %v", compiledPatternsLen, test.expectedLen)
			}
		})
	}
}


func TestCompilePatterns(t *testing.T) {
	tests := []struct {
		name		string
		patterns	[]Pattern
		expectedLen	int
	}{
		{
			name: "input 0 patterns",	
			patterns: []Pattern {
			},
			expectedLen: 0,
		},
		{
			name: "input 1 patterns",	
			patterns: []Pattern {
				"pattern1",
			},
			expectedLen: 1,
		},
		{
			name: "input 2 patterns",	
			patterns: []Pattern {
				"pattern1",
				"pattern2",
			},
			expectedLen: 2,
		},
		{
			name: "input 3 patterns",	
			patterns: []Pattern {
				"pattern1",
				"pattern2",
				"pattern3",
			},
			expectedLen: 3,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			compiledPatterns := compilePatterns(test.patterns)
			compiledPatternsLen := len(compiledPatterns)

			for i, pattern := range compiledPatterns {
				if !(*regexp.Regexp)(pattern).MatchString(string(test.patterns[i])) {
					t.Errorf("MatchString() for %v inside %v false, want true", test.patterns[i], pattern)
				}
			}

			if compiledPatternsLen != test.expectedLen {
				t.Errorf("compilePatternMap() len = %v, want %v", compiledPatternsLen, test.expectedLen)
			}
		})
	}

}
