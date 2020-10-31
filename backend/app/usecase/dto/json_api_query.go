package dto

import (
	"github.com/mitchellh/mapstructure"
)

// JSONAPIQuery for search results
type JSONAPIQuery struct {
	Sort   []string          `form:"sort"`
	Filter map[string]string `form:"filter"`
	Page   uint              `form:"page"`
}

type queryBase interface {
	SetPage(page uint)
	SetSort(sort []string)
}

// Bind convert to the format of query
func (q JSONAPIQuery) Bind(query queryBase) error {
	err := mapstructure.Decode(q.Filter, query)
	if err != nil {
		return err
	}

	if q.Page == 0 {
		query.SetPage(1)
	} else {
		query.SetPage(q.Page)
	}

	query.SetSort(q.Sort)
	return nil
}
