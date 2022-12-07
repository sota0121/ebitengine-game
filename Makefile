.PHONY: build run

GO ?= go
APP_NAME ?= app
APP_VERSION ?= 0.0.1
MAIN_FILE ?= cmd/main.go
BIN_DIR ?= bin

build:
	$(GO) build -o $(BIN_DIR)/$(APP_NAME) -ldflags "-X main.version=$(APP_VERSION)" $(MAIN_FILE)

run:
	$(GO) run $(MAIN_FILE)
