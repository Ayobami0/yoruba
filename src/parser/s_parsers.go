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
