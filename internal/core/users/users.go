package users

import (
	"github.com/gin-gonic/gin"
	"github.com/pangolin-do-golang/thumb-processor-api/internal/adapters/db"
)

var users = []db.User{
	{Nickname: "user", Password: "user"},
	{Nickname: "test", Password: "test"},
	{Nickname: "prod", Password: "prod"},
}

func GetAllowedUsers() gin.Accounts {
	accounts := make(gin.Accounts)
	for _, user := range users {
		accounts[user.Nickname] = user.Password
	}
	return accounts
}

func CreateUser(nickname, password string) {
	users = append(users, db.User{Nickname: nickname, Password: password})
}
