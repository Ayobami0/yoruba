package parser

import (
	"fmt"

	"github.com/Ayobami0/yoruba/src/token"
)

func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead",
		t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}
