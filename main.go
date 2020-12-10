package main

import (
	"fmt"
	"os"
	"strconv"
)

func test() {
	file, err := os.Open("test.grim")
	if err != nil {
		panic(err)
	}

	lexer := NewLexer(file)
	for {
		pos, tok, lit := lexer.Lex()
		if tok == EOF {
			break
		}
		fmt.Println(strconv.Itoa(pos.Line) + ":" + strconv.Itoa(pos.Column) + "\t" + tokens[tok] + "\t" + lit)
	}
}

func main() {
	file, err := os.Open("test.grim")
	if err != nil {
		panic(err)
	}

	proc := NewProcessor(file)
	proc.Process()
}
