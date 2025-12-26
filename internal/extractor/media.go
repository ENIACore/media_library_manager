package extractor

import (
	"fmt"
	"log/slog"
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

		_, err := parseYear(segment)
		// Valid year, skip
		if err != nil {
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
	for key, regexps := range patterns.GetResolutionPatterns() {
		for _, re := range regexps {
			if matchSegments(segments, re) != nil {
				return key
			}
		}
	}
	return ""
}

// Helper function to return codec if left most segments are a codec or empty string if not
func parseCodec(segments []string) string {
	for key, regexps := range patterns.GetCodecPatterns() {
		for _, re := range regexps {
			if matchSegments(segments, re) != nil {
				return key
			}
		}
	}
	return ""
}

// Helper function to return media source if left most segments are a media source or empty string if not
func parseSource(segments []string) string {
	for key, regexps := range patterns.GetSourcePatterns() {
		for _, re := range regexps {
			if matchSegments(segments, re) != nil {
				return key
			}
		}
	}
	return ""
}

// Helper function to return audio if left most segments are a audio or empty string if not
func parseAudio(segments []string) string {
	for key, regexps := range patterns.GetAudioPatterns() {
		for _, re := range regexps {
			if matchSegments(segments, re) != nil {
				return key
			}
		}
	}
	return ""
}

func parseYear(s string) (int, error) {
	if len(s) != 4 {
		return 0, fmt.Errorf("expected year with len 4, got len %v", len(s)) 
	}

	year, err := strconv.Atoi(s)	
	if err != nil {
		return 0, fmt.Errorf("invalid year, %w", err)
	}

	if year < 1930 || year > time.Now().Year() {
		return 0, fmt.Errorf("invalid year, value %v is less than 1930 and greater than current year", year)	
	}

	return year, nil
}
