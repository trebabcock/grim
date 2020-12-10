package main

import (
	"bufio"
	"io"
	"unicode"
)

type Token int

type TokenFormat struct {
	Pos Position
	Tok Token
	Lit string
}

const (
	EOF = iota
	ILLEGAL
	IDENT
	NUM
	DEC
	STR

	PLUS
	MINUS
	ASTR
	FSLASH

	DPLUS
	DMINUS

	SEMI
	COMMA
	PERIOD

	CARET
	EXCL
	QUES
	PIPE
	DPIPE
	AMP
	DAMP
	BSLASH

	LPAREN
	RPAREN
	LSQUARE
	RSQUARE
	LCURLY
	RCURLY

	EQ
	EQEQ
	GRT
	LES
	GRTEQ
	LESEQ
	NOTEQ

	IF
	ELSE
	WHILE
	FOR
	FUNCTION
	OBJECT
	DEFINE
	TRUE
	FALSE
	ASM
	IMPORT
	SWITCH
	CASE

	FLOAT
	FLOAT32
	FLOAT64
	INT
	INT8
	INT16
	INT32
	INT64
	UINT
	UINT8
	UINT16
	UINT32
	UINT64
	STRING
	BOOL
	BYTE
	CHAR
	ENUM
	FN
	VAR
)

var tokens = []string{
	EOF:     "EOF",
	ILLEGAL: "ILLEGAL",
	IDENT:   "IDENT",

	NUM: "NUM",
	DEC: "DEC",
	STR: "STR",

	PLUS:   "+",
	MINUS:  "-",
	ASTR:   "*",
	FSLASH: "/",

	DPLUS:  "++",
	DMINUS: "--",

	SEMI:   ";",
	COMMA:  ",",
	PERIOD: ".",

	CARET:  "^",
	EXCL:   "!",
	QUES:   "?",
	PIPE:   "|",
	DPIPE:  "||",
	AMP:    "&",
	DAMP:   "&&",
	BSLASH: "\\",

	LPAREN:  "(",
	RPAREN:  ")",
	LSQUARE: "[",
	RSQUARE: "]",
	LCURLY:  "{",
	RCURLY:  "}",

	EQ:    "=",
	EQEQ:  "==",
	GRT:   ">",
	LES:   "<",
	GRTEQ: ">=",
	LESEQ: "<=",
	NOTEQ: "!=",

	IF:       "if",
	ELSE:     "else",
	WHILE:    "while",
	FOR:      "for",
	FUNCTION: "function",
	OBJECT:   "object",
	DEFINE:   "define",
	TRUE:     "true",
	FALSE:    "false",
	ASM:      "asm",
	IMPORT:   "import",
	SWITCH:   "switch",
	CASE:     "case",

	FLOAT:   "float",
	FLOAT32: "float32",
	FLOAT64: "float64",
	INT:     "int",
	INT8:    "int8",
	INT16:   "int16",
	INT32:   "int32",
	INT64:   "int64",
	UINT:    "uint",
	UINT8:   "uint8",
	UINT16:  "uint16",
	UINT32:  "uint32",
	UINT64:  "uing64",
	STRING:  "string",
	BOOL:    "bool",
	BYTE:    "byte",
	CHAR:    "char",
	ENUM:    "enum",
	FN:      "fn",
	VAR:     "var",
}

func (t *Token) String() string {
	return tokens[*t]
}

type Position struct {
	Line   int
	Column int
}

type Lexer struct {
	Pos    Position
	Reader *bufio.Reader
}

func NewLexer(reader io.Reader) *Lexer {
	return &Lexer{
		Pos:    Position{Line: 1, Column: 0},
		Reader: bufio.NewReader(reader),
	}
}

