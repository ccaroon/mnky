package token

import "strings"

const (
	ILLEGAL = "ILLEGAL"
	EOF     = ""

	// Identifier & literals
	IDENT = "IDENT"
	INT   = "INT"

	// Operators
	ASSIGN = "="
	PLUS   = "+"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

// Mainly for Unit Tests
func GetTokenType(name string) TokenType {
	id := strings.ToUpper(name)

	tknType, _ := TokenByName[id]

	return tknType
}

var TokenByName map[string]TokenType = map[string]TokenType{
	"ILLEGAL":   ILLEGAL,
	"EOF":       "",
	"IDENT":     IDENT,
	"INT":       INT,
	"ASSIGN":    ASSIGN,
	"PLUS":      PLUS,
	"COMMA":     COMMA,
	"SEMICOLON": SEMICOLON,
	"LPAREN":    LPAREN,
	"RPAREN":    RPAREN,
	"LBRACE":    LBRACE,
	"RBRACE":    RBRACE,
	"FUNCTION":  FUNCTION,
	"LET":       LET,
}
