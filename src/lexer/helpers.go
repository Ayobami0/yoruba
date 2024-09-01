package lexer

import (
	"errors"
	"strings"

	"github.com/Ayobami0/yoruba/src/token"
)

func (l *Lexer) buildComment() error {
	var err error

	l.reader.Scan()
	c := l.reader.Bytes()[0]
	if c != '[' {
		err = errors.New("")
	}
	l.reader.Scan()
	for {
		c = l.reader.Bytes()[0]
    if c == '\\' {
      l.reader.Scan()
      l.reader.Scan()
    }
		if c == ']' {
			l.reader.Scan()
			if c != ']' {
				err = errors.New("")
			}
			break
		}
		l.reader.Scan()
	}
	return err
}

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
		b := l.reader.Bytes()
		if len(b) == 0 {
			break
		}
		c := b[0]
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
		b := l.reader.Bytes()
		if len(b) == 0 {
			break
		}
		ch := b[0]
		if !token.IsDigit(ch) {
			break
		}
		num.WriteByte(ch)
		l.reader.Scan()
	}

	return num.String()
}
