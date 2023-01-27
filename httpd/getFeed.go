package httpd

import (
	"chi-example/entities"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

func GetFeed(feed *entities.Repo) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		queryText := chi.URLParam(r, "query")

		myStoredVariable := GetData(queryText)
		feed.AddAll(myStoredVariable.Hits)
		items := feed.GetAll()

		errResponse := json.NewEncoder(w).Encode(items)
		if errResponse != nil {
			log.Fatal(errResponse)
		}
	}
}
