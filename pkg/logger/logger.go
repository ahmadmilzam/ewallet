package logger

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"path/filepath"
	"runtime"
	"slices"

	"github.com/ahmadmilzam/ewallet/internal/rest/middleware"
)

const (
	timeFormat   = "2006-01-02T15:04:05.000"
	maskedString = "*"
	errorKey     = "error"
	requestIDKey = "request_id"
	sourceKey    = "caller"
)

func ParseLevel(level string) slog.Level {
	logLevel := slog.Level(4)
	err := logLevel.UnmarshalText([]byte(level))
	if err != nil {
		slog.Error("failed to parse log level", ErrAttr(err))
	}
	return logLevel
}

func InitializeLogger(option *Option) {
	opts := &slog.HandlerOptions{
		AddSource:   false,
		Level:       ParseLevel(option.Level),
		ReplaceAttr: formatAttr(timeFormat),
	}
	logger := slog.New(newCustomHandler(option.Writer, opts))
	slog.SetDefault(logger)
}

func ErrAttr(err error) slog.Attr {
	return slog.String(errorKey, err.Error())
}

func formatAttr(timeFormat string) func(groups []string, attr slog.Attr) slog.Attr {
	return func(groups []string, attr slog.Attr) slog.Attr {
		if attr.Key == slog.TimeKey {
			attr.Value = slog.StringValue(attr.Value.Time().Format(timeFormat))
		}
		if isSensitiveKey(attr.Key) {
			attr.Value = slog.StringValue(maskedString)
		}
		return attr
	}
}

func isSensitiveKey(key string) bool {
	sensitiveKeys := []string{"pin", "password", "secret", "token", "credential"}
	return slices.Contains(sensitiveKeys, key)
}

type customHandler struct {
	slog.Handler
}

func newCustomHandler(w io.Writer, opts *slog.HandlerOptions) *customHandler {
	return &customHandler{
		slog.NewJSONHandler(w, opts),
	}
}

func (h *customHandler) Handle(ctx context.Context, r slog.Record) error {
	if requestID, ok := ctx.Value(middleware.RequestIDKey).(string); ok {
		r.AddAttrs(slog.String(requestIDKey, requestID))
	}

	// only shown when using >= WARN level log
	if r.Level >= slog.LevelWarn {
		fs := runtime.CallersFrames([]uintptr{r.PC})
		f, _ := fs.Next()
		r.AddAttrs(slog.String(sourceKey, fmt.Sprintf("@%s:#%d", filepath.Base(f.Function), f.Line)))
	}

	return h.Handler.Handle(ctx, r)
}
