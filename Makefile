.ONESHELL:
.SHELL := /usr/bin/bash
.PHONY: build clean deploy undeploy run deps

CURRENT_FOLDER := $(shell pwd)

EXECUTABLES = docker docker-machine serverless dep sam npm pip python zip
K := $(foreach exec,$(EXECUTABLES),\
		$(if $(shell command -v $(exec) 2> /dev/null),,$(error "ERROR: $(exec) is missing)))

default: build

build:
	@dep ensure -v
	@env GOOS=linux GOARCH=amd64 go build -o bin/app app/*
ifeq ($(OS),Windows_NT)
	@build-lambda-zip -o lambda.zip bin/app
else
	@chmod +x bin/app
	@zip -jq lambda.zip bin/app
endif

clean:
	@rm -rf ./bin ./vendor Gopkg.lock

deploy:
	@serverless deploy --verbose

undeploy:
	@serverless remove --verbose

run:
ifeq ($(DOCKER_HOST),)
	$(error DOCKER_HOST is not set)
endif
	@MSYS2_ARG_CONV_EXCL="*" sam local invoke app --event test/event.json --docker-volume-basedir "$(CURRENT_FOLDER)"

deps:
	@npm install -g serverless
	@pip install aws-sam-cli
