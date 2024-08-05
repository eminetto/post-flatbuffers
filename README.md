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

Sobre benchmarks: https://www.practical-go-lessons.com/chap-34-benchmarks
@todo: criar gráficos com as informações abaixo
Running tool: /opt/homebrew/bin/go test -benchmem -run=^$ -coverprofile=/var/folders/vn/gff4w90d37xbfc_2tn3616h40000gn/T/vscode-goPkxuz8/go-code-cover -bench . github.com/eminetto/post-flatbuffers/events-api -failfast -v

goos: darwin
goarch: arm64
pkg: github.com/eminetto/post-flatbuffers/events-api
BenchmarkJSON
BenchmarkJSON-8          	  737263	      1529 ns/op	    2288 B/op	      26 allocs/op
BenchmarkFlatBuffers
BenchmarkFlatBuffers-8   	 2183271	       551.6 ns/op	    1768 B/op	      16 allocs/op
PASS
coverage: 82.6% of statements
ok  	github.com/eminetto/post-flatbuffers/events-api	4.230s

@todo pegar graficos de comparação do próprio site do flatbuffers
@todo unificar em um diretório só para ficar mais fácil publicar no github
@todo criar diagrama mostrando a ideia da PoC, como feito no ipad
