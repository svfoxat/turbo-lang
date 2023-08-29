package interpreter

import (
	"turbo/object"
)

func evalNumberPrefixExpression(operator string, right object.Object) object.Object {
	switch right.Type() {
	case object.INTEGER_OBJ:
		return evalIntegerPrefixExpression(operator, right.(*object.Integer))
	case object.FLOAT_OBJ:
		return evalFloatPrefixExpression(operator, right.(*object.FloatLiteral))
	default:
		return newError("unknown operator: %s%s", operator, right.Type())
	}
}

func evalNumberInfixExpression(operator string, left, right object.Object) object.Object {
	switch {
	case left.Type() == object.INTEGER_OBJ && right.Type() == object.INTEGER_OBJ:
		return evalIntegerInfixExpression(operator, left, right)
	case left.Type() == object.FLOAT_OBJ && right.Type() == object.FLOAT_OBJ:
		return evalFloatInfixExpression(operator, left, right)
	default:
		return newError("type mismatch: %s %s %s", left.Type(), operator, right.Type())
	}
}

func evalIntegerPrefixExpression(operator string, right *object.Integer) object.Object {
	switch operator {
	case "!":
		return getNativeBoolean(right.Value == 0)
	case "-":
		value := right.Value
		return &object.Integer{Value: -value}
	default:
		return newError("unknown operator: %s%s", operator, right.Type())
	}
}

func evalFloatPrefixExpression(operator string, right *object.FloatLiteral) object.Object {
	switch operator {
	case "!":
		return getNativeBoolean(right.Value == 0.0)
	case "-":
		value := right.Value
		return &object.FloatLiteral{Value: -value}
	default:
		return newError("unknown operator: %s%s", operator, right.Type())
	}
}

func evalIntegerInfixExpression(operator string, left, right object.Object) object.Object {
	leftValue := left.(*object.Integer).Value
	rightValue := right.(*object.Integer).Value

	switch operator {
	case "+":
		return &object.Integer{Value: leftValue + rightValue}
	case "-":
		return &object.Integer{Value: leftValue - rightValue}
	case "*":
		return &object.Integer{Value: leftValue * rightValue}
	case "/":
		return &object.Integer{Value: leftValue / rightValue}
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

func evalFloatInfixExpression(operator string, left, right object.Object) object.Object {
	leftValue := left.(*object.FloatLiteral).Value
	rightValue := right.(*object.FloatLiteral).Value

	switch operator {
	case "+":
		return &object.FloatLiteral{Value: leftValue + rightValue}
	case "-":
		return &object.FloatLiteral{Value: leftValue - rightValue}
	case "*":
		return &object.FloatLiteral{Value: leftValue * rightValue}
	case "/":
		return &object.FloatLiteral{Value: leftValue / rightValue}
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
