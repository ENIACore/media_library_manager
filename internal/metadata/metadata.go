package metadata

type MediaInfo struct {
    Title      []string
    Year       *int     // nil if not found
    Episode    *int     // nil = no pattern, 0 = pattern but no number, >0 = ep number
    Season     *int     // nil = no pattern, 0 = pattern but no number, >0 = season number
    Resolution string   // "" if not found
    Codec      string	
    Source     string
    Audio      string
    Language   string
}

type PathInfo struct {
    Dest   string
    Source string
    Ext    string      // "" if no ext
    Type   ContentType // Unkown if directory or no ext
}
