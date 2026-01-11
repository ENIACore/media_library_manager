package classifier

import (
	"github.com/ENIACore/media_library_manager/internal/metadata"
)

func isSubtitleDir(entry *metadata.Entry) bool {
	// Subtitle directory cannot have nested directories
	if entry.Height() > 1 {
		return false;
	}

	// All files in subtitle directory should be of type subtitle
	for _, child := range entry.Children {
		if child.Type != metadata.Subtitle {
			return false;
		}
	}

	return true
}

func isBonusDir(entry *metadata.Entry) bool {
	return false
}

func isMovieDir(entry *metadata.Entry) bool {
	return false
}

func isSeasonDir(entry *metadata.Entry) bool {
	return false
}

func isSeriesDir(entry *metadata.Entry) bool {
	return false
}

// Structure of media torrents
/*
Movie File
Episode File
Subtitle File
Bonus File

Subtitle Directory
└── Subtitle File(s)

Bonus Directory
├── Bonus File(s)
└── Subtitle File(s) (optional)

Movie Directory
├── Movie File
├── Subtitle File (optional)
├── Bonus File (optional)
└── Bonus Directory (optional)

Season Directory
├── Episode File(s)
├── Subtitle File(s) (optional)
├── Bonus Directory (optional)
└── Subtitle Directory (optional)

Series Directory
├── Season Directory(s)
├── Bonus Directory (optional)
└── Subtitle Directory (optional)
*/
