.DEFAULT_GOAL := dev

BINARY = kertificate
MODULE = ${BINARY}.io/${BINARY}
VERSION = 0.0.0
DOCKER_REGISTRY = antonjohansson

GO_VERSION = $(shell go version | awk -F\go '{print $$3}' | awk '{print $$1}')
COMMIT = $(shell git rev-parse HEAD)
BUILD_DATE = $(shell date --utc +'%Y-%m-%dT%H:%M:%SZ')
PACKAGE_LIST = $$(go list ./...)
OUTPUT_DIRECTORY = ./bin
WEB_DIRECTORY = ./web
LDFLAGS = -ldflags "\
	-X ${MODULE}/pkg/version.version=${VERSION} \
	-X ${MODULE}/pkg/version.goVersion=${GO_VERSION} \
	-X ${MODULE}/pkg/version.commit=${COMMIT} \
	-X ${MODULE}/pkg/version.buildDate=${BUILD_DATE} \
	"

install: install-web install-go

install-web:
	npm install --prefix ${WEB_DIRECTORY}

install-go:
	go get -v -d ./...

generate-go:
	go generate ./pkg/api/static

fmt:
	gofmt -s -d -e -w .

vet:
	go vet ${PACKAGE_LIST}

test: install-go
	go test ${PACKAGE_LIST}

linux: install-go generate-go
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build ${LDFLAGS} -o ${OUTPUT_DIRECTORY}/${BINARY}-linux-amd64 ./cmd/kertificate/*

darwin: install-go generate-go
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build ${LDFLAGS} -o ${OUTPUT_DIRECTORY}/${BINARY}-darwin-amd64 ./cmd/kertificate/*

windows: install-go generate-go
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build ${LDFLAGS} -o ${OUTPUT_DIRECTORY}/${BINARY}-windows-amd64.exe ./cmd/kertificate/*

build: build-web build-go

build-web: install-web
	npm run --prefix ${WEB_DIRECTORY} build

build-go: linux darwin windows

build-linux: build-web linux

docker:
	docker build -t ${DOCKER_REGISTRY}/${BINARY}:${VERSION} .
	docker tag ${DOCKER_REGISTRY}/${BINARY}:${VERSION} ${DOCKER_REGISTRY}/${BINARY}:latest

dev-go: install-go
	go get github.com/cespare/reflex
	reflex -sr '\.go$$' -- go run -tags dev cmd/kertificate/* start

dev-web: install-web
	npm start --prefix ${WEB_DIRECTORY}

dev: dev-web dev-go

clean:
	rm -f ./pkg/api/static/static_data.go
	rm -rf ${OUTPUT_DIRECTORY}
	rm -rf ${WEB_DIRECTORY}/dist
	rm -rf ${WEB_DIRECTORY}/node_modules
