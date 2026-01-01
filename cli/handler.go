package cli

import (
	"fmt"

	"github.com/KartikVerma24/taskCli/service"
)

func NewTaskHandler(t *NewTaskCommand, svc service.TaskService) error {
	taskId, svcErr := svc.AddNewTask(t.description)
	if svcErr != nil {
		return svcErr
	}

	fmt.Println("Added new task with id : ", taskId)
	return nil
}
