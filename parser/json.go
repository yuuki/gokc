package parser

import "encoding/json"

func (conf Configuration) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Blocks []Block `json:blocks`
	}{
		Blocks: conf.Blocks,
	})
}

func (b Block) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Name  string    `json:"name"`
		Stmts []StmtAny `json:"stmts"`
	}{
		Name:  b.name,
		Stmts: b.stmts,
	})
}
