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

func TestRestServer_Serve(t *testing.T) {
	// Create a new RestServer
	NewRestServer(&RestServerOptions{})

	// Create a new Gin engine
	r := gin.Default()
	r.Use(middleware.CorsMiddleware())

	// Register handlers
	handler.RegisterHealthCheck(r)
	handler.RegisterSwaggerHandlers(r)
	handler.RegisterUserRoutes(r)

	// Routes that need authentication
	authorizedGroup := r.Group("/", middleware.AuthMiddleware(users.GetAllowedUsers))
	handler.RegisterLoginHandlers(authorizedGroup)

	// Create a test HTTP server
	ts := httptest.NewServer(r)
	defer ts.Close()

	// Make a request to the /health endpoint
	resp, err := http.Get(ts.URL + "/health")
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	// Assert that the status code is 200 OK
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
