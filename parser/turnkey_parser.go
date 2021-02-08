// Code generated from Turnkey.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // Turnkey

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 19, 102,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 3, 2,
	3, 2, 3, 2, 3, 3, 3, 3, 5, 3, 30, 10, 3, 3, 4, 3, 4, 7, 4, 34, 10, 4, 12,
	4, 14, 4, 37, 11, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3,
	5, 3, 6, 3, 6, 7, 6, 50, 10, 6, 12, 6, 14, 6, 53, 11, 6, 3, 6, 5, 6, 56,
	10, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 5, 7, 63, 10, 7, 3, 8, 3, 8, 3, 8,
	3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 7, 8, 74, 10, 8, 12, 8, 14, 8, 77,
	11, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 7, 9, 88,
	10, 9, 12, 9, 14, 9, 91, 11, 9, 3, 10, 3, 10, 3, 11, 3, 11, 3, 12, 3, 12,
	3, 12, 3, 12, 3, 12, 3, 12, 2, 4, 14, 16, 13, 2, 4, 6, 8, 10, 12, 14, 16,
	18, 20, 22, 2, 4, 3, 2, 7, 8, 3, 2, 9, 10, 2, 102, 2, 24, 3, 2, 2, 2, 4,
	29, 3, 2, 2, 2, 6, 31, 3, 2, 2, 2, 8, 40, 3, 2, 2, 2, 10, 51, 3, 2, 2,
	2, 12, 62, 3, 2, 2, 2, 14, 64, 3, 2, 2, 2, 16, 78, 3, 2, 2, 2, 18, 92,
	3, 2, 2, 2, 20, 94, 3, 2, 2, 2, 22, 96, 3, 2, 2, 2, 24, 25, 5, 4, 3, 2,
	25, 26, 7, 2, 2, 3, 26, 3, 3, 2, 2, 2, 27, 30, 5, 12, 7, 2, 28, 30, 5,
	8, 5, 2, 29, 27, 3, 2, 2, 2, 29, 28, 3, 2, 2, 2, 30, 5, 3, 2, 2, 2, 31,
	35, 7, 14, 2, 2, 32, 34, 5, 12, 7, 2, 33, 32, 3, 2, 2, 2, 34, 37, 3, 2,
	2, 2, 35, 33, 3, 2, 2, 2, 35, 36, 3, 2, 2, 2, 36, 38, 3, 2, 2, 2, 37, 35,
	3, 2, 2, 2, 38, 39, 7, 15, 2, 2, 39, 7, 3, 2, 2, 2, 40, 41, 7, 5, 2, 2,
	41, 42, 7, 19, 2, 2, 42, 43, 7, 12, 2, 2, 43, 44, 5, 10, 6, 2, 44, 45,
	7, 13, 2, 2, 45, 46, 5, 6, 4, 2, 46, 9, 3, 2, 2, 2, 47, 48, 7, 19, 2, 2,
	48, 50, 7, 4, 2, 2, 49, 47, 3, 2, 2, 2, 50, 53, 3, 2, 2, 2, 51, 49, 3,
	2, 2, 2, 51, 52, 3, 2, 2, 2, 52, 55, 3, 2, 2, 2, 53, 51, 3, 2, 2, 2, 54,
	56, 7, 19, 2, 2, 55, 54, 3, 2, 2, 2, 55, 56, 3, 2, 2, 2, 56, 11, 3, 2,
	2, 2, 57, 63, 5, 14, 8, 2, 58, 63, 5, 16, 9, 2, 59, 63, 5, 18, 10, 2, 60,
	63, 5, 20, 11, 2, 61, 63, 5, 22, 12, 2, 62, 57, 3, 2, 2, 2, 62, 58, 3,
	2, 2, 2, 62, 59, 3, 2, 2, 2, 62, 60, 3, 2, 2, 2, 62, 61, 3, 2, 2, 2, 63,
	13, 3, 2, 2, 2, 64, 65, 8, 8, 1, 2, 65, 66, 7, 16, 2, 2, 66, 75, 3, 2,
	2, 2, 67, 68, 12, 5, 2, 2, 68, 69, 9, 2, 2, 2, 69, 74, 5, 14, 8, 6, 70,
	71, 12, 4, 2, 2, 71, 72, 9, 3, 2, 2, 72, 74, 5, 14, 8, 5, 73, 67, 3, 2,
	2, 2, 73, 70, 3, 2, 2, 2, 74, 77, 3, 2, 2, 2, 75, 73, 3, 2, 2, 2, 75, 76,
	3, 2, 2, 2, 76, 15, 3, 2, 2, 2, 77, 75, 3, 2, 2, 2, 78, 79, 8, 9, 1, 2,
	79, 80, 7, 17, 2, 2, 80, 89, 3, 2, 2, 2, 81, 82, 12, 5, 2, 2, 82, 83, 9,
	2, 2, 2, 83, 88, 5, 16, 9, 6, 84, 85, 12, 4, 2, 2, 85, 86, 9, 3, 2, 2,
	86, 88, 5, 16, 9, 5, 87, 81, 3, 2, 2, 2, 87, 84, 3, 2, 2, 2, 88, 91, 3,
	2, 2, 2, 89, 87, 3, 2, 2, 2, 89, 90, 3, 2, 2, 2, 90, 17, 3, 2, 2, 2, 91,
	89, 3, 2, 2, 2, 92, 93, 7, 18, 2, 2, 93, 19, 3, 2, 2, 2, 94, 95, 7, 19,
	2, 2, 95, 21, 3, 2, 2, 2, 96, 97, 7, 19, 2, 2, 97, 98, 7, 12, 2, 2, 98,
	99, 5, 10, 6, 2, 99, 100, 7, 13, 2, 2, 100, 23, 3, 2, 2, 2, 11, 29, 35,
	51, 55, 62, 73, 75, 87, 89,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "", "','", "'func'", "'turn'", "'*'", "'/'", "'+'", "'-'", "'='", "'('",
	"')'", "'{'", "'}'",
}
var symbolicNames = []string{
	"", "WS", "COMMA", "FUNC", "TURN", "MUL", "DIV", "ADD", "SUB", "ASSIGN",
	"LPAREN", "RPAREN", "LBRACE", "RBRACE", "INT", "FLOAT", "BOOL", "IDENT",
}

