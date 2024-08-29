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

func (e *ExpressionStatement) statementNode()       {}
func (e *ExpressionStatement) TokenLiteral() string { return e.Token.Literal }

type IfStatement struct {
	Token       token.Token
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

func (i *IfStatement) statementNode()       {}
func (i *IfStatement) TokenLiteral() string { return i.Token.Literal }

type FunctionStatement struct {
	Token      token.Token
	Ident      Identifier
	Parameters []*Identifier
	Body       *BlockStatement
}

func (fn *FunctionStatement) statementNode()       {}
func (fn *FunctionStatement) TokenLiteral() string { return fn.Token.Literal }

type LoopStatement struct {
	Token     token.Token
	Condition Expression
	Body      *BlockStatement
}

func (lp *LoopStatement) statementNode()       {}
func (lp *LoopStatement) TokenLiteral() string { return lp.Token.Literal }

type BlockStatement struct {
	Token      token.Token
	Statements []Statement
}

func (b *BlockStatement) statementNode()       {}
func (b *BlockStatement) TokenLiteral() string { return b.Token.Literal }
