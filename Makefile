.DEFAULT_GOAL := build

BINARY = pkims
VERSION = 0.0.0
DOCKER_REGISTRY = antonjohansson

GO_VERSION = $(shell go version | awk -F\go '{print $$3}' | awk '{print $$1}')
COMMIT = $(shell git rev-parse HEAD)
BUILD_DATE = $(shell date --utc +'%Y-%m-%dT%H:%M:%SZ')
PACKAGE_LIST = $$(go list ./...)
OUTPUT_DIRECTORY = ./build
LDFLAGS = -ldflags "\
	-X ${BINARY}/version.version=${VERSION} \
	-X ${BINARY}/version.goVersion=${GO_VERSION} \
	-X ${BINARY}/version.commit=${COMMIT} \
	-X ${BINARY}/version.buildDate=${BUILD_DATE} \
	"

install:
	go get -v -d ./...

fmt:
	gofmt -s -d -e -w .

vet:
	go vet ${PACKAGE_LIST}

test: install
	go test ${PACKAGE_LIST}

linux: install
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build ${LDFLAGS} -o ${OUTPUT_DIRECTORY}/${BINARY}-linux-amd64 .

darwin: install
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build ${LDFLAGS} -o ${OUTPUT_DIRECTORY}/${BINARY}-darwin-amd64 .

windows: install
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build ${LDFLAGS} -o ${OUTPUT_DIRECTORY}/${BINARY}-windows-amd64.exe .

build: linux darwin windows

docker:
	docker build -t ${DOCKER_REGISTRY}/${BINARY}:${VERSION} .
	docker tag ${DOCKER_REGISTRY}/${BINARY}:${VERSION} ${DOCKER_REGISTRY}/${BINARY}:latest

clean:
	rm -rf ${OUTPUT_DIRECTORY}