var ruleNames = []string{
	"start", "statement", "block", "func_def", "parameters", "expression",
	"int_expression", "float_expression", "bool_expression", "ident_expression",
	"call_expression",
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
	this.GrammarFileName = "Turnkey.g4"

	return this
}

// TurnkeyParser tokens.
const (
	TurnkeyParserEOF    = antlr.TokenEOF
	TurnkeyParserWS     = 1
	TurnkeyParserCOMMA  = 2
	TurnkeyParserFUNC   = 3
	TurnkeyParserTURN   = 4
	TurnkeyParserMUL    = 5
	TurnkeyParserDIV    = 6
	TurnkeyParserADD    = 7
	TurnkeyParserSUB    = 8
	TurnkeyParserASSIGN = 9
	TurnkeyParserLPAREN = 10
	TurnkeyParserRPAREN = 11
	TurnkeyParserLBRACE = 12
	TurnkeyParserRBRACE = 13
	TurnkeyParserINT    = 14
	TurnkeyParserFLOAT  = 15
	TurnkeyParserBOOL   = 16
	TurnkeyParserIDENT  = 17
)

// TurnkeyParser rules.
const (
	TurnkeyParserRULE_start            = 0
	TurnkeyParserRULE_statement        = 1
	TurnkeyParserRULE_block            = 2
	TurnkeyParserRULE_func_def         = 3
	TurnkeyParserRULE_parameters       = 4
	TurnkeyParserRULE_expression       = 5
	TurnkeyParserRULE_int_expression   = 6
	TurnkeyParserRULE_float_expression = 7
	TurnkeyParserRULE_bool_expression  = 8
	TurnkeyParserRULE_ident_expression = 9
	TurnkeyParserRULE_call_expression  = 10
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
	if listenerT, ok := listener.(TurnkeyListener); ok {
		listenerT.EnterStart(s)
	}
}

func (s *StartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TurnkeyListener); ok {
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
		p.SetState(22)
		p.Statement()
	}
	{
		p.SetState(23)
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
	if listenerT, ok := listener.(TurnkeyListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TurnkeyListener); ok {
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

	p.SetState(27)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case TurnkeyParserINT, TurnkeyParserFLOAT, TurnkeyParserBOOL, TurnkeyParserIDENT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(25)
			p.Expression()
		}

	case TurnkeyParserFUNC:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(26)
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
	if listenerT, ok := listener.(TurnkeyListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TurnkeyListener); ok {
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
		p.SetState(29)
		p.Match(TurnkeyParserLBRACE)
	}
	p.SetState(33)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<TurnkeyParserINT)|(1<<TurnkeyParserFLOAT)|(1<<TurnkeyParserBOOL)|(1<<TurnkeyParserIDENT))) != 0 {
		{
			p.SetState(30)
			p.Expression()
		}

		p.SetState(35)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(36)
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
	if listenerT, ok := listener.(TurnkeyListener); ok {
		listenerT.EnterFunc_def(s)
	}
}

