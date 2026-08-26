package ram

import (
	"errors"
)

var (
	TaskNotFound = errors.New("Task with uuid not found")
	GenerateUUIDError = errors.New("Error generate uuid")
	ServerNotAuthorization = errors.New("401 Unauthorized")
)