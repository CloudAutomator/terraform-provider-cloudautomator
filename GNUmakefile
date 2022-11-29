BUILD_DIR ?= $(CURDIR)/bin
TEST?=$$(go list ./... | grep -v '/vendor/')

HOSTNAME=registry.terraform.io
NAMESPACE=CloudAutomator
NAME=cloudautomator
BINARY=terraform-provider-${NAME}

GO_OS ?= $(shell go env GOOS)
GO_ARCH ?= $(shell go env GOARCH)

build:
	@if [ -z "$(VERSION)" ]; \
	then \
	  echo "Please provide a version. Example: make build VERSION=0.2.2" && exit 1; \
 	fi
	@go build -v -o "${BUILD_DIR}/${BINARY}_v$(VERSION)"

clean:
	@if [ -z "$(VERSION)" ]; \
    	then \
         echo "Please provide a version. Example: make clean VERSION=0.2.2" && exit 1; \
     	fi
	@if [ -d "${HOME}/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/$(VERSION)" ]; \
	then \
	  rm -rf "${HOME}/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/$(VERSION)"; \
 	fi

fmt:
	terraform fmt -recursive examples

install: build
	@mkdir -p "${HOME}/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/$(VERSION)/${GO_OS}_${GO_ARCH}"
	@mv "${BUILD_DIR}/${BINARY}_v$(VERSION)" "${HOME}/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/$(VERSION)/${GO_OS}_${GO_ARCH}"

test:
	TF_ACC= go test $(TEST) -v $(TESTARGS) -timeout 3m -parallel=4

testacc:
	TF_ACC=1 go test $(TEST) -v $(TESTARGS) -timeout 60m

.PHONY: build clean fmt install test testacc
