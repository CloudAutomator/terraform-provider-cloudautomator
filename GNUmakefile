BUILD_DIR ?= $(CURDIR)/bin

HOSTNAME=registry.terraform.io
NAMESPACE=penta515
NAME=cloudautomator
BINARY=terraform-provider-${NAME}

GO_OS ?= $(shell go env GOOS)
GO_ARCH ?= $(shell go env GOARCH)
GO_PACKAGES := $(shell go list ./... | grep -v vendor)

build:
	@if [ -z "$(VERSION)" ]; \
	then \
	  echo "Please provide a version. Example: make build VERSION=0.2.2" && exit 1; \
 	fi
	@go build -v -o "${BUILD_DIR}/${BINARY}_v$(VERSION)"

install: build
	@mkdir -p "${HOME}/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/$(VERSION)/${GO_OS}_${GO_ARCH}"
	@mv "${BUILD_DIR}/${BINARY}_v$(VERSION)" "${HOME}/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/$(VERSION)/${GO_OS}_${GO_ARCH}"

clean:
	@if [ -z "$(VERSION)" ]; \
    	then \
         echo "Please provide a version. Example: make clean VERSION=0.2.2" && exit 1; \
     	fi
	@if [ -d "${HOME}/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/$(VERSION)" ]; \
	then \
	  rm -rf "${HOME}/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/$(VERSION)"; \
 	fi

test-unit:
	@go test ${GO_PACKAGES} || exit 1
	@echo ${GO_PACKAGES} | xargs -t -n4 go test $(TESTARGS) -timeout=30s -parallel=4

test-acc-e2e:
	TF_ACC=1 go test -v -cover -timeout 60m ./internal/provider
