package parser

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/mail"
	"regexp"
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
	"lb_algo": LB_ALGO,
	"lb_kind": LB_KIND,
	"lvs_sched": LVS_SCHED,
	"lvs_method": LVS_METHOD,
	"rr": RR,
	"wrr": WRR,
	"lc": LC,
	"wlc": WLC,
	"lblc": LBLC,
	"sh": SH,
	"dh": DH,
	"NAT": NAT,
	"TUN": TUN,
	"persistence_timeout": PERSISTENCE_TIMEOUT,
	"protocol": PROTOCOL,
	"TCP": TCP,
	"UDP": UDP,
	"sorry_server": SORRY_SERVER,
	"real_server": REAL_SERVER,
	"weight": WEIGHT,
	"HTTP_GET": HTTP_GET,
	"url": URL,
	"path": PATH,
	"digest": DIGEST,
	"status_code": STATUS_CODE,
	"connect_timeout": CONNECT_TIMEOUT,
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
	return ch == '_' || ch == '.' || ch == '/' || ch == '@' || unicode.IsLetter(ch) || unicode.IsDigit(ch)
}

func (l *Lexer) Lex(lval *yySymType) int {
	token := int(l.Scan())
	s := l.TokenText()

	log.Printf("token text: %s\n", s)

	if net.ParseIP(s) != nil {
		token = IPADDR
	}

	if ok, _ := regexp.MatchString("[[:xdigit:]]{32}", s); ok {
		token = HEX32
	}

	if ok, _ := regexp.MatchString("/([[:alnum:]./-_])*", s); ok {
		token = PATHSTR
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

