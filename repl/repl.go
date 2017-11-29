package repl

import (
	"bufio"
	"fmt"
	"github.com/onidoru/goemon/evaluator"
	"github.com/onidoru/goemon/lexer"
	"github.com/onidoru/goemon/object"
	"github.com/onidoru/goemon/parser"
	"io"
)

const PROMPT = ">> "

// func Start(in io.Reader, out io.Writer) {
// 	scanner := bufio.NewScanner(in)
// 	for {
// 		fmt.Printf(PROMPT)
// 		scanned := scanner.Scan()
// 		if !scanned {
// 			return
// 		}
// 		line := scanner.Text()
// 		l := lexer.New(line)
// 		p := parser.New(l)
// 		program := p.ParseProgram()
// 		if len(p.Errors()) != 0 {
// 			printParserErrors(out, p.Errors())
// 			continue
// 		}
// 		io.WriteString(out, program.String())
// 		io.WriteString(out, "\n")
// 	}
// }

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	for {
		fmt.Printf(PROMPT)
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
		evaluated := evaluator.Eval(program, env)
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
