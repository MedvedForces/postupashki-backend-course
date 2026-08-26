package service

type Task interface{
	Post(content string) (string, error)
	GetStatus(task_id string) (string, error)
	GetResult(task_id string) (string, error)
} 