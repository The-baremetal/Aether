// Generated from d:/Programs/Programmed/fluxlang/antlr4/FLUXASSEMBLY/src/lua_grammar_antlr4.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast", "CheckReturnValue"})
public class lua_grammar_antlr4Parser extends Parser {
	static { RuntimeMetaData.checkVersion("4.13.1", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		T__0=1, T__1=2, T__2=3, T__3=4, T__4=5, T__5=6, T__6=7, T__7=8, T__8=9, 
		T__9=10, T__10=11, T__11=12, T__12=13, T__13=14, T__14=15, T__15=16, T__16=17, 
		T__17=18, T__18=19, T__19=20, T__20=21, T__21=22, T__22=23, T__23=24, 
		T__24=25, T__25=26, T__26=27, T__27=28, T__28=29, T__29=30, T__30=31, 
		T__31=32, T__32=33, T__33=34, T__34=35, T__35=36, T__36=37, T__37=38, 
		T__38=39, T__39=40, T__40=41, T__41=42, T__42=43, T__43=44, T__44=45, 
		T__45=46, T__46=47, T__47=48, T__48=49, T__49=50, T__50=51, T__51=52, 
		T__52=53, T__53=54, T__54=55, T__55=56, T__56=57, T__57=58, T__58=59, 
		T__59=60, T__60=61, T__61=62, T__62=63, T__63=64, T__64=65, T__65=66, 
		T__66=67, T__67=68, T__68=69, T__69=70, T__70=71, T__71=72, T__72=73, 
		T__73=74, T__74=75, T__75=76, T__76=77, T__77=78, T__78=79, T__79=80, 
		T__80=81, T__81=82, T__82=83, T__83=84, IDENTIFIER=85, BOOL=86, NIL=87, 
		NUMBER=88, STRING=89, WS=90, COMMENT=91;
	public static final int
		RULE_program = 0, RULE_statement = 1, RULE_assignStatement = 2, RULE_expression = 3, 
		RULE_primaryExpression = 4, RULE_literal = 5, RULE_variable = 6, RULE_functionCall = 7, 
		RULE_tableConstructor = 8, RULE_metatable = 9, RULE_metamethods = 10, 
		RULE_labelStatement = 11, RULE_metamethod = 12, RULE_field = 13, RULE_binaryOperation = 14, 
		RULE_unaryOperation = 15, RULE_controlFlowStatement = 16, RULE_ifStatement = 17, 
		RULE_whileStatement = 18, RULE_repeatStatement = 19, RULE_forStatement = 20, 
		RULE_breakStatement = 21, RULE_gotoStatement = 22, RULE_coroutineStatement = 23, 
		RULE_block = 24, RULE_localDeclaration = 25, RULE_functionDeclaration = 26, 
		RULE_returnStatement = 27, RULE_operatorGroup = 28, RULE_logicalOp = 29, 
		RULE_comparisonOp = 30, RULE_arithOp = 31, RULE_bitwiseOp = 32, RULE_assignOp = 33, 
		RULE_unaryOp = 34, RULE_concatOp = 35, RULE_varargOp = 36, RULE_compoundAssignOp = 37, 
		RULE_incrOp = 38, RULE_coalesceOp = 39, RULE_shiftAssignOp = 40;
	private static String[] makeRuleNames() {
		return new String[] {
			"program", "statement", "assignStatement", "expression", "primaryExpression", 
			"literal", "variable", "functionCall", "tableConstructor", "metatable", 
			"metamethods", "labelStatement", "metamethod", "field", "binaryOperation", 
			"unaryOperation", "controlFlowStatement", "ifStatement", "whileStatement", 
			"repeatStatement", "forStatement", "breakStatement", "gotoStatement", 
			"coroutineStatement", "block", "localDeclaration", "functionDeclaration", 
			"returnStatement", "operatorGroup", "logicalOp", "comparisonOp", "arithOp", 
			"bitwiseOp", "assignOp", "unaryOp", "concatOp", "varargOp", "compoundAssignOp", 
			"incrOp", "coalesceOp", "shiftAssignOp"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'='", "'+='", "'-='", "'*='", "'/='", "'('", "')'", "','", "'{'", 
			"'}'", "'__metatable'", "'::'", "'__add'", "'__sub'", "'__mul'", "'__div'", 
			"'__mod'", "'__pow'", "'__unm'", "'__concat'", "'__len'", "'__eq'", "'__lt'", 
			"'__le'", "'__tostring'", "'__pairs'", "'__ipairs'", "'__call'", "'#'", 
			"'if'", "'then'", "'elseif'", "'else'", "'end'", "'while'", "'do'", "'repeat'", 
			"'until'", "'for'", "'break'", "'goto'", "'coroutine'", "'.'", "'create'", 
			"'resume'", "'yield'", "'status'", "'local'", "'function'", "'return'", 
			"'and'", "'or'", "'=='", "'>='", "'<='", "'~='", "'>'", "'<'", "'+'", 
			"'-'", "'*'", "'/'", "'//'", "'%'", "'^'", "'&'", "'|'", "'~'", "'<<'", 
			"'>>'", "'//='", "'^='", "'&='", "'|='", "'not'", "'..'", "'...'", "'..='", 
			"'??='", "'+_'", "'-_'", "'??'", "'<<='", "'>>='", null, null, "'nil'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, "IDENTIFIER", "BOOL", "NIL", "NUMBER", "STRING", "WS", "COMMENT"
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
	public String getGrammarFileName() { return "lua_grammar_antlr4.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public lua_grammar_antlr4Parser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ProgramContext extends ParserRuleContext {
		public TerminalNode EOF() { return getToken(lua_grammar_antlr4Parser.EOF, 0); }
		public List<StatementContext> statement() {
			return getRuleContexts(StatementContext.class);
		}
		public StatementContext statement(int i) {
			return getRuleContext(StatementContext.class,i);
		}
		public ProgramContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_program; }
	}

	public final ProgramContext program() throws RecognitionException {
		ProgramContext _localctx = new ProgramContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_program);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(85);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 1978744046620672L) != 0) || _la==IDENTIFIER) {
				{
				{
				setState(82);
				statement();
				}
				}
				setState(87);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(88);
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

	@SuppressWarnings("CheckReturnValue")
	public static class StatementContext extends ParserRuleContext {
		public AssignStatementContext assignStatement() {
			return getRuleContext(AssignStatementContext.class,0);
		}
		public ControlFlowStatementContext controlFlowStatement() {
			return getRuleContext(ControlFlowStatementContext.class,0);
		}
		public FunctionDeclarationContext functionDeclaration() {
			return getRuleContext(FunctionDeclarationContext.class,0);
		}
		public ReturnStatementContext returnStatement() {
			return getRuleContext(ReturnStatementContext.class,0);
		}
		public LocalDeclarationContext localDeclaration() {
			return getRuleContext(LocalDeclarationContext.class,0);
		}
		public LabelStatementContext labelStatement() {
			return getRuleContext(LabelStatementContext.class,0);
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
			setState(96);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case IDENTIFIER:
				enterOuterAlt(_localctx, 1);
				{
				setState(90);
				assignStatement();
				}
				break;
			case T__29:
			case T__34:
			case T__36:
			case T__38:
			case T__39:
			case T__40:
			case T__41:
				enterOuterAlt(_localctx, 2);
				{
				setState(91);
				controlFlowStatement();
				}
				break;
			case T__48:
				enterOuterAlt(_localctx, 3);
				{
				setState(92);
				functionDeclaration();
				}
				break;
			case T__49:
				enterOuterAlt(_localctx, 4);
				{
				setState(93);
				returnStatement();
				}
				break;
			case T__47:
				enterOuterAlt(_localctx, 5);
				{
				setState(94);
				localDeclaration();
				}
				break;
			case T__11:
				enterOuterAlt(_localctx, 6);
				{
				setState(95);
				labelStatement();
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

	@SuppressWarnings("CheckReturnValue")
	public static class AssignStatementContext extends ParserRuleContext {
		public VariableContext variable() {
			return getRuleContext(VariableContext.class,0);
		}
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public AssignStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_assignStatement; }
	}

	public final AssignStatementContext assignStatement() throws RecognitionException {
		AssignStatementContext _localctx = new AssignStatementContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_assignStatement);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(98);
			variable();
			setState(99);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 62L) != 0)) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			setState(100);
			expression(0);
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

