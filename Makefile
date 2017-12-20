all: deps build run

.PHONY: deps
deps:
	go get -d -v github.com/gin-gonic/gin
	go get -d -v github.com/gin-contrib/static
	go get -d -v github.com/dustin/go-broadcast

.PHONY: build
build:
	go build -o app src/*.go

.PHONY: run
run:
	./app
