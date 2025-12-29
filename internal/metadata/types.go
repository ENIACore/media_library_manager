package metadata

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

type Node uint8

// Classification of node that describes purpose of file or directory
const (
	MovieFile	Node = iota	
	EpisodeFile
	SubtitleFile
	BonusFile
	
	SubtitleDir
	BonusDir
	MovieDir
	SeasonDir
	SeriesDir
)


type Format uint8

// Format of file that describes type of file
const (
    Video			Format = iota
    Subtitle		
	Audio		
)
