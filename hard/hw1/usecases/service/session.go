package service

import (
	"server/domain"
)

func GenerateSession(userId string) (*domain.Session, error){
	sessionId, err := GenerateUUID()
	if err != nil{
		return nil, err
	}

	return &domain.Session{UserId: userId, Id: sessionId}, nil
}