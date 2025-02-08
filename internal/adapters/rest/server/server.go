package server

import (
	"github.com/gin-gonic/gin"
	"github.com/pangolin-do-golang/thumb-processor-api/internal/adapters/rest/handler"
	"github.com/pangolin-do-golang/thumb-processor-api/internal/adapters/rest/middleware"
	"github.com/pangolin-do-golang/thumb-processor-api/internal/core/thumb"
	"github.com/pangolin-do-golang/thumb-processor-api/internal/core/users"
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
	authorizedGroup := r.Group("/", middleware.AuthMiddleware(users.GetAllowedUsers))

	handler.RegisterLoginHandlers(authorizedGroup)

	handler.NewThumbHandler(rs.thumbService).RegisterRoutes(r)

	err := r.Run("0.0.0.0:8080")
	if err != nil {
		panic(err)
	}
}
