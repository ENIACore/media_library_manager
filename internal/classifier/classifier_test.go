package classifier

import (
	"testing"

	"github.com/ENIACore/media_library_manager/internal/metadata"
)

func intPtr(i int) *int {
	return &i
}

/*
var movieFile = metadata.Entry{
	Parent: nil,
	Children: nil,
	MediaInfo: metadata.MediaInfo{
		Title:		[]string{
			"TEST",
			"MOVIE",
		},
		Year:		intPtr(2025),
		Episode:	nil,
		Season:		nil,
		Resolution:	"1080p",
		Codec:		"x264",
		Source:		"REMUX",
		Audio:		"Atmos",
		Language:	"ENGLISH",
	},
	PathInfo: metadata.PathInfo{
		Dest: "",
		Source: "/test movie 2025 1080p.x264.remux.atmos.english.mp4",
		Ext: "MP4",	
		Type: metadata.Video,
	},
}
var episodeFile = metadata.Entry{
	Parent: nil,
	Children: nil,
	MediaInfo: metadata.MediaInfo{
		Title:		[]string{
			"TEST",
			"EPISODE",
		},
		Year:		intPtr(2025),
		Episode:	intPtr(1),
		Season:		intPtr(1),
		Resolution:	"1080p",
		Codec:		"x264",
		Source:		"REMUX",
		Audio:		"Atmos",
		Language:	"ENGLISH",
	},
	PathInfo: metadata.PathInfo{
		Dest: "",
		Source: "/test episode S1E1 2025 1080p.x264.remux.atmos.english.mp4",
		Ext: "MP4",	
		Type: metadata.Video,
	},
}
var bonusFile = metadata.Entry{
	Parent: nil,
	Children: nil,
	MediaInfo: metadata.MediaInfo{
		Title:		[]string{
			"TEST",
			"BONUS",
		},
		Year:		intPtr(2025),
		Episode:	nil,
		Season:		nil,
		Resolution:	"1080p",
		Codec:		"x264",
		Source:		"REMUX",
		Audio:		"Atmos",
		Language:	"ENGLISH",
	},
	PathInfo: metadata.PathInfo{
		Dest: "",
		Source: "/test bonus 2025 1080p.x264.remux.atmos.english.mp4",
		Ext: "MP4",	
		Type: metadata.Video,
	},
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
func TestIsSubtitleDir(t *testing.T) {
	subtitleFile := metadata.Entry{
		Parent: nil,
		Children: nil,
		MediaInfo: metadata.MediaInfo{
			Title:		[]string{
				"SUBTITLE",
			},
			Year:		nil,
			Episode:	nil,
			Season:		nil,
			Resolution:	"",
			Codec:		"",
			Source:		"",
			Audio:		"",
			Language:	"ENGLISH",
		},
		PathInfo: metadata.PathInfo{
			Dest: "",
			Source: "/subtitle english.srt",
			Ext: "SRT",	
			Type: metadata.Subtitle,
		},
	}

	validSubtitleDir := metadata.Entry{
		Parent: nil,
		Children: []*metadata.Entry{
			&subtitleFile,
			&subtitleFile,
			&subtitleFile,
		},	
		PathInfo: metadata.PathInfo{
			Dest: "",
			Source: "/subtitles",
			Ext: "",	
			Type: metadata.Unknown,
		},
	}

	tests := []struct{
		name		string
		node		metadata.Entry
	}{
		{
			name:		"valid subtitle directory", node:		validSubtitleDir,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
		})
	}
}

func TestIsBonusDir(t *testing.T) {
	tests := []struct{
		name		string
		node		metadata.Entry
	}{

	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
		})
	}
}

func TestIsMovieDir(t *testing.T) {
	tests := []struct{
		name		string
		node		metadata.Entry
	}{

	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
		})
	}
}

func TestIsSeasonDir(t *testing.T) {
	tests := []struct{
		name		string
		node		metadata.Entry
	}{

	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
		})
	}
}

func TestIsSeriesDir(t *testing.T) {
	tests := []struct{
		name		string
		node		metadata.Entry
	}{

	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
		})
	}
}
