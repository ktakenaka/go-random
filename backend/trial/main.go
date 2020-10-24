package main

import (
	"container/list"
	"fmt"
	"reflect"
	"time"

	"gopkg.in/guregu/null.v4"
)

// Model base struct
type Model struct {
	ID        uint64    `csv:"-"`
	CreatedAt time.Time `csv:"登録日付"`
	UpdatedAt time.Time `csv:"更新日付"`
}

// Sample entity
type Sample struct {
	Model
	Title   string      `validate:"max=20,required" csv:"タイトル"`
	Content null.String `validate:"max=100" csv:"コンテント"`
	UserID  uint64      `validate:"required" csv:"ユーザーID"`
	Hoge    *int        `csv:"fuga"`
}

// Tag => Lookup, Get
func iterateFields(val reflect.Value, res map[string]interface{}) {
	for i := 0; i < val.Type().NumField(); i++ {
		field := val.Field(i)
		typ := val.Type().Field(i)
		csvTag := typ.Tag.Get("csv")

		fmt.Println(field)

		if csvTag == "-" {
			continue
		} else if csvTag == "" && field.Kind() == reflect.Struct {
			iterateFields(field, res)
		} else {
			switch field.Kind() {
			case reflect.Struct:
				res[csvTag] = field.Interface()
			case reflect.Ptr:
				res[csvTag] = field.Elem().Interface()
			default:
				res[csvTag] = field.Interface()
			}
		}
	}
	fmt.Println(res)
}

// NewListElm constructor for list Element
func NewListElm(v interface{}) *list.Element {
	return &list.Element{Value: v}
}

func main() {
	hoge := 1
	s := reflect.ValueOf(Sample{
		Model:  Model{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		Title:  "title",
		UserID: 1,
		Hoge:   &hoge,
	})
	res := make(map[string]interface{})
	iterateFields(s, res)

	fmt.Println("===== linked list =====")
	lis := list.New()
	el1 := NewListElm("el1")
	el2 := NewListElm("el2")
	el3 := NewListElm("el3")
	lis.PushBack(el1)
	lis.PushBack(el2)
	lis.InsertBefore("Hello", el1)
	lis.PushFront(el3)

	lis.MoveAfter(el2, el3)

	var e *list.Element
	e = lis.Front()
	for i := 0; i < lis.Len(); i++ {
		fmt.Printf("iteration %v\n", i)
		fmt.Println(fmt.Sprintf("%v", e.Value))
		e = e.Next()
	}

	fmt.Println("after refactoring")
	var el *list.Element
	for el = lis.Front(); el != nil; el = el.Next() {
		fmt.Println(el.Value)
	}
}
