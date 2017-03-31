package parser

func generateTestTokens(toks []int) []*Token {
	tokens := make([]*Token, 0, len(toks))
	for _, tok := range toks {
		tokens = append(tokens, &Token{
			value:    tok,
			filename: "dummy.conf", // Whatever is fine
			line:     10,           // Whatever is fine
			column:   2,            // Whatever is fine
		})
	}
	return tokens
}
