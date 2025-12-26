package patterns

import (
	`regexp`
	`sync`
)

type Pattern string

type CompiledPattern regexp.Regexp

type PatternGroup struct {
	Key			string
	Patterns	[]Pattern	
}

type CompiledPatternGroup struct {
	Key			string
	Patterns	[]*CompiledPattern
}

var LanguagePatternGroups = []PatternGroup{
	{Key: `ENGLISH`, Patterns: []Pattern{ `ENGLISH`, `ENG`, `EN`, }},
	{Key: `SPANISH`, Patterns: []Pattern{ `SPANISH`, `CASTELLANO`, `SPA`, `ES`, `ESPAÑOL`, }},
	{Key: `FRENCH`, Patterns: []Pattern{ `FRENCH`, `FRA`, `FR`, }},
	{Key: `GERMAN`, Patterns: []Pattern{ `GERMAN`, `DEUTSCH`, `GER`, `DE`, `GERMAN`, }},
	{Key: `ITALIAN`, Patterns: []Pattern{ `ITALIAN`, `ITA`, `IT`, `ITALIANO`, }},
	{Key: `PORTUGUESE`, Patterns: []Pattern{ `PORTUGUESE`, `PORTUGUES`, `POR`, `PT`, }},
	{Key: `BRAZILIAN_PORTUGUESE`, Patterns: []Pattern{ `BRAZILIAN`, `BRAZIL`, `BR`, `PORTUGUESE.BR`, `PT.BR`, }},
	{Key: `RUSSIAN`, Patterns: []Pattern{ `RUSSIAN`, `RUS`, `RU`, }},
	{Key: `JAPANESE`, Patterns: []Pattern{ `JAPANESE`, `JAP`, `JPN`, `JP`, `JA`, }},
	{Key: `KOREAN`, Patterns: []Pattern{ `KOREAN`, `KOR`, `KO`, `KR`, }},
	{Key: `ARABIC`, Patterns: []Pattern{ `ARABIC`, `ARA`, `AR`, }},
	{Key: `HEBREW`, Patterns: []Pattern{ `HEBREW`, `HEB`, `HE`, }},
	{Key: `THAI`, Patterns: []Pattern{ `THAI`, `THA`, `TH`, }},
	{Key: `TURKISH`, Patterns: []Pattern{ `TURKISH`, `TUR`, `TR`, }},
	{Key: `GREEK`, Patterns: []Pattern{ `GREEK`, `GRE`, `EL`, }},
	{Key: `POLISH`, Patterns: []Pattern{ `POLISH`, `POL`, `PL`, `POLSKI`, }},
	{Key: `HUNGARIAN`, Patterns: []Pattern{ `HUNGARIAN`, `HUN`, `HU`, `MAGYAR`, }},
	{Key: `CZECH`, Patterns: []Pattern{ `CZECH`, `CZE`, `CS`, }},
	{Key: `CHINESE`, Patterns: []Pattern{ `CHINESE`, `CHI`, `ZH`, }},
}

var ExtrasPatterns = []Pattern{
	`EXTRA[S]?`,
	`FEATURETTE[S]?`,
	`BEHIND\.THE\.SCENE[S]?`,
	`BTS`,
	`DELETED\.SCENE[S]?`,
	`MAKING\.OF`,
	`TRAILER`,
	`BONUS`,
	`DOCUMENTARY`,
	`DOCUMENTARIES`,
}

var (
	GetLanguagePatternGroups = sync.OnceValue(func() []CompiledPatternGroup {
		return compilePatternGroups(LanguagePatternGroups)
	})
	GetExtrasPatterns = sync.OnceValue(func() []*CompiledPattern {
		return compilePatterns(ExtrasPatterns)
	})
)

func compilePatternGroups(patternGroups []PatternGroup) []CompiledPatternGroup {
	res := make([]CompiledPatternGroup, len(patternGroups))
	for i, group := range patternGroups {
		patterns := make([]*CompiledPattern, len(group.Patterns))
		for j, pattern := range group.Patterns {
			patterns[j] = (*CompiledPattern)(regexp.MustCompile(string(pattern)))
		}
		res[i] = CompiledPatternGroup{
			Key: 		group.Key,
			Patterns: 	patterns,
		}
	}
	return res
}

func compilePatterns(patterns []Pattern) []*CompiledPattern {
	result := make([]*CompiledPattern, len(patterns))
	for i, pattern := range patterns {
		result[i] = (*CompiledPattern)(regexp.MustCompile(string(pattern)))
	}
	return result
}
