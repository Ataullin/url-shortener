package auth

import (
	"fmt"
	"net/http"
	"url-shortner/configs"
	"url-shortner/pkg/request"
	"url-shortner/pkg/response"
)

type Handler struct {
	*configs.Config
}

type HandlerDeps struct {
	*configs.Config
}

func NewAuthHandler(router *http.ServeMux, deps HandlerDeps) {
	handler := &Handler{
		Config: deps.Config,
	}
	router.HandleFunc("POST /auth/login", handler.Login())
	router.HandleFunc("POST /auth/register", handler.Register())
}

func (handler *Handler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		body, err := request.HandleBody[LoginRequest](&w, req)
		if err != nil {
			return
		}
		fmt.Println(body)
		data := LoginResponse{
			Token: "123",
		}
		response.Json(w, data, 200)
	}
}

func (handler *Handler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := request.HandleBody[RegisterRequest](&w, r)
		if err != nil {
			return
		}
		fmt.Println(body)
	}
}
