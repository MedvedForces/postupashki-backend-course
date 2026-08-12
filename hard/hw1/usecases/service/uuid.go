package service

import (
	"github.com/nu7hatch/gouuid"
)

func GenerateUUID() (string, error){ //пока это заглушка, надо переписать
	uuid, err := uuid.NewV4()
	if err != nil{
		return "", err
	}

	return uuid.String(), nil
}