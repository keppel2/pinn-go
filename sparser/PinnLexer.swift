// Generated from Pinn.g4 by ANTLR 4.7.2
import Antlr4

open class PinnLexer: Lexer {

	internal static var _decisionToDFA: [DFA] = {
          var decisionToDFA = [DFA]()
          let length = PinnLexer._ATN.getNumberOfDecisions()
          for i in 0..<length {
          	    decisionToDFA.append(DFA(PinnLexer._ATN.getDecisionState(i)!, i))
          }
           return decisionToDFA
     }()

	internal static let _sharedContextCache = PredictionContextCache()

	public
	static let T__0=1, T__1=2, T__2=3, T__3=4, T__4=5, T__5=6, T__6=7, T__7=8, 
            T__8=9, T__9=10, T__10=11, T__11=12, T__12=13, T__13=14, T__14=15, 
            T__15=16, T__16=17, T__17=18, T__18=19, T__19=20, T__20=21, 
            T__21=22, T__22=23, T__23=24, T__24=25, T__25=26, T__26=27, 
            T__27=28, T__28=29, T__29=30, T__30=31, T__31=32, T__32=33, 
            T__33=34, T__34=35, T__35=36, T__36=37, T__37=38, T__38=39, 
            T__39=40, T__40=41, T__41=42, T__42=43, T__43=44, TWODOTS=45, 
            LSQUARE=46, LPAREN=47, MAP=48, SLICE=49, FILL=50, TYPES=51, 
            RANGE=52, COLON=53, CE=54, IOTA=55, BINOP=56, BOOL=57, ID=58, 
            CHAR=59, INT=60, FLOAT=61, WS=62, STRING=63

	public
	static let channelNames: [String] = [
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	]

	public
	static let modeNames: [String] = [
		"DEFAULT_MODE"
	]

	public
	static let ruleNames: [String] = [
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8", 
		"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16", 
		"T__17", "T__18", "T__19", "T__20", "T__21", "T__22", "T__23", "T__24", 
		"T__25", "T__26", "T__27", "T__28", "T__29", "T__30", "T__31", "T__32", 
		"T__33", "T__34", "T__35", "T__36", "T__37", "T__38", "T__39", "T__40", 
		"T__41", "T__42", "T__43", "TWODOTS", "LSQUARE", "LPAREN", "MAP", "SLICE", 
		"FILL", "TYPES", "RANGE", "COLON", "CE", "IOTA", "BINOP", "BOOL", "ID", 
		"CHAR", "INT", "FLOAT", "WS", "STRING", "DECIMAL_DIGIT", "DECIMAL_DIGITS", 
		"DECIMAL_EXPONENT", "HEX_DIGIT", "HEX_DIGITS", "HEX_MANTISSA", "HEX_EXPONENT", 
		"OCTAL_DIGIT", "BINARY_DIGIT"
	]

	private static let _LITERAL_NAMES: [String?] = [
		nil, "';'", "'func'", "','", "')'", "'{'", "'}'", "'='", "']'", "'++'", 
		"'--'", "'print'", "'printB'", "'printH'", "'delete'", "'len'", "'strLen'", 
		"'stringValue'", "'+'", "'-'", "'!'", "'^'", "'=='", "'!='", "'>'", "'<'", 
		"'>='", "'<='", "'&&'", "'||'", "'?'", "'return'", "'if'", "'else'", "'guard'", 
		"'while'", "'repeat'", "'for'", "'case'", "'switch'", "'default'", "'break'", 
		"'continue'", "'fallthrough'", "'debug'", "'@'", "'['", "'('", "'map'", 
		"'slice'", "'fill'", nil, "'range'", "':'", "':='", "'iota'"
	]
	private static let _SYMBOLIC_NAMES: [String?] = [
		nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, 
		nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, 
		nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, 
		nil, nil, nil, "TWODOTS", "LSQUARE", "LPAREN", "MAP", "SLICE", "FILL", 
		"TYPES", "RANGE", "COLON", "CE", "IOTA", "BINOP", "BOOL", "ID", "CHAR", 
		"INT", "FLOAT", "WS", "STRING"
	]
	public
	static let VOCABULARY = Vocabulary(_LITERAL_NAMES, _SYMBOLIC_NAMES)


	override open
	func getVocabulary() -> Vocabulary {
		return PinnLexer.VOCABULARY
	}

	public
	required init(_ input: CharStream) {
	    RuntimeMetaData.checkVersion("4.7.2", RuntimeMetaData.VERSION)
		super.init(input)
		_interp = LexerATNSimulator(self, PinnLexer._ATN, PinnLexer._decisionToDFA, PinnLexer._sharedContextCache)
	}

	override open
	func getGrammarFileName() -> String { return "Pinn.g4" }

	override open
	func getRuleNames() -> [String] { return PinnLexer.ruleNames }

	override open
	func getSerializedATN() -> String { return PinnLexer._serializedATN }

	override open
	func getChannelNames() -> [String] { return PinnLexer.channelNames }

	override open
	func getModeNames() -> [String] { return PinnLexer.modeNames }

	override open
	func getATN() -> ATN { return PinnLexer._ATN }


	public
	static let _serializedATN: String = PinnLexerATN().jsonString

	public
	static let _ATN: ATN = ATNDeserializer().deserializeFromJson(_serializedATN)
}