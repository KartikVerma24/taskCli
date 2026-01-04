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

var inValidID = -1

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

func (t *TaskJsonFSRepo) SaveTask(task *task.Task) (int, error) {
	t.mu.Lock()
	defer t.mu.Unlock()

	if task.GetId() == 0 {
		t.lastId++
		task.SetId(t.lastId)
	}

	t.data[task.GetId()] = task
	saveDataErr := t.saveDataToFile()
	if saveDataErr != nil {
		return inValidID, saveDataErr
	}

	return task.GetId(), nil
}

func (t *TaskJsonFSRepo) saveDataToFile() error {
	records := make([]taskRecordJson, 0, len(t.data))

	for _, t := range t.data {
		rec := JsonRecordFromTask(t)
		records = append(records, rec)
	}

	store := jsonData{
		LastId: t.lastId,
		Tasks:  records,
	}

	bytes, err := json.MarshalIndent(store, "", "  ")
	if err != nil {
		return err
	}

	return writeFileAtomic(t.filePath, bytes)
}

// func (t *TaskJsonFSRepo) UpdateTask(task *task.Task) (int, error) {
// 	t.mu.Lock()
// 	defer t.mu.Unlock()

// 	tId := task.GetId()

// 	t.data[tId] = task
// 	return tId, nil
// }

// func (t *TaskInMemRepo) FindByID(id int) (*task.Task, error) {
// 	task, exists := t.data[id]
// 	if !exists {
// 		return nil, database.ErrTaskNotFound
// 	}

// 	return task, nil
// }

// func (t *TaskInMemRepo) Delete(id int) error {
// 	t.mu.Lock()
// 	defer t.mu.Unlock()

// 	delete(t.data, id)
// 	return nil
// }

// func (t *TaskInMemRepo) FindAll() ([]*task.Task, error) {
// 	allTaskSlice := make([]*task.Task, 0)

// 	for _, val := range t.data {
// 		allTaskSlice = append(allTaskSlice, val)
// 	}

// 	return allTaskSlice, nil
// }
