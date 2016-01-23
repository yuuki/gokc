%{
package parser
import (
    "io"
)

const MAX_PORT_NUM int = 65535

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
%token           INTERFACE VIRTUAL_ROUTER_ID NOPREEMPT PRIORITY ADVERT_INT VIRTUAL_IPADDRESS
%token           VIRTUAL_SERVER
%token           DELAY_LOOP LB_ALGO LB_KIND LVS_SCHED LVS_METHOD RR WRR LC WLC LBLC SH DH NAT DR TUN PERSISTENCE_TIMEOUT PROTOCOL TCP UDP SORRY_SERVER

%%
configuration:  main_statements configuration | main_statements { }

main_statements:  { }
| global { }
| vrrp_instance_block { }
| virtual_server_block { }

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
| ADVERT_INT POSITIVE_INT { }
| VIRTUAL_IPADDRESS LB vips RB

virtual_server_block: VIRTUAL_SERVER iporfw LB virtual_server_statements RB

virtual_server_statements: virtual_server_statement virtual_server_statements | virtual_server_statement

virtual_server_statement: { }
| DELAY_LOOP POSITIVE_INT { }
| LB_ALGO lb_algo { }
| LB_KIND lb_kind { }
| LVS_SCHED lb_algo { }
| LVS_METHOD lb_kind { }
| PERSISTENCE_TIMEOUT POSITIVE_INT { }
| PROTOCOL protocol { }
| SORRY_SERVER IPADDR POSITIVE_INT { }

iporfw: { }
| IPADDR { }
| IPADDR POSITIVE_INT { }

lb_algo: { }
| RR   { }
| WRR  { }
| LC   { }
| WLC  { }
| LBLC { }
| SH   { }
| DH   { }

lb_kind: { }
| NAT	{ }	
| DR  { }
| TUN	{ }

protocol: { }
| TCP { }
| UDP { }

vips: vip vips | vip { }

vip: { }
| IPADDR { }

mail_statements:  mail_statement mail_statements |  mail_statement { }

mail_statement:	 STRING	{ }

%%

func Parse(r io.Reader) Checker {
    l := NewLexer(r)
    yyParse(l)
    return l.result
}

