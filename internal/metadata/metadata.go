package metadata

import (
	"github.com/ENIACore/media_library_manager/internal/config"
)


type MediaMetadata struct {
	name		string
	year		int
	episode		int
	season		int

	resolution	string
	codec		string
	source		string
	audio		string

	language	string
}

type FileMetadata struct {
	newPath			string
	originalPath	string
	ext				string
	format			config.FormatType
}