func (s *Func_defContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TurnkeyListener); ok {
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
		p.SetState(38)
		p.Match(TurnkeyParserFUNC)
	}
	{
		p.SetState(39)
		p.Match(TurnkeyParserIDENT)
	}
	{
		p.SetState(40)
		p.Match(TurnkeyParserLPAREN)
	}
	{
		p.SetState(41)
		p.Parameters()
	}
	{
		p.SetState(42)
		p.Match(TurnkeyParserRPAREN)
	}
	{
		p.SetState(43)
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
	if listenerT, ok := listener.(TurnkeyListener); ok {
		listenerT.EnterParameters(s)
	}
}

func (s *ParametersContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TurnkeyListener); ok {
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
	p.SetState(49)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(45)
				p.Match(TurnkeyParserIDENT)
			}
			{
				p.SetState(46)
				p.Match(TurnkeyParserCOMMA)
			}

		}
		p.SetState(51)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())
	}
	p.SetState(53)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == TurnkeyParserIDENT {
		{
			p.SetState(52)
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

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TurnkeyListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TurnkeyListener); ok {
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

	p.SetState(60)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(55)
			p.int_expression(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(56)
			p.float_expression(0)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(57)
			p.Bool_expression()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(58)
			p.Ident_expression()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(59)
			p.Call_expression()
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

func (s *Int_expressionContext) INT() antlr.TerminalNode {
	return s.GetToken(TurnkeyParserINT, 0)
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
	if listenerT, ok := listener.(TurnkeyListener); ok {
		listenerT.EnterInt_expression(s)
	}
}

func (s *Int_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TurnkeyListener); ok {
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
	{
		p.SetState(63)
		p.Match(TurnkeyParserINT)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(73)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(71)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
			case 1:
				localctx = NewInt_expressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, TurnkeyParserRULE_int_expression)
				p.SetState(65)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(66)

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
					p.SetState(67)
					p.int_expression(4)
				}

			case 2:
				localctx = NewInt_expressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, TurnkeyParserRULE_int_expression)
				p.SetState(68)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(69)

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
					p.SetState(70)
					p.int_expression(3)
				}

			}

		}
		p.SetState(75)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())
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

func (s *Float_expressionContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(TurnkeyParserFLOAT, 0)
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
	if listenerT, ok := listener.(TurnkeyListener); ok {
		listenerT.EnterFloat_expression(s)
	}
}

func (s *Float_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TurnkeyListener); ok {
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
	{
		p.SetState(77)
		p.Match(TurnkeyParserFLOAT)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(87)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(85)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) {
			case 1:
				localctx = NewFloat_expressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, TurnkeyParserRULE_float_expression)
				p.SetState(79)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(80)

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
					p.SetState(81)
					p.float_expression(4)
				}

			case 2:
				localctx = NewFloat_expressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, TurnkeyParserRULE_float_expression)
				p.SetState(82)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(83)

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
					p.SetState(84)
					p.float_expression(3)
				}

			}

		}
		p.SetState(89)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())
	}

	return localctx
}

// IBool_expressionContext is an interface to support dynamic dispatch.
type IBool_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBool_expressionContext differentiates from other interfaces.
	IsBool_expressionContext()
}

type Bool_expressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
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

func (s *Bool_expressionContext) BOOL() antlr.TerminalNode {
	return s.GetToken(TurnkeyParserBOOL, 0)
}

func (s *Bool_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Bool_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Bool_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TurnkeyListener); ok {
		listenerT.EnterBool_expression(s)
	}
}

func (s *Bool_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TurnkeyListener); ok {
		listenerT.ExitBool_expression(s)
	}
}

func (p *TurnkeyParser) Bool_expression() (localctx IBool_expressionContext) {
	localctx = NewBool_expressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, TurnkeyParserRULE_bool_expression)

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
		p.SetState(90)
		p.Match(TurnkeyParserBOOL)
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
	if listenerT, ok := listener.(TurnkeyListener); ok {
		listenerT.EnterIdent_expression(s)
	}
}

func (s *Ident_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TurnkeyListener); ok {
		listenerT.ExitIdent_expression(s)
	}
}

func (p *TurnkeyParser) Ident_expression() (localctx IIdent_expressionContext) {
	localctx = NewIdent_expressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, TurnkeyParserRULE_ident_expression)

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
		p.SetState(92)
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
	if listenerT, ok := listener.(TurnkeyListener); ok {
		listenerT.EnterCall_expression(s)
	}
}

func (s *Call_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TurnkeyListener); ok {
		listenerT.ExitCall_expression(s)
	}
}

func (p *TurnkeyParser) Call_expression() (localctx ICall_expressionContext) {
	localctx = NewCall_expressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, TurnkeyParserRULE_call_expression)

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
		p.SetState(94)
		p.Match(TurnkeyParserIDENT)
	}
	{
		p.SetState(95)
		p.Match(TurnkeyParserLPAREN)
	}
	{
		p.SetState(96)
		p.Parameters()
	}
	{
		p.SetState(97)
		p.Match(TurnkeyParserRPAREN)
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
