package ram

import (
	"server/usecases/service"
	"server/domain"
)

type UserStorage struct{
	user map[string]domain.User //хранит domain.User по его id
	session map[string]domain.Session //хранит domain.Session по ее id
}

func NewUserStorage() *UserStorage{
	return &UserStorage{
		user: make(map[string]domain.User), 
		session: make(map[string]domain.Session),
	}
}

func (storage *UserStorage) PostRegister(user domain.User) error{
	id, err := service.GenerateUUID()
	if err != nil{
		return err
	}

	storage.user[id] = user
	return nil
}

func (storage *UserStorage) PostLogin(user domain.User) (string, error){
	session, err := service.GenerateSession(user.Id)
	if err != nil{
		return "", err
	}

	storage.session[session.Id] = domain.Session{UserId : user.Id, Id : session.Id}
	return session.Id, nil
}



