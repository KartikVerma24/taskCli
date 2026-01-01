package inmemory

import (
	"sync"

	"github.com/KartikVerma24/taskCli/database"
	"github.com/KartikVerma24/taskCli/domain/task"
)

type TaskInMemRepo struct {
	mu     sync.Mutex
	data   map[int]*task.Task
	nextID int
}

func NewTaskInMemRepo() *TaskInMemRepo {
	return &TaskInMemRepo{
		data:   make(map[int]*task.Task),
		nextID: 1,
	}
}

func (t *TaskInMemRepo) SaveTask(task *task.Task) (int, error) {
	t.mu.Lock()
	defer t.mu.Unlock()

	taskId := t.nextID
	t.nextID++

	t.data[taskId] = task
	return taskId, nil
}

func (t *TaskInMemRepo) FindByID(id int) (*task.Task, error) {
	task, exists := t.data[id]
	if !exists {
		return nil, database.ErrTaskNotFound
	}

	return task, nil
}

func (t *TaskInMemRepo) Delete(id int) error {
	t.mu.Lock()
	defer t.mu.Unlock()

	_, exists := t.data[id]
	if !exists {
		return database.ErrTaskNotFound
	}

	delete(t.data, id)
	return nil
}
