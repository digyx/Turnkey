// Code generated from Turnkey.g4 by ANTLR 4.8. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 19, 118,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 3, 2, 6, 2, 39, 10, 2, 13, 2, 14, 2, 40, 3, 2, 3, 2, 3, 3, 3,
	3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3,
	6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 3, 11, 3, 12,
	3, 12, 3, 13, 3, 13, 3, 14, 3, 14, 3, 15, 3, 15, 7, 15, 77, 10, 15, 12,
	15, 14, 15, 80, 11, 15, 3, 16, 3, 16, 7, 16, 84, 10, 16, 12, 16, 14, 16,
	87, 11, 16, 3, 16, 3, 16, 6, 16, 91, 10, 16, 13, 16, 14, 16, 92, 3, 17,
	3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 5, 17, 104, 10,
	17, 3, 18, 7, 18, 107, 10, 18, 12, 18, 14, 18, 110, 11, 18, 3, 18, 3, 18,
	7, 18, 114, 10, 18, 12, 18, 14, 18, 117, 11, 18, 2, 2, 19, 3, 3, 5, 4,
	7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 25, 14,
	27, 15, 29, 16, 31, 17, 33, 18, 35, 19, 3, 2, 8, 5, 2, 11, 12, 15, 15,
	34, 34, 3, 2, 51, 59, 3, 2, 50, 59, 3, 2, 97, 97, 4, 2, 67, 92, 99, 124,
	6, 2, 50, 59, 67, 92, 97, 97, 99, 124, 2, 124, 2, 3, 3, 2, 2, 2, 2, 5,
	3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13,
	3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2,
	21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2,
	2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2,
	2, 3, 38, 3, 2, 2, 2, 5, 44, 3, 2, 2, 2, 7, 46, 3, 2, 2, 2, 9, 51, 3, 2,
	2, 2, 11, 56, 3, 2, 2, 2, 13, 58, 3, 2, 2, 2, 15, 60, 3, 2, 2, 2, 17, 62,
	3, 2, 2, 2, 19, 64, 3, 2, 2, 2, 21, 66, 3, 2, 2, 2, 23, 68, 3, 2, 2, 2,
	25, 70, 3, 2, 2, 2, 27, 72, 3, 2, 2, 2, 29, 74, 3, 2, 2, 2, 31, 81, 3,
	2, 2, 2, 33, 103, 3, 2, 2, 2, 35, 108, 3, 2, 2, 2, 37, 39, 9, 2, 2, 2,
	38, 37, 3, 2, 2, 2, 39, 40, 3, 2, 2, 2, 40, 38, 3, 2, 2, 2, 40, 41, 3,
	2, 2, 2, 41, 42, 3, 2, 2, 2, 42, 43, 8, 2, 2, 2, 43, 4, 3, 2, 2, 2, 44,
	45, 7, 46, 2, 2, 45, 6, 3, 2, 2, 2, 46, 47, 7, 104, 2, 2, 47, 48, 7, 119,
	2, 2, 48, 49, 7, 112, 2, 2, 49, 50, 7, 101, 2, 2, 50, 8, 3, 2, 2, 2, 51,
	52, 7, 118, 2, 2, 52, 53, 7, 119, 2, 2, 53, 54, 7, 116, 2, 2, 54, 55, 7,
	112, 2, 2, 55, 10, 3, 2, 2, 2, 56, 57, 7, 44, 2, 2, 57, 12, 3, 2, 2, 2,
	58, 59, 7, 49, 2, 2, 59, 14, 3, 2, 2, 2, 60, 61, 7, 45, 2, 2, 61, 16, 3,
	2, 2, 2, 62, 63, 7, 47, 2, 2, 63, 18, 3, 2, 2, 2, 64, 65, 7, 63, 2, 2,
	65, 20, 3, 2, 2, 2, 66, 67, 7, 42, 2, 2, 67, 22, 3, 2, 2, 2, 68, 69, 7,
	43, 2, 2, 69, 24, 3, 2, 2, 2, 70, 71, 7, 125, 2, 2, 71, 26, 3, 2, 2, 2,
	72, 73, 7, 127, 2, 2, 73, 28, 3, 2, 2, 2, 74, 78, 9, 3, 2, 2, 75, 77, 9,
	4, 2, 2, 76, 75, 3, 2, 2, 2, 77, 80, 3, 2, 2, 2, 78, 76, 3, 2, 2, 2, 78,
	79, 3, 2, 2, 2, 79, 30, 3, 2, 2, 2, 80, 78, 3, 2, 2, 2, 81, 85, 9, 3, 2,
	2, 82, 84, 9, 4, 2, 2, 83, 82, 3, 2, 2, 2, 84, 87, 3, 2, 2, 2, 85, 83,
	3, 2, 2, 2, 85, 86, 3, 2, 2, 2, 86, 88, 3, 2, 2, 2, 87, 85, 3, 2, 2, 2,
	88, 90, 7, 48, 2, 2, 89, 91, 9, 4, 2, 2, 90, 89, 3, 2, 2, 2, 91, 92, 3,
	2, 2, 2, 92, 90, 3, 2, 2, 2, 92, 93, 3, 2, 2, 2, 93, 32, 3, 2, 2, 2, 94,
	95, 7, 118, 2, 2, 95, 96, 7, 116, 2, 2, 96, 97, 7, 119, 2, 2, 97, 104,
	7, 103, 2, 2, 98, 99, 7, 104, 2, 2, 99, 100, 7, 99, 2, 2, 100, 101, 7,
	110, 2, 2, 101, 102, 7, 117, 2, 2, 102, 104, 7, 103, 2, 2, 103, 94, 3,
	2, 2, 2, 103, 98, 3, 2, 2, 2, 104, 34, 3, 2, 2, 2, 105, 107, 9, 5, 2, 2,
	106, 105, 3, 2, 2, 2, 107, 110, 3, 2, 2, 2, 108, 106, 3, 2, 2, 2, 108,
	109, 3, 2, 2, 2, 109, 111, 3, 2, 2, 2, 110, 108, 3, 2, 2, 2, 111, 115,
	9, 6, 2, 2, 112, 114, 9, 7, 2, 2, 113, 112, 3, 2, 2, 2, 114, 117, 3, 2,
	2, 2, 115, 113, 3, 2, 2, 2, 115, 116, 3, 2, 2, 2, 116, 36, 3, 2, 2, 2,
	117, 115, 3, 2, 2, 2, 10, 2, 40, 78, 85, 92, 103, 108, 115, 3, 8, 2, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "", "','", "'func'", "'turn'", "'*'", "'/'", "'+'", "'-'", "'='", "'('",
	"')'", "'{'", "'}'",
}

