package main

import (
	"chi-example/httpd"
	"chi-example/newsfeed"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
)

func main() {
	port := ":3000"
	feed := newsfeed.New()
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/feed/{query}", func(w http.ResponseWriter, r *http.Request) {
		queryText := chi.URLParam(r, "query")

		myStoredVariable := httpd.GetData(queryText)
		feed.AddAll(myStoredVariable.Hits)
		items := feed.GetAll()

		errResponse := json.NewEncoder(w).Encode(items)
		if errResponse != nil {
			log.Fatal(errResponse)
		}
	})
	fmt.Printf("Server on: %T\n", port)
	err := http.ListenAndServe(port, r)
	if err != nil {
		return
	}
}
