package presenter

import (
	"time"
)

type Sample interface {
	GetID() int
	GetTitle() string
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
}

type SamplePresenter struct {
	ID        int `json:"id"`
	Title     string `json:"title"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func NewSamplePresenter (s Sample) SamplePresenter {
	return SamplePresenter {
		ID:        s.GetID(),
		Title:     s.GetTitle(),
		CreatedAt: s.GetCreatedAt().UTC().Format(time.UnixDate),
		UpdatedAt: s.GetUpdatedAt().UTC().Format(time.UnixDate),
	}
}
