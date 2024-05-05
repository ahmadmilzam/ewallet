package middleware

// responseWriter is a minimal wrapper for http.ResponseWriter that allows the
// written HTTP status code to be captured for logging.
/*
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
invalid memory address or nil pointer dereference\ngoroutine 4 [running]:
net/http.(*conn).serve.func1()\n\t/usr/local/go/src/net/http/server.go:1898
/usr/local/go/src/runtime/panic.go:770
github.com/ahmadmilzam/ewallet/pkg/logger.ErrAttr({0x0?, 0x0?})\n\t
Users/ahmadmilzam/Development/Golang Project/ewallet/pkg/logger/logger.go:36
github.com/ahmadmilzam/ewallet/internal/usecase.(*AppUsecase).GetAccount(0x140001caa10, {0x100866ad0?, 0x100bea420?}, {0x140000105b0?, 0x140001630a8?})\n
Users/ahmadmilzam/Development/Golang Project/ewallet/internal/usecase/account.go:144
github.com/ahmadmilzam/ewallet/internal/usecase.(*AppUsecase).CreateTransfer(0x140001caa10, {0x100866ad0, 0x100bea420}, {{0x140000105a0, 0xc}, {0x140000105b0, 0xe}, 0x2710, {0x140000105c0, 0x5}, ...})\
Users/ahmadmilzam/Development/Golang Project/ewallet/internal/usecase/transfer.go:60
github.com/ahmadmilzam/ewallet/internal/rest/v1.(*TransferHandler).createTransfer(0x140001ee7f0, {0x1008662a0, 0x140002920e0}, 0x14000000b40)
Users/ahmadmilzam/Development/Golang Project/ewallet/internal/rest/v1/transfer.go:36
net/http.HandlerFunc.ServeHTTP(0x14000000a20?, {0x1008662a0?, 0x140002920e0?}, 0x2?)
/usr/local/go/src/net/http/server.go:2166 +0x38
github.com/gorilla/mux.(*Router).ServeHTTP(0x140001fa0c0, {0x1008662a0, 0x140002920e0}, 0x14000000900)
Users/ahmadmilzam/Development/go/pkg/mod/github.com/gorilla/mux@v1.8.1/mux.go:212
net/http.serverHandler.ServeHTTP({0x100864fb8?}, {0x1008662a0?, 0x140002920e0?}, 0x6?)
/usr/local/go/src/net/http/server.go:3137 +0xbc\nnet/http.(*conn).serve(0x140002883f0, {0x100866e88, 0x1400007e1e0})
/usr/local/go/src/net/http/server.go:2039 +0x508\ncreated by net/http.(*Server).Serve in goroutine 36
/usr/local/go/src/net/http/server.go:3285 +0x3f0"}
*/
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
