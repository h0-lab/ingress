.PHONY : default deps test build image docs

export VERSION = 0.9.0

HARDWARE = $(shell uname -m)
OS := $(shell uname)

default: deps test build

deps:
	@echo "Configuring Last.Backend Dependencies"
	go get -u github.com/golang/dep/cmd/dep
	dep ensure

test:
	@echo "Testing Last.Backend ingress image"
	@sh ./hack/run-coverage.sh

docs: docs/*
	@echo "Build Last.Backend ingress Documentation"
	@sh ./hack/build-docs.sh

build:
	@echo "== Pre-building configuration"
	mkdir -p build/linux && mkdir -p build/darwin
	@echo "== Building Last.Backend ingress server"
	@bash ./hack/build-cross.sh

install:
	@echo "== Install binaries"
	@bash ./hack/install-cross.sh

image:
	@echo "== Pre-building configuration"
	@sh ./hack/build-images.sh

run:
	@echo "== Run ingress local from sources"
	@go run ./cmd/ingress/ingress.go --config=./contrib/config.yml