package repository

import (
	"server/domain"
)

type UserStorageI interface{
	PostRegister(user domain.User) error
	PostLogin(user domain.User) (string, error)
}