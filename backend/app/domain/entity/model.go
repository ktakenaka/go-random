package entity

import (
	"time"
)

// Model base struct
type Model struct {
	ID        uint64    `csv:"-"`
	CreatedAt time.Time `csv:"登録日付"`
	UpdatedAt time.Time `csv:"更新日付"`
}
