package user

import (
	"server/repository"
	"server/domain"
)

type Service struct {
	Storage repository.UserStorageI
}

func NewService(storage repository.UserStorageI) *Service{
	return &Service{
		storage,
	}
}

func (service *Service) PostRegister(user domain.User) error{ //content это какое то поле/поля из структуры(распрасил request и записал в struct)
	err := service.Storage.PostRegister(user)

	if err != nil{
		return err
	}

	return nil
}

func (service *Service) PostLogin(user domain.User) (string, error){ //content это какое то поле/поля из структуры(распрасил request и записал в struct)
	sessionId, err := service.Storage.PostLogin(user)

	if err != nil{
		return "", err
	}

	return sessionId, nil
}