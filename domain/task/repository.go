package task

type Repo interface {
	SaveTask(task *Task) (int, error)
	FindByID(id int) (*Task, error)
	Delete(id int) error
	FindAll() ([]*Task, error)
	UpdateTask(task *Task) (int, error)
}