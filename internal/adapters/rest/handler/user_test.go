package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreateUserHandler(t *testing.T) {
	// Gin setup
	gin.SetMode(gin.TestMode)
	router := gin.New()
	RegisterUserRoutes(router)

	// Test cases
	tests := []struct {
		name             string
		requestBody      CreateUserRequest
		expectedStatus   int
		expectedResponse map[string]string
	}{
		{
			name: "Success",
			requestBody: CreateUserRequest{
				Nickname: "testuser",
				Password: "password123",
			},
			expectedStatus:   http.StatusCreated,
			expectedResponse: map[string]string{"message": "User created successfully"},
		},
		{
			name: "Validation Error - Empty Nickname",
			requestBody: CreateUserRequest{
				Password: "password123",
			},
			expectedStatus:   http.StatusBadRequest,
			expectedResponse: map[string]string{"error": "Key: 'CreateUserRequest.Nickname' Error:Field validation for 'Nickname' failed on the 'required' tag"},
		},
		{
			name: "Validation Error - Empty Password",
			requestBody: CreateUserRequest{
				Nickname: "testuser",
			},
			expectedStatus:   http.StatusBadRequest,
			expectedResponse: map[string]string{"error": "Key: 'CreateUserRequest.Password' Error:Field validation for 'Password' failed on the 'required' tag"},
		},
		{
			name:             "Validation Error - Both Empty",
			requestBody:      CreateUserRequest{},
			expectedStatus:   http.StatusBadRequest,
			expectedResponse: map[string]string{"error": "Key: 'CreateUserRequest.Nickname' Error:Field validation for 'Nickname' failed on the 'required' tag\nKey: 'CreateUserRequest.Password' Error:Field validation for 'Password' failed on the 'required' tag"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create the request body in JSON format
			requestBodyJSON, _ := json.Marshal(tt.requestBody)
			req, _ := http.NewRequest("POST", "/user", bytes.NewBuffer(requestBodyJSON))
			req.Header.Set("Content-Type", "application/json")

			// Simulate the request
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			// Check the status code
			assert.Equal(t, tt.expectedStatus, w.Code)

			// Check the response body
			var response map[string]string
			json.Unmarshal(w.Body.Bytes(), &response)
			assert.Equal(t, tt.expectedResponse, response)
		})
	}
}
