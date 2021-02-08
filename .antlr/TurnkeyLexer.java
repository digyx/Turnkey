// Generated from /home/digyx/Documents/Turnkey/Turnkey.g4 by ANTLR 4.8
import org.antlr.v4.runtime.Lexer;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.TokenStream;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.misc.*;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class TurnkeyLexer extends Lexer {
	static { RuntimeMetaData.checkVersion("4.8", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		WS=1, COMMA=2, FUNC=3, TURN=4, MUL=5, DIV=6, ADD=7, SUB=8, ASSIGN=9, LPAREN=10, 
		RPAREN=11, LBRACE=12, RBRACE=13, INT=14, FLOAT=15, BOOL=16, IDENT=17;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"WS", "COMMA", "FUNC", "TURN", "MUL", "DIV", "ADD", "SUB", "ASSIGN", 
			"LPAREN", "RPAREN", "LBRACE", "RBRACE", "INT", "FLOAT", "BOOL", "IDENT"
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


	public TurnkeyLexer(CharStream input) {
		super(input);
		_interp = new LexerATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@Override
	public String getGrammarFileName() { return "Turnkey.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public String[] getChannelNames() { return channelNames; }

	@Override
	public String[] getModeNames() { return modeNames; }

	@Override
	public ATN getATN() { return _ATN; }

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2\23v\b\1\4\2\t\2\4"+
		"\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t"+
		"\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\3\2\6\2\'\n\2\r\2\16\2(\3\2\3\2\3\3\3\3\3\4\3\4\3\4\3\4\3\4\3\5\3\5\3"+
		"\5\3\5\3\5\3\6\3\6\3\7\3\7\3\b\3\b\3\t\3\t\3\n\3\n\3\13\3\13\3\f\3\f\3"+
		"\r\3\r\3\16\3\16\3\17\3\17\7\17M\n\17\f\17\16\17P\13\17\3\20\3\20\7\20"+
		"T\n\20\f\20\16\20W\13\20\3\20\3\20\6\20[\n\20\r\20\16\20\\\3\21\3\21\3"+
		"\21\3\21\3\21\3\21\3\21\3\21\3\21\5\21h\n\21\3\22\7\22k\n\22\f\22\16\22"+
		"n\13\22\3\22\3\22\7\22r\n\22\f\22\16\22u\13\22\2\2\23\3\3\5\4\7\5\t\6"+
		"\13\7\r\b\17\t\21\n\23\13\25\f\27\r\31\16\33\17\35\20\37\21!\22#\23\3"+
		"\2\b\5\2\13\f\17\17\"\"\3\2\63;\3\2\62;\3\2aa\4\2C\\c|\6\2\62;C\\aac|"+
		"\2|\2\3\3\2\2\2\2\5\3\2\2\2\2\7\3\2\2\2\2\t\3\2\2\2\2\13\3\2\2\2\2\r\3"+
		"\2\2\2\2\17\3\2\2\2\2\21\3\2\2\2\2\23\3\2\2\2\2\25\3\2\2\2\2\27\3\2\2"+
		"\2\2\31\3\2\2\2\2\33\3\2\2\2\2\35\3\2\2\2\2\37\3\2\2\2\2!\3\2\2\2\2#\3"+
		"\2\2\2\3&\3\2\2\2\5,\3\2\2\2\7.\3\2\2\2\t\63\3\2\2\2\138\3\2\2\2\r:\3"+
		"\2\2\2\17<\3\2\2\2\21>\3\2\2\2\23@\3\2\2\2\25B\3\2\2\2\27D\3\2\2\2\31"+
		"F\3\2\2\2\33H\3\2\2\2\35J\3\2\2\2\37Q\3\2\2\2!g\3\2\2\2#l\3\2\2\2%\'\t"+
		"\2\2\2&%\3\2\2\2\'(\3\2\2\2(&\3\2\2\2()\3\2\2\2)*\3\2\2\2*+\b\2\2\2+\4"+
		"\3\2\2\2,-\7.\2\2-\6\3\2\2\2./\7h\2\2/\60\7w\2\2\60\61\7p\2\2\61\62\7"+
		"e\2\2\62\b\3\2\2\2\63\64\7v\2\2\64\65\7w\2\2\65\66\7t\2\2\66\67\7p\2\2"+
		"\67\n\3\2\2\289\7,\2\29\f\3\2\2\2:;\7\61\2\2;\16\3\2\2\2<=\7-\2\2=\20"+
		"\3\2\2\2>?\7/\2\2?\22\3\2\2\2@A\7?\2\2A\24\3\2\2\2BC\7*\2\2C\26\3\2\2"+
		"\2DE\7+\2\2E\30\3\2\2\2FG\7}\2\2G\32\3\2\2\2HI\7\177\2\2I\34\3\2\2\2J"+
		"N\t\3\2\2KM\t\4\2\2LK\3\2\2\2MP\3\2\2\2NL\3\2\2\2NO\3\2\2\2O\36\3\2\2"+
		"\2PN\3\2\2\2QU\t\3\2\2RT\t\4\2\2SR\3\2\2\2TW\3\2\2\2US\3\2\2\2UV\3\2\2"+
		"\2VX\3\2\2\2WU\3\2\2\2XZ\7\60\2\2Y[\t\4\2\2ZY\3\2\2\2[\\\3\2\2\2\\Z\3"+
		"\2\2\2\\]\3\2\2\2] \3\2\2\2^_\7v\2\2_`\7t\2\2`a\7w\2\2ah\7g\2\2bc\7h\2"+
		"\2cd\7c\2\2de\7n\2\2ef\7u\2\2fh\7g\2\2g^\3\2\2\2gb\3\2\2\2h\"\3\2\2\2"+
		"ik\t\5\2\2ji\3\2\2\2kn\3\2\2\2lj\3\2\2\2lm\3\2\2\2mo\3\2\2\2nl\3\2\2\2"+
		"os\t\6\2\2pr\t\7\2\2qp\3\2\2\2ru\3\2\2\2sq\3\2\2\2st\3\2\2\2t$\3\2\2\2"+
		"us\3\2\2\2\n\2(NU\\gls\3\b\2\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}