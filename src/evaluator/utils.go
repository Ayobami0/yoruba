package evaluator

import (
	"fmt"

	"github.com/Ayobami0/yoruba/src/object"
)

var (
	TRUE  = &object.Boolean{Value: true}
	FALSE = &object.Boolean{Value: false}
)

func nativeBoolToBooleanObject(input bool) *object.Boolean {
	if input {
		return TRUE
	}
	return FALSE
}

func evalMinusPrefixOperatorExpression(right object.Object) object.Object {
	if right.Type() != object.NUMBER_OBJ {
		return newError(PREFIX_OPERATOR_UNKNOWN, "-", right.Type())
	}
	value := right.(*object.Number).Value
	return &object.Number{Value: -value}
}

func evalIntegerInfixExpression(operator string, left, right object.Object) object.Object {
	leftVal := left.(*object.Number).Value
	rightVal := right.(*object.Number).Value
	switch operator {
	case "+":
		return &object.Number{Value: leftVal + rightVal}
	case "-":
		return &object.Number{Value: leftVal - rightVal}
	case "*":
		return &object.Number{Value: leftVal * rightVal}
	case "/":
		return &object.Number{Value: leftVal / rightVal}
	case "kereju":
		return nativeBoolToBooleanObject(leftVal < rightVal)
	case "tobiju":
		return nativeBoolToBooleanObject(leftVal > rightVal)
	case "baje":
		return nativeBoolToBooleanObject(leftVal == rightVal)
	case "kobaje":
		return nativeBoolToBooleanObject(leftVal != rightVal)
	default:
		return newError(INFIX_OPERATOR_UNKNOWN, left.Type(), operator, right.Type())
	}
}

// Checks if an object is a truthy value or not
func isTruthy(obj object.Object) bool {
	switch obj {
	case nil:
		return false
	case TRUE:
		return true
	case FALSE:
		return false
	default:
		return true
	}
}

// Does the actual function call and unwraps the return value.
func applyFunction(fn object.Object, args []object.Object) object.Object {
	switch fn := fn.(type) {
	case *object.Function:
		extendedEnv := extendFunctionEnv(fn, args)
		evaluated := Eval(fn.Body, extendedEnv)
		return unwrapReturnValue(evaluated)
	case *object.Builtin:
		return fn.Fn(args...)
	default:
		return newError(IDENT_NOT_FUNCTION, fn.Type())
	}
}

// Creates a new [object.Environment] and encloses it with the current [object.Environment] to
// allow access to local function bindings.
//
// It then binds the argument of the function to the new environment.
func extendFunctionEnv(fn *object.Function, args []object.Object) *object.Environment {
	env := object.NewEnclosedEnvironment(fn.Env)

	for paramIdx, param := range fn.Parameters {
		env.Set(param.Value, args[paramIdx])
	}

	return env
}

// Unwraps return value and stop it from "bubbling up" past several functions
// in cases of nested functions. Only the current function context is returned.
func unwrapReturnValue(obj object.Object) object.Object {
	if returnValue, ok := obj.(*object.ReturnValue); ok {
		return returnValue.Value
	}

	return obj
}

func newError(f string, args ...interface{}) *object.Error {
	return &object.Error{Message: fmt.Sprintf(f, args...)}
}

func isError(obj object.Object) bool {
	if obj != nil {
		return obj.Type() == object.ERROR_OBJ
	}
	return false
}
