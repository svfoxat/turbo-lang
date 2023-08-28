package main

import (
	"os"
	"turbo/repl"
)

func main() {
	repl.Start(os.Stdin, os.Stdout)
}
