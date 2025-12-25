package extractor

import (
	"strings"
	"log/slog"
	"regexp"
)

func sanitizeName(name string, logger *slog.Logger) string {
	log := logger.With("func", "sanitizeName")
	log.Debug("processing filename", "name", name)

	name = strings.ToUpper(name)
	name = strings.ReplaceAll(name, "'", "")
	name = strings.ReplaceAll(name, "\"", "")

	re := regexp.MustCompile("[^A-Z0-9]+")
	name = string(re.ReplaceAll([]byte(name), []byte(".")))

	name = strings.Trim(name, " .")

	log.Debug("returning sanitized name", "name", name)
	return name
}

// matchSegments joins dot-separated segments and attempts a full regex match.
// It joins only as many segments as needed based on the number of literal dots in the pattern. 
// Returns the match slice (full match + capture groups) if the entire joined string matches, or nil otherwise.
func matchSegments(segments []string, re *regexp.Regexp) []string {

	// Determine how many segments to join based on dots in pattern
    numDots := strings.Count(re.String(), `\.`)
    end := min(numDots+1, len(segments))
    
    str := strings.Join(segments[:end], ".")
    
    if match := re.FindStringSubmatch(str); match != nil && match[0] == str {
        return match
    }
    return nil

}
