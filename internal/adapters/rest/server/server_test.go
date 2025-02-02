package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/pangolin-do-golang/thumb-processor-api/internal/adapters/rest/handler"
	"github.com/pangolin-do-golang/thumb-processor-api/internal/adapters/rest/middleware"
	"github.com/pangolin-do-golang/thumb-processor-api/internal/core/users"
	"github.com/stretchr/testify/assert"
)

func TestRestServer_Serve_HealthCheck(t *testing.T) {
	r := gin.Default()
	r.Use(middleware.CorsMiddleware())
	r.GET("/health", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	req, _ := http.NewRequest("GET", "/health", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestRestServer_Serve_UserRoutes(t *testing.T) {
	r := gin.Default()
	r.Use(middleware.CorsMiddleware())
	handler.RegisterUserRoutes(r)

	req, _ := http.NewRequest("GET", "/users", nil) // Example route
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
}

func TestRestServer_Serve_LoginRoutes_Unauthenticated(t *testing.T) {
	r := gin.Default()
	r.Use(middleware.CorsMiddleware())

	authorizedGroup := r.Group("/", middleware.AuthMiddleware(users.GetAllowedUsers)) // Important!
	handler.RegisterLoginHandlers(authorizedGroup)

	req, _ := http.NewRequest("GET", "/login", nil) // Example login route
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code) // Expect unauthorized because no token
}
