package main

import (
	"chi-example/entities"
	"chi-example/httpd"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func main() {
	port := ":3000"
	feed := entities.New()
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Route("/feed", func(r chi.Router) {
		r.Get("/", httpd.GetFeed(feed))
		r.Get("/{query}", httpd.GetFeed(feed))
	})
	fmt.Printf("Server on: %s\n", port)
	err := http.ListenAndServe(port, r)
	if err != nil {
		return
	}
}
