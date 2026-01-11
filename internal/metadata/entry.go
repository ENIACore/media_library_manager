package metadata

type Entry struct {
	Parent		*Entry
	Children	[]*Entry
	Depth		int			// Root level entry should be Depth 0

	MediaInfo
	PathInfo
}
