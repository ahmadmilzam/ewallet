package api

import (
	"net/http"

	v1 "github.com/ahmadmilzam/ewallet/internal/api/v1"
	"github.com/ahmadmilzam/ewallet/internal/usecase"
	"github.com/gin-gonic/gin"
)

// func NewServer(store entity.Store) *Server {
// 	gin.SetMode(gin.ReleaseMode)
// 	// router := gin.New()
// 	router := gin.Default()
// 	router.HandleMethodNotAllowed = true
// 	// v1.RegisterRoutes(router)

// 	return &Server{Router: router}
// }

// // Server wraps an HTTP router and acts as an HTTP server
// type Server struct {
// 	Router *gin.Engine
// }

// func (s *Server) Serve(baseCtx context.Context, addr string) error {
// 	ctx, cancel := context.WithCancel(baseCtx)
// 	defer cancel()

// 	srv := &http.Server{
// 		Addr:    addr,
// 		Handler: s.Router,
// 	}

// 	go func() {
// 		defer cancel()
// 		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
// 			fmt.Printf("server exited with error: %s\n", err.Error())
// 		}
// 	}()

// 	return s.waitForInterrupt(ctx, srv)
// }

// const gracefulPeriod = 1 * time.Second

// func (s *Server) waitForInterrupt(ctx context.Context, srv *http.Server) error {
// 	<-ctx.Done()

// 	shutdownCtx, cancel := context.WithTimeout(context.Background(), gracefulPeriod)
// 	defer cancel()
// 	return srv.Shutdown(shutdownCtx)
// }

func NewRouter(handler *gin.Engine, u usecase.AccountUsecaseInterface) {
	// Options -.
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	// K8s probe for kubernetes health checks -.
	handler.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, "The server is up and running")
	})

	// Handling a page not found endpoint -.
	handler.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "The requested page is not found.Please try later!"})
	})

	// Routers -.
	rgroupv1 := handler.Group("/api/v1")
	{
		v1.NewAccountRoute(rgroupv1, u)
	}
}
