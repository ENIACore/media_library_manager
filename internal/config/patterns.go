// Package torrent provides constants for parsing information about a torrent file's quality and contents.
//
// Attributes:
//   - VideoExtensions: Set of file extensions for video files
//   - SubtitleExtensions: Set of file extensions for subtitle files
//   - AudioExtensions: Set of file extensions for audio files
//   - ResolutionPatterns: Map of common patterns used to describe video resolution in file names
//   - CodecPatterns: Map of common patterns used to describe video codec in file names
//   - SourcePatterns: Map of common patterns used to describe video quality & source in file names
//   - AudioPatterns: Map of common patterns used to describe audio in file names
//   - SeasonsPatterns: List of patterns to match season identifiers with capture groups
//   - EpisodesPatterns: List of patterns to match episode identifiers with capture groups
//   - ExtrasPatterns: List of patterns to match extra content
package config

import (
	"regexp"
	"sync"
)

// File extensions
var VideoExtensions = map[string]struct{}{
	"MP4":  {}, "MKV": {}, "AVI": {}, "MOV": {}, "FLV": {}, "WMV": {}, "WEBM": {}, "M4V": {}, "TS": {}, "M2TS": {},
	"MPG":  {}, "MPEG": {}, "VOB": {}, "3GP": {}, "OGV": {}, "RMVB": {}, "RM": {}, "DIVX": {}, "F4V": {},
}

var SubtitleExtensions = map[string]struct{}{
	"SRT": {}, "ASS": {}, "SSA": {}, "SUB": {}, "VTT": {}, "SBV": {}, "JSON": {}, "SMI": {}, "LRC": {},
	"PSB": {}, "IDX": {}, "USF": {}, "TTML": {},
}

var AudioExtensions = map[string]struct{}{
	"MP3": {}, "FLAC": {}, "AAC": {}, "OGG": {}, "WMA": {}, "M4A": {}, "OPUS": {}, "WAV": {},
	"APE": {}, "WV": {}, "DTS": {}, "AC3": {}, "MKA": {},
}

// ResolutionPatterns maps resolution names to regex patterns (using regexp.MatchString with fullmatch behavior)
var ResolutionPatterns = map[string][]string{
	// Matches 8k, 4320, 4320P, 4320I, 7680X4320, FULLUHD
	"8K": {`8K`, `4320[PI]?`, `7680X4320`, `FULLUHD`},
	// Matches 4k, UHD, 2160, 2160P, 2160I, 3840X2160
	"4K": {`4K`, `UHD`, `2160[PI]?`, `3840X2160`},
	// Matches 2K, 1440, 1440P, 1440I, 2560X1440, QHD, WQHD
	"2K": {`2K`, `1440[PI]?`, `2560X1440`, `QHD`, `WQHD`},
	// Matches 1080, 1080P, 1080I, FHD, 1920X1080, FULLHD
	"1080p": {`1080[PI]?`, `FHD`, `1920X1080`, `FULLHD`},
	// Matches 720, 720P, 720I, 1280X720
	"720p": {`720[PI]?`, `1280X720`},
	// Matches 576, 576P, 576I, PAL
	"576p": {`576[PI]?`, `PAL`},
	// Matches 480, 480P, 480I, NTSC
	"480p": {`480[PI]?`, `NTSC`},
	// Matches 360, 360P, 360I
	"360p": {`360[PI]?`},
	// Matches 240, 240P, 240I
	"240p": {`240[PI]?`},
}

// CodecPatterns maps codec names to regex patterns (using regexp.MatchString with fullmatch behavior)
var CodecPatterns = map[string][]string{
	// Matches AV1, SVT-AV1, SVTAV1
	"AV1": {`AV1`, `SVT\.AV1`, `SVTAV1`, `AOV1`},
	// Matches VP9
	"VP9": {`VP9`},
	// Matches VP8
	"VP8": {`VP8`},
	// Matches x265, X265, X.265, H265, H.265, HEVC, HEVC10, HEVC10BIT
	"x265": {`X265`, `X\.265`, `H265`, `H\.265`, `HEVC`, `HEVC10`, `HEVC10BIT`, `H265P`},
	// Matches x264, X264, X.264, H264, H.264, AVC, AVC1
	"x264": {`X264`, `X\.264`, `H264`, `H\.264`, `AVC`, `AVC1`, `H264P`},
	// Matches x263, X263, X.263, H263, H.263
	"x263": {`X263`, `X\.263`, `H263`, `H\.263`},
	// Matches XVID, XVID-AF
	"XVID": {`XVID`, `XVID\.AF`},
	// Matches DIVX, DIV3, DIVX6
	"DIVX": {`DIVX`, `DIV3`, `DIVX6`},
	// Matches MPEG-4, MPEG4, MP4V
	"MPEG4": {`MPEG\.4`, `MPEG4`, `MP4V`},
	// Matches MPEG-2, MPEG2, MP2V
	"MPEG2": {`MPEG\.2`, `MPEG2`, `MP2V`},
	// Matches MPEG-1, MPEG1, MP1V
	"MPEG1": {`MPEG\.1`, `MPEG1`, `MP1V`},
	// Matches VC-1, VC1, WMV3, WVC1
	"VC1": {`VC\.1`, `VC1`, `WMV3`, `WVC1`},
	// Matches THEORA
	"THEORA": {`THEORA`},
	// Matches PRORES, PRORES422, PRORES4444
	"PRORES": {`PRORES`, `PRORES422`, `PRORES4444`, `PRORES422HQ`},
	// Matches DNxHD, DNXHD, DNxHR, DNXHR
	"DNxHD": {`DNXHD`, `DNXHR`},
}

