all: build-binary

UNIT_TEST_PACKAGES=$(shell  go list ./... | grep -v "vendor")
APP_EXECUTABLE=grpc-replay

fmt:
	go fmt ./...

test:
	GOCACHE=off go test $(UNIT_TEST_PACKAGES)

deps:
	dep ensure

install:
	mkdir -p $(GOPATH)/bin/
	go install $(APP_EXECUTABLE)

build-binary: deps install
