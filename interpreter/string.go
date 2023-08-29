package interpreter

import (
	"turbo/object"
)

func evalStringInfixExpression(
	operator string,
	left, right object.Object,
) object.Object {
	if operator != "+" {
		return newError("unknown operator: %s %s %s",
			left.Type(), operator, right.Type())
	}

	leftVal := left.(*object.StringLiteral).Value
	rightVal := right.(*object.StringLiteral).Value
	return &object.StringLiteral{Value: leftVal + rightVal}
}
