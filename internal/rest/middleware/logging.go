package middleware

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/ahmadmilzam/ewallet/pkg/logger"
)

// responseWriter is a minimal wrapper for http.ResponseWriter that allows the
// written HTTP status code to be captured for logging.
type responseWriter struct {
	http.ResponseWriter
	status      int
	wroteHeader bool
}

func wrapResponseWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{ResponseWriter: w}
}

func (rw *responseWriter) Status() int {
	return rw.status
}

func (rw *responseWriter) WriteHeader(code int) {
	if rw.wroteHeader {
		return
	}

	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
	rw.wroteHeader = true
}

// Logging: logs the incoming HTTP request & its duration.
func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var level string

		defer func() {
			if err := recover(); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}
		}()

		start := time.Now()
		wrapped := wrapResponseWriter(w)
		next.ServeHTTP(wrapped, r)
		statusCode := wrapped.status

		switch {
		case statusCode >= 400:
			level = "WARN"
		case statusCode >= 500:
			level = "ERROR"
		default:
			level = "INFO"
		}

		slog.LogAttrs(
			r.Context(),
			logger.ParseLevel(level),
			"Request log",
			slog.String("duration", time.Since(start).String()),
			slog.String("method", r.Method),
			slog.String("path", r.URL.EscapedPath()),
			slog.Any("status", statusCode),
			slog.String("msg", "request"),
		)
	})
}

/*
func RequestLog() Middleware {
	return func(ctx *gin.Context) {
		t1 := time.Now().UnixNano() / int64(time.Millisecond)
		ctx.Next()
		t2 := time.Now().UnixNano() / int64(time.Millisecond)
		diff := t2 - t1

		msg, msg_exist := ctx.Get("msg")
		if !msg_exist {
			msg = ""
		}
		err, err_exist := ctx.Get("err")
		if !err_exist {
			err = ""
		}

		// access the status we are sending
		status := ctx.Writer.Status()

		switch {
		case status >= 500:
			slog.Error(
				msg.(string),
				"type", "Unexpected Error",
				"error", err,
				"duration", float64(diff),
				"path", ctx.Request.URL.EscapedPath(),
				"method", ctx.Request.Method,
				"status", status)
		case status >= 400:
			slog.Warn(
				msg.(string),
				"type", "Expected Error",
				"error", err,
				"duration", float64(diff),
				"path", ctx.Request.URL.EscapedPath(),
				"method", ctx.Request.Method,
				"status", status)
		default:
			slog.Info(
				"Request processed",
				"duration", float64(diff),
				"path", ctx.Request.URL.EscapedPath(),
				"method", ctx.Request.Method,
				"status", status)
		}
	}
}
*/
