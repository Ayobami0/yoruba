package ast


import "github.com/Ayobami0/yoruba/src/token"

type StringLiteral struct {
  Token token.Token
  Value string
}

func (s *StringLiteral) expressionNode() {
}

func (s *StringLiteral) TokenLiteral() string {
  return s.Token.Literal
}
