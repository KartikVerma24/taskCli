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

func ListAllTasks(svc service.TaskService) error {
	allTasks, err := svc.ListAllTasks()
	if err != nil {
		return err
	}

	fmt.Println("Task list ===============> ")

	for _, val := range allTasks {
		fmt.Printf("%v | %s | %s | %s | %s | %s", val.Id, val.Description, val.Status, val.Priority, val.StartedTime, val.CompletionTime)
		fmt.Println()
	}

	return nil
}