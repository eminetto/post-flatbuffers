package main

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/eminetto/post-flatbuffers/events"
	"github.com/eminetto/post-flatbuffers/events_pb"
	flatbuffers "github.com/google/flatbuffers/go"
	"google.golang.org/protobuf/proto"
)

func benchSetup() {
	os.Setenv("DEBUG", "false")
}

func BenchmarkJSON(b *testing.B) {
	benchSetup()
	r := handlers()
	payload := fmt.Sprintf(`{
		"type": "button.clicked",
		"source": "Login",
		"subject": "user1000",
		"time": "2018-04-05T17:31:00Z",
		"data": "User clicked because X"}`)
	for i := 0; i < b.N; i++ {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/json", strings.NewReader(payload))
		r.ServeHTTP(w, req)
		if w.Code != http.StatusCreated {
			b.Errorf("expected status 201, got %d", w.Code)
		}
	}
}

func BenchmarkFlatBuffers(b *testing.B) {
	benchSetup()
	r := handlers()
	builder := flatbuffers.NewBuilder(1024)
	evtType := builder.CreateString("button.clicked")
	evtSource := builder.CreateString("service-b")
	evtSubject := builder.CreateString("user1000")
	evtTime := builder.CreateString("2018-04-05T17:31:00Z")
	evtData := builder.CreateString("User clicked because X")

	events.EventStart(builder)
	events.EventAddType(builder, evtType)
	events.EventAddSource(builder, evtSource)
	events.EventAddSubject(builder, evtSubject)
	events.EventAddTime(builder, evtTime) //@todo use time.Time
	events.EventAddData(builder, evtData)
	evt := events.EventEnd(builder)
	builder.Finish(evt)

	buff := builder.FinishedBytes()
	for i := 0; i < b.N; i++ {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/fb", bytes.NewReader(buff))
		r.ServeHTTP(w, req)
		if w.Code != http.StatusCreated {
			b.Errorf("expected status 201, got %d", w.Code)
		}
	}
}

func BenchmarkProtobuffer(b *testing.B) {
	benchSetup()
	r := handlers()
	evt := events_pb.Event{
		Type:    "button.clicked",
		Subject: "user1000",
		Source:  "service-b",
		Time:    "2018-04-05T17:31:00Z",
		Data:    "User clicked because X",
	}
	payload, err := proto.Marshal(&evt)
	if err != nil {
		panic(err)
	}
	for i := 0; i < b.N; i++ {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/pb", bytes.NewReader(payload))
		r.ServeHTTP(w, req)
		if w.Code != http.StatusCreated {
			b.Errorf("expected status 201, got %d", w.Code)
		}
	}
}
