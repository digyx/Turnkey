// Code generated from TurnkeyParser.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // TurnkeyParser

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 29, 146,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 5, 3, 34, 10, 3, 3,
	4, 3, 4, 7, 4, 38, 10, 4, 12, 4, 14, 4, 41, 11, 4, 3, 4, 3, 4, 3, 5, 3,
	5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 7, 6, 54, 10, 6, 12, 6, 14,
	6, 57, 11, 6, 3, 6, 5, 6, 60, 10, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7,
	3, 7, 5, 7, 69, 10, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 5, 8, 77, 10,
	8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 7, 8, 85, 10, 8, 12, 8, 14, 8, 88,
	11, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 5, 9, 96, 10, 9, 3, 9, 3, 9,
	3, 9, 3, 9, 3, 9, 3, 9, 7, 9, 104, 10, 9, 12, 9, 14, 9, 107, 11, 9, 3,
	10, 3, 10, 3, 10, 3, 10, 5, 10, 113, 10, 10, 3, 10, 3, 10, 3, 10, 7, 10,
	118, 10, 10, 12, 10, 14, 10, 121, 11, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3,
	11, 3, 11, 3, 11, 3, 11, 7, 11, 131, 10, 11, 12, 11, 14, 11, 134, 11, 11,
	3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 3,
	14, 2, 6, 14, 16, 18, 20, 15, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24,
	26, 2, 5, 3, 2, 7, 8, 3, 2, 9, 10, 3, 2, 12, 13, 2, 151, 2, 28, 3, 2, 2,
	2, 4, 33, 3, 2, 2, 2, 6, 35, 3, 2, 2, 2, 8, 44, 3, 2, 2, 2, 10, 55, 3,
	2, 2, 2, 12, 68, 3, 2, 2, 2, 14, 76, 3, 2, 2, 2, 16, 95, 3, 2, 2, 2, 18,
	112, 3, 2, 2, 2, 20, 122, 3, 2, 2, 2, 22, 135, 3, 2, 2, 2, 24, 137, 3,
	2, 2, 2, 26, 142, 3, 2, 2, 2, 28, 29, 5, 4, 3, 2, 29, 30, 7, 2, 2, 3, 30,
	3, 3, 2, 2, 2, 31, 34, 5, 12, 7, 2, 32, 34, 5, 8, 5, 2, 33, 31, 3, 2, 2,
	2, 33, 32, 3, 2, 2, 2, 34, 5, 3, 2, 2, 2, 35, 39, 7, 21, 2, 2, 36, 38,
	5, 12, 7, 2, 37, 36, 3, 2, 2, 2, 38, 41, 3, 2, 2, 2, 39, 37, 3, 2, 2, 2,
	39, 40, 3, 2, 2, 2, 40, 42, 3, 2, 2, 2, 41, 39, 3, 2, 2, 2, 42, 43, 7,
	22, 2, 2, 43, 7, 3, 2, 2, 2, 44, 45, 7, 5, 2, 2, 45, 46, 7, 26, 2, 2, 46,
	47, 7, 19, 2, 2, 47, 48, 5, 10, 6, 2, 48, 49, 7, 20, 2, 2, 49, 50, 5, 6,
	4, 2, 50, 9, 3, 2, 2, 2, 51, 52, 7, 26, 2, 2, 52, 54, 7, 4, 2, 2, 53, 51,
	3, 2, 2, 2, 54, 57, 3, 2, 2, 2, 55, 53, 3, 2, 2, 2, 55, 56, 3, 2, 2, 2,
	56, 59, 3, 2, 2, 2, 57, 55, 3, 2, 2, 2, 58, 60, 7, 26, 2, 2, 59, 58, 3,
	2, 2, 2, 59, 60, 3, 2, 2, 2, 60, 11, 3, 2, 2, 2, 61, 69, 5, 14, 8, 2, 62,
	69, 5, 16, 9, 2, 63, 69, 5, 18, 10, 2, 64, 69, 5, 20, 11, 2, 65, 69, 5,
	22, 12, 2, 66, 69, 5, 24, 13, 2, 67, 69, 5, 26, 14, 2, 68, 61, 3, 2, 2,
	2, 68, 62, 3, 2, 2, 2, 68, 63, 3, 2, 2, 2, 68, 64, 3, 2, 2, 2, 68, 65,
	3, 2, 2, 2, 68, 66, 3, 2, 2, 2, 68, 67, 3, 2, 2, 2, 69, 13, 3, 2, 2, 2,
	70, 71, 8, 8, 1, 2, 71, 72, 7, 19, 2, 2, 72, 73, 5, 14, 8, 2, 73, 74, 7,
	20, 2, 2, 74, 77, 3, 2, 2, 2, 75, 77, 7, 23, 2, 2, 76, 70, 3, 2, 2, 2,
	76, 75, 3, 2, 2, 2, 77, 86, 3, 2, 2, 2, 78, 79, 12, 5, 2, 2, 79, 80, 9,
	2, 2, 2, 80, 85, 5, 14, 8, 6, 81, 82, 12, 4, 2, 2, 82, 83, 9, 3, 2, 2,
	83, 85, 5, 14, 8, 5, 84, 78, 3, 2, 2, 2, 84, 81, 3, 2, 2, 2, 85, 88, 3,
	2, 2, 2, 86, 84, 3, 2, 2, 2, 86, 87, 3, 2, 2, 2, 87, 15, 3, 2, 2, 2, 88,
	86, 3, 2, 2, 2, 89, 90, 8, 9, 1, 2, 90, 91, 7, 19, 2, 2, 91, 92, 5, 16,
	9, 2, 92, 93, 7, 20, 2, 2, 93, 96, 3, 2, 2, 2, 94, 96, 7, 24, 2, 2, 95,
	89, 3, 2, 2, 2, 95, 94, 3, 2, 2, 2, 96, 105, 3, 2, 2, 2, 97, 98, 12, 5,
	2, 2, 98, 99, 9, 2, 2, 2, 99, 104, 5, 16, 9, 6, 100, 101, 12, 4, 2, 2,
	101, 102, 9, 3, 2, 2, 102, 104, 5, 16, 9, 5, 103, 97, 3, 2, 2, 2, 103,
	100, 3, 2, 2, 2, 104, 107, 3, 2, 2, 2, 105, 103, 3, 2, 2, 2, 105, 106,
	3, 2, 2, 2, 106, 17, 3, 2, 2, 2, 107, 105, 3, 2, 2, 2, 108, 109, 8, 10,
	1, 2, 109, 110, 7, 11, 2, 2, 110, 113, 5, 18, 10, 5, 111, 113, 7, 25, 2,
	2, 112, 108, 3, 2, 2, 2, 112, 111, 3, 2, 2, 2, 113, 119, 3, 2, 2, 2, 114,
	115, 12, 4, 2, 2, 115, 116, 9, 4, 2, 2, 116, 118, 5, 18, 10, 5, 117, 114,
	3, 2, 2, 2, 118, 121, 3, 2, 2, 2, 119, 117, 3, 2, 2, 2, 119, 120, 3, 2,
	2, 2, 120, 19, 3, 2, 2, 2, 121, 119, 3, 2, 2, 2, 122, 123, 8, 11, 1, 2,
	123, 124, 7, 27, 2, 2, 124, 125, 7, 28, 2, 2, 125, 126, 7, 29, 2, 2, 126,
	132, 3, 2, 2, 2, 127, 128, 12, 4, 2, 2, 128, 129, 7, 9, 2, 2, 129, 131,
	5, 20, 11, 5, 130, 127, 3, 2, 2, 2, 131, 134, 3, 2, 2, 2, 132, 130, 3,
	2, 2, 2, 132, 133, 3, 2, 2, 2, 133, 21, 3, 2, 2, 2, 134, 132, 3, 2, 2,
	2, 135, 136, 7, 26, 2, 2, 136, 23, 3, 2, 2, 2, 137, 138, 7, 26, 2, 2, 138,
	139, 7, 19, 2, 2, 139, 140, 5, 10, 6, 2, 140, 141, 7, 20, 2, 2, 141, 25,
	3, 2, 2, 2, 142, 143, 7, 6, 2, 2, 143, 144, 5, 24, 13, 2, 144, 27, 3, 2,
	2, 2, 16, 33, 39, 55, 59, 68, 76, 84, 86, 95, 103, 105, 112, 119, 132,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "", "','", "'func'", "'turn'", "'*'", "'/'", "'+'", "'-'", "'!'", "'=='",
	"'!='", "'>'", "'>='", "'<'", "'<='", "'='", "'('", "')'", "'{'", "'}'",
}
var symbolicNames = []string{
	"", "WS", "COMMA", "FUNC", "TURN", "MUL", "DIV", "ADD", "SUB", "BANG",
	"EQ", "NOT_EQ", "GT", "GT_EQ", "LT", "LT_EQ", "ASSIGN", "LPAREN", "RPAREN",
	"LBRACE", "RBRACE", "INT", "FLOAT", "BOOL", "IDENT", "OPEN_STRING", "STRING",
	"CLOSE_STRING",
}

