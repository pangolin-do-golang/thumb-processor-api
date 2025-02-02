package thumb

import (
	"github.com/google/uuid"
	"testing"
)

func TestGetThumb(t *testing.T) {
	service := NewThumbService()

	expectedID := uuid.New()

	_, err := service.GetThumb(expectedID)

	if err != nil {
		t.Errorf("GetThumb returned an error: %v", err)
		return
	}
}
