package patterns

import (
	"regexp"
	"sync"
)

var VideoExtensionPatterns = []string{
	`MP4`, `MKV`, `AVI`, `MOV`, `FLV`, `WMV`, `WEBM`, `M4V`, `TS`, `M2TS`,
	`MPG`, `MPEG`, `VOB`, `3GP`, `OGV`, `RMVB`, `RM`, `DIVX`, `F4V`,
}

var SubtitleExtensionPatterns = []string{
	`SRT`, `ASS`, `SSA`, `SUB`, `VTT`, `SBV`, `JSON`, `SMI`, `LRC`,
	`PSB`, `IDX`, `USF`, `TTML`,
}

var AudioExtensionPatterns = []string{
	`MP3`, `FLAC`, `AAC`, `OGG`, `WMA`, `M4A`, `OPUS`, `WAV`,
	`APE`, `WV`, `DTS`, `AC3`, `MKA`,
}

var (
	GetVideoExtensionPatterns = sync.OnceValue(func() []*regexp.Regexp {
		return compilePatternSlice(AudioExtensionPatterns)
	})
	GetSubtitleExtensionPatterns = sync.OnceValue(func() []*regexp.Regexp {
		return compilePatternSlice(SubtitleExtensionPatterns)
	})
	GetAudioExtensionPatterns = sync.OnceValue(func() []*regexp.Regexp {
		return compilePatternSlice(AudioExtensionPatterns)
	})
)
