package extractor

import (
	"github.com/ENIACore/media_library_manager/internal/metadata"
	"log/slog"
)

func Extract(path string, logger *slog.Logger) metadata.Metadata {
	return metadata.Metadata{
		Media:		extractMedia(path),
		Path:		extractPath(path),
	}
}
