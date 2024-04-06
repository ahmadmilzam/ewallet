package server

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// WriteJSON writes the given value as JSON encoded data to the response writer with given
// status. 'v' must be compatible with JSON Marshal.
func WriteJSON(wr http.ResponseWriter, status int, v interface{}) {
	wr.Header().Set("Content-type", "application/json; charset=utf-8")
	wr.WriteHeader(status)
	if v == nil {
		return
	}
	if err := json.NewEncoder(wr).Encode(v); err != nil {
		panic(fmt.Errorf("writeJSON failed: %s", err.Error()))
	}
}

// ReadJSON reads the body of the given request as JSON encoded data and unmarshalls it into
// the 'into' pointer. If the `into` has a `Validate() error` it will be used to validate the
// data after unmarshal.
func ReadJSON(req *http.Request, into interface{}) error {
	if err := json.NewDecoder(req.Body).Decode(into); err != nil {
		return fmt.Errorf("failed to decode body: %s", err.Error())
	}
	if v, ok := into.(interface{ Validate() error }); ok {
		return v.Validate()
	}
	return nil
}

func staticHandler(status int, message string) http.Handler {
	return http.HandlerFunc(func(wr http.ResponseWriter, req *http.Request) {
		WriteJSON(wr, status, map[string]string{"message": message})
	})
}

// New initialises a new API server with a router and default ping handler.
func New() Server {
	gin.SetMode(gin.DebugMode)
	router := gin.New()
	router.HandleMethodNotAllowed = true
	return Server{
		Router: router,
	}
}

// Server wraps an HTTP router and acts as an HTTP server
type Server struct {
	Router *gin.Engine
}

// Register registers the given routes to the server.
func (s *Server) Register(routes ...[]Route) (*Server, error) {
	for _, list := range routes {
		for _, route := range list {
			switch route.Method {
			case http.MethodGet:
				s.Router.GET(route.Path, route.Handler...)
				break
			case http.MethodPost:
				s.Router.POST(route.Path, route.Handler...)
				break
			case http.MethodPut:
				s.Router.PUT(route.Path, route.Handler...)
			default:
				return nil, fmt.Errorf("the http method %s is not supported", route.Method)
			}
		}
	}

	return s, nil
}

// Serve starts the server on given addr and blocks until the underlying server
// exits with error or the context is cancelled. If the context is cancelled, a
// graceful shutdown is performed with configured grace period.
func (s *Server) Serve(baseCtx context.Context, addr string) error {
	ctx, cancel := context.WithCancel(baseCtx)
	defer cancel()

	srv := &http.Server{
		Addr:    addr,
		Handler: s.Router,
	}

	go func() {
		defer cancel()
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("server exited with error: %s\n", err.Error())
		}
	}()

	return s.waitForInterrupt(ctx, srv)
}

const gracefulPeriod = 1 * time.Second

func (s *Server) waitForInterrupt(ctx context.Context, srv *http.Server) error {
	<-ctx.Done()

	shutdownCtx, cancel := context.WithTimeout(context.Background(), gracefulPeriod)
	defer cancel()
	return srv.Shutdown(shutdownCtx)
}

// Route represents an HTTP endpoint and the handler to be used.
type Route struct {
	Method  string
	Path    string
	Handler []gin.HandlerFunc
}
