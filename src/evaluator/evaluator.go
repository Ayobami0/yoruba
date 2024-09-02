package evaluator

import (
	"github.com/Ayobami0/yoruba/src/ast"
	"github.com/Ayobami0/yoruba/src/object"
)

func Eval(node ast.Node, env *object.Environment) object.Object {
	switch node := node.(type) {
	case *ast.Program:
		return evalProgram(node, env)
	case *ast.LetStatement:
		val := Eval(node.Value, env)
		env.Set(node.Name.Value, val)
	case *ast.BreakStatement:
		return &object.Break{}
		// case *ast.Continue: // TODO: continue
	// 	return &object.Continue{}
	case *ast.Identifier:
		return evalIdentifier(node, env)
	case *ast.FunctionStatement:
		return evalFunctionStatement(node, env)
	case *ast.CallExpression:
		function := Eval(node.Function, env)
		args := evalExpressions(node.Arguments, env)
		// TODO: Error Handling
		// if len(args) == 1 { // && isError(args[0]) {
		// 	return args[0]
		// }
		return applyFunction(function, args)
	case *ast.BlockStatement:
		return evalBlockStatement(node, env)
	case *ast.LoopStatement:
		evalLoopStatement(node, env)
	case *ast.PrefixExpression:
		right := Eval(node.Right, env)
		return evalPrefixExpression(node.Operator, right)
	case *ast.IfStatement:
		return evalIfStatement(node, env)
	case *ast.InfixExpression:
		left := Eval(node.Left, env)
		right := Eval(node.Right, env)
		return evalInfixExpression(node.Operator, left, right)
	case *ast.ReturnStatement:
		rtn := Eval(node.ReturnValue, env)
		return &object.ReturnValue{Value: rtn}
	case *ast.ExpressionStatement:
		return Eval(node.Expression, env)
	case *ast.NumberLiteral:
		return &object.Number{Value: node.Value}
	case *ast.StringLiteral:
		return &object.String{Value: node.Value}
	case *ast.BooleanLiteral:
		return nativeBoolToBooleanObject(node.Value)
	}
	return nil
}

func evalLoopStatement(node *ast.LoopStatement, env *object.Environment) object.Object {
	cond := Eval(node.Condition, env)
	condV, ok := cond.(*object.Boolean)

	if !ok {
		return nil
	} // TODO: Error Handling

	var val object.Object

	for !condV.Value {
		val = Eval(node.Body, env)

		if val != nil && val.Type() == object.BREAK_OBJ {
			break
		}
		// Re-evaluate the condition again
		cond = Eval(node.Condition, env)
		condV = cond.(*object.Boolean)
	}

	return val
}

func evalFunctionStatement(node *ast.FunctionStatement, env *object.Environment) object.Object {
	fn := &object.Function{Parameters: node.Parameters, Body: node.Body, Env: env}

	env.Set(node.Ident.Value, fn)
	return fn
}

func evalExpressions(exps []ast.Expression, env *object.Environment) []object.Object {
	var result []object.Object

	for _, e := range exps {
		evaluated := Eval(e, env)
		// TODO: Error Handling
		result = append(result, evaluated)
	}

	return result
}

func evalIdentifier(node *ast.Identifier, env *object.Environment) object.Object {
	val, ok := env.Get(node.Value)

	if !ok {
		return nil
		// TODO: ERROR HANDLING
		// return newError("identifier not found: " + node.Value)
	}

	return val
}

func evalInfixExpression(s string, left, right object.Object) object.Object {

	switch {
	case left.Type() == object.NUMBER_OBJ && right.Type() == object.NUMBER_OBJ:
		return evalIntegerInfixExpression(s, left, right)
	case s == "baje":
		return nativeBoolToBooleanObject(left == right)
	case s == "kobaje":
		return nativeBoolToBooleanObject(left != right)
	case left.Type() == object.BOOLEAN_OBJ && right.Type() == object.BOOLEAN_OBJ:
		l := left.(*object.Boolean)
		r := right.(*object.Boolean)
		if s == "ati" {
			return nativeBoolToBooleanObject(l.Value && r.Value)
		} else if s == "tabi" {
			return nativeBoolToBooleanObject(l.Value && r.Value)
		} else {
			return nil
		}
	default:
		return nil
	}
}

func evalStatements(statements []ast.Statement, env *object.Environment) object.Object {
	var res object.Object

	for _, stmt := range statements {
		res = Eval(stmt, env)

		if returnValue, ok := res.(*object.ReturnValue); ok {
			return returnValue.Value
		}
	}

	return res
}

func evalPrefixExpression(operator string, right object.Object) object.Object {
	switch operator {
	case "-":
		return evalMinusPrefixOperatorExpression(right)
	default:
		// TODO: ERROR HANDLING
		return nil
	}
}

func evalIfStatement(stmt *ast.IfStatement, env *object.Environment) object.Object {
	cond := Eval(stmt.Condition, env)

	if isTruthy(cond) {
		return Eval(stmt.Consequence, env)
	} else if stmt.Alternative != nil {
		return Eval(stmt.Alternative, env)
	}
	return nil
}

func evalProgram(program *ast.Program, env *object.Environment) object.Object {
	var result object.Object

	for _, statement := range program.Statements {
		result = Eval(statement, env)
		if returnValue, ok := result.(*object.ReturnValue); ok {
			return returnValue.Value
		}
	}

	return result
}

func evalBlockStatement(block *ast.BlockStatement, env *object.Environment) object.Object {
	var result object.Object
	for _, statement := range block.Statements {
		result = Eval(statement, env)
		if result != nil && result.Type() == object.RETURN_VALUE_OBJ {
			return result
		} else if result != nil && result.Type() == object.BREAK_OBJ {
			return result
		}
	}
	return result
}
