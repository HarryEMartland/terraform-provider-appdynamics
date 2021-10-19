
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

test-create:
	go clean -testcache
	go test ./...  -v -run TestAccAppDDashboard_Create

test-update:
	go clean -testcache
	go test ./...  -v -run TestAccAppDDashboard_Update

test-widget:
	go clean -testcache
	go test ./...  -v -run TestAccDataSourceAppdService_basic

a:
	go clean -testcache
	go test ./...  -v -run TestAccDataSource

test-associations:
	go clean -testcache
	go test ./...  -v -run TestAccTierTemplateAssociation_Create

install:
	go get

build-install:
	make build
	mv terraform-provider-appdynamics ~/.terraform.d/plugins/registry.terraform.io/hashicorp/appdynamics/0.1.0/linux_amd64/terraform-provider-appdynamics_v0.1.0

build-linux:
	GOOS=linux GOARCH=amd64 $(GOBUILD)
build-mac:
	GOOS=darwin GOARCH=amd64 $(GOBUILD)

build-all:
	$(foreach GOOS, $(PLATFORMS),\
	$(foreach GOARCH, $(ARCHITECTURES), $(shell export GOOS=$(GOOS); export GOARCH=$(GOARCH); export CGO_ENABLED=0; go build -v -o $(BINARY)_$(VERSION)-$(GOOS)-$(GOARCH))))
