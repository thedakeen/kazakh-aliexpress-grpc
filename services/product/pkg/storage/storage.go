package storage

import "errors"

var (
	ErrNoRecordFound  = errors.New("no records found")
	ErrCategoryExists = errors.New("category already exists")
)
