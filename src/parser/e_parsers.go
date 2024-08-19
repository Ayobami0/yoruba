package parser

import (
	"github.com/Ayobami0/yoruba/src/ast"
)

// Parser for string expression
func (p *Parser) parseStringLiteral() ast.Expression {
  exp := &ast.StringLiteral{Token: p.curToken, Value: p.curToken.Literal}

  return exp
}
