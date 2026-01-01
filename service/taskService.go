package service

import (
	"strings"

	"github.com/KartikVerma24/taskCli/domain/task"
)

type TaskService struct {
	taskRepo task.Repo
}

func NewTaskService(r task.Repo) *TaskService {
	return &TaskService{
		taskRepo: r,
	}
}

func (s *TaskService) AddNewTask(content string) (int, error) {
	newTask, err := task.NewTask(content)
	if err != nil {
		return -1, err
	}

	savedTaskId, saveTaskErr := s.taskRepo.SaveTask(newTask)
	if saveTaskErr != nil {
		return -1, saveTaskErr
	}

	return savedTaskId, nil
}

func (s *TaskService) ChangeTaskPriority(priority string, id int) error {
	mappedPriority, errMapping := MapPriority(strings.ToLower(priority))
	if errMapping != nil {
		return errMapping
	}

	taskFromRepo, getTaskErr := s.taskRepo.FindByID(id)
	if getTaskErr != nil {
		return getTaskErr
	}

	changePriorityErr := taskFromRepo.ChangePriority(mappedPriority)
	if changePriorityErr != nil {
		return changePriorityErr
	}

	return nil	
}

// func (s *TaskService) ChangesTaskStatus(status string, id int) error {

// }