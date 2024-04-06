package handler

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/ahmadmilzam/ewallet/model"
	"github.com/gin-gonic/gin"
)

// New initialises a new API server with a router and default ping handler.
func NewServer(store model.Store) *Server {
	gin.SetMode(gin.DebugMode)
	router := gin.New()
	router.HandleMethodNotAllowed = true
	return &Server{
		Router: router,
		Store:  store,
	}
}

// Server wraps an HTTP router and acts as an HTTP server
type Server struct {
	Router *gin.Engine
	Store  model.Store
}

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
