package storage

import "errors"

var (
	ErrUserExists    = errors.New("user already exists")
	ErrNoRecordFound = errors.New("no records found")
)
