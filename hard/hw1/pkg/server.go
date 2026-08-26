package pkg

import (
	"server/usecases/service/task"
	"net/http"
	"github.com/go-chi/chi/v5"
)


func RunServer(r chi.Router, addr string) error{
	server := &http.Server{
		Addr:   addr,
		Handler: r,
	}

	err := server.ListenAndServe()
	return err
}

//

type Server struct{
	service task.Service
}

func NewServer(service task.Service) *Server{
	return &Server{
		service,
	}
}