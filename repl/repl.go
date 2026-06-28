package repl

import (
	"bufio"
	"fmt"

	"github.com/ccaroon/mnky/token"

	"io"

	"github.com/ccaroon/mnky/lexer"
)

const PROMPT = ">>> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprint(out, PROMPT)

		input := scanner.Scan()
		if !input {
			err := scanner.Err()
			if err != nil {
				fmt.Fprint(out, err.Error())
			} else {
				fmt.Fprintln(out, "Exiting")
				break
			}
		}

		line := scanner.Text()
		lexer := lexer.New(line)

		for tkn := lexer.NextToken(); tkn.Type != token.EOF; tkn = lexer.NextToken() {
			fmt.Fprintf(out, "%+v\n", tkn)
		}
	}
}
