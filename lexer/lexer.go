package lexer

import "github.com/ccaroon/mnky/token"

type Lexer struct {
	input        string
	position     int  // postion in input (point to current char)
	readPosition int  // reading position in input (after current char)
	ch           byte // char being examined
}

func New(input string) *Lexer {
	lex := Lexer{input: input}

	lex.readChar()

	return &lex
}

func (lex *Lexer) readChar() {
	if lex.readPosition >= len(lex.input) {
		lex.ch = 0
	} else {
		lex.ch = lex.input[lex.readPosition]
	}

	lex.position = lex.readPosition
	lex.readPosition += 1
}

func (lex *Lexer) NextToken() token.Token {
	var tkn token.Token

	switch lex.ch {
	// EOF
	case 0:
		tkn = token.NewToken(token.EOF, token.EOF)
	// Simple, single char token
	default:
		tkn = token.NewToken(token.TokenType(lex.ch), string(lex.ch))
	}

	lex.readChar()

	return tkn
}
