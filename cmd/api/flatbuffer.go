package main

import (
	"io"
	"net/http"

	"github.com/eminetto/post-flatbuffers/events"
)

func processFB() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body := r.Body
		data, _ := io.ReadAll(body)
		e := events.GetRootAsEvent(data, 0)
		saveEvent(string(e.Type()), string(e.Source()), string(e.Subject()), string(e.Time()), string(e.Data()))
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("flatbuffer received"))
	}
}
