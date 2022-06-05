BUILD_DIR ?= $(CURDIR)/bin
TEST?=$$(go list ./... | grep -v 'vendor')
HOSTNAME=registry.terraform.io
NAMESPACE=penta515
NAME=cloudautomator
VERSION=0.0.1
BINARY=terraform-provider-${NAME}
GO_OS ?= $(shell go env GOOS)
GO_ARCH ?= $(shell go env GOARCH)

build:
	@go build -v -o "${BUILD_DIR}/${BINARY}"

install: build
	@mkdir -p "${HOME}/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/$(VERSION)/${GO_OS}_${GO_ARCH}"
	@mv "${BUILD_DIR}/${BINARY}" "${HOME}/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/$(VERSION)/${GO_OS}_${GO_ARCH}"

clean:
	@if [ -d "${HOME}/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/$(VERSION)" ]; \
	then \
	  rm -rf "${HOME}/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/$(VERSION)"; \
 	fi

test-acc:
	TF_ACC=1 go test -v -cover -timeout 60m ${TEST}
