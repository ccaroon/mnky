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

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
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

func (lex *Lexer) readIdentifier() string {
	position := lex.position
	for isLetter(lex.ch) {
		lex.readChar()
	}

	return lex.input[position:lex.position]
}

func (lex *Lexer) NextToken() token.Token {
	var tType token.TokenType
	var literal string

	switch lex.ch {
	// Simple, single char tokens
	case '=':
		tType = token.ASSIGN
		literal = string(lex.ch)
	case ';':
		tType = token.SEMICOLON
		literal = string(lex.ch)
	case '(':
		tType = token.LPAREN
		literal = string(lex.ch)
	case ')':
		tType = token.RPAREN
		literal = string(lex.ch)
	case ',':
		tType = token.COMMA
		literal = string(lex.ch)
	case '+':
		tType = token.PLUS
		literal = string(lex.ch)
	case '{':
		tType = token.LBRACE
		literal = string(lex.ch)
	case '}':
		tType = token.RBRACE
		literal = string(lex.ch)
	case 0:
		tType = token.EOF
		literal = token.EOF
	default:
		if isLetter(lex.ch) {
			tType = token.IDENT
			literal = lex.readIdentifier()
		} else {
			tType = token.ILLEGAL
			literal = string(lex.ch)
		}
	}

	tkn := token.NewToken(tType, literal)

	lex.readChar()

	return tkn
}
