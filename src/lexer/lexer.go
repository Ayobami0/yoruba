package lexer

import (
	"bufio"
	"io"

	"github.com/Ayobami0/yoruba/src/token"
)

type Lexer struct {
	reader bufio.Scanner
}

func New(r io.Reader) *Lexer {
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanBytes)
	s.Scan()

	return &Lexer{reader: *s}
}

func (l *Lexer) NextToken() token.Token {
	var t token.Token

	bList := l.reader.Bytes()

	if len(bList) == 0 {
		t.Literal = ""
		t.Type = token.EOF
		return t
	}

	ch := bList[0]

	for token.IsSpace(ch) {
		b := l.reader.Scan()

		if !b {
			t.Literal = ""
			t.Type = token.EOF
			return t
		}
		ch = l.reader.Bytes()[0]
	}

	switch ch {
	case ',':
		t = token.New(token.COMMA, ch)
	case '}':
		t = token.New(token.R_GROUP, ch)
	case '{':
		t = token.New(token.L_GROUP, ch)
	case '+':
		t = token.New(token.PLUS, ch)
	case '-':
		t = token.New(token.MINUS, ch)
	case '/':
		t = token.New(token.DIVIDE, ch)
	case '*':
		t = token.New(token.TIMES, ch)
	case '"':
		str := l.buildStr()
		t.Literal = str
		t.Type = token.STR
	case '[':
		l.buildComment()
		l.reader.Scan()
		return l.NextToken()
	case 0:
		t = token.New(token.EOF, ch)
	default:
		if token.IsAlpha(ch) {
			ident := l.buildIdent()
			t.Literal = ident
			t.Type = token.LookUp(ident)
		} else if token.IsDigit(ch) {
			num := l.buildInt()
			t.Literal = num
			t.Type = token.NUM
		} else {
			t = token.New(token.ILLEGAL, ch)
			l.reader.Scan()
		}
		return t
	}
	l.reader.Scan()
	return t
}
