// Generated from /home/digyx/Documents/Turnkey/Turnkey.g4 by ANTLR 4.8
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class TurnkeyParser extends Parser {
	static { RuntimeMetaData.checkVersion("4.8", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		WS=1, COMMA=2, FUNC=3, TURN=4, MUL=5, DIV=6, ADD=7, SUB=8, ASSIGN=9, LPAREN=10, 
		RPAREN=11, LBRACE=12, RBRACE=13, INT=14, FLOAT=15, BOOL=16, IDENT=17;
	public static final int
		RULE_start = 0, RULE_statement = 1, RULE_block = 2, RULE_func_def = 3, 
		RULE_func_call = 4, RULE_expression = 5, RULE_int_expression = 6, RULE_float_expression = 7;
	private static String[] makeRuleNames() {
		return new String[] {
			"start", "statement", "block", "func_def", "func_call", "expression", 
			"int_expression", "float_expression"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, null, "','", "'func'", "'turn'", "'*'", "'/'", "'+'", "'-'", "'='", 
			"'('", "')'", "'{'", "'}'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "WS", "COMMA", "FUNC", "TURN", "MUL", "DIV", "ADD", "SUB", "ASSIGN", 
			"LPAREN", "RPAREN", "LBRACE", "RBRACE", "INT", "FLOAT", "BOOL", "IDENT"
		};
	}
	private static final String[] _SYMBOLIC_NAMES = makeSymbolicNames();
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}

	@Override
	public String getGrammarFileName() { return "Turnkey.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public TurnkeyParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	public static class StartContext extends ParserRuleContext {
		public StatementContext statement() {
			return getRuleContext(StatementContext.class,0);
		}
		public TerminalNode EOF() { return getToken(TurnkeyParser.EOF, 0); }
		public StartContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_start; }
	}

	public final StartContext start() throws RecognitionException {
		StartContext _localctx = new StartContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_start);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(16);
			statement();
			setState(17);
			match(EOF);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class StatementContext extends ParserRuleContext {
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public Func_defContext func_def() {
			return getRuleContext(Func_defContext.class,0);
		}
		public StatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_statement; }
	}

	public final StatementContext statement() throws RecognitionException {
		StatementContext _localctx = new StatementContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_statement);
		try {
			setState(21);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case INT:
			case FLOAT:
				enterOuterAlt(_localctx, 1);
				{
				setState(19);
				expression();
				}
				break;
			case FUNC:
				enterOuterAlt(_localctx, 2);
				{
				setState(20);
				func_def();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class BlockContext extends ParserRuleContext {
		public TerminalNode LBRACE() { return getToken(TurnkeyParser.LBRACE, 0); }
		public TerminalNode RBRACE() { return getToken(TurnkeyParser.RBRACE, 0); }
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public BlockContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_block; }
	}

	public final BlockContext block() throws RecognitionException {
		BlockContext _localctx = new BlockContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_block);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(23);
			match(LBRACE);
			setState(27);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==INT || _la==FLOAT) {
				{
				{
				setState(24);
				expression();
				}
				}
				setState(29);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(30);
			match(RBRACE);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Func_defContext extends ParserRuleContext {
		public TerminalNode FUNC() { return getToken(TurnkeyParser.FUNC, 0); }
		public TerminalNode LPAREN() { return getToken(TurnkeyParser.LPAREN, 0); }
		public TerminalNode RPAREN() { return getToken(TurnkeyParser.RPAREN, 0); }
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public List<TerminalNode> IDENT() { return getTokens(TurnkeyParser.IDENT); }
		public TerminalNode IDENT(int i) {
			return getToken(TurnkeyParser.IDENT, i);
		}
		public List<TerminalNode> COMMA() { return getTokens(TurnkeyParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(TurnkeyParser.COMMA, i);
		}
		public Func_defContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_func_def; }
	}

	public final Func_defContext func_def() throws RecognitionException {
		Func_defContext _localctx = new Func_defContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_func_def);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(32);
			match(FUNC);
			setState(33);
			match(LPAREN);
			setState(38);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==IDENT) {
				{
				{
				setState(34);
				match(IDENT);
				setState(35);
				match(COMMA);
				}
				}
				setState(40);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(41);
			match(RPAREN);
			setState(42);
			block();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Func_callContext extends ParserRuleContext {
		public TerminalNode IDENT() { return getToken(TurnkeyParser.IDENT, 0); }
		public TerminalNode LPAREN() { return getToken(TurnkeyParser.LPAREN, 0); }
		public TerminalNode RPAREN() { return getToken(TurnkeyParser.RPAREN, 0); }
		public TerminalNode COMMA() { return getToken(TurnkeyParser.COMMA, 0); }
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public Func_callContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_func_call; }
	}

	public final Func_callContext func_call() throws RecognitionException {
		Func_callContext _localctx = new Func_callContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_func_call);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(44);
			match(IDENT);
			setState(45);
			match(LPAREN);
			{
			setState(49);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==INT || _la==FLOAT) {
				{
				{
				setState(46);
				expression();
				}
				}
				setState(51);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(52);
			match(COMMA);
			}
			setState(54);
			match(RPAREN);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ExpressionContext extends ParserRuleContext {
		public Int_expressionContext int_expression() {
			return getRuleContext(Int_expressionContext.class,0);
		}
		public Float_expressionContext float_expression() {
			return getRuleContext(Float_expressionContext.class,0);
		}
		public ExpressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expression; }
	}

	public final ExpressionContext expression() throws RecognitionException {
		ExpressionContext _localctx = new ExpressionContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_expression);
		try {
			setState(58);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case INT:
				enterOuterAlt(_localctx, 1);
				{
				setState(56);
				int_expression(0);
				}
				break;
			case FLOAT:
				enterOuterAlt(_localctx, 2);
				{
				setState(57);
				float_expression(0);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Int_expressionContext extends ParserRuleContext {
		public Token op;
		public TerminalNode INT() { return getToken(TurnkeyParser.INT, 0); }
		public List<Int_expressionContext> int_expression() {
			return getRuleContexts(Int_expressionContext.class);
		}
		public Int_expressionContext int_expression(int i) {
			return getRuleContext(Int_expressionContext.class,i);
		}
		public TerminalNode MUL() { return getToken(TurnkeyParser.MUL, 0); }
		public TerminalNode DIV() { return getToken(TurnkeyParser.DIV, 0); }
		public TerminalNode ADD() { return getToken(TurnkeyParser.ADD, 0); }
		public TerminalNode SUB() { return getToken(TurnkeyParser.SUB, 0); }
		public Int_expressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_int_expression; }
	}

	public final Int_expressionContext int_expression() throws RecognitionException {
		return int_expression(0);
	}

	private Int_expressionContext int_expression(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		Int_expressionContext _localctx = new Int_expressionContext(_ctx, _parentState);
		Int_expressionContext _prevctx = _localctx;
		int _startState = 12;
		enterRecursionRule(_localctx, 12, RULE_int_expression, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(61);
			match(INT);
			}
			_ctx.stop = _input.LT(-1);
			setState(71);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,6,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(69);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,5,_ctx) ) {
					case 1:
						{
						_localctx = new Int_expressionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_int_expression);
						setState(63);
						if (!(precpred(_ctx, 3))) throw new FailedPredicateException(this, "precpred(_ctx, 3)");
						setState(64);
						((Int_expressionContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==MUL || _la==DIV) ) {
							((Int_expressionContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(65);
						int_expression(4);
						}
						break;
					case 2:
						{
						_localctx = new Int_expressionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_int_expression);
						setState(66);
						if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
						setState(67);
						((Int_expressionContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==ADD || _la==SUB) ) {
							((Int_expressionContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(68);
						int_expression(3);
						}
						break;
					}
					} 
				}
				setState(73);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,6,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	public static class Float_expressionContext extends ParserRuleContext {
		public Token op;
		public TerminalNode FLOAT() { return getToken(TurnkeyParser.FLOAT, 0); }
		public List<Float_expressionContext> float_expression() {
			return getRuleContexts(Float_expressionContext.class);
		}
		public Float_expressionContext float_expression(int i) {
			return getRuleContext(Float_expressionContext.class,i);
		}
		public TerminalNode MUL() { return getToken(TurnkeyParser.MUL, 0); }
		public TerminalNode DIV() { return getToken(TurnkeyParser.DIV, 0); }
		public TerminalNode ADD() { return getToken(TurnkeyParser.ADD, 0); }
		public TerminalNode SUB() { return getToken(TurnkeyParser.SUB, 0); }
		public Float_expressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_float_expression; }
	}

	public final Float_expressionContext float_expression() throws RecognitionException {
		return float_expression(0);
	}

	private Float_expressionContext float_expression(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		Float_expressionContext _localctx = new Float_expressionContext(_ctx, _parentState);
		Float_expressionContext _prevctx = _localctx;
		int _startState = 14;
		enterRecursionRule(_localctx, 14, RULE_float_expression, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(75);
			match(FLOAT);
			}
			_ctx.stop = _input.LT(-1);
			setState(85);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,8,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(83);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,7,_ctx) ) {
					case 1:
						{
						_localctx = new Float_expressionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_float_expression);
						setState(77);
						if (!(precpred(_ctx, 3))) throw new FailedPredicateException(this, "precpred(_ctx, 3)");
						setState(78);
						((Float_expressionContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==MUL || _la==DIV) ) {
							((Float_expressionContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(79);
						float_expression(4);
						}
						break;
					case 2:
						{
						_localctx = new Float_expressionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_float_expression);
						setState(80);
						if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
						setState(81);
						((Float_expressionContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==ADD || _la==SUB) ) {
							((Float_expressionContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(82);
						float_expression(3);
						}
						break;
					}
					} 
				}
				setState(87);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,8,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	public boolean sempred(RuleContext _localctx, int ruleIndex, int predIndex) {
		switch (ruleIndex) {
		case 6:
			return int_expression_sempred((Int_expressionContext)_localctx, predIndex);
		case 7:
			return float_expression_sempred((Float_expressionContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean int_expression_sempred(Int_expressionContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 3);
		case 1:
			return precpred(_ctx, 2);
		}
		return true;
	}
	private boolean float_expression_sempred(Float_expressionContext _localctx, int predIndex) {
		switch (predIndex) {
		case 2:
			return precpred(_ctx, 3);
		case 3:
			return precpred(_ctx, 2);
		}
		return true;
	}

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3\23[\4\2\t\2\4\3\t"+
		"\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\3\2\3\2\3\2\3\3\3\3"+
		"\5\3\30\n\3\3\4\3\4\7\4\34\n\4\f\4\16\4\37\13\4\3\4\3\4\3\5\3\5\3\5\3"+
		"\5\7\5\'\n\5\f\5\16\5*\13\5\3\5\3\5\3\5\3\6\3\6\3\6\7\6\62\n\6\f\6\16"+
		"\6\65\13\6\3\6\3\6\3\6\3\6\3\7\3\7\5\7=\n\7\3\b\3\b\3\b\3\b\3\b\3\b\3"+
		"\b\3\b\3\b\7\bH\n\b\f\b\16\bK\13\b\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t"+
		"\7\tV\n\t\f\t\16\tY\13\t\3\t\2\4\16\20\n\2\4\6\b\n\f\16\20\2\4\3\2\7\b"+
		"\3\2\t\n\2[\2\22\3\2\2\2\4\27\3\2\2\2\6\31\3\2\2\2\b\"\3\2\2\2\n.\3\2"+
		"\2\2\f<\3\2\2\2\16>\3\2\2\2\20L\3\2\2\2\22\23\5\4\3\2\23\24\7\2\2\3\24"+
		"\3\3\2\2\2\25\30\5\f\7\2\26\30\5\b\5\2\27\25\3\2\2\2\27\26\3\2\2\2\30"+
		"\5\3\2\2\2\31\35\7\16\2\2\32\34\5\f\7\2\33\32\3\2\2\2\34\37\3\2\2\2\35"+
		"\33\3\2\2\2\35\36\3\2\2\2\36 \3\2\2\2\37\35\3\2\2\2 !\7\17\2\2!\7\3\2"+
		"\2\2\"#\7\5\2\2#(\7\f\2\2$%\7\23\2\2%\'\7\4\2\2&$\3\2\2\2\'*\3\2\2\2("+
		"&\3\2\2\2()\3\2\2\2)+\3\2\2\2*(\3\2\2\2+,\7\r\2\2,-\5\6\4\2-\t\3\2\2\2"+
		"./\7\23\2\2/\63\7\f\2\2\60\62\5\f\7\2\61\60\3\2\2\2\62\65\3\2\2\2\63\61"+
		"\3\2\2\2\63\64\3\2\2\2\64\66\3\2\2\2\65\63\3\2\2\2\66\67\7\4\2\2\678\3"+
		"\2\2\289\7\r\2\29\13\3\2\2\2:=\5\16\b\2;=\5\20\t\2<:\3\2\2\2<;\3\2\2\2"+
		"=\r\3\2\2\2>?\b\b\1\2?@\7\20\2\2@I\3\2\2\2AB\f\5\2\2BC\t\2\2\2CH\5\16"+
		"\b\6DE\f\4\2\2EF\t\3\2\2FH\5\16\b\5GA\3\2\2\2GD\3\2\2\2HK\3\2\2\2IG\3"+
		"\2\2\2IJ\3\2\2\2J\17\3\2\2\2KI\3\2\2\2LM\b\t\1\2MN\7\21\2\2NW\3\2\2\2"+
		"OP\f\5\2\2PQ\t\2\2\2QV\5\20\t\6RS\f\4\2\2ST\t\3\2\2TV\5\20\t\5UO\3\2\2"+
		"\2UR\3\2\2\2VY\3\2\2\2WU\3\2\2\2WX\3\2\2\2X\21\3\2\2\2YW\3\2\2\2\13\27"+
		"\35(\63<GIUW";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}