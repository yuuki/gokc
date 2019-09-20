package parser

import (
	"fmt"
	"io"
	"net"
	"net/mail"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"text/scanner"
	"unicode"

	"github.com/yuuki/gokc/log"
)

const (
	EOF = 0
)

var symbolTables = map[string]int{
	"{": LB,
	"}": RB,

	"global_defs":                     GLOBALDEFS,
	"notification_email":              NOTIFICATION_EMAIL,
	"notification_email_from":         NOTIFICATION_EMAIL_FROM,
	"smtp_server":                     SMTP_SERVER,
	"smtp_connect_timeout":            SMTP_CONNECT_TIMEOUT,
	"router_id":                       ROUTER_ID,
	"lvs_id":                          LVS_ID,
	"vrrp_mcast_group4":               VRRP_MCAST_GROUP4,
	"vrrp_mcast_group6":               VRRP_MCAST_GROUP6,
	"vrrp_garp_master_delay":          VRRP_GARP_MASTER_DELAY,
	"vrrp_garp_master_repeat":         VRRP_GARP_MASTER_REPEAT,
	"vrrp_garp_master_refresh":        VRRP_GARP_MASTER_REFRESH,
	"vrrp_garp_master_refresh_repeat": VRRP_GARP_MASTER_REFRESH_REPEAT,
	"vrrp_version":                    VRRP_VERSION,
	"enable_script_security":          ENABLE_SCRIPT_SECURITY,

	"static_ipaddress": STATIC_IPADDRESS,
	"static_routes":    STATIC_ROUTES,
	"static_rules":     STATIC_RULES,

	"vrrp_sync_group": VRRP_SYNC_GROUP,
	"group":           GROUP,

	"vrrp_instance":              VRRP_INSTANCE,
	"use_vmac":                   USE_VMAC,
	"version":                    VERSION,
	"vmac_xmit_base":             VMAC_XMIT_BASE,
	"native_ipv6":                NATIVE_IPV6,
	"interface":                  INTERFACE,
	"mcast_src_ip":               MCAST_SRC_IP,
	"unicast_src_ip":             UNICAST_SRC_IP,
	"unicast_peer":               UNICAST_PEER,
	"lvs_sync_daemon_interface":  LVS_SYNC_DAEMON_INTERFACE,
	"virtual_router_id":          VIRTUAL_ROUTER_ID,
	"nopreempt":                  NOPREEMPT,
	"preempt_delay":              PREEMPT_DELAY,
	"priority":                   PRIORITY,
	"advert_int":                 ADVERT_INT,
	"virtual_ipaddress":          VIRTUAL_IPADDRESS,
	"virtual_ipaddress_excluded": VIRTUAL_IPADDRESS_EXCLUDED,
	"virtual_routes":             VIRTUAL_ROUTES,
	"state":                      STATE,
	"MASTER":                     MASTER,
	"BACKUP":                     BACKUP,
	"garp_master_delay":          GARP_MASTER_DELAY,
	"garp_master_repeat":         GARP_MASTER_REPEAT,
	"garp_master_refresh":        GARP_MASTER_REFRESH,
	"garp_master_refresh_repeat": GARP_MASTER_REFRESH_REPEAT,
	"smtp_alert":                 SMTP_ALERT,
	"authentication":             AUTHENTICATION,
	"auth_type":                  AUTH_TYPE,
	"auth_pass":                  AUTH_PASS,
	"PASS":                       PASS,
	"AH":                         AH,
	"label":                      LABEL,
	"dev":                        DEV,
	"scope":                      SCOPE,
	"site":                       SITE,
	"link":                       LINK,
	"host":                       HOST,
	"nowhere":                    NOWHERE,
	"global":                     GLOBAL,
	"brd":                        BRD,
	"src":                        SRC,
	"from":                       FROM,
	"to":                         TO,
	"via":                        VIA,
	"gw":                         GW,
	"or":                         OR,
	"table":                      TABLE,
	"metric":                     METRIC,
	"blackhole":                  BLACKHOLE,
	"track_interface":            TRACK_INTERFACE,
	"track_script":               TRACK_SCRIPT,
	"dont_track_primary":         DONT_TRACK_PRIMARY,
	"notify_master":              NOTIFY_MASTER,
	"notify_backup":              NOTIFY_BACKUP,
	"notify_fault":               NOTIFY_FAULT,
	"notify_stop":                NOTIFY_STOP,
	"notify":                     NOTIFY,

	"vrrp_script": VRRP_SCRIPT,
	"script":      SCRIPT,
	"interval":    INTERVAL,
	"timeout":     TIMEOUT,
	"fall":        FALL,
	"rise":        RISE,

	"virtual_server_group": VIRTUAL_SERVER_GROUP,
	"fwmark":               FWMARK,

	"virtual_server":      VIRTUAL_SERVER,
	"delay_loop":          DELAY_LOOP,
	"lb_algo":             LB_ALGO,
	"lb_kind":             LB_KIND,
	"lvs_sched":           LVS_SCHED,
	"lvs_method":          LVS_METHOD,
	"rr":                  RR,
	"wrr":                 WRR,
	"lc":                  LC,
	"wlc":                 WLC,
	"fo":                  FO,
	"ovf":                 OVF,
	"lblc":                LBLC,
	"lblcr":               LBLCR,
	"sh":                  SH,
	"dh":                  DH,
	"sed":                 SED,
	"nq":                  NQ,
	"NAT":                 NAT,
	"DR":                  DR,
	"TUN":                 TUN,
	"persistence_timeout": PERSISTENCE_TIMEOUT,
	"protocol":            PROTOCOL,
	"TCP":                 TCP,
	"UDP":                 UDP,
	"sorry_server":        SORRY_SERVER,
	"real_server":         REAL_SERVER,
	"weight":              WEIGHT,
	"inhibit_on_failure":  INHIBIT_ON_FAILURE,
	"TCP_CHECK":           TCP_CHECK,
	"HTTP_GET":            HTTP_GET,
	"SSL_GET":             SSL_GET,
	"SMTP_CHECK":          SMTP_CHECK,
	"DNS_CHECK":           DNS_CHECK,
	"MISC_CHECK":          MISC_CHECK,
	"url":                 URL,
	"path":                PATH,
	"digest":              DIGEST,
	"status_code":         STATUS_CODE,
	"connect_timeout":     CONNECT_TIMEOUT,
	"connect_port":        CONNECT_PORT,
	"connect_ip":          CONNECT_IP,
	"bindto":              BINDTO,
	"bind_port":           BIND_PORT,
	"retry":               RETRY,
	"helo_name":           HELO_NAME,
	"delay_before_retry":  DELAY_BEFORE_RETRY,
	"type":                TYPE,
	"name":                NAME,
	"misc_path":           MISC_PATH,
	"misc_timeout":        MISC_TIMEOUT,
	"warmup":              WARMUP,
	"misc_dynamic":        MISC_DYNAMIC,
	"nb_get_retry":        NB_GET_RETRY,
	"virtualhost":         VIRTUALHOST,
	"alpha":               ALPHA,
	"omega":               OMEGA,
	"quorum":              QUORUM,
	"hysteresis":          HYSTERESIS,
	"quorum_up":           QUORUM_UP,
	"quorum_down":         QUORUM_DOWN,
}

