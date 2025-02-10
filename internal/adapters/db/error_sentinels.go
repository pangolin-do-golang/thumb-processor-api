package db

import "errors"

var (
	RequiredUserIDError = errors.New("logged user ID is required")
)
