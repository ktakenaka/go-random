package entity

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type Sample struct {
	ID        int    `gorm:"primary_key"`
	Title     string `validate:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
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

var validate *validator.Validate

func (s *Sample) Validate() error {
	validate = validator.New()
	err := validate.Struct(s)
	return err
}