type Tokenizer struct {
	scanner  scanner.Scanner
	filename string
}

func NewTokenizer(src io.Reader, filename string) *Tokenizer {
	var t Tokenizer
	t.scanner.Init(src)
	t.scanner.Mode &^= scanner.ScanInts | scanner.ScanFloats | scanner.ScanChars | scanner.ScanRawStrings | scanner.ScanComments | scanner.SkipComments
	t.scanner.IsIdentRune = isIdentRune
	t.filename = filename
	return &t
}

func (t *Tokenizer) NextAll() ([]*Token, error) {
	var result []*Token

	for {
		token, s := t.scanNextToken()

		for s == "include" {
			token, s = t.scanNextToken()

			tokens, err := t.scanInclude(s)
			if err != nil {
				return nil, err
			}
			result = append(result, tokens...)

			token, s = t.scanNextToken()
		}

		if token == scanner.EOF {
			break
		}

		if token == scanner.Ident || token == scanner.String {
			token = STRING
		}

		if _, err := strconv.Atoi(s); err == nil {
			token = NUMBER
		}

		if ip := net.ParseIP(s); ip != nil {
			if ip.To4() != nil {
				token = IPV4
			} else if ip.To16() != nil {
				token = IPV6
			} else {
				log.Infof("warning: %s may be IP address?", s)
			}
		}

		if _, _, err := net.ParseCIDR(s); err == nil {
			token = IP_CIDR
		}

		// IPADDR_RANGE(XXX.YYY.ZZZ.WWW-VVV)
		if ss := strings.Split(s, "-"); len(ss) == 2 {
			if net.ParseIP(ss[0]) != nil {
				if ok, _ := regexp.MatchString(`^[\d]{1,3}$`, ss[1]); ok {
					token = IPADDR_RANGE
				}
			}
		}

		if ok, _ := regexp.MatchString(`^[[:xdigit:]]{32}$`, s); ok {
			token = HEX32
		}

		if ok, _ := regexp.MatchString(`/^([[:alnum:]./-_])*`, s); ok {
			token = PATHSTR
		}

		if _, err := mail.ParseAddress(s); err == nil {
			token = EMAIL
		}

		if _, ok := symbolTables[s]; ok {
			token = symbolTables[s]
		}

		result = append(result, &Token{
			value:    token,
			filename: t.filename,
			line:     t.scanner.Line,
			column:   t.scanner.Column,
		})
	}

	return result, nil
}

