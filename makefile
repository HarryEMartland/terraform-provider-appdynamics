
GOBUILD=go build -o terraform-provider-appdynamics


BINARY=terraform-provider-appdynamics
VERSION=$(shell git describe --tags --dirty)
PLATFORMS=darwin linux windows
ARCHITECTURES=amd64

build:
	$(GOBUILD)
	chmod +x ./terraform-provider-appdynamics
test:
	go test ./...  -v

r:
	go test ./...  -v -run TestAccAppDDashboard_basic
install:
	go get

build-install:
	make build
	mv terraform-provider-appdynamics ~/.terraform.d/plugins/github.com/HarryEMartland/terraform-provider-appdynamics

build-linux:
	GOOS=linux GOARCH=amd64 $(GOBUILD)
build-mac:
	GOOS=darwin GOARCH=amd64 $(GOBUILD)

build-all:
	$(foreach GOOS, $(PLATFORMS),\
	$(foreach GOARCH, $(ARCHITECTURES), $(shell export GOOS=$(GOOS); export GOARCH=$(GOARCH); export CGO_ENABLED=0; go build -v -o $(BINARY)_$(VERSION)-$(GOOS)-$(GOARCH))))
