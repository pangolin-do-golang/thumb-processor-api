package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// RegisterHealthCheck registers the health check endpoint with the given Gin engine.
//
// @Summary      Health Check
// @Description  Checks the health status of the application.
// @Produce      json
// @Success     200 {object} map[string]interface{}
// @Router      /health [get]
func RegisterHealthCheck(r *gin.Engine) {
	r.GET("/health", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})
}
