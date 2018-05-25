GOBUILD=go build
ENV=env GOOS=linux

all: clean build

build: main

vendor:
	govendor update github.com/pydio/cells-sdk-go


main:
	go build -ldflags "-X github.com/pydio/cells-sdk-go/config.version=${CELLS_VERSION} -X github.com/pydio/cells-sdk-go/config.BuildStamp=`date -u +%Y-%m-%dT%H:%M:%S` -X github.com/pydio/cells-sdk-go/config.BuildRevision=`git rev-parse HEAD`" -o cells-sdk-go main.go

dev:
	go build -ldflags "-X github.com/pydio/cells-sdk-go/config.version=0.2.0 -X github.com/pydio/cells-sdk-go/config.BuildStamp=2018-01-01T00:00:00 -X github.com/pydio/cells-sdk-go/config.BuildRevision=dev" -o cells-sdk-go main.go

cleanall: stop clean rm

clean:
	rm -f cells-sdk-go
