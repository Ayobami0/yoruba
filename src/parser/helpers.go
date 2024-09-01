package parser

import (
	"github.com/Ayobami0/yoruba/src/ast"
	"github.com/Ayobami0/yoruba/src/token"
)

// Type definations for infix and prefix handlers
type (
	prefixParseFn func() ast.Expression
	infixParseFn  func(ast.Expression) ast.Expression
)

// Binding Power / Precendence
const (
	_ int = iota
	LOWEST
	EQUALS
	LESSGREATER
	SUM
	PRODUCT
	PREFIX
	CALL
)

// Map of token types to their precedence
var precedences = map[token.TokenType]int{
	token.EQL:    EQUALS,
  token.AND: EQUALS,
  token.OR: EQUALS,
	token.NOTEQL: EQUALS,
	token.L_THAN: LESSGREATER,
	token.G_THAN: LESSGREATER,
	token.PLUS:   SUM,
	token.MINUS:  SUM,
	token.DIVIDE: PRODUCT,
	token.TIMES:  PRODUCT,
}

// Returns the precedence of the token ahead of the current token
// If precedence doesn't exist, return LOWEST
func (p *Parser) peekPrecedence() int {
	if p, ok := precedences[p.peekToken.Type]; ok {
		return p
	}
	return LOWEST
}

// Returns the precedence of the current token
// If precedence doesn't exist, return LOWEST
func (p *Parser) curPrecedence() int {
	if p, ok := precedences[p.curToken.Type]; ok {
		return p
	}
	return LOWEST
}

// Parses every statement passed to the program
func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	case token.RTN_PREFIX:
		return p.parseReturnStatement()
	case token.IF:
		return p.parseIfStatement()
	case token.FUNCTION:
		return p.parseFunctionStatement()
	case token.TILL:
		return p.parseLoopStatement()
	default:
		return p.parseExpressionStatement()
	}
}

// Checks if next token ahead is of type t
func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

// Checks if current token is of type t
func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

// Confirms if the expected of type t and advances the curToken pointer
func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		p.peekError(t)
		return false
	}
}

// Advances current token by one token on each successive calls.
func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

// Registers a new prefix handler function
func (p *Parser) registerPrefix(tokenType token.TokenType, fn prefixParseFn) {
	p.prefixParseFns[tokenType] = fn
}

// Registers a new infix handler function
func (p *Parser) registerInfix(tokenType token.TokenType, fn infixParseFn) {
	p.infixParseFns[tokenType] = fn
}
