package middleware

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestAuthMiddleware(t *testing.T) {
	// Test Cases:
	testCases := []struct {
		name             string
		username         string
		password         string
		allowedUsersFunc func() gin.Accounts
		expectedStatus   int
		expectedMessage  string
	}{
		{
			name:             "Valid Credentials",
			username:         "testuser",
			password:         "testpassword",
			allowedUsersFunc: func() gin.Accounts { return gin.Accounts{"testuser": "testpassword"} },
			expectedStatus:   http.StatusOK, // Or 200 if you're checking for a successful continuation
			expectedMessage:  "",            // No message expected on success
		},
		{
			name:             "Invalid Username",
			username:         "wronguser",
			password:         "testpassword",
			allowedUsersFunc: func() gin.Accounts { return gin.Accounts{"testuser": "testpassword"} },
			expectedStatus:   http.StatusUnauthorized,
			expectedMessage:  "Unauthorized",
		},
		{
			name:             "Invalid Password",
			username:         "testuser",
			password:         "wrongpassword",
			allowedUsersFunc: func() gin.Accounts { return gin.Accounts{"testuser": "testpassword"} },
			expectedStatus:   http.StatusUnauthorized,
			expectedMessage:  "Unauthorized",
		},
		{
			name:             "No Credentials Provided",
			username:         "",
			password:         "",
			allowedUsersFunc: func() gin.Accounts { return gin.Accounts{"testuser": "testpassword"} },
			expectedStatus:   http.StatusUnauthorized,
			expectedMessage:  "Unauthorized",
		},
		{
			name:             "Empty Accounts",
			username:         "testuser",
			password:         "testpassword",
			allowedUsersFunc: func() gin.Accounts { return gin.Accounts{} }, // Empty accounts
			expectedStatus:   http.StatusUnauthorized,
			expectedMessage:  "Unauthorized",
		},
		{
			name:             "Nil Accounts",
			username:         "testuser",
			password:         "testpassword",
			allowedUsersFunc: func() gin.Accounts { return nil }, // nil accounts
			expectedStatus:   http.StatusUnauthorized,
			expectedMessage:  "Unauthorized",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Setup
			r := gin.New()
			r.Use(AuthMiddleware(tc.allowedUsersFunc))

			// A dummy handler to simulate what happens after authentication
			var handlerCalled bool
			r.GET("/test", func(c *gin.Context) {
				handlerCalled = true
				c.Status(http.StatusOK) // Or whatever your success status is
			})

			req, _ := http.NewRequest("GET", "/test", nil)
			if tc.username != "" {
				req.SetBasicAuth(tc.username, tc.password)
			}

			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			// Assertions
			assert.Equal(t, tc.expectedStatus, w.Code)

			if tc.expectedMessage != "" {
				var response map[string]string
				err := json.Unmarshal(w.Body.Bytes(), &response)
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedMessage, response["message"])
			}

			if tc.expectedStatus == http.StatusOK { // Check if the handler was called only on success
				assert.True(t, handlerCalled, "Handler should be called on successful authentication")
			} else {
				assert.False(t, handlerCalled, "Handler should not be called on failed authentication")
			}
		})
	}
}
