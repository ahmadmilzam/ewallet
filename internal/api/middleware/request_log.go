package middleware

import (
	"log/slog"
	"time"

	"github.com/gin-gonic/gin"
)

func RequestLog() Middleware {
	return func(ctx *gin.Context) {
		t1 := time.Now().UnixNano() / int64(time.Millisecond)
		ctx.Next()

		t2 := time.Now().UnixNano() / int64(time.Millisecond)
		diff := t2 - t1

		// access the status we are sending
		status := ctx.Writer.Status()

		if status >= 500 {
			slog.Error("error", "duration", float64(diff), "path", ctx.Request.URL.EscapedPath(), "method", ctx.Request.Method)
		} else {
			slog.Info("info", "duration", float64(diff), "path", ctx.Request.URL.EscapedPath(), "method", ctx.Request.Method)
		}

	}
}
