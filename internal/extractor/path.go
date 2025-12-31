package extractor

import (
	"github.com/ENIACore/media_library_manager/internal/patterns"
	"github.com/ENIACore/media_library_manager/internal/metadata"
	"log/slog"
	"regexp"
)

func ExtractPath(path string, logger *slog.Logger) metadata.PathInfo {
	//log := logger.With("func", "ExtractPath")

	return metadata.PathInfo{}
}

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
