
GOBUILD=go build -o terraform-provider-appd

build:
	$(GOBUILD)
	chmod +x ./terraform-provider-appd
test:
	go test ./appd -v
install:
	go get

build-linux:
	GOOS=linux GOARCH=amd64 $(GOBUILD)
build-mac:
	GOOS=darwin GOARCH=amd64 $(GOBUILD)