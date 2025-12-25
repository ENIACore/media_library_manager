package metadata

// NodeType represents the type of a filesystem node
type NodeType string

const (
    // UnknownType
    Unknown NodeType = "UNKNOWN"
    
    // DirectoryType
    SeriesFolder   NodeType = "SERIES_FOLDER"
    SeasonFolder   NodeType = "SEASON_FOLDER"
    MovieFolder    NodeType = "MOVIE_FOLDER"
    SubtitleFolder NodeType = "SUBTITLE_FOLDER"
    ExtrasFolder   NodeType = "EXTRAS_FOLDER"
    
    // FileType
    MovieFile    NodeType = "MOVIE_FILE"
    EpisodeFile  NodeType = "EPISODE_FILE"
    SubtitleFile NodeType = "SUBTITLE_FILE"
    ExtrasFile   NodeType = "EXTRAS_FILE"
)

// FormatType represents the format of a file
type FormatType string

const (
    Video    FormatType = "VIDEO"
    Subtitle FormatType = "SUBTITLE"
    UnknownFormat FormatType = "UNKNOWN"
)
