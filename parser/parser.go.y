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

%token <integer> NUM
%token <symbol>	 STRING NOTIFICATION_EMAIL
%token           LB RB
%token           GLOBALDEFS

%%
configuration:  main_statements configuration | main_statements { }

main_statements:  { }
| global { }

global:	GLOBALDEFS LB global_statements RB

global_statements:	global_statement global_statements | global_statement

global_statement:
| NOTIFICATION_EMAIL LB mail_statements RB  { }
| NOTIFICATION_EMAIL STRING  { }

mail_statements:  mail_statement mail_statements |  mail_statement { }

mail_statement:	 STRING	{ }
%%

func Parse(r io.Reader) Checker {
    l := new(Lexer)
    l.Init(r)
    yyParse(l)
    return l.result
}
