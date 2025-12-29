package extractor

import (
	"github.com/ENIACore/media_library_manager/internal/metadata"
	"log/slog"
	"fmt"
	"path/filepath"
	"strings"
)

func Extract(path string, logger *slog.Logger) metadata.Metadata {
	log := logger.With("func", "Extract")
	log.Info("extracting path", "path", path)
	filename := filepath.Base(path)	

	sanitizedName := strings.Split(sanitizeName(filename), ".")
	media := metadata.MediaMetadata{}
	title := extractTitle(sanitizedName)
	sanitizedName = sanitizedName[len(title):]

	media.Title = strings.Join(title, ".")
	media.Year = extractYear(sanitizedName)
	media.Episode = extractEpisode(sanitizedName)
	media.Season = extractSeason(sanitizedName)
	media.Resolution = extractResolution(sanitizedName)
	media.Codec = extractCodec(sanitizedName)
	media.Source = extractSource(sanitizedName)
	media.Audio = extractAudio(sanitizedName)
	log.Debug("successfully extracted media metadata", "media-metadata", fmt.Sprintf("%+v", media))

	return metadata.Metadata{
		Media:	media,
		Path:	metadata.PathMetadata{

		},
	}
}
