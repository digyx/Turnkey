package main

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/digyx/Turnkey/parser"
)

type turnkeyListener struct {
	*parser.BaseTurnkeyListener
}

func main() {
	testParser()
}

func testParser() {
	is := antlr.NewInputStream("1.9 + 2.7 * 3.2")

	lexer := parser.NewTurnkeyLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewTurnkeyParser(stream)

	antlr.ParseTreeWalkerDefault.Walk(&turnkeyListener{}, p.Start())
}
