package token

import "strings"

const (
	ILLEGAL   = "ILLEGAL"
	EOF       = "EOF"
	IDENT     = "IDENT"
	INT       = "INT"
	ASSIGN    = "ASSIGN"
	PLUS      = "PLUS"
	COMMA     = "COMMA"
	SEMICOLON = "SEMICOLON"
	LPAREN    = "LPAREN"
	RPAREN    = "RPAREN"
	LBRACE    = "LBRACE"
	RBRACE    = "RBRACE"
	FUNCTION  = "FUNCTION"
	LET       = "LET"
)

var keywords map[string]TokenType = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

func LookupIdent(ident string) TokenType {
	var tknType TokenType

	if tType, ok := keywords[ident]; ok {
		tknType = tType
	} else {
		tknType = IDENT
	}

	return tknType
}

// Mainly for Unit Tests
func GetTokenType(name string) TokenType {
	id := strings.ToUpper(name)

	tknType, _ := TokenByName[id]

	return tknType
}

var TokenByName map[string]TokenType = map[string]TokenType{
	"ILLEGAL":   ILLEGAL,
	"EOF":       EOF,
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
