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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 17, 116,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 3, 2, 6, 2, 35,
	10, 2, 13, 2, 14, 2, 36, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4,
	3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8,
	3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 3, 11, 3, 12, 3, 12, 7, 12, 67, 10, 12,
	12, 12, 14, 12, 70, 11, 12, 3, 13, 3, 13, 7, 13, 74, 10, 13, 12, 13, 14,
	13, 77, 11, 13, 3, 13, 3, 13, 6, 13, 81, 10, 13, 13, 13, 14, 13, 82, 3,
	14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 5, 14, 94,
	10, 14, 3, 15, 3, 15, 6, 15, 98, 10, 15, 13, 15, 14, 15, 99, 3, 15, 3,
	15, 3, 16, 7, 16, 105, 10, 16, 12, 16, 14, 16, 108, 11, 16, 3, 16, 3, 16,
	7, 16, 112, 10, 16, 12, 16, 14, 16, 115, 11, 16, 2, 2, 17, 3, 3, 5, 4,
	7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 25, 14,
	27, 15, 29, 16, 31, 17, 3, 2, 9, 5, 2, 11, 12, 15, 15, 34, 34, 3, 2, 51,
	59, 3, 2, 50, 59, 5, 2, 50, 59, 67, 92, 99, 124, 3, 2, 97, 97, 3, 2, 99,
	124, 6, 2, 50, 59, 67, 92, 97, 97, 99, 124, 2, 123, 2, 3, 3, 2, 2, 2, 2,
	5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2,
	13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2,
	2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2,
	2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 3, 34, 3, 2, 2, 2, 5, 40, 3, 2,
	2, 2, 7, 45, 3, 2, 2, 2, 9, 50, 3, 2, 2, 2, 11, 52, 3, 2, 2, 2, 13, 54,
	3, 2, 2, 2, 15, 56, 3, 2, 2, 2, 17, 58, 3, 2, 2, 2, 19, 60, 3, 2, 2, 2,
	21, 62, 3, 2, 2, 2, 23, 64, 3, 2, 2, 2, 25, 71, 3, 2, 2, 2, 27, 93, 3,
	2, 2, 2, 29, 95, 3, 2, 2, 2, 31, 106, 3, 2, 2, 2, 33, 35, 9, 2, 2, 2, 34,
	33, 3, 2, 2, 2, 35, 36, 3, 2, 2, 2, 36, 34, 3, 2, 2, 2, 36, 37, 3, 2, 2,
	2, 37, 38, 3, 2, 2, 2, 38, 39, 8, 2, 2, 2, 39, 4, 3, 2, 2, 2, 40, 41, 7,
	104, 2, 2, 41, 42, 7, 119, 2, 2, 42, 43, 7, 112, 2, 2, 43, 44, 7, 101,
	2, 2, 44, 6, 3, 2, 2, 2, 45, 46, 7, 118, 2, 2, 46, 47, 7, 119, 2, 2, 47,
	48, 7, 116, 2, 2, 48, 49, 7, 112, 2, 2, 49, 8, 3, 2, 2, 2, 50, 51, 7, 44,
	2, 2, 51, 10, 3, 2, 2, 2, 52, 53, 7, 49, 2, 2, 53, 12, 3, 2, 2, 2, 54,
	55, 7, 45, 2, 2, 55, 14, 3, 2, 2, 2, 56, 57, 7, 47, 2, 2, 57, 16, 3, 2,
	2, 2, 58, 59, 7, 63, 2, 2, 59, 18, 3, 2, 2, 2, 60, 61, 7, 42, 2, 2, 61,
	20, 3, 2, 2, 2, 62, 63, 7, 43, 2, 2, 63, 22, 3, 2, 2, 2, 64, 68, 9, 3,
	2, 2, 65, 67, 9, 4, 2, 2, 66, 65, 3, 2, 2, 2, 67, 70, 3, 2, 2, 2, 68, 66,
	3, 2, 2, 2, 68, 69, 3, 2, 2, 2, 69, 24, 3, 2, 2, 2, 70, 68, 3, 2, 2, 2,
	71, 75, 9, 3, 2, 2, 72, 74, 9, 4, 2, 2, 73, 72, 3, 2, 2, 2, 74, 77, 3,
	2, 2, 2, 75, 73, 3, 2, 2, 2, 75, 76, 3, 2, 2, 2, 76, 78, 3, 2, 2, 2, 77,
	75, 3, 2, 2, 2, 78, 80, 7, 48, 2, 2, 79, 81, 9, 4, 2, 2, 80, 79, 3, 2,
	2, 2, 81, 82, 3, 2, 2, 2, 82, 80, 3, 2, 2, 2, 82, 83, 3, 2, 2, 2, 83, 26,
	3, 2, 2, 2, 84, 85, 7, 118, 2, 2, 85, 86, 7, 116, 2, 2, 86, 87, 7, 119,
	2, 2, 87, 94, 7, 103, 2, 2, 88, 89, 7, 104, 2, 2, 89, 90, 7, 99, 2, 2,
	90, 91, 7, 110, 2, 2, 91, 92, 7, 117, 2, 2, 92, 94, 7, 103, 2, 2, 93, 84,
	3, 2, 2, 2, 93, 88, 3, 2, 2, 2, 94, 28, 3, 2, 2, 2, 95, 97, 7, 36, 2, 2,
	96, 98, 9, 5, 2, 2, 97, 96, 3, 2, 2, 2, 98, 99, 3, 2, 2, 2, 99, 97, 3,
	2, 2, 2, 99, 100, 3, 2, 2, 2, 100, 101, 3, 2, 2, 2, 101, 102, 7, 36, 2,
	2, 102, 30, 3, 2, 2, 2, 103, 105, 9, 6, 2, 2, 104, 103, 3, 2, 2, 2, 105,
	108, 3, 2, 2, 2, 106, 104, 3, 2, 2, 2, 106, 107, 3, 2, 2, 2, 107, 109,
	3, 2, 2, 2, 108, 106, 3, 2, 2, 2, 109, 113, 9, 7, 2, 2, 110, 112, 9, 8,
	2, 2, 111, 110, 3, 2, 2, 2, 112, 115, 3, 2, 2, 2, 113, 111, 3, 2, 2, 2,
	113, 114, 3, 2, 2, 2, 114, 32, 3, 2, 2, 2, 115, 113, 3, 2, 2, 2, 11, 2,
	36, 68, 75, 82, 93, 99, 106, 113, 3, 8, 2, 2,
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
	"", "", "'func'", "'turn'", "'*'", "'/'", "'+'", "'-'", "'='", "'('", "')'",
}

var lexerSymbolicNames = []string{
	"", "WS", "FUNC", "TURN", "MUL", "DIV", "ADD", "SUB", "ASSIGN", "LPAREN",
	"RPAREN", "INT", "FLOAT", "BOOL", "STRING", "IDENT",
}

var lexerRuleNames = []string{
	"WS", "FUNC", "TURN", "MUL", "DIV", "ADD", "SUB", "ASSIGN", "LPAREN", "RPAREN",
	"INT", "FLOAT", "BOOL", "STRING", "IDENT",
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
	TurnkeyLexerFUNC   = 2
	TurnkeyLexerTURN   = 3
	TurnkeyLexerMUL    = 4
	TurnkeyLexerDIV    = 5
	TurnkeyLexerADD    = 6
	TurnkeyLexerSUB    = 7
	TurnkeyLexerASSIGN = 8
	TurnkeyLexerLPAREN = 9
	TurnkeyLexerRPAREN = 10
	TurnkeyLexerINT    = 11
	TurnkeyLexerFLOAT  = 12
	TurnkeyLexerBOOL   = 13
	TurnkeyLexerSTRING = 14
	TurnkeyLexerIDENT  = 15
)