var ruleNames = []string{
	"start", "statement", "block", "func_def", "parameters", "expression",
	"int_expression", "float_expression", "bool_expression", "string_expression",
	"ident_expression", "call_expression", "turn_expression",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type TurnkeyParser struct {
	*antlr.BaseParser
}

func NewTurnkeyParser(input antlr.TokenStream) *TurnkeyParser {
	this := new(TurnkeyParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "TurnkeyParser.g4"

	return this
}

// TurnkeyParser tokens.
const (
	TurnkeyParserEOF          = antlr.TokenEOF
	TurnkeyParserWS           = 1
	TurnkeyParserCOMMA        = 2
	TurnkeyParserFUNC         = 3
	TurnkeyParserTURN         = 4
	TurnkeyParserMUL          = 5
	TurnkeyParserDIV          = 6
	TurnkeyParserADD          = 7
	TurnkeyParserSUB          = 8
	TurnkeyParserBANG         = 9
	TurnkeyParserEQ           = 10
	TurnkeyParserNOT_EQ       = 11
	TurnkeyParserGT           = 12
	TurnkeyParserGT_EQ        = 13
	TurnkeyParserLT           = 14
	TurnkeyParserLT_EQ        = 15
	TurnkeyParserASSIGN       = 16
	TurnkeyParserLPAREN       = 17
	TurnkeyParserRPAREN       = 18
	TurnkeyParserLBRACE       = 19
	TurnkeyParserRBRACE       = 20
	TurnkeyParserINT          = 21
	TurnkeyParserFLOAT        = 22
	TurnkeyParserBOOL         = 23
	TurnkeyParserIDENT        = 24
	TurnkeyParserOPEN_STRING  = 25
	TurnkeyParserSTRING       = 26
	TurnkeyParserCLOSE_STRING = 27
)

// TurnkeyParser rules.
const (
	TurnkeyParserRULE_start             = 0
	TurnkeyParserRULE_statement         = 1
	TurnkeyParserRULE_block             = 2
	TurnkeyParserRULE_func_def          = 3
	TurnkeyParserRULE_parameters        = 4
	TurnkeyParserRULE_expression        = 5
	TurnkeyParserRULE_int_expression    = 6
	TurnkeyParserRULE_float_expression  = 7
	TurnkeyParserRULE_bool_expression   = 8
	TurnkeyParserRULE_string_expression = 9
	TurnkeyParserRULE_ident_expression  = 10
	TurnkeyParserRULE_call_expression   = 11
	TurnkeyParserRULE_turn_expression   = 12
)

// IStartContext is an interface to support dynamic dispatch.
type IStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStartContext differentiates from other interfaces.
	IsStartContext()
}

type StartContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStartContext() *StartContext {
	var p = new(StartContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TurnkeyParserRULE_start
	return p
}

func (*StartContext) IsStartContext() {}

func NewStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartContext {
	var p = new(StartContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TurnkeyParserRULE_start

	return p
}

func (s *StartContext) GetParser() antlr.Parser { return s.parser }

func (s *StartContext) Statement() IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *StartContext) EOF() antlr.TerminalNode {
	return s.GetToken(TurnkeyParserEOF, 0)
}

func (s *StartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StartContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TurnkeyParserListener); ok {
		listenerT.EnterStart(s)
	}
}

func (s *StartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TurnkeyParserListener); ok {
		listenerT.ExitStart(s)
	}
}

