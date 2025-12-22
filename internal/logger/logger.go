package logger

import (
	"io"
	"fmt"
	"context"
	"log/slog"
	"os"
	"path/filepath"
	"sync"
	"time"
	"github.com/ENIACore/media_library_manager/internal/config"
)

var getSessionTimestamp = sync.OnceValue(func() string {
	return time.Now().Format("2006-01-02_15-04-05")
})

func getLogFile(filename string) io.Writer {
	cfg := config.Load()

	dirpath := filepath.Join(cfg.ManagerPath, "logs", getSessionTimestamp())
	err := os.MkdirAll(dirpath, 0755)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Warning: failed to create log directory: %v, using Stdout instead of log file %v", err, filename)
		return os.Stdout
	}

	logpath := filepath.Join(dirpath, filename)
	file, err := os.OpenFile(logpath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Warning: failed to create log file: %v, using Stdout instead", filename)
		return os.Stdout
	}
	return file

}

type multiHandler struct {
	handlers []slog.Handler
}

var logger = sync.OnceValue(func() *slog.Logger {
	cfg := config.Load()

	debugFile := getLogFile("DEBUG.log")
	infoFile := getLogFile("INFO.log")
	warnFile := getLogFile("WARN.log")

	handler := &multiHandler{
		handlers: []slog.Handler{
			slog.NewTextHandler(debugFile, &slog.HandlerOptions{Level: slog.LevelDebug}),
			slog.NewTextHandler(infoFile, &slog.HandlerOptions{Level: slog.LevelInfo}),
			slog.NewTextHandler(warnFile, &slog.HandlerOptions{Level: slog.LevelWarn}),
		},
	}

	return slog.New(handler).With("dry-run", cfg.DryRun)
})

func (h *multiHandler) Enabled(ctx context.Context, level slog.Level) bool {
	for _, handler := range h.handlers {
		if handler.Enabled(ctx, level) {
			return true
		}
	}
	return false
}

func (h *multiHandler) Handle(ctx context.Context, r slog.Record) error {
	for _, handler := range h.handlers {
		if handler.Enabled(ctx, r.Level) {
			if err := handler.Handle(ctx, r); err != nil {
				return err
			}
		}
	}
	return nil
}

func (h *multiHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	handlers := make([]slog.Handler, len(h.handlers))
	for i, handler := range h.handlers {
		handlers[i] = handler.WithAttrs(attrs)
	}
	return &multiHandler{handlers: handlers}
}

func (h *multiHandler) WithGroup(name string) slog.Handler {
	handlers := make([]slog.Handler, len(h.handlers))
	for i, handler := range h.handlers {
		handlers[i] = handler.WithGroup(name)
	}
	return &multiHandler{handlers: handlers}
}

func Debug(msg string, args ...any) { logger().Debug(msg, args...) }
func Info(msg string, args ...any)  { logger().Info(msg, args...) }
func Warn(msg string, args ...any)  { logger().Warn(msg, args...) }
func Error(msg string, args ...any) { logger().Error(msg, args...) }
