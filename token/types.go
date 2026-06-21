package token

import "strings"

const (
	ASSIGN       = "ASSIGN"
	ASTERISK     = "ASTERISK"
	BANG         = "BANG"
	COMMA        = "COMMA"
	ELSE         = "ELSE"
	EOF          = "EOF"
	EQ           = "EQ"
	FALSE        = "FALSE"
	FUNCTION     = "FUNCTION"
	GREATER_THAN = "GREATER_THAN"
	IDENT        = "IDENT"
	IF           = "IF"
	ILLEGAL      = "ILLEGAL"
	INT          = "INT"
	LBRACE       = "LBRACE"
	LESS_THAN    = "LESS_THAN"
	LET          = "LET"
	LPAREN       = "LPAREN"
	MINUS        = "MINUS"
	NOT_EQ       = "NOT_EQ"
	PLUS         = "PLUS"
	RBRACE       = "RBRACE"
	RETURN       = "RETURN"
	RPAREN       = "RPAREN"
	SEMICOLON    = "SEMICOLON"
	SLASH        = "SLASH"
	TRUE         = "TRUE"
)

var keywords map[string]TokenType = map[string]TokenType{
	"else":   ELSE,
	"false":  FALSE,
	"fn":     FUNCTION,
	"if":     IF,
	"let":    LET,
	"return": RETURN,
	"true":   TRUE,
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
	"ASSIGN":       ASSIGN,
	"ASTERISK":     ASTERISK,
	"BANG":         BANG,
	"COMMA":        COMMA,
	"ELSE":         ELSE,
	"EOF":          EOF,
	"EQ":           EQ,
	"FALSE":        FALSE,
	"FUNCTION":     FUNCTION,
	"GREATER_THAN": GREATER_THAN,
	"IDENT":        IDENT,
	"IF":           IF,
	"ILLEGAL":      ILLEGAL,
	"INT":          INT,
	"LBRACE":       LBRACE,
	"LESS_THAN":    LESS_THAN,
	"LET":          LET,
	"LPAREN":       LPAREN,
	"MINUS":        MINUS,
	"NOT_EQ":       NOT_EQ,
	"PLUS":         PLUS,
	"RBRACE":       RBRACE,
	"RETURN":       RETURN,
	"RPAREN":       RPAREN,
	"SEMICOLON":    SEMICOLON,
	"SLASH":        SLASH,
	"TRUE":         TRUE,
}
