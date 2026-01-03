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
