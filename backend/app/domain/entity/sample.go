package entity

import (
	"time"

	validator "github.com/go-playground/validator/v10"
)

type Sample struct {
	ID        int    `gorm:"primary_key"`
	Title     string `validate:"max=20, required"`
	Content   string `validate:"max=100"`
	UserID    uint64 `validate:required`
	User      User
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (s *Sample) TableName() string {
	return "samples"
}

func (s *Sample) Validate() error {
	validate = validator.New()
	err := validate.Struct(s)
	return err
}
