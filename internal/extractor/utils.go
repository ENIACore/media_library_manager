package extractor

import (
	"log/slog"
	"strings"
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

// Used to match one or more FULL parts starting from the left
/*
func matchParts(parts []string, re *regexp.Regexp) string {
	

	return ""
}
*/