func (p *TurnkeyParser) Start() (localctx IStartContext) {
	localctx = NewStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, TurnkeyParserRULE_start)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(26)
		p.Statement()
	}
	{
		p.SetState(27)
		p.Match(TurnkeyParserEOF)
	}

	return localctx
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TurnkeyParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TurnkeyParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *StatementContext) Func_def() IFunc_defContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunc_defContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunc_defContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TurnkeyParserListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TurnkeyParserListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (p *TurnkeyParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, TurnkeyParserRULE_statement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(31)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case TurnkeyParserTURN, TurnkeyParserBANG, TurnkeyParserLPAREN, TurnkeyParserINT, TurnkeyParserFLOAT, TurnkeyParserBOOL, TurnkeyParserIDENT, TurnkeyParserOPEN_STRING:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(29)
			p.Expression()
		}

	case TurnkeyParserFUNC:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(30)
			p.Func_def()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TurnkeyParserRULE_block
	return p
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TurnkeyParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(TurnkeyParserLBRACE, 0)
}

func (s *BlockContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(TurnkeyParserRBRACE, 0)
}

func (s *BlockContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *BlockContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TurnkeyParserListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TurnkeyParserListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (p *TurnkeyParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, TurnkeyParserRULE_block)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(33)
		p.Match(TurnkeyParserLBRACE)
	}
	p.SetState(37)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<TurnkeyParserTURN)|(1<<TurnkeyParserBANG)|(1<<TurnkeyParserLPAREN)|(1<<TurnkeyParserINT)|(1<<TurnkeyParserFLOAT)|(1<<TurnkeyParserBOOL)|(1<<TurnkeyParserIDENT)|(1<<TurnkeyParserOPEN_STRING))) != 0 {
		{
			p.SetState(34)
			p.Expression()
		}

		p.SetState(39)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(40)
		p.Match(TurnkeyParserRBRACE)
	}

	return localctx
}

