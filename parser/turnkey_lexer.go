// Code generated from TurnkeyLexer.g4 by ANTLR 4.8. DO NOT EDIT.

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 22, 138,
	8, 1, 8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6,
	4, 7, 9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12,
	9, 12, 4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9,
	17, 4, 18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 3, 2, 6, 2,
	46, 10, 2, 13, 2, 14, 2, 47, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4,
	3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8,
	3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 3, 11, 3, 12, 3, 12, 3, 13, 3, 13,
	3, 14, 3, 14, 3, 15, 3, 15, 7, 15, 84, 10, 15, 12, 15, 14, 15, 87, 11,
	15, 3, 16, 3, 16, 7, 16, 91, 10, 16, 12, 16, 14, 16, 94, 11, 16, 3, 16,
	3, 16, 6, 16, 98, 10, 16, 13, 16, 14, 16, 99, 3, 17, 3, 17, 3, 17, 3, 17,
	3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 5, 17, 111, 10, 17, 3, 18, 7, 18, 114,
	10, 18, 12, 18, 14, 18, 117, 11, 18, 3, 18, 3, 18, 7, 18, 121, 10, 18,
	12, 18, 14, 18, 124, 11, 18, 3, 19, 3, 19, 3, 19, 3, 19, 3, 20, 6, 20,
	131, 10, 20, 13, 20, 14, 20, 132, 3, 21, 3, 21, 3, 21, 3, 21, 2, 2, 22,
	4, 3, 6, 4, 8, 5, 10, 6, 12, 7, 14, 8, 16, 9, 18, 10, 20, 11, 22, 12, 24,
	13, 26, 14, 28, 15, 30, 16, 32, 17, 34, 18, 36, 19, 38, 20, 40, 21, 42,
	22, 4, 2, 3, 9, 5, 2, 11, 12, 15, 15, 34, 34, 3, 2, 51, 59, 3, 2, 50, 59,
	3, 2, 97, 97, 4, 2, 67, 92, 99, 124, 6, 2, 50, 59, 67, 92, 97, 97, 99,
	124, 4, 2, 36, 36, 94, 94, 2, 144, 2, 4, 3, 2, 2, 2, 2, 6, 3, 2, 2, 2,
	2, 8, 3, 2, 2, 2, 2, 10, 3, 2, 2, 2, 2, 12, 3, 2, 2, 2, 2, 14, 3, 2, 2,
	2, 2, 16, 3, 2, 2, 2, 2, 18, 3, 2, 2, 2, 2, 20, 3, 2, 2, 2, 2, 22, 3, 2,
	2, 2, 2, 24, 3, 2, 2, 2, 2, 26, 3, 2, 2, 2, 2, 28, 3, 2, 2, 2, 2, 30, 3,
	2, 2, 2, 2, 32, 3, 2, 2, 2, 2, 34, 3, 2, 2, 2, 2, 36, 3, 2, 2, 2, 2, 38,
	3, 2, 2, 2, 3, 40, 3, 2, 2, 2, 3, 42, 3, 2, 2, 2, 4, 45, 3, 2, 2, 2, 6,
	51, 3, 2, 2, 2, 8, 53, 3, 2, 2, 2, 10, 58, 3, 2, 2, 2, 12, 63, 3, 2, 2,
	2, 14, 65, 3, 2, 2, 2, 16, 67, 3, 2, 2, 2, 18, 69, 3, 2, 2, 2, 20, 71,
	3, 2, 2, 2, 22, 73, 3, 2, 2, 2, 24, 75, 3, 2, 2, 2, 26, 77, 3, 2, 2, 2,
	28, 79, 3, 2, 2, 2, 30, 81, 3, 2, 2, 2, 32, 88, 3, 2, 2, 2, 34, 110, 3,
	2, 2, 2, 36, 115, 3, 2, 2, 2, 38, 125, 3, 2, 2, 2, 40, 130, 3, 2, 2, 2,
	42, 134, 3, 2, 2, 2, 44, 46, 9, 2, 2, 2, 45, 44, 3, 2, 2, 2, 46, 47, 3,
	2, 2, 2, 47, 45, 3, 2, 2, 2, 47, 48, 3, 2, 2, 2, 48, 49, 3, 2, 2, 2, 49,
	50, 8, 2, 2, 2, 50, 5, 3, 2, 2, 2, 51, 52, 7, 46, 2, 2, 52, 7, 3, 2, 2,
	2, 53, 54, 7, 104, 2, 2, 54, 55, 7, 119, 2, 2, 55, 56, 7, 112, 2, 2, 56,
	57, 7, 101, 2, 2, 57, 9, 3, 2, 2, 2, 58, 59, 7, 118, 2, 2, 59, 60, 7, 119,
	2, 2, 60, 61, 7, 116, 2, 2, 61, 62, 7, 112, 2, 2, 62, 11, 3, 2, 2, 2, 63,
	64, 7, 44, 2, 2, 64, 13, 3, 2, 2, 2, 65, 66, 7, 49, 2, 2, 66, 15, 3, 2,
	2, 2, 67, 68, 7, 45, 2, 2, 68, 17, 3, 2, 2, 2, 69, 70, 7, 47, 2, 2, 70,
	19, 3, 2, 2, 2, 71, 72, 7, 63, 2, 2, 72, 21, 3, 2, 2, 2, 73, 74, 7, 42,
	2, 2, 74, 23, 3, 2, 2, 2, 75, 76, 7, 43, 2, 2, 76, 25, 3, 2, 2, 2, 77,
	78, 7, 125, 2, 2, 78, 27, 3, 2, 2, 2, 79, 80, 7, 127, 2, 2, 80, 29, 3,
	2, 2, 2, 81, 85, 9, 3, 2, 2, 82, 84, 9, 4, 2, 2, 83, 82, 3, 2, 2, 2, 84,
	87, 3, 2, 2, 2, 85, 83, 3, 2, 2, 2, 85, 86, 3, 2, 2, 2, 86, 31, 3, 2, 2,
	2, 87, 85, 3, 2, 2, 2, 88, 92, 9, 3, 2, 2, 89, 91, 9, 4, 2, 2, 90, 89,
	3, 2, 2, 2, 91, 94, 3, 2, 2, 2, 92, 90, 3, 2, 2, 2, 92, 93, 3, 2, 2, 2,
	93, 95, 3, 2, 2, 2, 94, 92, 3, 2, 2, 2, 95, 97, 7, 48, 2, 2, 96, 98, 9,
	4, 2, 2, 97, 96, 3, 2, 2, 2, 98, 99, 3, 2, 2, 2, 99, 97, 3, 2, 2, 2, 99,
	100, 3, 2, 2, 2, 100, 33, 3, 2, 2, 2, 101, 102, 7, 118, 2, 2, 102, 103,
	7, 116, 2, 2, 103, 104, 7, 119, 2, 2, 104, 111, 7, 103, 2, 2, 105, 106,
	7, 104, 2, 2, 106, 107, 7, 99, 2, 2, 107, 108, 7, 110, 2, 2, 108, 109,
	7, 117, 2, 2, 109, 111, 7, 103, 2, 2, 110, 101, 3, 2, 2, 2, 110, 105, 3,
	2, 2, 2, 111, 35, 3, 2, 2, 2, 112, 114, 9, 5, 2, 2, 113, 112, 3, 2, 2,
	2, 114, 117, 3, 2, 2, 2, 115, 113, 3, 2, 2, 2, 115, 116, 3, 2, 2, 2, 116,
	118, 3, 2, 2, 2, 117, 115, 3, 2, 2, 2, 118, 122, 9, 6, 2, 2, 119, 121,
	9, 7, 2, 2, 120, 119, 3, 2, 2, 2, 121, 124, 3, 2, 2, 2, 122, 120, 3, 2,
	2, 2, 122, 123, 3, 2, 2, 2, 123, 37, 3, 2, 2, 2, 124, 122, 3, 2, 2, 2,
	125, 126, 7, 36, 2, 2, 126, 127, 3, 2, 2, 2, 127, 128, 8, 19, 3, 2, 128,
	39, 3, 2, 2, 2, 129, 131, 10, 8, 2, 2, 130, 129, 3, 2, 2, 2, 131, 132,
	3, 2, 2, 2, 132, 130, 3, 2, 2, 2, 132, 133, 3, 2, 2, 2, 133, 41, 3, 2,
	2, 2, 134, 135, 7, 36, 2, 2, 135, 136, 3, 2, 2, 2, 136, 137, 8, 21, 4,
	2, 137, 43, 3, 2, 2, 2, 12, 2, 3, 47, 85, 92, 99, 110, 115, 122, 132, 5,
	8, 2, 2, 7, 3, 2, 6, 2, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE", "IN_STRING",
}

