package lexer

import (
	"bytes"
	"testing"

	"github.com/Ayobami0/yoruba/src/token"
)

func TestLexer(t *testing.T) {
	b := bytes.NewBufferString(`
jeki name je "Ayobami"

ise sum 2,3 se
  da 2 + 3 pada
pari

ti name baje "Ayobami" lehinna
  ko sum 2,3
abi 
    ti name kobaje "Oludemi" lehinna
      ko sum 3,4
    abi lehinna
      ko sum 5,6
    pari
pari
`)

	l := New(b)

	tests := []struct {
		eType    token.TokenType
		eLiteral string
	}{
		{eType: token.LET, eLiteral: "jeki"},
		{eType: token.IDENT, eLiteral: "name"},
		{eType: token.ASSIGNMENT, eLiteral: "je"},
		{eType: token.STR, eLiteral: "Ayobami"},
		{eType: token.FUNCTION, eLiteral: "ise"},
		{eType: token.IDENT, eLiteral: "sum"},
		{eType: token.NUM, eLiteral: "2"},
		{eType: token.COMMA, eLiteral: ","},
		{eType: token.NUM, eLiteral: "3"},
		{eType: token.EXECUTE, eLiteral: "se"},
		{eType: token.RTN_PREFIX, eLiteral: "da"},
		{eType: token.NUM, eLiteral: "2"},
		{eType: token.PLUS, eLiteral: "+"},
		{eType: token.NUM, eLiteral: "3"},
		{eType: token.RTN_SURFIX, eLiteral: "pada"},
		{eType: token.END, eLiteral: "pari"},
		{eType: token.IF, eLiteral: "ti"},
		{eType: token.IDENT, eLiteral: "name"},
		{eType: token.EQL, eLiteral: "baje"},
		{eType: token.STR, eLiteral: "Ayobami"},
		{eType: token.THEN, eLiteral: "lehinna"},
		{eType: token.IDENT, eLiteral: "ko"},
		{eType: token.IDENT, eLiteral: "sum"},
		{eType: token.NUM, eLiteral: "2"},
		{eType: token.COMMA, eLiteral: ","},
		{eType: token.NUM, eLiteral: "3"},
		{eType: token.ELSE, eLiteral: "abi"},
		{eType: token.IF, eLiteral: "ti"},
		{eType: token.IDENT, eLiteral: "name"},
		{eType: token.NOTEQL, eLiteral: "kobaje"},
		{eType: token.STR, eLiteral: "Oludemi"},
		{eType: token.THEN, eLiteral: "lehinna"},
		{eType: token.IDENT, eLiteral: "ko"},
		{eType: token.IDENT, eLiteral: "sum"},
		{eType: token.NUM, eLiteral: "3"},
		{eType: token.COMMA, eLiteral: ","},
		{eType: token.NUM, eLiteral: "4"},
		{eType: token.ELSE, eLiteral: "abi"},
		{eType: token.THEN, eLiteral: "lehinna"},
		{eType: token.IDENT, eLiteral: "ko"},
		{eType: token.IDENT, eLiteral: "sum"},
		{eType: token.NUM, eLiteral: "5"},
		{eType: token.COMMA, eLiteral: ","},
		{eType: token.NUM, eLiteral: "6"},
		{eType: token.END, eLiteral: "pari"},
		{eType: token.END, eLiteral: "pari"},
		{eType: token.EOF, eLiteral: ""},
	}

	for i, tt := range tests {
		tok := l.NextToken()
		if tok.Type != tt.eType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.eType, tok.Type)
		}
		if tok.Literal != tt.eLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.eLiteral, tok.Literal)
		}
	}
}
