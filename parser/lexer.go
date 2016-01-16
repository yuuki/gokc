package parser

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/mail"
	"strconv"
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
	lex.Mode &^= scanner.ScanInts | scanner.ScanFloats | scanner.ScanChars | scanner.ScanRawStrings | scanner.ScanComments | scanner.SkipComments
	lex.IsIdentRune = isIdentRune
	return &lex
}

func isIdentRune(ch rune, i int) bool {
	return ch == '_' || ch == '.' || ch == '@' || unicode.IsLetter(ch) || unicode.IsDigit(ch)
}

func (l *Lexer) Lex(lval *yySymType) int {
	token := int(l.Scan())
	s := l.TokenText()

	log.Printf("token text: %s\n", s)

	if net.ParseIP(s) != nil {
		token = IPADDR
	}

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
	} else if s == "smtp_server" {
		token = SMTP_SERVER
	} else if s == "smtp_connect_timeout" {
		token = SMTP_CONNECT_TIMEOUT
	} else if s == "router_id" {
		token = ROUTER_ID
	} else if s == "vrrp_instance" {
		token = VRRP_INSTANCE
	} else if s == "interface" {
		token = INTERFACE
	}

	// Is Integer?
	if val, err := strconv.Atoi(s); err == nil {
		token = NUMBER
		if val >= 0 {
			return POSITIVE_INT
		}
	}

	if token == scanner.Ident {
		token = STRING
	}
	return token
}

func (l *Lexer) Error(e string) {
	fmt.Printf("Error Line %d, Pos %d\n", l.Line, l.Column)
	panic(e)
}

