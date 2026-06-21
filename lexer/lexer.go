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

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
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

func (lex *Lexer) unreadChar() {
	if lex.readPosition > 0 {
		lex.position -= 1
		lex.ch = lex.input[lex.position]
		lex.readPosition -= 1
	}
}

func (lex *Lexer) peekChar() byte {
	var nextChar byte

	if lex.readPosition >= len(lex.input) {
		nextChar = 0
	} else {
		nextChar = lex.input[lex.readPosition]
	}

	return nextChar
}

// Reads a sequence of bytes from input, return as string
func (lex *Lexer) readSequence(testFunc func(byte) bool) string {
	startPos := lex.position
	for testFunc(lex.ch) {
		lex.readChar()
	}
	endPos := lex.position

	// Unread last char
	lex.unreadChar()

	return lex.input[startPos:endPos]
}

func (lex *Lexer) skipWhitespace() {
	for lex.ch == ' ' || lex.ch == '\t' || lex.ch == '\n' || lex.ch == '\r' {
		lex.readChar()
	}
}

func (lex *Lexer) NextToken() token.Token {
	var tType token.TokenType
	var literal string

	lex.skipWhitespace()

	switch lex.ch {
	// Simple, single char tokens
	case '!':
		if lex.peekChar() == '=' {
			lex.readChar()
			tType = token.NOT_EQ
			literal = "!="
		} else {
			tType = token.BANG
			literal = string(lex.ch)
		}
	case '=':
		if lex.peekChar() == '=' {
			lex.readChar()
			tType = token.EQ
			literal = "=="
		} else {
			tType = token.ASSIGN
			literal = string(lex.ch)
		}
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
	case '-':
		tType = token.MINUS
		literal = string(lex.ch)
	case '/':
		tType = token.SLASH
		literal = string(lex.ch)
	case '*':
		tType = token.ASTERISK
		literal = string(lex.ch)
	case '<':
		tType = token.LESS_THAN
		literal = string(lex.ch)
	case '>':
		tType = token.GREATER_THAN
		literal = string(lex.ch)
	case 0:
		tType = token.EOF
		literal = ""
	default:
		if isLetter(lex.ch) {
			literal = lex.readSequence(isLetter)
			tType = token.LookupIdent(literal)
		} else if isDigit(lex.ch) {
			literal = lex.readSequence(isDigit)
			tType = token.INT
		} else {
			tType = token.ILLEGAL
			literal = string(lex.ch)
		}
	}

	tkn := token.NewToken(tType, literal)

	lex.readChar()

	return tkn
}
