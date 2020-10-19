package dto

import "time"

// CreateSample data transfer object
type CreateSample struct {
	Title   string
	Content *string
	UserID  uint64
}

// UpdateSample data transfer object
type UpdateSample struct {
	ID      uint64
	Title   string
	Content *string
	UserID  uint64
}

// ExportSample for csv export
type ExportSample struct {
	ID        uint64
	CreatedAt time.Time
	UpdatedAt time.Time
	Title     string
	Content   *string
	UserID    uint64
}

// ImportSample for csv import
type ImportSample struct {
	Title   string
	Content *string
	UserID  uint64
}
