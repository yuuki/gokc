package parser

import (
	"io"
	"net"
	"net/mail"
	"regexp"
	"strconv"
	"text/scanner"
	"unicode"

	"github.com/yuuki1/gokc/log"
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
	"lvs_id": LVS_ID,

	"vrrp_sync_group": VRRP_SYNC_GROUP,
	"group": GROUP,

	"vrrp_instance": VRRP_INSTANCE,
	"interface": INTERFACE,
	"lvs_sync_daemon_interface": LVS_SYNC_DAEMON_INTERFACE,
	"virtual_router_id": VIRTUAL_ROUTER_ID,
	"nopreempt": NOPREEMPT,
	"priority": PRIORITY,
	"advert_int": ADVERT_INT,
	"virtual_ipaddress": VIRTUAL_IPADDRESS,
	"state": STATE,
	"MASTER": MASTER,
	"BACKUP": BACKUP,
	"garp_master_delay": GARP_MASTER_DELAY,
	"smtp_alert": SMTP_ALERT,
	"authentication": AUTHENTICATION,
	"auth_type": AUTH_TYPE,
	"auth_pass": AUTH_PASS,
	"PASS": PASS,
	"AH": AH,
	"label": LABEL,
	"dev": DEV,
	"brd": BRD,
	"track_interface": TRACK_INTERFACE,
	"track_script": TRACK_SCRIPT,
	"notify_master": NOTIFY_MASTER,
	"notify_backup": NOTIFY_BACKUP,
	"notify_fault": NOTIFY_FAULT,
	"notify_stop": NOTIFY_STOP,
	"notify": NOTIFY,

	"vrrp_script": VRRP_SCRIPT,
	"script": SCRIPT,
	"interval": INTERVAL,
	"fall": FALL,
	"rise": RISE,

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
	"nb_get_retry": NB_GET_RETRY,
	"delay_before_retry": DELAY_BEFORE_RETRY,

	"include": INCLUDE,
}

type Lexer struct {
	scanner.Scanner
	errors []LexError
}

type LexError struct {
	Message string
	Line int
	Column int
}

func NewLexer(src io.Reader) *Lexer {
	var lex Lexer
	lex.Init(src)
	lex.Mode &^= scanner.ScanInts | scanner.ScanFloats | scanner.ScanChars | scanner.ScanRawStrings | scanner.ScanComments | scanner.SkipComments
	lex.IsIdentRune = isIdentRune
	return &lex
}

func isIdentRune(ch rune, i int) bool {
	return ch == '_' || ch == '.' || ch == '/' || ch == ':' || ch == '-' || ch == '+' || ch == '*' || ch == '@' || unicode.IsLetter(ch) || unicode.IsDigit(ch)
}

func (l *Lexer) scanNextToken() (int, string) {
	token := int(l.Scan())
	s := l.TokenText()

	for s == "!" || s == "#" {
		l.skipComments()

		token = int(l.Scan())
		s = l.TokenText()
	}

	log.Debugf("token text: %s\n", s)

	return token, s
}

func (l *Lexer) skipComments() {
	ch := l.Next()
	for ch != '\n' && ch >= 0 {
		ch = l.Next()
	}
}

func (l *Lexer) Lex(lval *yySymType) int {
	token, s := l.scanNextToken()

	if token == scanner.Ident || token == scanner.String {
		token = STRING
	}

	if _, err := strconv.Atoi(s); err == nil {
		token = NUMBER
	}

	if net.ParseIP(s) != nil {
		token = IPADDR
	}

	if _, _, err := net.ParseCIDR(s); err == nil {
		token = IP_CIDR
	}

	if ok, _ := regexp.MatchString("[[:xdigit:]]{32}", s); ok {
		token = HEX32
	}

	if ok, _ := regexp.MatchString("/^([[:alnum:]./-_])*", s); ok {
		token = PATHSTR
	}

	if _, err := mail.ParseAddress(s); err == nil {
		token = EMAIL
	}

	if _, ok := SYMBOL_TABLES[s]; ok {
		token = SYMBOL_TABLES[s]
	}

	return token
}

func (l *Lexer) Error(e string) {
	lexerr := LexError{Line: l.Line, Column: l.Column, Message: e}
	l.errors = append(l.errors, lexerr)
}

