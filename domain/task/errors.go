package task

import "errors"

var (
	ErrEmptyContent        = errors.New("task content can't be empty")
	ErrInvalidStatusChange = errors.New("invalid status transition")
	ErrTaskAlreadyDone     = errors.New("task already completed")
	ErrNoStatusToChange    = errors.New("current and new status are same")
	ErrNoPriorityToChange  = errors.New("current and new priority are same")
)
