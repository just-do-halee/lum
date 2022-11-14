all: examples

test:
	go test

examples: test
	go test ./examples/**
