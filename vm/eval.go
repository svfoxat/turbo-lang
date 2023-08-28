package vm

import (
	"interpreter/ast"
)

// global constants
var (
	NULL  = &Null{}
	TRUE  = &Boolean{Value: true}
	FALSE = &Boolean{Value: false}
)

func Eval(node ast.Node) Object {
	switch node := node.(type) {

	// Statements
	case *ast.Program:
		return evalStatements(node.Statements)
	case *ast.ExpressionStatement:
		return Eval(node.Expression)

	// Expressions
	case *ast.IntegerLiteral:
		return &Integer{Value: node.Value}
	case *ast.BooleanLiteral:
		return getNativeBoolean(node.Value)
	case *ast.PrefixExpression:
		right := Eval(node.Right)
		return evalPrefixExpression(node.Operator, right)
	}
	return nil
}

func evalStatements(stmts []ast.Statement) Object {
	var result Object

	for _, stmt := range stmts {
		result = Eval(stmt)
	}
	return result
}

func evalPrefixExpression(operator string, right Object) Object {
	switch operator {
	case "!":
		return evalBangOperatorExpression(right)
	default:
		return NULL
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
