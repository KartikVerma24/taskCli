package task

import (
	"time"

	"github.com/KartikVerma24/taskCli/domain"
)

type StatusOfTask int
type PriorityOfTask int

const (
	Todo StatusOfTask = iota
	WIP
	Done
	Cancelled
)

const (
	Low PriorityOfTask = iota
	Medium
	High
	Critical
)

// Entity
type Task struct {
	id          int
	content     string
	status      StatusOfTask
	priotity    PriorityOfTask
	startedAt   time.Time
	completedAt time.Time
}

// Entity Behaviours

// constructors
func NewTask(content string) (*Task, error) {
	if content == "" {
		return nil, domain.ErrEmptyContent
	}

	return &Task{
		content:   content,
		status:    Todo,
		priotity:  Medium,
		startedAt: time.Now(),
	}, nil
}

func RehydrateTask(id int, content string, status StatusOfTask, priority PriorityOfTask, startedAt time.Time, completedAt time.Time) (*Task, error) {
	if content == "" {
		return nil, domain.ErrEmptyContent
	}

	if id == 0 {
		return nil, domain.ErrInvalidTask
	}

	return &Task{
		id:          id,
		content:     content,
		status:      status,
		priotity:    priority,
		startedAt:   startedAt,
		completedAt: completedAt,
	}, nil
}

// state changing methods
func (t *Task) ChangeStatus(newStatus StatusOfTask) error {
	if t.status == newStatus {
		return domain.ErrNoStatusToChange
	}

	if !isValidStatusTransition(t.status, newStatus) {
		return domain.ErrInvalidStatusChange
	}

	t.status = newStatus

	if t.status == Done {
		t.completedAt = time.Now()
	}

	return nil
}

func (t *Task) MarkAsDone() error {
	return t.ChangeStatus(Done)
}

func (t *Task) ChangePriority(newPriority PriorityOfTask) error {
	if t.priotity == newPriority {
		return domain.ErrNoPriorityToChange
	}

	t.priotity = newPriority
	return nil
}

func (t *Task) SetId(id int) error {
	t.id = id
	return nil
}

func (t *Task) SetPriority(newPriority PriorityOfTask) error {
	t.priotity = newPriority
	return nil
}

// state getting methods
func (t *Task) GetId() int {
	return t.id
}

func (t *Task) GetTaskStatus() StatusOfTask {
	return t.status
}

func (t *Task) GetPriority() PriorityOfTask {
	return t.priotity
}

func (t *Task) GetContent() string {
	return t.content
}

func (t *Task) GetStartTime() time.Time {
	return t.startedAt
}

func (t *Task) GetCompletionTime() time.Time {
	return t.completedAt
}
