package main

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

type event struct {
	ID      uuid.UUID
	Type    string `json:"type"`
	Source  string `json:"source"`
	Subject string `json:"subject"`
	Time    string `json:"time"`
	Data    string `json:"data"`
}

func processJSON() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var e event
		err := json.NewDecoder(r.Body).Decode(&e)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		saveEvent(e.Type, e.Source, e.Subject, e.Time, e.Data)
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("json received"))
	}
}
