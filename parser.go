package main

import (
	"bufio"
)

/*
	This file is pretty much useless right now. Documentation will
	be made better when more code is added.
*/

// Parser is the parser type
type Parser struct {
	Lexer *Lexer
}

// NewParser creates a new parser
func NewParser(reader *bufio.Reader) *Parser {
	l := NewLexer(reader)
	return &Parser{
		Lexer: l,
	}
}

// Parse will one day parse the lexer's output into an AST
func (p *Parser) Parse() {

}
