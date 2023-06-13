package ast

import (
	"interpreter/token"
	"testing"
)

// let x = 5 in AST form
func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{Type: token.LET, Literal: "let"},
				Name: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "x"},
					Value: "x",
				},
				Value: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "5"},
					Value: "5",
				},
			},
		},
	}

	if program.String() != "let x = 5;" {
		t.Errorf("program.String() wrong. got=%q", program.String())
	}
}
