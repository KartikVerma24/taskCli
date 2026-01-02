package cli

import "errors"

var (
	ErrInvalidCommand      = errors.New("invalid command provided in input")
	ErrDescriptionRequired = errors.New("description is required for the command")
	ErrEmptyInputs         = errors.New("atleast one input is required for the command")
	ErrInvalidTask         = errors.New("invalid task, please select correct id. Use list-all to get all tasks")
)
