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
