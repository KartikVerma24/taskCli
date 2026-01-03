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
	case "list-all":
		return ListAllTasks(svc)
	case "change":
		parsedParam, parseErr := ChangeTaskParser(inputs[1:])
		if parseErr != nil {
			return parseErr
		}
		return ChangeTaskHandler(parsedParam, svc)
	case "done":
		parsedParam, parseErr := DoneTaskParser(inputs[1:])
		if parseErr != nil {
			return parseErr
		}
		return DoneTaskHandler(parsedParam, svc)
	case "delete":
		parsedParam, parseErr := DeleteTaskParser(inputs[1:])
		if parseErr != nil {
			return parseErr
		}
		return DeleteTaskHandler(parsedParam, svc)
	case "help":
		return PrintHelp()
	default:
		return InvalidCommandHandler()
	}
}
