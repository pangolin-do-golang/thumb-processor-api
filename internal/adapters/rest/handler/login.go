package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// RegisterLoginHandlers registers the login handler with the given Gin router group.
//
// @Summary User Login
// @Description Authenticates a user using Basic Authentication and returns user information.
// @Tags Auth
// @Accept json
// @Produce json
// @Param Authorization header string true "Basic Authentication credentials (username:password)"
// @Success 200 {object} interface{} "Successful login"
// @Failure 401 {object} interface{} "Unauthorized"
// @Router /login [get]
func RegisterLoginHandlers(authorizedGroup *gin.RouterGroup) {
	authorizedGroup.GET("login", func(c *gin.Context) {
		user, _, ok := c.Request.BasicAuth()

		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"nickname": user,
			"status":   "logged",
		})
	})
}
