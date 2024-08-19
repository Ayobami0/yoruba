package lexer

import (
	"strings"

	"github.com/Ayobami0/yoruba/src/token"
)

func (l *Lexer) buildStr() string {
	var str strings.Builder
	l.reader.Scan()
	for {
    c := l.reader.Bytes()[0]
		if c == '\n' || c == '"' {
			break
		}
		str.WriteByte(c)
		l.reader.Scan()
	}
	return str.String()
}

func (l *Lexer) buildIdent() string {
	var ident strings.Builder

	for {
    c := l.reader.Bytes()[0]
		if !token.IsAlpha(c) {
			break
		}
		ident.WriteByte(c)
		l.reader.Scan()
	}
	return ident.String()
}

func (l *Lexer) buildInt() string {
	var num strings.Builder

	for {
    ch := l.reader.Bytes()[0]
		if !token.IsDigit(ch) {
			break
		}
		num.WriteByte(ch)
		l.reader.Scan()
	}

	return num.String()
}

