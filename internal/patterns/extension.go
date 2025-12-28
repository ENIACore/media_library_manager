package patterns

import (
	"sync"
)

var VideoExtensionPatterns = []Pattern{
	`MP4`, `MKV`, `AVI`, `MOV`, `FLV`, `WMV`, `WEBM`, `M4V`, `TS`, `M2TS`,
	`MPG`, `MPEG`, `VOB`, `3GP`, `OGV`, `RMVB`, `RM`, `DIVX`, `F4V`,
}

var SubtitleExtensionPatterns = []Pattern{
	`SRT`, `ASS`, `SSA`, `SUB`, `VTT`, `SBV`, `JSON`, `SMI`, `LRC`,
	`PSB`, `IDX`, `USF`, `TTML`,
}

var AudioExtensionPatterns = []Pattern{
	`MP3`, `FLAC`, `AAC`, `OGG`, `WMA`, `M4A`, `OPUS`, `WAV`,
	`APE`, `WV`, `DTS`, `AC3`, `MKA`,
}

var (
	GetVideoExtensionPatterns = sync.OnceValue(func() []*CompiledPattern {
		return compilePatterns(VideoExtensionPatterns)
	})
	GetSubtitleExtensionPatterns = sync.OnceValue(func() []*CompiledPattern {
		return compilePatterns(SubtitleExtensionPatterns)
	})
	GetAudioExtensionPatterns = sync.OnceValue(func() []*CompiledPattern {
		return compilePatterns(AudioExtensionPatterns)
	})
)
