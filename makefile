
GOBUILD=go build -o terraform-provider-appdynamics

build:
	$(GOBUILD)
	chmod +x ./terraform-provider-appdynamics
test:
	go test ./appdynamics -v
install:
	go get

build-linux:
	GOOS=linux GOARCH=amd64 $(GOBUILD)
build-mac:
	GOOS=darwin GOARCH=amd64 $(GOBUILD)