var lexerLiteralNames = []string{
	"", "", "','", "'func'", "'turn'", "'*'", "'/'", "'+'", "'-'", "'='", "'('",
	"')'", "'{'", "'}'",
}

var lexerSymbolicNames = []string{
	"", "WS", "COMMA", "FUNC", "TURN", "MUL", "DIV", "ADD", "SUB", "ASSIGN",
	"LPAREN", "RPAREN", "LBRACE", "RBRACE", "INT", "FLOAT", "BOOL", "IDENT",
	"OPEN_STRING", "STRING", "CLOSE_STRING",
}

var lexerRuleNames = []string{
	"WS", "COMMA", "FUNC", "TURN", "MUL", "DIV", "ADD", "SUB", "ASSIGN", "LPAREN",
	"RPAREN", "LBRACE", "RBRACE", "INT", "FLOAT", "BOOL", "IDENT", "OPEN_STRING",
	"STRING", "CLOSE_STRING",
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
	l.GrammarFileName = "TurnkeyLexer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// TurnkeyLexer tokens.
const (
	TurnkeyLexerWS           = 1
	TurnkeyLexerCOMMA        = 2
	TurnkeyLexerFUNC         = 3
	TurnkeyLexerTURN         = 4
	TurnkeyLexerMUL          = 5
	TurnkeyLexerDIV          = 6
	TurnkeyLexerADD          = 7
	TurnkeyLexerSUB          = 8
	TurnkeyLexerASSIGN       = 9
	TurnkeyLexerLPAREN       = 10
	TurnkeyLexerRPAREN       = 11
	TurnkeyLexerLBRACE       = 12
	TurnkeyLexerRBRACE       = 13
	TurnkeyLexerINT          = 14
	TurnkeyLexerFLOAT        = 15
	TurnkeyLexerBOOL         = 16
	TurnkeyLexerIDENT        = 17
	TurnkeyLexerOPEN_STRING  = 18
	TurnkeyLexerSTRING       = 19
	TurnkeyLexerCLOSE_STRING = 20
)

// TurnkeyLexerIN_STRING is the TurnkeyLexer mode.
const TurnkeyLexerIN_STRING = 1
