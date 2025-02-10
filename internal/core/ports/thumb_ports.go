package ports

import "github.com/google/uuid"

type CreateProcessRequest struct {
	UserEmail string `json:"user_email"`
	Url       string `json:"url"`
}

type UpdateProcessRequest struct {
	ID            uuid.UUID `json:"id"`
	Status        string    `json:"status"`
	Error         string    `json:"error,omitempty"`
	ThumbnailPath string    `json:"thumbnailPath"`
}
