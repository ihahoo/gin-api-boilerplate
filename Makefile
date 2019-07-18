.PHONY: start build run lint clean doc dev docker-build docker-image docker-image-staging docker-image-dev start-docker-dev stop-docker-dev

GONAME=api

default: build

start:
	@GIN_MODE=release ./bin/$(GONAME) 

build:
	@go build -mod vendor -o bin/$(GONAME) 

run:
	@./bin/$(GONAME) 

lint:
	@golint

clean:
	@go clean && rm -rf ./bin/$(GONAME) && rm -f gin-bin

doc:
	godoc -http=:6060 -index

dev:
	@fresh

docker-build: clean
	@docker-compose -f docker/development/docker-compose.yml run --rm api make build

docker-image: docker-build 
	@docker build -t my-api:latest .

docker-image-staging: docker-build 
	@docker build -t my-api:staging .

docker-image-dev:
	@docker-compose -f docker/development/docker-compose.yml run --rm

start-docker-dev:
	@docker-compose up -d

stop-docker-dev:
	@docker-compose down