// SourcePatterns maps source/quality names to regex patterns (using regexp.MatchString with fullmatch behavior)
var SourcePatterns = map[string][]string{
	// Matches REMUX
	"REMUX": {`REMUX`},
	// Matches BluRay variants
	"BluRay": {`BLURAY`, `BDRIP`, `BD\.RIP`, `BR\.RIP`, `BRRIP`, `BDMV`, `BDISO`, `BD25`, `BD50`, `BD66`, `BD100`},
	// Matches WEB-DL, WEBDL, WEB DL
	"WEB-DL": {`WEB\.DL`, `WEBDL`},
	// Matches WEBRip, WEB-RIP, WEBRIP, WEB RIP
	"WEBRip": {`WEBRIP`, `WEB-RIP`, `WEB\.RIP`},
	// Matches WEB (but not WEB-DL or WEBRip)
	"WEB": {`WEB`},
	// Matches HDRip, HD-RIP, HDRIP, HD RIP
	"HDRip": {`HDRIP`, `HD\.RIP`},
	// Matches DVDRip, DVD-RIP, DVDRIP, DVD RIP
	"DVDRip": {`DVDRIP`, `DVD\.RIP`},
	// Matches DVD, DVDSCR, DVD5, DVD9
	"DVD": {`DVD`, `DVDSCR`, `DVD5`, `DVD9`},
	// Matches HDTV, HDTVRIP, DTTV, PDTV, SDTV, LDTV
	"HDTV": {`HDTV`, `HDTVRIP`, `DTTV`, `PDTV`, `SDTV`, `LDTV`},
	// Matches TELECINE, TC
	"TELECINE": {`TELECINE`, `TC`},
	// Matches TELESYNC, TS
	"TELESYNC": {`TELESYNC`, `TS`},
	// Matches SCREENER, SCR, DVDSCR, BDSCR
	"SCREENER": {`SCREENER`, `SCR`, `DVDSCR`, `BDSCR`},
	// Matches CAMRIP, CAM, HDCAM
	"CAM": {`CAMRIP`, `CAM`, `HDCAM`},
	// Matches WORKPRINT, WP
	"WORKPRINT": {`WORKPRINT`, `WP`},
	// Matches PPV, PPVRIP
	"PPV": {`PPV`, `PPVRIP`},
	// Matches VODRIP, VOD
	"VODRip": {`VODRIP`, `VOD`},
	// Matches HC, HCHDCAM
	"HC": {`HC`, `HCHDCAM`},
	// Matches LINE
	"LINE": {`LINE`},
	// Matches HDTS, HD-TS, HDTS
	"HDTS": {`HDTS`, `HD\.TS`},
	// Matches HDTC, HD-TC, HDTC
	"HDTC": {`HDTC`, `HD\.TC`},
	// Matches TVRIP, SATRIP, DTTVRIP
	"TVRip": {`TVRIP`, `SATRIP`, `DTTVRIP`},
}

