.PHONY: build test clean prepare update docker

GO=CGO_ENABLED=0 go

MICROSERVICES=cmd/device-random

.PHONY: $(MICROSERVICES)

DOCKERS=docker_device_random_go
.PHONY: $(DOCKERS)

VERSION=$(shell cat ./VERSION)
GIT_SHA=$(shell git rev-parse HEAD)
GOFLAGS=-ldflags "-X github.com/edgexfoundry/device-random.Version=$(VERSION)"

build: $(MICROSERVICES)
	go build ./...

cmd/device-random:
	$(GO) build $(GOFLAGS) -o $@ ./cmd

test:
	go test ./... -cover

clean:
	rm -f $(MICROSERVICES)

prepare:
	glide install

update:
	glide update

docker: $(DOCKERS)

docker_device_random_go:
	docker build \
		--label "git_sha=$(GIT_SHA) \
		-t edgexfoundry/docker-device-random-go:$(GIT_SHA) \
		-t edgexfoundry/docker-device-random-go:$(VERSION)-dev \
		.
