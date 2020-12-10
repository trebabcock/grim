package main

import (
	"bufio"
)

type Parser struct {
	Lexer *Lexer
}

func NewPaser(reader *bufio.Reader) *Parser {
	l := NewLexer(reader)
	return &Parser{
		Lexer: l,
	}
}

func (p *Parser) Parse() {

}