// IFunc_defContext is an interface to support dynamic dispatch.
type IFunc_defContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunc_defContext differentiates from other interfaces.
	IsFunc_defContext()
}

type Func_defContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunc_defContext() *Func_defContext {
	var p = new(Func_defContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TurnkeyParserRULE_func_def
	return p
}

func (*Func_defContext) IsFunc_defContext() {}

func NewFunc_defContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Func_defContext {
	var p = new(Func_defContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TurnkeyParserRULE_func_def

	return p
}

func (s *Func_defContext) GetParser() antlr.Parser { return s.parser }

func (s *Func_defContext) FUNC() antlr.TerminalNode {
	return s.GetToken(TurnkeyParserFUNC, 0)
}

func (s *Func_defContext) IDENT() antlr.TerminalNode {
	return s.GetToken(TurnkeyParserIDENT, 0)
}

func (s *Func_defContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(TurnkeyParserLPAREN, 0)
}

func (s *Func_defContext) Parameters() IParametersContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParametersContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParametersContext)
}

func (s *Func_defContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(TurnkeyParserRPAREN, 0)
}

func (s *Func_defContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *Func_defContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Func_defContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Func_defContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TurnkeyParserListener); ok {
		listenerT.EnterFunc_def(s)
	}
}

func (s *Func_defContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TurnkeyParserListener); ok {
		listenerT.ExitFunc_def(s)
	}
}

func (p *TurnkeyParser) Func_def() (localctx IFunc_defContext) {
	localctx = NewFunc_defContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, TurnkeyParserRULE_func_def)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(42)
		p.Match(TurnkeyParserFUNC)
	}
	{
		p.SetState(43)
		p.Match(TurnkeyParserIDENT)
	}
	{
		p.SetState(44)
		p.Match(TurnkeyParserLPAREN)
	}
	{
		p.SetState(45)
		p.Parameters()
	}
	{
		p.SetState(46)
		p.Match(TurnkeyParserRPAREN)
	}
	{
		p.SetState(47)
		p.Block()
	}

	return localctx
}

// IParametersContext is an interface to support dynamic dispatch.
type IParametersContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParametersContext differentiates from other interfaces.
	IsParametersContext()
}

type ParametersContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParametersContext() *ParametersContext {
	var p = new(ParametersContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TurnkeyParserRULE_parameters
	return p
}

func (*ParametersContext) IsParametersContext() {}

func NewParametersContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParametersContext {
	var p = new(ParametersContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TurnkeyParserRULE_parameters

	return p
}

func (s *ParametersContext) GetParser() antlr.Parser { return s.parser }

func (s *ParametersContext) AllIDENT() []antlr.TerminalNode {
	return s.GetTokens(TurnkeyParserIDENT)
}

func (s *ParametersContext) IDENT(i int) antlr.TerminalNode {
	return s.GetToken(TurnkeyParserIDENT, i)
}

func (s *ParametersContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(TurnkeyParserCOMMA)
}

func (s *ParametersContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(TurnkeyParserCOMMA, i)
}

func (s *ParametersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParametersContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParametersContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TurnkeyParserListener); ok {
		listenerT.EnterParameters(s)
	}
}

func (s *ParametersContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TurnkeyParserListener); ok {
		listenerT.ExitParameters(s)
	}
}

func (p *TurnkeyParser) Parameters() (localctx IParametersContext) {
	localctx = NewParametersContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, TurnkeyParserRULE_parameters)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(53)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(49)
				p.Match(TurnkeyParserIDENT)
			}
			{
				p.SetState(50)
				p.Match(TurnkeyParserCOMMA)
			}

		}
		p.SetState(55)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())
	}
	p.SetState(57)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == TurnkeyParserIDENT {
		{
			p.SetState(56)
			p.Match(TurnkeyParserIDENT)
		}

	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TurnkeyParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TurnkeyParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Int_expression() IInt_expressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInt_expressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInt_expressionContext)
}

func (s *ExpressionContext) Float_expression() IFloat_expressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFloat_expressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFloat_expressionContext)
}

func (s *ExpressionContext) Bool_expression() IBool_expressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBool_expressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBool_expressionContext)
}

func (s *ExpressionContext) String_expression() IString_expressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IString_expressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IString_expressionContext)
}

func (s *ExpressionContext) Ident_expression() IIdent_expressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdent_expressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdent_expressionContext)
}

func (s *ExpressionContext) Call_expression() ICall_expressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICall_expressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICall_expressionContext)
}

