
# JSON x Flatbuffers x Protocol Buffers

## To test JSON

```
curl -v -X "POST" "http://localhost:3000/json" \
     -H 'Accept: application/json' \
     -H 'Content-Type: application/json' \
     -d $'{
  "type": "button.clicked",
  "source": "Login",
  "subject": "user1000",
  "time": "2018-04-05T17:31:00Z",
  "data": "User clicked because X"
}'
```

## To test Flatbuffers

```
brew install flatbuffers
flatc --go event.fbs
go run cmd/fb/main.go
```

## To test Protobuff

```
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
brew install protobuf
protoc -I=. --go_out=./ event.proto

```
