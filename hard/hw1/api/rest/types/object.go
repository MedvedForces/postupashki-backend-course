package types

import (
	"net/http"
	taskError "server/repository/ram"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"io"
	errorsUser "server/repository/ram"
)

//user

type ObjectSessionHandlerResponse struct{
	Value string `json:"token"`
}

//task

type ObjectTaskUUIDHandlerResponse struct{
	Value string `json:"task_id"`
}

type ObjectTypeStatusHandlerResponse struct{
	Value string `json:"status"`
}

type ObjectTypeResultHandlerResponse struct{
	Value string `json:"result"`
}


type ObjectProcessError struct{
	StatusCode int
}

func (obj *ObjectProcessError) ProcessError(w http.ResponseWriter, err error, resp any){
	if err == taskError.TaskNotFound{
		http.Error(w, "Task not found with task_id", http.StatusNotFound)
		return
	} else if err == taskError.GenerateUUIDError{
		http.Error(w, "Error Generate uuid", http.StatusInternalServerError) // какая то супер базовая обработка ошибок, можно доработать потом
		return
	} else if err == errorsUser.ServerNotAuthorization{
		http.Error(w, "request don't have Authorization", 401) // какая то супер базовая обработка ошибок, можно доработать потом
		return
	} else if err != nil{
		http.Error(w, "Internal Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(obj.StatusCode)

	if err := json.NewEncoder(w).Encode(resp); err != nil{
		http.Error(w, "Internal Error", http.StatusInternalServerError)
	}
}

type objectTaskHandlerRequest struct{ //эта структура используется для парсинка Task - и методов Post, GetStatus, GetResult
	Content string
}

type objectUserHandlerRequest struct{ //эта структура используется для парсинка User
	UserName string `json:"username"`
	Password string `json:"password"`
}

func CreateObjectUserHandlerRequest(r *http.Request) (*objectUserHandlerRequest, error){
	var reqObject objectUserHandlerRequest
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(body, &reqObject); err != nil{
		return nil, err
	}

	return &reqObject, nil
}

func CreateObjectTaskIdHandlerRequest(r *http.Request) (*objectTaskHandlerRequest, error){
	content := chi.URLParam(r, "task_id")
	return &objectTaskHandlerRequest{Content: content}, nil
}

func CreateObjectTaskContentHandlerRequest(r *http.Request) (*objectTaskHandlerRequest, error){
	content := "content" //пока зашлушка
	return &objectTaskHandlerRequest{Content: content}, nil
}