package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/pangolin-do-golang/thumb-processor-api/internal/core/users"
	"net/http"
)

func AuthMiddleware(allowedUsersFunc func() gin.Accounts) gin.HandlerFunc {
	return func(c *gin.Context) {
		accounts := allowedUsersFunc() // Get the latest user list

		username, password, ok := c.Request.BasicAuth()
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			return
		}

		user, ok := accounts[username]
		if !ok || user != password {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			return
		}

		loggedUser := users.GetUserByNickname(username)

		c.Set("logged_user_id", loggedUser.ID)

		c.Next() // Continue to the handler
	}
}
