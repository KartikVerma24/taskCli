package cli

import "flag"

func NewTaskParse(inputs []string) (*NewTaskCommand, error) {
	fs := flag.NewFlagSet("add", flag.ContinueOnError)
	desc := fs.String("desc", "", "description of the task being added")
	priority := fs.String("priority", "", "priority of the task")

	parsingErr := fs.Parse(inputs)
	if parsingErr != nil {
		return nil, parsingErr
	}

	if *desc == "" {
		return nil, ErrDescriptionRequired
	}

	return &NewTaskCommand{
		description: *desc,
		priority:    *priority,
	}, nil
}

func ChangeTaskParser(inputs []string) (*ChangeTaskCommand, error) {
	fs := flag.NewFlagSet("change", flag.ContinueOnError)
	id := fs.Int("id", 0, "task id for which changes need to be done")
	status := fs.String("status", "", "new status of the task")
	priority := fs.String("priority", "", "new priority of the task")

	parsingErr := fs.Parse(inputs)
	if parsingErr != nil {
		return nil, parsingErr
	}

	if *id == 0 {
		return nil, ErrInvalidTask
	}

	if *status == "" && *priority == "" {
		return nil, ErrEmptyInputs
	}

	return &ChangeTaskCommand{
		id:          *id,
		newStatus:   *status,
		newPriority: *priority,
	}, nil
}

func DoneTaskParser(inputs []string) (*DoneTaskCommand, error) {
	fs := flag.NewFlagSet("done", flag.ContinueOnError)
	id := fs.Int("id", 0, "task id which needs to be marked as done")

	parsingErr := fs.Parse(inputs)
	if parsingErr != nil {
		return nil, parsingErr
	}

	if *id == 0 {
		return nil, ErrInvalidTask
	}

	return &DoneTaskCommand{
		id: *id,
	}, nil
}

func DeleteTaskParser(inputs []string) (*DeleteTaskCommand, error) {
	fs := flag.NewFlagSet("delete", flag.ContinueOnError)
	id := fs.Int("id", 0, "task id which needs to be deleted")
	
	parsingErr := fs.Parse(inputs)
	if parsingErr != nil {
		return nil, parsingErr
	}

	if *id == 0 {
		return nil, ErrInvalidTask
	}

	return &DeleteTaskCommand{
		id: *id,
	}, nil
}