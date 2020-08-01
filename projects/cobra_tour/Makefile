# Go parameters
GOBUILD=go build
GOCLEAN=go clean
GOTEST=gotestsum --format short-verbose --no-summary=output --junitfile unit-tests.xml --

all: build
build:
	$(GOBUILD)

race:
	$(GOBUILD) -race

test:
	$(GOTEST) -gcflags=-l -v ./...

clean:
	$(GOCLEAN)
