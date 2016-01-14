all: build

yacc:
	go tool yacc -o parser/parser.go -v parser/parser.output parser/parser.go.y

build: yacc
	go build .

