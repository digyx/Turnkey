package repl

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/digyx/Turnkey/parser"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprint(out, PROMPT)

		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		is := antlr.NewInputStream(line)

		l := parser.NewTurnkeyLexer(is)
		fmt.Fprint(out, reprLexer(l))
	}
}

func reprLexer(lexer *parser.TurnkeyLexer) string {
	out := []string{}

	for {
		t := lexer.NextToken()
		if t.GetTokenType() == antlr.TokenEOF {
			break
		}

		repr := fmt.Sprintf("%s %s\n", lexer.SymbolicNames[t.GetTokenType()], t.GetText())
		out = append(out, repr)
	}

	return strings.Join(out, "")
}