func skipComments(scanner *scanner.Scanner) {
	ch := scanner.Next()
	for ch != '\n' && ch >= 0 {
		ch = scanner.Next()
	}
}

func (t *Tokenizer) scanNextToken() (int, string) {
	token := int(t.scanner.Scan())
	s := t.scanner.TokenText()

	for s == "!" || s == "#" {
		skipComments(&t.scanner)

		token = int(t.scanner.Scan())
		s = t.scanner.TokenText()
	}

	log.Debugf("token text: %s\n", s)

	return token, s
}

func (t *Tokenizer) scanInclude(rawfilename string) ([]*Token, error) {
	curDir, err := filepath.Abs(".")
	if err != nil {
		return nil, err
	}

	baseDir := filepath.Dir(t.filename)
	os.Chdir(baseDir)
	defer os.Chdir(curDir)

	rawpaths, err := filepath.Glob(rawfilename)
	if err != nil {
		return nil, err
	}

	if len(rawpaths) < 1 {
		return nil, fmt.Errorf("warning: %s: No such file or directory", rawfilename)
	}

	var result []*Token
	for _, rawpath := range rawpaths {
		log.Verbosef("--> Parsing ... %s\n", rawpath)

		f, err := os.Open(rawpath)
		if err != nil {
			return nil, err
		}

		child := NewTokenizer(f, rawpath)
		tokens, err := child.NextAll()
		if err != nil {
			return nil, err
		}
		result = append(result, tokens...)

		f.Close()
	}

	return result, nil
}

type Token struct {
	value    int
	filename string
	line     int
	column   int
}

type Lexer struct {
	tokens []*Token
	pos    int
	e      error
}

type Error struct {
	Message  string
	Filename string
	Line     int
	Column   int
}

func (e *Error) Error() string {
	return e.Message
}

func NewLexer(tokens []*Token) *Lexer {
	return &Lexer{tokens: tokens, pos: -1}
}

func isIdentRune(ch rune, i int) bool {
	return ch == '_' || ch == '.' || ch == '/' || ch == ':' || ch == '-' || ch == '+' || ch == '*' || ch == '?' || ch == '=' || ch == '&' || ch == '@' || unicode.IsLetter(ch) || unicode.IsDigit(ch)
}

func (l *Lexer) curToken() *Token {
	return l.tokens[l.pos]
}

func (l *Lexer) nextToken() *Token {
	l.pos++
	return l.tokens[l.pos]
}

func (l *Lexer) Lex(lval *yySymType) int {
	if (len(l.tokens) - 1) == l.pos {
		return EOF
	}
	token := l.nextToken()
	return token.value
}

func (l *Lexer) Error(msg string) {
	token := l.curToken()
	l.e = &Error{
		Filename: token.filename,
		Line:     token.line,
		Column:   token.column,
		Message:  msg,
	}
}

func Parse(src io.Reader, filename string) error {
	yyErrorVerbose = true
	t := NewTokenizer(src, filename)
	tokens, err := t.NextAll()
	if err != nil {
		return err
	}
	l := NewLexer(tokens)
	if ret := yyParse(l); ret != 0 {
		return l.e
	}
	return l.e
}
