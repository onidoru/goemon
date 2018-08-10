package main

import (
	"github.com/Onidoru/goemon/repl"
	"os"
)

func main() {
	repl.Start(os.Stdin, os.Stdout)
}
