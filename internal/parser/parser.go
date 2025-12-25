package parser

import (
	"os"
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

func ParseTree(path string, parent *Node) (*Node, error) {
    info, err := os.Stat(path)
    if err != nil {
        return nil, err
    }

    node := &Node{
        parent: parent,
		children: nil,
		Metadata:	extractor.Extract(path),
    }
	if !info.IsDir() {
		return node, nil
	}

	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	children := make([]*Node, 0, len(entries))
	for _, entry := range entries {
		childPath := filepath.Join(path, entry.Name())
		child, err := ParseTree(childPath, node)
		if err != nil {
			return nil, err
		}
		if child != nil {
			children = append(children, child)
		}
	}
	node.children = children
    
    return node, nil
}
