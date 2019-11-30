BINARY_NAME = gtoy
GOBIN = ${GOPATH}/bin
VERSION = v0.1.0
REVISION = $(shell git rev-parse --short HEAD)
SRCS = $(shell find . -type f -name '*.go')
LDFLAGS = "-s -w -X \"main.Versioon=${VERSION}\" -X \"main.Revision=${REVISION}\" -extldflags \"-static\""

VERBOSE_FLAG=$(if $(VERBOSE), -v)
u := $(if $(update), -u)

export GO11MODULE=on

all: build

.PHONY: deps
deps:
			go get ${u} -d ${VERBOSE_FLAG}
			go mod tidy

.PHONY: test
test: deps
			go test ${VERBOSE_FLAG} ./...

.PHONY: build
build: deps
			go build ${VERBOSE_FLAG} -o ./bin/${BINARY_NAME} -ldflags=${LDFLAGS}

.PHONY: install
install: build
			cp -f ./bin/${BINARY_NAME} ${GOBIN}/

.PHONY: uninstall
uninstall:
			rm ${GOBIN}/${BINARY_NAME}

.PHONY: release
release: clean
			GOOS=darwin GOARCH=amd64 go build -o ${BINARY_NAME} && zip MacOS.zip ./${BINARY_NAME} && rm -rf ./${BINARY_NAME}
			GOOS=linux GOARCH=amd64 go build -o ${BINARY_NAME} && zip Linux.zip ./${BINARY_NAME} && rm -rf ./${BINARY_NAME}

.PHONY: clean
clean:
			go clean
