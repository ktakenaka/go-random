package dto

import (
	"github.com/mitchellh/mapstructure"
)

// JSONAPIQuery for search results
type JSONAPIQuery struct {
	Sort   []string          `form:"sort"`
	Filter map[string]string `form:"filter"`
	Page   map[string]string `form:"page"`
}

type queryBase interface {
	SetPage(page map[string]string)
	SetSort(sort []string)
}

// Bind convert to the format of query
func (q JSONAPIQuery) Bind(query queryBase) error {
	err := mapstructure.Decode(q.Filter, query)
	if err != nil {
		return err
	}

	query.SetPage(q.Page)
	query.SetSort(q.Sort)
	return nil
}
