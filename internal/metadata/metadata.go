package metadata

type MediaMetadata struct {
	title		string
	year		int
	episode		int
	season		int

	resolution	string
	codec		string
	source		string
	audio		string

	language	string
}

type PathMetadata struct {
	dest			string
	source			string
	ext				string
	format			FormatType
}

type Metadata struct {
	Media	MediaMetadata
	Path	PathMetadata
}
