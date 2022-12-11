.PHONY: build clean test run

GO ?= go
MAIN_FILE ?= cmd/main.go
OUT_DIR ?= out
BINARY_NAME ?= ebitenx
BIN_DIR ?= $(OUT_DIR)/bin
BIN_PATH ?= $(BIN_DIR)/$(BINARY_NAME)
COVER_PROFILE ?= $(OUT_DIR)/coverage.out

build:
	$(GO) build -o $(BIN_PATH) -ldflags "-s -w" $(MAIN_FILE)

run:
	$(GO) run $(MAIN_FILE)

test:
	$(GO) test -coverprofile=$(COVER_PROFILE) ./...

clean:
	rm -rf $(OUT_DIR)/*
