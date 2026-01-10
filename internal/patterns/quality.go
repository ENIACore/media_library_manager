package patterns

import (
	"sync"
)

var ResolutionPatternGroups = []PatternGroup{
	// Matches 8k, 4320, 4320P, 4320I, 7680X4320, FULLUHD
	{Key: `8K`, Patterns: []Pattern{`8K`, `4320[PI]?`, `7680X4320`, `FULLUHD`}},
	// Matches 4k, UHD, 2160, 2160P, 2160I, 3840X2160
	{Key: `4K`, Patterns: []Pattern{`4K`, `UHD`, `2160[PI]?`, `3840X2160`}},
	// Matches 2K, 1440, 1440P, 1440I, 2560X1440, QHD, WQHD
	{Key: `2K`, Patterns: []Pattern{`2K`, `1440[PI]?`, `2560X1440`, `QHD`, `WQHD`}},
	// Matches 1080, 1080P, 1080I, FHD, 1920X1080, FULLHD
	{Key: `1080P`, Patterns: []Pattern{`1080[PI]?`, `FHD`, `1920X1080`, `FULLHD`}},
	// Matches 720, 720P, 720I, 1280X720
	{Key: `720P`, Patterns: []Pattern{`720[PI]?`, `1280X720`}},
	// Matches 576, 576P, 576I, PAL
	{Key: `576P`, Patterns: []Pattern{`576[PI]?`, `PAL`}},
	// Matches 480, 480P, 480I, NTSC
	{Key: `480P`, Patterns: []Pattern{`480[PI]?`, `NTSC`}},
	// Matches 360, 360P, 360I
	{Key: `360P`, Patterns: []Pattern{`360[PI]?`}},
	// Matches 240, 240P, 240I
	{Key: `240P`, Patterns: []Pattern{`240[PI]?`}},
}