// AudioPatterns maps audio codec/channel names to regex patterns (using regexp.MatchString with fullmatch behavior)
var AudioPatterns = map[string][]string{
	// Matches ATMOS, DOLBY-ATMOS, DOLBY ATMOS, DOLBYATMOS
	"Atmos": {`ATMOS`, `DOLBY-ATMOS`, `DOLBY\.ATMOS`, `DOLBYATMOS`},
	// Matches DTS-X, DTSX, DTS X
	"DTS-X": {`DTSX`, `DTS\.X`},
	// Matches DTS-HD-MA, DTS-HD, DTSHD-MA, DTSHD
	"DTS-HD": {`DTS\.HD\.MA`, `DTSHD-MA`, `DTSHD\.MA`, `DTS\.HD`, `DTSHD`},
	// Matches DTS-MA, DTSMA
	"DTS-MA": {`DTS\.MA`, `DTSMA`},
	// Matches DTS-ES, DTSES, DTS ES
	"DTS-ES": {`DTS\.ES`, `DTSES`},
	// Matches DTS (standalone, not part of other DTS variants)
	"DTS": {`DTS`},
	// Matches TrueHD, TRUE-HD, TRUEHD, TRUE HD
	"TrueHD": {`TRUEHD`, `TRUE\.HD`},
	// Matches DD+, DDP, E-AC-3, EAC3, DD-PLUS, DDPLUS
	"DD+": {`DDP`, `E\.AC\.3`, `EAC3`, `DD\.PLUS`, `DDPLUS`},
	// Matches DD, AC3, DOLBY-DIGITAL, DOLBY DIGITAL (but not DD+ or variants)
	"DD": {`DD`, `AC3`, `DOLBY-DIGITAL`, `DOLBY\.DIGITAL`, `DOLBYDIGITAL`},
	// Matches AAC, HE-AAC, HEAAC, HE AAC
	"AAC": {`AAC`, `HE\.AAC`, `HEAAC`},
	// Matches FLAC
	"FLAC": {`FLAC`},
	// Matches MP3
	"MP3": {`MP3`},
	// Matches LPCM, PCM
	"LPCM": {`LPCM`, `PCM`},
	// Matches OGG, VORBIS
	"OGG": {`OGG`, `VORBIS`},
	// Matches OPUS
	"OPUS": {`OPUS`},
	// Matches 5.1, 5-1, 51, 6CH
	"5.1": {`5\.1`, `51`, `6CH`},
	// Matches 7.1, 7-1, 71, 8CH
	"7.1": {`7\.1`, `71`, `8CH`},
	// Matches 2.0, 2-0, 20, STEREO, 2CH
	"2.0": {`2\.0`, `20`, `STEREO`, `2CH`},
	// Matches DUAL-AUDIO, DUAL AUDIO, DUAL
	"DUAL": {`DUAL\.AUDIO`, `DUAL`},
}

// SeasonsPatterns contains regex patterns to match season identifiers.
// If season number exists, it is capture group 1 for all matches.
var SeasonPatterns = []string{
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

// EpisodesPatterns contains regex patterns to match episode identifiers.
// If episode number exists, it is capture group 1 for all matches.
var EpisodePatterns = []string{
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

// ExtrasPatterns contains patterns for matching extra content
var ExtrasPatterns = []string{
	`EXTRA[S]?`,
	`FEATURETTE[S]?`,
	`BEHIND\.THE\.SCENE[S]?`,
	`BTS`,
	`DELETED\.SCENE[S]?`,
	`MAKING\.OF`,
	`TRAILER`,
	`BONUS`,
	`DOCUMENTARY`,
	`DOCUMENTARIES`,
}

var (
	GetResolutionPatterns = sync.OnceValue(func() map[string][]*regexp.Regexp {
		return compilePatternMap(ResolutionPatterns)
	})
	GetCodecPatterns = sync.OnceValue(func() map[string][]*regexp.Regexp {
		return compilePatternMap(CodecPatterns)
	})
	GetSourcePatterns = sync.OnceValue(func() map[string][]*regexp.Regexp {
		return compilePatternMap(SourcePatterns)
	})
	GetAudioPatterns = sync.OnceValue(func() map[string][]*regexp.Regexp {
		return compilePatternMap(AudioPatterns)
	})
	GetSeasonPatterns = sync.OnceValue(func() []*regexp.Regexp {
		return compilePatternSlice(SeasonPatterns)
	})
	GetEpisodePatterns = sync.OnceValue(func() []*regexp.Regexp {
		return compilePatternSlice(EpisodePatterns)
	})
	GetExtrasPatterns = sync.OnceValue(func() []*regexp.Regexp {
		return compilePatternSlice(ExtrasPatterns)
	})
)

func compilePatternMap(patterns map[string][]string) map[string][]*regexp.Regexp {
	result := make(map[string][]*regexp.Regexp, len(patterns))
	for key, patternList := range patterns {
		compiled := make([]*regexp.Regexp, len(patternList))
		for i, pattern := range patternList {
			compiled[i] = regexp.MustCompile("^" + pattern + "$")
		}
		result[key] = compiled
	}
	return result
}

func compilePatternSlice(patterns []string) []*regexp.Regexp {
	result := make([]*regexp.Regexp, len(patterns))
	for i, pattern := range patterns {
		result[i] = regexp.MustCompile("^" + pattern + "$")
	}
	return result
}
