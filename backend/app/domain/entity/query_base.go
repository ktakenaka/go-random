package entity

import (
	"fmt"
	"strconv"
	"strings"

	"gorm.io/gorm"
)

const (
	defaultLimit = 20
)

// QueryBase for json api filters
type QueryBase struct {
	Filters map[string]filter
	Page    page
	Sort    []sortItem
}

// ========== Filters ==========
type filter struct {
	Condition condition
	Value     value
}

type condition string
type value interface{}

const (
	likeClause  = condition("LIKE ?")
	nlikeClause = condition("NOT LIKE ?")
	eqClause    = condition("= ?")
	neClause    = condition("<> ?")
	geClause    = condition(">= ?")
	gtClause    = condition("> ?")
	leClause    = condition("<=")
	ltClause    = condition("<")
	inClause    = condition("IN ?")
	ninClause   = condition("NOT IN ?")
	btwClause   = condition("BETWEEN ? AND ?")
)

var conditions = map[string]condition{
	"like":  likeClause,
	"nlike": nlikeClause,
	"eq":    eqClause,
	"ne":    neClause,
	"ge":    geClause,
	"gt":    gtClause,
	"le":    leClause,
	"lt":    ltClause,
	"in":    inClause,
	"nin":   ninClause,
	"btw":   btwClause,
}

func (c condition) sqlValue(v string) (value /*SQL clause*/, bool /*ok*/) {
	if c == btwClause {
		values := strings.Split(v, ",")
		if len(values) != 2 || values[0] == "" || values[1] == "" {
			return "", false
		}
		return values /*[]string, length == 2*/, true
	}

	if c == inClause || c == ninClause {
		values := strings.Split(v, ",")

		var cleaned []string
		for _, v := range values {
			if v != "" {
				cleaned = append(cleaned, v)
			}
		}

		if len(cleaned) == 0 {
			return value(""), false
		}
		return cleaned /*[]string*/, true
	}

	if c == likeClause || c == nlikeClause {
		return "%" + v + "%", true
	}
	return v, true
}

// SetFilters converts map[string]string to []filter
func (q *QueryBase) SetFilters(params map[string]string) {
	result := make(map[string]filter)

	for k, v := range params {
		f := strings.SplitN(v, ":", 2)
		if len(f) != 2 || f[0] == "" || f[1] == "" {
			continue
		}

		condition, ok := conditions[f[0]]
		if !ok {
			continue
		}

		sqlv, ok := condition.sqlValue(f[1])
		if !ok {
			continue
		}

		result[k] = filter{Condition: condition, Value: sqlv}
	}

	q.Filters = result
}

// FIXME: shouldn't use gorm in entity for the point of clean architecture
// It may be one option to move these logics to repository or pkg
type gormDB interface {
	Where(query interface{}, args ...interface{}) (tx *gorm.DB)
}

// AddClause add conditions
// TODO: enable to convert other types, int, time...
func (q *QueryBase) AddClause(column string, db gormDB) {
	f, ok := q.Filters[column]
	if !ok {
		return
	}

	if f.Condition == btwClause {
		db.Where(
			column+" "+string(f.Condition),
			f.Value.([]string)[0],
			f.Value.([]string)[1],
		)
		return
	}
	db.Where(column+" "+string(f.Condition), f.Value)
}

// ========== Page ==========
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

// ========== Sort ==========
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

// TODO: prevent from SQL Injection
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
