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
%token <symbol>	 STRING
%token           LB RB
%token           GLOBALDEFS

%%
configuration:  main_statements configuration | main_statements { }

main_statements:  { }
|global { }

global:	GLOBALDEFS LB STRING RB
%%

func Parse(r io.Reader) Checker {
    l := new(Lexer)
    l.Init(r)
    yyParse(l)
    return l.result
}
