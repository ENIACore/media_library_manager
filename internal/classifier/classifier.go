package classifier

import (
	"github.com/ENIACore/media_library_manager/internal/metadata"
)

func isMovieFile(entry *metadata.Entry) bool {
	return false
}

func isEpisodeFile(entry *metadata.Entry) bool {
	return false
}

func isSubtitleFile(entry *metadata.Entry) bool {
	return false
}

func isBonusFile(entry *metadata.Entry) bool {
	return false
}

/*
func isSubtitleDir(entry *metadata.Entry) bool {
	return false
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
*/

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
