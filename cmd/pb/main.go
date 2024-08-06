package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	"github.com/eminetto/post-flatbuffers/events_pb"
	"google.golang.org/protobuf/proto"
)

func main() {
	evt := events_pb.Event{
		Type:    "button.clicked",
		Subject: "user1000",
		Source:  "service-b",
		Time:    "2018-04-05T17:31:00Z",
		Data:    "User clicked because X",
	}
	b, err := proto.Marshal(&evt)
	if err != nil {
		panic(err)
	}
	u := "http://localhost:3000/pb"
	req, err := http.NewRequest("POST", u, bytes.NewReader(b))
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
