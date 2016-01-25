BIN := gokc

all: build

yacc: deps
	go tool yacc -o parser/parser.go -v parser/parser.output parser/parser.go.y

build: yacc
	go build -o $(BIN) ./cmd

test: build
	find ./testdata -type f | xargs -I{} ./$(BIN) -f {}
	find ./keepalived/doc/samples/keepalived.conf.* -type f | xargs -I{} ./$(BIN) -f {}

deps:
	go get -d -v ./...

clean:
	go clean

.PHONY: yacc build clean deps test
