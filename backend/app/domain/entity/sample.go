package entity

import (
	"fmt"
	"reflect"
	"time"

	validator "github.com/go-playground/validator/v10"
	"gopkg.in/guregu/null.v4"
)

// Sample entity
type Sample struct {
	ID        uint64      `csv:"-"`
	CreatedAt time.Time   `csv:"登録日付"`
	UpdatedAt time.Time   `csv:"更新日付"`
	Title     string      `validate:"max=20,required" csv:"タイトル"`
	Content   null.String `validate:"max=100" csv:"コンテント"`
	UserID    uint64      `validate:"required" csv:"ユーザーID"`
	Hoge      *int        `csv:"fuga"`
}

// Validate with validator v10
func (s *Sample) Validate() error {
	validate := validator.New()
	err := validate.Struct(s)
	return err
}

// CSVTags trial implementation of tags
func (s *Sample) CSVTags() {
	rt, rv := reflect.TypeOf(s), reflect.ValueOf(s)

	for i := 0; i < rt.NumField(); i++ {
		field := rt.Field(i)
		csv := field.Tag.Get("csv")

		fmt.Printf("[Tag] csv:%s\n", csv)
		fmt.Printf("[Value] %s\n\n", rv.Field(i).Interface())
	}
}

// Equal trial implementation of goderive
// make gen
func (s *Sample) Equal(s2 *Sample) bool {
	return deriveEqual(s, s2)
}
