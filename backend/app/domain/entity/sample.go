package entity

import (
	"fmt"
	"reflect"
	"time"

	validator "github.com/go-playground/validator/v10"
)

// Sample entity
type Sample struct {
	ID        uint64 `gorm:"primary_key"`
	Title     string `validate:"max=20,required" csv:"タイトル"`
	Content   string `validate:"max=100" csv:"コンテント"`
	UserID    uint64 `validate:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Validate with validator v10
func (s *Sample) Validate() error {
	validate := validator.New()
	err := validate.Struct(s)
	return err
}

// Tags trial implementation of tags
func (s *Sample) Tags() {
	rt, rv := reflect.TypeOf(*s), reflect.ValueOf(*s)

	for i := 0; i < rt.NumField(); i++ {
		field := rt.Field(i)
		csv := field.Tag.Get("csv")

		fmt.Printf("[Tag] csv:%s\n", csv)
		fmt.Printf("[Value] %s\n\n", rv.Field(i).Interface())
	}
}
