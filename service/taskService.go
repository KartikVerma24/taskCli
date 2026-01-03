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

var inValidTaskId = -1

func (s *TaskService) AddNewTask(content string, priority string) (int, error) {
	newTask, err := task.NewTask(content)
	if err != nil {
		return inValidTaskId, err
	}

	if priority != "" {
		mappedPriority, errMapping := MapPriority(strings.ToLower(priority))
		if errMapping != nil {
			return inValidTaskId, errMapping
		}
		newTask.SetPriority(mappedPriority)
	}

	savedTaskId, saveTaskErr := s.taskRepo.SaveTask(newTask)
	if saveTaskErr != nil {
		return inValidTaskId, saveTaskErr
	}

	return savedTaskId, nil
}

func (s *TaskService) ChangeTaskStatus(id int, changeStatus string) error {
	mappedStatus, errMapping := MapStatus(strings.ToLower(changeStatus))
	if errMapping != nil {
		return errMapping
	}

	taskFromRepo, getTaskErr := s.taskRepo.FindByID(id)
	if getTaskErr != nil {
		return getTaskErr
	}

	changeStatusErr := taskFromRepo.ChangeStatus(mappedStatus)
	if changeStatusErr != nil {
		return changeStatusErr
	}

	_, updateChangesErr := s.taskRepo.UpdateTask(taskFromRepo)
	if updateChangesErr != nil {
		return updateChangesErr
	}

	return nil
}

func (s *TaskService) ChangeTaskPriority(id int, changePriority string) error {
	mappedPriority, errMapping := MapPriority(strings.ToLower(changePriority))
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

	_, updateChangesErr := s.taskRepo.UpdateTask(taskFromRepo)
	if updateChangesErr != nil {
		return updateChangesErr
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

	_, updateChangesErr := s.taskRepo.UpdateTask(taskFromRepo)
	if updateChangesErr != nil {
		return updateChangesErr
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

func (s *TaskService) DeleteTask(id int) error {
	_, getTaskErr := s.taskRepo.FindByID(id)
	if getTaskErr != nil {
		return getTaskErr
	}

	deleteErr := s.taskRepo.Delete(id)
	if deleteErr != nil {
		return deleteErr
	}

	return nil
}
