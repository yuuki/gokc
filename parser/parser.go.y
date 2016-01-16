%{
package parser
import (
    "io"
)

type Checker interface{}
%}

%union {
  integer   int
  symbol    string
};

%token <integer> NUMBER POSITIVE_INT
%token <symbol>	 ID STRING EMAIL IPADDR
%token           LB RB
%token           GLOBALDEFS
%token           NOTIFICATION_EMAIL NOTIFICATION_EMAIL_FROM SMTP_SERVER SMTP_CONNECT_TIMEOUT ROUTER_ID
%token           VRRP_INSTANCE
%token           INTERFACE VIRTUAL_ROUTER_ID NOPREEMPT PRIORITY

%%
configuration:  main_statements configuration | main_statements { }

main_statements:  { }
| global { }
| vrrp_instance_block { }

global:	GLOBALDEFS LB global_statements RB

global_statements:	global_statement global_statements | global_statement

global_statement:
| NOTIFICATION_EMAIL LB mail_statements RB  { }
| NOTIFICATION_EMAIL STRING  { }
| NOTIFICATION_EMAIL_FROM EMAIL { }
| SMTP_SERVER IPADDR  { }
| SMTP_SERVER STRING  { }
| SMTP_CONNECT_TIMEOUT POSITIVE_INT { }
| ROUTER_ID STRING { }

vrrp_instance_block: VRRP_INSTANCE STRING LB vrrp_instance_statements RB

vrrp_instance_statements: vrrp_instance_statement vrrp_instance_statements | vrrp_instance_statement

vrrp_instance_statement: { }
| INTERFACE STRING { }
| VIRTUAL_ROUTER_ID STRING { }
| VIRTUAL_ROUTER_ID POSITIVE_INT { }
| NOPREEMPT
| PRIORITY POSITIVE_INT { }

mail_statements:  mail_statement mail_statements |  mail_statement { }

mail_statement:	 STRING	{ }

%%

func Parse(r io.Reader) Checker {
    l := NewLexer(r)
    yyParse(l)
    return l.result
}
