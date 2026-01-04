package jsonfs

import (
	"github.com/KartikVerma24/taskCli/domain/task"
)

func TaskFromJsonRecord(r taskRecordJson) (*task.Task, error) {
	t, rehydrateErr := task.RehydrateTask(
		r.Id,
		r.Content,
		task.StatusOfTask(r.Status),
		task.PriorityOfTask(r.Priotity),
		r.StartedAt,
		r.CompletedAt,
	)

	if rehydrateErr != nil {
		return nil, rehydrateErr
	}

	return t, nil
}

func JsonRecordFromTask(t *task.Task) (taskRecordJson) {
	var record taskRecordJson
	record.Id = t.GetId()
	record.Content = t.GetContent()
	record.Status = int(t.GetTaskStatus())
	record.Priotity = int(t.GetPriority())
	record.StartedAt = t.GetStartTime()
	
	if !t.GetCompletionTime().IsZero() {
		record.CompletedAt = t.GetCompletionTime()
	}

	return record
}