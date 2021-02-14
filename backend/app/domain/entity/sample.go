package entity

import (
	"time"
)

// Sample entity
type Sample struct {
	Base
	CreatedAt time.Time
	UpdatedAt time.Time
	Title     string  `validate:"required,max=20,min=5"`
	Content   *string `validate:"max=100"`
	UserID    string  `validate:"required"`
}

// SampleQuery sql filter
type SampleQuery struct {
	QueryBase
}

// Validate with validator v10
func (s *Sample) Validate() error {
	err := validate.Struct(s)
	return err
}

// Equal trial implementation of goderive
// make gen
func (s *Sample) Equal(s2 *Sample) bool {
	return deriveEqual(s, s2)
}
