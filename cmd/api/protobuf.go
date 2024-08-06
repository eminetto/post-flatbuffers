package main

import (
	"io"
	"net/http"

	"github.com/eminetto/post-flatbuffers/events_pb"
	"google.golang.org/protobuf/proto"
)

func processPB() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body := r.Body
		data, _ := io.ReadAll(body)

		e := events_pb.Event{}
		err := proto.Unmarshal(data, &e)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		saveEvent(e.GetType(), e.GetSource(), e.GetSubject(), e.GetTime(), e.GetData())
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("protobuf received"))
	}
}
