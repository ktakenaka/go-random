package dto

import "gopkg.in/guregu/null.v4"

// CreateSample data transfer object
type CreateSample struct {
	Title   string
	Content string
	UserID  uint64
}

// UpdateSample data transfer object
type UpdateSample struct {
	ID      uint64
	Title   string
	Content null.String
	UserID  uint64
}
