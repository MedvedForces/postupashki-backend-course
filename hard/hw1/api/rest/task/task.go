package rest

import (
	"net/http"
	"server/api/rest/types"
	"server/usecases/service"
	"github.com/go-chi/chi/v5"
	"strings"
	errorsUser "server/repository/ram"
)

type Task struct{
	service service.Task
}

func NewTask(service service.Task) *Task{
	return &Task{
		service,
	}
}

func (task *Task) SettingRouting(r chi.Router){
	r.Route("/task", func(r chi.Router){
		r.Post("/", task.PostHandler)
	})

	r.Route("/status", func(r chi.Router){
		r.Get("/{task_id}", task.GetStatus)
	})

	r.Route("/result", func(r chi.Router){
		r.Get("/{task_id}", task.GetResult)
	})
}

func (task *Task) isHeaderAuthorization(r *http.Request) (bool, string){
	for name, value := range r.Header{
		if name == "Authorization"{
			return true, strings.Join(value, "")
		}
	}

	return false, ""
}

// @Summary Отправить таску на обработку
// @Description Отправить таску на обработку
// @Tags Таска
// @Produce json
// @Param task_body body string true "Таска которую нужно решить в string формате"
// @Param Authorization header string false "Bearer {auth_token}"
// @Success 200 {object} types.ObjectTypeResultHandlerResponse
// @Failure 400 {string} string "Bad request"
// @Failure 404 {string} string "Task not found"
// @Router /task [post]
func (task *Task) PostHandler(w http.ResponseWriter, r *http.Request) {
	var obj types.ObjectProcessError
	ok, _ := task.isHeaderAuthorization(r) //пока session никак не используется поэтому заглушка

	if !ok{
		statusCode := 401
		obj = types.ObjectProcessError{StatusCode : statusCode}
		obj.ProcessError(w, errorsUser.ServerNotAuthorization, nil)
		return
	}

	reqObjectParsing, _ := types.CreateObjectTaskContentHandlerRequest(r) //ошибка пока не обрабатывается

	uuid, err := task.service.Post(reqObjectParsing.Content)

	respObject := types.ObjectTaskUUIDHandlerResponse{Value : uuid}

	statusCode := 201

	obj = types.ObjectProcessError{StatusCode : statusCode}

	obj.ProcessError(w, err, respObject)
}

// @Summary Готовность таски
// @Description Получить статус таски по task_id
// @Tags Таска
// @Produce json
// @Param task_id path string true "task id который получен от сервера"
// @Param Authorization header string false "Bearer {auth_token}"
// @Success 200 {object} types.ObjectTypeResultHandlerResponse
// @Failure 400 {string} string "Bad request"
// @Failure 404 {string} string "Task not found"
// @Router /status/{task_id} [get]
func (task *Task) GetStatus(w http.ResponseWriter, r *http.Request){
	var obj types.ObjectProcessError
	ok, _ := task.isHeaderAuthorization(r) //пока session никак не используется поэтому заглушка

	if !ok{
		statusCode := 401
		obj = types.ObjectProcessError{StatusCode : statusCode}
		obj.ProcessError(w, errorsUser.ServerNotAuthorization, nil)
		return
	}

	reqObjectParsing, _ := types.CreateObjectTaskIdHandlerRequest(r) //ошибка пока не обрабатывается

	status, err := task.service.GetStatus(reqObjectParsing.Content)

	respObject := types.ObjectTypeStatusHandlerResponse{Value : status}

	statusCode := 200

	obj = types.ObjectProcessError{StatusCode : statusCode}

	obj.ProcessError(w, err, respObject)
}

// @Summary Получить результат таски
// @Description Получить результат таски по task_id
// @Tags Таска
// @Produce json
// @Param task_id path string true "task id который получен от сервера"
// @Param Authorization header string false "Bearer {auth_token}"
// @Success 200 {object} types.ObjectTypeResultHandlerResponse
// @Failure 400 {string} string "Bad request"
// @Failure 404 {string} string "Task not found"
// @Router /result/{task_id} [get]
func (task *Task) GetResult(w http.ResponseWriter, r *http.Request){
	var obj types.ObjectProcessError
	ok, _ := task.isHeaderAuthorization(r) //пока session никак не используется поэтому заглушка

	if !ok{
		statusCode := 401
		obj = types.ObjectProcessError{StatusCode : statusCode}
		obj.ProcessError(w, errorsUser.ServerNotAuthorization, nil)
		return
	}

	reqObjectParsing, _ := types.CreateObjectTaskIdHandlerRequest(r) //ошибка пока не обрабатывается

	result, err := task.service.GetResult(reqObjectParsing.Content)

	respObject := types.ObjectTypeResultHandlerResponse{Value : result}

	statusCode := 200

	obj = types.ObjectProcessError{StatusCode : statusCode}

	obj.ProcessError(w, err, respObject)
}
