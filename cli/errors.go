package cli

import "errors"

var (
	ErrInvalidCommand      = errors.New("invalid command provided in input")
	ErrDescriptionRequired = errors.New("description is required for the command")
)
