.PHONY: test test-cover-set test-cover-atomic test-cover-count clean

test:
	go test -v ./...

test-cover-set:
	go test -covermode=set -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

test-cover-atomic:
	go test -covermode=atomic -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

test-cover-count:
	go test -covermode=count -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

clean:
	go clean -testcache
