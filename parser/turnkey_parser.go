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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 17, 46, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 3, 2, 3, 2, 3, 2, 3, 3, 3,
	3, 5, 3, 16, 10, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	4, 7, 4, 27, 10, 4, 12, 4, 14, 4, 30, 11, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3,
	5, 3, 5, 3, 5, 3, 5, 3, 5, 7, 5, 41, 10, 5, 12, 5, 14, 5, 44, 11, 5, 3,
	5, 2, 4, 6, 8, 6, 2, 4, 6, 8, 2, 4, 3, 2, 6, 7, 3, 2, 8, 9, 2, 46, 2, 10,
	3, 2, 2, 2, 4, 15, 3, 2, 2, 2, 6, 17, 3, 2, 2, 2, 8, 31, 3, 2, 2, 2, 10,
	11, 5, 4, 3, 2, 11, 12, 7, 2, 2, 3, 12, 3, 3, 2, 2, 2, 13, 16, 5, 6, 4,
	2, 14, 16, 5, 8, 5, 2, 15, 13, 3, 2, 2, 2, 15, 14, 3, 2, 2, 2, 16, 5, 3,
	2, 2, 2, 17, 18, 8, 4, 1, 2, 18, 19, 7, 13, 2, 2, 19, 28, 3, 2, 2, 2, 20,
	21, 12, 5, 2, 2, 21, 22, 9, 2, 2, 2, 22, 27, 5, 6, 4, 6, 23, 24, 12, 4,
	2, 2, 24, 25, 9, 3, 2, 2, 25, 27, 5, 6, 4, 5, 26, 20, 3, 2, 2, 2, 26, 23,
	3, 2, 2, 2, 27, 30, 3, 2, 2, 2, 28, 26, 3, 2, 2, 2, 28, 29, 3, 2, 2, 2,
	29, 7, 3, 2, 2, 2, 30, 28, 3, 2, 2, 2, 31, 32, 8, 5, 1, 2, 32, 33, 7, 14,
	2, 2, 33, 42, 3, 2, 2, 2, 34, 35, 12, 5, 2, 2, 35, 36, 9, 2, 2, 2, 36,
	41, 5, 8, 5, 6, 37, 38, 12, 4, 2, 2, 38, 39, 9, 3, 2, 2, 39, 41, 5, 8,
	5, 5, 40, 34, 3, 2, 2, 2, 40, 37, 3, 2, 2, 2, 41, 44, 3, 2, 2, 2, 42, 40,
	3, 2, 2, 2, 42, 43, 3, 2, 2, 2, 43, 9, 3, 2, 2, 2, 44, 42, 3, 2, 2, 2,
	7, 15, 26, 28, 40, 42,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "", "'func'", "'turn'", "'*'", "'/'", "'+'", "'-'", "'='", "'('", "')'",
}
var symbolicNames = []string{
	"", "WS", "FUNC", "TURN", "MUL", "DIV", "ADD", "SUB", "ASSIGN", "LPAREN",
	"RPAREN", "INT", "FLOAT", "BOOL", "STRING", "IDENT",
}

var ruleNames = []string{
	"start", "expression", "int_expression", "float_expression",
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
	TurnkeyParserFUNC   = 2
	TurnkeyParserTURN   = 3
	TurnkeyParserMUL    = 4
	TurnkeyParserDIV    = 5
	TurnkeyParserADD    = 6
	TurnkeyParserSUB    = 7
	TurnkeyParserASSIGN = 8
	TurnkeyParserLPAREN = 9
	TurnkeyParserRPAREN = 10
	TurnkeyParserINT    = 11
	TurnkeyParserFLOAT  = 12
	TurnkeyParserBOOL   = 13
	TurnkeyParserSTRING = 14
	TurnkeyParserIDENT  = 15
)

// TurnkeyParser rules.
const (
	TurnkeyParserRULE_start            = 0
	TurnkeyParserRULE_expression       = 1
	TurnkeyParserRULE_int_expression   = 2
	TurnkeyParserRULE_float_expression = 3
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

func (s *StartContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
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
		p.SetState(8)
		p.Expression()
	}
	{
		p.SetState(9)
		p.Match(TurnkeyParserEOF)
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
	p.EnterRule(localctx, 2, TurnkeyParserRULE_expression)

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

	p.SetState(13)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case TurnkeyParserINT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(11)
			p.int_expression(0)
		}

	case TurnkeyParserFLOAT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(12)
			p.float_expression(0)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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
	_startState := 4
	p.EnterRecursionRule(localctx, 4, TurnkeyParserRULE_int_expression, _p)
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
		p.SetState(16)
		p.Match(TurnkeyParserINT)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(26)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(24)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
			case 1:
				localctx = NewInt_expressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, TurnkeyParserRULE_int_expression)
				p.SetState(18)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(19)

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
					p.SetState(20)
					p.int_expression(4)
				}

			case 2:
				localctx = NewInt_expressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, TurnkeyParserRULE_int_expression)
				p.SetState(21)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(22)

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
					p.SetState(23)
					p.int_expression(3)
				}

			}

		}
		p.SetState(28)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())
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
	_startState := 6
	p.EnterRecursionRule(localctx, 6, TurnkeyParserRULE_float_expression, _p)
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
		p.SetState(30)
		p.Match(TurnkeyParserFLOAT)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(40)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(38)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
			case 1:
				localctx = NewFloat_expressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, TurnkeyParserRULE_float_expression)
				p.SetState(32)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(33)

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
					p.SetState(34)
					p.float_expression(4)
				}

			case 2:
				localctx = NewFloat_expressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, TurnkeyParserRULE_float_expression)
				p.SetState(35)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(36)

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
					p.SetState(37)
					p.float_expression(3)
				}

			}

		}
		p.SetState(42)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())
	}

	return localctx
}

func (p *TurnkeyParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 2:
		var t *Int_expressionContext = nil
		if localctx != nil {
			t = localctx.(*Int_expressionContext)
		}
		return p.Int_expression_Sempred(t, predIndex)

	case 3:
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
