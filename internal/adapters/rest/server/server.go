package server

import (
	"github.com/gin-gonic/gin"
	"github.com/pangolin-do-golang/thumb-processor-api/internal/adapters/rest/handler"
	"github.com/pangolin-do-golang/thumb-processor-api/internal/adapters/rest/middleware"
	"github.com/pangolin-do-golang/thumb-processor-api/internal/core/users"
)

type RestServer struct {
}

type RestServerOptions struct {
}

func NewRestServer(_ *RestServerOptions) *RestServer {
	return &RestServer{}
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

	err := r.Run("0.0.0.0:8080")
	if err != nil {
		panic(err)
	}
}
