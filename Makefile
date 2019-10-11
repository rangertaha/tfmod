# Go parameters
GOCMD=go
GOINSTALL=$(GOCMD) install
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GODOC=$(GOCMD)doc

BINARY_NAME=tfmod
VERSION=$(shell grep -e 'VERSION = ".*"' pkg/version.go | cut -d= -f2 | sed  s/[[:space:]]*\"//g)

.PHONY: help

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-10s\033[0m %s\n", $$1, $$2}'

build: deps build ## Build the binaries for Windows, OSX, and Linux
	mkdir -p builds
	cd cmd; $(GOBUILD) -o ../builds/$(BINARY_NAME) -v
	cd cmd; env GOOS=darwin GOARCH=amd64 $(GOBUILD) -o ../builds/$(BINARY_NAME)-$(VERSION)-darwin-amd64 -v
	cd cmd; env GOOS=linux GOARCH=amd64 $(GOBUILD) -o ../builds/$(BINARY_NAME)-$(VERSION)-linux-amd64 -v
	cd cmd; env GOOS=windows GOARCH=amd64 $(GOBUILD) -o ../builds/$(BINARY_NAME)-$(VERSION)-windows-amd64.exe -v

install: ## Install the binary
	$(GOINSTALL)

deps: ## Install Go library dependencies
	$(GOGET) ./...

docker: image ## Build docker image and upload to docker hub
	docker login

image: clean ## Build docker image
	docker build -t $(BINARY_NAME) .

test: deps ## Run unit test
	$(GOTEST) -v ./...

clean: ## Remove files and images created by the build
	$(GOCLEAN)
	rm -fr builds
	docker rmi $(docker images --filter "tfmod" -q --no-trunc)

run: ## Run server docker image: http://localhost:8080
	docker run -it --rm -p 8080:8080 $(BINARY_NAME)

doc: ## Go documentation
	$(GODOC) -http=:6060

push: image ## Push build to
	docker tag $(BINARY_NAME) $(GCR_HOST)/$(GCP_PROJECT_ID)/$(BINARY_NAME)
	docker push $(GCR_HOST)/$(GCP_PROJECT_ID)/$(BINARY_NAME)
