package main

import (
	"fmt"
	"strings"
	"testing"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/digyx/Turnkey/parser"
)

func TestLexer(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{`5`, "INT 5"},
		{`5.0`, "FLOAT 5.0"},
		{`true`, "BOOL true"},
		{`alex`, "IDENT alex"},
		{`func this()`, "FUNC:IDENT this:LPAREN:RPAREN"},
		{`func this(a, b)`, "FUNC:IDENT this:LPAREN:IDENT a:IDENT b:RPAREN"},
		{`this()`, "IDENT this:LPAREN:RPAREN"},
		{`this(a, b)`, "IDENT this:LPAREN:IDENT a:IDENT b:RPAREN"},
	}

	for _, test := range tests {
		is := antlr.NewInputStream(test.input)
		lexer := parser.NewTurnkeyLexer(is)

		result := repr(lexer)

		if result != test.expected {
			t.Errorf("expected %s got %s", test.expected, result)
		}
	}
}

// ==================== Helper Functions ====================
func repr(lexer *parser.TurnkeyLexer) string {
	result := []string{}

	for {
		tok := lexer.NextToken()

		if tok.GetTokenType() == antlr.TokenEOF {
			break
		}

		var repr string

		switch tok.GetTokenType() {
		case parser.TurnkeyLexerIDENT, parser.TurnkeyLexerFLOAT,
			parser.TurnkeyLexerINT, parser.TurnkeyLexerBOOL:
			tokType := lexer.SymbolicNames[tok.GetTokenType()]
			repr = fmt.Sprintf("%s %s", tokType, tok.GetText())
		case parser.TurnkeyLexerCOMMA:
			continue
		default:
			repr = lexer.SymbolicNames[tok.GetTokenType()]
		}

		result = append(result, repr)
	}

	return strings.Join(result, ":")
}
