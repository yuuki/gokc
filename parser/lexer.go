package parser

import (
	"fmt"
	"text/scanner"
)

type Lexer struct {
	scanner.Scanner
	result Checker
}

func (l *Lexer) Lex(lval *yySymType) int {
	token := int(l.Scan())
	s := l.TokenText()

	fmt.Printf("token text: %s\n", s)

	if s == "{" {
		token = LB
	} else if s == "global_defs" {
		token = GLOBALDEFS
	}

	if token == scanner.Ident {
		token = STRING
	}
	if token == scanner.String {
		token = STRING
	}
	if token == scanner.Int {
		token = NUM
	}
	return token
}

func (l *Lexer) Error(e string) {
	fmt.Printf("Error Line %d, Pos %d\n", l.Line, l.Column)
	panic(e)
}
