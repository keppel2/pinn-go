// Generated from Pinn.g4 by ANTLR 4.7.2
import Antlr4

open class PinnParser: Parser {

	internal static var _decisionToDFA: [DFA] = {
          var decisionToDFA = [DFA]()
          let length = PinnParser._ATN.getNumberOfDecisions()
          for i in 0..<length {
            decisionToDFA.append(DFA(PinnParser._ATN.getDecisionState(i)!, i))
           }
           return decisionToDFA
     }()

	internal static let _sharedContextCache = PredictionContextCache()

	public
	enum Tokens: Int {
		case EOF = -1, T__0 = 1, T__1 = 2, T__2 = 3, T__3 = 4, T__4 = 5, T__5 = 6, 
                 T__6 = 7, T__7 = 8, T__8 = 9, T__9 = 10, T__10 = 11, T__11 = 12, 
                 T__12 = 13, T__13 = 14, T__14 = 15, T__15 = 16, T__16 = 17, 
                 T__17 = 18, T__18 = 19, T__19 = 20, T__20 = 21, T__21 = 22, 
                 T__22 = 23, T__23 = 24, T__24 = 25, T__25 = 26, T__26 = 27, 
                 T__27 = 28, T__28 = 29, T__29 = 30, T__30 = 31, T__31 = 32, 
                 T__32 = 33, T__33 = 34, T__34 = 35, T__35 = 36, T__36 = 37, 
                 T__37 = 38, T__38 = 39, T__39 = 40, T__40 = 41, T__41 = 42, 
                 T__42 = 43, T__43 = 44, TWODOTS = 45, LSQUARE = 46, LPAREN = 47, 
                 MAP = 48, SLICE = 49, FILL = 50, TYPES = 51, RANGE = 52, 
                 COLON = 53, CE = 54, IOTA = 55, BINOP = 56, BOOL = 57, 
                 ID = 58, CHAR = 59, INT = 60, FLOAT = 61, WS = 62, STRING = 63
	}

	public
	static let RULE_file = 0, RULE_function = 1, RULE_testRule = 2, RULE_block = 3, 
            RULE_fvarDecl = 4, RULE_varDecl = 5, RULE_kind = 6, RULE_simpleStatement = 7, 
            RULE_indexExpr = 8, RULE_funcExpr = 9, RULE_expr = 10, RULE_exprList = 11, 
            RULE_returnStatement = 12, RULE_ifStatement = 13, RULE_guardStatement = 14, 
            RULE_whStatement = 15, RULE_repeatStatement = 16, RULE_foStatement = 17, 
            RULE_caseStatement = 18, RULE_switchStatement = 19, RULE_statement = 20, 
            RULE_pset = 21

