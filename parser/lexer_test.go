package parser

import (
	"bytes"
	"fmt"
	"os"
	"reflect"
	"testing"
)

const (
	testdataDir = "../testdata/"
)

func TestNewTokenizer(t *testing.T) {
	buf := bytes.NewReader([]byte("hello world"))
	tokenizer := NewTokenizer(buf, "buffer")
	if tokenizer.scanner.Pos().Offset != 0 {
		t.Errorf("scanner pos got %v, want %v", tokenizer.scanner.Pos().Offset, 0)
	}
	if tokenizer.filename != "buffer" {
		t.Errorf("filename got %v, want %v", tokenizer.filename, "buffer")
	}
}

var (
	tokenMap = map[string][]int{
		"unicast_mcast": {
			VRRP_INSTANCE, STRING, LB, MCAST_SRC_IP, IPV4,
			UNICAST_SRC_IP, IPV4, UNICAST_PEER, LB,
			IPV4, IPV4, IPV4, RB, RB,
		},
	}
)

func TestTokenizer_NextAll(t *testing.T) {
	tests := []struct {
		filename string
		toks     []int
	}{
		{
			"unicast_mcast.conf", tokenMap["unicast_mcast"],
		},
	}

	for _, tt := range tests {
		path := testdataDir + tt.filename
		f, err := os.Open(path)
		if err != nil {
			panic(err)
		}

		tokenizer := NewTokenizer(f, path)
		tokens, err := tokenizer.NextAll()
		if err != nil {
			t.Errorf("shoud not raise error: %v", err)
		}
		fmt.Printf("%+v", len(tokens))
		for i, token := range tokens {
			got, want := token.value, tt.toks[i]
			if got != want {
				t.Errorf("token got %v, want %v (%v)", got, want, i)
			}
		}
	}
}

func TestNewLexer(t *testing.T) {
	tokens := []*Token{
		{VRRP_INSTANCE, "dummy.conf", 2, 10},
	}
	l := NewLexer(tokens)
	if !reflect.DeepEqual(l.tokens, tokens) {
		t.Errorf("NewLexer got %v, want %v", l.tokens, tokens)
	}
	if l.pos != -1 {
		t.Errorf("pos got %v, want %v", l.pos, -1)
	}
	if l.e != nil {
		t.Errorf("lexer error: got %v, want %v", l.e)
	}
}

func TestLexer_Lex(t *testing.T) {
	tests := []struct {
		desc   string
		tokens []*Token
	}{
		{"unicast_mcast", generateTestTokens(tokenMap["unicast_mcast"])},
	}
	for _, tt := range tests {
		l := NewLexer(tt.tokens)
		ret := yyParse(l)
		if ret != 0 {
			t.Errorf("yyParse got %v, want %v", ret, 0)
		}
		if l.e != nil {
			t.Errorf("yyParse error shoud be nil: %v", l.e)
		}
	}
}
