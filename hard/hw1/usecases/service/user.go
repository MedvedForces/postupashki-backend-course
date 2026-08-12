package service

import (
	"server/domain"
)

type User interface{
	PostRegister(user domain.User) error
	PostLogin(user domain.User) (string, error)
} 