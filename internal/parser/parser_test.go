package parser

import (
	"testing"
	"os"
	"path/filepath"
)


func TestParseTree(t *testing.T) {
	tmp := createDummyLibrary(t)
	ParseTree(tmp)
}

func createDummyLibrary(t *testing.T) string {
	dir := t.TempDir()

	movieDir := filepath.Join(dir, "my dummy movie")
	err := os.Mkdir(movieDir, 0755)
	if err != nil {
		t.Fatalf("Unable to create dummy movie dir %v, error %v\n", movieDir, err)
	}

	movie := filepath.Join(movieDir, "my dummy movie.mp4")
	_, err = os.Create(movie)
	if err != nil {
		t.Fatalf("Unable to create dummy movie %v, error %v", movie, err)
	}

	return dir
}
