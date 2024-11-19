run:
	go run main.go

build:
	go build ./...

build-app:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o soundlib main.go