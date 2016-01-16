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
%token <symbol>	 STRING EMAIL IPADDR
%token           LB RB
%token           GLOBALDEFS
%token           NOTIFICATION_EMAIL NOTIFICATION_EMAIL_FROM SMTP_SERVER SMTP_CONNECT_TIMEOUT

%%
configuration:  main_statements configuration | main_statements { }

main_statements:  { }
| global { }

global:	GLOBALDEFS LB global_statements RB

global_statements:	global_statement global_statements | global_statement

global_statement:
| NOTIFICATION_EMAIL LB mail_statements RB  { }
| NOTIFICATION_EMAIL STRING  { }
| NOTIFICATION_EMAIL_FROM EMAIL { }
| SMTP_SERVER IPADDR  { }
| SMTP_SERVER STRING  { }
| SMTP_CONNECT_TIMEOUT POSITIVE_INT { }

mail_statements:  mail_statement mail_statements |  mail_statement { }

mail_statement:	 STRING	{ }
%%

func Parse(r io.Reader) Checker {
    l := NewLexer(r)
    yyParse(l)
    return l.result
}
