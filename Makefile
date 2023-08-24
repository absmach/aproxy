# Copyright (c) Mainflux
# SPDX-License-Identifier: Apache-2.0

PROGRAM = aproxy
AMDM_DOCKER_IMAGE_NAME_PREFIX ?= amdm
SOURCES = $(wildcard *.go) cmd/main.go
CGO_ENABLED ?= 0
GOARCH ?= amd64
VERSION ?= $(shell git describe --abbrev=0 --tags 2>/dev/null || echo "0.13.0")
COMMIT ?= $(shell git rev-parse HEAD)
TIME ?= $(shell date +%F_%T)

all: $(PROGRAM)

.PHONY: all clean $(PROGRAM)

define make_docker
	docker build \
		--no-cache \
		--build-arg SVC=$(PROGRAM) \
		--build-arg GOARCH=$(GOARCH) \
		--build-arg GOARM=$(GOARM) \
		--build-arg VERSION=$(VERSION) \
		--build-arg COMMIT=$(COMMIT) \
		--build-arg TIME=$(TIME) \
		--tag=$(AMDM_DOCKER_IMAGE_NAME_PREFIX)/$(PROGRAM) \
		-f docker/Dockerfile .
endef

define make_docker_dev
	docker build \
		--no-cache \
		--build-arg SVC=$(PROGRAM) \
		--tag=$(AMDM_DOCKER_IMAGE_NAME_PREFIX)/$(PROGRAM) \
		-f docker/Dockerfile.dev ./build
endef

$(PROGRAM): $(SOURCES)
	CGO_ENABLED=$(CGO_ENABLED) GOOS=$(GOOS) GOARCH=$(GOARCH) GOARM=$(GOARM) \
	go build -mod=vendor -ldflags "-s -w \
	-X 'github.com/mainflux/mainflux.BuildTime=$(TIME)' \
	-X 'github.com/mainflux/mainflux.Version=$(VERSION)' \
	-X 'github.com/mainflux/mainflux.Commit=$(COMMIT)'" \
	-o ./build/$(AMDM_DOCKER_IMAGE_NAME_PREFIX)-$(PROGRAM) cmd/main.go

clean:
	rm -rf $(PROGRAM)

docker-image:
	$(call make_docker)
docker-dev:
	$(call make_docker_dev)

run:
	docker compose -f ./docker/docker-compose.yml up