func (s *ExpressionContext) Turn_expression() ITurn_expressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITurn_expressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITurn_expressionContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TurnkeyParserListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TurnkeyParserListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *TurnkeyParser) Expression() (localctx IExpressionContext) {
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, TurnkeyParserRULE_expression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(66)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(59)
			p.int_expression(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(60)
			p.float_expression(0)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(61)
			p.bool_expression(0)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(62)
			p.string_expression(0)
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(63)
			p.Ident_expression()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(64)
			p.Call_expression()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(65)
			p.Turn_expression()
		}

	}

	return localctx
}

// IInt_expressionContext is an interface to support dynamic dispatch.
type IInt_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// IsInt_expressionContext differentiates from other interfaces.
	IsInt_expressionContext()
}

type Int_expressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	op     antlr.Token
}

func NewEmptyInt_expressionContext() *Int_expressionContext {
	var p = new(Int_expressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TurnkeyParserRULE_int_expression
	return p
}

func (*Int_expressionContext) IsInt_expressionContext() {}

func NewInt_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Int_expressionContext {
	var p = new(Int_expressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TurnkeyParserRULE_int_expression

	return p
}

func (s *Int_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Int_expressionContext) GetOp() antlr.Token { return s.op }

func (s *Int_expressionContext) SetOp(v antlr.Token) { s.op = v }

func (s *Int_expressionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(TurnkeyParserLPAREN, 0)
}

func (s *Int_expressionContext) AllInt_expression() []IInt_expressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IInt_expressionContext)(nil)).Elem())
	var tst = make([]IInt_expressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IInt_expressionContext)
		}
	}

	return tst
}

func (s *Int_expressionContext) Int_expression(i int) IInt_expressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInt_expressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IInt_expressionContext)
}

func (s *Int_expressionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(TurnkeyParserRPAREN, 0)
}

func (s *Int_expressionContext) INT() antlr.TerminalNode {
	return s.GetToken(TurnkeyParserINT, 0)
}

func (s *Int_expressionContext) MUL() antlr.TerminalNode {
	return s.GetToken(TurnkeyParserMUL, 0)
}

func (s *Int_expressionContext) DIV() antlr.TerminalNode {
	return s.GetToken(TurnkeyParserDIV, 0)
}

func (s *Int_expressionContext) ADD() antlr.TerminalNode {
	return s.GetToken(TurnkeyParserADD, 0)
}

func (s *Int_expressionContext) SUB() antlr.TerminalNode {
	return s.GetToken(TurnkeyParserSUB, 0)
}

func (s *Int_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Int_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Int_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TurnkeyParserListener); ok {
		listenerT.EnterInt_expression(s)
	}
}

func (s *Int_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TurnkeyParserListener); ok {
		listenerT.ExitInt_expression(s)
	}
}

func (p *TurnkeyParser) Int_expression() (localctx IInt_expressionContext) {
	return p.int_expression(0)
}

func (p *TurnkeyParser) int_expression(_p int) (localctx IInt_expressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewInt_expressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IInt_expressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 12
	p.EnterRecursionRule(localctx, 12, TurnkeyParserRULE_int_expression, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(74)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case TurnkeyParserLPAREN:
		{
			p.SetState(69)
			p.Match(TurnkeyParserLPAREN)
		}
		{
			p.SetState(70)
			p.int_expression(0)
		}
		{
			p.SetState(71)
			p.Match(TurnkeyParserRPAREN)
		}

	case TurnkeyParserINT:
		{
			p.SetState(73)
			p.Match(TurnkeyParserINT)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(84)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(82)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) {
			case 1:
				localctx = NewInt_expressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, TurnkeyParserRULE_int_expression)
				p.SetState(76)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(77)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*Int_expressionContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == TurnkeyParserMUL || _la == TurnkeyParserDIV) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*Int_expressionContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(78)
					p.int_expression(4)
				}

			case 2:
				localctx = NewInt_expressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, TurnkeyParserRULE_int_expression)
				p.SetState(79)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(80)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*Int_expressionContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == TurnkeyParserADD || _la == TurnkeyParserSUB) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*Int_expressionContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(81)
					p.int_expression(3)
				}

			}

		}
		p.SetState(86)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext())
	}

	return localctx
}

// IFloat_expressionContext is an interface to support dynamic dispatch.
type IFloat_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// IsFloat_expressionContext differentiates from other interfaces.
	IsFloat_expressionContext()
}

type Float_expressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	op     antlr.Token
}

func NewEmptyFloat_expressionContext() *Float_expressionContext {
	var p = new(Float_expressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TurnkeyParserRULE_float_expression
	return p
}

func (*Float_expressionContext) IsFloat_expressionContext() {}

func NewFloat_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Float_expressionContext {
	var p = new(Float_expressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TurnkeyParserRULE_float_expression

	return p
}

func (s *Float_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Float_expressionContext) GetOp() antlr.Token { return s.op }

func (s *Float_expressionContext) SetOp(v antlr.Token) { s.op = v }

func (s *Float_expressionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(TurnkeyParserLPAREN, 0)
}

func (s *Float_expressionContext) AllFloat_expression() []IFloat_expressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFloat_expressionContext)(nil)).Elem())
	var tst = make([]IFloat_expressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFloat_expressionContext)
		}
	}

	return tst
}

