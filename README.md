
## Test JSON

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

## Test Flatbuffers

```
flatc --go event.fbs
go run cmd/fb/main.go
```

## Test Protobuff

```
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
brew install protobuf
protoc -I=. --go_out=./ event.proto

```
## Todo
Sobre benchmarks: https://www.practical-go-lessons.com/chap-34-benchmarks
- [x] criar gráficos com as informações abaixo
```
Running tool: /opt/homebrew/bin/go test -benchmem -run=^$ -coverprofile=/var/folders/vn/gff4w90d37xbfc_2tn3616h40000gn/T/vscode-gojAS4GO/go-code-cover -bench . github.com/eminetto/post-flatbuffers/cmd/api -failfast -v

goos: darwin
goarch: arm64
pkg: github.com/eminetto/post-flatbuffers/cmd/api
BenchmarkJSON
BenchmarkJSON-8          	  658386	      1732 ns/op	    2288 B/op	      26 allocs/op
BenchmarkFlatBuffers
BenchmarkFlatBuffers-8   	 1749194	       640.5 ns/op	    1856 B/op	      21 allocs/op
BenchmarkProtobuffer
BenchmarkProtobuffer-8   	 1497356	       696.9 ns/op	    1952 B/op	      21 allocs/op
PASS
coverage: 77.5% of statements
ok  	github.com/eminetto/post-flatbuffers/cmd/api	5.042s

```

- [x] adicionar protobuffers
- [x] unificar em um diretório só para ficar mais fácil publicar no github
- [ ] criar diagrama mostrando a ideia da PoC, como feito no ipad
