package classifier

import (
	"testing"

	"github.com/ENIACore/media_library_manager/internal/metadata"
	"github.com/ENIACore/media_library_manager/internal/parser"
)


func TestIsSubtitleDir(t *testing.T) {
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
	movieFile := parser.Entry{
		Parent: nil,
		Children: nil,
		MediaInfo: metadata.MediaInfo{
			Title:		[]string{
				"TEST",
				"MOVIE",
			},
			Year:		2025,
			Episode:	-1,
			Season:		-1,
			Resolution:	"",
			Codec:		"",
			Source:		"",
			Audio:		"",
			Language:	"",
		},
		PathInfo: metadata.PathInfo{
			Dest: "",
			Source: "/dir/test movie 2025.mp4",
			Ext: "MP4",	
			Type: metadata.Video,
		},
	}
	subtitleFile := parser.Entry{
		Parent: nil,
		Children: nil,
		MediaInfo: metadata.MediaInfo{
			Title:		[]string{
				"TEST",
				"MOVIE",
			},
			Year:		-1,
			Episode:	-1,
			Season:		-1,
			Resolution:	"",
			Codec:		"",
			Source:		"",
			Audio:		"",
			Language:	"",
		},
		PathInfo: metadata.PathInfo{
			Dest: "",
			Source: "/dir/subtitles/subtitle english.srt",
			Ext: "SRT",	
			Type: metadata.Subtitle,
		},
	}

	validSubtitleDir := parser.Entry{
		Parent: nil,
		Children: []*parser.Entry{
			&subtitleFile,
			&subtitleFile,
			&subtitleFile,
		},	
		PathInfo: metadata.PathInfo{
			Dest: "",
			Source: "/dir/subtitles",
			Ext: "",	
			Type: metadata.Unknown,
		},
	}

	tests := []struct{
		name		string
		node		parser.Entry
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
		node		parser.Entry
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
		node		parser.Entry
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
		node		parser.Entry
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
		node		parser.Entry
	}{

	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
		})
	}
}
