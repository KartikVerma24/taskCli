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
