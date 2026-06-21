package lexer_test

import (
	"fmt"
	"os"
	"strings"

	"github.com/ccaroon/mnky/lexer"
	"github.com/ccaroon/mnky/token"
	"go.yaml.in/yaml/v3"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type testFile struct {
	Input    string   `yaml:"input"`
	Expected []string `yaml:"expected"`
}
type testData struct {
	input          string
	expectedTokens []token.Token
}

func readTestFile(filename string) testData {
	var data testData

	path := fmt.Sprintf("./test-code/%s", filename)
	fileContents, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	var fileData testFile
	err = yaml.Unmarshal(fileContents, &fileData)
	if err != nil {
		panic(err)
	}

	data.input = fileData.Input
	for _, expectation := range fileData.Expected {
		tokenData := strings.SplitN(expectation, "|", 2)

		data.expectedTokens = append(
			data.expectedTokens,
			token.Token{
				Type:    token.GetTokenType(tokenData[0]),
				Literal: tokenData[1],
			},
		)
	}

	return data
}

func runYAMLTest(tests testData) {
	lex := lexer.New(tests.input)

	for _, expected := range tests.expectedTokens {
		token := lex.NextToken()

		Expect(token.Type).To(Equal(expected.Type))
		Expect(token.Literal).To(Equal(expected.Literal))
	}
}

var _ = Describe("Lexer", func() {

	It("can get the next token for basic-tokens.yml", func() {
		tests := readTestFile("basic-tokens.yml")
		runYAMLTest(tests)
	})

	It("can get the next token for simple-mnky1.yml", func() {
		tests := readTestFile("simple-mnky1.yml")
		runYAMLTest(tests)
	})

	It("can get the next token for simple-mnky2.yml", func() {
		tests := readTestFile("simple-mnky2.yml")
		runYAMLTest(tests)
	})

})
