package extractor

import (
	"github.com/ENIACore/media_library_manager/internal/patterns"
	"github.com/ENIACore/media_library_manager/internal/metadata"
	"regexp"
)

func ParsePath(path string) metadata.PathMetadata {
	return metadata.PathMetadata{}
}


// Extracts extension and filetype from filename parts, returns "" if extension not present
// Enables extensions with multiple '.' seperators
/*
func extractFormat(segments []string) (string, metadata.FormatType) {
	for i := range segments {
		candidates := segments[i:]
		if ext := parseVideoExt(candidates); ext != "" {
			return ext, metadata.Video
		}
		if ext := parseSubtitleExt(candidates); ext != "" {
			return ext, metadata.Subtitle
		}
		if ext := parseAudioExt(candidates); ext != "" {
			return ext, metadata.Audio
		}
	}

	return "", metadata.UnknownFormat
}
*/

func parseVideoExt(segments []string) string {
	for _, re := range patterns.GetVideoExtensionPatterns() {
		match := matchSegments(segments, (*regexp.Regexp)(re))
		if match != nil {
			return match[0]
		}
	}
	return ""
}

func parseSubtitleExt(segments []string) string {
	for _, re := range patterns.GetSubtitleExtensionPatterns() {
		match := matchSegments(segments, (*regexp.Regexp)(re))
		if match != nil {
			return match[0]
		}
	}
	return ""
}

func parseAudioExt(segments []string) string {
	for _, re := range patterns.GetAudioExtensionPatterns() {
		match := matchSegments(segments, (*regexp.Regexp)(re))
		if match != nil {
			return match[0]
		}
	}
	return ""
}
