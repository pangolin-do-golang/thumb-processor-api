package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// RegisterHealthCheck registers the health check endpoint with the given Gin engine.
//
// @Tags Health Check
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
