package ast

import "github.com/Ayobami0/yoruba/src/token"

type Identifier struct {
  Token token.Token
  Value string
}
func (i *Identifier) expressionNode()  {
}

func (i Identifier) TokenLiteral() string {
  return i.Token.Literal
}
