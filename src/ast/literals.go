package ast

import (
	"github.com/Ayobami0/yoruba/src/token"
)

type StringLiteral struct {
	Token token.Token
	Value string
}
func (s *StringLiteral) expressionNode() {}
func (s *StringLiteral) TokenLiteral() string { return s.Token.Literal }

type NumberLiteral struct {
	Token token.Token
	Value int64
}
func (n *NumberLiteral) expressionNode() {}
func (n *NumberLiteral) TokenLiteral() string { return n.Token.Literal }

type BooleanLiteral struct {
	Token token.Token
	Value bool
}
func (b *BooleanLiteral) expressionNode()      {}
func (b *BooleanLiteral) TokenLiteral() string { return b.Token.Literal }
