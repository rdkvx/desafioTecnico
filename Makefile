PKGS ?= $(shell go list ./...)

run:
	go run main.go