func (s *Float_expressionContext) Float_expression(i int) IFloat_expressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFloat_expressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFloat_expressionContext)
}

func (s *Float_expressionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(TurnkeyParserRPAREN, 0)
}

func (s *Float_expressionContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(TurnkeyParserFLOAT, 0)
}

func (s *Float_expressionContext) MUL() antlr.TerminalNode {
	return s.GetToken(TurnkeyParserMUL, 0)
}

func (s *Float_expressionContext) DIV() antlr.TerminalNode {
	return s.GetToken(TurnkeyParserDIV, 0)
}

func (s *Float_expressionContext) ADD() antlr.TerminalNode {
	return s.GetToken(TurnkeyParserADD, 0)
}

func (s *Float_expressionContext) SUB() antlr.TerminalNode {
	return s.GetToken(TurnkeyParserSUB, 0)
}

func (s *Float_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Float_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Float_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TurnkeyParserListener); ok {
		listenerT.EnterFloat_expression(s)
	}
}

func (s *Float_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TurnkeyParserListener); ok {
		listenerT.ExitFloat_expression(s)
	}
}

func (p *TurnkeyParser) Float_expression() (localctx IFloat_expressionContext) {
	return p.float_expression(0)
}

func (p *TurnkeyParser) float_expression(_p int) (localctx IFloat_expressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewFloat_expressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IFloat_expressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 14
	p.EnterRecursionRule(localctx, 14, TurnkeyParserRULE_float_expression, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(93)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case TurnkeyParserLPAREN:
		{
			p.SetState(88)
			p.Match(TurnkeyParserLPAREN)
		}
		{
			p.SetState(89)
			p.float_expression(0)
		}
		{
			p.SetState(90)
			p.Match(TurnkeyParserRPAREN)
		}

	case TurnkeyParserFLOAT:
		{
			p.SetState(92)
			p.Match(TurnkeyParserFLOAT)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(103)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(101)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) {
			case 1:
				localctx = NewFloat_expressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, TurnkeyParserRULE_float_expression)
				p.SetState(95)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(96)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*Float_expressionContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == TurnkeyParserMUL || _la == TurnkeyParserDIV) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*Float_expressionContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(97)
					p.float_expression(4)
				}

			case 2:
				localctx = NewFloat_expressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, TurnkeyParserRULE_float_expression)
				p.SetState(98)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(99)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*Float_expressionContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == TurnkeyParserADD || _la == TurnkeyParserSUB) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*Float_expressionContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(100)
					p.float_expression(3)
				}

			}

		}
		p.SetState(105)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext())
	}

	return localctx
}

// IBool_expressionContext is an interface to support dynamic dispatch.
type IBool_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// IsBool_expressionContext differentiates from other interfaces.
	IsBool_expressionContext()
}

type Bool_expressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	op     antlr.Token
}

func NewEmptyBool_expressionContext() *Bool_expressionContext {
	var p = new(Bool_expressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TurnkeyParserRULE_bool_expression
	return p
}

func (*Bool_expressionContext) IsBool_expressionContext() {}

func NewBool_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Bool_expressionContext {
	var p = new(Bool_expressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TurnkeyParserRULE_bool_expression

	return p
}

func (s *Bool_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Bool_expressionContext) GetOp() antlr.Token { return s.op }

func (s *Bool_expressionContext) SetOp(v antlr.Token) { s.op = v }

func (s *Bool_expressionContext) BANG() antlr.TerminalNode {
	return s.GetToken(TurnkeyParserBANG, 0)
}

func (s *Bool_expressionContext) AllBool_expression() []IBool_expressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBool_expressionContext)(nil)).Elem())
	var tst = make([]IBool_expressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBool_expressionContext)
		}
	}

	return tst
}

func (s *Bool_expressionContext) Bool_expression(i int) IBool_expressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBool_expressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBool_expressionContext)
}

func (s *Bool_expressionContext) BOOL() antlr.TerminalNode {
	return s.GetToken(TurnkeyParserBOOL, 0)
}

func (s *Bool_expressionContext) EQ() antlr.TerminalNode {
	return s.GetToken(TurnkeyParserEQ, 0)
}

func (s *Bool_expressionContext) NOT_EQ() antlr.TerminalNode {
	return s.GetToken(TurnkeyParserNOT_EQ, 0)
}

func (s *Bool_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Bool_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Bool_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TurnkeyParserListener); ok {
		listenerT.EnterBool_expression(s)
	}
}

