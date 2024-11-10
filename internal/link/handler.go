package link

import (
	"fmt"
	"net/http"
	"url-shortner/pkg/request"
	"url-shortner/pkg/response"
)

type LinkHandler struct {
	LinkRepository *LinkRepository
}

type LinkHandlerDeps struct {
	LinkRepository *LinkRepository
}

func NewLinkHandler(router *http.ServeMux, deps LinkHandlerDeps) {
	handler := &LinkHandler{
		LinkRepository: deps.LinkRepository,
	}
	router.HandleFunc("POST /link", handler.Create())
	router.HandleFunc("GET /link/{hash}", handler.GoTo())
	router.HandleFunc("PATH /link/{id}", handler.Update())
	router.HandleFunc("DELETE /link/{id}", handler.Delete())
}

func (handlerLink LinkHandler) Create() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := request.HandleBody[LinkCreateRequest](&w, r)
		if err != nil {
			return
		}
		link := NewLink(body.Url)
		createdLink, err := handlerLink.LinkRepository.Create(link)
		if err != nil {
			response.Json(w, err.Error(), 500)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		response.Json(w, createdLink, http.StatusCreated)
	}
}

func (handlerLink LinkHandler) GoTo() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		hash := r.PathValue("hash")
		byHash, err := handlerLink.LinkRepository.GetByHash(hash)
		if err != nil {
			response.Json(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Redirect(w, r, byHash.Url, http.StatusTemporaryRedirect)
	}
}

func (handlerLink LinkHandler) Update() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

func (handlerLink LinkHandler) Delete() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		fmt.Println("Delete link with id", id)
	}
}
