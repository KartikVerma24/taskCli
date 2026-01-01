package cli

import (
	"github.com/KartikVerma24/taskCli/service"
)

func RunCommands(inputs []string, svc service.TaskService) error {
	if len(inputs) == 0 {
		return ErrInvalidCommand
	}

	cmd := inputs[0]
	switch cmd {
	case "add":
		parsedParam, parseErr := NewTaskParse(inputs[1:])
		if parseErr != nil {
			return parseErr
		}
		return NewTaskHandler(parsedParam, svc)
	default:
		return ErrInvalidCommand
	}
}
