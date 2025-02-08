package users

import (
	"testing"

	"github.com/pangolin-do-golang/thumb-processor-api/internal/adapters/db"
	"github.com/stretchr/testify/assert" // Use testify for assertions
)

func TestGetAllowedUsers(t *testing.T) {
	// Test case 1: Empty users slice
	originalUsers := users // Store the original users slice
	users = []db.User{}
	accounts := GetAllowedUsers()
	assert.Empty(t, accounts, "Accounts should be empty when users slice is empty")

	// Test case 2: Populated users slice
	users = originalUsers // Restore the original users slice
	accounts = GetAllowedUsers()
	assert.Len(t, accounts, len(users), "Accounts should have the same length as users slice")
	for _, user := range users {
		assert.Contains(t, accounts, user.Nickname, "Account should contain all nicknames")
		assert.Equal(t, user.Password, accounts[user.Nickname], "Password should match for each user")
	}

	// Test case 3: Check specific user
	assert.Equal(t, "user", users[0].Nickname)
	assert.Equal(t, "user", GetAllowedUsers()["user"])

}

func TestCreateUser(t *testing.T) {
	// Test case 1: Add a new user
	initialLength := len(users)
	CreateUser("newuser", "newpassword")
	assert.Len(t, users, initialLength+1, "Users slice should have one more element")
	newUser := users[len(users)-1]
	assert.Equal(t, "newuser", newUser.Nickname, "New user nickname should be correct")
	assert.Equal(t, "newpassword", newUser.Password, "New user password should be correct")

	// Test case 2: Add duplicate user (check if it's handled correctly or if an error is acceptable)
	CreateUser("newuser", "anotherpassword")                                          //Same user name
	assert.Len(t, users, initialLength+2, "Users slice should have one more element") // It's appending the same user name again
	newUser2 := users[len(users)-1]
	assert.Equal(t, "newuser", newUser2.Nickname, "Duplicate user nickname should be correct")
	assert.Equal(t, "anotherpassword", newUser2.Password, "Duplicate user password should be correct") //Password is different

	// Restore original user list after the tests
	users = []db.User{
		{Nickname: "user", Password: "user"},
		{Nickname: "test", Password: "test"},
		{Nickname: "prod", Password: "prod"},
	}
}

func TestGetUserByNickname(t *testing.T) {
	// Test case 1: User exists
	existingUser := "user"
	user := GetUserByNickname(existingUser)

	if user == nil {
		t.Errorf("GetUserByNickname(%s) returned nil, expected a user", existingUser)
	} else if user.Nickname != existingUser {
		t.Errorf("GetUserByNickname(%s) returned user with nickname %s, expected %s", existingUser, user.Nickname, existingUser)
	}

	// Test case 2: User does not exist
	nonExistingUser := "nonexistent"
	user = GetUserByNickname(nonExistingUser)

	if user != nil {
		t.Errorf("GetUserByNickname(%s) returned a user, expected nil", nonExistingUser)
	}

	// Test case 3: Empty nickname
	emptyNickname := ""
	user = GetUserByNickname(emptyNickname)

	if user != nil {
		t.Errorf("GetUserByNickname(%s) returned a user, expected nil", emptyNickname)
	}

	// Test case 4: Check if the returned user is a copy, not a reference to the original slice
	originalUsers := users // Keep a copy of the original users slice

	testNickname := "test_copy"
	testPassword := "password_copy"
	CreateUser(testNickname, testPassword) // Add a new user

	retrievedUser := GetUserByNickname(testNickname)
	if retrievedUser == nil {
		t.Errorf("GetUserByNickname(%s) returned nil, expected a user", testNickname)
	}

	retrievedUser.Password = "modified_password" // Modify the retrieved user's password

	originalRetrievedUser := GetUserByNickname(testNickname) //Get again the user

	if originalRetrievedUser.Password == "modified_password" {
		t.Errorf("GetUserByNickname returned a pointer to the original user in the slice. Should return a copy")
	}
	
	users = originalUsers

}