var CodecPatternGroups = []PatternGroup{
	// Matches AV1, SVT-AV1, SVTAV1
	{Key: `AV1`, Patterns: []Pattern{`AV1`, `SVT\.AV1`, `SVTAV1`, `AOV1`}},
	// Matches VP9
	{Key: `VP9`, Patterns: []Pattern{`VP9`}},
	// Matches VP8
	{Key: `VP8`, Patterns: []Pattern{`VP8`}},
	// Matches x265, X265, X.265, H265, H.265, HEVC, HEVC10, HEVC10BIT
	{Key: `X265`, Patterns: []Pattern{`X265`, `X\.265`, `H265`, `H\.265`, `HEVC`, `HEVC10`, `HEVC10BIT`, `H265P`}},
	// Matches x264, X264, X.264, H264, H.264, AVC, AVC1
	{Key: `X264`, Patterns: []Pattern{`X264`, `X\.264`, `H264`, `H\.264`, `AVC`, `AVC1`, `H264P`}},
	// Matches x263, X263, X.263, H263, H.263
	{Key: `X263`, Patterns: []Pattern{`X263`, `X\.263`, `H263`, `H\.263`}},
	// Matches XVID, XVID-AF
	{Key: `XVID`, Patterns: []Pattern{`XVID`, `XVID\.AF`}},
	// Matches DIVX, DIV3, DIVX6
	{Key: `DIVX`, Patterns: []Pattern{`DIVX`, `DIV3`, `DIVX6`}},
	// Matches MPEG-4, MPEG4, MP4V
	{Key: `MPEG4`, Patterns: []Pattern{`MPEG\.4`, `MPEG4`, `MP4V`}},
	// Matches MPEG-2, MPEG2, MP2V
	{Key: `MPEG2`, Patterns: []Pattern{`MPEG\.2`, `MPEG2`, `MP2V`}},
	// Matches MPEG-1, MPEG1, MP1V
	{Key: `MPEG1`, Patterns: []Pattern{`MPEG\.1`, `MPEG1`, `MP1V`}},
	// Matches VC-1, VC1, WMV3, WVC1
	{Key: `VC1`, Patterns: []Pattern{`VC\.1`, `VC1`, `WMV3`, `WVC1`}},
	// Matches THEORA
	{Key: `THEORA`, Patterns: []Pattern{`THEORA`}},
	// Matches PRORES, PRORES422, PRORES4444
	{Key: `PRORES`, Patterns: []Pattern{`PRORES`, `PRORES422`, `PRORES4444`, `PRORES422HQ`}},
	// Matches DNxHD, DNXHD, DNxHR, DNXHR
	{Key: `DNXHD`, Patterns: []Pattern{`DNXHD`, `DNXHR`}},
}
var SourcePatternGroups = []PatternGroup{
	// Matches REMUX
	{Key: `REMUX`, Patterns: []Pattern{`REMUX`}},
	// Matches BluRay variants
	{Key: `BLURAY`, Patterns: []Pattern{`BLURAY`, `BDRIP`, `BD\.RIP`, `BR\.RIP`, `BRRIP`, `BDMV`, `BDISO`, `BD25`, `BD50`, `BD66`, `BD100`}},
	// Matches WEB-DL, WEBDL, WEB DL
	{Key: `WEB-DL`, Patterns: []Pattern{`WEB\.DL`, `WEBDL`}},
	// Matches WEBRip, WEB-RIP, WEBRIP, WEB RIP
	{Key: `WEBRIP`, Patterns: []Pattern{`WEBRIP`, `WEB-RIP`, `WEB\.RIP`}},
	// Matches WEB (but not WEB-DL or WEBRip)
	{Key: `WEB`, Patterns: []Pattern{`WEB`}},
	// Matches HDRip, HD-RIP, HDRIP, HD RIP
	{Key: `HDRIP`, Patterns: []Pattern{`HDRIP`, `HD\.RIP`}},
	// Matches DVDRip, DVD-RIP, DVDRIP, DVD RIP
	{Key: `DVDRIP`, Patterns: []Pattern{`DVDRIP`, `DVD\.RIP`}},
	// Matches DVD, DVDSCR, DVD5, DVD9
	{Key: `DVD`, Patterns: []Pattern{`DVD`, `DVDSCR`, `DVD5`, `DVD9`}},
	// Matches HDTV, HDTVRIP, DTTV, PDTV, SDTV, LDTV
	{Key: `HDTV`, Patterns: []Pattern{`HDTV`, `HDTVRIP`, `DTTV`, `PDTV`, `SDTV`, `LDTV`}},
	// Matches TELECINE, TC
	{Key: `TELECINE`, Patterns: []Pattern{`TELECINE`, `TC`}},
	// Matches TELESYNC, TS
	{Key: `TELESYNC`, Patterns: []Pattern{`TELESYNC`, `TS`}},
	// Matches SCREENER, SCR, DVDSCR, BDSCR
	{Key: `SCREENER`, Patterns: []Pattern{`SCREENER`, `SCR`, `DVDSCR`, `BDSCR`}},
	// Matches CAMRIP, CAM, HDCAM
	{Key: `CAM`, Patterns: []Pattern{`CAMRIP`, `CAM`, `HDCAM`}},
	// Matches WORKPRINT, WP
	{Key: `WORKPRINT`, Patterns: []Pattern{`WORKPRINT`, `WP`}},
	// Matches PPV, PPVRIP
	{Key: `PPV`, Patterns: []Pattern{`PPV`, `PPVRIP`}},
	// Matches VODRIP, VOD
	{Key: `VODRIP`, Patterns: []Pattern{`VODRIP`, `VOD`}},
	// Matches HC, HCHDCAM
	{Key: `HC`, Patterns: []Pattern{`HC`, `HCHDCAM`}},
	// Matches LINE
	{Key: `LINE`, Patterns: []Pattern{`LINE`}},
	// Matches HDTS, HD-TS, HDTS
	{Key: `HDTS`, Patterns: []Pattern{`HDTS`, `HD\.TS`}},
	// Matches HDTC, HD-TC, HDTC
	{Key: `HDTC`, Patterns: []Pattern{`HDTC`, `HD\.TC`}},
	// Matches TVRIP, SATRIP, DTTVRIP
	{Key: `TVRIP`, Patterns: []Pattern{`TVRIP`, `SATRIP`, `DTTVRIP`}},
}

