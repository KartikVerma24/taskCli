package jsonfs

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/KartikVerma24/taskCli/domain/task"
)

type TaskJsonFSRepo struct {
	mu       sync.Mutex
	data     map[int]*task.Task
	lastId   int
	filePath string
}

type taskRecordJson struct {
	Id          int       `json:"id"`
	Content     string    `json:"content"`
	Status      int       `json:"status"`
	Priotity    int       `json:"priority"`
	StartedAt   time.Time `json:"startedAt"`
	CompletedAt time.Time `json:"completedAt"`
}

type jsonData struct {
	LastId int              `json:"lastID"`
	Tasks  []taskRecordJson `json:"tasks"`
}

func NewTaskJsonFSRepo(filePath string) (*TaskJsonFSRepo, error) {
	repo := &TaskJsonFSRepo{
		data:     make(map[int]*task.Task),
		lastId:   0,
		filePath: filePath,
	}

	// to check if the file exists or not
	_, existsErr := os.Stat(filePath)
	if os.IsNotExist(existsErr) {
		return nil, existsErr
	}

	// if it does, load the details
	errLoad := repo.loadFileData()
	if errLoad != nil {
		return nil, errLoad
	}

	return repo, nil
}

func (r *TaskJsonFSRepo) loadFileData() error {
	file, errOpenFile := os.Open(r.filePath)
	if errOpenFile != nil {
		return errOpenFile
	}

	var allTaskData jsonData

	decoder := json.NewDecoder(file)
	errDecode := decoder.Decode(&allTaskData)
	if errDecode != nil {
		return errDecode
	}

	for _, singleTaskJson := range allTaskData.Tasks {
		tempTask, err := TaskFromJsonRecord(singleTaskJson)
		if err != nil {
			fmt.Println("Data not loaded for task : ", tempTask.GetId())
			continue
		}
		r.data[tempTask.GetId()] = tempTask
	}

	r.lastId = allTaskData.LastId

	return nil
}
