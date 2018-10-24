.PHONY: all build version

.EXPORT_ALL_VARIABLES:

APPVER=v1.2.0
CGO_ENABLED=0
GO111MODULE=on

GITVER=$(shell git rev-parse --short HEAD)
GOVER=$(shell go version)
BUILDTIME=$(shell date +%FT%T%z)

HTTP_PROXY=socks5://127.0.0.1:1080/
HTTPS_PROXY=socks5://127.0.0.1:1080/

all: build version

build:
	go build -v -ldflags "-X 'github.com/openqt/diinfo/inspector.AppVersion=${APPVER}' -X 'github.com/openqt/diinfo/inspector.GoVersion=${GOVER}' -X 'github.com/openqt/diinfo/inspector.GitVersion=${GITVER}' -X 'github.com/openqt/diinfo/inspector.BuildTime=${BUILDTIME}'"

version:
	./diinfo version

update:
	#godep save -v
	go mod vendor -v
