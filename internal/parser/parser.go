package parser

import (
	"os"
	"fmt"
	"path/filepath"
	"github.com/ENIACore/media_library_manager/internal/metadata"
	"github.com/ENIACore/media_library_manager/internal/extractor"
	"log/slog"
)


type Node struct {
	parent		*Node
	children	[]*Node

	metadata.Metadata
}

func ParseTree(path string, parent *Node, logger *slog.Logger) (*Node, error) {

    info, err := os.Stat(path)
    if err != nil {
		return nil, fmt.Errorf("stat path %s, %w", path, err)
    }

    node := &Node{
        parent:		parent,
		Metadata:	extractor.Extract(path, logger),
    }
	if !info.IsDir() {
		return node, nil
	}

	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, fmt.Errorf("read dir %s, %w", path, err)
	}

	children := make([]*Node, 0, len(entries))
	for _, entry := range entries {
		childPath := filepath.Join(path, entry.Name())
		child, err := ParseTree(childPath, node, logger)
		if err != nil {
			return nil, err
		}
		children = append(children, child)
	}
	node.children = children
    
    return node, nil
}
