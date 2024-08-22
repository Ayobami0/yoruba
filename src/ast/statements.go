package ast

import "github.com/Ayobami0/yoruba/src/token"

type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (l *LetStatement) statementNode() {
}

func (l *LetStatement) TokenLiteral() string {
	return l.Token.Literal
}

type ReturnStatement struct {
	PrefixToken token.Token
	SurfixToken token.Token

	ReturnValue *Identifier
}

func (r *ReturnStatement) statementNode() {
}

func (r *ReturnStatement) TokenLiteral() string {
	return r.PrefixToken.Literal + " " + r.SurfixToken.Literal
}

type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

func (e *ExpressionStatement) statementNode()      {}
func (e *ExpressionStatement) TokenLiteral() string { return e.Token.Literal }
