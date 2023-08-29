package interpreter

import (
	"turbo/object"
)

var builtins = map[string]*object.Builtin{
	"len":   {Fn: lenFunction},
	"print": {Fn: printFunction},
}

func lenFunction(args ...object.Object) object.Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. got=%d, want=1",
			len(args))
	}
	switch arg := args[0].(type) {
	case *object.StringLiteral:
		return &object.Integer{Value: int64(len(arg.Value))}
	// case *object.Array:
	// 	return &object.Integer{Value: int64(len(arg.Elements))}
	default:
		return newError("argument to `len` not supported, got %s",
			args[0].Type())
	}
}

func printFunction(args ...object.Object) object.Object {
	for _, arg := range args {
		switch arg.Type() {
		case object.INTEGER_OBJ:
			print(arg.(*object.Integer).Inspect())
		case object.BOOLEAN_OBJ:
			print(arg.(*object.Boolean).Inspect())
		case object.STRING_OBJ:
			print(arg.(*object.StringLiteral).Inspect())
		case object.FLOAT_OBJ:
			print(arg.(*object.FloatLiteral).Inspect())
		// case object.ARRAY_OBJ:
		// 	print(arg.(*object.Array).Inspect())
		// case object.HASH_OBJ:
		// 	print(arg.(*object.Hash).Inspect())
		case object.NULL_OBJ:
			print(arg.(*object.Null).Inspect())
		default:
			print(arg.Inspect())
		}
	}
	print("\n")
	return nil
}
