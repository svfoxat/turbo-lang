package object

import (
	"bytes"
	"fmt"
	"strings"
	"turbo/ast"
)

type ObjectType string

type Object interface {
	Type() ObjectType
	Inspect() string
}

type BuiltinFunction func(args ...Object) Object

type Builtin struct {
	Fn BuiltinFunction
}

func (b *Builtin) Type() ObjectType {
	return FUNCTION_OBJ
}
func (b *Builtin) Inspect() string {
	return "builtin function"
}

const (
	INTEGER_OBJ      = "INTEGER"
	BOOLEAN_OBJ      = "BOOLEAN"
	NULL_OBJ         = "NULL"
	RETURN_VALUE_OBJ = "RETURN_VALUE"
	ERROR_OBJ        = "ERROR"
	FUNCTION_OBJ     = "FUNCTION"
	STRING_OBJ       = "STRING"
	FLOAT_OBJ        = "FLOAT"
)

type Integer struct {
	Value int64
}

func (i *Integer) Type() ObjectType {
	return INTEGER_OBJ
}
func (i *Integer) Inspect() string {
	return fmt.Sprintf("%d", i.Value)
}

type Boolean struct {
	Value bool
}

func (b *Boolean) Type() ObjectType {
	return BOOLEAN_OBJ
}
func (b *Boolean) Inspect() string {
	return fmt.Sprintf("%t", b.Value)
}

type Null struct{}

func (n *Null) Type() ObjectType {
	return NULL_OBJ
}
func (n *Null) Inspect() string {
	return "null"
}

type ReturnValue struct {
	Value Object
}

func (rv *ReturnValue) Type() ObjectType {
	return RETURN_VALUE_OBJ
}
func (rv *ReturnValue) Inspect() string {
	return fmt.Sprintf("%s", rv.Value.Inspect())
}

type Error struct {
	Message string
}

func (e *Error) Type() ObjectType {
	return ERROR_OBJ
}
func (e *Error) Inspect() string {
	return fmt.Sprintf("[Error %s]", e.Message)
}

type FunctionLiteral struct {
	Parameters []*ast.Identifier
	Body       *ast.BlockStatement
	Env        *Environment
}

func (f *FunctionLiteral) Type() ObjectType {
	return FUNCTION_OBJ
}
func (f *FunctionLiteral) Inspect() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range f.Parameters {
		params = append(params, p.String())
	}

	out.WriteString("fn")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") {\n")
	out.WriteString(f.Body.String())
	out.WriteString("\n}")

	return out.String()
}

type StringLiteral struct {
	Value string
}

func (s *StringLiteral) Type() ObjectType {
	return STRING_OBJ
}
func (s *StringLiteral) Inspect() string {
	return s.Value
}

type FloatLiteral struct {
	Value float64
}

func (f *FloatLiteral) Type() ObjectType {
	return FLOAT_OBJ
}
func (f *FloatLiteral) Inspect() string {
	return fmt.Sprintf("%f", f.Value)
}
