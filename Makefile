GOFMT=gofmt
GC=go build

VERSION := $(shell git describe --always --tags --long)
BUILD_NODE_PAR = -ldflags "-X git.ont.io/ontid/otf/utils.Version=$(VERSION)" #-race

ARCH=$(shell uname -m)
SRC_FILES = $(shell git ls-files | grep -e .go$ | grep -v _test.go)

agent-otf: $(SRC_FILES)
	$(GC)  $(BUILD_NODE_PAR) -o agent-otf main.go
 


agent-otf-cross: agent-otf-windows agent-otf-linux agent-otf-darwin

agent-otf-windows:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 $(GC) $(BUILD_NODE_PAR) -o agent-otf-windows-amd64.exe main.go

agent-otf-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GC) $(BUILD_NODE_PAR) -o agent-otf-linux-amd64 main.go

agent-otf-darwin:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 $(GC) $(BUILD_NODE_PAR) -o agent-otf-darwin-amd64 main.go

tools-cross: tools-windows tools-linux tools-darwin

format:
	$(GOFMT) -w main.go

clean:
	rm -rf *.8 *.o *.out *.6 *exe
	rm -rf agent-otf agent-otf-*
