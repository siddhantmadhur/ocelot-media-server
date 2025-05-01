PKG := github.com/siddhantmadhur/ocelot-media-server
VERSION ?= $(shell git describe --match 'v[0-9]*' --dirty='.m' --always --tags)

GO_LDFLAGS ?= -w -X ${PKG}/internal.Version=${VERSION}

EXECNAME = ocelot

.PHONY: all build lint test

all: lint test build

build: # build the golang binary for this project
	go build ${BUILD_FLAGS} -ldflags "${GO_LDFLAGS}" -o build/${EXECNAME} ./cmd
