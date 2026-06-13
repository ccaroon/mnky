package lexer_test

import (
	"github.com/ccaroon/mnky/lexer"
	"github.com/ccaroon/mnky/token"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Lexer", func() {
	var simpleInputTests = []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, token.ASSIGN},
		{token.PLUS, token.PLUS},
		{token.LPAREN, token.LPAREN},
		{token.RPAREN, token.RPAREN},
		{token.LBRACE, token.LBRACE},
		{token.RBRACE, token.RBRACE},
		{token.COMMA, token.COMMA},
		{token.SEMICOLON, token.SEMICOLON},
		{token.EOF, ""},
	}

	It("can get the next token for simple input", func() {
		var input string = `=+(){},;`
		lex := lexer.New(input)

		for _, tt := range simpleInputTests {
			token := lex.NextToken()

			Expect(token.Type).To(Equal(tt.expectedType))
			Expect(token.Literal).To(Equal(tt.expectedLiteral))
		}
	})

	// 	It("can get the next token for small program", func() {
	// 		var input string = `let five = 5;
	// let ten = 10;

	// let add = fn(x, y) {
	//   x + y
	// };

	// let result = add(five, ten);
	// `
	// 	})

})
