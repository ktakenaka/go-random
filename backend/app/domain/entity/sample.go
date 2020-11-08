package entity

import (
	"time"

	validator "github.com/go-playground/validator/v10"
)

// Sample entity
type Sample struct {
	ID        uint64
	CreatedAt time.Time
	UpdatedAt time.Time
	Title     string  `validate:"required,max=20"`
	Content   *string `validate:"max=100"`
	UserID    uint64  `validate:"required"`
}

// SampleQuery sql filter
type SampleQuery struct {
	Title   string `column:"title"`
	Content string `column:"content"`

	QueryBase
}

// Validate with validator v10
func (s *Sample) Validate() error {
	validate := validator.New()
	err := validate.Struct(s)
	return err
}

// Equal trial implementation of goderive
// make gen
func (s *Sample) Equal(s2 *Sample) bool {
	return deriveEqual(s, s2)
}