func (l *Lexer) Lex() (Position, Token, string) {
	for {
		r, _, err := l.Reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				return l.Pos, EOF, "EOF"
			}
			panic(err)
		}
		l.Pos.Column++

		switch r {
		case '\n':
			l.resetPosition()
		case ';':
			return l.Pos, SEMI, tokens[SEMI]
		case ',':
			return l.Pos, COMMA, tokens[COMMA]
		case '.':
			return l.Pos, PERIOD, tokens[PERIOD]
		case '^':
			return l.Pos, CARET, tokens[CARET]
		case '(':
			return l.Pos, LPAREN, tokens[LPAREN]
		case ')':
			return l.Pos, RPAREN, tokens[RPAREN]
		case '[':
			return l.Pos, LSQUARE, tokens[LSQUARE]
		case ']':
			return l.Pos, RSQUARE, tokens[RSQUARE]
		case '{':
			return l.Pos, LCURLY, tokens[LCURLY]
		case '}':
			return l.Pos, RCURLY, tokens[RCURLY]
		case '*':
			return l.Pos, ASTR, tokens[ASTR]
		case '/':
			return l.Pos, FSLASH, tokens[FSLASH]
		case '\\':
			return l.Pos, BSLASH, tokens[BSLASH]
		case '?':
			return l.Pos, QUES, tokens[QUES]
		case '!':
			startPos := l.Pos
			if l.peek() == '=' {
				l.skip()
				return startPos, NOTEQ, tokens[NOTEQ]
			}
			return startPos, EXCL, tokens[EXCL]
		case '|':
			startPos := l.Pos
			if l.peek() == '|' {
				l.skip()
				return startPos, DPIPE, tokens[DPIPE]
			}
			return startPos, PIPE, tokens[PIPE]
		case '&':
			startPos := l.Pos
			if l.peek() == '&' {
				l.skip()
				return startPos, DAMP, tokens[DAMP]
			}
			return startPos, AMP, tokens[AMP]
		case '+':
			startPos := l.Pos
			if l.peek() == '+' {
				l.skip()
				return startPos, DPLUS, tokens[DPLUS]
			}
			return startPos, PLUS, tokens[PLUS]
		case '-':
			startPos := l.Pos
			if l.peek() == '-' {
				l.skip()
				return startPos, DMINUS, tokens[DMINUS]
			}
			return startPos, MINUS, tokens[MINUS]
		case '=':
			startPos := l.Pos
			if l.peek() == '=' {
				l.skip()
				return startPos, EQEQ, tokens[EQEQ]
			}
			return startPos, EQ, tokens[EQ]
		case '>':
			startPos := l.Pos
			if l.peek() == '=' {
				l.skip()
				return startPos, GRTEQ, tokens[GRTEQ]
			}
			return startPos, GRT, tokens[GRT]
		case '<':
			startPos := l.Pos
			if l.peek() == '=' {
				l.skip()
				return startPos, LESEQ, tokens[LESEQ]
			}
			return startPos, LES, tokens[LES]
		case '"':
			startPos := l.Pos
			lit := l.lexString()
			return startPos, STR, lit
		default:
			if unicode.IsSpace(r) {
				continue
			} else if unicode.IsDigit(r) {
				startPos := l.Pos
				l.backup()
				lit, t := l.lexNum()
				return startPos, t, lit
			} else if unicode.IsLetter(r) {
				startPos := l.Pos
				l.backup()
				lit := l.lexIdent()
				switch lit {
				case "if":
					return startPos, IF, lit
				case "else":
					return startPos, ELSE, lit
				case "WHILE":
					return startPos, WHILE, lit
				case "for":
					return startPos, FOR, lit
				case "function":
					return startPos, FUNCTION, lit
				case "object":
					return startPos, OBJECT, lit
				case "define":
					return startPos, DEFINE, lit
				case "true":
					return startPos, TRUE, lit
				case "false":
					return startPos, FALSE, lit
				case "asm":
					return startPos, ASM, lit
				case "import":
					return startPos, IMPORT, lit
				case "int":
					return startPos, INT32, lit
				case "int8":
					return startPos, INT8, lit
				case "int16":
					return startPos, INT16, lit
				case "int32":
					return startPos, INT32, lit
				case "int64":
					return startPos, INT64, lit
				case "uint":
					return startPos, UINT32, lit
				case "uint8":
					return startPos, UINT8, lit
				case "uint16":
					return startPos, UINT16, lit
				case "uint32":
					return startPos, UINT32, lit
				case "uint64":
					return startPos, UINT64, lit
				case "float":
					return startPos, FLOAT32, lit
				case "foat32":
					return startPos, FLOAT32, lit
				case "float64":
					return startPos, FLOAT64, lit
				case "bool":
					return startPos, BOOL, lit
				case "string":
					return startPos, STRING, lit
				case "byte":
					return startPos, BYTE, lit
				case "char":
					return startPos, CHAR, lit
				case "enum":
					return startPos, ENUM, lit
				case "fn":
					return startPos, FN, lit
				case "var":
					return startPos, VAR, lit
				case "switch":
					return startPos, SWITCH, lit
				case "case":
					return startPos, CASE, lit
				default:
					return startPos, IDENT, lit
				}
			} else {
				return l.Pos, ILLEGAL, string(r)
			}
		}
	}
}

func (l *Lexer) resetPosition() {
	l.Pos.Line++
	l.Pos.Column = 0
}

func (l *Lexer) backup() {
	if err := l.Reader.UnreadRune(); err != nil {
		panic(err)
	}

	l.Pos.Column--
}

func (l *Lexer) skip() {
	_, _, err := l.Reader.ReadRune()
	if err != nil {
		if err == io.EOF {
			return
		}
		panic(err)
	}
	l.Pos.Column++
}

func (l *Lexer) peek() rune {
	r, _, err := l.Reader.ReadRune()
	if err != nil {
		if err == io.EOF {
			return r
		}
		panic(err)
	}
	l.Pos.Column++
	l.backup()
	return r
}

func (l *Lexer) lexNum() (string, Token) {
	var lit string
	var t Token = NUM
	for {
		r, _, err := l.Reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				return lit, t
			}
			panic(err)
		}

		l.Pos.Column++
		if unicode.IsDigit(r) {
			lit = lit + string(r)
		} else if r == '.' {
			t = DEC
			lit = lit + string(r)
		} else {
			l.backup()
			return lit, t
		}
	}
}

func (l *Lexer) lexIdent() string {
	var lit string
	for {
		r, _, err := l.Reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				return lit
			}
			panic(err)
		}
		l.Pos.Column++
		if unicode.IsLetter(r) {
			lit = lit + string(r)
		} else {
			l.backup()
			return lit
		}
	}
}

func (l *Lexer) lexString() string {
	var lit string
	for {
		r, _, err := l.Reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				return lit
			}
			panic(err)
		}
		l.Pos.Column++
		if r == '"' {
			return lit
		} else if r == '\\' {
			lit = lit + string(r)
			if l.peek() == '"' {
				lit = lit + "\""
				l.skip()
			}
		} else {
			lit = lit + string(r)
		}
	}
}
