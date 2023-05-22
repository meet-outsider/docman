# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=main
BINARY__MAC=$(BINARY_NAME)_mac
BINARY_LINUX=$(BINARY_NAME)_linux
all:  test build
build:
		$(GOBUILD) -o $(BINARY_NAME) main.go
test:
		$(GOTEST) -v ./test/...
clean:
		$(GOCLEAN)
		rm -f $(BINARY_NAME)
		rm -f $(BINARY__MAC)
		rm -f $(BINARY_LINUX)
run:
		$(GOBUILD) -o $(BINARY_NAME) main.go
		./$(BINARY_NAME)
deps:
		$(GOGET) get
build-linux:
		CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_NAME) main.go
build-mac:
		CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 $(GOBUILD) -o $(BINARY_NAME) main.go
build-all:
		CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 $(GOBUILD) -o $(BINARY__MAC) main.go
		CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_LINUX) main.go
docker:
		docker build -t $(BINARY_NAME) .
