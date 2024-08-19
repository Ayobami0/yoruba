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
	return &stmt
}

// Parser for return statement
func (p *Parser) parseReturnStatement() ast.Statement {
	stmt := ast.ReturnStatement{PrefixToken: p.curToken}

	// TODO: Shiuld expect an expression or identifier

	if !p.expectPeek(token.RTN_PREFIX) {
    return nil
	}

	return &stmt
}
