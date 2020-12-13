package dto

import "time"

// CreateSample data transfer object
type CreateSample struct {
	Title   string
	Content *string
	UserID  string
}

// UpdateSample data transfer object
type UpdateSample struct {
	ID      string
	Title   string
	Content *string
	UserID  string
}

// ExportSample for csv export
type ExportSample struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
	Title     string
	Content   *string
	UserID    string
}

// ImportSample for csv import
type ImportSample struct {
	Title   string
	Content *string
	UserID  string
}
