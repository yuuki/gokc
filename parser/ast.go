package parser

// Token represents the token of the configuration.
type Token struct {
	tok      int
	lit      string
	filename string
	line     int
	column   int
}

type Configuration struct {
	Blocks []BlockAny
}

type BlockAny interface{} //FIXME

type Block struct {
	name  string // <STRING>: vrrp_script <STRING> { ... }
	stmts []StmtAny
}

type BlockArgs struct { // virtual_server <IP ADDRESS> <PORT> { }
	Name  string    `json:"name"`
	Args  []string  `json:"args"`
	Stmts []StmtAny `json:"stmts"`
}

type SubBlockArgs BlockArgs // real_server <IP ADDRESS> <PORT> { ... }

type StmtAny interface{} //FIXME

type Stmt struct { // state: MASTER
	Key string `json:"key"`
	Val Value  `json:"value"`
}

type StmtMulti struct { // virtual_ipaddress { vips }
	Key  string  `json:"key"`
	Vals []Value `json:"values"`
}

type StmtValue = Value // static_ipaddress { address_options }

type Value interface{}

// VIPAddr represents virtual_ipaddress
type VIPAddr struct {
	Addr string `json:"addr"`
}
