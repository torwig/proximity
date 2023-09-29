BINARY_NAME = proxx

## help: print the help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

## tidy: runs additional commands/tools to tidy the codebase
.PHONY: tidy
tidy:
	go mod tidy
	go mod verify
	goimports -w .

## test: runs tests
.PHONY: test
test:
	go test -v -race ./...

## lint: runs golangci-lint
.PHONY: lint
lint:
	golangci-lint run ./...

## build: builds the application/service
.PHONY: build
build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $(BINARY_NAME) cmd/$(BINARY_NAME)/main.go

## build: builds the application/service for Windows, x64
.PHONY: build_win
build_win:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o $(BINARY_NAME).exe cmd/$(BINARY_NAME)/main.go
