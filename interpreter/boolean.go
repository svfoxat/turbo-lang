package interpreter

import (
	"turbo/object"
)

func evalBangOperatorExpression(right object.Object) object.Object {
	switch right {
	case TRUE:
		return FALSE
	case FALSE:
		return TRUE
	case NULL:
		return FALSE
	}

	switch right.Type() {
	case object.INTEGER_OBJ:
		if right.(*object.Integer).Value != 0 {
			return FALSE
		} else {
			return TRUE
		}
	case object.FLOAT_OBJ:
		if right.(*object.FloatLiteral).Value != 0.0 {
			return FALSE
		} else {
			return TRUE
		}
	}

	return FALSE
}

func getNativeBoolean(input bool) *object.Boolean {
	if input {
		return TRUE
	} else {
		return FALSE
	}
}

func isTruthy(input object.Object) bool {
	switch input {
	case NULL:
		return false
	case TRUE:
		return true
	case FALSE:
		return false
	case input.(*object.Integer):
		if input.(*object.Integer).Value != 0 {
			return true
		} else {
			return false
		}
	}
	return false
}
