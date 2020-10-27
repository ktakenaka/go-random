package dto

// JSONAPIQuery for search results
type JSONAPIQuery struct {
	Sort   []string          `form:"sort"`
	Filter map[string]string `form:"filter"`
	Page   uint64            `form:"page"`
}
