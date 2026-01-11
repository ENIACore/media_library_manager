package patterns

import (
	"regexp"
	"sync"
)

type Pattern string

type CompiledPattern regexp.Regexp

type PatternGroup struct {
	Key      string
	Patterns []Pattern
}

type CompiledPatternGroup struct {
	Key      string
	Patterns []*CompiledPattern
}

// Un-useful metadata of file pertaining to typical patterns found in torrent files
var MiscPatterns = []Pattern{
	// === UNUSED QUALITY INDICATORS ===
	// HDR variants
	`HDR`, `HDR10`, `HDR10PLUS`, `HDR10\+`, `DOLBY\.VISION`, `DOLBYVISION`, `DV`, `HLG`,
	// Bit depth
	`10BIT`, `10\.BIT`, `8BIT`, `8\.BIT`, `12BIT`, `12\.BIT`,
	// Color
	`SDR`,

	// === EDITION / VERSION ===
	`REMASTERED`, `REMASTER`,
	`EXTENDED`, `EXTENDED\.CUT`, `EXTENDED\.EDITION`,
	`UNRATED`,
	`UNCUT`,
	`DIRECTORS\.CUT`, `DC`,
	`THEATRICAL`, `THEATRICAL\.CUT`,
	`CRITERION`, `CC`,
	`SPECIAL\.EDITION`, `SE`,
	`ANNIVERSARY`, `ANNIVERSARY\.EDITION`,
	`COLLECTORS`, `COLLECTORS\.EDITION`, `CE`,
	`LIMITED`, `LIMITED\.EDITION`,
	`PROPER`,
	`REPACK`, `RERIP`,
	`REAL`,
	`RETAIL`,
	`FINAL\.CUT`,
	`IMAX`,
	`OPEN\.MATTE`, `OPENMATTE`,
	`3D`, `HSBS`, `HOU`, `HALF\.SBS`, `FULL\.SBS`,

	// === RELEASE INFO ===
	`INTERNAL`, `INT`,
	`NFO`, `NFOFIX`,
	`SAMPLE`,
	`PROOF`,
	`READNFO`, `READ\.NFO`,
	`DIRFIX`,
	`NFOFIX`,
	`SYNCFIX`,
	`SAMPLEFIX`,
	`SUBBED`, `SUBS`, `SUB`,
	`DUBBED`, `DUB`,
	`HARDCODED`, `HC`,
	`MULTISUBS`, `MULTI\.SUBS`, `MULTISUB`,
	`MULTI`, `MULTILANG`, `MULTI\.LANG`, `MULTi`,

	// === SCENE / P2P RELEASE GROUPS ===
	`YIFY`, `YTS`, `YTS\.MX`, `YTS\.AM`, `YTS\.LT`, `YTS\.AG`,
	`RARBG`,
	`ETRG`, `ETTV`, `ETHD`,
	`PSA`, `PSARIPS`,
	`GALAXYRG`, `GALAXY\.RG`, `GALAXYTV`,
	`SPARKS`,
	`GECKOS`,
	`AMIABLE`,
	`DRONES`,
	`FGT`,
	`EVO`,
	`CMRG`,
	`TIGOLE`, `QXRTRIGOLE`,
	`FLUX`,
	`NTG`,
	`EPSILON`,
	`PLAYNOW`,
	`HDETG`,
	`DIMENSION`,
	`LOL`,
	`KILLERS`,
	`AVS`,
	`SVA`,
	`FLEET`,
	`NTEB`,
	`PAHE`, `PAHE\.IN`, `PAHE\.PH`,
	`MKVKING`,
	`ION10`,
	`AMZN`,
	`NF`,
	`HULU`,
	`DSNP`,
	`ATVP`,
	`PCOK`,
	`HMAX`, `HBO`,
	`MAX`,
	`PMTP`,
	`CRAV`,
	`SHO`,
	`STAN`,
	`CRITERION`,

	// === TV SPECIFIC ===
	`COMPLETE`, `COMPLETE\.SERIES`,
	`MINISERIES`, `MINI\.SERIES`,
	`PILOT`,
	`FINALE`,
	`HDTV`,

	// === MISC NOISE ===
	`XXX`,
	`HINDI\.DUBBED`, `TAMIL\.DUBBED`, `TELUGU\.DUBBED`,
	`CONVERT`,
	`COLORIZED`,
	`RESTORED`,
	`AI\.UPSCALE`, `UPSCALED`, `AI\.ENHANCED`,
}

