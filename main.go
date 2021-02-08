package main

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/digyx/Turnkey/parser"
)

type turnkeyListener struct {
	*parser.BaseTurnkeyListener
}

func main() {
	testLexer()
	testParser()
}

func testParser() {
	is := antlr.NewInputStream("1.9 + 2.7 * 3.2")

	lexer := parser.NewTurnkeyLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewTurnkeyParser(stream)

	antlr.ParseTreeWalkerDefault.Walk(&turnkeyListener{}, p.Start())
}

func testLexer() {
	is := antlr.NewInputStream(`"HelloWorld"`)

	lexer := parser.NewTurnkeyLexer(is)

	for {
		t := lexer.NextToken()
		if t.GetTokenType() == antlr.TokenEOF {
			break
		}

		fmt.Printf("%s (%q)\n", lexer.SymbolicNames[t.GetTokenType()], t.GetText())
	}
}
