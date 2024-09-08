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

type BuiltinFunction func(args ...Object) Object

const (
	RETURN_VALUE_OBJ = "RETURN_VALUE"
	BREAK_OBJ        = "BREAK"
	CONTINUE_OBJ     = "CONTINUE"
	NUMBER_OBJ       = "NUMBER"
	BOOLEAN_OBJ      = "BOOLEAN"
	STRING_OBJ       = "STRING"
	FUNCTION_OBJ     = "FUNCTION"
	BUILTIN_OBJ      = "BUILTIN"
	ERROR_OBJ        = "ERROR"
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

type Builtin struct {
	Fn BuiltinFunction
}

func (b *Builtin) Type() ObjectType { return BUILTIN_OBJ }
func (b *Builtin) Inspect() string  { return "builtin function" }

type Error struct {
	Message string
	ErrType string
}

func (e *Error) Type() ObjectType { return ERROR_OBJ }
func (e *Error) Inspect() string  { return "aṣiṣe: " + e.Message }
