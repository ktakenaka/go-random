package entity

import (
	"errors"
	"time"

	"github.com/ktakenaka/go-random/app/interface/persistence/mysql"
)

type Sample struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Title     string
}

func (s *Sample) TableName() string {
	return "samples"
}

var db = mysql.DBConnection()

func CreateSample(title string) error {
	now := time.Now()

	sample := Sample{
		Title:     title,
		CreatedAt: now,
		UpdatedAt: now,
	}

	db.Create(&sample)
	if db.NewRecord(sample) {
		return errors.New("failed to create Record")
	}
	return nil
}

func ListSamples() ([]Sample, error) {
	var samples []Sample
	err := db.Find(&samples).Error

	return samples, err
}

func FindSample(id string) (Sample, error) {
	var sample Sample
	err := db.First(&sample, id).Error

	return sample, err
}

func (s Sample) Save() error {
	err := db.Save(&s).Error
	return err
}

func (s Sample) Delete() error {
	err := db.Delete(&s).Error
	return err
}
