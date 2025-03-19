MAJOR_VERSION = 0
MINOR_VERSION = 1
PATCH_VERSION = 0

PKG := github.com/siddhantmadhur/ocelot-media-server
VERSION ?= $(shell git describe --match 'v[0-9]*' --dirty='.m' --always --tags)

GO_LDFLAGS ?= -w -X ${PKG}/internal.Version=${VERSION}

EXECNAME = ocelotmediaserver

.PHONY: all build lint test

all: lint test build

build: # build the golang binary for this project
	go build ${BUILD_FLAGS} -ldflags "${GO_LDFLAGS}" -o build/${EXECNAME} ./cmd
