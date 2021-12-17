
GOBUILD=go build -o terraform-provider-appdynamics


BINARY=terraform-provider-appdynamics
VERSION=$(shell git describe --tags --dirty)
PLATFORMS=darwin linux windows
ARCHITECTURES=amd64
TEST_VERSION=0.1.0-6

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

test-dashboard-import:
	go clean -testcache
	go test ./...  -v -run TestAccImportExportDashboard_Create

test-widget:
	go clean -testcache
	go test ./...  -v -run TestAccDataSourceAppdService_basic

test-health-rules:
	go clean -testcache
	go test ./...  -v -run TestAccAppDHealthRule

test-appd-service:
	go clean -testcache
	go test ./...  -v -run TestAccDataSource

test-associations:
	go clean -testcache
	go test ./...  -v -run TestAccTierTemplateAssociation_Create

test-collector-create:
	go clean -testcache
	go test ./...  -v -run TestAccAppDCollector_Create

install:
	go get

build-install:
	make build
	mkdir -p ~/.terraform.d/plugins/registry.terraform.io/worldremit/appdynamics/${TEST_VERSION}/linux_amd64
	mv terraform-provider-appdynamics ~/.terraform.d/plugins/registry.terraform.io/worldremit/appdynamics/${TEST_VERSION}/linux_amd64/terraform-provider-appdynamics_v${TEST_VERSION}

build-linux:
	GOOS=linux GOARCH=amd64 $(GOBUILD)
build-mac:
	GOOS=darwin GOARCH=amd64 $(GOBUILD)

build-all:
	$(foreach GOOS, $(PLATFORMS),\
	$(foreach GOARCH, $(ARCHITECTURES), $(shell export GOOS=$(GOOS); export GOARCH=$(GOARCH); export CGO_ENABLED=0; go build -v -o $(BINARY)_$(VERSION)-$(GOOS)-$(GOARCH))))
