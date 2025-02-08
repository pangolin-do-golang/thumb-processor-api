package entity

import "github.com/google/uuid"

const (
	ThumbProcessStatusPending  string = "pending"
	ThumbProcessStatusComplete string = "complete"
	ThumbProcessStatusFailed   string = "failed"
)

var AllowedProcessStatus = map[string]bool{
	ThumbProcessStatusPending:  true,
	ThumbProcessStatusComplete: true,
	ThumbProcessStatusFailed:   true,
}

type ThumbProcess struct {
	ID        uuid.UUID         `json:"id"`
	Video     ThumbProcessVideo `json:"video"`
	Status    string            `json:"status"`
	Error     string            `json:"error,omitempty"`
	Thumbnail ThumbProcessThumb `json:"thumbnail"`
}

type ThumbProcessVideo struct {
	Path string `json:"path"`
}

type ThumbProcessThumb struct {
	Path string `json:"url"`
}

func NewThumbProcess(url string) *ThumbProcess {
	return &ThumbProcess{
		Video: ThumbProcessVideo{
			Path: url,
		},
		Status: ThumbProcessStatusPending,
	}
}
