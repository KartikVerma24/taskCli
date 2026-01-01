package database

import "errors"

var (
	ErrTaskNotFound = errors.New("no task found for provided id")
)