	public
	static let ruleNames: [String] = [
		"file", "function", "testRule", "block", "fvarDecl", "varDecl", "kind", 
		"simpleStatement", "indexExpr", "funcExpr", "expr", "exprList", "returnStatement", 
		"ifStatement", "guardStatement", "whStatement", "repeatStatement", "foStatement", 
		"caseStatement", "switchStatement", "statement", "pset"
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
	func getGrammarFileName() -> String { return "Pinn.g4" }

	override open
	func getRuleNames() -> [String] { return PinnParser.ruleNames }

	override open
	func getSerializedATN() -> String { return PinnParser._serializedATN }

	override open
	func getATN() -> ATN { return PinnParser._ATN }


	override open
	func getVocabulary() -> Vocabulary {
	    return PinnParser.VOCABULARY
	}

	override public
	init(_ input:TokenStream) throws {
	    RuntimeMetaData.checkVersion("4.7.2", RuntimeMetaData.VERSION)
		try super.init(input)
		_interp = ParserATNSimulator(self,PinnParser._ATN,PinnParser._decisionToDFA, PinnParser._sharedContextCache)
	}


	public class FileContext: ParserRuleContext {
			open
			func function() -> [FunctionContext] {
				return getRuleContexts(FunctionContext.self)
			}
			open
			func function(_ i: Int) -> FunctionContext? {
				return getRuleContext(FunctionContext.self, i)
			}
			open
			func varDecl() -> [VarDeclContext] {
				return getRuleContexts(VarDeclContext.self)
			}
			open
			func varDecl(_ i: Int) -> VarDeclContext? {
				return getRuleContext(VarDeclContext.self, i)
			}
		override open
		func getRuleIndex() -> Int {
			return PinnParser.RULE_file
		}
	}
	@discardableResult
	 open func file() throws -> FileContext {
		var _localctx: FileContext = FileContext(_ctx, getState())
		try enterRule(_localctx, 0, PinnParser.RULE_file)
		var _la: Int = 0
		defer {
	    		try! exitRule()
	    }
		do {
		 	try enterOuterAlt(_localctx, 1)
		 	setState(48) 
		 	try _errHandler.sync(self)
		 	_la = try _input.LA(1)
		 	repeat {
		 		setState(48)
		 		try _errHandler.sync(self)
		 		switch (PinnParser.Tokens(rawValue: try _input.LA(1))!) {
		 		case .T__1:
		 			setState(44)
		 			try function()

		 			break

		 		case .ID:
		 			setState(45)
		 			try varDecl()
		 			setState(46)
		 			try match(PinnParser.Tokens.T__0.rawValue)

		 			break
		 		default:
		 			throw ANTLRException.recognition(e: NoViableAltException(self))
		 		}

		 		setState(50); 
		 		try _errHandler.sync(self)
		 		_la = try _input.LA(1)
		 	} while (//closure
		 	 { () -> Bool in
		 	      let testSet: Bool = _la == PinnParser.Tokens.T__1.rawValue || _la == PinnParser.Tokens.ID.rawValue
		 	      return testSet
		 	 }())

		}
		catch ANTLRException.recognition(let re) {
			_localctx.exception = re
			_errHandler.reportError(self, re)
			try _errHandler.recover(self, re)
		}

		return _localctx
	}

	public class FunctionContext: ParserRuleContext {
			open
			func ID() -> TerminalNode? {
				return getToken(PinnParser.Tokens.ID.rawValue, 0)
			}
			open
			func LPAREN() -> TerminalNode? {
				return getToken(PinnParser.Tokens.LPAREN.rawValue, 0)
			}
			open
			func block() -> BlockContext? {
				return getRuleContext(BlockContext.self, 0)
			}
			open
			func fvarDecl() -> [FvarDeclContext] {
				return getRuleContexts(FvarDeclContext.self)
			}
			open
			func fvarDecl(_ i: Int) -> FvarDeclContext? {
				return getRuleContext(FvarDeclContext.self, i)
			}
			open
			func kind() -> KindContext? {
				return getRuleContext(KindContext.self, 0)
			}
		override open
		func getRuleIndex() -> Int {
			return PinnParser.RULE_function
		}
	}
	@discardableResult
	 open func function() throws -> FunctionContext {
		var _localctx: FunctionContext = FunctionContext(_ctx, getState())
		try enterRule(_localctx, 2, PinnParser.RULE_function)
		var _la: Int = 0
		defer {
	    		try! exitRule()
	    }
		do {
		 	try enterOuterAlt(_localctx, 1)
		 	setState(52)
		 	try match(PinnParser.Tokens.T__1.rawValue)
		 	setState(53)
		 	try match(PinnParser.Tokens.ID.rawValue)
		 	setState(54)
		 	try match(PinnParser.Tokens.LPAREN.rawValue)
		 	setState(63)
		 	try _errHandler.sync(self)
		 	_la = try _input.LA(1)
		 	if (//closure
		 	 { () -> Bool in
		 	      let testSet: Bool = _la == PinnParser.Tokens.ID.rawValue
		 	      return testSet
		 	 }()) {
		 		setState(55)
		 		try fvarDecl()
		 		setState(60)
		 		try _errHandler.sync(self)
		 		_la = try _input.LA(1)
		 		while (//closure
		 		 { () -> Bool in
		 		      let testSet: Bool = _la == PinnParser.Tokens.T__2.rawValue
		 		      return testSet
		 		 }()) {
		 			setState(56)
		 			try match(PinnParser.Tokens.T__2.rawValue)
		 			setState(57)
		 			try fvarDecl()


		 			setState(62)
		 			try _errHandler.sync(self)
		 			_la = try _input.LA(1)
		 		}

		 	}

		 	setState(65)
		 	try match(PinnParser.Tokens.T__3.rawValue)
		 	setState(67)
		 	try _errHandler.sync(self)
		 	_la = try _input.LA(1)
		 	if (//closure
		 	 { () -> Bool in
		 	      let testSet: Bool = _la == PinnParser.Tokens.LSQUARE.rawValue || _la == PinnParser.Tokens.TYPES.rawValue
		 	      return testSet
		 	 }()) {
		 		setState(66)
		 		try kind()

		 	}

		 	setState(69)
		 	try block()

		}
		catch ANTLRException.recognition(let re) {
			_localctx.exception = re
			_errHandler.reportError(self, re)
			try _errHandler.recover(self, re)
		}

		return _localctx
	}

	public class TestRuleContext: ParserRuleContext {
			open
			func ID() -> [TerminalNode] {
				return getTokens(PinnParser.Tokens.ID.rawValue)
			}
			open
			func ID(_ i:Int) -> TerminalNode? {
				return getToken(PinnParser.Tokens.ID.rawValue, i)
			}
		override open
		func getRuleIndex() -> Int {
			return PinnParser.RULE_testRule
		}
	}
	@discardableResult
	 open func testRule() throws -> TestRuleContext {
		var _localctx: TestRuleContext = TestRuleContext(_ctx, getState())
		try enterRule(_localctx, 4, PinnParser.RULE_testRule)
		defer {
	    		try! exitRule()
	    }
		do {
		 	try enterOuterAlt(_localctx, 1)
		 	setState(71)
		 	try match(PinnParser.Tokens.ID.rawValue)
		 	setState(72)
		 	try match(PinnParser.Tokens.ID.rawValue)

		}
		catch ANTLRException.recognition(let re) {
			_localctx.exception = re
			_errHandler.reportError(self, re)
			try _errHandler.recover(self, re)
		}

		return _localctx
	}

	public class BlockContext: ParserRuleContext {
			open
			func statement() -> [StatementContext] {
				return getRuleContexts(StatementContext.self)
			}
			open
			func statement(_ i: Int) -> StatementContext? {
				return getRuleContext(StatementContext.self, i)
			}
		override open
		func getRuleIndex() -> Int {
			return PinnParser.RULE_block
		}
	}
	@discardableResult
	 open func block() throws -> BlockContext {
		var _localctx: BlockContext = BlockContext(_ctx, getState())
		try enterRule(_localctx, 6, PinnParser.RULE_block)
		var _la: Int = 0
		defer {
	    		try! exitRule()
	    }
		do {
		 	try enterOuterAlt(_localctx, 1)
		 	setState(74)
		 	try match(PinnParser.Tokens.T__4.rawValue)
		 	setState(78)
		 	try _errHandler.sync(self)
		 	_la = try _input.LA(1)
		 	while (//closure
		 	 { () -> Bool in
		 	      let testSet: Bool = {  () -> Bool in
		 	   let testArray: [Int] = [_la, PinnParser.Tokens.T__0.rawValue,PinnParser.Tokens.T__4.rawValue,PinnParser.Tokens.T__10.rawValue,PinnParser.Tokens.T__11.rawValue,PinnParser.Tokens.T__12.rawValue,PinnParser.Tokens.T__13.rawValue,PinnParser.Tokens.T__14.rawValue,PinnParser.Tokens.T__15.rawValue,PinnParser.Tokens.T__16.rawValue,PinnParser.Tokens.T__17.rawValue,PinnParser.Tokens.T__18.rawValue,PinnParser.Tokens.T__19.rawValue,PinnParser.Tokens.T__20.rawValue,PinnParser.Tokens.T__30.rawValue,PinnParser.Tokens.T__31.rawValue,PinnParser.Tokens.T__33.rawValue,PinnParser.Tokens.T__34.rawValue,PinnParser.Tokens.T__35.rawValue,PinnParser.Tokens.T__36.rawValue,PinnParser.Tokens.T__38.rawValue,PinnParser.Tokens.T__40.rawValue,PinnParser.Tokens.T__41.rawValue,PinnParser.Tokens.T__42.rawValue,PinnParser.Tokens.T__43.rawValue,PinnParser.Tokens.LPAREN.rawValue,PinnParser.Tokens.IOTA.rawValue,PinnParser.Tokens.BOOL.rawValue,PinnParser.Tokens.ID.rawValue,PinnParser.Tokens.CHAR.rawValue,PinnParser.Tokens.INT.rawValue,PinnParser.Tokens.FLOAT.rawValue,PinnParser.Tokens.STRING.rawValue]
		 	    return  Utils.testBitLeftShiftArray(testArray, 0)
		 	}()
		 	      return testSet
		 	 }()) {
		 		setState(75)
		 		try statement()


		 		setState(80)
		 		try _errHandler.sync(self)
		 		_la = try _input.LA(1)
		 	}
		 	setState(81)
		 	try match(PinnParser.Tokens.T__5.rawValue)

		}
		catch ANTLRException.recognition(let re) {
			_localctx.exception = re
			_errHandler.reportError(self, re)
			try _errHandler.recover(self, re)
		}

		return _localctx
	}

	public class FvarDeclContext: ParserRuleContext {
			open
			func ID() -> TerminalNode? {
				return getToken(PinnParser.Tokens.ID.rawValue, 0)
			}
			open
			func kind() -> KindContext? {
				return getRuleContext(KindContext.self, 0)
			}
		override open
		func getRuleIndex() -> Int {
			return PinnParser.RULE_fvarDecl
		}
	}
	@discardableResult
	 open func fvarDecl() throws -> FvarDeclContext {
		var _localctx: FvarDeclContext = FvarDeclContext(_ctx, getState())
		try enterRule(_localctx, 8, PinnParser.RULE_fvarDecl)
		defer {
	    		try! exitRule()
	    }
		do {
		 	try enterOuterAlt(_localctx, 1)
		 	setState(83)
		 	try match(PinnParser.Tokens.ID.rawValue)
		 	setState(84)
		 	try kind()

		}
		catch ANTLRException.recognition(let re) {
			_localctx.exception = re
			_errHandler.reportError(self, re)
			try _errHandler.recover(self, re)
		}

		return _localctx
	}

	public class VarDeclContext: ParserRuleContext {
			open
			func ID() -> TerminalNode? {
				return getToken(PinnParser.Tokens.ID.rawValue, 0)
			}
			open
			func kind() -> KindContext? {
				return getRuleContext(KindContext.self, 0)
			}
			open
			func exprList() -> ExprListContext? {
				return getRuleContext(ExprListContext.self, 0)
			}
			open
			func CE() -> TerminalNode? {
				return getToken(PinnParser.Tokens.CE.rawValue, 0)
			}
			open
			func expr() -> ExprContext? {
				return getRuleContext(ExprContext.self, 0)
			}
		override open
		func getRuleIndex() -> Int {
			return PinnParser.RULE_varDecl
		}
	}
	@discardableResult
	 open func varDecl() throws -> VarDeclContext {
		var _localctx: VarDeclContext = VarDeclContext(_ctx, getState())
		try enterRule(_localctx, 10, PinnParser.RULE_varDecl)
		var _la: Int = 0
		defer {
	    		try! exitRule()
	    }
		do {
		 	setState(95)
		 	try _errHandler.sync(self)
		 	switch(try getInterpreter().adaptivePredict(_input,7, _ctx)) {
		 	case 1:
		 		try enterOuterAlt(_localctx, 1)
		 		setState(86)
		 		try match(PinnParser.Tokens.ID.rawValue)
		 		setState(87)
		 		try kind()
		 		setState(90)
		 		try _errHandler.sync(self)
		 		_la = try _input.LA(1)
		 		if (//closure
		 		 { () -> Bool in
		 		      let testSet: Bool = _la == PinnParser.Tokens.T__6.rawValue
		 		      return testSet
		 		 }()) {
		 			setState(88)
		 			try match(PinnParser.Tokens.T__6.rawValue)
		 			setState(89)
		 			try exprList()

		 		}


		 		break
		 	case 2:
		 		try enterOuterAlt(_localctx, 2)
		 		setState(92)
		 		try match(PinnParser.Tokens.ID.rawValue)
		 		setState(93)
		 		try match(PinnParser.Tokens.CE.rawValue)
		 		setState(94)
		 		try expr(0)

		 		break
		 	default: break
		 	}
		}
		catch ANTLRException.recognition(let re) {
			_localctx.exception = re
			_errHandler.reportError(self, re)
			try _errHandler.recover(self, re)
		}

		return _localctx
	}

	public class KindContext: ParserRuleContext {
			open
			func TYPES() -> TerminalNode? {
				return getToken(PinnParser.Tokens.TYPES.rawValue, 0)
			}
			open
			func LSQUARE() -> TerminalNode? {
				return getToken(PinnParser.Tokens.LSQUARE.rawValue, 0)
			}
			open
			func MAP() -> TerminalNode? {
				return getToken(PinnParser.Tokens.MAP.rawValue, 0)
			}
			open
			func SLICE() -> TerminalNode? {
				return getToken(PinnParser.Tokens.SLICE.rawValue, 0)
			}
			open
			func FILL() -> TerminalNode? {
				return getToken(PinnParser.Tokens.FILL.rawValue, 0)
			}
			open
			func expr() -> ExprContext? {
				return getRuleContext(ExprContext.self, 0)
			}
		override open
		func getRuleIndex() -> Int {
			return PinnParser.RULE_kind
		}
	}
	@discardableResult
	 open func kind() throws -> KindContext {
		var _localctx: KindContext = KindContext(_ctx, getState())
		try enterRule(_localctx, 12, PinnParser.RULE_kind)
		var _la: Int = 0
		defer {
	    		try! exitRule()
	    }
		do {
		 	try enterOuterAlt(_localctx, 1)
		 	setState(105)
		 	try _errHandler.sync(self)
		 	_la = try _input.LA(1)
		 	if (//closure
		 	 { () -> Bool in
		 	      let testSet: Bool = _la == PinnParser.Tokens.LSQUARE.rawValue
		 	      return testSet
		 	 }()) {
		 		setState(97)
		 		try match(PinnParser.Tokens.LSQUARE.rawValue)
		 		setState(102)
		 		try _errHandler.sync(self)
		 		switch (PinnParser.Tokens(rawValue: try _input.LA(1))!) {
		 		case .MAP:
		 			setState(98)
		 			try match(PinnParser.Tokens.MAP.rawValue)

		 			break

		 		case .SLICE:
		 			setState(99)
		 			try match(PinnParser.Tokens.SLICE.rawValue)

		 			break

		 		case .FILL:
		 			setState(100)
		 			try match(PinnParser.Tokens.FILL.rawValue)

		 			break
		 		case .T__10:fallthrough
		 		case .T__11:fallthrough
		 		case .T__12:fallthrough
		 		case .T__13:fallthrough
		 		case .T__14:fallthrough
		 		case .T__15:fallthrough
		 		case .T__16:fallthrough
		 		case .T__17:fallthrough
		 		case .T__18:fallthrough
		 		case .T__19:fallthrough
		 		case .T__20:fallthrough
		 		case .LPAREN:fallthrough
		 		case .IOTA:fallthrough
		 		case .BOOL:fallthrough
		 		case .ID:fallthrough
		 		case .CHAR:fallthrough
		 		case .INT:fallthrough
		 		case .FLOAT:fallthrough
		 		case .STRING:
		 			setState(101)
		 			try expr(0)

		 			break
		 		default:
		 			throw ANTLRException.recognition(e: NoViableAltException(self))
		 		}
		 		setState(104)
		 		try match(PinnParser.Tokens.T__7.rawValue)

		 	}

		 	setState(107)
		 	try match(PinnParser.Tokens.TYPES.rawValue)

		}
		catch ANTLRException.recognition(let re) {
			_localctx.exception = re
			_errHandler.reportError(self, re)
			try _errHandler.recover(self, re)
		}

		return _localctx
	}

	public class SimpleStatementContext: ParserRuleContext {
			open
			func pset() -> PsetContext? {
				return getRuleContext(PsetContext.self, 0)
			}
			open
			func ID() -> TerminalNode? {
				return getToken(PinnParser.Tokens.ID.rawValue, 0)
			}
			open
			func LSQUARE() -> TerminalNode? {
				return getToken(PinnParser.Tokens.LSQUARE.rawValue, 0)
			}
			open
			func expr() -> ExprContext? {
				return getRuleContext(ExprContext.self, 0)
			}
		override open
		func getRuleIndex() -> Int {
			return PinnParser.RULE_simpleStatement
		}
	}
	@discardableResult
	 open func simpleStatement() throws -> SimpleStatementContext {
		var _localctx: SimpleStatementContext = SimpleStatementContext(_ctx, getState())
		try enterRule(_localctx, 14, PinnParser.RULE_simpleStatement)
		var _la: Int = 0
		defer {
	    		try! exitRule()
	    }
		do {
		 	setState(118)
		 	try _errHandler.sync(self)
		 	switch(try getInterpreter().adaptivePredict(_input,11, _ctx)) {
		 	case 1:
		 		try enterOuterAlt(_localctx, 1)
		 		setState(109)
		 		try pset()

		 		break
		 	case 2:
		 		try enterOuterAlt(_localctx, 2)
		 		setState(110)
		 		try match(PinnParser.Tokens.ID.rawValue)
		 		setState(115)
		 		try _errHandler.sync(self)
		 		_la = try _input.LA(1)
		 		if (//closure
		 		 { () -> Bool in
		 		      let testSet: Bool = _la == PinnParser.Tokens.LSQUARE.rawValue
		 		      return testSet
		 		 }()) {
		 			setState(111)
		 			try match(PinnParser.Tokens.LSQUARE.rawValue)
		 			setState(112)
		 			try expr(0)
		 			setState(113)
		 			try match(PinnParser.Tokens.T__7.rawValue)

		 		}

		 		setState(117)
		 		_la = try _input.LA(1)
		 		if (!(//closure
		 		 { () -> Bool in
		 		      let testSet: Bool = _la == PinnParser.Tokens.T__8.rawValue || _la == PinnParser.Tokens.T__9.rawValue
		 		      return testSet
		 		 }())) {
		 		try _errHandler.recoverInline(self)
		 		}
		 		else {
		 			_errHandler.reportMatch(self)
		 			try consume()
		 		}

		 		break
		 	default: break
		 	}
		}
		catch ANTLRException.recognition(let re) {
			_localctx.exception = re
			_errHandler.reportError(self, re)
			try _errHandler.recover(self, re)
		}

		return _localctx
	}

	public class IndexExprContext: ParserRuleContext {
		open var first: ExprContext!
		open var second: ExprContext!
			open
			func ID() -> TerminalNode? {
				return getToken(PinnParser.Tokens.ID.rawValue, 0)
			}
			open
			func LSQUARE() -> TerminalNode? {
				return getToken(PinnParser.Tokens.LSQUARE.rawValue, 0)
			}
			open
			func TWODOTS() -> TerminalNode? {
				return getToken(PinnParser.Tokens.TWODOTS.rawValue, 0)
			}
			open
			func COLON() -> TerminalNode? {
				return getToken(PinnParser.Tokens.COLON.rawValue, 0)
			}
			open
			func expr() -> [ExprContext] {
				return getRuleContexts(ExprContext.self)
			}
			open
			func expr(_ i: Int) -> ExprContext? {
				return getRuleContext(ExprContext.self, i)
			}
		override open
		func getRuleIndex() -> Int {
			return PinnParser.RULE_indexExpr
		}
	}
	@discardableResult
	 open func indexExpr() throws -> IndexExprContext {
		var _localctx: IndexExprContext = IndexExprContext(_ctx, getState())
		try enterRule(_localctx, 16, PinnParser.RULE_indexExpr)
		var _la: Int = 0
		defer {
	    		try! exitRule()
	    }
		do {
		 	setState(135)
		 	try _errHandler.sync(self)
		 	switch(try getInterpreter().adaptivePredict(_input,14, _ctx)) {
		 	case 1:
		 		try enterOuterAlt(_localctx, 1)
		 		setState(120)
		 		try match(PinnParser.Tokens.ID.rawValue)
		 		setState(121)
		 		try match(PinnParser.Tokens.LSQUARE.rawValue)
		 		setState(123)
		 		try _errHandler.sync(self)
		 		_la = try _input.LA(1)
		 		if (//closure
		 		 { () -> Bool in
		 		      let testSet: Bool = {  () -> Bool in
		 		   let testArray: [Int] = [_la, PinnParser.Tokens.T__10.rawValue,PinnParser.Tokens.T__11.rawValue,PinnParser.Tokens.T__12.rawValue,PinnParser.Tokens.T__13.rawValue,PinnParser.Tokens.T__14.rawValue,PinnParser.Tokens.T__15.rawValue,PinnParser.Tokens.T__16.rawValue,PinnParser.Tokens.T__17.rawValue,PinnParser.Tokens.T__18.rawValue,PinnParser.Tokens.T__19.rawValue,PinnParser.Tokens.T__20.rawValue,PinnParser.Tokens.LPAREN.rawValue,PinnParser.Tokens.IOTA.rawValue,PinnParser.Tokens.BOOL.rawValue,PinnParser.Tokens.ID.rawValue,PinnParser.Tokens.CHAR.rawValue,PinnParser.Tokens.INT.rawValue,PinnParser.Tokens.FLOAT.rawValue,PinnParser.Tokens.STRING.rawValue]
		 		    return  Utils.testBitLeftShiftArray(testArray, 0)
		 		}()
		 		      return testSet
		 		 }()) {
		 			setState(122)
		 			try {
		 					let assignmentValue = try expr(0)
		 					_localctx.castdown(IndexExprContext.self).first = assignmentValue
		 			     }()


		 		}

		 		setState(125)
		 		_la = try _input.LA(1)
		 		if (!(//closure
		 		 { () -> Bool in
		 		      let testSet: Bool = _la == PinnParser.Tokens.TWODOTS.rawValue || _la == PinnParser.Tokens.COLON.rawValue
		 		      return testSet
		 		 }())) {
		 		try _errHandler.recoverInline(self)
		 		}
		 		else {
		 			_errHandler.reportMatch(self)
		 			try consume()
		 		}
		 		setState(127)
		 		try _errHandler.sync(self)
		 		_la = try _input.LA(1)
		 		if (//closure
		 		 { () -> Bool in
		 		      let testSet: Bool = {  () -> Bool in
		 		   let testArray: [Int] = [_la, PinnParser.Tokens.T__10.rawValue,PinnParser.Tokens.T__11.rawValue,PinnParser.Tokens.T__12.rawValue,PinnParser.Tokens.T__13.rawValue,PinnParser.Tokens.T__14.rawValue,PinnParser.Tokens.T__15.rawValue,PinnParser.Tokens.T__16.rawValue,PinnParser.Tokens.T__17.rawValue,PinnParser.Tokens.T__18.rawValue,PinnParser.Tokens.T__19.rawValue,PinnParser.Tokens.T__20.rawValue,PinnParser.Tokens.LPAREN.rawValue,PinnParser.Tokens.IOTA.rawValue,PinnParser.Tokens.BOOL.rawValue,PinnParser.Tokens.ID.rawValue,PinnParser.Tokens.CHAR.rawValue,PinnParser.Tokens.INT.rawValue,PinnParser.Tokens.FLOAT.rawValue,PinnParser.Tokens.STRING.rawValue]
		 		    return  Utils.testBitLeftShiftArray(testArray, 0)
		 		}()
		 		      return testSet
		 		 }()) {
		 			setState(126)
		 			try {
		 					let assignmentValue = try expr(0)
		 					_localctx.castdown(IndexExprContext.self).second = assignmentValue
		 			     }()


		 		}

		 		setState(129)
		 		try match(PinnParser.Tokens.T__7.rawValue)

		 		break
		 	case 2:
		 		try enterOuterAlt(_localctx, 2)
		 		setState(130)
		 		try match(PinnParser.Tokens.ID.rawValue)
		 		setState(131)
		 		try match(PinnParser.Tokens.LSQUARE.rawValue)
		 		setState(132)
		 		try expr(0)
		 		setState(133)
		 		try match(PinnParser.Tokens.T__7.rawValue)

		 		break
		 	default: break
		 	}
		}
		catch ANTLRException.recognition(let re) {
			_localctx.exception = re
			_errHandler.reportError(self, re)
			try _errHandler.recover(self, re)
		}

		return _localctx
	}

	public class FuncExprContext: ParserRuleContext {
			open
			func LPAREN() -> TerminalNode? {
				return getToken(PinnParser.Tokens.LPAREN.rawValue, 0)
			}
			open
			func exprList() -> ExprListContext? {
				return getRuleContext(ExprListContext.self, 0)
			}
			open
			func expr() -> ExprContext? {
				return getRuleContext(ExprContext.self, 0)
			}
			open
			func ID() -> TerminalNode? {
				return getToken(PinnParser.Tokens.ID.rawValue, 0)
			}
		override open
		func getRuleIndex() -> Int {
			return PinnParser.RULE_funcExpr
		}
	}
	@discardableResult
	 open func funcExpr() throws -> FuncExprContext {
		var _localctx: FuncExprContext = FuncExprContext(_ctx, getState())
		try enterRule(_localctx, 18, PinnParser.RULE_funcExpr)
		defer {
	    		try! exitRule()
	    }
		do {
		 	setState(173)
		 	try _errHandler.sync(self)
		 	switch (PinnParser.Tokens(rawValue: try _input.LA(1))!) {
		 	case .T__10:
		 		try enterOuterAlt(_localctx, 1)
		 		setState(137)
		 		try match(PinnParser.Tokens.T__10.rawValue)
		 		setState(138)
		 		try match(PinnParser.Tokens.LPAREN.rawValue)
		 		setState(139)
		 		try exprList()
		 		setState(140)
		 		try match(PinnParser.Tokens.T__3.rawValue)

		 		break

		 	case .T__11:
		 		try enterOuterAlt(_localctx, 2)
		 		setState(142)
		 		try match(PinnParser.Tokens.T__11.rawValue)
		 		setState(143)
		 		try match(PinnParser.Tokens.LPAREN.rawValue)
		 		setState(144)
		 		try expr(0)
		 		setState(145)
		 		try match(PinnParser.Tokens.T__3.rawValue)

		 		break

		 	case .T__12:
		 		try enterOuterAlt(_localctx, 3)
		 		setState(147)
		 		try match(PinnParser.Tokens.T__12.rawValue)
		 		setState(148)
		 		try match(PinnParser.Tokens.LPAREN.rawValue)
		 		setState(149)
		 		try expr(0)
		 		setState(150)
		 		try match(PinnParser.Tokens.T__3.rawValue)

		 		break

		 	case .T__13:
		 		try enterOuterAlt(_localctx, 4)
		 		setState(152)
		 		try match(PinnParser.Tokens.T__13.rawValue)
		 		setState(153)
		 		try match(PinnParser.Tokens.LPAREN.rawValue)
		 		setState(154)
		 		try match(PinnParser.Tokens.ID.rawValue)
		 		setState(155)
		 		try match(PinnParser.Tokens.T__2.rawValue)
		 		setState(156)
		 		try expr(0)
		 		setState(157)
		 		try match(PinnParser.Tokens.T__3.rawValue)

		 		break

		 	case .T__14:
		 		try enterOuterAlt(_localctx, 5)
		 		setState(159)
		 		try match(PinnParser.Tokens.T__14.rawValue)
		 		setState(160)
		 		try match(PinnParser.Tokens.LPAREN.rawValue)
		 		setState(161)
		 		try match(PinnParser.Tokens.ID.rawValue)
		 		setState(162)
		 		try match(PinnParser.Tokens.T__3.rawValue)

		 		break

		 	case .T__15:
		 		try enterOuterAlt(_localctx, 6)
		 		setState(163)
		 		try match(PinnParser.Tokens.T__15.rawValue)
		 		setState(164)
		 		try match(PinnParser.Tokens.LPAREN.rawValue)
		 		setState(165)
		 		try expr(0)
		 		setState(166)
		 		try match(PinnParser.Tokens.T__3.rawValue)

		 		break

		 	case .T__16:
		 		try enterOuterAlt(_localctx, 7)
		 		setState(168)
		 		try match(PinnParser.Tokens.T__16.rawValue)
		 		setState(169)
		 		try match(PinnParser.Tokens.LPAREN.rawValue)
		 		setState(170)
		 		try expr(0)
		 		setState(171)
		 		try match(PinnParser.Tokens.T__3.rawValue)

		 		break
		 	default:
		 		throw ANTLRException.recognition(e: NoViableAltException(self))
		 	}
		}
		catch ANTLRException.recognition(let re) {
			_localctx.exception = re
			_errHandler.reportError(self, re)
			try _errHandler.recover(self, re)
		}

		return _localctx
	}


	public class ExprContext: ParserRuleContext {
		open var firstExpr: ExprContext!
		open var secondExpr: ExprContext!
			open
			func funcExpr() -> FuncExprContext? {
				return getRuleContext(FuncExprContext.self, 0)
			}
			open
			func ID() -> TerminalNode? {
				return getToken(PinnParser.Tokens.ID.rawValue, 0)
			}
			open
			func LPAREN() -> TerminalNode? {
				return getToken(PinnParser.Tokens.LPAREN.rawValue, 0)
			}
			open
			func exprList() -> ExprListContext? {
				return getRuleContext(ExprListContext.self, 0)
			}
			open
			func indexExpr() -> IndexExprContext? {
				return getRuleContext(IndexExprContext.self, 0)
			}
			open
			func expr() -> [ExprContext] {
				return getRuleContexts(ExprContext.self)
			}
			open
			func expr(_ i: Int) -> ExprContext? {
				return getRuleContext(ExprContext.self, i)
			}
			open
			func FLOAT() -> TerminalNode? {
				return getToken(PinnParser.Tokens.FLOAT.rawValue, 0)
			}
			open
			func INT() -> TerminalNode? {
				return getToken(PinnParser.Tokens.INT.rawValue, 0)
			}
			open
			func BOOL() -> TerminalNode? {
				return getToken(PinnParser.Tokens.BOOL.rawValue, 0)
			}
			open
			func STRING() -> TerminalNode? {
				return getToken(PinnParser.Tokens.STRING.rawValue, 0)
			}
			open
			func CHAR() -> TerminalNode? {
				return getToken(PinnParser.Tokens.CHAR.rawValue, 0)
			}
			open
			func IOTA() -> TerminalNode? {
				return getToken(PinnParser.Tokens.IOTA.rawValue, 0)
			}
			open
			func BINOP() -> TerminalNode? {
				return getToken(PinnParser.Tokens.BINOP.rawValue, 0)
			}
			open
			func COLON() -> TerminalNode? {
				return getToken(PinnParser.Tokens.COLON.rawValue, 0)
			}
			open
			func TWODOTS() -> TerminalNode? {
				return getToken(PinnParser.Tokens.TWODOTS.rawValue, 0)
			}
		override open
		func getRuleIndex() -> Int {
			return PinnParser.RULE_expr
		}
	}

	 public final  func expr( ) throws -> ExprContext   {
		return try expr(0)
	}
	@discardableResult
	private func expr(_ _p: Int) throws -> ExprContext   {
		let _parentctx: ParserRuleContext? = _ctx
		var _parentState: Int = getState()
		var _localctx: ExprContext = ExprContext(_ctx, _parentState)
		var  _prevctx: ExprContext = _localctx
		var _startState: Int = 20
		try enterRecursionRule(_localctx, 20, PinnParser.RULE_expr, _p)
		var _la: Int = 0
		defer {
	    		try! unrollRecursionContexts(_parentctx)
	    }
		do {
			var _alt: Int
			try enterOuterAlt(_localctx, 1)
			setState(197)
			try _errHandler.sync(self)
			switch(try getInterpreter().adaptivePredict(_input,17, _ctx)) {
			case 1:
				setState(176)
				try funcExpr()

				break
			case 2:
				setState(177)
				try match(PinnParser.Tokens.ID.rawValue)
				setState(178)
				try match(PinnParser.Tokens.LPAREN.rawValue)
				setState(180)
				try _errHandler.sync(self)
				_la = try _input.LA(1)
				if (//closure
				 { () -> Bool in
				      let testSet: Bool = {  () -> Bool in
				   let testArray: [Int] = [_la, PinnParser.Tokens.T__10.rawValue,PinnParser.Tokens.T__11.rawValue,PinnParser.Tokens.T__12.rawValue,PinnParser.Tokens.T__13.rawValue,PinnParser.Tokens.T__14.rawValue,PinnParser.Tokens.T__15.rawValue,PinnParser.Tokens.T__16.rawValue,PinnParser.Tokens.T__17.rawValue,PinnParser.Tokens.T__18.rawValue,PinnParser.Tokens.T__19.rawValue,PinnParser.Tokens.T__20.rawValue,PinnParser.Tokens.LPAREN.rawValue,PinnParser.Tokens.IOTA.rawValue,PinnParser.Tokens.BOOL.rawValue,PinnParser.Tokens.ID.rawValue,PinnParser.Tokens.CHAR.rawValue,PinnParser.Tokens.INT.rawValue,PinnParser.Tokens.FLOAT.rawValue,PinnParser.Tokens.STRING.rawValue]
				    return  Utils.testBitLeftShiftArray(testArray, 0)
				}()
				      return testSet
				 }()) {
					setState(179)
					try exprList()

				}

				setState(182)
				try match(PinnParser.Tokens.T__3.rawValue)

				break
			case 3:
				setState(183)
				try indexExpr()

				break
			case 4:
				setState(184)
				_la = try _input.LA(1)
				if (!(//closure
				 { () -> Bool in
				      let testSet: Bool = {  () -> Bool in
				   let testArray: [Int] = [_la, PinnParser.Tokens.T__17.rawValue,PinnParser.Tokens.T__18.rawValue,PinnParser.Tokens.T__19.rawValue,PinnParser.Tokens.T__20.rawValue]
				    return  Utils.testBitLeftShiftArray(testArray, 0)
				}()
				      return testSet
				 }())) {
				try _errHandler.recoverInline(self)
				}
				else {
					_errHandler.reportMatch(self)
					try consume()
				}
				setState(185)
				try expr(14)

				break
			case 5:
				setState(186)
				try match(PinnParser.Tokens.LPAREN.rawValue)
				setState(187)
				try expr(0)
				setState(188)
				try match(PinnParser.Tokens.T__3.rawValue)

				break
			case 6:
				setState(190)
				try match(PinnParser.Tokens.ID.rawValue)

				break
			case 7:
				setState(191)
				try match(PinnParser.Tokens.FLOAT.rawValue)

				break
			case 8:
				setState(192)
				try match(PinnParser.Tokens.INT.rawValue)

				break
			case 9:
				setState(193)
				try match(PinnParser.Tokens.BOOL.rawValue)

				break
			case 10:
				setState(194)
				try match(PinnParser.Tokens.STRING.rawValue)

				break
			case 11:
				setState(195)
				try match(PinnParser.Tokens.CHAR.rawValue)

				break
			case 12:
				setState(196)
				try match(PinnParser.Tokens.IOTA.rawValue)

				break
			default: break
			}
			_ctx!.stop = try _input.LT(-1)
			setState(219)
			try _errHandler.sync(self)
			_alt = try getInterpreter().adaptivePredict(_input,19,_ctx)
			while (_alt != 2 && _alt != ATN.INVALID_ALT_NUMBER) {
				if ( _alt==1 ) {
					if _parseListeners != nil {
					   try triggerExitRuleEvent()
					}
					_prevctx = _localctx
					setState(217)
					try _errHandler.sync(self)
					switch(try getInterpreter().adaptivePredict(_input,18, _ctx)) {
					case 1:
						_localctx = ExprContext(_parentctx, _parentState);
						try pushNewRecursionContext(_localctx, _startState, PinnParser.RULE_expr)
						setState(199)
						if (!(precpred(_ctx, 13))) {
						    throw ANTLRException.recognition(e:FailedPredicateException(self, "precpred(_ctx, 13)"))
						}
						setState(200)
						_la = try _input.LA(1)
						if (!(//closure
						 { () -> Bool in
						      let testSet: Bool = {  () -> Bool in
						   let testArray: [Int] = [_la, PinnParser.Tokens.T__17.rawValue,PinnParser.Tokens.T__18.rawValue,PinnParser.Tokens.T__20.rawValue,PinnParser.Tokens.BINOP.rawValue]
						    return  Utils.testBitLeftShiftArray(testArray, 0)
						}()
						      return testSet
						 }())) {
						try _errHandler.recoverInline(self)
						}
						else {
							_errHandler.reportMatch(self)
							try consume()
						}
						setState(201)
						try expr(14)

						break
					case 2:
						_localctx = ExprContext(_parentctx, _parentState);
						try pushNewRecursionContext(_localctx, _startState, PinnParser.RULE_expr)
						setState(202)
						if (!(precpred(_ctx, 12))) {
						    throw ANTLRException.recognition(e:FailedPredicateException(self, "precpred(_ctx, 12)"))
						}
						setState(203)
						_la = try _input.LA(1)
						if (!(//closure
						 { () -> Bool in
						      let testSet: Bool = {  () -> Bool in
						   let testArray: [Int] = [_la, PinnParser.Tokens.T__21.rawValue,PinnParser.Tokens.T__22.rawValue,PinnParser.Tokens.T__23.rawValue,PinnParser.Tokens.T__24.rawValue,PinnParser.Tokens.T__25.rawValue,PinnParser.Tokens.T__26.rawValue]
						    return  Utils.testBitLeftShiftArray(testArray, 0)
						}()
						      return testSet
						 }())) {
						try _errHandler.recoverInline(self)
						}
						else {
							_errHandler.reportMatch(self)
							try consume()
						}
						setState(204)
						try expr(13)

						break
					case 3:
						_localctx = ExprContext(_parentctx, _parentState);
						try pushNewRecursionContext(_localctx, _startState, PinnParser.RULE_expr)
						setState(205)
						if (!(precpred(_ctx, 11))) {
						    throw ANTLRException.recognition(e:FailedPredicateException(self, "precpred(_ctx, 11)"))
						}
						setState(206)
						_la = try _input.LA(1)
						if (!(//closure
						 { () -> Bool in
						      let testSet: Bool = _la == PinnParser.Tokens.T__27.rawValue || _la == PinnParser.Tokens.T__28.rawValue
						      return testSet
						 }())) {
						try _errHandler.recoverInline(self)
						}
						else {
							_errHandler.reportMatch(self)
							try consume()
						}
						setState(207)
						try expr(12)

						break
					case 4:
						_localctx = ExprContext(_parentctx, _parentState);
						try pushNewRecursionContext(_localctx, _startState, PinnParser.RULE_expr)
						setState(208)
						if (!(precpred(_ctx, 10))) {
						    throw ANTLRException.recognition(e:FailedPredicateException(self, "precpred(_ctx, 10)"))
						}
						setState(209)
						try match(PinnParser.Tokens.T__29.rawValue)
						setState(210)
						try expr(0)
						setState(211)
						try match(PinnParser.Tokens.COLON.rawValue)
						setState(212)
						try expr(11)

						break
					case 5:
						_localctx = ExprContext(_parentctx, _parentState);
						_localctx.firstExpr = _prevctx
						_localctx.firstExpr = _prevctx;
						try pushNewRecursionContext(_localctx, _startState, PinnParser.RULE_expr)
						setState(214)
						if (!(precpred(_ctx, 8))) {
						    throw ANTLRException.recognition(e:FailedPredicateException(self, "precpred(_ctx, 8)"))
						}
						setState(215)
						_la = try _input.LA(1)
						if (!(//closure
						 { () -> Bool in
						      let testSet: Bool = _la == PinnParser.Tokens.TWODOTS.rawValue || _la == PinnParser.Tokens.COLON.rawValue
						      return testSet
						 }())) {
						try _errHandler.recoverInline(self)
						}
						else {
							_errHandler.reportMatch(self)
							try consume()
						}
						setState(216)
						try {
								let assignmentValue = try expr(9)
								_localctx.castdown(ExprContext.self).secondExpr = assignmentValue
						     }()


						break
					default: break
					}
			 
				}
				setState(221)
				try _errHandler.sync(self)
				_alt = try getInterpreter().adaptivePredict(_input,19,_ctx)
			}

		}
		catch ANTLRException.recognition(let re) {
			_localctx.exception = re
			_errHandler.reportError(self, re)
			try _errHandler.recover(self, re)
		}

		return _localctx;
	}

	public class ExprListContext: ParserRuleContext {
			open
			func expr() -> [ExprContext] {
				return getRuleContexts(ExprContext.self)
			}
			open
			func expr(_ i: Int) -> ExprContext? {
				return getRuleContext(ExprContext.self, i)
			}
		override open
		func getRuleIndex() -> Int {
			return PinnParser.RULE_exprList
		}
	}
	@discardableResult
	 open func exprList() throws -> ExprListContext {
		var _localctx: ExprListContext = ExprListContext(_ctx, getState())
		try enterRule(_localctx, 22, PinnParser.RULE_exprList)
		var _la: Int = 0
		defer {
	    		try! exitRule()
	    }
		do {
		 	try enterOuterAlt(_localctx, 1)
		 	setState(222)
		 	try expr(0)
		 	setState(227)
		 	try _errHandler.sync(self)
		 	_la = try _input.LA(1)
		 	while (//closure
		 	 { () -> Bool in
		 	      let testSet: Bool = _la == PinnParser.Tokens.T__2.rawValue
		 	      return testSet
		 	 }()) {
		 		setState(223)
		 		try match(PinnParser.Tokens.T__2.rawValue)
		 		setState(224)
		 		try expr(0)


		 		setState(229)
		 		try _errHandler.sync(self)
		 		_la = try _input.LA(1)
		 	}

		}
		catch ANTLRException.recognition(let re) {
			_localctx.exception = re
			_errHandler.reportError(self, re)
			try _errHandler.recover(self, re)
		}

		return _localctx
	}

	public class ReturnStatementContext: ParserRuleContext {
			open
			func expr() -> ExprContext? {
				return getRuleContext(ExprContext.self, 0)
			}
		override open
		func getRuleIndex() -> Int {
			return PinnParser.RULE_returnStatement
		}
	}
	@discardableResult
	 open func returnStatement() throws -> ReturnStatementContext {
		var _localctx: ReturnStatementContext = ReturnStatementContext(_ctx, getState())
		try enterRule(_localctx, 24, PinnParser.RULE_returnStatement)
		var _la: Int = 0
		defer {
	    		try! exitRule()
	    }
		do {
		 	try enterOuterAlt(_localctx, 1)
		 	setState(230)
		 	try match(PinnParser.Tokens.T__30.rawValue)
		 	setState(232)
		 	try _errHandler.sync(self)
		 	_la = try _input.LA(1)
		 	if (//closure
		 	 { () -> Bool in
		 	      let testSet: Bool = {  () -> Bool in
		 	   let testArray: [Int] = [_la, PinnParser.Tokens.T__10.rawValue,PinnParser.Tokens.T__11.rawValue,PinnParser.Tokens.T__12.rawValue,PinnParser.Tokens.T__13.rawValue,PinnParser.Tokens.T__14.rawValue,PinnParser.Tokens.T__15.rawValue,PinnParser.Tokens.T__16.rawValue,PinnParser.Tokens.T__17.rawValue,PinnParser.Tokens.T__18.rawValue,PinnParser.Tokens.T__19.rawValue,PinnParser.Tokens.T__20.rawValue,PinnParser.Tokens.LPAREN.rawValue,PinnParser.Tokens.IOTA.rawValue,PinnParser.Tokens.BOOL.rawValue,PinnParser.Tokens.ID.rawValue,PinnParser.Tokens.CHAR.rawValue,PinnParser.Tokens.INT.rawValue,PinnParser.Tokens.FLOAT.rawValue,PinnParser.Tokens.STRING.rawValue]
		 	    return  Utils.testBitLeftShiftArray(testArray, 0)
		 	}()
		 	      return testSet
		 	 }()) {
		 		setState(231)
		 		try expr(0)

		 	}

		 	setState(234)
		 	try match(PinnParser.Tokens.T__0.rawValue)

		}
		catch ANTLRException.recognition(let re) {
			_localctx.exception = re
			_errHandler.reportError(self, re)
			try _errHandler.recover(self, re)
		}

		return _localctx
	}

	public class IfStatementContext: ParserRuleContext {
			open
			func expr() -> ExprContext? {
				return getRuleContext(ExprContext.self, 0)
			}
			open
			func statement() -> [StatementContext] {
				return getRuleContexts(StatementContext.self)
			}
			open
			func statement(_ i: Int) -> StatementContext? {
				return getRuleContext(StatementContext.self, i)
			}
		override open
		func getRuleIndex() -> Int {
			return PinnParser.RULE_ifStatement
		}
	}
	@discardableResult
	 open func ifStatement() throws -> IfStatementContext {
		var _localctx: IfStatementContext = IfStatementContext(_ctx, getState())
		try enterRule(_localctx, 26, PinnParser.RULE_ifStatement)
		defer {
	    		try! exitRule()
	    }
		do {
		 	try enterOuterAlt(_localctx, 1)
		 	setState(236)
		 	try match(PinnParser.Tokens.T__31.rawValue)
		 	setState(237)
		 	try expr(0)
		 	setState(238)
		 	try statement()
		 	setState(241)
		 	try _errHandler.sync(self)
		 	switch (try getInterpreter().adaptivePredict(_input,22,_ctx)) {
		 	case 1:
		 		setState(239)
		 		try match(PinnParser.Tokens.T__32.rawValue)
		 		setState(240)
		 		try statement()

		 		break
		 	default: break
		 	}

		}
		catch ANTLRException.recognition(let re) {
			_localctx.exception = re
			_errHandler.reportError(self, re)
			try _errHandler.recover(self, re)
		}

		return _localctx
	}

	public class GuardStatementContext: ParserRuleContext {
			open
			func expr() -> ExprContext? {
				return getRuleContext(ExprContext.self, 0)
			}
			open
			func block() -> BlockContext? {
				return getRuleContext(BlockContext.self, 0)
			}
		override open
		func getRuleIndex() -> Int {
			return PinnParser.RULE_guardStatement
		}
	}
	@discardableResult
	 open func guardStatement() throws -> GuardStatementContext {
		var _localctx: GuardStatementContext = GuardStatementContext(_ctx, getState())
		try enterRule(_localctx, 28, PinnParser.RULE_guardStatement)
		defer {
	    		try! exitRule()
	    }
		do {
		 	try enterOuterAlt(_localctx, 1)
		 	setState(243)
		 	try match(PinnParser.Tokens.T__33.rawValue)
		 	setState(244)
		 	try expr(0)
		 	setState(245)
		 	try match(PinnParser.Tokens.T__32.rawValue)
		 	setState(246)
		 	try block()

		}
		catch ANTLRException.recognition(let re) {
			_localctx.exception = re
			_errHandler.reportError(self, re)
			try _errHandler.recover(self, re)
		}

		return _localctx
	}

	public class WhStatementContext: ParserRuleContext {
			open
			func expr() -> ExprContext? {
				return getRuleContext(ExprContext.self, 0)
			}
			open
			func block() -> BlockContext? {
				return getRuleContext(BlockContext.self, 0)
			}
		override open
		func getRuleIndex() -> Int {
			return PinnParser.RULE_whStatement
		}
	}
	@discardableResult
	 open func whStatement() throws -> WhStatementContext {
		var _localctx: WhStatementContext = WhStatementContext(_ctx, getState())
		try enterRule(_localctx, 30, PinnParser.RULE_whStatement)
		defer {
	    		try! exitRule()
	    }
		do {
		 	try enterOuterAlt(_localctx, 1)
		 	setState(248)
		 	try match(PinnParser.Tokens.T__34.rawValue)
		 	setState(249)
		 	try expr(0)
		 	setState(250)
		 	try block()

		}
		catch ANTLRException.recognition(let re) {
			_localctx.exception = re
			_errHandler.reportError(self, re)
			try _errHandler.recover(self, re)
		}

		return _localctx
	}

	public class RepeatStatementContext: ParserRuleContext {
			open
			func block() -> BlockContext? {
				return getRuleContext(BlockContext.self, 0)
			}
			open
			func expr() -> ExprContext? {
				return getRuleContext(ExprContext.self, 0)
			}
		override open
		func getRuleIndex() -> Int {
			return PinnParser.RULE_repeatStatement
		}
	}
	@discardableResult
	 open func repeatStatement() throws -> RepeatStatementContext {
		var _localctx: RepeatStatementContext = RepeatStatementContext(_ctx, getState())
		try enterRule(_localctx, 32, PinnParser.RULE_repeatStatement)
		defer {
	    		try! exitRule()
	    }
		do {
		 	try enterOuterAlt(_localctx, 1)
		 	setState(252)
		 	try match(PinnParser.Tokens.T__35.rawValue)
		 	setState(253)
		 	try block()
		 	setState(254)
		 	try match(PinnParser.Tokens.T__34.rawValue)
		 	setState(255)
		 	try expr(0)

		}
		catch ANTLRException.recognition(let re) {
			_localctx.exception = re
			_errHandler.reportError(self, re)
			try _errHandler.recover(self, re)
		}

		return _localctx
	}

	public class FoStatementContext: ParserRuleContext {
		open var fss: SimpleStatementContext!
		open var sss: SimpleStatementContext!
			open
			func expr() -> ExprContext? {
				return getRuleContext(ExprContext.self, 0)
			}
			open
			func block() -> BlockContext? {
				return getRuleContext(BlockContext.self, 0)
			}
			open
			func simpleStatement() -> [SimpleStatementContext] {
				return getRuleContexts(SimpleStatementContext.self)
			}
			open
			func simpleStatement(_ i: Int) -> SimpleStatementContext? {
				return getRuleContext(SimpleStatementContext.self, i)
			}
			open
			func varDecl() -> VarDeclContext? {
				return getRuleContext(VarDeclContext.self, 0)
			}
			open
			func ID() -> [TerminalNode] {
				return getTokens(PinnParser.Tokens.ID.rawValue)
			}
			open
			func ID(_ i:Int) -> TerminalNode? {
				return getToken(PinnParser.Tokens.ID.rawValue, i)
			}
			open
			func RANGE() -> TerminalNode? {
				return getToken(PinnParser.Tokens.RANGE.rawValue, 0)
			}
		override open
		func getRuleIndex() -> Int {
			return PinnParser.RULE_foStatement
		}
	}
	@discardableResult
	 open func foStatement() throws -> FoStatementContext {
		var _localctx: FoStatementContext = FoStatementContext(_ctx, getState())
		try enterRule(_localctx, 34, PinnParser.RULE_foStatement)
		defer {
	    		try! exitRule()
	    }
		do {
		 	setState(284)
		 	try _errHandler.sync(self)
		 	switch(try getInterpreter().adaptivePredict(_input,24, _ctx)) {
		 	case 1:
		 		try enterOuterAlt(_localctx, 1)
		 		setState(257)
		 		try match(PinnParser.Tokens.T__36.rawValue)
		 		setState(260)
		 		try _errHandler.sync(self)
		 		switch (try getInterpreter().adaptivePredict(_input,23,_ctx)) {
		 		case 1:
		 			setState(258)
		 			try varDecl()

		 			break
		 		case 2:
		 			setState(259)
		 			try {
		 					let assignmentValue = try simpleStatement()
		 					_localctx.castdown(FoStatementContext.self).fss = assignmentValue
		 			     }()


		 			break
		 		default: break
		 		}
		 		setState(262)
		 		try match(PinnParser.Tokens.T__0.rawValue)
		 		setState(263)
		 		try expr(0)
		 		setState(264)
		 		try match(PinnParser.Tokens.T__0.rawValue)
		 		setState(265)
		 		try {
		 				let assignmentValue = try simpleStatement()
		 				_localctx.castdown(FoStatementContext.self).sss = assignmentValue
		 		     }()

		 		setState(266)
		 		try block()

		 		break
		 	case 2:
		 		try enterOuterAlt(_localctx, 2)
		 		setState(268)
		 		try match(PinnParser.Tokens.T__36.rawValue)
		 		setState(269)
		 		try match(PinnParser.Tokens.ID.rawValue)
		 		setState(270)
		 		try match(PinnParser.Tokens.T__2.rawValue)
		 		setState(271)
		 		try match(PinnParser.Tokens.ID.rawValue)
		 		setState(272)
		 		try match(PinnParser.Tokens.T__6.rawValue)
		 		setState(273)
		 		try match(PinnParser.Tokens.RANGE.rawValue)
		 		setState(274)
		 		try expr(0)
		 		setState(275)
		 		try block()

		 		break
		 	case 3:
		 		try enterOuterAlt(_localctx, 3)
		 		setState(277)
		 		try match(PinnParser.Tokens.T__36.rawValue)
		 		setState(278)
		 		try match(PinnParser.Tokens.ID.rawValue)
		 		setState(279)
		 		try match(PinnParser.Tokens.T__6.rawValue)
		 		setState(280)
		 		try match(PinnParser.Tokens.RANGE.rawValue)
		 		setState(281)
		 		try expr(0)
		 		setState(282)
		 		try block()

		 		break
		 	default: break
		 	}
		}
		catch ANTLRException.recognition(let re) {
			_localctx.exception = re
			_errHandler.reportError(self, re)
			try _errHandler.recover(self, re)
		}

		return _localctx
	}

	public class CaseStatementContext: ParserRuleContext {
			open
			func exprList() -> ExprListContext? {
				return getRuleContext(ExprListContext.self, 0)
			}
			open
			func COLON() -> TerminalNode? {
				return getToken(PinnParser.Tokens.COLON.rawValue, 0)
			}
			open
			func statement() -> [StatementContext] {
				return getRuleContexts(StatementContext.self)
			}
			open
			func statement(_ i: Int) -> StatementContext? {
				return getRuleContext(StatementContext.self, i)
			}
		override open
		func getRuleIndex() -> Int {
			return PinnParser.RULE_caseStatement
		}
	}
	@discardableResult
	 open func caseStatement() throws -> CaseStatementContext {
		var _localctx: CaseStatementContext = CaseStatementContext(_ctx, getState())
		try enterRule(_localctx, 36, PinnParser.RULE_caseStatement)
		var _la: Int = 0
		defer {
	    		try! exitRule()
	    }
		do {
		 	try enterOuterAlt(_localctx, 1)
		 	setState(286)
		 	try match(PinnParser.Tokens.T__37.rawValue)
		 	setState(287)
		 	try exprList()
		 	setState(288)
		 	try match(PinnParser.Tokens.COLON.rawValue)
		 	setState(292)
		 	try _errHandler.sync(self)
		 	_la = try _input.LA(1)
		 	while (//closure
		 	 { () -> Bool in
		 	      let testSet: Bool = {  () -> Bool in
		 	   let testArray: [Int] = [_la, PinnParser.Tokens.T__0.rawValue,PinnParser.Tokens.T__4.rawValue,PinnParser.Tokens.T__10.rawValue,PinnParser.Tokens.T__11.rawValue,PinnParser.Tokens.T__12.rawValue,PinnParser.Tokens.T__13.rawValue,PinnParser.Tokens.T__14.rawValue,PinnParser.Tokens.T__15.rawValue,PinnParser.Tokens.T__16.rawValue,PinnParser.Tokens.T__17.rawValue,PinnParser.Tokens.T__18.rawValue,PinnParser.Tokens.T__19.rawValue,PinnParser.Tokens.T__20.rawValue,PinnParser.Tokens.T__30.rawValue,PinnParser.Tokens.T__31.rawValue,PinnParser.Tokens.T__33.rawValue,PinnParser.Tokens.T__34.rawValue,PinnParser.Tokens.T__35.rawValue,PinnParser.Tokens.T__36.rawValue,PinnParser.Tokens.T__38.rawValue,PinnParser.Tokens.T__40.rawValue,PinnParser.Tokens.T__41.rawValue,PinnParser.Tokens.T__42.rawValue,PinnParser.Tokens.T__43.rawValue,PinnParser.Tokens.LPAREN.rawValue,PinnParser.Tokens.IOTA.rawValue,PinnParser.Tokens.BOOL.rawValue,PinnParser.Tokens.ID.rawValue,PinnParser.Tokens.CHAR.rawValue,PinnParser.Tokens.INT.rawValue,PinnParser.Tokens.FLOAT.rawValue,PinnParser.Tokens.STRING.rawValue]
		 	    return  Utils.testBitLeftShiftArray(testArray, 0)
		 	}()
		 	      return testSet
		 	 }()) {
		 		setState(289)
		 		try statement()


		 		setState(294)
		 		try _errHandler.sync(self)
		 		_la = try _input.LA(1)
		 	}

		}
		catch ANTLRException.recognition(let re) {
			_localctx.exception = re
			_errHandler.reportError(self, re)
			try _errHandler.recover(self, re)
		}

		return _localctx
	}

	public class SwitchStatementContext: ParserRuleContext {
			open
			func expr() -> ExprContext? {
				return getRuleContext(ExprContext.self, 0)
			}
			open
			func caseStatement() -> [CaseStatementContext] {
				return getRuleContexts(CaseStatementContext.self)
			}
			open
			func caseStatement(_ i: Int) -> CaseStatementContext? {
				return getRuleContext(CaseStatementContext.self, i)
			}
			open
			func COLON() -> TerminalNode? {
				return getToken(PinnParser.Tokens.COLON.rawValue, 0)
			}
			open
			func statement() -> [StatementContext] {
				return getRuleContexts(StatementContext.self)
			}
			open
			func statement(_ i: Int) -> StatementContext? {
				return getRuleContext(StatementContext.self, i)
			}
		override open
		func getRuleIndex() -> Int {
			return PinnParser.RULE_switchStatement
		}
	}
	@discardableResult
	 open func switchStatement() throws -> SwitchStatementContext {
		var _localctx: SwitchStatementContext = SwitchStatementContext(_ctx, getState())
		try enterRule(_localctx, 38, PinnParser.RULE_switchStatement)
		var _la: Int = 0
		defer {
	    		try! exitRule()
	    }
		do {
		 	try enterOuterAlt(_localctx, 1)
		 	setState(295)
		 	try match(PinnParser.Tokens.T__38.rawValue)
		 	setState(296)
		 	try expr(0)
		 	setState(297)
		 	try match(PinnParser.Tokens.T__4.rawValue)
		 	setState(301)
		 	try _errHandler.sync(self)
		 	_la = try _input.LA(1)
		 	while (//closure
		 	 { () -> Bool in
		 	      let testSet: Bool = _la == PinnParser.Tokens.T__37.rawValue
		 	      return testSet
		 	 }()) {
		 		setState(298)
		 		try caseStatement()


		 		setState(303)
		 		try _errHandler.sync(self)
		 		_la = try _input.LA(1)
		 	}
		 	setState(312)
		 	try _errHandler.sync(self)
		 	_la = try _input.LA(1)
		 	if (//closure
		 	 { () -> Bool in
		 	      let testSet: Bool = _la == PinnParser.Tokens.T__39.rawValue
		 	      return testSet
		 	 }()) {
		 		setState(304)
		 		try match(PinnParser.Tokens.T__39.rawValue)
		 		setState(305)
		 		try match(PinnParser.Tokens.COLON.rawValue)
		 		setState(309)
		 		try _errHandler.sync(self)
		 		_la = try _input.LA(1)
		 		while (//closure
		 		 { () -> Bool in
		 		      let testSet: Bool = {  () -> Bool in
		 		   let testArray: [Int] = [_la, PinnParser.Tokens.T__0.rawValue,PinnParser.Tokens.T__4.rawValue,PinnParser.Tokens.T__10.rawValue,PinnParser.Tokens.T__11.rawValue,PinnParser.Tokens.T__12.rawValue,PinnParser.Tokens.T__13.rawValue,PinnParser.Tokens.T__14.rawValue,PinnParser.Tokens.T__15.rawValue,PinnParser.Tokens.T__16.rawValue,PinnParser.Tokens.T__17.rawValue,PinnParser.Tokens.T__18.rawValue,PinnParser.Tokens.T__19.rawValue,PinnParser.Tokens.T__20.rawValue,PinnParser.Tokens.T__30.rawValue,PinnParser.Tokens.T__31.rawValue,PinnParser.Tokens.T__33.rawValue,PinnParser.Tokens.T__34.rawValue,PinnParser.Tokens.T__35.rawValue,PinnParser.Tokens.T__36.rawValue,PinnParser.Tokens.T__38.rawValue,PinnParser.Tokens.T__40.rawValue,PinnParser.Tokens.T__41.rawValue,PinnParser.Tokens.T__42.rawValue,PinnParser.Tokens.T__43.rawValue,PinnParser.Tokens.LPAREN.rawValue,PinnParser.Tokens.IOTA.rawValue,PinnParser.Tokens.BOOL.rawValue,PinnParser.Tokens.ID.rawValue,PinnParser.Tokens.CHAR.rawValue,PinnParser.Tokens.INT.rawValue,PinnParser.Tokens.FLOAT.rawValue,PinnParser.Tokens.STRING.rawValue]
		 		    return  Utils.testBitLeftShiftArray(testArray, 0)
		 		}()
		 		      return testSet
		 		 }()) {
		 			setState(306)
		 			try statement()


		 			setState(311)
		 			try _errHandler.sync(self)
		 			_la = try _input.LA(1)
		 		}

		 	}

		 	setState(314)
		 	try match(PinnParser.Tokens.T__5.rawValue)

		}
		catch ANTLRException.recognition(let re) {
			_localctx.exception = re
			_errHandler.reportError(self, re)
			try _errHandler.recover(self, re)
		}

		return _localctx
	}

	public class StatementContext: ParserRuleContext {
			open
			func expr() -> ExprContext? {
				return getRuleContext(ExprContext.self, 0)
			}
			open
			func varDecl() -> VarDeclContext? {
				return getRuleContext(VarDeclContext.self, 0)
			}
			open
			func simpleStatement() -> SimpleStatementContext? {
				return getRuleContext(SimpleStatementContext.self, 0)
			}
			open
			func ifStatement() -> IfStatementContext? {
				return getRuleContext(IfStatementContext.self, 0)
			}
			open
			func guardStatement() -> GuardStatementContext? {
				return getRuleContext(GuardStatementContext.self, 0)
			}
			open
			func whStatement() -> WhStatementContext? {
				return getRuleContext(WhStatementContext.self, 0)
			}
			open
			func repeatStatement() -> RepeatStatementContext? {
				return getRuleContext(RepeatStatementContext.self, 0)
			}
			open
			func switchStatement() -> SwitchStatementContext? {
				return getRuleContext(SwitchStatementContext.self, 0)
			}
			open
			func returnStatement() -> ReturnStatementContext? {
				return getRuleContext(ReturnStatementContext.self, 0)
			}
			open
			func foStatement() -> FoStatementContext? {
				return getRuleContext(FoStatementContext.self, 0)
			}
			open
			func block() -> BlockContext? {
				return getRuleContext(BlockContext.self, 0)
			}
			open
			func LPAREN() -> TerminalNode? {
				return getToken(PinnParser.Tokens.LPAREN.rawValue, 0)
			}
		override open
		func getRuleIndex() -> Int {
			return PinnParser.RULE_statement
		}
	}
	@discardableResult
	 open func statement() throws -> StatementContext {
		var _localctx: StatementContext = StatementContext(_ctx, getState())
		try enterRule(_localctx, 40, PinnParser.RULE_statement)
		defer {
	    		try! exitRule()
	    }
		do {
		 	setState(346)
		 	try _errHandler.sync(self)
		 	switch(try getInterpreter().adaptivePredict(_input,29, _ctx)) {
		 	case 1:
		 		try enterOuterAlt(_localctx, 1)
		 		setState(316)
		 		try expr(0)
		 		setState(317)
		 		try match(PinnParser.Tokens.T__0.rawValue)

		 		break
		 	case 2:
		 		try enterOuterAlt(_localctx, 2)
		 		setState(319)
		 		try varDecl()
		 		setState(320)
		 		try match(PinnParser.Tokens.T__0.rawValue)

		 		break
		 	case 3:
		 		try enterOuterAlt(_localctx, 3)
		 		setState(322)
		 		try simpleStatement()
		 		setState(323)
		 		try match(PinnParser.Tokens.T__0.rawValue)

		 		break
		 	case 4:
		 		try enterOuterAlt(_localctx, 4)
		 		setState(325)
		 		try ifStatement()

		 		break
		 	case 5:
		 		try enterOuterAlt(_localctx, 5)
		 		setState(326)
		 		try guardStatement()

		 		break
		 	case 6:
		 		try enterOuterAlt(_localctx, 6)
		 		setState(327)
		 		try whStatement()

		 		break
		 	case 7:
		 		try enterOuterAlt(_localctx, 7)
		 		setState(328)
		 		try repeatStatement()
		 		setState(329)
		 		try match(PinnParser.Tokens.T__0.rawValue)

		 		break
		 	case 8:
		 		try enterOuterAlt(_localctx, 8)
		 		setState(331)
		 		try switchStatement()

		 		break
		 	case 9:
		 		try enterOuterAlt(_localctx, 9)
		 		setState(332)
		 		try returnStatement()

		 		break
		 	case 10:
		 		try enterOuterAlt(_localctx, 10)
		 		setState(333)
		 		try foStatement()

		 		break
		 	case 11:
		 		try enterOuterAlt(_localctx, 11)
		 		setState(334)
		 		try block()

		 		break
		 	case 12:
		 		try enterOuterAlt(_localctx, 12)
		 		setState(335)
		 		try match(PinnParser.Tokens.T__40.rawValue)
		 		setState(336)
		 		try match(PinnParser.Tokens.T__0.rawValue)

		 		break
		 	case 13:
		 		try enterOuterAlt(_localctx, 13)
		 		setState(337)
		 		try match(PinnParser.Tokens.T__41.rawValue)
		 		setState(338)
		 		try match(PinnParser.Tokens.T__0.rawValue)

		 		break
		 	case 14:
		 		try enterOuterAlt(_localctx, 14)
		 		setState(339)
		 		try match(PinnParser.Tokens.T__42.rawValue)
		 		setState(340)
		 		try match(PinnParser.Tokens.T__0.rawValue)

		 		break
		 	case 15:
		 		try enterOuterAlt(_localctx, 15)
		 		setState(341)
		 		try match(PinnParser.Tokens.T__43.rawValue)
		 		setState(342)
		 		try match(PinnParser.Tokens.LPAREN.rawValue)
		 		setState(343)
		 		try match(PinnParser.Tokens.T__3.rawValue)
		 		setState(344)
		 		try match(PinnParser.Tokens.T__0.rawValue)

		 		break
		 	case 16:
		 		try enterOuterAlt(_localctx, 16)
		 		setState(345)
		 		try match(PinnParser.Tokens.T__0.rawValue)

		 		break
		 	default: break
		 	}
		}
		catch ANTLRException.recognition(let re) {
			_localctx.exception = re
			_errHandler.reportError(self, re)
			try _errHandler.recover(self, re)
		}

		return _localctx
	}

	public class PsetContext: ParserRuleContext {
		override open
		func getRuleIndex() -> Int {
			return PinnParser.RULE_pset
		}
	 
		open
		func copyFrom(_ ctx: PsetContext) {
			super.copyFrom(ctx)
		}
	}
	public class SimpleSetContext: PsetContext {
			open
			func ID() -> TerminalNode? {
				return getToken(PinnParser.Tokens.ID.rawValue, 0)
			}
			open
			func expr() -> [ExprContext] {
				return getRuleContexts(ExprContext.self)
			}
			open
			func expr(_ i: Int) -> ExprContext? {
				return getRuleContext(ExprContext.self, i)
			}
			open
			func LSQUARE() -> TerminalNode? {
				return getToken(PinnParser.Tokens.LSQUARE.rawValue, 0)
			}

		public
		init(_ ctx: PsetContext) {
			super.init()
			copyFrom(ctx)
		}
	}
	public class CompoundSetContext: PsetContext {
			open
			func ID() -> TerminalNode? {
				return getToken(PinnParser.Tokens.ID.rawValue, 0)
			}
			open
			func expr() -> [ExprContext] {
				return getRuleContexts(ExprContext.self)
			}
			open
			func expr(_ i: Int) -> ExprContext? {
				return getRuleContext(ExprContext.self, i)
			}
			open
			func BINOP() -> TerminalNode? {
				return getToken(PinnParser.Tokens.BINOP.rawValue, 0)
			}
			open
			func LSQUARE() -> TerminalNode? {
				return getToken(PinnParser.Tokens.LSQUARE.rawValue, 0)
			}

		public
		init(_ ctx: PsetContext) {
			super.init()
			copyFrom(ctx)
		}
	}
	@discardableResult
	 open func pset() throws -> PsetContext {
		var _localctx: PsetContext = PsetContext(_ctx, getState())
		try enterRule(_localctx, 42, PinnParser.RULE_pset)
		var _la: Int = 0
		defer {
	    		try! exitRule()
	    }
		do {
		 	setState(367)
		 	try _errHandler.sync(self)
		 	switch(try getInterpreter().adaptivePredict(_input,32, _ctx)) {
		 	case 1:
		 		_localctx =  SimpleSetContext(_localctx);
		 		try enterOuterAlt(_localctx, 1)
		 		setState(348)
		 		try match(PinnParser.Tokens.ID.rawValue)
		 		setState(353)
		 		try _errHandler.sync(self)
		 		_la = try _input.LA(1)
		 		if (//closure
		 		 { () -> Bool in
		 		      let testSet: Bool = _la == PinnParser.Tokens.LSQUARE.rawValue
		 		      return testSet
		 		 }()) {
		 			setState(349)
		 			try match(PinnParser.Tokens.LSQUARE.rawValue)
		 			setState(350)
		 			try expr(0)
		 			setState(351)
		 			try match(PinnParser.Tokens.T__7.rawValue)

		 		}

		 		setState(355)
		 		try match(PinnParser.Tokens.T__6.rawValue)
		 		setState(356)
		 		try expr(0)

		 		break
		 	case 2:
		 		_localctx =  CompoundSetContext(_localctx);
		 		try enterOuterAlt(_localctx, 2)
		 		setState(357)
		 		try match(PinnParser.Tokens.ID.rawValue)
		 		setState(362)
		 		try _errHandler.sync(self)
		 		_la = try _input.LA(1)
		 		if (//closure
		 		 { () -> Bool in
		 		      let testSet: Bool = _la == PinnParser.Tokens.LSQUARE.rawValue
		 		      return testSet
		 		 }()) {
		 			setState(358)
		 			try match(PinnParser.Tokens.LSQUARE.rawValue)
		 			setState(359)
		 			try expr(0)
		 			setState(360)
		 			try match(PinnParser.Tokens.T__7.rawValue)

		 		}

		 		setState(364)
		 		_la = try _input.LA(1)
		 		if (!(//closure
		 		 { () -> Bool in
		 		      let testSet: Bool = {  () -> Bool in
		 		   let testArray: [Int] = [_la, PinnParser.Tokens.T__17.rawValue,PinnParser.Tokens.T__18.rawValue,PinnParser.Tokens.T__20.rawValue,PinnParser.Tokens.BINOP.rawValue]
		 		    return  Utils.testBitLeftShiftArray(testArray, 0)
		 		}()
		 		      return testSet
		 		 }())) {
		 		try _errHandler.recoverInline(self)
		 		}
		 		else {
		 			_errHandler.reportMatch(self)
		 			try consume()
		 		}
		 		setState(365)
		 		try match(PinnParser.Tokens.T__6.rawValue)
		 		setState(366)
		 		try expr(0)

		 		break
		 	default: break
		 	}
		}
		catch ANTLRException.recognition(let re) {
			_localctx.exception = re
			_errHandler.reportError(self, re)
			try _errHandler.recover(self, re)
		}

		return _localctx
	}

	override open
	func sempred(_ _localctx: RuleContext?, _ ruleIndex: Int,  _ predIndex: Int)throws -> Bool {
		switch (ruleIndex) {
		case  10:
			return try expr_sempred(_localctx?.castdown(ExprContext.self), predIndex)
	    default: return true
		}
	}
	private func expr_sempred(_ _localctx: ExprContext!,  _ predIndex: Int) throws -> Bool {
		switch (predIndex) {
		    case 0:return precpred(_ctx, 13)
		    case 1:return precpred(_ctx, 12)
		    case 2:return precpred(_ctx, 11)
		    case 3:return precpred(_ctx, 10)
		    case 4:return precpred(_ctx, 8)
		    default: return true
		}
	}


	public
	static let _serializedATN = PinnParserATN().jsonString

	public
	static let _ATN = ATNDeserializer().deserializeFromJson(_serializedATN)
}