package ram

import (
	"server/usecases/service"
	"server/domain"
)

type TaskStorage struct{
	taskStatus map[string]string //статус таски по ее task_id
	taskResult map[string]string //результат таски по ее task_id [пока результат это просто какой то мусор]
}

func NewTaskStorage() *TaskStorage{
	return &TaskStorage{
		taskStatus: make(map[string]string), 
		taskResult: make(map[string]string), 
	}
}

func (storage *TaskStorage) PostProgress() (string, error){
	taskId, err := service.GenerateUUID()
	if err != nil{
		return "", err
	}

	storage.taskStatus[taskId] = "in_progress"
	return taskId, nil
}

func (storage *TaskStorage) PostResult(task domain.Task){
	storage.taskStatus[task.Id] = "ready"
	storage.taskResult[task.Id] = task.Result
}

func (storage *TaskStorage) GetStatus(task domain.Task) (string, error){
	if status, ok := storage.taskStatus[task.Id]; ok {
		return status, nil
	} else{
		return "", TaskNotFound
	}
}

func (storage *TaskStorage) GetResult(task domain.Task) (string, error){
	status, _ := storage.GetStatus(task)
	if status == "in_progress"{
		return "", nil
	}

	if result, ok := storage.taskResult[task.Id]; ok {
		return result, nil
	} else{
		return "", TaskNotFound
	}
}

