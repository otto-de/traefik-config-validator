.DEFAULT_GOAL := all

VERSION ?= SNAPSHOT

EXE ?= traefik-config-validator

IMAGE_NAMESPACE ?= otto-de
IMAGE_NAME ?= $(EXE)
IMAGE_FULLNAME ?= $(IMAGE_NAMESPACE)/$(IMAGE_NAME)
IMAGE_REGISTRY ?= ghcr.io
IMAGE_BUILD_CONTEXT ?= .
IMAGE_BUILD_MANIFEST ?= Dockerfile

all: clean verify test lint build

clean:
	$(RM) -f $(EXE)
	$(RM) -rf dist

verify:
	go mod verify
	go mod tidy

build:
	go build -o $(EXE) cmd/traefik-config-validator/main.go

test:
	go test -v -race -count=1 ./...

lint:
	golangci-lint run -v -c .golangci.yml

fmt:
	golangci-lint run -v --fix -c .golangci.yml

docker-image:
	docker build -t $(IMAGE_FULLNAME):$(VERSION) -f $(IMAGE_BUILD_MANIFEST) --build-arg VERSION=$(VERSION) $(IMAGE_BUILD_CONTEXT)
	docker tag $(IMAGE_FULLNAME):$(VERSION) $(IMAGE_REGISTRY)/$(IMAGE_FULLNAME):$(VERSION)

release:
	goreleaser release --rm-dist

release/snapshot:
	goreleaser release --rm-dist --snapshot
