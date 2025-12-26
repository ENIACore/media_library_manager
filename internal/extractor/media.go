// TODO: Make patterns return in order of specificity
package extractor

import (
	"log/slog"
	"regexp"
	//"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/ENIACore/media_library_manager/internal/patterns"
	//"github.com/ENIACore/media_library_manager/internal/patterns"
)

// Extracts movie or series title from file or directory
// Extracts using these expected patterns:
//		-	<title>.<year (optional)>.<resolution, codec, source, or audio>...
//		-	<title>.<year (optional)>.<season or ep>...
//		-	<title>.<year (optional)>.<file ext>
//		-	<title>.<year (optional)>
func extractTitle(segments []string, logger *slog.Logger) string {
	//log := logger.With("func", "extractTitle")

	var sb strings.Builder
	for i, segment := range segments {
		candidates := segments[i:] 

		year := parseYear(segment)
		// Valid year, skip
		if year != -1 {
			continue
		}

		if parseResolution(candidates) != "" {
			break
		} 
		if parseCodec(candidates) != "" {
			break
		}
		if parseSource(candidates) != "" {
			break
		}
		if parseAudio(candidates) != "" {
			break
		}


	}

	return sb.String()
}

// Helper function to return resolution if left most segments are a resolution or empty string if not
func parseResolution(segments []string) string {
	for _, group := range patterns.GetResolutionPatternGroups() {
		for _, re := range group.Patterns {
			if matchSegments(segments, (*regexp.Regexp)(re)) != nil {
				return group.Key
			}
		}
	}
	return ""
}

// Helper function to return codec if left most segments are a codec or empty string if not
func parseCodec(segments []string) string {
	for _, group := range patterns.GetCodecPatternGroups() {
		for _, re := range group.Patterns {
			if matchSegments(segments, (*regexp.Regexp)(re)) != nil {
				return group.Key
			}
		}
	}
	return ""
}

// Helper function to return media source if left most segments are a media source or empty string if not
func parseSource(segments []string) string {
	for _, group := range patterns.GetSourcePatternGroups() {
		for _, re := range group.Patterns {
			if matchSegments(segments, (*regexp.Regexp)(re)) != nil {
				return group.Key
			}
		}
	}
	return ""
}

// Helper function to return audio if left most segments are a audio or empty string if not
func parseAudio(segments []string) string {
	for _, group := range patterns.GetAudioPatternGroups() {
		for _, re := range group.Patterns {
			if matchSegments(segments, (*regexp.Regexp)(re)) != nil {
				return group.Key
			}
		}
	}
	return ""
}

// Returns -1 for invalid year, otherwise returns year
func parseYear(s string) int {
	if len(s) != 4 {
		return -1
	}

	year, err := strconv.Atoi(s)	
	if err != nil {
		return -1
	}

	if year < 1930 || year > time.Now().Year() {
		return -1
	}

	return year
}


// Returns -1 for SEASON pattern not matched, 0 for match without number, > 0 for season number
func parseSeason(segments []string) int {
	for _, re := range patterns.GetSeasonPatterns() {
		match := matchSegments(segments, (*regexp.Regexp)(re))

		if match == nil {
			continue
		}
		if len(match) == 1 {
			return 0	
		}

		if season, err := strconv.Atoi(match[1]); err == nil {
			return season
		}
		return 0 // matched but couldn't parse number
		
	}
	return -1
}
