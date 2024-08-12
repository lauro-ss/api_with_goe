all: build run

build:
	@go build -o cmd/main cmd/main.go

run: 
	@cmd/main