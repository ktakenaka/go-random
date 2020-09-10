package presenter

import (
	"time"
)

type SampleRequest struct {
	Title string `json:"title"`
}

type SampleResponse struct {
	ID           int    `json:"id"`
	Title        string `json:"title"`
	CreatedAtUTC string `json:"created_at"`
	UpdatedAtUTC string `json:"updated_at"`
}

func (s *SampleResponse) CreatedAt(createdAt time.Time) {
	s.CreatedAtUTC = createdAt.UTC().Format(time.UnixDate)
}

func (s *SampleResponse) UpdatedAt(updatedAt time.Time) {
	s.UpdatedAtUTC = updatedAt.UTC().Format(time.UnixDate)
}
