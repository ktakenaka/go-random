package entity

import (
	"fmt"
	"strings"
)

// QueryBase for json api filters
type QueryBase struct {
	Page uint
	Sort []sortItem
}

type sortItem string

func (s sortItem) ToOrderBy() string {
	if strings.HasPrefix(string(s), "-") {
		return fmt.Sprintf("%s DESC", s[1:])
	}
	return string(s)
}

// SetPage setter for Page
func (q *QueryBase) SetPage(page uint) {
	q.Page = page
}

// SetSort setter for Sort
func (q *QueryBase) SetSort(sort []string) {
	items := make([]sortItem, len(sort))
	for i, item := range sort {
		items[i] = sortItem(item)
	}
	q.Sort = items
}
