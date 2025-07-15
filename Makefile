.PHONY: build test clean

build:
	go build -o bin/checker ./cmd/checker

test:
	go test -v ./...

clean:
	rm -rf bin/
	go clean

run:
	./bin/checker -path $(path)

install:
	go install ./cmd/checker 