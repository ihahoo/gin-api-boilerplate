.PHONY: start build run lint clean doc dev

GONAME=$(shell basename "$(PWD)")

default: build

start:
	@GIN_MODE=release ./bin/$(GONAME) 

build:
	@go build -o bin/$(GONAME) 

run:
	@./bin/$(GONAME) 

lint:
	@golint

clean:
	@go clean && rm -rf ./bin/$(GONAME) && rm -f gin-bin

doc:
	godoc -http=:6060 -index

dev:
	@gin -a 8080 -p 3030 run main.go
