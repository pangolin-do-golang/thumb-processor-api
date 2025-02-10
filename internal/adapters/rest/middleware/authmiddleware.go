package middleware

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pangolin-do-golang/thumb-processor-api/internal/core/users"
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

		c.Set("user", username)
		loggedUser := users.GetUserByNickname(username)

		ctx := context.WithValue(c.Request.Context(), "logged_user_id", loggedUser.ID)
		c.Request = c.Request.WithContext(ctx)

		c.Next() // Continue to the handler
	}
}
