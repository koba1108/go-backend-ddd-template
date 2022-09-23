package model

import "fmt"

var (
	ErrSortKeyRequired = fmt.Errorf("sort key is required")
	ErrLimitRequired   = fmt.Errorf("limit key is required")
	ErrOffsetRequired  = fmt.Errorf("offset key is required")
)

type Pagination struct {
	SortKey string `json:"sortKey"`
	IsAsc   bool   `json:"isAsc"`
	Limit   int    `json:"limit"`
	Offset  int    `json:"offset"`
}

func NewPagination(limit, offset int, sortKey string, isAsc bool) (*Pagination, error) {
	if sortKey == "" {
		return nil, ErrSortKeyRequired
	}
	if limit < 1 {
		return nil, ErrLimitRequired
	}
	if offset < 1 {
		return nil, ErrOffsetRequired
	}
	return &Pagination{
		Limit:   limit,
		Offset:  offset,
		IsAsc:   isAsc,
		SortKey: sortKey,
	}, nil
}
