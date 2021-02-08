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
		WS=1, FUNC=2, TURN=3, MUL=4, DIV=5, ADD=6, SUB=7, ASSIGN=8, LPAREN=9, 
		RPAREN=10, INT=11, FLOAT=12, BOOL=13, STRING=14, IDENT=15;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"WS", "FUNC", "TURN", "MUL", "DIV", "ADD", "SUB", "ASSIGN", "LPAREN", 
			"RPAREN", "INT", "FLOAT", "BOOL", "STRING", "IDENT"
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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2\21t\b\1\4\2\t\2\4"+
		"\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t"+
		"\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\3\2\6\2#\n\2\r\2\16"+
		"\2$\3\2\3\2\3\3\3\3\3\3\3\3\3\3\3\4\3\4\3\4\3\4\3\4\3\5\3\5\3\6\3\6\3"+
		"\7\3\7\3\b\3\b\3\t\3\t\3\n\3\n\3\13\3\13\3\f\3\f\7\fC\n\f\f\f\16\fF\13"+
		"\f\3\r\3\r\7\rJ\n\r\f\r\16\rM\13\r\3\r\3\r\6\rQ\n\r\r\r\16\rR\3\16\3\16"+
		"\3\16\3\16\3\16\3\16\3\16\3\16\3\16\5\16^\n\16\3\17\3\17\6\17b\n\17\r"+
		"\17\16\17c\3\17\3\17\3\20\7\20i\n\20\f\20\16\20l\13\20\3\20\3\20\7\20"+
		"p\n\20\f\20\16\20s\13\20\2\2\21\3\3\5\4\7\5\t\6\13\7\r\b\17\t\21\n\23"+
		"\13\25\f\27\r\31\16\33\17\35\20\37\21\3\2\t\5\2\13\f\17\17\"\"\3\2\63"+
		";\3\2\62;\5\2\62;C\\c|\3\2aa\3\2c|\6\2\62;C\\aac|\2{\2\3\3\2\2\2\2\5\3"+
		"\2\2\2\2\7\3\2\2\2\2\t\3\2\2\2\2\13\3\2\2\2\2\r\3\2\2\2\2\17\3\2\2\2\2"+
		"\21\3\2\2\2\2\23\3\2\2\2\2\25\3\2\2\2\2\27\3\2\2\2\2\31\3\2\2\2\2\33\3"+
		"\2\2\2\2\35\3\2\2\2\2\37\3\2\2\2\3\"\3\2\2\2\5(\3\2\2\2\7-\3\2\2\2\t\62"+
		"\3\2\2\2\13\64\3\2\2\2\r\66\3\2\2\2\178\3\2\2\2\21:\3\2\2\2\23<\3\2\2"+
		"\2\25>\3\2\2\2\27@\3\2\2\2\31G\3\2\2\2\33]\3\2\2\2\35_\3\2\2\2\37j\3\2"+
		"\2\2!#\t\2\2\2\"!\3\2\2\2#$\3\2\2\2$\"\3\2\2\2$%\3\2\2\2%&\3\2\2\2&\'"+
		"\b\2\2\2\'\4\3\2\2\2()\7h\2\2)*\7w\2\2*+\7p\2\2+,\7e\2\2,\6\3\2\2\2-."+
		"\7v\2\2./\7w\2\2/\60\7t\2\2\60\61\7p\2\2\61\b\3\2\2\2\62\63\7,\2\2\63"+
		"\n\3\2\2\2\64\65\7\61\2\2\65\f\3\2\2\2\66\67\7-\2\2\67\16\3\2\2\289\7"+
		"/\2\29\20\3\2\2\2:;\7?\2\2;\22\3\2\2\2<=\7*\2\2=\24\3\2\2\2>?\7+\2\2?"+
		"\26\3\2\2\2@D\t\3\2\2AC\t\4\2\2BA\3\2\2\2CF\3\2\2\2DB\3\2\2\2DE\3\2\2"+
		"\2E\30\3\2\2\2FD\3\2\2\2GK\t\3\2\2HJ\t\4\2\2IH\3\2\2\2JM\3\2\2\2KI\3\2"+
		"\2\2KL\3\2\2\2LN\3\2\2\2MK\3\2\2\2NP\7\60\2\2OQ\t\4\2\2PO\3\2\2\2QR\3"+
		"\2\2\2RP\3\2\2\2RS\3\2\2\2S\32\3\2\2\2TU\7v\2\2UV\7t\2\2VW\7w\2\2W^\7"+
		"g\2\2XY\7h\2\2YZ\7c\2\2Z[\7n\2\2[\\\7u\2\2\\^\7g\2\2]T\3\2\2\2]X\3\2\2"+
		"\2^\34\3\2\2\2_a\7$\2\2`b\t\5\2\2a`\3\2\2\2bc\3\2\2\2ca\3\2\2\2cd\3\2"+
		"\2\2de\3\2\2\2ef\7$\2\2f\36\3\2\2\2gi\t\6\2\2hg\3\2\2\2il\3\2\2\2jh\3"+
		"\2\2\2jk\3\2\2\2km\3\2\2\2lj\3\2\2\2mq\t\7\2\2np\t\b\2\2on\3\2\2\2ps\3"+
		"\2\2\2qo\3\2\2\2qr\3\2\2\2r \3\2\2\2sq\3\2\2\2\13\2$DKR]cjq\3\b\2\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}