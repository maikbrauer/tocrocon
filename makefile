default: help

build:
	GOARCH=amd64 GOOS=linux	go build .
	zip main.zip main
	mkdir -p bin
	mv main.zip bin/
	rm main

test:
	go test -v ./...

clean:
	rm  bin/*

help:
	@echo 'Usage: make (build | test | clean)'