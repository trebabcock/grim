package main

import (
	//"fmt"

	"bytes"
	"io"
	"io/ioutil"
)

// Processor holds a Lexer, Buffer, and list of Imports
type Processor struct {
	Lexer   *Lexer
	Buffer  *bytes.Buffer
	Imports []string
}

// NewProcessor creates a new processor
func NewProcessor(reader io.Reader) *Processor {
	l := NewLexer(reader)

	return &Processor{
		Lexer:   l,
		Buffer:  bytes.NewBuffer([]byte{}),
		Imports: []string{},
	}
}

// CheckImports checks to see if the buffer has any imports
func (p *Processor) CheckImports() bool {
	for {
		_, token, _ := p.Lexer.Lex()
		if token == IMPORT {
			pos, tok, lit := p.Lexer.Lex()
			if tok != STR {
				compilerError(pos, "module path", lit, tok)
			}
			if !p.SearchImports(lit) {
				p.Imports = append(p.Imports, lit)
			}
			p.Lexer.Reset()
			return true
		}

		if token == EOF {
			break
		}
	}
	p.Lexer.Reset()
	return false
}

// SearchImports checks to make sure something isn't getting imported twice
func (p *Processor) SearchImports(imp string) bool {
	for _, i := range p.Imports {
		if i == imp {
			return true
		}
	}
	return false
}

// Init sets up the buffer and runs the processing
func (p *Processor) Init() *bytes.Buffer {
	bts := []byte{}
	_, err := p.Lexer.Reader.Read(bts)
	if err != nil {
		panic(err)
	}
	_, err = p.Buffer.Write(bts)
	if err != nil {
		panic(err)
	}
	for {
		if p.CheckImports() {
			p.Process()
		} else {
			break
		}
	}
	return p.Buffer
}

// Process add all files from imported paths to the buffer
func (p *Processor) Process() *bytes.Buffer {
	for _, i := range p.Imports {
		var dir string
		dir = "./"
		files, err := ioutil.ReadDir("./" + i + "/")
		if err != nil {
			panic(err)
		}
		for _, f := range files {
			if !f.IsDir() {
				println(i + "/" + f.Name())
				contents, err := ioutil.ReadFile(dir + i + "/" + f.Name())
				if err != nil {
					genericError(err)
				}
				p.Buffer.Write(contents)
			}
		}
	}
	adf := []byte{}
	_, err := p.Buffer.Read(adf)
	if err != nil {
		panic(err)
	}
	println(string(adf))
	return p.Buffer
}

/*
	read input file(s) and find imports

	add contents of imported files to buffer
	along with original file contents

	send new buffer through processor again
	repeat until there are no more imports
*/
