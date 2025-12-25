package extractor

import (
	"github.com/ENIACore/media_library_manager/internal/metadata"
)

func Extract(path string) metadata.Metadata {
	return metadata.Metadata{
		Media:		extractMedia(path),
		Path:		extractPath(path),
	}
}
