package entity

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	defaultLimit = 20
)

// QueryBase for json api filters
type QueryBase struct {
	Page page
	Sort []sortItem
}

type page struct {
	Offset int
	Limit  int
}

// SetPage setter for Page
func (q *QueryBase) SetPage(param map[string]string) {
	offset, err := strconv.Atoi(param["offset"])
	if err != nil {
		q.Page.Offset = 0
	} else {
		q.Page.Offset = offset
	}

	limit, err := strconv.Atoi(param["limit"])
	if err != nil {
		q.Page.Limit = defaultLimit
	} else {
		q.Page.Limit = limit
	}
}

// GetLimit pagenation limit
func (q QueryBase) GetLimit() int {
	return q.Page.Limit
}

// GetOffset pagenation offset
func (q QueryBase) GetOffset() int {
	return q.Page.Offset
}

type sortItem string

// SetSort setter for Sort
func (q *QueryBase) SetSort(sort []string) {
	items := make([]sortItem, len(sort))
	for i, item := range sort {
		items[i] = sortItem(item)
	}
	q.Sort = items
}

// IsOrderByNeeded necessary
func (q QueryBase) IsOrderByNeeded() bool {
	return len(q.Sort) > 0
}

func (s sortItem) toOrderBy() string {
	if strings.HasPrefix(string(s), "-") {
		return fmt.Sprintf("%s DESC", s[1:])
	}
	return string(s)
}

// ToOrderBy constructs ORDER BY clause
func (q QueryBase) ToOrderBy() string {
	clauses := make([]string, len(q.Sort))
	for i := 0; i < len(q.Sort); i++ {
		clauses[i] = q.Sort[i].toOrderBy()
	}
	return strings.Join(clauses, ",")
}
