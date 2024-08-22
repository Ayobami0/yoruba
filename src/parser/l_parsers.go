package parser

import (
	"fmt"
	"strconv"

	"github.com/Ayobami0/yoruba/src/ast"
)

// Parser for string expression
func (p *Parser) parseStringLiteral() ast.Expression {
	exp := &ast.StringLiteral{Token: p.curToken, Value: p.curToken.Literal}

	return exp
}

func (p *Parser) parseNumberLiteral() ast.Expression {
	exp := &ast.NumberLiteral{Token: p.curToken}
	value, err := strconv.ParseInt(p.curToken.Literal, 0, 64)

	if err != nil {
		msg := fmt.Sprintf("could not parse %q as number", p.curToken.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}

	exp.Value = value

	return exp
}
