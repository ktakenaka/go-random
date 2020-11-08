package dto

import (
	"strings"

	"github.com/mitchellh/mapstructure"
)

// JSONAPIQuery for search results
type JSONAPIQuery struct {
	Sort   []string          `form:"sort"`
	Filter map[string]string `form:"filter"`
	Page   map[string]string `form:"page"`
}

type ginContext interface {
	Query(key string) string
	QueryMap(key string) map[string]string
}

// NewJSONAPIQueryFromContext constructor
func NewJSONAPIQueryFromContext(ctx ginContext) (*JSONAPIQuery, error) {
	var q JSONAPIQuery

	sort := ctx.Query("sort")
	if sort != "" {
		q.Sort = strings.Split(sort, ",")
	}

	q.Filter = ctx.QueryMap("filter")
	q.Page = ctx.QueryMap("page")
	return &q, nil
}

type queryBase interface {
	SetFilters(filter map[string]string)
	SetPage(page map[string]string)
	SetSort(sort []string)
}

// Bind convert to the format of query
func (q JSONAPIQuery) Bind(query queryBase) error {
	err := mapstructure.Decode(q.Filter, query)
	if err != nil {
		return err
	}

	query.SetFilters(q.Filter)
	query.SetPage(q.Page)
	query.SetSort(q.Sort)
	return nil
}