func (s *Bool_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TurnkeyParserListener); ok {
		listenerT.ExitBool_expression(s)
	}
}

func (p *TurnkeyParser) Bool_expression() (localctx IBool_expressionContext) {
	return p.bool_expression(0)
}

func (p *TurnkeyParser) bool_expression(_p int) (localctx IBool_expressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewBool_expressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IBool_expressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 16
	p.EnterRecursionRule(localctx, 16, TurnkeyParserRULE_bool_expression, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(110)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case TurnkeyParserBANG:
		{
			p.SetState(107)
			p.Match(TurnkeyParserBANG)
		}
		{
			p.SetState(108)
			p.bool_expression(3)
		}

	case TurnkeyParserBOOL:
		{
			p.SetState(109)
			p.Match(TurnkeyParserBOOL)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(117)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewBool_expressionContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, TurnkeyParserRULE_bool_expression)
			p.SetState(112)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(113)

				var _lt = p.GetTokenStream().LT(1)

				localctx.(*Bool_expressionContext).op = _lt

				_la = p.GetTokenStream().LA(1)

				if !(_la == TurnkeyParserEQ || _la == TurnkeyParserNOT_EQ) {
					var _ri = p.GetErrorHandler().RecoverInline(p)

					localctx.(*Bool_expressionContext).op = _ri
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(114)
				p.bool_expression(3)
			}

		}
		p.SetState(119)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext())
	}

	return localctx
}

// IString_expressionContext is an interface to support dynamic dispatch.
type IString_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// IsString_expressionContext differentiates from other interfaces.
	IsString_expressionContext()
}

type String_expressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	op     antlr.Token
}

func NewEmptyString_expressionContext() *String_expressionContext {
	var p = new(String_expressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TurnkeyParserRULE_string_expression
	return p
}

func (*String_expressionContext) IsString_expressionContext() {}

func NewString_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *String_expressionContext {
	var p = new(String_expressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TurnkeyParserRULE_string_expression

	return p
}

func (s *String_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *String_expressionContext) GetOp() antlr.Token { return s.op }

func (s *String_expressionContext) SetOp(v antlr.Token) { s.op = v }

func (s *String_expressionContext) OPEN_STRING() antlr.TerminalNode {
	return s.GetToken(TurnkeyParserOPEN_STRING, 0)
}

func (s *String_expressionContext) STRING() antlr.TerminalNode {
	return s.GetToken(TurnkeyParserSTRING, 0)
}

func (s *String_expressionContext) CLOSE_STRING() antlr.TerminalNode {
	return s.GetToken(TurnkeyParserCLOSE_STRING, 0)
}

func (s *String_expressionContext) AllString_expression() []IString_expressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IString_expressionContext)(nil)).Elem())
	var tst = make([]IString_expressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IString_expressionContext)
		}
	}

	return tst
}

func (s *String_expressionContext) String_expression(i int) IString_expressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IString_expressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IString_expressionContext)
}

func (s *String_expressionContext) ADD() antlr.TerminalNode {
	return s.GetToken(TurnkeyParserADD, 0)
}

func (s *String_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *String_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *String_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TurnkeyParserListener); ok {
		listenerT.EnterString_expression(s)
	}
}

func (s *String_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TurnkeyParserListener); ok {
		listenerT.ExitString_expression(s)
	}
}

func (p *TurnkeyParser) String_expression() (localctx IString_expressionContext) {
	return p.string_expression(0)
}

func (p *TurnkeyParser) string_expression(_p int) (localctx IString_expressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewString_expressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IString_expressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 18
	p.EnterRecursionRule(localctx, 18, TurnkeyParserRULE_string_expression, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(121)
		p.Match(TurnkeyParserOPEN_STRING)
	}
	{
		p.SetState(122)
		p.Match(TurnkeyParserSTRING)
	}
	{
		p.SetState(123)
		p.Match(TurnkeyParserCLOSE_STRING)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(130)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewString_expressionContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, TurnkeyParserRULE_string_expression)
			p.SetState(125)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(126)

				var _m = p.Match(TurnkeyParserADD)

				localctx.(*String_expressionContext).op = _m
			}
			{
				p.SetState(127)
				p.string_expression(3)
			}

		}
		p.SetState(132)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext())
	}

	return localctx
}

// IIdent_expressionContext is an interface to support dynamic dispatch.
type IIdent_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIdent_expressionContext differentiates from other interfaces.
	IsIdent_expressionContext()
}

