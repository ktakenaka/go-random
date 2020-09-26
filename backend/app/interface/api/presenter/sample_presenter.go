package presenter

import (
	"time"
)

type SampleRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type SampleResponse struct {
	ID           uint64 `json:"id"`
	Title        string `json:"title"`
	Content      string `json:"content"`
	CreatedAtUTC string `json:"created_at"`
	UpdatedAtUTC string `json:"updated_at"`
}

func (s *SampleResponse) CreatedAt(createdAt time.Time) {
	s.CreatedAtUTC = createdAt.UTC().Format(time.UnixDate)
}

func (s *SampleResponse) UpdatedAt(updatedAt time.Time) {
	s.UpdatedAtUTC = updatedAt.UTC().Format(time.UnixDate)
}
