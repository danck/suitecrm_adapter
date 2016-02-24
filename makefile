MAKEFLAGS += --warn-undefined-variables
SHELL := /bin/bash
.DEFAULT_GOAL := build

ROOT := $(shell pwd)
#PACKAGE := gitlab.com/danck/hawai-suitecrm

clean:
	rm -rf build cover

build:
	mkdir build
	cd $(ROOT)/build
	go build $(ROOT)