var LanguagePatternGroups = []PatternGroup{
	{Key: `ENGLISH`, Patterns: []Pattern{`ENGLISH`, `ENG`, `EN`}},
	{Key: `SPANISH`, Patterns: []Pattern{`SPANISH`, `CASTELLANO`, `SPA`, `ES`, `ESPAÑOL`}},
	{Key: `FRENCH`, Patterns: []Pattern{`FRENCH`, `FRA`, `FR`}},
	{Key: `GERMAN`, Patterns: []Pattern{`GERMAN`, `DEUTSCH`, `GER`, `DE`, `GERMAN`}},
	{Key: `ITALIAN`, Patterns: []Pattern{`ITALIAN`, `ITA`, `ITALIANO`}},
	{Key: `PORTUGUESE`, Patterns: []Pattern{`PORTUGUESE`, `PORTUGUES`, `POR`, `PT`}},
	{Key: `BRAZILIAN_PORTUGUESE`, Patterns: []Pattern{`BRAZILIAN`, `BRAZIL`, `BR`, `PORTUGUESE.BR`, `PT.BR`}},
	{Key: `RUSSIAN`, Patterns: []Pattern{`RUSSIAN`, `RUS`, `RU`}},
	{Key: `JAPANESE`, Patterns: []Pattern{`JAPANESE`, `JAP`, `JPN`, `JP`, `JA`}},
	{Key: `KOREAN`, Patterns: []Pattern{`KOREAN`, `KOR`, `KO`, `KR`}},
	{Key: `ARABIC`, Patterns: []Pattern{`ARABIC`, `ARA`, `AR`}},
	{Key: `HEBREW`, Patterns: []Pattern{`HEBREW`, `HEB`, `HE`}},
	{Key: `THAI`, Patterns: []Pattern{`THAI`, `THA`, `TH`}},
	{Key: `TURKISH`, Patterns: []Pattern{`TURKISH`, `TUR`, `TR`}},
	{Key: `GREEK`, Patterns: []Pattern{`GREEK`, `GRE`, `EL`}},
	{Key: `POLISH`, Patterns: []Pattern{`POLISH`, `POL`, `PL`, `POLSKI`}},
	{Key: `HUNGARIAN`, Patterns: []Pattern{`HUNGARIAN`, `HUN`, `HU`, `MAGYAR`}},
	{Key: `CZECH`, Patterns: []Pattern{`CZECH`, `CZE`, `CS`}},
	{Key: `CHINESE`, Patterns: []Pattern{`CHINESE`, `CHI`, `ZH`}},
}

var BonusPatternGroups = []PatternGroup{
	// Behind the scenes / Making of
	{Key: `BEHIND_THE_SCENES`, Patterns: []Pattern{
		`BEHIND\.THE\.SCENE[S]?`,
		`BTS`,
		`MAKING\.OF`,
		`MAKING`,
		`THE\.MAKING\.OF`,
	}},
	// Deleted / Extended scenes
	{Key: `DELETED_SCENE`, Patterns: []Pattern{
		`DELETED\.SCENE[S]?`,
		`DELETED`,
		`EXTENDED\.SCENE[S]?`,
		`ALTERNATE\.SCENE[S]?`,
		`ADDITIONAL\.SCENE[S]?`,
	}},
	// Featurettes
	{Key: `FEATURETTE`, Patterns: []Pattern{
		`FEATURETTE[S]?`,
		`FEATURE[S]?`,
		`SHORT[S]?`,
		`MINI\.FEATURE[S]?`,
	}},
	// Interviews
	{Key: `INTERVIEW`, Patterns: []Pattern{
		`INTERVIEW[S]?`,
		`CAST\.INTERVIEW[S]?`,
		`Q\.?AND\.?A`,
		`QA`,
	}},
	// Bloopers / Gag reel
	{Key: `BLOOPER`, Patterns: []Pattern{
		`BLOOPER[S]?`,
		`GAG\.?REEL[S]?`,
		`OUTTAKE[S]?`,
	}},
	// Trailers / Promos
	{Key: `TRAILER`, Patterns: []Pattern{
		`TRAILER[S]?`,
		`TEASER[S]?`,
		`PROMO[S]?`,
		`TV\.SPOT[S]?`,
		`TVSPOT[S]?`,
	}},
	// Commentary
	{Key: `COMMENTARY`, Patterns: []Pattern{
		`COMMENTARY`,
		`AUDIO\.COMMENTARY`,
		`DIRECTOR[S]?\.COMMENTARY`,
	}},
	// Documentary
	{Key: `DOCUMENTARY`, Patterns: []Pattern{
		`DOCUMENTARY`,
		`DOCUMENTARIES`,
		`DOC[S]?`,
	}},
	// General extras / bonus
	{Key: `EXTRA`, Patterns: []Pattern{
		`EXTRA[S]?`,
		`BONUS`,
		`BONUS\.CONTENT`,
		`BONUS\.MATERIAL[S]?`,
		`SPECIAL\.FEATURE[S]?`,
		`SUPPLEMENTAL[S]?`,
		`SUPPLEMENT[S]?`,
	}},
}

var (
	GetLanguagePatternGroups = sync.OnceValue(func() []CompiledPatternGroup {
		return compilePatternGroups(LanguagePatternGroups)
	})
	GetBonusPatternGroups = sync.OnceValue(func() []CompiledPatternGroup {
		return compilePatternGroups(BonusPatternGroups)
	})
	GetMiscPatterns = sync.OnceValue(func() []*CompiledPattern {
		return compilePatterns(MiscPatterns)
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
			Key:      group.Key,
			Patterns: patterns,
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
