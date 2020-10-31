package entity

// QueryBase for json api filters
type QueryBase struct {
	Page uint
	Sort []string
}

// SetPage setter for Page
func (q *QueryBase) SetPage(page uint) {
	q.Page = page
}

// SetSort setter for Sort
func (q *QueryBase) SetSort(sort []string) {
	q.Sort = sort
}
