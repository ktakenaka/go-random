package dto

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
	Content string
	UserID  uint64
}
