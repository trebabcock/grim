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
	buf := proc.Init()

	lexer := NewLexer(buf)
	for {
		pos, tok, lit := lexer.Lex()
		if tok == EOF {
			break
		}
		fmt.Println(strconv.Itoa(pos.Line) + ":" + strconv.Itoa(pos.Column) + "\t" + tokens[tok] + "\t" + lit)
	}
}

func compilerError(pos Position, expected string, actual string, token Token) {
	fmt.Println(strconv.Itoa(pos.Line) + ":" + strconv.Itoa(pos.Column) + "\tExpected " + expected + ", got " + actual + " (" + tokens[token] + ")")
	os.Exit(1)
}

func genericError(err error) {
	fmt.Println(err.Error())
	os.Exit(1)
}
