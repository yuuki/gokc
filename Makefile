BIN := gokc

all: build

yacc: deps
	goyacc -o parser/parser.go -v parser/parser.output parser/parser.go.y

build: yacc
	go build -o $(BIN) ./cmd/...

buildlinux: yacc
	GOOS=linux GOARCH=amd64 make build

test: build
	find ./testdata -type f | xargs -I{} ./$(BIN) -f {}
	find ./keepalived/doc/samples/keepalived.conf.* -type f | xargs -I{} ./$(BIN) -f {}

patch: gobump
	./script/release.sh patch

minor: gobump
	./script/release.sh minor

gobump:
	go get github.com/motemen/gobump/cmd/gobump

deps:
	go get golang.org/x/tools/cmd/goyacc
	go get -d -v ./...

clean:
	go clean

.PHONY: yacc build clean deps test buildlinux patch minor gobump
