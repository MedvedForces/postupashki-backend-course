package rest

import (
	"net/http"
	"server/api/rest/types"
	"server/usecases/service"
	"github.com/go-chi/chi/v5"
	"server/domain"
)

type User struct{
	service service.User
}

func NewUser(service service.User) *User{
	return &User{
		service,
	}
}

func (user *User) SettingRouting(r chi.Router){

	r.Route("/", func(r chi.Router){
		r.Post("/register", user.PostRegisterHandler)
		r.Post("/login", user.PostLoginHandler)
	})
}

// @Summary Авторизация
// @Description Залогиниться на сайте и получить id сессии
// @Tags user
// @Produce json
// @Param User_body body string true "username и password"
// @Success 200 {object} types.ObjectTypeResultHandlerResponse
// @Failure 400 {string} string "Bad request"
// @Failure 404 {string} string "User not found"
// @Router /login [post]
func (User *User) PostLoginHandler(w http.ResponseWriter, r *http.Request) {
	reqObjectParsing, err := types.CreateObjectUserHandlerRequest(r)

	statusCode := 200

	obj := types.ObjectProcessError{StatusCode : statusCode}

	if err != nil{
		obj.ProcessError(w, err, nil)
		return
	}

	sessionId, err := User.service.PostLogin(domain.User{
		Name: reqObjectParsing.UserName, 
		Password: reqObjectParsing.Password,
		Id: "",
	})

	respObject := types.ObjectSessionHandlerResponse{Value : sessionId} // создать types.ObjectTaskUUIDHandlerResponse и заменить в будущем

	obj.ProcessError(w, err, respObject)
}

// @Summary Регистрация
// @Description Зарегаться на сайте
// @Tags user
// @Produce json
// @Param User_body body string true "username и password"
// @Success 200 {object} types.ObjectTypeResultHandlerResponse
// @Failure 400 {string} string "Bad request"
// @Failure 404 {string} string "User not found"
// @Router /register [post]
func (User *User) PostRegisterHandler(w http.ResponseWriter, r *http.Request) {
	reqObjectParsing, err := types.CreateObjectUserHandlerRequest(r)
	
	statusCode := 201
	obj := types.ObjectProcessError{StatusCode : statusCode}

	if err != nil{
		obj.ProcessError(w, err, nil)
		return
	}

	err = User.service.PostRegister(domain.User{
		Name: reqObjectParsing.UserName, 
		Password: reqObjectParsing.Password,
		Id: "",
	})

	obj.ProcessError(w, err, nil)
}