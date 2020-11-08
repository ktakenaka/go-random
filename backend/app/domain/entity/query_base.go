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
	Sort    map[string]string /*column: ASC or DESC*/
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

// AddWhereClause add conditions
// FIXME: shouldn't use gorm in entity for the point of clean architecture
func (q *QueryBase) AddWhereClause(columns []string, tx *gorm.DB) {
	for _, c := range columns {
		f, ok := q.Filters[c]
		if !ok {
			continue
		}

		if f.Condition == btwClause {
			tx.Where(
				c+" "+string(f.Condition),
				f.Value.([]string)[0],
				f.Value.([]string)[1],
			)
			continue
		}
		tx.Where(c+" "+string(f.Condition), f.Value)
	}
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
const (
	orderTypeASC  = "ASC"
	orderTypeDESC = "DESC"
)

// SetSort setter for Sort
func (q *QueryBase) SetSort(sort []string) {
	items := make(map[string]string)

	for _, item := range sort {
		if item == "" || item == "-" {
			continue
		}

		if strings.HasPrefix(item, "-") {
			items[item[1:]] = orderTypeDESC
			continue
		}

		items[item] = orderTypeASC
	}
	q.Sort = items
}

// ToOrderBy constructs ORDER BY clause
func (q QueryBase) ToOrderBy(columns []string) string {
	var clauses []string
	for _, c := range columns {
		if typ, ok := q.Sort[c]; ok {
			clauses = append(clauses, fmt.Sprintf("%s %s", c, typ))
		}
	}
	return strings.Join(clauses, ",")
}
