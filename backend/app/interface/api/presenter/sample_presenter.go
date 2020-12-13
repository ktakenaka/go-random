package presenter

import "time"

// SampleRequest for request
type SampleRequest struct {
	Title   string  `json:"title"`
	Content *string `json:"content"`
}

// SampleResponse for response
type SampleResponse struct {
	ID           string  `json:"id"`
	Title        string  `json:"title"`
	Content      *string `json:"content"`
	CreatedAtUTC string  `json:"created_at"`
	UpdatedAtUTC string  `json:"updated_at"`
}

// CreatedAt convert
func (s *SampleResponse) CreatedAt(createdAt time.Time) {
	s.CreatedAtUTC = createdAt.UTC().Format(time.UnixDate)
}

// UpdatedAt convert
func (s *SampleResponse) UpdatedAt(updatedAt time.Time) {
	s.UpdatedAtUTC = updatedAt.UTC().Format(time.UnixDate)
}

// SampleCSVPresenter for CSV import/export
type SampleCSVPresenter struct {
	ID           string  `csv:"ID"`
	CreatedAtStr string  `csv:"登録日付"`
	UpdatedAtStr string  `csv:"更新日付"`
	Title        string  `csv:"タイトル"`
	Content      *string `csv:"コンテント"`
	UserID       string  `csv:"ユーザーID"` // TODO: make this to UserName
}

// CreatedAt convert
func (s *SampleCSVPresenter) CreatedAt(createdAt time.Time) {
	s.CreatedAtStr = createdAt.String()
}

// UpdatedAt convert
func (s *SampleCSVPresenter) UpdatedAt(updatedAt time.Time) {
	s.UpdatedAtStr = updatedAt.String()
}