var lexerSymbolicNames = []string{
	"", "WS", "COMMA", "FUNC", "TURN", "MUL", "DIV", "ADD", "SUB", "ASSIGN",
	"LPAREN", "RPAREN", "LBRACE", "RBRACE", "INT", "FLOAT", "BOOL", "IDENT",
}

var lexerRuleNames = []string{
	"WS", "COMMA", "FUNC", "TURN", "MUL", "DIV", "ADD", "SUB", "ASSIGN", "LPAREN",
	"RPAREN", "LBRACE", "RBRACE", "INT", "FLOAT", "BOOL", "IDENT",
}

type TurnkeyLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewTurnkeyLexer(input antlr.CharStream) *TurnkeyLexer {

	l := new(TurnkeyLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Turnkey.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// TurnkeyLexer tokens.
const (
	TurnkeyLexerWS     = 1
	TurnkeyLexerCOMMA  = 2
	TurnkeyLexerFUNC   = 3
	TurnkeyLexerTURN   = 4
	TurnkeyLexerMUL    = 5
	TurnkeyLexerDIV    = 6
	TurnkeyLexerADD    = 7
	TurnkeyLexerSUB    = 8
	TurnkeyLexerASSIGN = 9
	TurnkeyLexerLPAREN = 10
	TurnkeyLexerRPAREN = 11
	TurnkeyLexerLBRACE = 12
	TurnkeyLexerRBRACE = 13
	TurnkeyLexerINT    = 14
	TurnkeyLexerFLOAT  = 15
	TurnkeyLexerBOOL   = 16
	TurnkeyLexerIDENT  = 17
)
