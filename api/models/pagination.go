package models

import (
	"errors"
	"strconv"
)

var (
	ErrInvalidPageParameter  = errors.New("invalid page parameter")
	ErrInvalidLimitParameter = errors.New("invalid limit parameter")
)

// Define um limite máximo seguro
const (
	DefaultLimit = 10
	MaxLimit     = 1000
)

type Pagination struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

func NewPagination(pageStr, limitStr string) *Pagination {
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		limit = DefaultLimit
	}

	if limit > MaxLimit {
		limit = MaxLimit
	}

	return &Pagination{
		Page:  page,
		Limit: limit,
	}
}

type PaginatedResponse struct {
	Data       any   `json:"data"`
	Total      int64 `json:"total"`
	TotalPages int   `json:"totalPages"`
	Page       int   `json:"page"`
	Limit      int   `json:"limit"`
}
