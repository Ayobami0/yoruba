package parser

import (
	"fmt"

	"github.com/Ayobami0/yoruba/src/token"
)

func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("O ti ṣe yẹ ami atẹle lati jẹ %s ṣugbọn o jẹ %s", t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}
