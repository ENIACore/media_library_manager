package classifier

import (
	"github.com/ENIACore/media_library_manager/internal/parser"
)


func isSubtitleDir(entry *parser.Entry) bool {
	return false
}

func isBonusDir(entry *parser.Entry) bool {
	return false
}

func isMovieDir(entry *parser.Entry) bool {
	return false
}

func isSeasonDir(entry *parser.Entry) bool {
	return false
}

func isSeriesDir(entry *parser.Entry) bool {
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
