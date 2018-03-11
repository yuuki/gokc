package parser

// Token represents the token of the configuration.
type Token struct {
	tok      int
	lit      string
	filename string
	line     int
	column   int
}

type Block struct {
	name  string // <STRING>: vrrp_script <STRING> { ... }
	stmts []Stmt
}

type Stmt struct {
	name string
	val  Value
}

type Value interface{}
