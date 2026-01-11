package parser

import (
	"os"
	"fmt"
	"path/filepath"
	"github.com/ENIACore/media_library_manager/internal/metadata"
	"github.com/ENIACore/media_library_manager/internal/extractor"
	"log/slog"
)


type Entry struct {
	Parent		*Entry
	Children	[]*Entry
	Depth		int			// Root level entry should be Depth 0

	metadata.MediaInfo
	metadata.PathInfo
}

func ParseTree(path string, parent *Entry, depth int, logger *slog.Logger) (*Entry, error) {

    info, err := os.Stat(path)
    if err != nil {
		return nil, fmt.Errorf("stat path %s, %w", path, err)
    }

    node := &Entry{
        Parent:		parent,
		Depth:		depth,
		MediaInfo: extractor.ExtractMedia(path, logger),
		PathInfo: extractor.ExtractPath(path, logger),
    }
	if !info.IsDir() {
		return node, nil
	}

	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, fmt.Errorf("read dir %s, %w", path, err)
	}

	children := make([]*Entry, 0, len(entries))
	for _, entry := range entries {
		childPath := filepath.Join(path, entry.Name())
		child, err := ParseTree(childPath, node, depth + 1, logger)
		if err != nil {
			return nil, err
		}
		children = append(children, child)
	}
	node.Children = children
    
    return node, nil
}
