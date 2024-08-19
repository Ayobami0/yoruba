package main

import (
	"fmt"
	"os"

	"github.com/Ayobami0/yoruba/src/lexer"
	"github.com/Ayobami0/yoruba/src/token"
)

func main() {
	args := os.Args

	if len(args) <= 1 {
		fmt.Println("Usage: yoru [FILE]")
		os.Exit(1)
	}
	fname := args[1]
	f, err := os.Open(fname)

	if err != nil {
		switch {
		case os.IsNotExist(err):
			fmt.Printf("ERROR: file %s not found\n", fname)
		}
		os.Exit(1)
	}

	l := lexer.New(f)

  for t := l.NextToken(); t.Type != token.EOF; t = l.NextToken() {
    fmt.Println(t)
  }
}
