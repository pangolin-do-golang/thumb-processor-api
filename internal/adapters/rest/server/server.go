package server

import (
	"github.com/gin-gonic/gin"
	"github.com/pangolin-do-golang/thumb-processor-api/internal/adapters/rest/handler"
	"github.com/pangolin-do-golang/thumb-processor-api/internal/adapters/rest/middleware"
	"github.com/pangolin-do-golang/thumb-processor-api/internal/core/thumb"
)

type RestServer struct {
	thumbService thumb.IThumbService
}

type RestServerOptions struct {
	ThumService thumb.IThumbService
}

func NewRestServer(opts *RestServerOptions) *RestServer {
	return &RestServer{
		thumbService: opts.ThumService,
	}
}

func (rs RestServer) Serve() {
	r := gin.Default()
	r.Use(middleware.CorsMiddleware())

	handler.RegisterHealthCheck(r)

	handler.RegisterSwaggerHandlers(r)

	handler.RegisterUserRoutes(r)

	// Rotes that need authentication
	handler.RegisterLoginHandlers(r.Group("/"))

	handler.NewThumbHandler(rs.thumbService).RegisterRoutes(r.Group("/"))

	err := r.Run("0.0.0.0:8080")
	if err != nil {
		panic(err)
	}
}
