package repository

import (
	"server/domain"
)

type TaskStorageI interface{
	PostResult(task domain.Task)
	PostProgress() (string, error)
	GetStatus(task domain.Task) (string, error)
	GetResult(task domain.Task) (string, error)
}