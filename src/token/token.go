package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

func New(t TokenType, l byte) Token {
  return Token{Type: t, Literal: string(l)}
}
