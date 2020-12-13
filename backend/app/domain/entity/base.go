package entity

import (
	"crypto/rand"
	"time"

	"gorm.io/gorm"

	ulid "github.com/oklog/ulid/v2"
)

// ID primary key of tables
type ID string

// IsZero check if it's a zero value
func (id ID) IsZero() bool {
	return string(id) == ""
}

// String string implement interface of Stringer
func (id ID) String() string {
	return string(id)
}

// Base contains basic information of a entity
type Base struct {
	ID ID
}

// BeforeCreate define hook used by GORM
func (b *Base) BeforeCreate(_ *gorm.DB) error {
	if !b.ID.IsZero() {
		return nil
	}

	t := time.Now()
	entropy := ulid.Monotonic(rand.Reader, 0)
	id := ulid.MustNew(ulid.Timestamp(t), entropy)
	b.ID = ID(id.String())
	return nil
}
