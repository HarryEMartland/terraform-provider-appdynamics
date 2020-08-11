
GOBUILD=go build -o terraform-provider-appdynamics

build:
	$(GOBUILD)
	chmod +x ./terraform-provider-appdynamics
test:
	go test ./...  -v
install:
	go get

build-install:
	make build
	mv terraform-provider-appdynamics ~/.terraform.d/plugins/github.com/HarryEMartland/terraform-provider-appdynamics

build-linux:
	GOOS=linux GOARCH=amd64 $(GOBUILD)
build-mac:
	GOOS=darwin GOARCH=amd64 $(GOBUILD)