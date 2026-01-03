package cli

type NewTaskCommand struct {
	priority    string
	description string
}

type ChangeTaskCommand struct {
	id          int
	newStatus   string
	newPriority string
	description string
}

type DoneTaskCommand struct {
	id int
}

type DeleteTaskCommand struct {
	id int
}
