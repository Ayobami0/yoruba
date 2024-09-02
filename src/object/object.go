package object

import (
	"fmt"
	"strings"

	"github.com/Ayobami0/yoruba/src/ast"
)

type ObjectType string

type Object interface {
	Type() ObjectType
	Inspect() string
}

const (
	RETURN_VALUE_OBJ = "RETURN_VALUE"
	BREAK_OBJ        = "BREAK"
	CONTINUE_OBJ     = "CONTINUE"
	NUMBER_OBJ       = "NUMBER"
	BOOLEAN_OBJ      = "BOOLEAN"
	STRING_OBJ       = "STRING"
	FUNCTION_OBJ     = "FUNCTION"
)

type Number struct {
	Value int64
}

func (n *Number) Inspect() string { return fmt.Sprintf("%d", n.Value) }
func (n Number) Type() ObjectType { return NUMBER_OBJ }

type Boolean struct {
	Value bool
}

func (b *Boolean) Type() ObjectType { return BOOLEAN_OBJ }
func (b *Boolean) Inspect() string  { return fmt.Sprintf("%t", b.Value) }

type String struct {
	Value string
}

func (s *String) Inspect() string { return s.Value }
func (n String) Type() ObjectType { return STRING_OBJ }

type ReturnValue struct {
	Value Object
}

func (rv *ReturnValue) Type() ObjectType { return RETURN_VALUE_OBJ }
func (rv *ReturnValue) Inspect() string  { return rv.Value.Inspect() }

type Function struct {
	Parameters []*ast.Identifier
	Body       *ast.BlockStatement
	Env        *Environment
}

func (fn *Function) Inspect() string {
	var s strings.Builder

	for i, p := range fn.Parameters {
		s.WriteString(p.Value)

		if i != len(fn.Parameters)-1 {
			s.WriteString(",")
		}
	}
	return fmt.Sprintf("ise<%s>", s.String())
}
func (fn Function) Type() ObjectType { return FUNCTION_OBJ }

type Break struct{}

func (bk *Break) Type() ObjectType { return BREAK_OBJ }
func (bk *Break) Inspect() string  { return "nil" }

type Continue struct{}

func (c *Continue) Type() ObjectType { return BREAK_OBJ }
func (c *Continue) Inspect() string  { return "nil" }