type Ident_expressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdent_expressionContext() *Ident_expressionContext {
	var p = new(Ident_expressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TurnkeyParserRULE_ident_expression
	return p
}

func (*Ident_expressionContext) IsIdent_expressionContext() {}

func NewIdent_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Ident_expressionContext {
	var p = new(Ident_expressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TurnkeyParserRULE_ident_expression

	return p
}

func (s *Ident_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Ident_expressionContext) IDENT() antlr.TerminalNode {
	return s.GetToken(TurnkeyParserIDENT, 0)
}

func (s *Ident_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Ident_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Ident_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TurnkeyParserListener); ok {
		listenerT.EnterIdent_expression(s)
	}
}

func (s *Ident_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TurnkeyParserListener); ok {
		listenerT.ExitIdent_expression(s)
	}
}

func (p *TurnkeyParser) Ident_expression() (localctx IIdent_expressionContext) {
	localctx = NewIdent_expressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, TurnkeyParserRULE_ident_expression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(133)
		p.Match(TurnkeyParserIDENT)
	}

	return localctx
}

// ICall_expressionContext is an interface to support dynamic dispatch.
type ICall_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCall_expressionContext differentiates from other interfaces.
	IsCall_expressionContext()
}

type Call_expressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCall_expressionContext() *Call_expressionContext {
	var p = new(Call_expressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TurnkeyParserRULE_call_expression
	return p
}

func (*Call_expressionContext) IsCall_expressionContext() {}

func NewCall_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Call_expressionContext {
	var p = new(Call_expressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TurnkeyParserRULE_call_expression

	return p
}

func (s *Call_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Call_expressionContext) IDENT() antlr.TerminalNode {
	return s.GetToken(TurnkeyParserIDENT, 0)
}

func (s *Call_expressionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(TurnkeyParserLPAREN, 0)
}

func (s *Call_expressionContext) Parameters() IParametersContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParametersContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParametersContext)
}

func (s *Call_expressionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(TurnkeyParserRPAREN, 0)
}

func (s *Call_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Call_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Call_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TurnkeyParserListener); ok {
		listenerT.EnterCall_expression(s)
	}
}

func (s *Call_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TurnkeyParserListener); ok {
		listenerT.ExitCall_expression(s)
	}
}

func (p *TurnkeyParser) Call_expression() (localctx ICall_expressionContext) {
	localctx = NewCall_expressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, TurnkeyParserRULE_call_expression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(135)
		p.Match(TurnkeyParserIDENT)
	}
	{
		p.SetState(136)
		p.Match(TurnkeyParserLPAREN)
	}
	{
		p.SetState(137)
		p.Parameters()
	}
	{
		p.SetState(138)
		p.Match(TurnkeyParserRPAREN)
	}

	return localctx
}

// ITurn_expressionContext is an interface to support dynamic dispatch.
type ITurn_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTurn_expressionContext differentiates from other interfaces.
	IsTurn_expressionContext()
}

type Turn_expressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTurn_expressionContext() *Turn_expressionContext {
	var p = new(Turn_expressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TurnkeyParserRULE_turn_expression
	return p
}

func (*Turn_expressionContext) IsTurn_expressionContext() {}

func NewTurn_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Turn_expressionContext {
	var p = new(Turn_expressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TurnkeyParserRULE_turn_expression

	return p
}

func (s *Turn_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Turn_expressionContext) TURN() antlr.TerminalNode {
	return s.GetToken(TurnkeyParserTURN, 0)
}

func (s *Turn_expressionContext) Call_expression() ICall_expressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICall_expressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICall_expressionContext)
}

func (s *Turn_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Turn_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Turn_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TurnkeyParserListener); ok {
		listenerT.EnterTurn_expression(s)
	}
}

func (s *Turn_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TurnkeyParserListener); ok {
		listenerT.ExitTurn_expression(s)
	}
}

func (p *TurnkeyParser) Turn_expression() (localctx ITurn_expressionContext) {
	localctx = NewTurn_expressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, TurnkeyParserRULE_turn_expression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(140)
		p.Match(TurnkeyParserTURN)
	}
	{
		p.SetState(141)
		p.Call_expression()
	}

	return localctx
}

func (p *TurnkeyParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 6:
		var t *Int_expressionContext = nil
		if localctx != nil {
			t = localctx.(*Int_expressionContext)
		}
		return p.Int_expression_Sempred(t, predIndex)

	case 7:
		var t *Float_expressionContext = nil
		if localctx != nil {
			t = localctx.(*Float_expressionContext)
		}
		return p.Float_expression_Sempred(t, predIndex)

	case 8:
		var t *Bool_expressionContext = nil
		if localctx != nil {
			t = localctx.(*Bool_expressionContext)
		}
		return p.Bool_expression_Sempred(t, predIndex)

	case 9:
		var t *String_expressionContext = nil
		if localctx != nil {
			t = localctx.(*String_expressionContext)
		}
		return p.String_expression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *TurnkeyParser) Int_expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *TurnkeyParser) Float_expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 2:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *TurnkeyParser) Bool_expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 4:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *TurnkeyParser) String_expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 5:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
