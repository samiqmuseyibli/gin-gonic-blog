ROOT_DIR := $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))

BIN_DIR = $(ROOT_DIR)/bin
BIN = $(BIN_DIR)/gin-gonic-blog 
APP_DIR = /go/src/github.com/samigmuseyibli/gin-gonic-blog

build:
	@(echo "-> Compiling")
	@(mkdir -p $(BIN_DIR))
	@(go build -mod=vendor -o $(BIN_DIR)/gin-gonic-blog ./cmd/gin-gonic-blog/main.go)
	@(echo "-> gin-gonic-blog binary created")

run:
	go run main.go