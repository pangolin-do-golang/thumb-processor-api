package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestRegisterHealthCheck(t *testing.T) {
	// Create a new Gin engine
	router := gin.New()

	// Register the health check endpoint
	RegisterHealthCheck(router)

	// Create a request to the /health endpoint
	req, err := http.NewRequest(http.MethodGet, "/health", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder
	w := httptest.NewRecorder()

	// Serve the request to the router
	router.ServeHTTP(w, req)

	// Assert that the status code is 200 OK
	assert.Equal(t, http.StatusOK, w.Code)
}
