package vm

import (
	"fmt"
	"turbo/ast"
)

// global constants
var (
	NULL  = &Null{}
	TRUE  = &Boolean{Value: true}
	FALSE = &Boolean{Value: false}
)

func Eval(node ast.Node, env *Environment) Object {
	switch node := node.(type) {

	// Statements
	case *ast.Program:
		return evalProgram(node, env)
	case *ast.ExpressionStatement:
		return Eval(node.Expression, env)

	// Expressions
	case *ast.IntegerLiteral:
		return &Integer{Value: node.Value}
	case *ast.BooleanLiteral:
		return getNativeBoolean(node.Value)
	case *ast.PrefixExpression:
		right := Eval(node.Right, env)
		if isError(right) {
			return right
		}
		return evalPrefixExpression(node.Operator, right)
	case *ast.InfixExpression:
		left := Eval(node.Left, env)
		right := Eval(node.Right, env)
		if isError(left) {
			return left
		}
		if isError(right) {
			return right
		}

		return evalInfixExpression(node.Operator, left, right)

	case *ast.BlockStatement:
		return evalBlockStatement(node.Statements, env)
	case *ast.IfExpression:
		return evalIfExpression(node, env)

	case *ast.ReturnStatement:
		val := Eval(node.ReturnValue, env)
		if isError(val) {
			return val
		}
		return &ReturnValue{Value: val}

	case *ast.LetStatement:
		val := Eval(node.Value, env)
		if isError(val) {
			return val
		}
		env.Set(node.Name.Value, val)

	case *ast.Identifier:
		return evalIdentifier(node, env)
	}
	return nil
}

func evalIdentifier(node *ast.Identifier, env *Environment) Object {
	if val, ok := env.Get(node.Value); ok {
		return val
	}
	return newError("identifier not found: " + node.Value)
}

func evalProgram(program *ast.Program, env *Environment) Object {
	var result Object

	for _, stmt := range program.Statements {
		result = Eval(stmt, env)
		if ret, ok := result.(*ReturnValue); ok {
			return ret.Value
		}

		switch result.(type) {
		case *ReturnValue:
			return result.(*ReturnValue).Value
		case *Error:
			return result
		}
	}
	return result
}

func evalBlockStatement(stmts []ast.Statement, env *Environment) Object {
	var result Object

	for _, stmt := range stmts {
		result = Eval(stmt, env)

		if result != nil &&
			(result.Type() == RETURN_VALUE_OBJ || result.Type() == ERROR_OBJ) {
			return result
		}
	}
	return result
}

func evalIfExpression(ie *ast.IfExpression, env *Environment) Object {
	condition := Eval(ie.Condition, env)
	if isError(condition) {
		return condition
	}

	if isTruthy(condition) {
		return Eval(ie.Consequence, env)
	} else if ie.Alternative != nil {
		return Eval(ie.Alternative, env)
	} else {
		return NULL
	}
}

func evalPrefixExpression(operator string, right Object) Object {
	switch operator {
	case "!":
		return evalBangOperatorExpression(right)
	case "-":
		if right.Type() != INTEGER_OBJ {
			return newError("unknown operator: -%s", right.Type())
		}
		value := right.(*Integer).Value
		return &Integer{Value: -value}
	default:
		return newError("unknown operator: %s%s", operator, right.Type())
	}
}

func evalInfixExpression(operator string, left, right Object) Object {
	// slog.Info("[INFIX]", "OPERATOR", operator, "LEFT", left.Type(), "RIGHT", right.Type())

	switch {
	case left.Type() == INTEGER_OBJ && right.Type() == INTEGER_OBJ:
		return evalIntegerInfixExpression(operator, left, right)
	case operator == "==":
		return getNativeBoolean(left == right)
	case operator == "!=":
		return getNativeBoolean(left != right)
	case left.Type() != right.Type():
		return newError("type mismatch: %s %s %s", left.Type(), operator, right.Type())
	default:
		return newError("unknown operator: %s %s %s", left.Type(), operator, right.Type())
	}
}

func evalIntegerInfixExpression(operator string, left, right Object) Object {
	leftValue := left.(*Integer).Value
	rightValue := right.(*Integer).Value

	switch operator {
	case "+":
		return &Integer{Value: leftValue + rightValue}
	case "-":
		return &Integer{Value: leftValue - rightValue}
	case "*":
		return &Integer{Value: leftValue * rightValue}
	case "/":
		return &Integer{Value: leftValue / rightValue}
	case "<":
		return getNativeBoolean(leftValue < rightValue)
	case ">":
		return getNativeBoolean(leftValue > rightValue)
	case "==":
		return getNativeBoolean(leftValue == rightValue)
	case "!=":
		return getNativeBoolean(leftValue != rightValue)
	default:
		return newError("unknown operator: %s %s %s", left.Type(), operator, right.Type())
	}
}

// what should be returned when banged?
// !5 -> false
// !0 -> true
func evalBangOperatorExpression(right Object) Object {
	switch right {
	case TRUE:
		return FALSE
	case FALSE:
		return TRUE
	case NULL:
		return FALSE
	case right.(*Integer):
		if right.(*Integer).Value != 0 {
			return FALSE
		} else {
			return TRUE
		}
	}
	return FALSE
}

func getNativeBoolean(input bool) *Boolean {
	if input {
		return TRUE
	} else {
		return FALSE
	}
}

func isTruthy(input Object) bool {
	switch input {
	case NULL:
		return false
	case TRUE:
		return true
	case FALSE:
		return false
	case input.(*Integer):
		if input.(*Integer).Value != 0 {
			return true
		} else {
			return false
		}
	}
	return false
}

func isError(input Object) bool {
	if input != nil {
		return input.Type() == ERROR_OBJ
	}
	return false
}

func newError(format string, a ...interface{}) *Error {
	return &Error{Message: fmt.Sprintf(format, a...)}
}
