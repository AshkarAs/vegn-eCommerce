run:
	go run ./cmd

build:
	go build -o ./cmd/vegnExecutableFile ./cmd/main.go

buildrun:
	./cmd/vegnExecutableFile

swaggo:
	swag init -g ./cmd/main.go

swaggoformat:
	swag fmt	

test:
	go test -v ./...
		