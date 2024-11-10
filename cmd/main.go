package main

import (
	"fmt"
	"net/http"
	"url-shortner/configs"
	"url-shortner/internal/auth"
	"url-shortner/internal/link"
	"url-shortner/pkg/db"
)

func main() {
	config := configs.LoadConfig()
	db := db.NewDb(config)

	linkRepository := link.NewLinkRepository(db)

	router := http.NewServeMux()
	auth.NewAuthHandler(router, auth.HandlerDeps{
		Config: config,
	})
	link.NewLinkHandler(router, link.LinkHandlerDeps{
		LinkRepository: linkRepository,
	})

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	fmt.Println("server listening on port 8081")
	server.ListenAndServe()
}
