package repl

import (
	"bufio"
	"fmt"
	"io"
	"turbo/interpreter"
	"turbo/lexer"
	"turbo/parser"
)

const PROMPT = ">>> "

func Start(in io.Reader, out io.Writer) {
	printBanner()
	scanner := bufio.NewScanner(in)
	env := interpreter.NewEnvironment()

	for {
		fmt.Fprint(out, PROMPT)
		scanned := scanner.Scan()

		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		evaluated := interpreter.Eval(program, env)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

func printParserErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}

func printBanner() {
	fmt.Println("Seas hawara, I bims 1 Interpreter lol")
}
