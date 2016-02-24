MAKEFLAGS += --warn-undefined-variables
SHELL := /bin/bash
.DEFAULT_GOAL := build

.PHONY: clean lint

ROOT := $(shell pwd)
PACKAGE := gitlab.com/danck/hawai-suitecrm

clean:
	rm -rf build cover
	rm hawai-suitecrm

build:
	go build -v

rebuild: clean build

test:
	go test -race ./...

lint:
	go vet ./...
	golint ./...

package:
	sudo docker build -t hawai-suitecrm $(ROOT)

all: clean build lint test package

run:
	sudo docker run --publish 8080:8080 --name test --rm hawai-suitecrm
