MAKEFLAGS += --warn-undefined-variables
SHELL := /bin/bash
.DEFAULT_GOAL := build

.PHONY: clean lint

ROOT := $(shell pwd)
PACKAGE := HAWAI/repos/hawai-crm

clean:
	rm -rf build cover
	rm hawai-crm

build:
	go build -v

rebuild: clean build

test:
	go test -v -race ./...

lint:
	go vet ./...
	golint ./...

package:
	sudo docker build -t hawai-crm $(ROOT)

all: clean build lint test package

run:
	sudo docker run --publish 32001:32001 --name crm --rm hawai-crm
