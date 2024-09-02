package object

import "strings"

type Environment struct {
	store map[string]Object
	outer *Environment
}

// Retrives a reference from the environment
func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]

  // Checks if the the reference is defined in the current scope
	if !ok && e.outer != nil {
		obj, ok = e.outer.Get(name)
	}

	return obj, ok
}

// Sets a reference in the environment
func (e *Environment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}

func (e Environment) Debug() string {
	var debug strings.Builder
	debug.WriteString("==========START==========\n")
	for k, v := range e.store {
		debug.WriteString(k)
		debug.WriteString(": ")
		if v != nil {
			debug.WriteString(v.Inspect())
		} else {
			debug.WriteString("nil")
		}
		debug.WriteString("\n")
	}
	debug.WriteString("==========END==========")

	return debug.String()
}

// New environment
func NewEnvironment() *Environment {
	s := make(map[string]Object)
	return &Environment{store: s, outer: nil}
}

// Creates a inner scope. The outer environment is registered as the
// enclosing environment or outer scope
func NewEnclosedEnvironment(outer *Environment) *Environment {
	env := NewEnvironment()
	env.outer = outer

	return env
}
