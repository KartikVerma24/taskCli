package service

import (
	"strings"

	"github.com/KartikVerma24/taskCli/domain/task"
)

type TaskService struct {
	taskRepo task.Repo
}

type TaskView struct {
	Id             int
	Description    string
	Status         string
	Priority       string
	StartedTime    string
	CompletionTime string
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

func (s *TaskService) ChangeTaskPriorityStatus(priority string, id int, status string, flag string) error {
	var mappedPriority task.PriorityOfTask
	var mappedStatus task.StatusOfTask
	var errMapping error

	if strings.ToLower(flag) == "p" {
		mappedPriority, errMapping = MapPriority(strings.ToLower(priority))
		if errMapping != nil {
			return errMapping
		}
	} else {
		mappedStatus, errMapping = MapStatus(status)
		if errMapping != nil {
			return errMapping
		}
	}

	taskFromRepo, getTaskErr := s.taskRepo.FindByID(id)
	if getTaskErr != nil {
		return getTaskErr
	}

	if strings.ToLower(flag) == "p" {
		changePriorityErr := taskFromRepo.ChangePriority(mappedPriority)
		if changePriorityErr != nil {
			return changePriorityErr
		}
	} else {
		changeStatusErr := taskFromRepo.ChangeStatus(mappedStatus)
		if changeStatusErr != nil {
			return changeStatusErr
		}
	}

	_, saveChangesErr := s.taskRepo.SaveTask(taskFromRepo)
	if saveChangesErr != nil {
		return saveChangesErr
	}

	return nil
}

func (s *TaskService) CompleteTask(id int) error {
	taskFromRepo, getTaskErr := s.taskRepo.FindByID(id)
	if getTaskErr != nil {
		return getTaskErr
	}

	markDoneErr := taskFromRepo.MarkAsDone()
	if markDoneErr != nil {
		return markDoneErr
	}

	_, saveChangesErr := s.taskRepo.SaveTask(taskFromRepo)
	if saveChangesErr != nil {
		return saveChangesErr
	}

	return nil
}

func (s *TaskService) ListAllTasks() ([]TaskView, error) {
	allTasks, getListErr := s.taskRepo.FindAll()
	if getListErr != nil {
		return nil, getListErr
	}

	view := make([]TaskView, 0)
	for _, task := range allTasks {
		view = append(view, TaskView{
			Id:             task.GetId(),
			Description:    task.GetContent(),
			Status:         ReverseStatusMapping(task.GetTaskStatus()),
			Priority:       ReversePriorityMapping(task.GetPriority()),
			StartedTime:    GetTimeString(task.GetStartTime()),
			CompletionTime: GetTimeString(task.GetCompletionTime()),
		})
	}

	return view, nil
}