	@SuppressWarnings("CheckReturnValue")
	public static class ExpressionContext extends ParserRuleContext {
		public PrimaryExpressionContext primaryExpression() {
			return getRuleContext(PrimaryExpressionContext.class,0);
		}
		public UnaryOpContext unaryOp() {
			return getRuleContext(UnaryOpContext.class,0);
		}
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public OperatorGroupContext operatorGroup() {
			return getRuleContext(OperatorGroupContext.class,0);
		}
		public ExpressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expression; }
	}

	public final ExpressionContext expression() throws RecognitionException {
		return expression(0);
	}

	private ExpressionContext expression(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		ExpressionContext _localctx = new ExpressionContext(_ctx, _parentState);
		ExpressionContext _prevctx = _localctx;
		int _startState = 6;
		enterRecursionRule(_localctx, 6, RULE_expression, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(107);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,2,_ctx) ) {
			case 1:
				{
				setState(103);
				primaryExpression();
				}
				break;
			case 2:
				{
				setState(104);
				unaryOp();
				setState(105);
				expression(1);
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(119);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,4,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(117);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,3,_ctx) ) {
					case 1:
						{
						_localctx = new ExpressionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(109);
						if (!(precpred(_ctx, 3))) throw new FailedPredicateException(this, "precpred(_ctx, 3)");
						setState(110);
						operatorGroup();
						setState(111);
						expression(4);
						}
						break;
					case 2:
						{
						_localctx = new ExpressionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(113);
						if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
						setState(114);
						operatorGroup();
						setState(115);
						expression(2);
						}
						break;
					}
					} 
				}
				setState(121);
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

	@SuppressWarnings("CheckReturnValue")
	public static class PrimaryExpressionContext extends ParserRuleContext {
		public LiteralContext literal() {
			return getRuleContext(LiteralContext.class,0);
		}
		public VariableContext variable() {
			return getRuleContext(VariableContext.class,0);
		}
		public FunctionCallContext functionCall() {
			return getRuleContext(FunctionCallContext.class,0);
		}
		public UnaryOperationContext unaryOperation() {
			return getRuleContext(UnaryOperationContext.class,0);
		}
		public TableConstructorContext tableConstructor() {
			return getRuleContext(TableConstructorContext.class,0);
		}
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public PrimaryExpressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_primaryExpression; }
	}

	public final PrimaryExpressionContext primaryExpression() throws RecognitionException {
		PrimaryExpressionContext _localctx = new PrimaryExpressionContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_primaryExpression);
		try {
			setState(131);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,5,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(122);
				literal();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(123);
				variable();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(124);
				functionCall();
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(125);
				unaryOperation();
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(126);
				tableConstructor();
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(127);
				match(T__5);
				setState(128);
				expression(0);
				setState(129);
				match(T__6);
				}
				break;
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

	@SuppressWarnings("CheckReturnValue")
	public static class LiteralContext extends ParserRuleContext {
		public TerminalNode BOOL() { return getToken(lua_grammar_antlr4Parser.BOOL, 0); }
		public TerminalNode NIL() { return getToken(lua_grammar_antlr4Parser.NIL, 0); }
		public TerminalNode NUMBER() { return getToken(lua_grammar_antlr4Parser.NUMBER, 0); }
		public TerminalNode STRING() { return getToken(lua_grammar_antlr4Parser.STRING, 0); }
		public LiteralContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_literal; }
	}

	public final LiteralContext literal() throws RecognitionException {
		LiteralContext _localctx = new LiteralContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_literal);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(133);
			_la = _input.LA(1);
			if ( !(((((_la - 86)) & ~0x3f) == 0 && ((1L << (_la - 86)) & 15L) != 0)) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
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

	@SuppressWarnings("CheckReturnValue")
	public static class VariableContext extends ParserRuleContext {
		public TerminalNode IDENTIFIER() { return getToken(lua_grammar_antlr4Parser.IDENTIFIER, 0); }
		public VariableContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_variable; }
	}

	public final VariableContext variable() throws RecognitionException {
		VariableContext _localctx = new VariableContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_variable);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(135);
			match(IDENTIFIER);
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

	@SuppressWarnings("CheckReturnValue")
	public static class FunctionCallContext extends ParserRuleContext {
		public TerminalNode IDENTIFIER() { return getToken(lua_grammar_antlr4Parser.IDENTIFIER, 0); }
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public FunctionCallContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_functionCall; }
	}

	public final FunctionCallContext functionCall() throws RecognitionException {
		FunctionCallContext _localctx = new FunctionCallContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_functionCall);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(137);
			match(IDENTIFIER);
			setState(138);
			match(T__5);
			setState(147);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 1152921505143718464L) != 0) || ((((_la - 75)) & ~0x3f) == 0 && ((1L << (_la - 75)) & 31745L) != 0)) {
				{
				setState(139);
				expression(0);
				setState(144);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==T__7) {
					{
					{
					setState(140);
					match(T__7);
					setState(141);
					expression(0);
					}
					}
					setState(146);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				}
			}

			setState(149);
			match(T__6);
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

	@SuppressWarnings("CheckReturnValue")
	public static class TableConstructorContext extends ParserRuleContext {
		public List<FieldContext> field() {
			return getRuleContexts(FieldContext.class);
		}
		public FieldContext field(int i) {
			return getRuleContext(FieldContext.class,i);
		}
		public MetatableContext metatable() {
			return getRuleContext(MetatableContext.class,0);
		}
		public TableConstructorContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_tableConstructor; }
	}

	public final TableConstructorContext tableConstructor() throws RecognitionException {
		TableConstructorContext _localctx = new TableConstructorContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_tableConstructor);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(151);
			match(T__8);
			setState(160);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 1152921505143718464L) != 0) || ((((_la - 75)) & ~0x3f) == 0 && ((1L << (_la - 75)) & 31745L) != 0)) {
				{
				setState(152);
				field();
				setState(157);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,8,_ctx);
				while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
					if ( _alt==1 ) {
						{
						{
						setState(153);
						match(T__7);
						setState(154);
						field();
						}
						} 
					}
					setState(159);
					_errHandler.sync(this);
					_alt = getInterpreter().adaptivePredict(_input,8,_ctx);
				}
				}
			}

			setState(164);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__7) {
				{
				setState(162);
				match(T__7);
				setState(163);
				metatable();
				}
			}

			setState(166);
			match(T__9);
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

	@SuppressWarnings("CheckReturnValue")
	public static class MetatableContext extends ParserRuleContext {
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public MetatableContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_metatable; }
	}

	public final MetatableContext metatable() throws RecognitionException {
		MetatableContext _localctx = new MetatableContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_metatable);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(168);
			match(T__10);
			setState(169);
			match(T__0);
			setState(170);
			expression(0);
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

	@SuppressWarnings("CheckReturnValue")
	public static class MetamethodsContext extends ParserRuleContext {
		public List<MetamethodContext> metamethod() {
			return getRuleContexts(MetamethodContext.class);
		}
		public MetamethodContext metamethod(int i) {
			return getRuleContext(MetamethodContext.class,i);
		}
		public MetamethodsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_metamethods; }
	}

	public final MetamethodsContext metamethods() throws RecognitionException {
		MetamethodsContext _localctx = new MetamethodsContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_metamethods);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(172);
			metamethod();
			setState(177);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__7) {
				{
				{
				setState(173);
				match(T__7);
				setState(174);
				metamethod();
				}
				}
				setState(179);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
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

	@SuppressWarnings("CheckReturnValue")
	public static class LabelStatementContext extends ParserRuleContext {
		public TerminalNode IDENTIFIER() { return getToken(lua_grammar_antlr4Parser.IDENTIFIER, 0); }
		public LabelStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_labelStatement; }
	}

	public final LabelStatementContext labelStatement() throws RecognitionException {
		LabelStatementContext _localctx = new LabelStatementContext(_ctx, getState());
		enterRule(_localctx, 22, RULE_labelStatement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(180);
			match(T__11);
			setState(181);
			match(IDENTIFIER);
			setState(182);
			match(T__11);
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

	@SuppressWarnings("CheckReturnValue")
	public static class MetamethodContext extends ParserRuleContext {
		public MetamethodContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_metamethod; }
	}

	public final MetamethodContext metamethod() throws RecognitionException {
		MetamethodContext _localctx = new MetamethodContext(_ctx, getState());
		enterRule(_localctx, 24, RULE_metamethod);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(184);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 536862720L) != 0)) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
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

	@SuppressWarnings("CheckReturnValue")
	public static class FieldContext extends ParserRuleContext {
		public TerminalNode IDENTIFIER() { return getToken(lua_grammar_antlr4Parser.IDENTIFIER, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public FieldContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_field; }
	}

	public final FieldContext field() throws RecognitionException {
		FieldContext _localctx = new FieldContext(_ctx, getState());
		enterRule(_localctx, 26, RULE_field);
		try {
			setState(190);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,12,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(186);
				match(IDENTIFIER);
				setState(187);
				match(T__0);
				setState(188);
				expression(0);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(189);
				expression(0);
				}
				break;
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

	@SuppressWarnings("CheckReturnValue")
	public static class BinaryOperationContext extends ParserRuleContext {
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public ArithOpContext arithOp() {
			return getRuleContext(ArithOpContext.class,0);
		}
		public BitwiseOpContext bitwiseOp() {
			return getRuleContext(BitwiseOpContext.class,0);
		}
		public ComparisonOpContext comparisonOp() {
			return getRuleContext(ComparisonOpContext.class,0);
		}
		public LogicalOpContext logicalOp() {
			return getRuleContext(LogicalOpContext.class,0);
		}
		public ConcatOpContext concatOp() {
			return getRuleContext(ConcatOpContext.class,0);
		}
		public CompoundAssignOpContext compoundAssignOp() {
			return getRuleContext(CompoundAssignOpContext.class,0);
		}
		public CoalesceOpContext coalesceOp() {
			return getRuleContext(CoalesceOpContext.class,0);
		}
		public ShiftAssignOpContext shiftAssignOp() {
			return getRuleContext(ShiftAssignOpContext.class,0);
		}
		public IncrOpContext incrOp() {
			return getRuleContext(IncrOpContext.class,0);
		}
		public BinaryOperationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_binaryOperation; }
	}

	public final BinaryOperationContext binaryOperation() throws RecognitionException {
		BinaryOperationContext _localctx = new BinaryOperationContext(_ctx, getState());
		enterRule(_localctx, 28, RULE_binaryOperation);
		try {
			setState(228);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,13,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(192);
				expression(0);
				setState(193);
				arithOp();
				setState(194);
				expression(0);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(196);
				expression(0);
				setState(197);
				bitwiseOp();
				setState(198);
				expression(0);
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(200);
				expression(0);
				setState(201);
				comparisonOp();
				setState(202);
				expression(0);
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(204);
				expression(0);
				setState(205);
				logicalOp();
				setState(206);
				expression(0);
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(208);
				expression(0);
				setState(209);
				concatOp();
				setState(210);
				expression(0);
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(212);
				expression(0);
				setState(213);
				compoundAssignOp();
				setState(214);
				expression(0);
				}
				break;
			case 7:
				enterOuterAlt(_localctx, 7);
				{
				setState(216);
				expression(0);
				setState(217);
				coalesceOp();
				setState(218);
				expression(0);
				}
				break;
			case 8:
				enterOuterAlt(_localctx, 8);
				{
				setState(220);
				expression(0);
				setState(221);
				shiftAssignOp();
				setState(222);
				expression(0);
				}
				break;
			case 9:
				enterOuterAlt(_localctx, 9);
				{
				setState(224);
				expression(0);
				setState(225);
				incrOp();
				setState(226);
				expression(0);
				}
				break;
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

	@SuppressWarnings("CheckReturnValue")
	public static class UnaryOperationContext extends ParserRuleContext {
		public UnaryOpContext unaryOp() {
			return getRuleContext(UnaryOpContext.class,0);
		}
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public UnaryOperationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_unaryOperation; }
	}

	public final UnaryOperationContext unaryOperation() throws RecognitionException {
		UnaryOperationContext _localctx = new UnaryOperationContext(_ctx, getState());
		enterRule(_localctx, 30, RULE_unaryOperation);
		try {
			setState(235);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,14,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(230);
				unaryOp();
				setState(231);
				expression(0);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(233);
				match(T__28);
				setState(234);
				expression(0);
				}
				break;
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

	@SuppressWarnings("CheckReturnValue")
	public static class ControlFlowStatementContext extends ParserRuleContext {
		public IfStatementContext ifStatement() {
			return getRuleContext(IfStatementContext.class,0);
		}
		public WhileStatementContext whileStatement() {
			return getRuleContext(WhileStatementContext.class,0);
		}
		public RepeatStatementContext repeatStatement() {
			return getRuleContext(RepeatStatementContext.class,0);
		}
		public ForStatementContext forStatement() {
			return getRuleContext(ForStatementContext.class,0);
		}
		public BreakStatementContext breakStatement() {
			return getRuleContext(BreakStatementContext.class,0);
		}
		public GotoStatementContext gotoStatement() {
			return getRuleContext(GotoStatementContext.class,0);
		}
		public CoroutineStatementContext coroutineStatement() {
			return getRuleContext(CoroutineStatementContext.class,0);
		}
		public ControlFlowStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_controlFlowStatement; }
	}

	public final ControlFlowStatementContext controlFlowStatement() throws RecognitionException {
		ControlFlowStatementContext _localctx = new ControlFlowStatementContext(_ctx, getState());
		enterRule(_localctx, 32, RULE_controlFlowStatement);
		try {
			setState(244);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__29:
				enterOuterAlt(_localctx, 1);
				{
				setState(237);
				ifStatement();
				}
				break;
			case T__34:
				enterOuterAlt(_localctx, 2);
				{
				setState(238);
				whileStatement();
				}
				break;
			case T__36:
				enterOuterAlt(_localctx, 3);
				{
				setState(239);
				repeatStatement();
				}
				break;
			case T__38:
				enterOuterAlt(_localctx, 4);
				{
				setState(240);
				forStatement();
				}
				break;
			case T__39:
				enterOuterAlt(_localctx, 5);
				{
				setState(241);
				breakStatement();
				}
				break;
			case T__40:
				enterOuterAlt(_localctx, 6);
				{
				setState(242);
				gotoStatement();
				}
				break;
			case T__41:
				enterOuterAlt(_localctx, 7);
				{
				setState(243);
				coroutineStatement();
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

	@SuppressWarnings("CheckReturnValue")
	public static class IfStatementContext extends ParserRuleContext {
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public List<BlockContext> block() {
			return getRuleContexts(BlockContext.class);
		}
		public BlockContext block(int i) {
			return getRuleContext(BlockContext.class,i);
		}
		public IfStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_ifStatement; }
	}

	public final IfStatementContext ifStatement() throws RecognitionException {
		IfStatementContext _localctx = new IfStatementContext(_ctx, getState());
		enterRule(_localctx, 34, RULE_ifStatement);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(246);
			match(T__29);
			setState(247);
			expression(0);
			setState(248);
			match(T__30);
			setState(249);
			block();
			setState(257);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__31) {
				{
				{
				setState(250);
				match(T__31);
				setState(251);
				expression(0);
				setState(252);
				match(T__30);
				setState(253);
				block();
				}
				}
				setState(259);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(262);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__32) {
				{
				setState(260);
				match(T__32);
				setState(261);
				block();
				}
			}

			setState(264);
			match(T__33);
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

	@SuppressWarnings("CheckReturnValue")
	public static class WhileStatementContext extends ParserRuleContext {
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public WhileStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_whileStatement; }
	}

	public final WhileStatementContext whileStatement() throws RecognitionException {
		WhileStatementContext _localctx = new WhileStatementContext(_ctx, getState());
		enterRule(_localctx, 36, RULE_whileStatement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(266);
			match(T__34);
			setState(267);
			expression(0);
			setState(268);
			match(T__35);
			setState(269);
			block();
			setState(270);
			match(T__33);
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

	@SuppressWarnings("CheckReturnValue")
	public static class RepeatStatementContext extends ParserRuleContext {
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public RepeatStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_repeatStatement; }
	}

	public final RepeatStatementContext repeatStatement() throws RecognitionException {
		RepeatStatementContext _localctx = new RepeatStatementContext(_ctx, getState());
		enterRule(_localctx, 38, RULE_repeatStatement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(272);
			match(T__36);
			setState(273);
			block();
			setState(274);
			match(T__37);
			setState(275);
			expression(0);
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

	@SuppressWarnings("CheckReturnValue")
	public static class ForStatementContext extends ParserRuleContext {
		public TerminalNode IDENTIFIER() { return getToken(lua_grammar_antlr4Parser.IDENTIFIER, 0); }
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public ForStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_forStatement; }
	}

	public final ForStatementContext forStatement() throws RecognitionException {
		ForStatementContext _localctx = new ForStatementContext(_ctx, getState());
		enterRule(_localctx, 40, RULE_forStatement);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(277);
			match(T__38);
			setState(278);
			match(IDENTIFIER);
			setState(279);
			match(T__0);
			setState(280);
			expression(0);
			setState(281);
			match(T__7);
			setState(282);
			expression(0);
			setState(285);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__7) {
				{
				setState(283);
				match(T__7);
				setState(284);
				expression(0);
				}
			}

			setState(287);
			match(T__35);
			setState(288);
			block();
			setState(289);
			match(T__33);
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

	@SuppressWarnings("CheckReturnValue")
	public static class BreakStatementContext extends ParserRuleContext {
		public BreakStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_breakStatement; }
	}

	public final BreakStatementContext breakStatement() throws RecognitionException {
		BreakStatementContext _localctx = new BreakStatementContext(_ctx, getState());
		enterRule(_localctx, 42, RULE_breakStatement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(291);
			match(T__39);
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

	@SuppressWarnings("CheckReturnValue")
	public static class GotoStatementContext extends ParserRuleContext {
		public TerminalNode IDENTIFIER() { return getToken(lua_grammar_antlr4Parser.IDENTIFIER, 0); }
		public GotoStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_gotoStatement; }
	}

	public final GotoStatementContext gotoStatement() throws RecognitionException {
		GotoStatementContext _localctx = new GotoStatementContext(_ctx, getState());
		enterRule(_localctx, 44, RULE_gotoStatement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(293);
			match(T__40);
			setState(294);
			match(IDENTIFIER);
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

	@SuppressWarnings("CheckReturnValue")
	public static class CoroutineStatementContext extends ParserRuleContext {
		public CoroutineStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_coroutineStatement; }
	}

	public final CoroutineStatementContext coroutineStatement() throws RecognitionException {
		CoroutineStatementContext _localctx = new CoroutineStatementContext(_ctx, getState());
		enterRule(_localctx, 46, RULE_coroutineStatement);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(296);
			match(T__41);
			setState(297);
			match(T__42);
			setState(298);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 263882790666240L) != 0)) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
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

	@SuppressWarnings("CheckReturnValue")
	public static class BlockContext extends ParserRuleContext {
		public List<StatementContext> statement() {
			return getRuleContexts(StatementContext.class);
		}
		public StatementContext statement(int i) {
			return getRuleContext(StatementContext.class,i);
		}
		public BlockContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_block; }
	}

	public final BlockContext block() throws RecognitionException {
		BlockContext _localctx = new BlockContext(_ctx, getState());
		enterRule(_localctx, 48, RULE_block);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(303);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 1978744046620672L) != 0) || _la==IDENTIFIER) {
				{
				{
				setState(300);
				statement();
				}
				}
				setState(305);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
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

	@SuppressWarnings("CheckReturnValue")
	public static class LocalDeclarationContext extends ParserRuleContext {
		public TerminalNode IDENTIFIER() { return getToken(lua_grammar_antlr4Parser.IDENTIFIER, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public LocalDeclarationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_localDeclaration; }
	}

	public final LocalDeclarationContext localDeclaration() throws RecognitionException {
		LocalDeclarationContext _localctx = new LocalDeclarationContext(_ctx, getState());
		enterRule(_localctx, 50, RULE_localDeclaration);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(306);
			match(T__47);
			setState(307);
			match(IDENTIFIER);
			setState(310);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__0) {
				{
				setState(308);
				match(T__0);
				setState(309);
				expression(0);
				}
			}

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

	@SuppressWarnings("CheckReturnValue")
	public static class FunctionDeclarationContext extends ParserRuleContext {
		public List<TerminalNode> IDENTIFIER() { return getTokens(lua_grammar_antlr4Parser.IDENTIFIER); }
		public TerminalNode IDENTIFIER(int i) {
			return getToken(lua_grammar_antlr4Parser.IDENTIFIER, i);
		}
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public FunctionDeclarationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_functionDeclaration; }
	}

	public final FunctionDeclarationContext functionDeclaration() throws RecognitionException {
		FunctionDeclarationContext _localctx = new FunctionDeclarationContext(_ctx, getState());
		enterRule(_localctx, 52, RULE_functionDeclaration);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(312);
			match(T__48);
			setState(313);
			match(IDENTIFIER);
			setState(314);
			match(T__5);
			setState(323);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==IDENTIFIER) {
				{
				setState(315);
				match(IDENTIFIER);
				setState(320);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==T__7) {
					{
					{
					setState(316);
					match(T__7);
					setState(317);
					match(IDENTIFIER);
					}
					}
					setState(322);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				}
			}

			setState(325);
			match(T__6);
			setState(326);
			block();
			setState(327);
			match(T__33);
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

	@SuppressWarnings("CheckReturnValue")
	public static class ReturnStatementContext extends ParserRuleContext {
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public ReturnStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_returnStatement; }
	}

	public final ReturnStatementContext returnStatement() throws RecognitionException {
		ReturnStatementContext _localctx = new ReturnStatementContext(_ctx, getState());
		enterRule(_localctx, 54, RULE_returnStatement);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(329);
			match(T__49);
			setState(338);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,24,_ctx) ) {
			case 1:
				{
				setState(330);
				expression(0);
				setState(335);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==T__7) {
					{
					{
					setState(331);
					match(T__7);
					setState(332);
					expression(0);
					}
					}
					setState(337);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				}
				break;
			}
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

	@SuppressWarnings("CheckReturnValue")
	public static class OperatorGroupContext extends ParserRuleContext {
		public LogicalOpContext logicalOp() {
			return getRuleContext(LogicalOpContext.class,0);
		}
		public ComparisonOpContext comparisonOp() {
			return getRuleContext(ComparisonOpContext.class,0);
		}
		public ArithOpContext arithOp() {
			return getRuleContext(ArithOpContext.class,0);
		}
		public BitwiseOpContext bitwiseOp() {
			return getRuleContext(BitwiseOpContext.class,0);
		}
		public AssignOpContext assignOp() {
			return getRuleContext(AssignOpContext.class,0);
		}
		public UnaryOpContext unaryOp() {
			return getRuleContext(UnaryOpContext.class,0);
		}
		public ConcatOpContext concatOp() {
			return getRuleContext(ConcatOpContext.class,0);
		}
		public CompoundAssignOpContext compoundAssignOp() {
			return getRuleContext(CompoundAssignOpContext.class,0);
		}
		public IncrOpContext incrOp() {
			return getRuleContext(IncrOpContext.class,0);
		}
		public CoalesceOpContext coalesceOp() {
			return getRuleContext(CoalesceOpContext.class,0);
		}
		public ShiftAssignOpContext shiftAssignOp() {
			return getRuleContext(ShiftAssignOpContext.class,0);
		}
		public OperatorGroupContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_operatorGroup; }
	}

	public final OperatorGroupContext operatorGroup() throws RecognitionException {
		OperatorGroupContext _localctx = new OperatorGroupContext(_ctx, getState());
		enterRule(_localctx, 56, RULE_operatorGroup);
		try {
			setState(351);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,25,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(340);
				logicalOp();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(341);
				comparisonOp();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(342);
				arithOp();
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(343);
				bitwiseOp();
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(344);
				assignOp();
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(345);
				unaryOp();
				}
				break;
			case 7:
				enterOuterAlt(_localctx, 7);
				{
				setState(346);
				concatOp();
				}
				break;
			case 8:
				enterOuterAlt(_localctx, 8);
				{
				setState(347);
				compoundAssignOp();
				}
				break;
			case 9:
				enterOuterAlt(_localctx, 9);
				{
				setState(348);
				incrOp();
				}
				break;
			case 10:
				enterOuterAlt(_localctx, 10);
				{
				setState(349);
				coalesceOp();
				}
				break;
			case 11:
				enterOuterAlt(_localctx, 11);
				{
				setState(350);
				shiftAssignOp();
				}
				break;
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

	@SuppressWarnings("CheckReturnValue")
	public static class LogicalOpContext extends ParserRuleContext {
		public LogicalOpContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_logicalOp; }
	}

	public final LogicalOpContext logicalOp() throws RecognitionException {
		LogicalOpContext _localctx = new LogicalOpContext(_ctx, getState());
		enterRule(_localctx, 58, RULE_logicalOp);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(353);
			_la = _input.LA(1);
			if ( !(_la==T__50 || _la==T__51) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
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

	@SuppressWarnings("CheckReturnValue")
	public static class ComparisonOpContext extends ParserRuleContext {
		public ComparisonOpContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_comparisonOp; }
	}

	public final ComparisonOpContext comparisonOp() throws RecognitionException {
		ComparisonOpContext _localctx = new ComparisonOpContext(_ctx, getState());
		enterRule(_localctx, 60, RULE_comparisonOp);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(355);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 567453553048682496L) != 0)) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
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

	@SuppressWarnings("CheckReturnValue")
	public static class ArithOpContext extends ParserRuleContext {
		public ArithOpContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_arithOp; }
	}

	public final ArithOpContext arithOp() throws RecognitionException {
		ArithOpContext _localctx = new ArithOpContext(_ctx, getState());
		enterRule(_localctx, 62, RULE_arithOp);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(357);
			_la = _input.LA(1);
			if ( !(((((_la - 59)) & ~0x3f) == 0 && ((1L << (_la - 59)) & 127L) != 0)) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
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

	@SuppressWarnings("CheckReturnValue")
	public static class BitwiseOpContext extends ParserRuleContext {
		public BitwiseOpContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_bitwiseOp; }
	}

	public final BitwiseOpContext bitwiseOp() throws RecognitionException {
		BitwiseOpContext _localctx = new BitwiseOpContext(_ctx, getState());
		enterRule(_localctx, 64, RULE_bitwiseOp);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(359);
			_la = _input.LA(1);
			if ( !(((((_la - 66)) & ~0x3f) == 0 && ((1L << (_la - 66)) & 31L) != 0)) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
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

	@SuppressWarnings("CheckReturnValue")
	public static class AssignOpContext extends ParserRuleContext {
		public AssignOpContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_assignOp; }
	}

	public final AssignOpContext assignOp() throws RecognitionException {
		AssignOpContext _localctx = new AssignOpContext(_ctx, getState());
		enterRule(_localctx, 66, RULE_assignOp);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(361);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 62L) != 0) || ((((_la - 71)) & ~0x3f) == 0 && ((1L << (_la - 71)) & 15L) != 0)) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
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

	@SuppressWarnings("CheckReturnValue")
	public static class UnaryOpContext extends ParserRuleContext {
		public UnaryOpContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_unaryOp; }
	}

	public final UnaryOpContext unaryOp() throws RecognitionException {
		UnaryOpContext _localctx = new UnaryOpContext(_ctx, getState());
		enterRule(_localctx, 68, RULE_unaryOp);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(363);
			_la = _input.LA(1);
			if ( !(((((_la - 29)) & ~0x3f) == 0 && ((1L << (_la - 29)) & 70370891661313L) != 0)) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
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

	@SuppressWarnings("CheckReturnValue")
	public static class ConcatOpContext extends ParserRuleContext {
		public ConcatOpContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_concatOp; }
	}

	public final ConcatOpContext concatOp() throws RecognitionException {
		ConcatOpContext _localctx = new ConcatOpContext(_ctx, getState());
		enterRule(_localctx, 70, RULE_concatOp);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(365);
			match(T__75);
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

	@SuppressWarnings("CheckReturnValue")
	public static class VarargOpContext extends ParserRuleContext {
		public VarargOpContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_varargOp; }
	}

	public final VarargOpContext varargOp() throws RecognitionException {
		VarargOpContext _localctx = new VarargOpContext(_ctx, getState());
		enterRule(_localctx, 72, RULE_varargOp);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(367);
			match(T__76);
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

	@SuppressWarnings("CheckReturnValue")
	public static class CompoundAssignOpContext extends ParserRuleContext {
		public CompoundAssignOpContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_compoundAssignOp; }
	}

	public final CompoundAssignOpContext compoundAssignOp() throws RecognitionException {
		CompoundAssignOpContext _localctx = new CompoundAssignOpContext(_ctx, getState());
		enterRule(_localctx, 74, RULE_compoundAssignOp);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(369);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 60L) != 0) || ((((_la - 71)) & ~0x3f) == 0 && ((1L << (_la - 71)) & 387L) != 0)) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
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

	@SuppressWarnings("CheckReturnValue")
	public static class IncrOpContext extends ParserRuleContext {
		public IncrOpContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_incrOp; }
	}

	public final IncrOpContext incrOp() throws RecognitionException {
		IncrOpContext _localctx = new IncrOpContext(_ctx, getState());
		enterRule(_localctx, 76, RULE_incrOp);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(371);
			_la = _input.LA(1);
			if ( !(_la==T__79 || _la==T__80) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
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

	@SuppressWarnings("CheckReturnValue")
	public static class CoalesceOpContext extends ParserRuleContext {
		public CoalesceOpContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_coalesceOp; }
	}

	public final CoalesceOpContext coalesceOp() throws RecognitionException {
		CoalesceOpContext _localctx = new CoalesceOpContext(_ctx, getState());
		enterRule(_localctx, 78, RULE_coalesceOp);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(373);
			match(T__81);
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

	@SuppressWarnings("CheckReturnValue")
	public static class ShiftAssignOpContext extends ParserRuleContext {
		public ShiftAssignOpContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_shiftAssignOp; }
	}

	public final ShiftAssignOpContext shiftAssignOp() throws RecognitionException {
		ShiftAssignOpContext _localctx = new ShiftAssignOpContext(_ctx, getState());
		enterRule(_localctx, 80, RULE_shiftAssignOp);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(375);
			_la = _input.LA(1);
			if ( !(_la==T__82 || _la==T__83) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
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

	public boolean sempred(RuleContext _localctx, int ruleIndex, int predIndex) {
		switch (ruleIndex) {
		case 3:
			return expression_sempred((ExpressionContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean expression_sempred(ExpressionContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 3);
		case 1:
			return precpred(_ctx, 2);
		}
		return true;
	}

	public static final String _serializedATN =
		"\u0004\u0001[\u017a\u0002\u0000\u0007\u0000\u0002\u0001\u0007\u0001\u0002"+
		"\u0002\u0007\u0002\u0002\u0003\u0007\u0003\u0002\u0004\u0007\u0004\u0002"+
		"\u0005\u0007\u0005\u0002\u0006\u0007\u0006\u0002\u0007\u0007\u0007\u0002"+
		"\b\u0007\b\u0002\t\u0007\t\u0002\n\u0007\n\u0002\u000b\u0007\u000b\u0002"+
		"\f\u0007\f\u0002\r\u0007\r\u0002\u000e\u0007\u000e\u0002\u000f\u0007\u000f"+
		"\u0002\u0010\u0007\u0010\u0002\u0011\u0007\u0011\u0002\u0012\u0007\u0012"+
		"\u0002\u0013\u0007\u0013\u0002\u0014\u0007\u0014\u0002\u0015\u0007\u0015"+
		"\u0002\u0016\u0007\u0016\u0002\u0017\u0007\u0017\u0002\u0018\u0007\u0018"+
		"\u0002\u0019\u0007\u0019\u0002\u001a\u0007\u001a\u0002\u001b\u0007\u001b"+
		"\u0002\u001c\u0007\u001c\u0002\u001d\u0007\u001d\u0002\u001e\u0007\u001e"+
		"\u0002\u001f\u0007\u001f\u0002 \u0007 \u0002!\u0007!\u0002\"\u0007\"\u0002"+
		"#\u0007#\u0002$\u0007$\u0002%\u0007%\u0002&\u0007&\u0002\'\u0007\'\u0002"+
		"(\u0007(\u0001\u0000\u0005\u0000T\b\u0000\n\u0000\f\u0000W\t\u0000\u0001"+
		"\u0000\u0001\u0000\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001"+
		"\u0001\u0001\u0001\u0003\u0001a\b\u0001\u0001\u0002\u0001\u0002\u0001"+
		"\u0002\u0001\u0002\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001"+
		"\u0003\u0003\u0003l\b\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001"+
		"\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0005\u0003v\b"+
		"\u0003\n\u0003\f\u0003y\t\u0003\u0001\u0004\u0001\u0004\u0001\u0004\u0001"+
		"\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0003"+
		"\u0004\u0084\b\u0004\u0001\u0005\u0001\u0005\u0001\u0006\u0001\u0006\u0001"+
		"\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0005\u0007\u008f"+
		"\b\u0007\n\u0007\f\u0007\u0092\t\u0007\u0003\u0007\u0094\b\u0007\u0001"+
		"\u0007\u0001\u0007\u0001\b\u0001\b\u0001\b\u0001\b\u0005\b\u009c\b\b\n"+
		"\b\f\b\u009f\t\b\u0003\b\u00a1\b\b\u0001\b\u0001\b\u0003\b\u00a5\b\b\u0001"+
		"\b\u0001\b\u0001\t\u0001\t\u0001\t\u0001\t\u0001\n\u0001\n\u0001\n\u0005"+
		"\n\u00b0\b\n\n\n\f\n\u00b3\t\n\u0001\u000b\u0001\u000b\u0001\u000b\u0001"+
		"\u000b\u0001\f\u0001\f\u0001\r\u0001\r\u0001\r\u0001\r\u0003\r\u00bf\b"+
		"\r\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e"+
		"\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e"+
		"\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e"+
		"\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e"+
		"\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e"+
		"\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e"+
		"\u0003\u000e\u00e5\b\u000e\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u000f"+
		"\u0001\u000f\u0003\u000f\u00ec\b\u000f\u0001\u0010\u0001\u0010\u0001\u0010"+
		"\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0003\u0010\u00f5\b\u0010"+
		"\u0001\u0011\u0001\u0011\u0001\u0011\u0001\u0011\u0001\u0011\u0001\u0011"+
		"\u0001\u0011\u0001\u0011\u0001\u0011\u0005\u0011\u0100\b\u0011\n\u0011"+
		"\f\u0011\u0103\t\u0011\u0001\u0011\u0001\u0011\u0003\u0011\u0107\b\u0011"+
		"\u0001\u0011\u0001\u0011\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012"+
		"\u0001\u0012\u0001\u0012\u0001\u0013\u0001\u0013\u0001\u0013\u0001\u0013"+
		"\u0001\u0013\u0001\u0014\u0001\u0014\u0001\u0014\u0001\u0014\u0001\u0014"+
		"\u0001\u0014\u0001\u0014\u0001\u0014\u0003\u0014\u011e\b\u0014\u0001\u0014"+
		"\u0001\u0014\u0001\u0014\u0001\u0014\u0001\u0015\u0001\u0015\u0001\u0016"+
		"\u0001\u0016\u0001\u0016\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017"+
		"\u0001\u0018\u0005\u0018\u012e\b\u0018\n\u0018\f\u0018\u0131\t\u0018\u0001"+
		"\u0019\u0001\u0019\u0001\u0019\u0001\u0019\u0003\u0019\u0137\b\u0019\u0001"+
		"\u001a\u0001\u001a\u0001\u001a\u0001\u001a\u0001\u001a\u0001\u001a\u0005"+
		"\u001a\u013f\b\u001a\n\u001a\f\u001a\u0142\t\u001a\u0003\u001a\u0144\b"+
		"\u001a\u0001\u001a\u0001\u001a\u0001\u001a\u0001\u001a\u0001\u001b\u0001"+
		"\u001b\u0001\u001b\u0001\u001b\u0005\u001b\u014e\b\u001b\n\u001b\f\u001b"+
		"\u0151\t\u001b\u0003\u001b\u0153\b\u001b\u0001\u001c\u0001\u001c\u0001"+
		"\u001c\u0001\u001c\u0001\u001c\u0001\u001c\u0001\u001c\u0001\u001c\u0001"+
		"\u001c\u0001\u001c\u0001\u001c\u0003\u001c\u0160\b\u001c\u0001\u001d\u0001"+
		"\u001d\u0001\u001e\u0001\u001e\u0001\u001f\u0001\u001f\u0001 \u0001 \u0001"+
		"!\u0001!\u0001\"\u0001\"\u0001#\u0001#\u0001$\u0001$\u0001%\u0001%\u0001"+
		"&\u0001&\u0001\'\u0001\'\u0001(\u0001(\u0001(\u0000\u0001\u0006)\u0000"+
		"\u0002\u0004\u0006\b\n\f\u000e\u0010\u0012\u0014\u0016\u0018\u001a\u001c"+
		"\u001e \"$&(*,.02468:<>@BDFHJLNP\u0000\r\u0001\u0000\u0001\u0005\u0001"+
		"\u0000VY\u0001\u0000\r\u001c\u0001\u0000,/\u0001\u000034\u0001\u00005"+
		":\u0001\u0000;A\u0001\u0000BF\u0002\u0000\u0001\u0005GJ\u0003\u0000\u001d"+
		"\u001d<<KK\u0003\u0000\u0002\u0005GHNO\u0001\u0000PQ\u0001\u0000ST\u0187"+
		"\u0000U\u0001\u0000\u0000\u0000\u0002`\u0001\u0000\u0000\u0000\u0004b"+
		"\u0001\u0000\u0000\u0000\u0006k\u0001\u0000\u0000\u0000\b\u0083\u0001"+
		"\u0000\u0000\u0000\n\u0085\u0001\u0000\u0000\u0000\f\u0087\u0001\u0000"+
		"\u0000\u0000\u000e\u0089\u0001\u0000\u0000\u0000\u0010\u0097\u0001\u0000"+
		"\u0000\u0000\u0012\u00a8\u0001\u0000\u0000\u0000\u0014\u00ac\u0001\u0000"+
		"\u0000\u0000\u0016\u00b4\u0001\u0000\u0000\u0000\u0018\u00b8\u0001\u0000"+
		"\u0000\u0000\u001a\u00be\u0001\u0000\u0000\u0000\u001c\u00e4\u0001\u0000"+
		"\u0000\u0000\u001e\u00eb\u0001\u0000\u0000\u0000 \u00f4\u0001\u0000\u0000"+
		"\u0000\"\u00f6\u0001\u0000\u0000\u0000$\u010a\u0001\u0000\u0000\u0000"+
		"&\u0110\u0001\u0000\u0000\u0000(\u0115\u0001\u0000\u0000\u0000*\u0123"+
		"\u0001\u0000\u0000\u0000,\u0125\u0001\u0000\u0000\u0000.\u0128\u0001\u0000"+
		"\u0000\u00000\u012f\u0001\u0000\u0000\u00002\u0132\u0001\u0000\u0000\u0000"+
		"4\u0138\u0001\u0000\u0000\u00006\u0149\u0001\u0000\u0000\u00008\u015f"+
		"\u0001\u0000\u0000\u0000:\u0161\u0001\u0000\u0000\u0000<\u0163\u0001\u0000"+
		"\u0000\u0000>\u0165\u0001\u0000\u0000\u0000@\u0167\u0001\u0000\u0000\u0000"+
		"B\u0169\u0001\u0000\u0000\u0000D\u016b\u0001\u0000\u0000\u0000F\u016d"+
		"\u0001\u0000\u0000\u0000H\u016f\u0001\u0000\u0000\u0000J\u0171\u0001\u0000"+
		"\u0000\u0000L\u0173\u0001\u0000\u0000\u0000N\u0175\u0001\u0000\u0000\u0000"+
		"P\u0177\u0001\u0000\u0000\u0000RT\u0003\u0002\u0001\u0000SR\u0001\u0000"+
		"\u0000\u0000TW\u0001\u0000\u0000\u0000US\u0001\u0000\u0000\u0000UV\u0001"+
		"\u0000\u0000\u0000VX\u0001\u0000\u0000\u0000WU\u0001\u0000\u0000\u0000"+
		"XY\u0005\u0000\u0000\u0001Y\u0001\u0001\u0000\u0000\u0000Za\u0003\u0004"+
		"\u0002\u0000[a\u0003 \u0010\u0000\\a\u00034\u001a\u0000]a\u00036\u001b"+
		"\u0000^a\u00032\u0019\u0000_a\u0003\u0016\u000b\u0000`Z\u0001\u0000\u0000"+
		"\u0000`[\u0001\u0000\u0000\u0000`\\\u0001\u0000\u0000\u0000`]\u0001\u0000"+
		"\u0000\u0000`^\u0001\u0000\u0000\u0000`_\u0001\u0000\u0000\u0000a\u0003"+
		"\u0001\u0000\u0000\u0000bc\u0003\f\u0006\u0000cd\u0007\u0000\u0000\u0000"+
		"de\u0003\u0006\u0003\u0000e\u0005\u0001\u0000\u0000\u0000fg\u0006\u0003"+
		"\uffff\uffff\u0000gl\u0003\b\u0004\u0000hi\u0003D\"\u0000ij\u0003\u0006"+
		"\u0003\u0001jl\u0001\u0000\u0000\u0000kf\u0001\u0000\u0000\u0000kh\u0001"+
		"\u0000\u0000\u0000lw\u0001\u0000\u0000\u0000mn\n\u0003\u0000\u0000no\u0003"+
		"8\u001c\u0000op\u0003\u0006\u0003\u0004pv\u0001\u0000\u0000\u0000qr\n"+
		"\u0002\u0000\u0000rs\u00038\u001c\u0000st\u0003\u0006\u0003\u0002tv\u0001"+
		"\u0000\u0000\u0000um\u0001\u0000\u0000\u0000uq\u0001\u0000\u0000\u0000"+
		"vy\u0001\u0000\u0000\u0000wu\u0001\u0000\u0000\u0000wx\u0001\u0000\u0000"+
		"\u0000x\u0007\u0001\u0000\u0000\u0000yw\u0001\u0000\u0000\u0000z\u0084"+
		"\u0003\n\u0005\u0000{\u0084\u0003\f\u0006\u0000|\u0084\u0003\u000e\u0007"+
		"\u0000}\u0084\u0003\u001e\u000f\u0000~\u0084\u0003\u0010\b\u0000\u007f"+
		"\u0080\u0005\u0006\u0000\u0000\u0080\u0081\u0003\u0006\u0003\u0000\u0081"+
		"\u0082\u0005\u0007\u0000\u0000\u0082\u0084\u0001\u0000\u0000\u0000\u0083"+
		"z\u0001\u0000\u0000\u0000\u0083{\u0001\u0000\u0000\u0000\u0083|\u0001"+
		"\u0000\u0000\u0000\u0083}\u0001\u0000\u0000\u0000\u0083~\u0001\u0000\u0000"+
		"\u0000\u0083\u007f\u0001\u0000\u0000\u0000\u0084\t\u0001\u0000\u0000\u0000"+
		"\u0085\u0086\u0007\u0001\u0000\u0000\u0086\u000b\u0001\u0000\u0000\u0000"+
		"\u0087\u0088\u0005U\u0000\u0000\u0088\r\u0001\u0000\u0000\u0000\u0089"+
		"\u008a\u0005U\u0000\u0000\u008a\u0093\u0005\u0006\u0000\u0000\u008b\u0090"+
		"\u0003\u0006\u0003\u0000\u008c\u008d\u0005\b\u0000\u0000\u008d\u008f\u0003"+
		"\u0006\u0003\u0000\u008e\u008c\u0001\u0000\u0000\u0000\u008f\u0092\u0001"+
		"\u0000\u0000\u0000\u0090\u008e\u0001\u0000\u0000\u0000\u0090\u0091\u0001"+
		"\u0000\u0000\u0000\u0091\u0094\u0001\u0000\u0000\u0000\u0092\u0090\u0001"+
		"\u0000\u0000\u0000\u0093\u008b\u0001\u0000\u0000\u0000\u0093\u0094\u0001"+
		"\u0000\u0000\u0000\u0094\u0095\u0001\u0000\u0000\u0000\u0095\u0096\u0005"+
		"\u0007\u0000\u0000\u0096\u000f\u0001\u0000\u0000\u0000\u0097\u00a0\u0005"+
		"\t\u0000\u0000\u0098\u009d\u0003\u001a\r\u0000\u0099\u009a\u0005\b\u0000"+
		"\u0000\u009a\u009c\u0003\u001a\r\u0000\u009b\u0099\u0001\u0000\u0000\u0000"+
		"\u009c\u009f\u0001\u0000\u0000\u0000\u009d\u009b\u0001\u0000\u0000\u0000"+
		"\u009d\u009e\u0001\u0000\u0000\u0000\u009e\u00a1\u0001\u0000\u0000\u0000"+
		"\u009f\u009d\u0001\u0000\u0000\u0000\u00a0\u0098\u0001\u0000\u0000\u0000"+
		"\u00a0\u00a1\u0001\u0000\u0000\u0000\u00a1\u00a4\u0001\u0000\u0000\u0000"+
		"\u00a2\u00a3\u0005\b\u0000\u0000\u00a3\u00a5\u0003\u0012\t\u0000\u00a4"+
		"\u00a2\u0001\u0000\u0000\u0000\u00a4\u00a5\u0001\u0000\u0000\u0000\u00a5"+
		"\u00a6\u0001\u0000\u0000\u0000\u00a6\u00a7\u0005\n\u0000\u0000\u00a7\u0011"+
		"\u0001\u0000\u0000\u0000\u00a8\u00a9\u0005\u000b\u0000\u0000\u00a9\u00aa"+
		"\u0005\u0001\u0000\u0000\u00aa\u00ab\u0003\u0006\u0003\u0000\u00ab\u0013"+
		"\u0001\u0000\u0000\u0000\u00ac\u00b1\u0003\u0018\f\u0000\u00ad\u00ae\u0005"+
		"\b\u0000\u0000\u00ae\u00b0\u0003\u0018\f\u0000\u00af\u00ad\u0001\u0000"+
		"\u0000\u0000\u00b0\u00b3\u0001\u0000\u0000\u0000\u00b1\u00af\u0001\u0000"+
		"\u0000\u0000\u00b1\u00b2\u0001\u0000\u0000\u0000\u00b2\u0015\u0001\u0000"+
		"\u0000\u0000\u00b3\u00b1\u0001\u0000\u0000\u0000\u00b4\u00b5\u0005\f\u0000"+
		"\u0000\u00b5\u00b6\u0005U\u0000\u0000\u00b6\u00b7\u0005\f\u0000\u0000"+
		"\u00b7\u0017\u0001\u0000\u0000\u0000\u00b8\u00b9\u0007\u0002\u0000\u0000"+
		"\u00b9\u0019\u0001\u0000\u0000\u0000\u00ba\u00bb\u0005U\u0000\u0000\u00bb"+
		"\u00bc\u0005\u0001\u0000\u0000\u00bc\u00bf\u0003\u0006\u0003\u0000\u00bd"+
		"\u00bf\u0003\u0006\u0003\u0000\u00be\u00ba\u0001\u0000\u0000\u0000\u00be"+
		"\u00bd\u0001\u0000\u0000\u0000\u00bf\u001b\u0001\u0000\u0000\u0000\u00c0"+
		"\u00c1\u0003\u0006\u0003\u0000\u00c1\u00c2\u0003>\u001f\u0000\u00c2\u00c3"+
		"\u0003\u0006\u0003\u0000\u00c3\u00e5\u0001\u0000\u0000\u0000\u00c4\u00c5"+
		"\u0003\u0006\u0003\u0000\u00c5\u00c6\u0003@ \u0000\u00c6\u00c7\u0003\u0006"+
		"\u0003\u0000\u00c7\u00e5\u0001\u0000\u0000\u0000\u00c8\u00c9\u0003\u0006"+
		"\u0003\u0000\u00c9\u00ca\u0003<\u001e\u0000\u00ca\u00cb\u0003\u0006\u0003"+
		"\u0000\u00cb\u00e5\u0001\u0000\u0000\u0000\u00cc\u00cd\u0003\u0006\u0003"+
		"\u0000\u00cd\u00ce\u0003:\u001d\u0000\u00ce\u00cf\u0003\u0006\u0003\u0000"+
		"\u00cf\u00e5\u0001\u0000\u0000\u0000\u00d0\u00d1\u0003\u0006\u0003\u0000"+
		"\u00d1\u00d2\u0003F#\u0000\u00d2\u00d3\u0003\u0006\u0003\u0000\u00d3\u00e5"+
		"\u0001\u0000\u0000\u0000\u00d4\u00d5\u0003\u0006\u0003\u0000\u00d5\u00d6"+
		"\u0003J%\u0000\u00d6\u00d7\u0003\u0006\u0003\u0000\u00d7\u00e5\u0001\u0000"+
		"\u0000\u0000\u00d8\u00d9\u0003\u0006\u0003\u0000\u00d9\u00da\u0003N\'"+
		"\u0000\u00da\u00db\u0003\u0006\u0003\u0000\u00db\u00e5\u0001\u0000\u0000"+
		"\u0000\u00dc\u00dd\u0003\u0006\u0003\u0000\u00dd\u00de\u0003P(\u0000\u00de"+
		"\u00df\u0003\u0006\u0003\u0000\u00df\u00e5\u0001\u0000\u0000\u0000\u00e0"+
		"\u00e1\u0003\u0006\u0003\u0000\u00e1\u00e2\u0003L&\u0000\u00e2\u00e3\u0003"+
		"\u0006\u0003\u0000\u00e3\u00e5\u0001\u0000\u0000\u0000\u00e4\u00c0\u0001"+
		"\u0000\u0000\u0000\u00e4\u00c4\u0001\u0000\u0000\u0000\u00e4\u00c8\u0001"+
		"\u0000\u0000\u0000\u00e4\u00cc\u0001\u0000\u0000\u0000\u00e4\u00d0\u0001"+
		"\u0000\u0000\u0000\u00e4\u00d4\u0001\u0000\u0000\u0000\u00e4\u00d8\u0001"+
		"\u0000\u0000\u0000\u00e4\u00dc\u0001\u0000\u0000\u0000\u00e4\u00e0\u0001"+
		"\u0000\u0000\u0000\u00e5\u001d\u0001\u0000\u0000\u0000\u00e6\u00e7\u0003"+
		"D\"\u0000\u00e7\u00e8\u0003\u0006\u0003\u0000\u00e8\u00ec\u0001\u0000"+
		"\u0000\u0000\u00e9\u00ea\u0005\u001d\u0000\u0000\u00ea\u00ec\u0003\u0006"+
		"\u0003\u0000\u00eb\u00e6\u0001\u0000\u0000\u0000\u00eb\u00e9\u0001\u0000"+
		"\u0000\u0000\u00ec\u001f\u0001\u0000\u0000\u0000\u00ed\u00f5\u0003\"\u0011"+
		"\u0000\u00ee\u00f5\u0003$\u0012\u0000\u00ef\u00f5\u0003&\u0013\u0000\u00f0"+
		"\u00f5\u0003(\u0014\u0000\u00f1\u00f5\u0003*\u0015\u0000\u00f2\u00f5\u0003"+
		",\u0016\u0000\u00f3\u00f5\u0003.\u0017\u0000\u00f4\u00ed\u0001\u0000\u0000"+
		"\u0000\u00f4\u00ee\u0001\u0000\u0000\u0000\u00f4\u00ef\u0001\u0000\u0000"+
		"\u0000\u00f4\u00f0\u0001\u0000\u0000\u0000\u00f4\u00f1\u0001\u0000\u0000"+
		"\u0000\u00f4\u00f2\u0001\u0000\u0000\u0000\u00f4\u00f3\u0001\u0000\u0000"+
		"\u0000\u00f5!\u0001\u0000\u0000\u0000\u00f6\u00f7\u0005\u001e\u0000\u0000"+
		"\u00f7\u00f8\u0003\u0006\u0003\u0000\u00f8\u00f9\u0005\u001f\u0000\u0000"+
		"\u00f9\u0101\u00030\u0018\u0000\u00fa\u00fb\u0005 \u0000\u0000\u00fb\u00fc"+
		"\u0003\u0006\u0003\u0000\u00fc\u00fd\u0005\u001f\u0000\u0000\u00fd\u00fe"+
		"\u00030\u0018\u0000\u00fe\u0100\u0001\u0000\u0000\u0000\u00ff\u00fa\u0001"+
		"\u0000\u0000\u0000\u0100\u0103\u0001\u0000\u0000\u0000\u0101\u00ff\u0001"+
		"\u0000\u0000\u0000\u0101\u0102\u0001\u0000\u0000\u0000\u0102\u0106\u0001"+
		"\u0000\u0000\u0000\u0103\u0101\u0001\u0000\u0000\u0000\u0104\u0105\u0005"+
		"!\u0000\u0000\u0105\u0107\u00030\u0018\u0000\u0106\u0104\u0001\u0000\u0000"+
		"\u0000\u0106\u0107\u0001\u0000\u0000\u0000\u0107\u0108\u0001\u0000\u0000"+
		"\u0000\u0108\u0109\u0005\"\u0000\u0000\u0109#\u0001\u0000\u0000\u0000"+
		"\u010a\u010b\u0005#\u0000\u0000\u010b\u010c\u0003\u0006\u0003\u0000\u010c"+
		"\u010d\u0005$\u0000\u0000\u010d\u010e\u00030\u0018\u0000\u010e\u010f\u0005"+
		"\"\u0000\u0000\u010f%\u0001\u0000\u0000\u0000\u0110\u0111\u0005%\u0000"+
		"\u0000\u0111\u0112\u00030\u0018\u0000\u0112\u0113\u0005&\u0000\u0000\u0113"+
		"\u0114\u0003\u0006\u0003\u0000\u0114\'\u0001\u0000\u0000\u0000\u0115\u0116"+
		"\u0005\'\u0000\u0000\u0116\u0117\u0005U\u0000\u0000\u0117\u0118\u0005"+
		"\u0001\u0000\u0000\u0118\u0119\u0003\u0006\u0003\u0000\u0119\u011a\u0005"+
		"\b\u0000\u0000\u011a\u011d\u0003\u0006\u0003\u0000\u011b\u011c\u0005\b"+
		"\u0000\u0000\u011c\u011e\u0003\u0006\u0003\u0000\u011d\u011b\u0001\u0000"+
		"\u0000\u0000\u011d\u011e\u0001\u0000\u0000\u0000\u011e\u011f\u0001\u0000"+
		"\u0000\u0000\u011f\u0120\u0005$\u0000\u0000\u0120\u0121\u00030\u0018\u0000"+
		"\u0121\u0122\u0005\"\u0000\u0000\u0122)\u0001\u0000\u0000\u0000\u0123"+
		"\u0124\u0005(\u0000\u0000\u0124+\u0001\u0000\u0000\u0000\u0125\u0126\u0005"+
		")\u0000\u0000\u0126\u0127\u0005U\u0000\u0000\u0127-\u0001\u0000\u0000"+
		"\u0000\u0128\u0129\u0005*\u0000\u0000\u0129\u012a\u0005+\u0000\u0000\u012a"+
		"\u012b\u0007\u0003\u0000\u0000\u012b/\u0001\u0000\u0000\u0000\u012c\u012e"+
		"\u0003\u0002\u0001\u0000\u012d\u012c\u0001\u0000\u0000\u0000\u012e\u0131"+
		"\u0001\u0000\u0000\u0000\u012f\u012d\u0001\u0000\u0000\u0000\u012f\u0130"+
		"\u0001\u0000\u0000\u0000\u01301\u0001\u0000\u0000\u0000\u0131\u012f\u0001"+
		"\u0000\u0000\u0000\u0132\u0133\u00050\u0000\u0000\u0133\u0136\u0005U\u0000"+
		"\u0000\u0134\u0135\u0005\u0001\u0000\u0000\u0135\u0137\u0003\u0006\u0003"+
		"\u0000\u0136\u0134\u0001\u0000\u0000\u0000\u0136\u0137\u0001\u0000\u0000"+
		"\u0000\u01373\u0001\u0000\u0000\u0000\u0138\u0139\u00051\u0000\u0000\u0139"+
		"\u013a\u0005U\u0000\u0000\u013a\u0143\u0005\u0006\u0000\u0000\u013b\u0140"+
		"\u0005U\u0000\u0000\u013c\u013d\u0005\b\u0000\u0000\u013d\u013f\u0005"+
		"U\u0000\u0000\u013e\u013c\u0001\u0000\u0000\u0000\u013f\u0142\u0001\u0000"+
		"\u0000\u0000\u0140\u013e\u0001\u0000\u0000\u0000\u0140\u0141\u0001\u0000"+
		"\u0000\u0000\u0141\u0144\u0001\u0000\u0000\u0000\u0142\u0140\u0001\u0000"+
		"\u0000\u0000\u0143\u013b\u0001\u0000\u0000\u0000\u0143\u0144\u0001\u0000"+
		"\u0000\u0000\u0144\u0145\u0001\u0000\u0000\u0000\u0145\u0146\u0005\u0007"+
		"\u0000\u0000\u0146\u0147\u00030\u0018\u0000\u0147\u0148\u0005\"\u0000"+
		"\u0000\u01485\u0001\u0000\u0000\u0000\u0149\u0152\u00052\u0000\u0000\u014a"+
		"\u014f\u0003\u0006\u0003\u0000\u014b\u014c\u0005\b\u0000\u0000\u014c\u014e"+
		"\u0003\u0006\u0003\u0000\u014d\u014b\u0001\u0000\u0000\u0000\u014e\u0151"+
		"\u0001\u0000\u0000\u0000\u014f\u014d\u0001\u0000\u0000\u0000\u014f\u0150"+
		"\u0001\u0000\u0000\u0000\u0150\u0153\u0001\u0000\u0000\u0000\u0151\u014f"+
		"\u0001\u0000\u0000\u0000\u0152\u014a\u0001\u0000\u0000\u0000\u0152\u0153"+
		"\u0001\u0000\u0000\u0000\u01537\u0001\u0000\u0000\u0000\u0154\u0160\u0003"+
		":\u001d\u0000\u0155\u0160\u0003<\u001e\u0000\u0156\u0160\u0003>\u001f"+
		"\u0000\u0157\u0160\u0003@ \u0000\u0158\u0160\u0003B!\u0000\u0159\u0160"+
		"\u0003D\"\u0000\u015a\u0160\u0003F#\u0000\u015b\u0160\u0003J%\u0000\u015c"+
		"\u0160\u0003L&\u0000\u015d\u0160\u0003N\'\u0000\u015e\u0160\u0003P(\u0000"+
		"\u015f\u0154\u0001\u0000\u0000\u0000\u015f\u0155\u0001\u0000\u0000\u0000"+
		"\u015f\u0156\u0001\u0000\u0000\u0000\u015f\u0157\u0001\u0000\u0000\u0000"+
		"\u015f\u0158\u0001\u0000\u0000\u0000\u015f\u0159\u0001\u0000\u0000\u0000"+
		"\u015f\u015a\u0001\u0000\u0000\u0000\u015f\u015b\u0001\u0000\u0000\u0000"+
		"\u015f\u015c\u0001\u0000\u0000\u0000\u015f\u015d\u0001\u0000\u0000\u0000"+
		"\u015f\u015e\u0001\u0000\u0000\u0000\u01609\u0001\u0000\u0000\u0000\u0161"+
		"\u0162\u0007\u0004\u0000\u0000\u0162;\u0001\u0000\u0000\u0000\u0163\u0164"+
		"\u0007\u0005\u0000\u0000\u0164=\u0001\u0000\u0000\u0000\u0165\u0166\u0007"+
		"\u0006\u0000\u0000\u0166?\u0001\u0000\u0000\u0000\u0167\u0168\u0007\u0007"+
		"\u0000\u0000\u0168A\u0001\u0000\u0000\u0000\u0169\u016a\u0007\b\u0000"+
		"\u0000\u016aC\u0001\u0000\u0000\u0000\u016b\u016c\u0007\t\u0000\u0000"+
		"\u016cE\u0001\u0000\u0000\u0000\u016d\u016e\u0005L\u0000\u0000\u016eG"+
		"\u0001\u0000\u0000\u0000\u016f\u0170\u0005M\u0000\u0000\u0170I\u0001\u0000"+
		"\u0000\u0000\u0171\u0172\u0007\n\u0000\u0000\u0172K\u0001\u0000\u0000"+
		"\u0000\u0173\u0174\u0007\u000b\u0000\u0000\u0174M\u0001\u0000\u0000\u0000"+
		"\u0175\u0176\u0005R\u0000\u0000\u0176O\u0001\u0000\u0000\u0000\u0177\u0178"+
		"\u0007\f\u0000\u0000\u0178Q\u0001\u0000\u0000\u0000\u001aU`kuw\u0083\u0090"+
		"\u0093\u009d\u00a0\u00a4\u00b1\u00be\u00e4\u00eb\u00f4\u0101\u0106\u011d"+
		"\u012f\u0136\u0140\u0143\u014f\u0152\u015f";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}