package entity

import (
	"time"

	"github.com/ktakenaka/go-random/app/config"
)

type Sample struct {
  ID        uint `gorm:"primary_key"`
  CreatedAt time.Time
  UpdatedAt time.Time
  Title     string
}

func (b *Sample) TableName() string {
	return "samples"
}

// TODO: move this to repository
func CreateSample(title string) {
	sample := Sample{Title: title, CreatedAt: time.Now(), UpdatedAt: time.Now()}
	db := config.DBConnection()
	db.Create(&sample)
}