var AudioPatternGroups = []PatternGroup{
	// Matches ATMOS, DOLBY-ATMOS, DOLBY ATMOS, DOLBYATMOS
	{Key: `ATMOS`, Patterns: []Pattern{`ATMOS`, `DOLBY-ATMOS`, `DOLBY\.ATMOS`, `DOLBYATMOS`}},
	// Matches DTS-X, DTSX, DTS X
	{Key: `DTS-X`, Patterns: []Pattern{`DTSX`, `DTS\.X`, `DTS`}},
	// Matches DTS-HD-MA, DTS-HD, DTSHD-MA, DTSHD
	{Key: `DTS-HD`, Patterns: []Pattern{`DTS\.HD\.MA`, `DTSHD-MA`, `DTSHD\.MA`, `DTS\.HD`, `DTSHD`}},
	// Matches DTS-MA, DTSMA
	{Key: `DTS-MA`, Patterns: []Pattern{`DTS\.MA`, `DTSMA`}},
	// Matches DTS-ES, DTSES, DTS ES
	{Key: `DTS-ES`, Patterns: []Pattern{`DTS\.ES`, `DTSES`}},
	// Matches DTS (standalone, not part of other DTS variants)
	{Key: `DTS`, Patterns: []Pattern{`DTS`}},
	// Matches TrueHD, TRUE-HD, TRUEHD, TRUE HD
	{Key: `TRUEHD`, Patterns: []Pattern{`TRUEHD`, `TRUE\.HD`}},
	// Matches DD+, DDP, E-AC-3, EAC3, DD-PLUS, DDPLUS
	{Key: `DD+`, Patterns: []Pattern{`DDP`, `E\.AC\.3`, `EAC3`, `DD\.PLUS`, `DDPLUS`}},
	// Matches DD, AC3, DOLBY-DIGITAL, DOLBY DIGITAL (but not DD+ or variants)
	{Key: `DD`, Patterns: []Pattern{`DD`, `AC3`, `DOLBY-DIGITAL`, `DOLBY\.DIGITAL`, `DOLBYDIGITAL`}},
	// Matches AAC, HE-AAC, HEAAC, HE AAC
	{Key: `AAC`, Patterns: []Pattern{`AAC`, `HE\.AAC`, `HEAAC`}},
	// Matches FLAC
	{Key: `FLAC`, Patterns: []Pattern{`FLAC`}},
	// Matches MP3
	{Key: `MP3`, Patterns: []Pattern{`MP3`}},
	// Matches LPCM, PCM
	{Key: `LPCM`, Patterns: []Pattern{`LPCM`, `PCM`}},
	// Matches OGG, VORBIS
	{Key: `OGG`, Patterns: []Pattern{`OGG`, `VORBIS`}},
	// Matches OPUS
	{Key: `OPUS`, Patterns: []Pattern{`OPUS`}},
	// Matches 5.1, 5-1, 51, 6CH
	{Key: `5.1`, Patterns: []Pattern{`5\.1`, `51`, `6CH`}},
	// Matches 7.1, 7-1, 71, 8CH
	{Key: `7.1`, Patterns: []Pattern{`7\.1`, `71`, `8CH`}},
	// Matches 2.0, 2-0, 20, STEREO, 2CH
	{Key: `2.0`, Patterns: []Pattern{`2\.0`, `20`, `STEREO`, `2CH`}},
	// Matches DUAL-AUDIO, DUAL AUDIO, DUAL
	{Key: `DUAL`, Patterns: []Pattern{`DUAL\.AUDIO`, `DUAL`}},
}

var (
	GetResolutionPatternGroups = sync.OnceValue(func() []CompiledPatternGroup {
		return compilePatternGroups(ResolutionPatternGroups)
	})
	GetCodecPatternGroups = sync.OnceValue(func() []CompiledPatternGroup {
		return compilePatternGroups(CodecPatternGroups)
	})
	GetSourcePatternGroups = sync.OnceValue(func() []CompiledPatternGroup {
		return compilePatternGroups(SourcePatternGroups)
	})
	GetAudioPatternGroups = sync.OnceValue(func() []CompiledPatternGroup {
		return compilePatternGroups(AudioPatternGroups)
	})
	GetQualityPatternGroups = sync.OnceValue(func() []CompiledPatternGroup {
		res := make([]CompiledPatternGroup, len(ResolutionPatternGroups) + len(CodecPatternGroups) + len(SourcePatternGroups) + len(AudioPatternGroups))

		for i, group := range GetResolutionPatternGroups() {
			res[i] = CompiledPatternGroup{
				Key: group.Key,
				Patterns: group.Patterns,
			}
		}
		for i, group := range GetCodecPatternGroups() {
			res[i] = CompiledPatternGroup{
				Key: group.Key,
				Patterns: group.Patterns,
			}
		}
		for i, group := range GetSourcePatternGroups() {
			res[i] = CompiledPatternGroup{
				Key: group.Key,
				Patterns: group.Patterns,
			}
		}
		for i, group := range GetAudioPatternGroups() {
			res[i] = CompiledPatternGroup{
				Key: group.Key,
				Patterns: group.Patterns,
			}
		}

		return res
	})
)
