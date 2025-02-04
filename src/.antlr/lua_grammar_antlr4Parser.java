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
		T__80=81, T__81=82, T__82=83, T__83=84, T__84=85, T__85=86, T__86=87, 
		T__87=88, T__88=89, T__89=90, T__90=91, T__91=92, T__92=93, T__93=94, 
		T__94=95, IDENTIFIER=96, BOOL=97, NIL=98, NUMBER=99, STRING=100, WS=101, 
		COMMENT=102;
	public static final int
		RULE_program = 0, RULE_statement = 1, RULE_assignStatement = 2, RULE_expression = 3, 
		RULE_primaryExpression = 4, RULE_literal = 5, RULE_variable = 6, RULE_functionCall = 7, 
		RULE_tableConstructor = 8, RULE_metatable = 9, RULE_metamethods = 10, 
		RULE_labelStatement = 11, RULE_metamethod = 12, RULE_field = 13, RULE_binaryOperation = 14, 
		RULE_unaryOperation = 15, RULE_controlFlowStatement = 16, RULE_ifStatement = 17, 
		RULE_whileStatement = 18, RULE_repeatStatement = 19, RULE_forStatement = 20, 
		RULE_breakStatement = 21, RULE_gotoStatement = 22, RULE_coroutineStatement = 23, 
		RULE_protectedCallStatement = 24, RULE_block = 25, RULE_localDeclaration = 26, 
		RULE_functionDeclaration = 27, RULE_returnStatement = 28, RULE_operatorGroup = 29, 
		RULE_logicalOp = 30, RULE_comparisonOp = 31, RULE_arithOp = 32, RULE_bitwiseOp = 33, 
		RULE_assignOp = 34, RULE_unaryOp = 35, RULE_concatOp = 36, RULE_varargOp = 37, 
		RULE_compoundAssignOp = 38, RULE_incrOp = 39, RULE_coalesceOp = 40, RULE_shiftAssignOp = 41, 
		RULE_identifierList = 42, RULE_expressionList = 43, RULE_functionExpression = 44, 
		RULE_selectStatement = 45;
	private static String[] makeRuleNames() {
		return new String[] {
			"program", "statement", "assignStatement", "expression", "primaryExpression", 
			"literal", "variable", "functionCall", "tableConstructor", "metatable", 
			"metamethods", "labelStatement", "metamethod", "field", "binaryOperation", 
			"unaryOperation", "controlFlowStatement", "ifStatement", "whileStatement", 
			"repeatStatement", "forStatement", "breakStatement", "gotoStatement", 
			"coroutineStatement", "protectedCallStatement", "block", "localDeclaration", 
			"functionDeclaration", "returnStatement", "operatorGroup", "logicalOp", 
			"comparisonOp", "arithOp", "bitwiseOp", "assignOp", "unaryOp", "concatOp", 
			"varargOp", "compoundAssignOp", "incrOp", "coalesceOp", "shiftAssignOp", 
			"identifierList", "expressionList", "functionExpression", "selectStatement"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'='", "'+='", "'-='", "'*='", "'/='", "'('", "')'", "','", "'{'", 
			"'}'", "'__metatable'", "'::'", "'__add'", "'__sub'", "'__mul'", "'__div'", 
			"'__mod'", "'__pow'", "'__unm'", "'__concat'", "'__len'", "'__eq'", "'__lt'", 
			"'__le'", "'__tostring'", "'__pairs'", "'__ipairs'", "'__call'", "'['", 
			"']'", "':'", "'#'", "'if'", "'then'", "'elseif'", "'else'", "'end'", 
			"'while'", "'do'", "'repeat'", "'until'", "'for'", "'in'", "'break'", 
			"'goto'", "'coroutine'", "'.'", "'create'", "'resume'", "'yield'", "'status'", 
			"'running'", "'wrap'", "'isyieldable'", "'pcall'", "'xpcall'", "'local'", 
			"'function'", "'return'", "'=>'", "'and'", "'or'", "'=='", "'>='", "'<='", 
			"'~='", "'>'", "'<'", "'+'", "'-'", "'*'", "'/'", "'//'", "'%'", "'^'", 
			"'&'", "'|'", "'~'", "'<<'", "'>>'", "'//='", "'^='", "'&='", "'|='", 
			"'not'", "'..'", "'...'", "'..='", "'??='", "'+_'", "'-_'", "'??'", "'<<='", 
			"'>>='", "'select'", null, null, "'nil'"
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
			null, null, null, null, null, null, null, null, null, null, null, null, 
			"IDENTIFIER", "BOOL", "NIL", "NUMBER", "STRING", "WS", "COMMENT"
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
			setState(95);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 1117021633916178432L) != 0) || _la==T__94 || _la==IDENTIFIER) {
				{
				{
				setState(92);
				statement();
				}
				}
				setState(97);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(98);
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
		public SelectStatementContext selectStatement() {
			return getRuleContext(SelectStatementContext.class,0);
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
			setState(107);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case IDENTIFIER:
				enterOuterAlt(_localctx, 1);
				{
				setState(100);
				assignStatement();
				}
				break;
			case T__32:
			case T__37:
			case T__39:
			case T__41:
			case T__43:
			case T__44:
			case T__45:
			case T__54:
			case T__55:
				enterOuterAlt(_localctx, 2);
				{
				setState(101);
				controlFlowStatement();
				}
				break;
			case T__57:
				enterOuterAlt(_localctx, 3);
				{
				setState(102);
				functionDeclaration();
				}
				break;
			case T__58:
				enterOuterAlt(_localctx, 4);
				{
				setState(103);
				returnStatement();
				}
				break;
			case T__56:
				enterOuterAlt(_localctx, 5);
				{
				setState(104);
				localDeclaration();
				}
				break;
			case T__11:
				enterOuterAlt(_localctx, 6);
				{
				setState(105);
				labelStatement();
				}
				break;
			case T__94:
				enterOuterAlt(_localctx, 7);
				{
				setState(106);
				selectStatement();
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
			setState(109);
			variable();
			setState(110);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 62L) != 0)) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			setState(111);
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
			setState(118);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,2,_ctx) ) {
			case 1:
				{
				setState(114);
				primaryExpression();
				}
				break;
			case 2:
				{
				setState(115);
				unaryOp();
				setState(116);
				expression(1);
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(130);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,4,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(128);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,3,_ctx) ) {
					case 1:
						{
						_localctx = new ExpressionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(120);
						if (!(precpred(_ctx, 3))) throw new FailedPredicateException(this, "precpred(_ctx, 3)");
						setState(121);
						operatorGroup();
						setState(122);
						expression(4);
						}
						break;
					case 2:
						{
						_localctx = new ExpressionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(124);
						if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
						setState(125);
						operatorGroup();
						setState(126);
						expression(2);
						}
						break;
					}
					} 
				}
				setState(132);
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
		public FunctionExpressionContext functionExpression() {
			return getRuleContext(FunctionExpressionContext.class,0);
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
			setState(143);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,5,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(133);
				literal();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(134);
				variable();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(135);
				functionCall();
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(136);
				unaryOperation();
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(137);
				tableConstructor();
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(138);
				functionExpression();
				}
				break;
			case 7:
				enterOuterAlt(_localctx, 7);
				{
				setState(139);
				match(T__5);
				setState(140);
				expression(0);
				setState(141);
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
			setState(145);
			_la = _input.LA(1);
			if ( !(((((_la - 97)) & ~0x3f) == 0 && ((1L << (_la - 97)) & 15L) != 0)) ) {
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
			setState(147);
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
			setState(149);
			match(IDENTIFIER);
			setState(150);
			match(T__5);
			setState(159);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 288230380446679616L) != 0) || ((((_la - 70)) & ~0x3f) == 0 && ((1L << (_la - 70)) & 2080407553L) != 0)) {
				{
				setState(151);
				expression(0);
				setState(156);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==T__7) {
					{
					{
					setState(152);
					match(T__7);
					setState(153);
					expression(0);
					}
					}
					setState(158);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				}
			}

			setState(161);
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
			setState(163);
			match(T__8);
			setState(172);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 288230380983550528L) != 0) || ((((_la - 70)) & ~0x3f) == 0 && ((1L << (_la - 70)) & 2080407553L) != 0)) {
				{
				setState(164);
				field();
				setState(169);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,8,_ctx);
				while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
					if ( _alt==1 ) {
						{
						{
						setState(165);
						match(T__7);
						setState(166);
						field();
						}
						} 
					}
					setState(171);
					_errHandler.sync(this);
					_alt = getInterpreter().adaptivePredict(_input,8,_ctx);
				}
				}
			}

			setState(176);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__7) {
				{
				setState(174);
				match(T__7);
				setState(175);
				metatable();
				}
			}

			setState(178);
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
			setState(180);
			match(T__10);
			setState(181);
			match(T__0);
			setState(182);
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
			setState(184);
			metamethod();
			setState(189);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__7) {
				{
				{
				setState(185);
				match(T__7);
				setState(186);
				metamethod();
				}
				}
				setState(191);
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
			setState(192);
			match(T__11);
			setState(193);
			match(IDENTIFIER);
			setState(194);
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
			setState(196);
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
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public FunctionExpressionContext functionExpression() {
			return getRuleContext(FunctionExpressionContext.class,0);
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
			setState(211);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,12,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(198);
				match(IDENTIFIER);
				setState(199);
				match(T__0);
				setState(200);
				expression(0);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(201);
				match(T__28);
				setState(202);
				expression(0);
				setState(203);
				match(T__29);
				setState(204);
				match(T__0);
				setState(205);
				expression(0);
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(207);
				expression(0);
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(208);
				match(IDENTIFIER);
				setState(209);
				match(T__30);
				setState(210);
				functionExpression();
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
			setState(249);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,13,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(213);
				expression(0);
				setState(214);
				arithOp();
				setState(215);
				expression(0);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(217);
				expression(0);
				setState(218);
				bitwiseOp();
				setState(219);
				expression(0);
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(221);
				expression(0);
				setState(222);
				comparisonOp();
				setState(223);
				expression(0);
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(225);
				expression(0);
				setState(226);
				logicalOp();
				setState(227);
				expression(0);
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(229);
				expression(0);
				setState(230);
				concatOp();
				setState(231);
				expression(0);
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(233);
				expression(0);
				setState(234);
				compoundAssignOp();
				setState(235);
				expression(0);
				}
				break;
			case 7:
				enterOuterAlt(_localctx, 7);
				{
				setState(237);
				expression(0);
				setState(238);
				coalesceOp();
				setState(239);
				expression(0);
				}
				break;
			case 8:
				enterOuterAlt(_localctx, 8);
				{
				setState(241);
				expression(0);
				setState(242);
				shiftAssignOp();
				setState(243);
				expression(0);
				}
				break;
			case 9:
				enterOuterAlt(_localctx, 9);
				{
				setState(245);
				expression(0);
				setState(246);
				incrOp();
				setState(247);
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
			setState(256);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,14,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(251);
				unaryOp();
				setState(252);
				expression(0);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(254);
				match(T__31);
				setState(255);
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
		public ProtectedCallStatementContext protectedCallStatement() {
			return getRuleContext(ProtectedCallStatementContext.class,0);
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
			setState(266);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__32:
				enterOuterAlt(_localctx, 1);
				{
				setState(258);
				ifStatement();
				}
				break;
			case T__37:
				enterOuterAlt(_localctx, 2);
				{
				setState(259);
				whileStatement();
				}
				break;
			case T__39:
				enterOuterAlt(_localctx, 3);
				{
				setState(260);
				repeatStatement();
				}
				break;
			case T__41:
				enterOuterAlt(_localctx, 4);
				{
				setState(261);
				forStatement();
				}
				break;
			case T__43:
				enterOuterAlt(_localctx, 5);
				{
				setState(262);
				breakStatement();
				}
				break;
			case T__44:
				enterOuterAlt(_localctx, 6);
				{
				setState(263);
				gotoStatement();
				}
				break;
			case T__45:
				enterOuterAlt(_localctx, 7);
				{
				setState(264);
				coroutineStatement();
				}
				break;
			case T__54:
			case T__55:
				enterOuterAlt(_localctx, 8);
				{
				setState(265);
				protectedCallStatement();
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
			setState(268);
			match(T__32);
			setState(269);
			expression(0);
			setState(270);
			match(T__33);
			setState(271);
			block();
			setState(279);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__34) {
				{
				{
				setState(272);
				match(T__34);
				setState(273);
				expression(0);
				setState(274);
				match(T__33);
				setState(275);
				block();
				}
				}
				setState(281);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(284);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__35) {
				{
				setState(282);
				match(T__35);
				setState(283);
				block();
				}
			}

			setState(286);
			match(T__36);
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
			setState(288);
			match(T__37);
			setState(289);
			expression(0);
			setState(290);
			match(T__38);
			setState(291);
			block();
			setState(292);
			match(T__36);
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
			setState(294);
			match(T__39);
			setState(295);
			block();
			setState(296);
			match(T__40);
			setState(297);
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
		public ForStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_forStatement; }
	 
		public ForStatementContext() { }
		public void copyFrom(ForStatementContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class GenericForContext extends ForStatementContext {
		public IdentifierListContext identifierList() {
			return getRuleContext(IdentifierListContext.class,0);
		}
		public ExpressionListContext expressionList() {
			return getRuleContext(ExpressionListContext.class,0);
		}
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public GenericForContext(ForStatementContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class NumericForContext extends ForStatementContext {
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
		public NumericForContext(ForStatementContext ctx) { copyFrom(ctx); }
	}

	public final ForStatementContext forStatement() throws RecognitionException {
		ForStatementContext _localctx = new ForStatementContext(_ctx, getState());
		enterRule(_localctx, 40, RULE_forStatement);
		int _la;
		try {
			setState(321);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,19,_ctx) ) {
			case 1:
				_localctx = new NumericForContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(299);
				match(T__41);
				setState(300);
				match(IDENTIFIER);
				setState(301);
				match(T__0);
				setState(302);
				expression(0);
				setState(303);
				match(T__7);
				setState(304);
				expression(0);
				setState(307);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==T__7) {
					{
					setState(305);
					match(T__7);
					setState(306);
					expression(0);
					}
				}

				setState(309);
				match(T__38);
				setState(310);
				block();
				setState(311);
				match(T__36);
				}
				break;
			case 2:
				_localctx = new GenericForContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(313);
				match(T__41);
				setState(314);
				identifierList();
				setState(315);
				match(T__42);
				setState(316);
				expressionList();
				setState(317);
				match(T__38);
				setState(318);
				block();
				setState(319);
				match(T__36);
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
			setState(323);
			match(T__43);
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
			setState(325);
			match(T__44);
			setState(326);
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
			setState(328);
			match(T__45);
			setState(329);
			match(T__46);
			setState(330);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 35747322042253312L) != 0)) ) {
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
	public static class ProtectedCallStatementContext extends ParserRuleContext {
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public TerminalNode IDENTIFIER() { return getToken(lua_grammar_antlr4Parser.IDENTIFIER, 0); }
		public ProtectedCallStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_protectedCallStatement; }
	}

	public final ProtectedCallStatementContext protectedCallStatement() throws RecognitionException {
		ProtectedCallStatementContext _localctx = new ProtectedCallStatementContext(_ctx, getState());
		enterRule(_localctx, 48, RULE_protectedCallStatement);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(332);
			_la = _input.LA(1);
			if ( !(_la==T__54 || _la==T__55) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			setState(334);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__30 || _la==T__46) {
				{
				setState(333);
				_la = _input.LA(1);
				if ( !(_la==T__30 || _la==T__46) ) {
				_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				}
			}

			setState(337);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==IDENTIFIER) {
				{
				setState(336);
				match(IDENTIFIER);
				}
			}

			}
			setState(339);
			match(T__5);
			setState(340);
			expression(0);
			setState(343);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__7) {
				{
				setState(341);
				match(T__7);
				setState(342);
				expression(0);
				}
			}

			setState(345);
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
		enterRule(_localctx, 50, RULE_block);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(350);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 1117021633916178432L) != 0) || _la==T__94 || _la==IDENTIFIER) {
				{
				{
				setState(347);
				statement();
				}
				}
				setState(352);
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
		public List<TerminalNode> IDENTIFIER() { return getTokens(lua_grammar_antlr4Parser.IDENTIFIER); }
		public TerminalNode IDENTIFIER(int i) {
			return getToken(lua_grammar_antlr4Parser.IDENTIFIER, i);
		}
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public ExpressionListContext expressionList() {
			return getRuleContext(ExpressionListContext.class,0);
		}
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public LocalDeclarationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_localDeclaration; }
	}

	public final LocalDeclarationContext localDeclaration() throws RecognitionException {
		LocalDeclarationContext _localctx = new LocalDeclarationContext(_ctx, getState());
		enterRule(_localctx, 52, RULE_localDeclaration);
		int _la;
		try {
			setState(388);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,28,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(353);
				match(T__56);
				setState(354);
				match(IDENTIFIER);
				setState(357);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==T__0) {
					{
					setState(355);
					match(T__0);
					setState(356);
					expression(0);
					}
				}

				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(359);
				match(T__56);
				setState(360);
				match(IDENTIFIER);
				setState(365);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==T__7) {
					{
					{
					setState(361);
					match(T__7);
					setState(362);
					match(IDENTIFIER);
					}
					}
					setState(367);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(368);
				match(T__0);
				setState(369);
				expressionList();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(370);
				match(T__56);
				setState(371);
				match(T__57);
				setState(372);
				match(IDENTIFIER);
				setState(373);
				match(T__5);
				setState(382);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==IDENTIFIER) {
					{
					setState(374);
					match(IDENTIFIER);
					setState(379);
					_errHandler.sync(this);
					_la = _input.LA(1);
					while (_la==T__7) {
						{
						{
						setState(375);
						match(T__7);
						setState(376);
						match(IDENTIFIER);
						}
						}
						setState(381);
						_errHandler.sync(this);
						_la = _input.LA(1);
					}
					}
				}

				setState(384);
				match(T__6);
				setState(385);
				block();
				setState(386);
				match(T__36);
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
	public static class FunctionDeclarationContext extends ParserRuleContext {
		public List<TerminalNode> IDENTIFIER() { return getTokens(lua_grammar_antlr4Parser.IDENTIFIER); }
		public TerminalNode IDENTIFIER(int i) {
			return getToken(lua_grammar_antlr4Parser.IDENTIFIER, i);
		}
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public VarargOpContext varargOp() {
			return getRuleContext(VarargOpContext.class,0);
		}
		public FunctionDeclarationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_functionDeclaration; }
	}

	public final FunctionDeclarationContext functionDeclaration() throws RecognitionException {
		FunctionDeclarationContext _localctx = new FunctionDeclarationContext(_ctx, getState());
		enterRule(_localctx, 54, RULE_functionDeclaration);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(390);
			match(T__57);
			setState(395);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,29,_ctx) ) {
			case 1:
				{
				setState(391);
				match(IDENTIFIER);
				setState(392);
				match(T__46);
				}
				break;
			case 2:
				{
				setState(393);
				match(IDENTIFIER);
				setState(394);
				match(T__30);
				}
				break;
			}
			setState(397);
			match(IDENTIFIER);
			setState(398);
			match(T__5);
			setState(412);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case IDENTIFIER:
				{
				setState(399);
				match(IDENTIFIER);
				setState(404);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,30,_ctx);
				while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
					if ( _alt==1 ) {
						{
						{
						setState(400);
						match(T__7);
						setState(401);
						match(IDENTIFIER);
						}
						} 
					}
					setState(406);
					_errHandler.sync(this);
					_alt = getInterpreter().adaptivePredict(_input,30,_ctx);
				}
				setState(409);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==T__7) {
					{
					setState(407);
					match(T__7);
					setState(408);
					varargOp();
					}
				}

				}
				break;
			case T__86:
				{
				setState(411);
				varargOp();
				}
				break;
			case T__6:
				break;
			default:
				break;
			}
			setState(414);
			match(T__6);
			setState(415);
			block();
			setState(416);
			match(T__36);
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
		public FunctionCallContext functionCall() {
			return getRuleContext(FunctionCallContext.class,0);
		}
		public ReturnStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_returnStatement; }
	}

	public final ReturnStatementContext returnStatement() throws RecognitionException {
		ReturnStatementContext _localctx = new ReturnStatementContext(_ctx, getState());
		enterRule(_localctx, 56, RULE_returnStatement);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(418);
			match(T__58);
			setState(428);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,34,_ctx) ) {
			case 1:
				{
				setState(419);
				expression(0);
				setState(424);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==T__7) {
					{
					{
					setState(420);
					match(T__7);
					setState(421);
					expression(0);
					}
					}
					setState(426);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				}
				break;
			case 2:
				{
				setState(427);
				functionCall();
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
		enterRule(_localctx, 58, RULE_operatorGroup);
		try {
			setState(442);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,35,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(430);
				logicalOp();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(431);
				comparisonOp();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(432);
				arithOp();
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(433);
				bitwiseOp();
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(434);
				assignOp();
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(435);
				unaryOp();
				}
				break;
			case 7:
				enterOuterAlt(_localctx, 7);
				{
				setState(436);
				concatOp();
				}
				break;
			case 8:
				enterOuterAlt(_localctx, 8);
				{
				setState(437);
				compoundAssignOp();
				}
				break;
			case 9:
				enterOuterAlt(_localctx, 9);
				{
				setState(438);
				incrOp();
				}
				break;
			case 10:
				enterOuterAlt(_localctx, 10);
				{
				setState(439);
				coalesceOp();
				}
				break;
			case 11:
				enterOuterAlt(_localctx, 11);
				{
				setState(440);
				shiftAssignOp();
				}
				break;
			case 12:
				enterOuterAlt(_localctx, 12);
				{
				setState(441);
				match(T__59);
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
		enterRule(_localctx, 60, RULE_logicalOp);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(444);
			_la = _input.LA(1);
			if ( !(_la==T__60 || _la==T__61) ) {
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
		enterRule(_localctx, 62, RULE_comparisonOp);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(446);
			_la = _input.LA(1);
			if ( !(((((_la - 63)) & ~0x3f) == 0 && ((1L << (_la - 63)) & 63L) != 0)) ) {
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
		enterRule(_localctx, 64, RULE_arithOp);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(448);
			_la = _input.LA(1);
			if ( !(((((_la - 69)) & ~0x3f) == 0 && ((1L << (_la - 69)) & 127L) != 0)) ) {
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
		enterRule(_localctx, 66, RULE_bitwiseOp);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(450);
			_la = _input.LA(1);
			if ( !(((((_la - 76)) & ~0x3f) == 0 && ((1L << (_la - 76)) & 31L) != 0)) ) {
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
		enterRule(_localctx, 68, RULE_assignOp);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(452);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 62L) != 0) || ((((_la - 81)) & ~0x3f) == 0 && ((1L << (_la - 81)) & 15L) != 0)) ) {
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
		enterRule(_localctx, 70, RULE_unaryOp);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(454);
			_la = _input.LA(1);
			if ( !(((((_la - 32)) & ~0x3f) == 0 && ((1L << (_la - 32)) & 9007474132647937L) != 0)) ) {
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
		enterRule(_localctx, 72, RULE_concatOp);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(456);
			match(T__85);
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
		enterRule(_localctx, 74, RULE_varargOp);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(458);
			match(T__86);
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
		enterRule(_localctx, 76, RULE_compoundAssignOp);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(460);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 60L) != 0) || ((((_la - 81)) & ~0x3f) == 0 && ((1L << (_la - 81)) & 387L) != 0)) ) {
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
		enterRule(_localctx, 78, RULE_incrOp);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(462);
			_la = _input.LA(1);
			if ( !(_la==T__89 || _la==T__90) ) {
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
		enterRule(_localctx, 80, RULE_coalesceOp);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(464);
			match(T__91);
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
		enterRule(_localctx, 82, RULE_shiftAssignOp);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(466);
			_la = _input.LA(1);
			if ( !(_la==T__92 || _la==T__93) ) {
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
	public static class IdentifierListContext extends ParserRuleContext {
		public List<TerminalNode> IDENTIFIER() { return getTokens(lua_grammar_antlr4Parser.IDENTIFIER); }
		public TerminalNode IDENTIFIER(int i) {
			return getToken(lua_grammar_antlr4Parser.IDENTIFIER, i);
		}
		public IdentifierListContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_identifierList; }
	}

	public final IdentifierListContext identifierList() throws RecognitionException {
		IdentifierListContext _localctx = new IdentifierListContext(_ctx, getState());
		enterRule(_localctx, 84, RULE_identifierList);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(468);
			match(IDENTIFIER);
			setState(473);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__7) {
				{
				{
				setState(469);
				match(T__7);
				setState(470);
				match(IDENTIFIER);
				}
				}
				setState(475);
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
	public static class ExpressionListContext extends ParserRuleContext {
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public ExpressionListContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expressionList; }
	}

	public final ExpressionListContext expressionList() throws RecognitionException {
		ExpressionListContext _localctx = new ExpressionListContext(_ctx, getState());
		enterRule(_localctx, 86, RULE_expressionList);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(476);
			expression(0);
			setState(481);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__7) {
				{
				{
				setState(477);
				match(T__7);
				setState(478);
				expression(0);
				}
				}
				setState(483);
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
	public static class FunctionExpressionContext extends ParserRuleContext {
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public List<TerminalNode> IDENTIFIER() { return getTokens(lua_grammar_antlr4Parser.IDENTIFIER); }
		public TerminalNode IDENTIFIER(int i) {
			return getToken(lua_grammar_antlr4Parser.IDENTIFIER, i);
		}
		public FunctionExpressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_functionExpression; }
	}

	public final FunctionExpressionContext functionExpression() throws RecognitionException {
		FunctionExpressionContext _localctx = new FunctionExpressionContext(_ctx, getState());
		enterRule(_localctx, 88, RULE_functionExpression);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(484);
			match(T__57);
			setState(487);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==IDENTIFIER) {
				{
				setState(485);
				match(IDENTIFIER);
				setState(486);
				match(T__30);
				}
			}

			setState(489);
			match(T__5);
			setState(498);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==IDENTIFIER) {
				{
				setState(490);
				match(IDENTIFIER);
				setState(495);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==T__7) {
					{
					{
					setState(491);
					match(T__7);
					setState(492);
					match(IDENTIFIER);
					}
					}
					setState(497);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				}
			}

			setState(500);
			match(T__6);
			setState(501);
			block();
			setState(502);
			match(T__36);
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
	public static class SelectStatementContext extends ParserRuleContext {
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public SelectStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_selectStatement; }
	}

	public final SelectStatementContext selectStatement() throws RecognitionException {
		SelectStatementContext _localctx = new SelectStatementContext(_ctx, getState());
		enterRule(_localctx, 90, RULE_selectStatement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(504);
			match(T__94);
			setState(505);
			match(T__5);
			setState(508);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,41,_ctx) ) {
			case 1:
				{
				setState(506);
				match(T__31);
				}
				break;
			case 2:
				{
				setState(507);
				expression(0);
				}
				break;
			}
			setState(510);
			match(T__7);
			setState(511);
			expression(0);
			setState(512);
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
		"\u0004\u0001f\u0203\u0002\u0000\u0007\u0000\u0002\u0001\u0007\u0001\u0002"+
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
		"(\u0007(\u0002)\u0007)\u0002*\u0007*\u0002+\u0007+\u0002,\u0007,\u0002"+
		"-\u0007-\u0001\u0000\u0005\u0000^\b\u0000\n\u0000\f\u0000a\t\u0000\u0001"+
		"\u0000\u0001\u0000\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001"+
		"\u0001\u0001\u0001\u0001\u0001\u0003\u0001l\b\u0001\u0001\u0002\u0001"+
		"\u0002\u0001\u0002\u0001\u0002\u0001\u0003\u0001\u0003\u0001\u0003\u0001"+
		"\u0003\u0001\u0003\u0003\u0003w\b\u0003\u0001\u0003\u0001\u0003\u0001"+
		"\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0005"+
		"\u0003\u0081\b\u0003\n\u0003\f\u0003\u0084\t\u0003\u0001\u0004\u0001\u0004"+
		"\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004"+
		"\u0001\u0004\u0001\u0004\u0003\u0004\u0090\b\u0004\u0001\u0005\u0001\u0005"+
		"\u0001\u0006\u0001\u0006\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007"+
		"\u0001\u0007\u0005\u0007\u009b\b\u0007\n\u0007\f\u0007\u009e\t\u0007\u0003"+
		"\u0007\u00a0\b\u0007\u0001\u0007\u0001\u0007\u0001\b\u0001\b\u0001\b\u0001"+
		"\b\u0005\b\u00a8\b\b\n\b\f\b\u00ab\t\b\u0003\b\u00ad\b\b\u0001\b\u0001"+
		"\b\u0003\b\u00b1\b\b\u0001\b\u0001\b\u0001\t\u0001\t\u0001\t\u0001\t\u0001"+
		"\n\u0001\n\u0001\n\u0005\n\u00bc\b\n\n\n\f\n\u00bf\t\n\u0001\u000b\u0001"+
		"\u000b\u0001\u000b\u0001\u000b\u0001\f\u0001\f\u0001\r\u0001\r\u0001\r"+
		"\u0001\r\u0001\r\u0001\r\u0001\r\u0001\r\u0001\r\u0001\r\u0001\r\u0001"+
		"\r\u0001\r\u0003\r\u00d4\b\r\u0001\u000e\u0001\u000e\u0001\u000e\u0001"+
		"\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001"+
		"\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001"+
		"\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001"+
		"\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001"+
		"\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001"+
		"\u000e\u0001\u000e\u0001\u000e\u0003\u000e\u00fa\b\u000e\u0001\u000f\u0001"+
		"\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0003\u000f\u0101\b\u000f\u0001"+
		"\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001"+
		"\u0010\u0001\u0010\u0003\u0010\u010b\b\u0010\u0001\u0011\u0001\u0011\u0001"+
		"\u0011\u0001\u0011\u0001\u0011\u0001\u0011\u0001\u0011\u0001\u0011\u0001"+
		"\u0011\u0005\u0011\u0116\b\u0011\n\u0011\f\u0011\u0119\t\u0011\u0001\u0011"+
		"\u0001\u0011\u0003\u0011\u011d\b\u0011\u0001\u0011\u0001\u0011\u0001\u0012"+
		"\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0013"+
		"\u0001\u0013\u0001\u0013\u0001\u0013\u0001\u0013\u0001\u0014\u0001\u0014"+
		"\u0001\u0014\u0001\u0014\u0001\u0014\u0001\u0014\u0001\u0014\u0001\u0014"+
		"\u0003\u0014\u0134\b\u0014\u0001\u0014\u0001\u0014\u0001\u0014\u0001\u0014"+
		"\u0001\u0014\u0001\u0014\u0001\u0014\u0001\u0014\u0001\u0014\u0001\u0014"+
		"\u0001\u0014\u0001\u0014\u0003\u0014\u0142\b\u0014\u0001\u0015\u0001\u0015"+
		"\u0001\u0016\u0001\u0016\u0001\u0016\u0001\u0017\u0001\u0017\u0001\u0017"+
		"\u0001\u0017\u0001\u0018\u0001\u0018\u0003\u0018\u014f\b\u0018\u0001\u0018"+
		"\u0003\u0018\u0152\b\u0018\u0001\u0018\u0001\u0018\u0001\u0018\u0001\u0018"+
		"\u0003\u0018\u0158\b\u0018\u0001\u0018\u0001\u0018\u0001\u0019\u0005\u0019"+
		"\u015d\b\u0019\n\u0019\f\u0019\u0160\t\u0019\u0001\u001a\u0001\u001a\u0001"+
		"\u001a\u0001\u001a\u0003\u001a\u0166\b\u001a\u0001\u001a\u0001\u001a\u0001"+
		"\u001a\u0001\u001a\u0005\u001a\u016c\b\u001a\n\u001a\f\u001a\u016f\t\u001a"+
		"\u0001\u001a\u0001\u001a\u0001\u001a\u0001\u001a\u0001\u001a\u0001\u001a"+
		"\u0001\u001a\u0001\u001a\u0001\u001a\u0005\u001a\u017a\b\u001a\n\u001a"+
		"\f\u001a\u017d\t\u001a\u0003\u001a\u017f\b\u001a\u0001\u001a\u0001\u001a"+
		"\u0001\u001a\u0001\u001a\u0003\u001a\u0185\b\u001a\u0001\u001b\u0001\u001b"+
		"\u0001\u001b\u0001\u001b\u0001\u001b\u0003\u001b\u018c\b\u001b\u0001\u001b"+
		"\u0001\u001b\u0001\u001b\u0001\u001b\u0001\u001b\u0005\u001b\u0193\b\u001b"+
		"\n\u001b\f\u001b\u0196\t\u001b\u0001\u001b\u0001\u001b\u0003\u001b\u019a"+
		"\b\u001b\u0001\u001b\u0003\u001b\u019d\b\u001b\u0001\u001b\u0001\u001b"+
		"\u0001\u001b\u0001\u001b\u0001\u001c\u0001\u001c\u0001\u001c\u0001\u001c"+
		"\u0005\u001c\u01a7\b\u001c\n\u001c\f\u001c\u01aa\t\u001c\u0001\u001c\u0003"+
		"\u001c\u01ad\b\u001c\u0001\u001d\u0001\u001d\u0001\u001d\u0001\u001d\u0001"+
		"\u001d\u0001\u001d\u0001\u001d\u0001\u001d\u0001\u001d\u0001\u001d\u0001"+
		"\u001d\u0001\u001d\u0003\u001d\u01bb\b\u001d\u0001\u001e\u0001\u001e\u0001"+
		"\u001f\u0001\u001f\u0001 \u0001 \u0001!\u0001!\u0001\"\u0001\"\u0001#"+
		"\u0001#\u0001$\u0001$\u0001%\u0001%\u0001&\u0001&\u0001\'\u0001\'\u0001"+
		"(\u0001(\u0001)\u0001)\u0001*\u0001*\u0001*\u0005*\u01d8\b*\n*\f*\u01db"+
		"\t*\u0001+\u0001+\u0001+\u0005+\u01e0\b+\n+\f+\u01e3\t+\u0001,\u0001,"+
		"\u0001,\u0003,\u01e8\b,\u0001,\u0001,\u0001,\u0001,\u0005,\u01ee\b,\n"+
		",\f,\u01f1\t,\u0003,\u01f3\b,\u0001,\u0001,\u0001,\u0001,\u0001-\u0001"+
		"-\u0001-\u0001-\u0003-\u01fd\b-\u0001-\u0001-\u0001-\u0001-\u0001-\u0000"+
		"\u0001\u0006.\u0000\u0002\u0004\u0006\b\n\f\u000e\u0010\u0012\u0014\u0016"+
		"\u0018\u001a\u001c\u001e \"$&(*,.02468:<>@BDFHJLNPRTVXZ\u0000\u000f\u0001"+
		"\u0000\u0001\u0005\u0001\u0000ad\u0001\u0000\r\u001c\u0001\u000006\u0001"+
		"\u000078\u0002\u0000\u001f\u001f//\u0001\u0000=>\u0001\u0000?D\u0001\u0000"+
		"EK\u0001\u0000LP\u0002\u0000\u0001\u0005QT\u0003\u0000  FFUU\u0003\u0000"+
		"\u0002\u0005QRXY\u0001\u0000Z[\u0001\u0000]^\u0225\u0000_\u0001\u0000"+
		"\u0000\u0000\u0002k\u0001\u0000\u0000\u0000\u0004m\u0001\u0000\u0000\u0000"+
		"\u0006v\u0001\u0000\u0000\u0000\b\u008f\u0001\u0000\u0000\u0000\n\u0091"+
		"\u0001\u0000\u0000\u0000\f\u0093\u0001\u0000\u0000\u0000\u000e\u0095\u0001"+
		"\u0000\u0000\u0000\u0010\u00a3\u0001\u0000\u0000\u0000\u0012\u00b4\u0001"+
		"\u0000\u0000\u0000\u0014\u00b8\u0001\u0000\u0000\u0000\u0016\u00c0\u0001"+
		"\u0000\u0000\u0000\u0018\u00c4\u0001\u0000\u0000\u0000\u001a\u00d3\u0001"+
		"\u0000\u0000\u0000\u001c\u00f9\u0001\u0000\u0000\u0000\u001e\u0100\u0001"+
		"\u0000\u0000\u0000 \u010a\u0001\u0000\u0000\u0000\"\u010c\u0001\u0000"+
		"\u0000\u0000$\u0120\u0001\u0000\u0000\u0000&\u0126\u0001\u0000\u0000\u0000"+
		"(\u0141\u0001\u0000\u0000\u0000*\u0143\u0001\u0000\u0000\u0000,\u0145"+
		"\u0001\u0000\u0000\u0000.\u0148\u0001\u0000\u0000\u00000\u014c\u0001\u0000"+
		"\u0000\u00002\u015e\u0001\u0000\u0000\u00004\u0184\u0001\u0000\u0000\u0000"+
		"6\u0186\u0001\u0000\u0000\u00008\u01a2\u0001\u0000\u0000\u0000:\u01ba"+
		"\u0001\u0000\u0000\u0000<\u01bc\u0001\u0000\u0000\u0000>\u01be\u0001\u0000"+
		"\u0000\u0000@\u01c0\u0001\u0000\u0000\u0000B\u01c2\u0001\u0000\u0000\u0000"+
		"D\u01c4\u0001\u0000\u0000\u0000F\u01c6\u0001\u0000\u0000\u0000H\u01c8"+
		"\u0001\u0000\u0000\u0000J\u01ca\u0001\u0000\u0000\u0000L\u01cc\u0001\u0000"+
		"\u0000\u0000N\u01ce\u0001\u0000\u0000\u0000P\u01d0\u0001\u0000\u0000\u0000"+
		"R\u01d2\u0001\u0000\u0000\u0000T\u01d4\u0001\u0000\u0000\u0000V\u01dc"+
		"\u0001\u0000\u0000\u0000X\u01e4\u0001\u0000\u0000\u0000Z\u01f8\u0001\u0000"+
		"\u0000\u0000\\^\u0003\u0002\u0001\u0000]\\\u0001\u0000\u0000\u0000^a\u0001"+
		"\u0000\u0000\u0000_]\u0001\u0000\u0000\u0000_`\u0001\u0000\u0000\u0000"+
		"`b\u0001\u0000\u0000\u0000a_\u0001\u0000\u0000\u0000bc\u0005\u0000\u0000"+
		"\u0001c\u0001\u0001\u0000\u0000\u0000dl\u0003\u0004\u0002\u0000el\u0003"+
		" \u0010\u0000fl\u00036\u001b\u0000gl\u00038\u001c\u0000hl\u00034\u001a"+
		"\u0000il\u0003\u0016\u000b\u0000jl\u0003Z-\u0000kd\u0001\u0000\u0000\u0000"+
		"ke\u0001\u0000\u0000\u0000kf\u0001\u0000\u0000\u0000kg\u0001\u0000\u0000"+
		"\u0000kh\u0001\u0000\u0000\u0000ki\u0001\u0000\u0000\u0000kj\u0001\u0000"+
		"\u0000\u0000l\u0003\u0001\u0000\u0000\u0000mn\u0003\f\u0006\u0000no\u0007"+
		"\u0000\u0000\u0000op\u0003\u0006\u0003\u0000p\u0005\u0001\u0000\u0000"+
		"\u0000qr\u0006\u0003\uffff\uffff\u0000rw\u0003\b\u0004\u0000st\u0003F"+
		"#\u0000tu\u0003\u0006\u0003\u0001uw\u0001\u0000\u0000\u0000vq\u0001\u0000"+
		"\u0000\u0000vs\u0001\u0000\u0000\u0000w\u0082\u0001\u0000\u0000\u0000"+
		"xy\n\u0003\u0000\u0000yz\u0003:\u001d\u0000z{\u0003\u0006\u0003\u0004"+
		"{\u0081\u0001\u0000\u0000\u0000|}\n\u0002\u0000\u0000}~\u0003:\u001d\u0000"+
		"~\u007f\u0003\u0006\u0003\u0002\u007f\u0081\u0001\u0000\u0000\u0000\u0080"+
		"x\u0001\u0000\u0000\u0000\u0080|\u0001\u0000\u0000\u0000\u0081\u0084\u0001"+
		"\u0000\u0000\u0000\u0082\u0080\u0001\u0000\u0000\u0000\u0082\u0083\u0001"+
		"\u0000\u0000\u0000\u0083\u0007\u0001\u0000\u0000\u0000\u0084\u0082\u0001"+
		"\u0000\u0000\u0000\u0085\u0090\u0003\n\u0005\u0000\u0086\u0090\u0003\f"+
		"\u0006\u0000\u0087\u0090\u0003\u000e\u0007\u0000\u0088\u0090\u0003\u001e"+
		"\u000f\u0000\u0089\u0090\u0003\u0010\b\u0000\u008a\u0090\u0003X,\u0000"+
		"\u008b\u008c\u0005\u0006\u0000\u0000\u008c\u008d\u0003\u0006\u0003\u0000"+
		"\u008d\u008e\u0005\u0007\u0000\u0000\u008e\u0090\u0001\u0000\u0000\u0000"+
		"\u008f\u0085\u0001\u0000\u0000\u0000\u008f\u0086\u0001\u0000\u0000\u0000"+
		"\u008f\u0087\u0001\u0000\u0000\u0000\u008f\u0088\u0001\u0000\u0000\u0000"+
		"\u008f\u0089\u0001\u0000\u0000\u0000\u008f\u008a\u0001\u0000\u0000\u0000"+
		"\u008f\u008b\u0001\u0000\u0000\u0000\u0090\t\u0001\u0000\u0000\u0000\u0091"+
		"\u0092\u0007\u0001\u0000\u0000\u0092\u000b\u0001\u0000\u0000\u0000\u0093"+
		"\u0094\u0005`\u0000\u0000\u0094\r\u0001\u0000\u0000\u0000\u0095\u0096"+
		"\u0005`\u0000\u0000\u0096\u009f\u0005\u0006\u0000\u0000\u0097\u009c\u0003"+
		"\u0006\u0003\u0000\u0098\u0099\u0005\b\u0000\u0000\u0099\u009b\u0003\u0006"+
		"\u0003\u0000\u009a\u0098\u0001\u0000\u0000\u0000\u009b\u009e\u0001\u0000"+
		"\u0000\u0000\u009c\u009a\u0001\u0000\u0000\u0000\u009c\u009d\u0001\u0000"+
		"\u0000\u0000\u009d\u00a0\u0001\u0000\u0000\u0000\u009e\u009c\u0001\u0000"+
		"\u0000\u0000\u009f\u0097\u0001\u0000\u0000\u0000\u009f\u00a0\u0001\u0000"+
		"\u0000\u0000\u00a0\u00a1\u0001\u0000\u0000\u0000\u00a1\u00a2\u0005\u0007"+
		"\u0000\u0000\u00a2\u000f\u0001\u0000\u0000\u0000\u00a3\u00ac\u0005\t\u0000"+
		"\u0000\u00a4\u00a9\u0003\u001a\r\u0000\u00a5\u00a6\u0005\b\u0000\u0000"+
		"\u00a6\u00a8\u0003\u001a\r\u0000\u00a7\u00a5\u0001\u0000\u0000\u0000\u00a8"+
		"\u00ab\u0001\u0000\u0000\u0000\u00a9\u00a7\u0001\u0000\u0000\u0000\u00a9"+
		"\u00aa\u0001\u0000\u0000\u0000\u00aa\u00ad\u0001\u0000\u0000\u0000\u00ab"+
		"\u00a9\u0001\u0000\u0000\u0000\u00ac\u00a4\u0001\u0000\u0000\u0000\u00ac"+
		"\u00ad\u0001\u0000\u0000\u0000\u00ad\u00b0\u0001\u0000\u0000\u0000\u00ae"+
		"\u00af\u0005\b\u0000\u0000\u00af\u00b1\u0003\u0012\t\u0000\u00b0\u00ae"+
		"\u0001\u0000\u0000\u0000\u00b0\u00b1\u0001\u0000\u0000\u0000\u00b1\u00b2"+
		"\u0001\u0000\u0000\u0000\u00b2\u00b3\u0005\n\u0000\u0000\u00b3\u0011\u0001"+
		"\u0000\u0000\u0000\u00b4\u00b5\u0005\u000b\u0000\u0000\u00b5\u00b6\u0005"+
		"\u0001\u0000\u0000\u00b6\u00b7\u0003\u0006\u0003\u0000\u00b7\u0013\u0001"+
		"\u0000\u0000\u0000\u00b8\u00bd\u0003\u0018\f\u0000\u00b9\u00ba\u0005\b"+
		"\u0000\u0000\u00ba\u00bc\u0003\u0018\f\u0000\u00bb\u00b9\u0001\u0000\u0000"+
		"\u0000\u00bc\u00bf\u0001\u0000\u0000\u0000\u00bd\u00bb\u0001\u0000\u0000"+
		"\u0000\u00bd\u00be\u0001\u0000\u0000\u0000\u00be\u0015\u0001\u0000\u0000"+
		"\u0000\u00bf\u00bd\u0001\u0000\u0000\u0000\u00c0\u00c1\u0005\f\u0000\u0000"+
		"\u00c1\u00c2\u0005`\u0000\u0000\u00c2\u00c3\u0005\f\u0000\u0000\u00c3"+
		"\u0017\u0001\u0000\u0000\u0000\u00c4\u00c5\u0007\u0002\u0000\u0000\u00c5"+
		"\u0019\u0001\u0000\u0000\u0000\u00c6\u00c7\u0005`\u0000\u0000\u00c7\u00c8"+
		"\u0005\u0001\u0000\u0000\u00c8\u00d4\u0003\u0006\u0003\u0000\u00c9\u00ca"+
		"\u0005\u001d\u0000\u0000\u00ca\u00cb\u0003\u0006\u0003\u0000\u00cb\u00cc"+
		"\u0005\u001e\u0000\u0000\u00cc\u00cd\u0005\u0001\u0000\u0000\u00cd\u00ce"+
		"\u0003\u0006\u0003\u0000\u00ce\u00d4\u0001\u0000\u0000\u0000\u00cf\u00d4"+
		"\u0003\u0006\u0003\u0000\u00d0\u00d1\u0005`\u0000\u0000\u00d1\u00d2\u0005"+
		"\u001f\u0000\u0000\u00d2\u00d4\u0003X,\u0000\u00d3\u00c6\u0001\u0000\u0000"+
		"\u0000\u00d3\u00c9\u0001\u0000\u0000\u0000\u00d3\u00cf\u0001\u0000\u0000"+
		"\u0000\u00d3\u00d0\u0001\u0000\u0000\u0000\u00d4\u001b\u0001\u0000\u0000"+
		"\u0000\u00d5\u00d6\u0003\u0006\u0003\u0000\u00d6\u00d7\u0003@ \u0000\u00d7"+
		"\u00d8\u0003\u0006\u0003\u0000\u00d8\u00fa\u0001\u0000\u0000\u0000\u00d9"+
		"\u00da\u0003\u0006\u0003\u0000\u00da\u00db\u0003B!\u0000\u00db\u00dc\u0003"+
		"\u0006\u0003\u0000\u00dc\u00fa\u0001\u0000\u0000\u0000\u00dd\u00de\u0003"+
		"\u0006\u0003\u0000\u00de\u00df\u0003>\u001f\u0000\u00df\u00e0\u0003\u0006"+
		"\u0003\u0000\u00e0\u00fa\u0001\u0000\u0000\u0000\u00e1\u00e2\u0003\u0006"+
		"\u0003\u0000\u00e2\u00e3\u0003<\u001e\u0000\u00e3\u00e4\u0003\u0006\u0003"+
		"\u0000\u00e4\u00fa\u0001\u0000\u0000\u0000\u00e5\u00e6\u0003\u0006\u0003"+
		"\u0000\u00e6\u00e7\u0003H$\u0000\u00e7\u00e8\u0003\u0006\u0003\u0000\u00e8"+
		"\u00fa\u0001\u0000\u0000\u0000\u00e9\u00ea\u0003\u0006\u0003\u0000\u00ea"+
		"\u00eb\u0003L&\u0000\u00eb\u00ec\u0003\u0006\u0003\u0000\u00ec\u00fa\u0001"+
		"\u0000\u0000\u0000\u00ed\u00ee\u0003\u0006\u0003\u0000\u00ee\u00ef\u0003"+
		"P(\u0000\u00ef\u00f0\u0003\u0006\u0003\u0000\u00f0\u00fa\u0001\u0000\u0000"+
		"\u0000\u00f1\u00f2\u0003\u0006\u0003\u0000\u00f2\u00f3\u0003R)\u0000\u00f3"+
		"\u00f4\u0003\u0006\u0003\u0000\u00f4\u00fa\u0001\u0000\u0000\u0000\u00f5"+
		"\u00f6\u0003\u0006\u0003\u0000\u00f6\u00f7\u0003N\'\u0000\u00f7\u00f8"+
		"\u0003\u0006\u0003\u0000\u00f8\u00fa\u0001\u0000\u0000\u0000\u00f9\u00d5"+
		"\u0001\u0000\u0000\u0000\u00f9\u00d9\u0001\u0000\u0000\u0000\u00f9\u00dd"+
		"\u0001\u0000\u0000\u0000\u00f9\u00e1\u0001\u0000\u0000\u0000\u00f9\u00e5"+
		"\u0001\u0000\u0000\u0000\u00f9\u00e9\u0001\u0000\u0000\u0000\u00f9\u00ed"+
		"\u0001\u0000\u0000\u0000\u00f9\u00f1\u0001\u0000\u0000\u0000\u00f9\u00f5"+
		"\u0001\u0000\u0000\u0000\u00fa\u001d\u0001\u0000\u0000\u0000\u00fb\u00fc"+
		"\u0003F#\u0000\u00fc\u00fd\u0003\u0006\u0003\u0000\u00fd\u0101\u0001\u0000"+
		"\u0000\u0000\u00fe\u00ff\u0005 \u0000\u0000\u00ff\u0101\u0003\u0006\u0003"+
		"\u0000\u0100\u00fb\u0001\u0000\u0000\u0000\u0100\u00fe\u0001\u0000\u0000"+
		"\u0000\u0101\u001f\u0001\u0000\u0000\u0000\u0102\u010b\u0003\"\u0011\u0000"+
		"\u0103\u010b\u0003$\u0012\u0000\u0104\u010b\u0003&\u0013\u0000\u0105\u010b"+
		"\u0003(\u0014\u0000\u0106\u010b\u0003*\u0015\u0000\u0107\u010b\u0003,"+
		"\u0016\u0000\u0108\u010b\u0003.\u0017\u0000\u0109\u010b\u00030\u0018\u0000"+
		"\u010a\u0102\u0001\u0000\u0000\u0000\u010a\u0103\u0001\u0000\u0000\u0000"+
		"\u010a\u0104\u0001\u0000\u0000\u0000\u010a\u0105\u0001\u0000\u0000\u0000"+
		"\u010a\u0106\u0001\u0000\u0000\u0000\u010a\u0107\u0001\u0000\u0000\u0000"+
		"\u010a\u0108\u0001\u0000\u0000\u0000\u010a\u0109\u0001\u0000\u0000\u0000"+
		"\u010b!\u0001\u0000\u0000\u0000\u010c\u010d\u0005!\u0000\u0000\u010d\u010e"+
		"\u0003\u0006\u0003\u0000\u010e\u010f\u0005\"\u0000\u0000\u010f\u0117\u0003"+
		"2\u0019\u0000\u0110\u0111\u0005#\u0000\u0000\u0111\u0112\u0003\u0006\u0003"+
		"\u0000\u0112\u0113\u0005\"\u0000\u0000\u0113\u0114\u00032\u0019\u0000"+
		"\u0114\u0116\u0001\u0000\u0000\u0000\u0115\u0110\u0001\u0000\u0000\u0000"+
		"\u0116\u0119\u0001\u0000\u0000\u0000\u0117\u0115\u0001\u0000\u0000\u0000"+
		"\u0117\u0118\u0001\u0000\u0000\u0000\u0118\u011c\u0001\u0000\u0000\u0000"+
		"\u0119\u0117\u0001\u0000\u0000\u0000\u011a\u011b\u0005$\u0000\u0000\u011b"+
		"\u011d\u00032\u0019\u0000\u011c\u011a\u0001\u0000\u0000\u0000\u011c\u011d"+
		"\u0001\u0000\u0000\u0000\u011d\u011e\u0001\u0000\u0000\u0000\u011e\u011f"+
		"\u0005%\u0000\u0000\u011f#\u0001\u0000\u0000\u0000\u0120\u0121\u0005&"+
		"\u0000\u0000\u0121\u0122\u0003\u0006\u0003\u0000\u0122\u0123\u0005\'\u0000"+
		"\u0000\u0123\u0124\u00032\u0019\u0000\u0124\u0125\u0005%\u0000\u0000\u0125"+
		"%\u0001\u0000\u0000\u0000\u0126\u0127\u0005(\u0000\u0000\u0127\u0128\u0003"+
		"2\u0019\u0000\u0128\u0129\u0005)\u0000\u0000\u0129\u012a\u0003\u0006\u0003"+
		"\u0000\u012a\'\u0001\u0000\u0000\u0000\u012b\u012c\u0005*\u0000\u0000"+
		"\u012c\u012d\u0005`\u0000\u0000\u012d\u012e\u0005\u0001\u0000\u0000\u012e"+
		"\u012f\u0003\u0006\u0003\u0000\u012f\u0130\u0005\b\u0000\u0000\u0130\u0133"+
		"\u0003\u0006\u0003\u0000\u0131\u0132\u0005\b\u0000\u0000\u0132\u0134\u0003"+
		"\u0006\u0003\u0000\u0133\u0131\u0001\u0000\u0000\u0000\u0133\u0134\u0001"+
		"\u0000\u0000\u0000\u0134\u0135\u0001\u0000\u0000\u0000\u0135\u0136\u0005"+
		"\'\u0000\u0000\u0136\u0137\u00032\u0019\u0000\u0137\u0138\u0005%\u0000"+
		"\u0000\u0138\u0142\u0001\u0000\u0000\u0000\u0139\u013a\u0005*\u0000\u0000"+
		"\u013a\u013b\u0003T*\u0000\u013b\u013c\u0005+\u0000\u0000\u013c\u013d"+
		"\u0003V+\u0000\u013d\u013e\u0005\'\u0000\u0000\u013e\u013f\u00032\u0019"+
		"\u0000\u013f\u0140\u0005%\u0000\u0000\u0140\u0142\u0001\u0000\u0000\u0000"+
		"\u0141\u012b\u0001\u0000\u0000\u0000\u0141\u0139\u0001\u0000\u0000\u0000"+
		"\u0142)\u0001\u0000\u0000\u0000\u0143\u0144\u0005,\u0000\u0000\u0144+"+
		"\u0001\u0000\u0000\u0000\u0145\u0146\u0005-\u0000\u0000\u0146\u0147\u0005"+
		"`\u0000\u0000\u0147-\u0001\u0000\u0000\u0000\u0148\u0149\u0005.\u0000"+
		"\u0000\u0149\u014a\u0005/\u0000\u0000\u014a\u014b\u0007\u0003\u0000\u0000"+
		"\u014b/\u0001\u0000\u0000\u0000\u014c\u014e\u0007\u0004\u0000\u0000\u014d"+
		"\u014f\u0007\u0005\u0000\u0000\u014e\u014d\u0001\u0000\u0000\u0000\u014e"+
		"\u014f\u0001\u0000\u0000\u0000\u014f\u0151\u0001\u0000\u0000\u0000\u0150"+
		"\u0152\u0005`\u0000\u0000\u0151\u0150\u0001\u0000\u0000\u0000\u0151\u0152"+
		"\u0001\u0000\u0000\u0000\u0152\u0153\u0001\u0000\u0000\u0000\u0153\u0154"+
		"\u0005\u0006\u0000\u0000\u0154\u0157\u0003\u0006\u0003\u0000\u0155\u0156"+
		"\u0005\b\u0000\u0000\u0156\u0158\u0003\u0006\u0003\u0000\u0157\u0155\u0001"+
		"\u0000\u0000\u0000\u0157\u0158\u0001\u0000\u0000\u0000\u0158\u0159\u0001"+
		"\u0000\u0000\u0000\u0159\u015a\u0005\u0007\u0000\u0000\u015a1\u0001\u0000"+
		"\u0000\u0000\u015b\u015d\u0003\u0002\u0001\u0000\u015c\u015b\u0001\u0000"+
		"\u0000\u0000\u015d\u0160\u0001\u0000\u0000\u0000\u015e\u015c\u0001\u0000"+
		"\u0000\u0000\u015e\u015f\u0001\u0000\u0000\u0000\u015f3\u0001\u0000\u0000"+
		"\u0000\u0160\u015e\u0001\u0000\u0000\u0000\u0161\u0162\u00059\u0000\u0000"+
		"\u0162\u0165\u0005`\u0000\u0000\u0163\u0164\u0005\u0001\u0000\u0000\u0164"+
		"\u0166\u0003\u0006\u0003\u0000\u0165\u0163\u0001\u0000\u0000\u0000\u0165"+
		"\u0166\u0001\u0000\u0000\u0000\u0166\u0185\u0001\u0000\u0000\u0000\u0167"+
		"\u0168\u00059\u0000\u0000\u0168\u016d\u0005`\u0000\u0000\u0169\u016a\u0005"+
		"\b\u0000\u0000\u016a\u016c\u0005`\u0000\u0000\u016b\u0169\u0001\u0000"+
		"\u0000\u0000\u016c\u016f\u0001\u0000\u0000\u0000\u016d\u016b\u0001\u0000"+
		"\u0000\u0000\u016d\u016e\u0001\u0000\u0000\u0000\u016e\u0170\u0001\u0000"+
		"\u0000\u0000\u016f\u016d\u0001\u0000\u0000\u0000\u0170\u0171\u0005\u0001"+
		"\u0000\u0000\u0171\u0185\u0003V+\u0000\u0172\u0173\u00059\u0000\u0000"+
		"\u0173\u0174\u0005:\u0000\u0000\u0174\u0175\u0005`\u0000\u0000\u0175\u017e"+
		"\u0005\u0006\u0000\u0000\u0176\u017b\u0005`\u0000\u0000\u0177\u0178\u0005"+
		"\b\u0000\u0000\u0178\u017a\u0005`\u0000\u0000\u0179\u0177\u0001\u0000"+
		"\u0000\u0000\u017a\u017d\u0001\u0000\u0000\u0000\u017b\u0179\u0001\u0000"+
		"\u0000\u0000\u017b\u017c\u0001\u0000\u0000\u0000\u017c\u017f\u0001\u0000"+
		"\u0000\u0000\u017d\u017b\u0001\u0000\u0000\u0000\u017e\u0176\u0001\u0000"+
		"\u0000\u0000\u017e\u017f\u0001\u0000\u0000\u0000\u017f\u0180\u0001\u0000"+
		"\u0000\u0000\u0180\u0181\u0005\u0007\u0000\u0000\u0181\u0182\u00032\u0019"+
		"\u0000\u0182\u0183\u0005%\u0000\u0000\u0183\u0185\u0001\u0000\u0000\u0000"+
		"\u0184\u0161\u0001\u0000\u0000\u0000\u0184\u0167\u0001\u0000\u0000\u0000"+
		"\u0184\u0172\u0001\u0000\u0000\u0000\u01855\u0001\u0000\u0000\u0000\u0186"+
		"\u018b\u0005:\u0000\u0000\u0187\u0188\u0005`\u0000\u0000\u0188\u018c\u0005"+
		"/\u0000\u0000\u0189\u018a\u0005`\u0000\u0000\u018a\u018c\u0005\u001f\u0000"+
		"\u0000\u018b\u0187\u0001\u0000\u0000\u0000\u018b\u0189\u0001\u0000\u0000"+
		"\u0000\u018b\u018c\u0001\u0000\u0000\u0000\u018c\u018d\u0001\u0000\u0000"+
		"\u0000\u018d\u018e\u0005`\u0000\u0000\u018e\u019c\u0005\u0006\u0000\u0000"+
		"\u018f\u0194\u0005`\u0000\u0000\u0190\u0191\u0005\b\u0000\u0000\u0191"+
		"\u0193\u0005`\u0000\u0000\u0192\u0190\u0001\u0000\u0000\u0000\u0193\u0196"+
		"\u0001\u0000\u0000\u0000\u0194\u0192\u0001\u0000\u0000\u0000\u0194\u0195"+
		"\u0001\u0000\u0000\u0000\u0195\u0199\u0001\u0000\u0000\u0000\u0196\u0194"+
		"\u0001\u0000\u0000\u0000\u0197\u0198\u0005\b\u0000\u0000\u0198\u019a\u0003"+
		"J%\u0000\u0199\u0197\u0001\u0000\u0000\u0000\u0199\u019a\u0001\u0000\u0000"+
		"\u0000\u019a\u019d\u0001\u0000\u0000\u0000\u019b\u019d\u0003J%\u0000\u019c"+
		"\u018f\u0001\u0000\u0000\u0000\u019c\u019b\u0001\u0000\u0000\u0000\u019c"+
		"\u019d\u0001\u0000\u0000\u0000\u019d\u019e\u0001\u0000\u0000\u0000\u019e"+
		"\u019f\u0005\u0007\u0000\u0000\u019f\u01a0\u00032\u0019\u0000\u01a0\u01a1"+
		"\u0005%\u0000\u0000\u01a17\u0001\u0000\u0000\u0000\u01a2\u01ac\u0005;"+
		"\u0000\u0000\u01a3\u01a8\u0003\u0006\u0003\u0000\u01a4\u01a5\u0005\b\u0000"+
		"\u0000\u01a5\u01a7\u0003\u0006\u0003\u0000\u01a6\u01a4\u0001\u0000\u0000"+
		"\u0000\u01a7\u01aa\u0001\u0000\u0000\u0000\u01a8\u01a6\u0001\u0000\u0000"+
		"\u0000\u01a8\u01a9\u0001\u0000\u0000\u0000\u01a9\u01ad\u0001\u0000\u0000"+
		"\u0000\u01aa\u01a8\u0001\u0000\u0000\u0000\u01ab\u01ad\u0003\u000e\u0007"+
		"\u0000\u01ac\u01a3\u0001\u0000\u0000\u0000\u01ac\u01ab\u0001\u0000\u0000"+
		"\u0000\u01ac\u01ad\u0001\u0000\u0000\u0000\u01ad9\u0001\u0000\u0000\u0000"+
		"\u01ae\u01bb\u0003<\u001e\u0000\u01af\u01bb\u0003>\u001f\u0000\u01b0\u01bb"+
		"\u0003@ \u0000\u01b1\u01bb\u0003B!\u0000\u01b2\u01bb\u0003D\"\u0000\u01b3"+
		"\u01bb\u0003F#\u0000\u01b4\u01bb\u0003H$\u0000\u01b5\u01bb\u0003L&\u0000"+
		"\u01b6\u01bb\u0003N\'\u0000\u01b7\u01bb\u0003P(\u0000\u01b8\u01bb\u0003"+
		"R)\u0000\u01b9\u01bb\u0005<\u0000\u0000\u01ba\u01ae\u0001\u0000\u0000"+
		"\u0000\u01ba\u01af\u0001\u0000\u0000\u0000\u01ba\u01b0\u0001\u0000\u0000"+
		"\u0000\u01ba\u01b1\u0001\u0000\u0000\u0000\u01ba\u01b2\u0001\u0000\u0000"+
		"\u0000\u01ba\u01b3\u0001\u0000\u0000\u0000\u01ba\u01b4\u0001\u0000\u0000"+
		"\u0000\u01ba\u01b5\u0001\u0000\u0000\u0000\u01ba\u01b6\u0001\u0000\u0000"+
		"\u0000\u01ba\u01b7\u0001\u0000\u0000\u0000\u01ba\u01b8\u0001\u0000\u0000"+
		"\u0000\u01ba\u01b9\u0001\u0000\u0000\u0000\u01bb;\u0001\u0000\u0000\u0000"+
		"\u01bc\u01bd\u0007\u0006\u0000\u0000\u01bd=\u0001\u0000\u0000\u0000\u01be"+
		"\u01bf\u0007\u0007\u0000\u0000\u01bf?\u0001\u0000\u0000\u0000\u01c0\u01c1"+
		"\u0007\b\u0000\u0000\u01c1A\u0001\u0000\u0000\u0000\u01c2\u01c3\u0007"+
		"\t\u0000\u0000\u01c3C\u0001\u0000\u0000\u0000\u01c4\u01c5\u0007\n\u0000"+
		"\u0000\u01c5E\u0001\u0000\u0000\u0000\u01c6\u01c7\u0007\u000b\u0000\u0000"+
		"\u01c7G\u0001\u0000\u0000\u0000\u01c8\u01c9\u0005V\u0000\u0000\u01c9I"+
		"\u0001\u0000\u0000\u0000\u01ca\u01cb\u0005W\u0000\u0000\u01cbK\u0001\u0000"+
		"\u0000\u0000\u01cc\u01cd\u0007\f\u0000\u0000\u01cdM\u0001\u0000\u0000"+
		"\u0000\u01ce\u01cf\u0007\r\u0000\u0000\u01cfO\u0001\u0000\u0000\u0000"+
		"\u01d0\u01d1\u0005\\\u0000\u0000\u01d1Q\u0001\u0000\u0000\u0000\u01d2"+
		"\u01d3\u0007\u000e\u0000\u0000\u01d3S\u0001\u0000\u0000\u0000\u01d4\u01d9"+
		"\u0005`\u0000\u0000\u01d5\u01d6\u0005\b\u0000\u0000\u01d6\u01d8\u0005"+
		"`\u0000\u0000\u01d7\u01d5\u0001\u0000\u0000\u0000\u01d8\u01db\u0001\u0000"+
		"\u0000\u0000\u01d9\u01d7\u0001\u0000\u0000\u0000\u01d9\u01da\u0001\u0000"+
		"\u0000\u0000\u01daU\u0001\u0000\u0000\u0000\u01db\u01d9\u0001\u0000\u0000"+
		"\u0000\u01dc\u01e1\u0003\u0006\u0003\u0000\u01dd\u01de\u0005\b\u0000\u0000"+
		"\u01de\u01e0\u0003\u0006\u0003\u0000\u01df\u01dd\u0001\u0000\u0000\u0000"+
		"\u01e0\u01e3\u0001\u0000\u0000\u0000\u01e1\u01df\u0001\u0000\u0000\u0000"+
		"\u01e1\u01e2\u0001\u0000\u0000\u0000\u01e2W\u0001\u0000\u0000\u0000\u01e3"+
		"\u01e1\u0001\u0000\u0000\u0000\u01e4\u01e7\u0005:\u0000\u0000\u01e5\u01e6"+
		"\u0005`\u0000\u0000\u01e6\u01e8\u0005\u001f\u0000\u0000\u01e7\u01e5\u0001"+
		"\u0000\u0000\u0000\u01e7\u01e8\u0001\u0000\u0000\u0000\u01e8\u01e9\u0001"+
		"\u0000\u0000\u0000\u01e9\u01f2\u0005\u0006\u0000\u0000\u01ea\u01ef\u0005"+
		"`\u0000\u0000\u01eb\u01ec\u0005\b\u0000\u0000\u01ec\u01ee\u0005`\u0000"+
		"\u0000\u01ed\u01eb\u0001\u0000\u0000\u0000\u01ee\u01f1\u0001\u0000\u0000"+
		"\u0000\u01ef\u01ed\u0001\u0000\u0000\u0000\u01ef\u01f0\u0001\u0000\u0000"+
		"\u0000\u01f0\u01f3\u0001\u0000\u0000\u0000\u01f1\u01ef\u0001\u0000\u0000"+
		"\u0000\u01f2\u01ea\u0001\u0000\u0000\u0000\u01f2\u01f3\u0001\u0000\u0000"+
		"\u0000\u01f3\u01f4\u0001\u0000\u0000\u0000\u01f4\u01f5\u0005\u0007\u0000"+
		"\u0000\u01f5\u01f6\u00032\u0019\u0000\u01f6\u01f7\u0005%\u0000\u0000\u01f7"+
		"Y\u0001\u0000\u0000\u0000\u01f8\u01f9\u0005_\u0000\u0000\u01f9\u01fc\u0005"+
		"\u0006\u0000\u0000\u01fa\u01fd\u0005 \u0000\u0000\u01fb\u01fd\u0003\u0006"+
		"\u0003\u0000\u01fc\u01fa\u0001\u0000\u0000\u0000\u01fc\u01fb\u0001\u0000"+
		"\u0000\u0000\u01fd\u01fe\u0001\u0000\u0000\u0000\u01fe\u01ff\u0005\b\u0000"+
		"\u0000\u01ff\u0200\u0003\u0006\u0003\u0000\u0200\u0201\u0005\u0007\u0000"+
		"\u0000\u0201[\u0001\u0000\u0000\u0000*_kv\u0080\u0082\u008f\u009c\u009f"+
		"\u00a9\u00ac\u00b0\u00bd\u00d3\u00f9\u0100\u010a\u0117\u011c\u0133\u0141"+
		"\u014e\u0151\u0157\u015e\u0165\u016d\u017b\u017e\u0184\u018b\u0194\u0199"+
		"\u019c\u01a8\u01ac\u01ba\u01d9\u01e1\u01e7\u01ef\u01f2\u01fc";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}