package presenter

import (
	"time"
)

type sample interface {
	GetID() int
	GetTitle() string
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
}

type SampleResponse struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func NewSampleResponse(s sample) SampleResponse {
	return SampleResponse{
		ID:        s.GetID(),
		Title:     s.GetTitle(),
		CreatedAt: s.GetCreatedAt().UTC().Format(time.UnixDate),
		UpdatedAt: s.GetUpdatedAt().UTC().Format(time.UnixDate),
	}
}
