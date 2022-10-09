.EXPORT_ALL_VARIABLES:
GOOS ?= $(uname -s)
GOARCH ?= $(uname -m)
LD_FLAGS := -ldflags="-s -w -X 'main.BuildDate=$(shell date)'"

build:
	mkdir -p bin/
	go build -v -o bin/gosh *.go

build_release:
	env CGO_ENABLED=0 GOOS=$(GOOS) GOARCH=$(GOARCH) go build $(LD_FLAGS) -o bin/gosh_release *.go

run: build
	chmod +x bin/gosh
	./bin/gosh

test:
	go test -v ./...