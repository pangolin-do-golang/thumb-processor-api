package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pangolin-do-golang/thumb-processor-api/internal/core/users"
)

type CreateUserRequest struct {
	Nickname string `json:"nickname" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// RegisterUserRoutes registers the user-related routes with the Gin engine.
// @Summary Create a new user
// @Description Creates a new user with the provided nickname and password.
// @Tags users
// @Accept json
// @Produce json
// @Param request body CreateUserRequest true "User creation request body"
// @Success 201 {object} interface{} "User created"
// @Failure 400 {object} interface{} "Bad request"
// @Router /user [post]
func RegisterUserRoutes(router *gin.Engine) {
	router.POST("/user", createUserHandler)
}

func createUserHandler(c *gin.Context) {
	var req CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	users.CreateUser(req.Nickname, req.Password)

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}
