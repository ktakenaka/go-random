package entity

import (
	"time"
)

type Sample struct {
	ID        int       `gorm:"primary_key" json:"id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (s *Sample) TableName() string {
	return "samples"
}

func NewSample(id int, title string) *Sample {
	now := time.Now()
	return &Sample{
		Title: title, CreatedAt: now, UpdatedAt: now,
	}
}

func (s *Sample) GetID() int {
	return s.ID
}

func (s *Sample) GetTitle() string {
	return s.Title
}
