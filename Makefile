GOFMT=gofmt
GC=go build

VERSION := $(shell git describe --always --tags --long)
BUILD_NODE_PAR = -ldflags "-X github.com/ontio/mercury/utils.Version=$(VERSION)" #-race

ARCH=$(shell uname -m)
SRC_FILES = $(shell git ls-files | grep -e .go$ | grep -v _test.go)

agent-mercury: $(SRC_FILES)
	$(GC)  $(BUILD_NODE_PAR) -o agent-mercury main.go
 


agent-mercury-cross: agent-mercury-windows agent-mercury-linux agent-mercury-darwin

agent-mercury-windows:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 $(GC) $(BUILD_NODE_PAR) -o agent-mercury-windows-amd64.exe main.go

agent-mercury-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GC) $(BUILD_NODE_PAR) -o agent-mercury-linux-amd64 main.go

agent-mercury-darwin:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 $(GC) $(BUILD_NODE_PAR) -o agent-mercury-darwin-amd64 main.go

tools-cross: tools-windows tools-linux tools-darwin

format:
	$(GOFMT) -w main.go

clean:
	rm -rf *.8 *.o *.out *.6 *exe
	rm -rf agent-mercury agent-mercury-*
