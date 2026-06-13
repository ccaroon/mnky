package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

func NewToken(tknType TokenType, literal string) Token {
	tkn := Token{
		Type: tknType, Literal: literal,
	}

	return tkn
}
