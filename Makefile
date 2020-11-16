Version := $(shell date "+%Y%m%d%H%M")
GitCommit := $(shell git rev-parse HEAD)
DIR := $(shell pwd)
LDFLAGS := -s -w -X main.Version=$(Version) -X main.GitCommit=$(GitCommit)

build:
	go build -race -ldflags "$(LDFLAGS)" -o build/debug/differ main.go

release:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "$(LDFLAGS)" -o build/release/differ main.go

.PHONY: run build release
