package main

import (
	"fmt"
	"github.com/ENIACore/media_library_manager/internal/processor"
)

func main() {
	result := processor.Main("message")
	fmt.Println(result)
}
