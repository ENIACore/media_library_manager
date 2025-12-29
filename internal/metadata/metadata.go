package metadata

type Media struct {
	Title		string
	Year		int
	Episode		int // -1 if no season pattern, 0 if season pattern, > 0 if season pattern with season number
	Season		int // -1 if no ep pattern, 0 if ep pattern, > 0 if ep pattern with ep number
	Resolution	string // Empty if not found
	Codec		string // Empty if not found
	Source		string // Empty if not found
	Audio		string // Empty if not found

	Language	string
}

type Path struct {
	Dest			string
	Source			string
	IsDir			bool
	Ext				string
	Format			FormatType
}
