// Expression Parsers
package parser

import (
	"fmt"

	"github.com/Ayobami0/yoruba/src/ast"
	"github.com/Ayobami0/yoruba/src/token"
)

func (p *Parser) noPrefixParseFnError(t token.TokenType) {
	msg := fmt.Sprintf("no prefix parse function for %s found", t)
	p.errors = append(p.errors, msg)
}

func (p *Parser) parseExpression(precedence int) ast.Expression {
	prefix := p.prefixParseFns[p.curToken.Type]
	if prefix == nil {
		p.noPrefixParseFnError(p.curToken.Type)
		return nil
	}
	leftExp := prefix()

	for precedence < p.peekPrecedence() {
		infix := p.infixParseFns[p.peekToken.Type]
		if infix == nil {
			return leftExp
		}
		p.nextToken()
		leftExp = infix(leftExp) // Recussive step
	}

	return leftExp
}

// NuDs - Null Denomination Function
func (p *Parser) parsePrefixExpression() ast.Expression {
	expression := &ast.PrefixExpression{
		Token:    p.curToken,
		Operator: p.curToken.Literal,
	}
	p.nextToken()
	expression.Right = p.parseExpression(PREFIX)
	return expression
}

// LeDs - Left Denomination Function
func (p *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
	expression := &ast.InfixExpression{
		Token:    p.curToken,
		Operator: p.curToken.Literal,
		Left:     left, // Assigns the left expression to the left field of the infix expression
	}
	precedence := p.curPrecedence()
	p.nextToken()
	expression.Right = p.parseExpression(precedence) // Parses the RHS expression
	return expression
}

// Parser for grouped expressions i.e (1 + 2) * 3
func (p *Parser) parseGroupedExpression() ast.Expression {
	p.nextToken() // Move to next token (now in parentises)

	expr := p.parseExpression(LOWEST) // Parse expression inside parenteses

	if !p.expectPeek(token.R_GROUP) {
		return nil
	}

	return expr
}

func (p *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
}

func (p *Parser) parseCallExpression() ast.Expression {
	exp := &ast.CallExpression{Token: p.curToken}

	p.nextToken()
	exp.Function = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if !p.peekTokenIs(token.F_CALL_SURFIX) {
		if !p.expectPeek(token.F_CALL_INFIX) {
			return nil
		}
		exp.Arguments = p.parseCallArguments()
	}

	if !p.expectPeek(token.F_CALL_SURFIX){
		return nil
	}

	return exp
}

func (p *Parser) parseCallArguments() []ast.Expression {
	args := []ast.Expression{}

	p.nextToken()
	args = append(args, p.parseExpression(LOWEST))

	for p.peekTokenIs(token.COMMA) {
		p.nextToken()
		p.nextToken()

		args = append(args, p.parseExpression(LOWEST))
	}

	return args
}
