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

	// TODO: Work on expression
	p.nextToken()
	return &stmt
}

// Parser for return statement
func (p *Parser) parseReturnStatement() ast.Statement {
	stmt := ast.ReturnStatement{PrefixToken: p.curToken}

	// TODO: Should expect an expression or identifier. Advance token for now
	for p.curToken.Type != token.RTN_SURFIX {
		p.nextToken()
	}

	// if !p.expectPeek(token.RTN_PREFIX) {
	//    return nil
	// }

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

func (p *Parser) parseBlockStatement() *ast.BlockStatement {
	blk := &ast.BlockStatement{Token: p.curToken}
	p.nextToken()

	for !p.curTokenIs(token.END) && !p.curTokenIs(token.EOF) && !p.curTokenIs(token.ELSE) {
		stmt := p.parseStatement()

		if stmt != nil {
			blk.Statements = append(blk.Statements, stmt)
		}
		p.nextToken()
	}
	return blk
}
