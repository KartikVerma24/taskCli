package service

import (
	"time"

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

var timeFormat = "02-Jan-2006 15:04:05"

func ReverseStatusMapping(status task.StatusOfTask) string {
	switch status {
	case task.Todo:
		return "Todo"
	case task.Done:
		return "Done"
	case task.WIP:
		return "WIP"
	case task.Cancelled:
		return "Cancelled"
	default:
		return "invalid status"
	}
}

func ReversePriorityMapping(priority task.PriorityOfTask) string {
	switch priority {
	case task.Low:
		return "Low"
	case task.Medium:
		return "Medium"
	case task.High:
		return "High"
	case task.Critical:
		return "Critical"
	default:
		return "invalid priority"
	}
}

func GetTimeString(taskTime time.Time) string {
	if taskTime.IsZero() {
		return "---"
	}
	return taskTime.Format(timeFormat)
}
