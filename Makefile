#!/usr/bin/make

test:
	cd portainers; go test -v -cover

run:
	cd client; go run main.go

init:
	cd portainers; go mod download
	cd client; go mod download
