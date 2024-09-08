package main

import (
	"fmt"
	"math"
	"os"
	"strings"

	"github.com/Ayobami0/yoruba/src/evaluator"
	"github.com/Ayobami0/yoruba/src/lexer"
	"github.com/Ayobami0/yoruba/src/object"
	"github.com/Ayobami0/yoruba/src/parser"
)

func main() {
	args := os.Args

	if len(args) <= 1 {
		fmt.Println("Usage: yoruba [FILE]")
		os.Exit(1)
	}
	fname := args[1]
	f, err := os.Open(fname)

	if err != nil {
		switch {
		case os.IsNotExist(err):
			fmt.Printf("aṣiṣe: %s ko wa ni be nibi\n", fname)
		}
		os.Exit(1)
	}

	l := lexer.New(f)
	p := parser.New(l)
	env := object.NewEnvironment()
	program := p.ParseProgram()

	if len(p.Errors()) != 0 {
		printParserErrors(p.Errors())
	} else {
    if obj := evaluator.Eval(program, env); obj.Type() == object.ERROR_OBJ {
      fmt.Println(obj.Inspect())
    }
	}

}

func printParserErrors(errors []string) {
	fmt.Println("Awọn aṣiṣe:")
	lenOfPad := len(errors)

	for i, msg := range errors {
		fmt.Printf("\t%d.%s %s\n", i+1, buildPadding(lenOfPad), msg)
	}
	fmt.Printf("Aṣiṣe %d lapapọ\n", lenOfPad)
}

func buildPadding(length int) string {
	l := int(math.Log10(float64(length))) + 1
	var pad strings.Builder
	for i := 0; i < l; i++ {
		pad.WriteByte(' ')

	}
	return pad.String()
}
