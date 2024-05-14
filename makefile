default: help

build:
	env GOOS=linux go build -ldflags="-s -w" -o bootstrap .
	mkdir -p bin/
	zip bin/bootstrap.zip bootstrap
	rm bootstrap

build-only-binary:
	env GOOS=linux go build -ldflags="-s -w" -o bootstrap .
	mkdir -p bin/
	mv bootstrap bin/

test:
	go test -v ./...

clean:
	rm  bin/*

help:
	@echo 'Usage: make (build | test | clean)'