package main

import (
	"fmt"
	"server/repository/ram"
	"server/usecases/service/task"
	"server/usecases/service/user"
	"github.com/go-chi/chi/v5"
    restTask "server/api/rest/task"
	restUser "server/api/rest/user"
	pkgHttp "server/pkg"
	_ "server/docs"
	httpSwager "github.com/swaggo/http-swagger"
)


// @title           Rest API для обработки users и tasks
// @version         1.0
// @description     Handlers для создания/поиска user/task

// @host      localhost:8000
// @BasePath  /

func createTaskApi() *restTask.Task{
	taskStorage := ram.NewTaskStorage()
	taskService := task.NewService(taskStorage)
	taskHandler := restTask.NewTask(taskService)
	return taskHandler
}

func createUserApi() *restUser.User{
	userStorage := ram.NewUserStorage()
	userService := user.NewService(userStorage)
	userHandler := restUser.NewUser(userService)
	return userHandler
}

func main(){
	fmt.Println("Hello I'm main")
	address := "127.0.0.1:8000"

	router := chi.NewRouter()
	router.Get("/swagger/*", httpSwager.WrapHandler)
	createTaskApi().SettingRouting(router)
	createUserApi().SettingRouting(router)

	if err := pkgHttp.RunServer(router, address); err != nil{
		fmt.Println("Ошибка создания сервера")
	}
}