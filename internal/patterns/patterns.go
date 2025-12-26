package patterns

import (
	`regexp`
	`sync`
)

var LanguagePatterns = map[string][]string{
    `ENGLISH`: { `ENGLISH`, `ENG`, `EN`, },
    `SPANISH`: { `SPANISH`, `CASTELLANO`, `SPA`, `ES`, `ESPAÑOL`, },
    `FRENCH`: { `FRENCH`, `FRA`, `FR`, },
    `GERMAN`: { `GERMAN`, `DEUTSCH`, `GER`, `DE`, `GERMAN`, },
    `ITALIAN`: { `ITALIAN`, `ITA`, `IT`, `ITALIANO`, },
    `PORTUGUESE`: { `PORTUGUESE`, `PORTUGUES`, `POR`, `PT`, },
    `BRAZILIAN_PORTUGUESE`: { `BRAZILIAN`, `BRAZIL`, `BR`, `PORTUGUESE.BR`, `PT.BR`, },
    `RUSSIAN`: { `RUSSIAN`, `RUS`, `RU`, },
    `JAPANESE`: { `JAPANESE`, `JAP`, `JPN`, `JP`, `JA`, },
    `KOREAN`: { `KOREAN`, `KOR`, `KO`, `KR`, },
    `ARABIC`: { `ARABIC`, `ARA`, `AR`, },
    `HEBREW`: { `HEBREW`, `HEB`, `HE`, },
    `THAI`: { `THAI`, `THA`, `TH`, },
    `TURKISH`: { `TURKISH`, `TUR`, `TR`, },
    `GREEK`: { `GREEK`, `GRE`, `EL`, },
    `POLISH`: { `POLISH`, `POL`, `PL`, `POLSKI`, },
    `HUNGARIAN`: { `HUNGARIAN`, `HUN`, `HU`, `MAGYAR`, },
    `CZECH`: { `CZECH`, `CZE`, `CS`, },
    `CHINESE`: { `CHINESE`, `CHI`, `ZH`, },
}

var ExtrasPatterns = []string{
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
	GetLanguagePatterns = sync.OnceValue(func() map[string][]*regexp.Regexp {
		return compilePatternMap(LanguagePatterns)
	})
	GetExtrasPatterns = sync.OnceValue(func() []*regexp.Regexp {
		return compilePatternSlice(ExtrasPatterns)
	})
)

func compilePatternMap(patterns map[string][]string) map[string][]*regexp.Regexp {
	result := make(map[string][]*regexp.Regexp, len(patterns))
	for key, patternList := range patterns {
		compiled := make([]*regexp.Regexp, len(patternList))
		for i, pattern := range patternList {
			compiled[i] = regexp.MustCompile(pattern)
			//compiled[i] = regexp.MustCompile(`^` + pattern + `$`)
		}
		result[key] = compiled
	}
	return result
}

func compilePatternSlice(patterns []string) []*regexp.Regexp {
	result := make([]*regexp.Regexp, len(patterns))
	for i, pattern := range patterns {
		result[i] = regexp.MustCompile(pattern)
	}
	return result
}
