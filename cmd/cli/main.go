package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	"github.com/eminetto/post-flatbuffers/events"
	flatbuffers "github.com/google/flatbuffers/go"
)

func main() {
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

	u := "http://localhost:3000/fb"
	req, err := http.NewRequest("POST", u, bytes.NewReader(buff))
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/x-binary")

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	fmt.Println("response Status:", response.Status)
	body, _ := io.ReadAll(response.Body)
	fmt.Println("response Body:", string(body))
}
