package task

import (
	"fmt"
	"server/repository"
	"server/domain"
)

type Service struct {
	Storage repository.TaskStorageI
}

func NewService(storage repository.TaskStorageI) *Service{
	return &Service{
		Storage : storage,
	}
}

func (service *Service) Post(content string) (string, error){ //content это какое то поле/поля из структуры(распрасил request и записал в struct)
	task_id, err := service.Storage.PostProgress()

	if err != nil{
		return "", err
	}

	go service.makeAndPostTask(task_id, content)

	// здесь нужно еще добавить обработку ошибки и вернуть ее

	return task_id, nil
}

func (service *Service) makeAndPostTask(task_id string, content string){
	result := service.makeTask(content)
	service.Storage.PostResult(domain.Task{
		Id : task_id, 
		Result : result,
	})

	fmt.Println("Таска сохранена!")
}

func (service *Service) GetStatus(task_id string) (string, error){
	status, err := service.Storage.GetStatus(domain.Task{
		Id : task_id, 
		Result : "",
	})

	if err != nil{
		return "", err
	}

	return status, nil
}

func (service *Service) GetResult(task_id string) (string, error){
	result, err := service.Storage.GetResult(domain.Task{
		Id : task_id, 
		Result : "",
	})
	
	if err != nil{
		return "", err
	}

	return result, nil
}