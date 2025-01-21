build-main:
	go build -o bin/main main.go

build-http-server:
	go build -o bin/http-server ./worker-http/main.go

build:
	make build-main
	make build-http-server
