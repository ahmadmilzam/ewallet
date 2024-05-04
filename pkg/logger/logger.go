package logger

import (
	"io"
	"log/slog"
	"path/filepath"
	"slices"
)

const (
	timeFormat   = "2006-01-02T15:04:05.000"
	maskedString = "*"
	errorKey     = "error"
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
		AddSource:   true,
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
		if attr.Key == slog.SourceKey {
			source := attr.Value.Any().(*slog.Source)
			source.Function = filepath.Base(source.Function)
			source.File = filepath.Base(source.File)
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
