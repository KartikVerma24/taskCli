package service

import "errors"

var (
	ErrPriorityInput = errors.New("invalid input for priority")
	ErrStatusInput   = errors.New("invalid input for status")
)
