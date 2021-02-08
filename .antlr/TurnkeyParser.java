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
		WS=1, FUNC=2, TURN=3, MUL=4, DIV=5, ADD=6, SUB=7, ASSIGN=8, LPAREN=9, 
		RPAREN=10, INT=11, FLOAT=12, BOOL=13, STRING=14, IDENT=15;
	public static final int
		RULE_start = 0, RULE_expression = 1, RULE_int_expression = 2, RULE_float_expression = 3, 
		RULE_string_expression = 4;
	private static String[] makeRuleNames() {
		return new String[] {
			"start", "expression", "int_expression", "float_expression", "string_expression"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, null, "'func'", "'turn'", "'*'", "'/'", "'+'", "'-'", "'='", "'('", 
			"')'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "WS", "FUNC", "TURN", "MUL", "DIV", "ADD", "SUB", "ASSIGN", "LPAREN", 
			"RPAREN", "INT", "FLOAT", "BOOL", "STRING", "IDENT"
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
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
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
			setState(10);
			expression();
			setState(11);
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

	public static class ExpressionContext extends ParserRuleContext {
		public Int_expressionContext int_expression() {
			return getRuleContext(Int_expressionContext.class,0);
		}
		public Float_expressionContext float_expression() {
			return getRuleContext(Float_expressionContext.class,0);
		}
		public String_expressionContext string_expression() {
			return getRuleContext(String_expressionContext.class,0);
		}
		public ExpressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expression; }
	}

	public final ExpressionContext expression() throws RecognitionException {
		ExpressionContext _localctx = new ExpressionContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_expression);
		try {
			setState(16);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case INT:
				enterOuterAlt(_localctx, 1);
				{
				setState(13);
				int_expression(0);
				}
				break;
			case FLOAT:
				enterOuterAlt(_localctx, 2);
				{
				setState(14);
				float_expression(0);
				}
				break;
			case STRING:
				enterOuterAlt(_localctx, 3);
				{
				setState(15);
				string_expression(0);
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
		int _startState = 4;
		enterRecursionRule(_localctx, 4, RULE_int_expression, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(19);
			match(INT);
			}
			_ctx.stop = _input.LT(-1);
			setState(29);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,2,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(27);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,1,_ctx) ) {
					case 1:
						{
						_localctx = new Int_expressionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_int_expression);
						setState(21);
						if (!(precpred(_ctx, 3))) throw new FailedPredicateException(this, "precpred(_ctx, 3)");
						setState(22);
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
						setState(23);
						int_expression(4);
						}
						break;
					case 2:
						{
						_localctx = new Int_expressionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_int_expression);
						setState(24);
						if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
						setState(25);
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
						setState(26);
						int_expression(3);
						}
						break;
					}
					} 
				}
				setState(31);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,2,_ctx);
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
		int _startState = 6;
		enterRecursionRule(_localctx, 6, RULE_float_expression, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(33);
			match(FLOAT);
			}
			_ctx.stop = _input.LT(-1);
			setState(43);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,4,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(41);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,3,_ctx) ) {
					case 1:
						{
						_localctx = new Float_expressionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_float_expression);
						setState(35);
						if (!(precpred(_ctx, 3))) throw new FailedPredicateException(this, "precpred(_ctx, 3)");
						setState(36);
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
						setState(37);
						float_expression(4);
						}
						break;
					case 2:
						{
						_localctx = new Float_expressionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_float_expression);
						setState(38);
						if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
						setState(39);
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
						setState(40);
						float_expression(3);
						}
						break;
					}
					} 
				}
				setState(45);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,4,_ctx);
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

	public static class String_expressionContext extends ParserRuleContext {
		public Token op;
		public TerminalNode STRING() { return getToken(TurnkeyParser.STRING, 0); }
		public List<String_expressionContext> string_expression() {
			return getRuleContexts(String_expressionContext.class);
		}
		public String_expressionContext string_expression(int i) {
			return getRuleContext(String_expressionContext.class,i);
		}
		public TerminalNode ADD() { return getToken(TurnkeyParser.ADD, 0); }
		public String_expressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_string_expression; }
	}

	public final String_expressionContext string_expression() throws RecognitionException {
		return string_expression(0);
	}

	private String_expressionContext string_expression(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		String_expressionContext _localctx = new String_expressionContext(_ctx, _parentState);
		String_expressionContext _prevctx = _localctx;
		int _startState = 8;
		enterRecursionRule(_localctx, 8, RULE_string_expression, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(47);
			match(STRING);
			}
			_ctx.stop = _input.LT(-1);
			setState(54);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,5,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new String_expressionContext(_parentctx, _parentState);
					pushNewRecursionContext(_localctx, _startState, RULE_string_expression);
					setState(49);
					if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
					setState(50);
					((String_expressionContext)_localctx).op = match(ADD);
					setState(51);
					string_expression(3);
					}
					} 
				}
				setState(56);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,5,_ctx);
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
		case 2:
			return int_expression_sempred((Int_expressionContext)_localctx, predIndex);
		case 3:
			return float_expression_sempred((Float_expressionContext)_localctx, predIndex);
		case 4:
			return string_expression_sempred((String_expressionContext)_localctx, predIndex);
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
	private boolean string_expression_sempred(String_expressionContext _localctx, int predIndex) {
		switch (predIndex) {
		case 4:
			return precpred(_ctx, 2);
		}
		return true;
	}

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3\21<\4\2\t\2\4\3\t"+
		"\3\4\4\t\4\4\5\t\5\4\6\t\6\3\2\3\2\3\2\3\3\3\3\3\3\5\3\23\n\3\3\4\3\4"+
		"\3\4\3\4\3\4\3\4\3\4\3\4\3\4\7\4\36\n\4\f\4\16\4!\13\4\3\5\3\5\3\5\3\5"+
		"\3\5\3\5\3\5\3\5\3\5\7\5,\n\5\f\5\16\5/\13\5\3\6\3\6\3\6\3\6\3\6\3\6\7"+
		"\6\67\n\6\f\6\16\6:\13\6\3\6\2\5\6\b\n\7\2\4\6\b\n\2\4\3\2\6\7\3\2\b\t"+
		"\2=\2\f\3\2\2\2\4\22\3\2\2\2\6\24\3\2\2\2\b\"\3\2\2\2\n\60\3\2\2\2\f\r"+
		"\5\4\3\2\r\16\7\2\2\3\16\3\3\2\2\2\17\23\5\6\4\2\20\23\5\b\5\2\21\23\5"+
		"\n\6\2\22\17\3\2\2\2\22\20\3\2\2\2\22\21\3\2\2\2\23\5\3\2\2\2\24\25\b"+
		"\4\1\2\25\26\7\r\2\2\26\37\3\2\2\2\27\30\f\5\2\2\30\31\t\2\2\2\31\36\5"+
		"\6\4\6\32\33\f\4\2\2\33\34\t\3\2\2\34\36\5\6\4\5\35\27\3\2\2\2\35\32\3"+
		"\2\2\2\36!\3\2\2\2\37\35\3\2\2\2\37 \3\2\2\2 \7\3\2\2\2!\37\3\2\2\2\""+
		"#\b\5\1\2#$\7\16\2\2$-\3\2\2\2%&\f\5\2\2&\'\t\2\2\2\',\5\b\5\6()\f\4\2"+
		"\2)*\t\3\2\2*,\5\b\5\5+%\3\2\2\2+(\3\2\2\2,/\3\2\2\2-+\3\2\2\2-.\3\2\2"+
		"\2.\t\3\2\2\2/-\3\2\2\2\60\61\b\6\1\2\61\62\7\20\2\2\628\3\2\2\2\63\64"+
		"\f\4\2\2\64\65\7\b\2\2\65\67\5\n\6\5\66\63\3\2\2\2\67:\3\2\2\28\66\3\2"+
		"\2\289\3\2\2\29\13\3\2\2\2:8\3\2\2\2\b\22\35\37+-8";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}