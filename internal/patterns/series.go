package patterns

import (
	"sync"
)

var SeasonPatterns = []Pattern{
	// Matches S<number>
	`S(\d+)`,
	// Matches S.<number>
	`S\.(\d+)`,
	// Matches SEA<number>
	`SEA(\d+)`,
	// Matches SEA.<number>
	`SEA\.(\d+)`,
	// Matches SEASON<number>
	`SEASON(\d+)`,
	// Matches SEASON.<number>
	`SEASON\.(\d+)`,
	// Matches SEASON
	`SEASON`,
	// Matches S<number>E<number>
	`S(\d+)E\d+`,
}

var EpisodePatterns = []Pattern{
	// Matches E<number>
	`E(\d+)`,
	// Matches E.<number>
	`E\.(\d+)`,
	// Matches EP<number>
	`EP(\d+)`,
	// Matches EP.<number>
	`EP\.(\d+)`,
	// Matches EPISODE<number>
	`EPISODE(\d+)`,
	// Matches EPISODE.<number>
	`EPISODE\.(\d+)`,
	// Matches EPISODE
	`EPISODE`,
	// Matches EP
	`EP`,
	// Matches S<number>E<number>
	`S\d+E(\d+)`,
	// Matches <number>X<number>
	`\d+X(\d+)`,
	// Matches <number>.X.<number> (e.g., 1.X.01, 2.X.15)
	`\d+\.X\.(\d+)`,
}

var (
	GetSeasonPatterns = sync.OnceValue(func() []*CompiledPattern {
		return compilePatterns(SeasonPatterns)
	})
	GetEpisodePatterns = sync.OnceValue(func() []*CompiledPattern {
		return compilePatterns(EpisodePatterns)
	})
)
