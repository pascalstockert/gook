gook:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/gook cmd/gook/main.go

gook-http:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/http-server cmd/gook-http/main.go

all:
	make gook
	make gook-http
