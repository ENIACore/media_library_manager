package classifier

import (
	"testing"

	"github.com/ENIACore/media_library_manager/internal/metadata"
	"github.com/ENIACore/media_library_manager/internal/parser"
)

func intPtr(i int) *int {
	return &i
}

var movieFile = parser.Entry{
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
		Source: "/dir/test movie 2025 1080p.x264.remux.atmos.english.mp4",
		Ext: "MP4",	
		Type: metadata.Video,
	},
}
var episodeFile = parser.Entry{
	Parent: nil,
	Children: nil,
	MediaInfo: metadata.MediaInfo{
		Title:		[]string{
			"TEST",
			"EPISODE",
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
		Source: "/dir/test movie 2025 1080p.x264.remux.atmos.english.mp4",
		Ext: "MP4",	
		Type: metadata.Video,
	},
}
var subtitleFile = parser.Entry{
	Parent: nil,
	Children: nil,
	MediaInfo: metadata.MediaInfo{
		Title:		[]string{
			"TEST",
			"MOVIE",
		},
		Year:		nil,
		Episode:	nil,
		Season:		nil,
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


func TestIsMovieFile(t *testing.T) {
	tests := []struct{
		name		string
		entry		*parser.Entry
		expected	bool
	}{
		{
			name:		"movie file",
			entry:		nil,
			expected:	false,
		},
		{
			name:		"episode file",
			entry:		nil,
			expected:	false,
		},
		{
			name:		"subtitle file",
			entry:		nil,
			expected:	false,
		},
		{
			name:		"bonus file",
			entry:		nil,
			expected:	false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

		})
	}
}

func TestIsEpisodeFile(t *testing.T) {
	tests := []struct{
		name		string
		entry		*parser.Entry
		expected	bool
	}{
		{
			name:		"movie file",
			entry:		nil,
			expected:	false,
		},
		{
			name:		"episode file",
			entry:		nil,
			expected:	false,
		},
		{
			name:		"subtitle file",
			entry:		nil,
			expected:	false,
		},
		{
			name:		"bonus file",
			entry:		nil,
			expected:	false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

		})
	}
}

func TestIsSubtitleFile(t *testing.T) {
	tests := []struct{
		name		string
		entry		*parser.Entry
		expected	bool
	}{
		{
			name:		"movie file",
			entry:		nil,
			expected:	false,
		},
		{
			name:		"episode file",
			entry:		nil,
			expected:	false,
		},
		{
			name:		"subtitle file",
			entry:		nil,
			expected:	false,
		},
		{
			name:		"bonus file",
			entry:		nil,
			expected:	false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

		})
	}
}

func TestIsBonusFile(t *testing.T) {
	tests := []struct{
		name		string
		entry		*parser.Entry
		expected	bool
	}{
		{
			name:		"movie file",
			entry:		nil,
			expected:	false,
		},
		{
			name:		"episode file",
			entry:		nil,
			expected:	false,
		},
		{
			name:		"subtitle file",
			entry:		nil,
			expected:	false,
		},
		{
			name:		"bonus file",
			entry:		nil,
			expected:	false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

		})
	}
}


/*
func TestIsSubtitleDir(t *testing.T) {

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
*/
