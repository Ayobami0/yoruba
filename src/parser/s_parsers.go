// Statement parsers
package parser

import (
	"github.com/Ayobami0/yoruba/src/ast"
	"github.com/Ayobami0/yoruba/src/token"
)

// Parser for let statement
func (p *Parser) parseLetStatement() ast.Statement {
	stmt := ast.LetStatement{Token: p.curToken}

	if !p.expectPeek(token.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if !p.expectPeek(token.ASSIGNMENT) {
		return nil
	}

	p.nextToken()

  stmt.Value = p.parseExpression(LOWEST)
	return &stmt
}

// Parser for return statement
func (p *Parser) parseReturnStatement() ast.Statement {
	stmt := ast.ReturnStatement{PrefixToken: p.curToken}

  p.nextToken()

	stmt.ReturnValue = p.parseExpression(LOWEST)

	if !p.expectPeek(token.RTN_SURFIX) {
		return nil
	}

	stmt.SurfixToken = p.curToken

	return &stmt
}

func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	stmt := &ast.ExpressionStatement{Token: p.curToken}
	stmt.Expression = p.parseExpression(LOWEST)

	return stmt
}

func (p *Parser) parseIfStatement() *ast.IfStatement {
	stmt := &ast.IfStatement{Token: p.curToken}
	p.nextToken()

	stmt.Condition = p.parseExpression(LOWEST)

	if !p.expectPeek(token.THEN) {
		return nil
	}

	stmt.Consequence = p.parseBlockStatement()

	switch p.curToken.Type {
	case token.ELSE:
		stmt.Alternative = p.parseBlockStatement()
	}

	if !p.curTokenIs(token.END) {
		return nil
	}

	return stmt
}

func (p *Parser) parseFunctionStatement() *ast.FunctionStatement {
	fn := &ast.FunctionStatement{Token: p.curToken}

	if !p.expectPeek(token.IDENT) {
		return nil
	}

	fn.Parameters = p.parseFunctionParameters()
	fn.Body = p.parseBlockStatement()

	if !p.curTokenIs(token.END) {
		return nil
	}

	return fn
}

func (p *Parser) parseFunctionParameters() []*ast.Identifier {

	var params []*ast.Identifier

	if p.peekTokenIs(token.EXECUTE) {
		p.nextToken()
		return params
	}

	p.nextToken()
	ident := ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
	params = append(params, &ident)

	for p.peekTokenIs(token.COMMA) {
		p.nextToken()
		p.nextToken()
		ident := ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

		params = append(params, &ident)
	}

	if !p.expectPeek(token.EXECUTE) {
		return nil
	}

	return params
}

func (p *Parser) parseBlockStatement() *ast.BlockStatement {
	blk := &ast.BlockStatement{Token: p.curToken}
	p.nextToken()

	for !p.curTokenIs(token.END) && !p.curTokenIs(token.EOF) && !p.curTokenIs(token.ELSE) {
		stmt := p.parseStatement() // Statements can also include if statements

		if stmt != nil {
			blk.Statements = append(blk.Statements, stmt)
		}
		p.nextToken()
	}
	return blk
}

func (p *Parser) parseLoopStatement() *ast.LoopStatement {
	loop := &ast.LoopStatement{Token: p.curToken}

	p.nextToken()

	loop.Condition = p.parseExpression(LOWEST)

	if !p.expectPeek(token.EXECUTE) {
		return nil
	}

	loop.Body = p.parseBlockStatement()

	if !p.curTokenIs(token.END) {
		return nil
	}

	return loop
}
