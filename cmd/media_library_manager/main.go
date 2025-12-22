package main

import (
	"github.com/ENIACore/media_library_manager/internal/logger"
)

func main() {
	logger := logger.GetLogger()
	logger.Info("test here")
}
