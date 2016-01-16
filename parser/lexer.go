package parser

import (
	"fmt"
	"io"
	"log"
	"net/mail"
	"text/scanner"
	"unicode"
)

type Lexer struct {
	scanner.Scanner
	result Checker
}

func NewLexer(src io.Reader) *Lexer {
	var lex Lexer
	lex.Init(src)
	lex.Mode &^= scanner.ScanChars | scanner.ScanRawStrings | scanner.ScanComments | scanner.SkipComments
	lex.IsIdentRune = isIdentRune
	return &lex
}

func isIdentRune(ch rune, i int) bool {
	return ch == '_' || ch == '.' || ch == '@' || unicode.IsLetter(ch) || unicode.IsDigit(ch) && i > 0
}

func (l *Lexer) Lex(lval *yySymType) int {
	token := int(l.Scan())
	s := l.TokenText()

	log.Printf("token text: %s\n", s)

	_, err := mail.ParseAddress(s)
	if err == nil {
		token = EMAIL
	}

	if s == "{" {
		token = LB
	} else if s == "}" {
		token = RB
	} else if s == "global_defs" {
		token = GLOBALDEFS
	} else if s == "notification_email" {
		token = NOTIFICATION_EMAIL
	} else if s == "notification_email_from" {
		token = NOTIFICATION_EMAIL_FROM
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

