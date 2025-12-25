package parser


import (
	"fmt"
	//"github.com/ENIACore/media_library_manager/internal/metadata"
	"path/filepath"
	"io/fs"
)


type Node struct {
	parent		*Node
	children	[]*Node

}

func ParseTree(path string) *Node {
	err := filepath.WalkDir(path, func(path string, d fs.DirEntry, err error) error {
		fmt.Println("================")
		fmt.Println("Path is, " + path)	
		fmt.Printf("Name is, %v\n", d.Name())
		fmt.Printf("IsDir is, %v\n", d.IsDir())
		info, _ := d.Info()
		fmt.Printf("Info is, %v\n", info.Name())
		fmt.Println("================")
		return err
	})
	fmt.Println(err)
	return nil
}
