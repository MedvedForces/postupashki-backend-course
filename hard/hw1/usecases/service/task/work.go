package task

import (
	"time"
)

func (service *Service) makeTask(content string) string{
	time.Sleep(5 * time.Second) //бурная деятельность

	result := "Результат таски"
	// сохраняю результат таски(пока что это какой то мусор)
	return result
}