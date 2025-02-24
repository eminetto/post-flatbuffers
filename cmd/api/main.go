package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/google/uuid"
)

func main() {
	r := handlers()
	http.ListenAndServe(":3000", r)
}

func handlers() *chi.Mux {
	r := chi.NewRouter()
	if os.Getenv("DEBUG") != "false" {
		r.Use(middleware.Logger)
	}
	r.Post("/json", processJSON())
	r.Post("/fb", processFB())
	r.Post("/pb", processPB())
	return r
}

func saveEvent(evType, source, subject, time, data string) {
	if os.Getenv("DEBUG") != "false" {
		id := uuid.New()
		q := fmt.Sprintf("insert into event values('%s', '%s', '%s', '%s', '%s', '%s')", id, evType, source, subject, time, data)
		fmt.Println(q)
	}
	// save event to database
}
