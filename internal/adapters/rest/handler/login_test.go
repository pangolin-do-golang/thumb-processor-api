package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestRegisterLoginHandlers(t *testing.T) {
	gin.SetMode(gin.TestMode) // Set Gin to test mode

	t.Run("Unauthorized - No Basic Auth", func(t *testing.T) {
		router := gin.New()
		authorizedGroup := router.Group("/")
		RegisterLoginHandlers(authorizedGroup)

		req, _ := http.NewRequest("GET", "/login", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
		var response map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "Unauthorized", response["error"])

	})

	t.Run("Authorized - Valid Basic Auth", func(t *testing.T) {
		router := gin.New()
		authorizedGroup := router.Group("/")
		RegisterLoginHandlers(authorizedGroup)

		req, _ := http.NewRequest("GET", "/login", nil)
		req.SetBasicAuth("testuser", "testpassword") // Set valid Basic Auth
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		var response map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "testuser", response["nickname"])
		assert.Equal(t, "logged", response["status"])
	})

	t.Run("Authorized - Empty Password", func(t *testing.T) {
		router := gin.New()
		authorizedGroup := router.Group("/")
		RegisterLoginHandlers(authorizedGroup)

		req, _ := http.NewRequest("GET", "/login", nil)
		req.SetBasicAuth("testuser", "") // Set valid Basic Auth, empty password
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		var response map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "testuser", response["nickname"])
		assert.Equal(t, "logged", response["status"])
	})
}
