package main

import (
	"fmt"
	"interpreter/repl"
	"os"
)

func main() {
	fmt.Printf("Hello! This is some interpreted programming language!\n")
	repl.Start(os.Stdin, os.Stdout)
}
