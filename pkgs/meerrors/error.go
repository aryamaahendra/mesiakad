package meerrors

import (
	"errors"
	"fmt"
)

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrUnuthorized    = errors.New("unauthorized")
)

type ErrNotUniqueInDB struct {
	Column string
}

func (err *ErrNotUniqueInDB) Error() string {
	return fmt.Sprintf("%s already in used", err.Column)
}

func NewErrNotUniqueInDB(column string) *ErrNotUniqueInDB {
	return &ErrNotUniqueInDB{Column: column}
}
