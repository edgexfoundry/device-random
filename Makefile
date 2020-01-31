.PHONY: build test clean prepare update docker

GO = CGO_ENABLED=0 GO111MODULE=on go

MICROSERVICES=cmd/device-random

.PHONY: $(MICROSERVICES)

DOCKERS=docker_device_random_go
.PHONY: $(DOCKERS)

VERSION=$(shell cat ./VERSION 2>/dev/null || echo 0.0.0)
GIT_SHA=$(shell git rev-parse HEAD)
GOFLAGS=-ldflags "-X github.com/edgexfoundry/device-random.Version=$(VERSION)"

build: $(MICROSERVICES)
	$(GO) build ./...

cmd/device-random:
	$(GO) build $(GOFLAGS) -o $@ ./cmd

test:
	$(GO) test ./... -coverprofile=coverage.out
	$(GO) vet ./...
	gofmt -l .
	[ "`gofmt -l .`" = "" ]
	./bin/test-attribution-txt.sh
	./bin/test-go-mod-tidy.sh

clean:
	rm -f $(MICROSERVICES)

docker: $(DOCKERS)

docker_device_random_go:
	docker build \
		--label "git_sha=$(GIT_SHA)" \
		-t edgexfoundry/docker-device-random-go:$(GIT_SHA) \
		-t edgexfoundry/docker-device-random-go:$(VERSION)-dev \
		.
