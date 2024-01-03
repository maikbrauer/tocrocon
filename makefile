default: help

build:
	env GOOS=linux go build -ldflags="-s -w" -o main .
	mkdir -p bin/
	zip bin/main.zip main
	rm bin/main

build-only-binary:
	env GOOS=linux go build -ldflags="-s -w" -o main .
	mkdir -p bin/
	mv main bin/

test:
	go test -v ./...

clean:
	rm  bin/*

help:
	@echo 'Usage: make (build | test | clean)'