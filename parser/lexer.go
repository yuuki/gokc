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

var SYMBOL_TABLES = map[string]int{
	"{": LB,
	"}": RB,

	"global_defs": GLOBALDEFS,
	"notification_email": NOTIFICATION_EMAIL,
	"notification_email_from": NOTIFICATION_EMAIL_FROM,
	"smtp_server": SMTP_SERVER,
	"smtp_connect_timeout": SMTP_CONNECT_TIMEOUT,
	"router_id": ROUTER_ID,

	"vrrp_instance": VRRP_INSTANCE,
	"interface": INTERFACE,
	"virtual_router_id": VIRTUAL_ROUTER_ID,
	"nopreempt": NOPREEMPT,
	"priority": PRIORITY,
	"advert_int": ADVERT_INT,
	"virtual_ipaddress": VIRTUAL_IPADDRESS,

	"virtual_server": VIRTUAL_SERVER,
	"delay_loop": DELAY_LOOP,
}

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

	if _, err := mail.ParseAddress(s); err == nil {
		token = EMAIL
	}

	if _, ok := SYMBOL_TABLES[s]; ok {
		token = SYMBOL_TABLES[s]
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

