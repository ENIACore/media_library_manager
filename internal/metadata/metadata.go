package metadata

type MediaMetadata struct {
	Title		string
	Year		int
	Episode		int
	Season		int

	Resolution	string
	Codec		string
	Source		string
	Audio		string

	Language	string
}

type PathMetadata struct {
	Dest			string
	Source			string
	Ext				string
	Format			FormatType
}

type Metadata struct {
	Media	MediaMetadata
	Path	PathMetadata
}
