package users

import (
	"github.com/gin-gonic/gin"
	"github.com/pangolin-do-golang/thumb-processor-api/internal/adapters/db"
	"math/rand"
)

var users = []db.User{
	{ID: 1, Nickname: "user", Password: "user"},
	{ID: 2, Nickname: "test", Password: "test"},
	{ID: 3, Nickname: "prod", Password: "prod"},
}

func GetAllowedUsers() gin.Accounts {
	accounts := make(gin.Accounts)
	for _, user := range users {
		accounts[user.Nickname] = user.Password
	}
	return accounts
}

func CreateUser(nickname, password string) {
	users = append(users, db.User{ID: rand.Int(), Nickname: nickname, Password: password})
}

func GetUserByNickname(nickname string) *db.User {
	for _, user := range users {
		if user.Nickname == nickname {
			return &user
		}
	}
	return nil
}
