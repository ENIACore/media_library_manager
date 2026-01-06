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

type EntryRole int8

const Unknown = -1

// Classification of node that describes purpose of file or directory
const (
	MovieFile		EntryRole = iota	
	EpisodeFile
	SubtitleFile
	BonusFile
	
	SubtitleDir
	BonusDir
	MovieDir
	SeasonDir
	SeriesDir
)


type ContentType int8

// Type of file
const (
    Video			ContentType = iota
    Subtitle		
	//Audio			- Audio classification not yet included		
)
