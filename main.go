package main

import (
	"chi-example/entities"
	"chi-example/httpd"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"net/http"
)

func main() {
	port := ":3000"
	feed := entities.New()
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
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
