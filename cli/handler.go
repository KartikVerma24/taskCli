package cli

import (
	"fmt"

	"github.com/KartikVerma24/taskCli/service"
)

func NewTaskHandler(t *NewTaskCommand, svc service.TaskService) error {
	taskId, svcErr := svc.AddNewTask(t.description, t.priority)
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

func ChangeTaskHandler(t *ChangeTaskCommand, svc service.TaskService) error {
	if t.newStatus != "" {
		changeStatusErr := svc.ChangeTaskStatus(t.id, t.newStatus)
		if changeStatusErr != nil {
			return changeStatusErr
		}
		fmt.Println("Changed status of task : ", t.id)
	}

	if t.newPriority != "" {
		changePriorityErr := svc.ChangeTaskPriority(t.id, t.newPriority)
		if changePriorityErr != nil {
			return changePriorityErr
		}
		fmt.Println("Changed priority of task : ", t.id)
	}

	return nil
}

func DoneTaskHandler(t *DoneTaskCommand, svc service.TaskService) error {
	markDoneErr := svc.CompleteTask(t.id)
	if markDoneErr != nil {
		return markDoneErr
	}

	fmt.Printf("task %v is done", t.id)
	fmt.Println()
	return nil
}