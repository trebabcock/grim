package main

import (
	//"fmt"
	"bufio"
	"io"
	"io/ioutil"
)

type Processor struct {
	Lexer *Lexer
}

func NewProcessor(reader io.Reader) *Processor {
	l := NewLexer(reader)
	return &Processor{
		Lexer: l,
	}
}

func (p *Processor) Process() *bufio.Reader {
	tokns := []TokenFormat{}
	imports := []string{}
	for {
		position, token, literal := p.Lexer.Lex()
		if token == EOF {
			break
		}
		if token == IMPORT {
			tokns = append(tokns, TokenFormat{position, token, literal})
			pos, tok, lit := p.Lexer.Lex()
			tokns = append(tokns, TokenFormat{pos, tok, lit})
			imports = append(imports, lit)
		}
	}
	for _, t := range tokns {
		println(tokens[t.Tok] + "\t" + t.Lit)
	}
	metafile := new(bufio.ReadWriter)
	for _, i := range imports {
		files, err := ioutil.ReadDir("$GRIMPATH" + i)
		if err != nil {
			f, er := ioutil.ReadDir("./" + i)
			if er != nil {
				panic(err)
			}
		}
		for _, f := range files {
			io.Copy(metafile.Writer, f)
		}
	}
	return metafile.Reader
}
