package service

import (
	"github.com/KartikVerma24/taskCli/domain/task"
)

func MapPriority(inputPriority string) (task.PriorityOfTask, error) {
	switch inputPriority {
	case "low", "l":
		return task.Low, nil
	case "medium", "m":
		return task.Medium, nil
	case "high", "h":
		return task.High, nil
	case "critical", "c":
		return task.Critical, nil
	default:
		return task.Medium, ErrPriorityInput
	}
}

func MapStatus(inputStatus string) (task.StatusOfTask, error) {
	switch inputStatus {
	case "todo": 
		return task.Todo, nil
	case "wip":
		return task.WIP, nil
	case "done":
		return task.Done, nil
	case "cancelled":
		return task.Cancelled, nil
	default:
		return task.Todo, ErrStatusInput
	}
}