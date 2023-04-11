
.PHONY: build clean test help default

BIN_NAME=go-mesa-state

GIT_COMMIT=$(shell git rev-parse HEAD)
GIT_DIRTY=$(shell test -n "`git status --porcelain`" && echo "+CHANGES" || true)
BUILD_DATE=$(shell date '+%Y-%m-%d-%H:%M:%S')
IMAGE_NAME := "asimazbunzel/go-mesa-state"

default: test

help:
	@echo 'Management commands for go-mesa-state:'
	@echo
	@echo 'Usage:'
	@echo '    make build           Compile the project.'
	@echo '    make clean           Clean the directory tree.'
	@echo

build:
	@echo "building ${BIN_NAME}"
	@echo "GOPATH=${GOPATH}"
	go build -ldflags "-X github.com/asimazbunzel/go-mesa-state/version.GitCommit=${GIT_COMMIT}${GIT_DIRTY} -X github.com/asimazbunzel/go-mesa-state/version.BuildDate=${BUILD_DATE}" -o bin/${BIN_NAME}

tidy:
	go mod tidy

clean:
	@test ! -e bin/${BIN_NAME} || rm bin/${BIN_NAME}

test:
	go test ./...

