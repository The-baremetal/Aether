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
		T__94=95, T__95=96, T__96=97, T__97=98, T__98=99, T__99=100, T__100=101, 
		T__101=102, T__102=103, T__103=104, T__104=105, T__105=106, T__106=107, 
		T__107=108, T__108=109, T__109=110, T__110=111, T__111=112, T__112=113, 
		T__113=114, T__114=115, T__115=116, IDENTIFIER=117, BOOL=118, NIL=119, 
		NUMBER=120, STRING=121, WS=122, COMMENT=123;
	public static final int
		RULE_program = 0, RULE_statement = 1, RULE_assignStatement = 2, RULE_expression = 3, 
		RULE_primaryExpression = 4, RULE_callChain = 5, RULE_literal = 6, RULE_variable = 7, 
		RULE_safeAccess = 8, RULE_functionCall = 9, RULE_tableConstructor = 10, 
		RULE_metatable = 11, RULE_metamethods = 12, RULE_labelStatement = 13, 
		RULE_matchArm = 14, RULE_pattern = 15, RULE_fieldPattern = 16, RULE_metamethod = 17, 
		RULE_field = 18, RULE_binaryOperation = 19, RULE_unaryOperation = 20, 
		RULE_controlFlowStatement = 21, RULE_ifStatement = 22, RULE_whileStatement = 23, 
		RULE_repeatStatement = 24, RULE_forStatement = 25, RULE_breakStatement = 26, 
		RULE_gotoStatement = 27, RULE_coroutineStatement = 28, RULE_protectedCallStatement = 29, 
		RULE_namedArgs = 30, RULE_block = 31, RULE_localDeclaration = 32, RULE_functionDeclaration = 33, 
		RULE_returnStatement = 34, RULE_operatorGroup = 35, RULE_logicalOp = 36, 
		RULE_comparisonOp = 37, RULE_arithOp = 38, RULE_bitwiseOp = 39, RULE_assignOp = 40, 
		RULE_unaryOp = 41, RULE_concatOp = 42, RULE_varargOp = 43, RULE_compoundAssignOp = 44, 
		RULE_incrOp = 45, RULE_coalesceOp = 46, RULE_shiftAssignOp = 47, RULE_nonNullAssertOp = 48, 
		RULE_identifierList = 49, RULE_expressionList = 50, RULE_functionExpression = 51, 
		RULE_selectStatement = 52, RULE_lambdaExpression = 53, RULE_typeAnnotation = 54, 
		RULE_typeSpec = 55, RULE_experimentalExpression = 56, RULE_safeNavigation = 57, 
		RULE_pipeOperator = 58, RULE_decoratorSyntax = 59;
	private static String[] makeRuleNames() {
		return new String[] {
			"program", "statement", "assignStatement", "expression", "primaryExpression", 
			"callChain", "literal", "variable", "safeAccess", "functionCall", "tableConstructor", 
			"metatable", "metamethods", "labelStatement", "matchArm", "pattern", 
			"fieldPattern", "metamethod", "field", "binaryOperation", "unaryOperation", 
			"controlFlowStatement", "ifStatement", "whileStatement", "repeatStatement", 
			"forStatement", "breakStatement", "gotoStatement", "coroutineStatement", 
			"protectedCallStatement", "namedArgs", "block", "localDeclaration", "functionDeclaration", 
			"returnStatement", "operatorGroup", "logicalOp", "comparisonOp", "arithOp", 
			"bitwiseOp", "assignOp", "unaryOp", "concatOp", "varargOp", "compoundAssignOp", 
			"incrOp", "coalesceOp", "shiftAssignOp", "nonNullAssertOp", "identifierList", 
			"expressionList", "functionExpression", "selectStatement", "lambdaExpression", 
			"typeAnnotation", "typeSpec", "experimentalExpression", "safeNavigation", 
			"pipeOperator", "decoratorSyntax"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'spawn'", "'match'", "'with'", "'end'", "'async'", "'await'", 
			"'!!'", "'('", "')'", "':'", "'.'", "'['", "']'", "'?.'", "'?['", "','", 
			"'{'", "'}'", "'__metatable'", "'='", "'::'", "'when'", "'=>'", "'|'", 
			"'__add'", "'__sub'", "'__mul'", "'__div'", "'__mod'", "'__pow'", "'__unm'", 
			"'__concat'", "'__len'", "'__eq'", "'__lt'", "'__le'", "'__tostring'", 
			"'__pairs'", "'__ipairs'", "'__call'", "'#'", "'if'", "'then'", "'elseif'", 
			"'else'", "'while'", "'do'", "'repeat'", "'until'", "'for'", "'in'", 
			"'break'", "'goto'", "'coroutine'", "'create'", "'resume'", "'yield'", 
			"'status'", "'running'", "'wrap'", "'isyieldable'", "'pcall'", "'xpcall'", 
			"'local'", "'function'", "'return'", "'and'", "'or'", "'=='", "'>='", 
			"'<='", "'~='", "'>'", "'<'", "'+'", "'-'", "'*'", "'/'", "'//'", "'%'", 
			"'^'", "'&'", "'~'", "'<<'", "'>>'", "'+='", "'-='", "'*='", "'/='", 
			"'//='", "'^='", "'&='", "'|='", "'not'", "'typeof'", "'..'", "'...'", 
			"'..='", "'??='", "'+_='", "'-_='", "'+_'", "'-_'", "'??'", "'<<='", 
			"'>>='", "'select'", "'number'", "'string'", "'boolean'", "'table'", 
			"'any'", "'[]'", "'table<'", "'|>'", "'@'", null, null, "'nil'"
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
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, null, "IDENTIFIER", "BOOL", 
			"NIL", "NUMBER", "STRING", "WS", "COMMENT"
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
			setState(123);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & -4578678679359455230L) != 0) || ((((_la - 64)) & ~0x3f) == 0 && ((1L << (_la - 64)) & 9015995347763207L) != 0)) {
				{
				{
				setState(120);
				statement();
				}
				}
				setState(125);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(126);
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
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
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
			setState(137);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case IDENTIFIER:
				enterOuterAlt(_localctx, 1);
				{
				setState(128);
				assignStatement();
				}
				break;
			case T__41:
			case T__45:
			case T__47:
			case T__49:
			case T__51:
			case T__52:
			case T__53:
			case T__61:
			case T__62:
				enterOuterAlt(_localctx, 2);
				{
				setState(129);
				controlFlowStatement();
				}
				break;
			case T__64:
				enterOuterAlt(_localctx, 3);
				{
				setState(130);
				functionDeclaration();
				}
				break;
			case T__65:
				enterOuterAlt(_localctx, 4);
				{
				setState(131);
				returnStatement();
				}
				break;
			case T__63:
				enterOuterAlt(_localctx, 5);
				{
				setState(132);
				localDeclaration();
				}
				break;
			case T__20:
				enterOuterAlt(_localctx, 6);
				{
				setState(133);
				labelStatement();
				}
				break;
			case T__106:
				enterOuterAlt(_localctx, 7);
				{
				setState(134);
				selectStatement();
				}
				break;
			case T__0:
				enterOuterAlt(_localctx, 8);
				{
				setState(135);
				match(T__0);
				setState(136);
				expression(0);
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
		public AssignOpContext assignOp() {
			return getRuleContext(AssignOpContext.class,0);
		}
		public IncrOpContext incrOp() {
			return getRuleContext(IncrOpContext.class,0);
		}
		public ShiftAssignOpContext shiftAssignOp() {
			return getRuleContext(ShiftAssignOpContext.class,0);
		}
		public CoalesceOpContext coalesceOp() {
			return getRuleContext(CoalesceOpContext.class,0);
		}
		public AssignStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_assignStatement; }
	}

	public final AssignStatementContext assignStatement() throws RecognitionException {
		AssignStatementContext _localctx = new AssignStatementContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_assignStatement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(139);
			variable(0);
			setState(144);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__19:
			case T__85:
			case T__86:
			case T__87:
			case T__88:
			case T__89:
			case T__90:
			case T__91:
			case T__92:
				{
				setState(140);
				assignOp();
				}
				break;
			case T__101:
			case T__102:
				{
				setState(141);
				incrOp();
				}
				break;
			case T__104:
			case T__105:
				{
				setState(142);
				shiftAssignOp();
				}
				break;
			case T__103:
				{
				setState(143);
				coalesceOp();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			setState(146);
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
		public List<OperatorGroupContext> operatorGroup() {
			return getRuleContexts(OperatorGroupContext.class);
		}
		public OperatorGroupContext operatorGroup(int i) {
			return getRuleContext(OperatorGroupContext.class,i);
		}
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public UnaryOpContext unaryOp() {
			return getRuleContext(UnaryOpContext.class,0);
		}
		public List<MatchArmContext> matchArm() {
			return getRuleContexts(MatchArmContext.class);
		}
		public MatchArmContext matchArm(int i) {
			return getRuleContext(MatchArmContext.class,i);
		}
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
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
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(177);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,5,_ctx) ) {
			case 1:
				{
				setState(149);
				primaryExpression();
				setState(155);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,3,_ctx);
				while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
					if ( _alt==1 ) {
						{
						{
						setState(150);
						operatorGroup();
						setState(151);
						expression(0);
						}
						} 
					}
					setState(157);
					_errHandler.sync(this);
					_alt = getInterpreter().adaptivePredict(_input,3,_ctx);
				}
				}
				break;
			case 2:
				{
				setState(158);
				unaryOp();
				setState(159);
				expression(5);
				}
				break;
			case 3:
				{
				setState(161);
				match(T__1);
				setState(162);
				expression(0);
				setState(163);
				match(T__2);
				setState(165); 
				_errHandler.sync(this);
				_la = _input.LA(1);
				do {
					{
					{
					setState(164);
					matchArm();
					}
					}
					setState(167); 
					_errHandler.sync(this);
					_la = _input.LA(1);
				} while ( _la==T__16 || ((((_la - 117)) & ~0x3f) == 0 && ((1L << (_la - 117)) & 31L) != 0) );
				setState(169);
				match(T__3);
				}
				break;
			case 4:
				{
				setState(171);
				match(T__4);
				setState(172);
				block();
				setState(173);
				match(T__3);
				}
				break;
			case 5:
				{
				setState(175);
				match(T__5);
				setState(176);
				expression(2);
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(191);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,7,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(189);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,6,_ctx) ) {
					case 1:
						{
						_localctx = new ExpressionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(179);
						if (!(precpred(_ctx, 7))) throw new FailedPredicateException(this, "precpred(_ctx, 7)");
						setState(180);
						operatorGroup();
						setState(181);
						expression(8);
						}
						break;
					case 2:
						{
						_localctx = new ExpressionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(183);
						if (!(precpred(_ctx, 6))) throw new FailedPredicateException(this, "precpred(_ctx, 6)");
						setState(184);
						operatorGroup();
						setState(185);
						expression(6);
						}
						break;
					case 3:
						{
						_localctx = new ExpressionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(187);
						if (!(precpred(_ctx, 1))) throw new FailedPredicateException(this, "precpred(_ctx, 1)");
						setState(188);
						match(T__6);
						}
						break;
					}
					} 
				}
				setState(193);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,7,_ctx);
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
		public LambdaExpressionContext lambdaExpression() {
			return getRuleContext(LambdaExpressionContext.class,0);
		}
		public List<CallChainContext> callChain() {
			return getRuleContexts(CallChainContext.class);
		}
		public CallChainContext callChain(int i) {
			return getRuleContext(CallChainContext.class,i);
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
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(205);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,8,_ctx) ) {
			case 1:
				{
				setState(194);
				literal();
				}
				break;
			case 2:
				{
				setState(195);
				variable(0);
				}
				break;
			case 3:
				{
				setState(196);
				functionCall();
				}
				break;
			case 4:
				{
				setState(197);
				unaryOperation();
				}
				break;
			case 5:
				{
				setState(198);
				tableConstructor();
				}
				break;
			case 6:
				{
				setState(199);
				functionExpression();
				}
				break;
			case 7:
				{
				setState(200);
				match(T__7);
				setState(201);
				expression(0);
				setState(202);
				match(T__8);
				}
				break;
			case 8:
				{
				setState(204);
				lambdaExpression();
				}
				break;
			}
			setState(210);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,9,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(207);
					callChain();
					}
					} 
				}
				setState(212);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,9,_ctx);
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
	public static class CallChainContext extends ParserRuleContext {
		public CallChainContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_callChain; }
	 
		public CallChainContext() { }
		public void copyFrom(CallChainContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class PropertyChainContext extends CallChainContext {
		public TerminalNode IDENTIFIER() { return getToken(lua_grammar_antlr4Parser.IDENTIFIER, 0); }
		public PropertyChainContext(CallChainContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class IndexChainContext extends CallChainContext {
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public IndexChainContext(CallChainContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class MethodChainContext extends CallChainContext {
		public TerminalNode IDENTIFIER() { return getToken(lua_grammar_antlr4Parser.IDENTIFIER, 0); }
		public ExpressionListContext expressionList() {
			return getRuleContext(ExpressionListContext.class,0);
		}
		public MethodChainContext(CallChainContext ctx) { copyFrom(ctx); }
	}

	public final CallChainContext callChain() throws RecognitionException {
		CallChainContext _localctx = new CallChainContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_callChain);
		int _la;
		try {
			setState(226);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__9:
				_localctx = new MethodChainContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(213);
				match(T__9);
				setState(214);
				match(IDENTIFIER);
				setState(215);
				match(T__7);
				setState(217);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 2199023386980L) != 0) || ((((_la - 65)) & ~0x3f) == 0 && ((1L << (_la - 65)) & 139611594354329601L) != 0)) {
					{
					setState(216);
					expressionList();
					}
				}

				setState(219);
				match(T__8);
				}
				break;
			case T__10:
				_localctx = new PropertyChainContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(220);
				match(T__10);
				setState(221);
				match(IDENTIFIER);
				}
				break;
			case T__11:
				_localctx = new IndexChainContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(222);
				match(T__11);
				setState(223);
				expression(0);
				setState(224);
				match(T__12);
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
		enterRule(_localctx, 12, RULE_literal);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(228);
			_la = _input.LA(1);
			if ( !(((((_la - 118)) & ~0x3f) == 0 && ((1L << (_la - 118)) & 15L) != 0)) ) {
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
		public VariableContext variable() {
			return getRuleContext(VariableContext.class,0);
		}
		public SafeAccessContext safeAccess() {
			return getRuleContext(SafeAccessContext.class,0);
		}
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public VariableContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_variable; }
	}

	public final VariableContext variable() throws RecognitionException {
		return variable(0);
	}

	private VariableContext variable(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		VariableContext _localctx = new VariableContext(_ctx, _parentState);
		VariableContext _prevctx = _localctx;
		int _startState = 14;
		enterRecursionRule(_localctx, 14, RULE_variable, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(231);
			match(IDENTIFIER);
			}
			_ctx.stop = _input.LT(-1);
			setState(245);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,13,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(243);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,12,_ctx) ) {
					case 1:
						{
						_localctx = new VariableContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_variable);
						setState(233);
						if (!(precpred(_ctx, 3))) throw new FailedPredicateException(this, "precpred(_ctx, 3)");
						setState(234);
						safeAccess();
						}
						break;
					case 2:
						{
						_localctx = new VariableContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_variable);
						setState(235);
						if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
						setState(236);
						match(T__11);
						setState(237);
						expression(0);
						setState(238);
						match(T__12);
						}
						break;
					case 3:
						{
						_localctx = new VariableContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_variable);
						setState(240);
						if (!(precpred(_ctx, 1))) throw new FailedPredicateException(this, "precpred(_ctx, 1)");
						setState(241);
						match(T__10);
						setState(242);
						match(IDENTIFIER);
						}
						break;
					}
					} 
				}
				setState(247);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,13,_ctx);
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
	public static class SafeAccessContext extends ParserRuleContext {
		public TerminalNode IDENTIFIER() { return getToken(lua_grammar_antlr4Parser.IDENTIFIER, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public SafeAccessContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_safeAccess; }
	}

	public final SafeAccessContext safeAccess() throws RecognitionException {
		SafeAccessContext _localctx = new SafeAccessContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_safeAccess);
		try {
			setState(254);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__13:
				enterOuterAlt(_localctx, 1);
				{
				setState(248);
				match(T__13);
				setState(249);
				match(IDENTIFIER);
				}
				break;
			case T__14:
				enterOuterAlt(_localctx, 2);
				{
				setState(250);
				match(T__14);
				setState(251);
				expression(0);
				setState(252);
				match(T__12);
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
	public static class FunctionCallContext extends ParserRuleContext {
		public VariableContext variable() {
			return getRuleContext(VariableContext.class,0);
		}
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public TerminalNode IDENTIFIER() { return getToken(lua_grammar_antlr4Parser.IDENTIFIER, 0); }
		public FunctionCallContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_functionCall; }
	}

	public final FunctionCallContext functionCall() throws RecognitionException {
		FunctionCallContext _localctx = new FunctionCallContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_functionCall);
		int _la;
		try {
			setState(286);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,19,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(256);
				variable(0);
				setState(257);
				match(T__7);
				setState(266);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 2199023386980L) != 0) || ((((_la - 65)) & ~0x3f) == 0 && ((1L << (_la - 65)) & 139611590059362305L) != 0)) {
					{
					setState(258);
					expression(0);
					setState(263);
					_errHandler.sync(this);
					_la = _input.LA(1);
					while (_la==T__15) {
						{
						{
						setState(259);
						match(T__15);
						setState(260);
						expression(0);
						}
						}
						setState(265);
						_errHandler.sync(this);
						_la = _input.LA(1);
					}
					}
				}

				setState(268);
				match(T__8);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(270);
				variable(0);
				setState(271);
				match(T__9);
				setState(272);
				match(IDENTIFIER);
				setState(273);
				match(T__7);
				setState(282);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 2199023386980L) != 0) || ((((_la - 65)) & ~0x3f) == 0 && ((1L << (_la - 65)) & 139611590059362305L) != 0)) {
					{
					setState(274);
					expression(0);
					setState(279);
					_errHandler.sync(this);
					_la = _input.LA(1);
					while (_la==T__15) {
						{
						{
						setState(275);
						match(T__15);
						setState(276);
						expression(0);
						}
						}
						setState(281);
						_errHandler.sync(this);
						_la = _input.LA(1);
					}
					}
				}

				setState(284);
				match(T__8);
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
		enterRule(_localctx, 20, RULE_tableConstructor);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(288);
			match(T__16);
			setState(297);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 2199023391076L) != 0) || ((((_la - 65)) & ~0x3f) == 0 && ((1L << (_la - 65)) & 139611590059362305L) != 0)) {
				{
				setState(289);
				field();
				setState(294);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,20,_ctx);
				while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
					if ( _alt==1 ) {
						{
						{
						setState(290);
						match(T__15);
						setState(291);
						field();
						}
						} 
					}
					setState(296);
					_errHandler.sync(this);
					_alt = getInterpreter().adaptivePredict(_input,20,_ctx);
				}
				}
			}

			setState(301);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__15) {
				{
				setState(299);
				match(T__15);
				setState(300);
				metatable();
				}
			}

			setState(303);
			match(T__17);
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
		public MetamethodsContext metamethods() {
			return getRuleContext(MetamethodsContext.class,0);
		}
		public MetatableContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_metatable; }
	}

	public final MetatableContext metatable() throws RecognitionException {
		MetatableContext _localctx = new MetatableContext(_ctx, getState());
		enterRule(_localctx, 22, RULE_metatable);
		try {
			setState(312);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__18:
				enterOuterAlt(_localctx, 1);
				{
				setState(305);
				match(T__18);
				setState(306);
				match(T__19);
				setState(307);
				expression(0);
				}
				break;
			case T__16:
				enterOuterAlt(_localctx, 2);
				{
				setState(308);
				match(T__16);
				setState(309);
				metamethods();
				setState(310);
				match(T__17);
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
		enterRule(_localctx, 24, RULE_metamethods);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(314);
			metamethod();
			setState(319);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__15) {
				{
				{
				setState(315);
				match(T__15);
				setState(316);
				metamethod();
				}
				}
				setState(321);
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
		enterRule(_localctx, 26, RULE_labelStatement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(322);
			match(T__20);
			setState(323);
			match(IDENTIFIER);
			setState(324);
			match(T__20);
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
	public static class MatchArmContext extends ParserRuleContext {
		public PatternContext pattern() {
			return getRuleContext(PatternContext.class,0);
		}
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public MatchArmContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_matchArm; }
	}

	public final MatchArmContext matchArm() throws RecognitionException {
		MatchArmContext _localctx = new MatchArmContext(_ctx, getState());
		enterRule(_localctx, 28, RULE_matchArm);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(326);
			pattern(0);
			setState(329);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__21) {
				{
				setState(327);
				match(T__21);
				setState(328);
				expression(0);
				}
			}

			setState(331);
			match(T__22);
			setState(332);
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
	public static class PatternContext extends ParserRuleContext {
		public LiteralContext literal() {
			return getRuleContext(LiteralContext.class,0);
		}
		public TableConstructorContext tableConstructor() {
			return getRuleContext(TableConstructorContext.class,0);
		}
		public TerminalNode IDENTIFIER() { return getToken(lua_grammar_antlr4Parser.IDENTIFIER, 0); }
		public List<FieldPatternContext> fieldPattern() {
			return getRuleContexts(FieldPatternContext.class);
		}
		public FieldPatternContext fieldPattern(int i) {
			return getRuleContext(FieldPatternContext.class,i);
		}
		public List<PatternContext> pattern() {
			return getRuleContexts(PatternContext.class);
		}
		public PatternContext pattern(int i) {
			return getRuleContext(PatternContext.class,i);
		}
		public PatternContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_pattern; }
	}

	public final PatternContext pattern() throws RecognitionException {
		return pattern(0);
	}

	private PatternContext pattern(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		PatternContext _localctx = new PatternContext(_ctx, _parentState);
		PatternContext _prevctx = _localctx;
		int _startState = 30;
		enterRecursionRule(_localctx, 30, RULE_pattern, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(349);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,27,_ctx) ) {
			case 1:
				{
				setState(335);
				literal();
				}
				break;
			case 2:
				{
				setState(336);
				tableConstructor();
				}
				break;
			case 3:
				{
				setState(337);
				match(IDENTIFIER);
				}
				break;
			case 4:
				{
				setState(338);
				match(T__16);
				setState(339);
				fieldPattern();
				setState(344);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==T__15) {
					{
					{
					setState(340);
					match(T__15);
					setState(341);
					fieldPattern();
					}
					}
					setState(346);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(347);
				match(T__17);
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(356);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,28,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new PatternContext(_parentctx, _parentState);
					pushNewRecursionContext(_localctx, _startState, RULE_pattern);
					setState(351);
					if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
					setState(352);
					match(T__23);
					setState(353);
					pattern(3);
					}
					} 
				}
				setState(358);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,28,_ctx);
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
	public static class FieldPatternContext extends ParserRuleContext {
		public TerminalNode IDENTIFIER() { return getToken(lua_grammar_antlr4Parser.IDENTIFIER, 0); }
		public PatternContext pattern() {
			return getRuleContext(PatternContext.class,0);
		}
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public FieldPatternContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_fieldPattern; }
	}

	public final FieldPatternContext fieldPattern() throws RecognitionException {
		FieldPatternContext _localctx = new FieldPatternContext(_ctx, getState());
		enterRule(_localctx, 32, RULE_fieldPattern);
		try {
			setState(369);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,29,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(359);
				match(IDENTIFIER);
				setState(360);
				match(T__19);
				setState(361);
				pattern(0);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(362);
				match(T__11);
				setState(363);
				expression(0);
				setState(364);
				match(T__12);
				setState(365);
				match(T__19);
				setState(366);
				pattern(0);
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(368);
				pattern(0);
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
	public static class MetamethodContext extends ParserRuleContext {
		public MetamethodContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_metamethod; }
	}

	public final MetamethodContext metamethod() throws RecognitionException {
		MetamethodContext _localctx = new MetamethodContext(_ctx, getState());
		enterRule(_localctx, 34, RULE_metamethod);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(371);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 2198989701120L) != 0)) ) {
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
		enterRule(_localctx, 36, RULE_field);
		try {
			setState(386);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,30,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(373);
				match(IDENTIFIER);
				setState(374);
				match(T__19);
				setState(375);
				expression(0);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(376);
				match(T__11);
				setState(377);
				expression(0);
				setState(378);
				match(T__12);
				setState(379);
				match(T__19);
				setState(380);
				expression(0);
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(382);
				expression(0);
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(383);
				match(IDENTIFIER);
				setState(384);
				match(T__9);
				setState(385);
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
		enterRule(_localctx, 38, RULE_binaryOperation);
		try {
			setState(424);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,31,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(388);
				expression(0);
				setState(389);
				arithOp();
				setState(390);
				expression(0);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(392);
				expression(0);
				setState(393);
				bitwiseOp();
				setState(394);
				expression(0);
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(396);
				expression(0);
				setState(397);
				comparisonOp();
				setState(398);
				expression(0);
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(400);
				expression(0);
				setState(401);
				logicalOp();
				setState(402);
				expression(0);
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(404);
				expression(0);
				setState(405);
				concatOp();
				setState(406);
				expression(0);
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(408);
				expression(0);
				setState(409);
				compoundAssignOp();
				setState(410);
				expression(0);
				}
				break;
			case 7:
				enterOuterAlt(_localctx, 7);
				{
				setState(412);
				expression(0);
				setState(413);
				coalesceOp();
				setState(414);
				expression(0);
				}
				break;
			case 8:
				enterOuterAlt(_localctx, 8);
				{
				setState(416);
				expression(0);
				setState(417);
				shiftAssignOp();
				setState(418);
				expression(0);
				}
				break;
			case 9:
				enterOuterAlt(_localctx, 9);
				{
				setState(420);
				expression(0);
				setState(421);
				incrOp();
				setState(422);
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
		enterRule(_localctx, 40, RULE_unaryOperation);
		try {
			setState(431);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,32,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(426);
				unaryOp();
				setState(427);
				expression(0);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(429);
				match(T__40);
				setState(430);
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
		enterRule(_localctx, 42, RULE_controlFlowStatement);
		try {
			setState(441);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__41:
				enterOuterAlt(_localctx, 1);
				{
				setState(433);
				ifStatement();
				}
				break;
			case T__45:
				enterOuterAlt(_localctx, 2);
				{
				setState(434);
				whileStatement();
				}
				break;
			case T__47:
				enterOuterAlt(_localctx, 3);
				{
				setState(435);
				repeatStatement();
				}
				break;
			case T__49:
				enterOuterAlt(_localctx, 4);
				{
				setState(436);
				forStatement();
				}
				break;
			case T__51:
				enterOuterAlt(_localctx, 5);
				{
				setState(437);
				breakStatement();
				}
				break;
			case T__52:
				enterOuterAlt(_localctx, 6);
				{
				setState(438);
				gotoStatement();
				}
				break;
			case T__53:
				enterOuterAlt(_localctx, 7);
				{
				setState(439);
				coroutineStatement();
				}
				break;
			case T__61:
			case T__62:
				enterOuterAlt(_localctx, 8);
				{
				setState(440);
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
		enterRule(_localctx, 44, RULE_ifStatement);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(443);
			match(T__41);
			setState(444);
			expression(0);
			setState(445);
			match(T__42);
			setState(446);
			block();
			setState(454);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__43) {
				{
				{
				setState(447);
				match(T__43);
				setState(448);
				expression(0);
				setState(449);
				match(T__42);
				setState(450);
				block();
				}
				}
				setState(456);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(459);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__44) {
				{
				setState(457);
				match(T__44);
				setState(458);
				block();
				}
			}

			setState(461);
			match(T__3);
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
		enterRule(_localctx, 46, RULE_whileStatement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(463);
			match(T__45);
			setState(464);
			expression(0);
			setState(465);
			match(T__46);
			setState(466);
			block();
			setState(467);
			match(T__3);
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
		enterRule(_localctx, 48, RULE_repeatStatement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(469);
			match(T__47);
			setState(470);
			block();
			setState(471);
			match(T__48);
			setState(472);
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
		enterRule(_localctx, 50, RULE_forStatement);
		int _la;
		try {
			setState(496);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,37,_ctx) ) {
			case 1:
				_localctx = new NumericForContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(474);
				match(T__49);
				setState(475);
				match(IDENTIFIER);
				setState(476);
				match(T__19);
				setState(477);
				expression(0);
				setState(478);
				match(T__15);
				setState(479);
				expression(0);
				setState(482);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==T__15) {
					{
					setState(480);
					match(T__15);
					setState(481);
					expression(0);
					}
				}

				setState(484);
				match(T__46);
				setState(485);
				block();
				setState(486);
				match(T__3);
				}
				break;
			case 2:
				_localctx = new GenericForContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(488);
				match(T__49);
				setState(489);
				identifierList();
				setState(490);
				match(T__50);
				setState(491);
				expressionList();
				setState(492);
				match(T__46);
				setState(493);
				block();
				setState(494);
				match(T__3);
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
		enterRule(_localctx, 52, RULE_breakStatement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(498);
			match(T__51);
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
		enterRule(_localctx, 54, RULE_gotoStatement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(500);
			match(T__52);
			setState(501);
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
		enterRule(_localctx, 56, RULE_coroutineStatement);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(503);
			match(T__53);
			setState(504);
			match(T__10);
			setState(505);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 4575657221408423936L) != 0)) ) {
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
		public ExpressionListContext expressionList() {
			return getRuleContext(ExpressionListContext.class,0);
		}
		public NamedArgsContext namedArgs() {
			return getRuleContext(NamedArgsContext.class,0);
		}
		public TerminalNode IDENTIFIER() { return getToken(lua_grammar_antlr4Parser.IDENTIFIER, 0); }
		public ProtectedCallStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_protectedCallStatement; }
	}

	public final ProtectedCallStatementContext protectedCallStatement() throws RecognitionException {
		ProtectedCallStatementContext _localctx = new ProtectedCallStatementContext(_ctx, getState());
		enterRule(_localctx, 58, RULE_protectedCallStatement);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(507);
			_la = _input.LA(1);
			if ( !(_la==T__61 || _la==T__62) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			setState(509);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__9 || _la==T__10) {
				{
				setState(508);
				_la = _input.LA(1);
				if ( !(_la==T__9 || _la==T__10) ) {
				_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				}
			}

			setState(512);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==IDENTIFIER) {
				{
				setState(511);
				match(IDENTIFIER);
				}
			}

			}
			setState(514);
			match(T__7);
			setState(517);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,40,_ctx) ) {
			case 1:
				{
				setState(515);
				expressionList();
				}
				break;
			case 2:
				{
				setState(516);
				namedArgs();
				}
				break;
			}
			setState(519);
			match(T__8);
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
	public static class NamedArgsContext extends ParserRuleContext {
		public List<TerminalNode> IDENTIFIER() { return getTokens(lua_grammar_antlr4Parser.IDENTIFIER); }
		public TerminalNode IDENTIFIER(int i) {
			return getToken(lua_grammar_antlr4Parser.IDENTIFIER, i);
		}
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public NamedArgsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_namedArgs; }
	}

	public final NamedArgsContext namedArgs() throws RecognitionException {
		NamedArgsContext _localctx = new NamedArgsContext(_ctx, getState());
		enterRule(_localctx, 60, RULE_namedArgs);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(521);
			match(IDENTIFIER);
			setState(522);
			match(T__19);
			setState(523);
			expression(0);
			setState(530);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__15) {
				{
				{
				setState(524);
				match(T__15);
				setState(525);
				match(IDENTIFIER);
				setState(526);
				match(T__19);
				setState(527);
				expression(0);
				}
				}
				setState(532);
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
		enterRule(_localctx, 62, RULE_block);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(536);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & -4578678679359455230L) != 0) || ((((_la - 64)) & ~0x3f) == 0 && ((1L << (_la - 64)) & 9015995347763207L) != 0)) {
				{
				{
				setState(533);
				statement();
				}
				}
				setState(538);
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
		public TypeAnnotationContext typeAnnotation() {
			return getRuleContext(TypeAnnotationContext.class,0);
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
		enterRule(_localctx, 64, RULE_localDeclaration);
		int _la;
		try {
			setState(577);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,48,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(539);
				match(T__63);
				setState(541);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==T__9) {
					{
					setState(540);
					typeAnnotation();
					}
				}

				setState(543);
				match(IDENTIFIER);
				setState(546);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==T__19) {
					{
					setState(544);
					match(T__19);
					setState(545);
					expression(0);
					}
				}

				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(548);
				match(T__63);
				setState(549);
				match(IDENTIFIER);
				setState(554);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==T__15) {
					{
					{
					setState(550);
					match(T__15);
					setState(551);
					match(IDENTIFIER);
					}
					}
					setState(556);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(557);
				match(T__19);
				setState(558);
				expressionList();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(559);
				match(T__63);
				setState(560);
				match(T__64);
				setState(561);
				match(IDENTIFIER);
				setState(562);
				match(T__7);
				setState(571);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==IDENTIFIER) {
					{
					setState(563);
					match(IDENTIFIER);
					setState(568);
					_errHandler.sync(this);
					_la = _input.LA(1);
					while (_la==T__15) {
						{
						{
						setState(564);
						match(T__15);
						setState(565);
						match(IDENTIFIER);
						}
						}
						setState(570);
						_errHandler.sync(this);
						_la = _input.LA(1);
					}
					}
				}

				setState(573);
				match(T__8);
				setState(574);
				block();
				setState(575);
				match(T__3);
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
		enterRule(_localctx, 66, RULE_functionDeclaration);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(579);
			match(T__64);
			setState(584);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,49,_ctx) ) {
			case 1:
				{
				setState(580);
				match(IDENTIFIER);
				setState(581);
				match(T__10);
				}
				break;
			case 2:
				{
				setState(582);
				match(IDENTIFIER);
				setState(583);
				match(T__9);
				}
				break;
			}
			setState(586);
			match(IDENTIFIER);
			setState(587);
			match(T__7);
			setState(601);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case IDENTIFIER:
				{
				setState(588);
				match(IDENTIFIER);
				setState(593);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,50,_ctx);
				while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
					if ( _alt==1 ) {
						{
						{
						setState(589);
						match(T__15);
						setState(590);
						match(IDENTIFIER);
						}
						} 
					}
					setState(595);
					_errHandler.sync(this);
					_alt = getInterpreter().adaptivePredict(_input,50,_ctx);
				}
				setState(598);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==T__15) {
					{
					setState(596);
					match(T__15);
					setState(597);
					varargOp();
					}
				}

				}
				break;
			case T__96:
				{
				setState(600);
				varargOp();
				}
				break;
			case T__8:
				break;
			default:
				break;
			}
			setState(603);
			match(T__8);
			setState(604);
			block();
			setState(605);
			match(T__3);
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
		enterRule(_localctx, 68, RULE_returnStatement);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(607);
			match(T__65);
			setState(617);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,54,_ctx) ) {
			case 1:
				{
				setState(608);
				expression(0);
				setState(613);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==T__15) {
					{
					{
					setState(609);
					match(T__15);
					setState(610);
					expression(0);
					}
					}
					setState(615);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				}
				break;
			case 2:
				{
				setState(616);
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
		public SafeAccessContext safeAccess() {
			return getRuleContext(SafeAccessContext.class,0);
		}
		public NonNullAssertOpContext nonNullAssertOp() {
			return getRuleContext(NonNullAssertOpContext.class,0);
		}
		public OperatorGroupContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_operatorGroup; }
	}

	public final OperatorGroupContext operatorGroup() throws RecognitionException {
		OperatorGroupContext _localctx = new OperatorGroupContext(_ctx, getState());
		enterRule(_localctx, 70, RULE_operatorGroup);
		try {
			setState(633);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,55,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(619);
				logicalOp();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(620);
				comparisonOp();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(621);
				arithOp();
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(622);
				bitwiseOp();
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(623);
				assignOp();
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(624);
				unaryOp();
				}
				break;
			case 7:
				enterOuterAlt(_localctx, 7);
				{
				setState(625);
				concatOp();
				}
				break;
			case 8:
				enterOuterAlt(_localctx, 8);
				{
				setState(626);
				compoundAssignOp();
				}
				break;
			case 9:
				enterOuterAlt(_localctx, 9);
				{
				setState(627);
				incrOp();
				}
				break;
			case 10:
				enterOuterAlt(_localctx, 10);
				{
				setState(628);
				coalesceOp();
				}
				break;
			case 11:
				enterOuterAlt(_localctx, 11);
				{
				setState(629);
				shiftAssignOp();
				}
				break;
			case 12:
				enterOuterAlt(_localctx, 12);
				{
				setState(630);
				match(T__22);
				}
				break;
			case 13:
				enterOuterAlt(_localctx, 13);
				{
				setState(631);
				safeAccess();
				}
				break;
			case 14:
				enterOuterAlt(_localctx, 14);
				{
				setState(632);
				nonNullAssertOp();
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
		enterRule(_localctx, 72, RULE_logicalOp);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(635);
			_la = _input.LA(1);
			if ( !(_la==T__66 || _la==T__67) ) {
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
		enterRule(_localctx, 74, RULE_comparisonOp);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(637);
			_la = _input.LA(1);
			if ( !(((((_la - 69)) & ~0x3f) == 0 && ((1L << (_la - 69)) & 63L) != 0)) ) {
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
		enterRule(_localctx, 76, RULE_arithOp);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(639);
			_la = _input.LA(1);
			if ( !(((((_la - 75)) & ~0x3f) == 0 && ((1L << (_la - 75)) & 127L) != 0)) ) {
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
		enterRule(_localctx, 78, RULE_bitwiseOp);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(641);
			_la = _input.LA(1);
			if ( !(((((_la - 24)) & ~0x3f) == 0 && ((1L << (_la - 24)) & 4323455642275676161L) != 0)) ) {
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
		enterRule(_localctx, 80, RULE_assignOp);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(643);
			_la = _input.LA(1);
			if ( !(_la==T__19 || ((((_la - 86)) & ~0x3f) == 0 && ((1L << (_la - 86)) & 255L) != 0)) ) {
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
		enterRule(_localctx, 82, RULE_unaryOp);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(645);
			_la = _input.LA(1);
			if ( !(((((_la - 41)) & ~0x3f) == 0 && ((1L << (_la - 41)) & 27026030170472449L) != 0)) ) {
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
		enterRule(_localctx, 84, RULE_concatOp);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(647);
			match(T__95);
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
		enterRule(_localctx, 86, RULE_varargOp);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(649);
			match(T__96);
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
		enterRule(_localctx, 88, RULE_compoundAssignOp);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(651);
			_la = _input.LA(1);
			if ( !(((((_la - 86)) & ~0x3f) == 0 && ((1L << (_la - 86)) & 61503L) != 0)) ) {
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
		enterRule(_localctx, 90, RULE_incrOp);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(653);
			_la = _input.LA(1);
			if ( !(_la==T__101 || _la==T__102) ) {
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
		enterRule(_localctx, 92, RULE_coalesceOp);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(655);
			match(T__103);
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
		enterRule(_localctx, 94, RULE_shiftAssignOp);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(657);
			_la = _input.LA(1);
			if ( !(_la==T__104 || _la==T__105) ) {
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
	public static class NonNullAssertOpContext extends ParserRuleContext {
		public NonNullAssertOpContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_nonNullAssertOp; }
	}

	public final NonNullAssertOpContext nonNullAssertOp() throws RecognitionException {
		NonNullAssertOpContext _localctx = new NonNullAssertOpContext(_ctx, getState());
		enterRule(_localctx, 96, RULE_nonNullAssertOp);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(659);
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
		enterRule(_localctx, 98, RULE_identifierList);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(661);
			match(IDENTIFIER);
			setState(666);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__15) {
				{
				{
				setState(662);
				match(T__15);
				setState(663);
				match(IDENTIFIER);
				}
				}
				setState(668);
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
		public VarargOpContext varargOp() {
			return getRuleContext(VarargOpContext.class,0);
		}
		public ExpressionListContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expressionList; }
	}

	public final ExpressionListContext expressionList() throws RecognitionException {
		ExpressionListContext _localctx = new ExpressionListContext(_ctx, getState());
		enterRule(_localctx, 100, RULE_expressionList);
		int _la;
		try {
			setState(678);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__1:
			case T__4:
			case T__5:
			case T__7:
			case T__16:
			case T__40:
			case T__64:
			case T__75:
			case T__82:
			case T__93:
			case T__94:
			case IDENTIFIER:
			case BOOL:
			case NIL:
			case NUMBER:
			case STRING:
				enterOuterAlt(_localctx, 1);
				{
				setState(669);
				expression(0);
				setState(674);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==T__15) {
					{
					{
					setState(670);
					match(T__15);
					setState(671);
					expression(0);
					}
					}
					setState(676);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				}
				break;
			case T__96:
				enterOuterAlt(_localctx, 2);
				{
				setState(677);
				varargOp();
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
	public static class FunctionExpressionContext extends ParserRuleContext {
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
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
		enterRule(_localctx, 102, RULE_functionExpression);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(680);
			match(T__64);
			setState(683);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==IDENTIFIER) {
				{
				setState(681);
				match(IDENTIFIER);
				setState(682);
				match(T__9);
				}
			}

			setState(685);
			match(T__7);
			setState(694);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==IDENTIFIER) {
				{
				setState(686);
				match(IDENTIFIER);
				setState(691);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==T__15) {
					{
					{
					setState(687);
					match(T__15);
					setState(688);
					match(IDENTIFIER);
					}
					}
					setState(693);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				}
			}

			setState(696);
			match(T__8);
			setState(702);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__0:
			case T__3:
			case T__20:
			case T__41:
			case T__45:
			case T__47:
			case T__49:
			case T__51:
			case T__52:
			case T__53:
			case T__61:
			case T__62:
			case T__63:
			case T__64:
			case T__65:
			case T__106:
			case IDENTIFIER:
				{
				setState(697);
				block();
				setState(698);
				match(T__3);
				}
				break;
			case T__22:
				{
				setState(700);
				match(T__22);
				setState(701);
				expression(0);
				}
				break;
			default:
				throw new NoViableAltException(this);
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
		enterRule(_localctx, 104, RULE_selectStatement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(704);
			match(T__106);
			setState(705);
			match(T__7);
			setState(708);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,63,_ctx) ) {
			case 1:
				{
				setState(706);
				match(T__40);
				}
				break;
			case 2:
				{
				setState(707);
				expression(0);
				}
				break;
			}
			setState(710);
			match(T__15);
			setState(711);
			expression(0);
			setState(712);
			match(T__8);
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
	public static class LambdaExpressionContext extends ParserRuleContext {
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public List<TerminalNode> IDENTIFIER() { return getTokens(lua_grammar_antlr4Parser.IDENTIFIER); }
		public TerminalNode IDENTIFIER(int i) {
			return getToken(lua_grammar_antlr4Parser.IDENTIFIER, i);
		}
		public LambdaExpressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_lambdaExpression; }
	}

	public final LambdaExpressionContext lambdaExpression() throws RecognitionException {
		LambdaExpressionContext _localctx = new LambdaExpressionContext(_ctx, getState());
		enterRule(_localctx, 106, RULE_lambdaExpression);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(714);
			match(T__7);
			setState(723);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==IDENTIFIER) {
				{
				setState(715);
				match(IDENTIFIER);
				setState(720);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==T__15) {
					{
					{
					setState(716);
					match(T__15);
					setState(717);
					match(IDENTIFIER);
					}
					}
					setState(722);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				}
			}

			setState(725);
			match(T__8);
			setState(726);
			match(T__22);
			setState(727);
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
	public static class TypeAnnotationContext extends ParserRuleContext {
		public TypeSpecContext typeSpec() {
			return getRuleContext(TypeSpecContext.class,0);
		}
		public TypeAnnotationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_typeAnnotation; }
	}

	public final TypeAnnotationContext typeAnnotation() throws RecognitionException {
		TypeAnnotationContext _localctx = new TypeAnnotationContext(_ctx, getState());
		enterRule(_localctx, 108, RULE_typeAnnotation);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(729);
			match(T__9);
			setState(730);
			typeSpec(0);
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
	public static class TypeSpecContext extends ParserRuleContext {
		public List<TypeSpecContext> typeSpec() {
			return getRuleContexts(TypeSpecContext.class);
		}
		public TypeSpecContext typeSpec(int i) {
			return getRuleContext(TypeSpecContext.class,i);
		}
		public TypeSpecContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_typeSpec; }
	}

	public final TypeSpecContext typeSpec() throws RecognitionException {
		return typeSpec(0);
	}

	private TypeSpecContext typeSpec(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		TypeSpecContext _localctx = new TypeSpecContext(_ctx, _parentState);
		TypeSpecContext _prevctx = _localctx;
		int _startState = 110;
		enterRecursionRule(_localctx, 110, RULE_typeSpec, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(745);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__107:
				{
				setState(733);
				match(T__107);
				}
				break;
			case T__108:
				{
				setState(734);
				match(T__108);
				}
				break;
			case T__109:
				{
				setState(735);
				match(T__109);
				}
				break;
			case T__110:
				{
				setState(736);
				match(T__110);
				}
				break;
			case T__64:
				{
				setState(737);
				match(T__64);
				}
				break;
			case T__111:
				{
				setState(738);
				match(T__111);
				}
				break;
			case T__113:
				{
				setState(739);
				match(T__113);
				setState(740);
				typeSpec(0);
				setState(741);
				match(T__15);
				setState(742);
				typeSpec(0);
				setState(743);
				match(T__72);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			_ctx.stop = _input.LT(-1);
			setState(751);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,67,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new TypeSpecContext(_parentctx, _parentState);
					pushNewRecursionContext(_localctx, _startState, RULE_typeSpec);
					setState(747);
					if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
					setState(748);
					match(T__112);
					}
					} 
				}
				setState(753);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,67,_ctx);
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
	public static class ExperimentalExpressionContext extends ParserRuleContext {
		public SafeNavigationContext safeNavigation() {
			return getRuleContext(SafeNavigationContext.class,0);
		}
		public PipeOperatorContext pipeOperator() {
			return getRuleContext(PipeOperatorContext.class,0);
		}
		public DecoratorSyntaxContext decoratorSyntax() {
			return getRuleContext(DecoratorSyntaxContext.class,0);
		}
		public ExperimentalExpressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_experimentalExpression; }
	}

	public final ExperimentalExpressionContext experimentalExpression() throws RecognitionException {
		ExperimentalExpressionContext _localctx = new ExperimentalExpressionContext(_ctx, getState());
		enterRule(_localctx, 112, RULE_experimentalExpression);
		try {
			setState(757);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,68,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(754);
				safeNavigation();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(755);
				pipeOperator();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(756);
				decoratorSyntax();
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
	public static class SafeNavigationContext extends ParserRuleContext {
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public TerminalNode IDENTIFIER() { return getToken(lua_grammar_antlr4Parser.IDENTIFIER, 0); }
		public SafeNavigationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_safeNavigation; }
	}

	public final SafeNavigationContext safeNavigation() throws RecognitionException {
		SafeNavigationContext _localctx = new SafeNavigationContext(_ctx, getState());
		enterRule(_localctx, 114, RULE_safeNavigation);
		try {
			setState(768);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,69,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(759);
				expression(0);
				setState(760);
				match(T__13);
				setState(761);
				match(IDENTIFIER);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(763);
				expression(0);
				setState(764);
				match(T__14);
				setState(765);
				expression(0);
				setState(766);
				match(T__12);
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
	public static class PipeOperatorContext extends ParserRuleContext {
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public PipeOperatorContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_pipeOperator; }
	}

	public final PipeOperatorContext pipeOperator() throws RecognitionException {
		PipeOperatorContext _localctx = new PipeOperatorContext(_ctx, getState());
		enterRule(_localctx, 116, RULE_pipeOperator);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(770);
			expression(0);
			setState(771);
			match(T__114);
			setState(772);
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
	public static class DecoratorSyntaxContext extends ParserRuleContext {
		public TerminalNode IDENTIFIER() { return getToken(lua_grammar_antlr4Parser.IDENTIFIER, 0); }
		public ExpressionListContext expressionList() {
			return getRuleContext(ExpressionListContext.class,0);
		}
		public DecoratorSyntaxContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_decoratorSyntax; }
	}

	public final DecoratorSyntaxContext decoratorSyntax() throws RecognitionException {
		DecoratorSyntaxContext _localctx = new DecoratorSyntaxContext(_ctx, getState());
		enterRule(_localctx, 118, RULE_decoratorSyntax);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(774);
			match(T__115);
			setState(775);
			match(IDENTIFIER);
			setState(781);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__7) {
				{
				setState(776);
				match(T__7);
				setState(778);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 2199023386980L) != 0) || ((((_la - 65)) & ~0x3f) == 0 && ((1L << (_la - 65)) & 139611594354329601L) != 0)) {
					{
					setState(777);
					expressionList();
					}
				}

				setState(780);
				match(T__8);
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

	public boolean sempred(RuleContext _localctx, int ruleIndex, int predIndex) {
		switch (ruleIndex) {
		case 3:
			return expression_sempred((ExpressionContext)_localctx, predIndex);
		case 7:
			return variable_sempred((VariableContext)_localctx, predIndex);
		case 15:
			return pattern_sempred((PatternContext)_localctx, predIndex);
		case 55:
			return typeSpec_sempred((TypeSpecContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean expression_sempred(ExpressionContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 7);
		case 1:
			return precpred(_ctx, 6);
		case 2:
			return precpred(_ctx, 1);
		}
		return true;
	}
	private boolean variable_sempred(VariableContext _localctx, int predIndex) {
		switch (predIndex) {
		case 3:
			return precpred(_ctx, 3);
		case 4:
			return precpred(_ctx, 2);
		case 5:
			return precpred(_ctx, 1);
		}
		return true;
	}
	private boolean pattern_sempred(PatternContext _localctx, int predIndex) {
		switch (predIndex) {
		case 6:
			return precpred(_ctx, 2);
		}
		return true;
	}
	private boolean typeSpec_sempred(TypeSpecContext _localctx, int predIndex) {
		switch (predIndex) {
		case 7:
			return precpred(_ctx, 2);
		}
		return true;
	}

	public static final String _serializedATN =
		"\u0004\u0001{\u0310\u0002\u0000\u0007\u0000\u0002\u0001\u0007\u0001\u0002"+
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
		"-\u0007-\u0002.\u0007.\u0002/\u0007/\u00020\u00070\u00021\u00071\u0002"+
		"2\u00072\u00023\u00073\u00024\u00074\u00025\u00075\u00026\u00076\u0002"+
		"7\u00077\u00028\u00078\u00029\u00079\u0002:\u0007:\u0002;\u0007;\u0001"+
		"\u0000\u0005\u0000z\b\u0000\n\u0000\f\u0000}\t\u0000\u0001\u0000\u0001"+
		"\u0000\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001"+
		"\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0003\u0001\u008a\b\u0001\u0001"+
		"\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0003\u0002\u0091"+
		"\b\u0002\u0001\u0002\u0001\u0002\u0001\u0003\u0001\u0003\u0001\u0003\u0001"+
		"\u0003\u0001\u0003\u0005\u0003\u009a\b\u0003\n\u0003\f\u0003\u009d\t\u0003"+
		"\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003"+
		"\u0001\u0003\u0004\u0003\u00a6\b\u0003\u000b\u0003\f\u0003\u00a7\u0001"+
		"\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001"+
		"\u0003\u0001\u0003\u0003\u0003\u00b2\b\u0003\u0001\u0003\u0001\u0003\u0001"+
		"\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001"+
		"\u0003\u0001\u0003\u0005\u0003\u00be\b\u0003\n\u0003\f\u0003\u00c1\t\u0003"+
		"\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004"+
		"\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0003\u0004"+
		"\u00ce\b\u0004\u0001\u0004\u0005\u0004\u00d1\b\u0004\n\u0004\f\u0004\u00d4"+
		"\t\u0004\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0003\u0005\u00da"+
		"\b\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001"+
		"\u0005\u0001\u0005\u0003\u0005\u00e3\b\u0005\u0001\u0006\u0001\u0006\u0001"+
		"\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001"+
		"\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001"+
		"\u0007\u0005\u0007\u00f4\b\u0007\n\u0007\f\u0007\u00f7\t\u0007\u0001\b"+
		"\u0001\b\u0001\b\u0001\b\u0001\b\u0001\b\u0003\b\u00ff\b\b\u0001\t\u0001"+
		"\t\u0001\t\u0001\t\u0001\t\u0005\t\u0106\b\t\n\t\f\t\u0109\t\t\u0003\t"+
		"\u010b\b\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001"+
		"\t\u0001\t\u0005\t\u0116\b\t\n\t\f\t\u0119\t\t\u0003\t\u011b\b\t\u0001"+
		"\t\u0001\t\u0003\t\u011f\b\t\u0001\n\u0001\n\u0001\n\u0001\n\u0005\n\u0125"+
		"\b\n\n\n\f\n\u0128\t\n\u0003\n\u012a\b\n\u0001\n\u0001\n\u0003\n\u012e"+
		"\b\n\u0001\n\u0001\n\u0001\u000b\u0001\u000b\u0001\u000b\u0001\u000b\u0001"+
		"\u000b\u0001\u000b\u0001\u000b\u0003\u000b\u0139\b\u000b\u0001\f\u0001"+
		"\f\u0001\f\u0005\f\u013e\b\f\n\f\f\f\u0141\t\f\u0001\r\u0001\r\u0001\r"+
		"\u0001\r\u0001\u000e\u0001\u000e\u0001\u000e\u0003\u000e\u014a\b\u000e"+
		"\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000f\u0001\u000f\u0001\u000f"+
		"\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0005\u000f"+
		"\u0157\b\u000f\n\u000f\f\u000f\u015a\t\u000f\u0001\u000f\u0001\u000f\u0003"+
		"\u000f\u015e\b\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0005\u000f\u0163"+
		"\b\u000f\n\u000f\f\u000f\u0166\t\u000f\u0001\u0010\u0001\u0010\u0001\u0010"+
		"\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010"+
		"\u0001\u0010\u0003\u0010\u0172\b\u0010\u0001\u0011\u0001\u0011\u0001\u0012"+
		"\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012"+
		"\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012"+
		"\u0003\u0012\u0183\b\u0012\u0001\u0013\u0001\u0013\u0001\u0013\u0001\u0013"+
		"\u0001\u0013\u0001\u0013\u0001\u0013\u0001\u0013\u0001\u0013\u0001\u0013"+
		"\u0001\u0013\u0001\u0013\u0001\u0013\u0001\u0013\u0001\u0013\u0001\u0013"+
		"\u0001\u0013\u0001\u0013\u0001\u0013\u0001\u0013\u0001\u0013\u0001\u0013"+
		"\u0001\u0013\u0001\u0013\u0001\u0013\u0001\u0013\u0001\u0013\u0001\u0013"+
		"\u0001\u0013\u0001\u0013\u0001\u0013\u0001\u0013\u0001\u0013\u0001\u0013"+
		"\u0001\u0013\u0001\u0013\u0003\u0013\u01a9\b\u0013\u0001\u0014\u0001\u0014"+
		"\u0001\u0014\u0001\u0014\u0001\u0014\u0003\u0014\u01b0\b\u0014\u0001\u0015"+
		"\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015"+
		"\u0001\u0015\u0003\u0015\u01ba\b\u0015\u0001\u0016\u0001\u0016\u0001\u0016"+
		"\u0001\u0016\u0001\u0016\u0001\u0016\u0001\u0016\u0001\u0016\u0001\u0016"+
		"\u0005\u0016\u01c5\b\u0016\n\u0016\f\u0016\u01c8\t\u0016\u0001\u0016\u0001"+
		"\u0016\u0003\u0016\u01cc\b\u0016\u0001\u0016\u0001\u0016\u0001\u0017\u0001"+
		"\u0017\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0018\u0001"+
		"\u0018\u0001\u0018\u0001\u0018\u0001\u0018\u0001\u0019\u0001\u0019\u0001"+
		"\u0019\u0001\u0019\u0001\u0019\u0001\u0019\u0001\u0019\u0001\u0019\u0003"+
		"\u0019\u01e3\b\u0019\u0001\u0019\u0001\u0019\u0001\u0019\u0001\u0019\u0001"+
		"\u0019\u0001\u0019\u0001\u0019\u0001\u0019\u0001\u0019\u0001\u0019\u0001"+
		"\u0019\u0001\u0019\u0003\u0019\u01f1\b\u0019\u0001\u001a\u0001\u001a\u0001"+
		"\u001b\u0001\u001b\u0001\u001b\u0001\u001c\u0001\u001c\u0001\u001c\u0001"+
		"\u001c\u0001\u001d\u0001\u001d\u0003\u001d\u01fe\b\u001d\u0001\u001d\u0003"+
		"\u001d\u0201\b\u001d\u0001\u001d\u0001\u001d\u0001\u001d\u0003\u001d\u0206"+
		"\b\u001d\u0001\u001d\u0001\u001d\u0001\u001e\u0001\u001e\u0001\u001e\u0001"+
		"\u001e\u0001\u001e\u0001\u001e\u0001\u001e\u0005\u001e\u0211\b\u001e\n"+
		"\u001e\f\u001e\u0214\t\u001e\u0001\u001f\u0005\u001f\u0217\b\u001f\n\u001f"+
		"\f\u001f\u021a\t\u001f\u0001 \u0001 \u0003 \u021e\b \u0001 \u0001 \u0001"+
		" \u0003 \u0223\b \u0001 \u0001 \u0001 \u0001 \u0005 \u0229\b \n \f \u022c"+
		"\t \u0001 \u0001 \u0001 \u0001 \u0001 \u0001 \u0001 \u0001 \u0001 \u0005"+
		" \u0237\b \n \f \u023a\t \u0003 \u023c\b \u0001 \u0001 \u0001 \u0001 "+
		"\u0003 \u0242\b \u0001!\u0001!\u0001!\u0001!\u0001!\u0003!\u0249\b!\u0001"+
		"!\u0001!\u0001!\u0001!\u0001!\u0005!\u0250\b!\n!\f!\u0253\t!\u0001!\u0001"+
		"!\u0003!\u0257\b!\u0001!\u0003!\u025a\b!\u0001!\u0001!\u0001!\u0001!\u0001"+
		"\"\u0001\"\u0001\"\u0001\"\u0005\"\u0264\b\"\n\"\f\"\u0267\t\"\u0001\""+
		"\u0003\"\u026a\b\"\u0001#\u0001#\u0001#\u0001#\u0001#\u0001#\u0001#\u0001"+
		"#\u0001#\u0001#\u0001#\u0001#\u0001#\u0001#\u0003#\u027a\b#\u0001$\u0001"+
		"$\u0001%\u0001%\u0001&\u0001&\u0001\'\u0001\'\u0001(\u0001(\u0001)\u0001"+
		")\u0001*\u0001*\u0001+\u0001+\u0001,\u0001,\u0001-\u0001-\u0001.\u0001"+
		".\u0001/\u0001/\u00010\u00010\u00011\u00011\u00011\u00051\u0299\b1\n1"+
		"\f1\u029c\t1\u00012\u00012\u00012\u00052\u02a1\b2\n2\f2\u02a4\t2\u0001"+
		"2\u00032\u02a7\b2\u00013\u00013\u00013\u00033\u02ac\b3\u00013\u00013\u0001"+
		"3\u00013\u00053\u02b2\b3\n3\f3\u02b5\t3\u00033\u02b7\b3\u00013\u00013"+
		"\u00013\u00013\u00013\u00013\u00033\u02bf\b3\u00014\u00014\u00014\u0001"+
		"4\u00034\u02c5\b4\u00014\u00014\u00014\u00014\u00015\u00015\u00015\u0001"+
		"5\u00055\u02cf\b5\n5\f5\u02d2\t5\u00035\u02d4\b5\u00015\u00015\u00015"+
		"\u00015\u00016\u00016\u00016\u00017\u00017\u00017\u00017\u00017\u0001"+
		"7\u00017\u00017\u00017\u00017\u00017\u00017\u00017\u00037\u02ea\b7\u0001"+
		"7\u00017\u00057\u02ee\b7\n7\f7\u02f1\t7\u00018\u00018\u00018\u00038\u02f6"+
		"\b8\u00019\u00019\u00019\u00019\u00019\u00019\u00019\u00019\u00019\u0003"+
		"9\u0301\b9\u0001:\u0001:\u0001:\u0001:\u0001;\u0001;\u0001;\u0001;\u0003"+
		";\u030b\b;\u0001;\u0003;\u030e\b;\u0001;\u0000\u0004\u0006\u000e\u001e"+
		"n<\u0000\u0002\u0004\u0006\b\n\f\u000e\u0010\u0012\u0014\u0016\u0018\u001a"+
		"\u001c\u001e \"$&(*,.02468:<>@BDFHJLNPRTVXZ\\^`bdfhjlnprtv\u0000\u000e"+
		"\u0001\u0000vy\u0001\u0000\u0019(\u0001\u00007=\u0001\u0000>?\u0001\u0000"+
		"\n\u000b\u0001\u0000CD\u0001\u0000EJ\u0001\u0000KQ\u0002\u0000\u0018\u0018"+
		"RU\u0002\u0000\u0014\u0014V]\u0004\u0000))LLSS^_\u0002\u0000V[be\u0001"+
		"\u0000fg\u0001\u0000ij\u0357\u0000{\u0001\u0000\u0000\u0000\u0002\u0089"+
		"\u0001\u0000\u0000\u0000\u0004\u008b\u0001\u0000\u0000\u0000\u0006\u00b1"+
		"\u0001\u0000\u0000\u0000\b\u00cd\u0001\u0000\u0000\u0000\n\u00e2\u0001"+
		"\u0000\u0000\u0000\f\u00e4\u0001\u0000\u0000\u0000\u000e\u00e6\u0001\u0000"+
		"\u0000\u0000\u0010\u00fe\u0001\u0000\u0000\u0000\u0012\u011e\u0001\u0000"+
		"\u0000\u0000\u0014\u0120\u0001\u0000\u0000\u0000\u0016\u0138\u0001\u0000"+
		"\u0000\u0000\u0018\u013a\u0001\u0000\u0000\u0000\u001a\u0142\u0001\u0000"+
		"\u0000\u0000\u001c\u0146\u0001\u0000\u0000\u0000\u001e\u015d\u0001\u0000"+
		"\u0000\u0000 \u0171\u0001\u0000\u0000\u0000\"\u0173\u0001\u0000\u0000"+
		"\u0000$\u0182\u0001\u0000\u0000\u0000&\u01a8\u0001\u0000\u0000\u0000("+
		"\u01af\u0001\u0000\u0000\u0000*\u01b9\u0001\u0000\u0000\u0000,\u01bb\u0001"+
		"\u0000\u0000\u0000.\u01cf\u0001\u0000\u0000\u00000\u01d5\u0001\u0000\u0000"+
		"\u00002\u01f0\u0001\u0000\u0000\u00004\u01f2\u0001\u0000\u0000\u00006"+
		"\u01f4\u0001\u0000\u0000\u00008\u01f7\u0001\u0000\u0000\u0000:\u01fb\u0001"+
		"\u0000\u0000\u0000<\u0209\u0001\u0000\u0000\u0000>\u0218\u0001\u0000\u0000"+
		"\u0000@\u0241\u0001\u0000\u0000\u0000B\u0243\u0001\u0000\u0000\u0000D"+
		"\u025f\u0001\u0000\u0000\u0000F\u0279\u0001\u0000\u0000\u0000H\u027b\u0001"+
		"\u0000\u0000\u0000J\u027d\u0001\u0000\u0000\u0000L\u027f\u0001\u0000\u0000"+
		"\u0000N\u0281\u0001\u0000\u0000\u0000P\u0283\u0001\u0000\u0000\u0000R"+
		"\u0285\u0001\u0000\u0000\u0000T\u0287\u0001\u0000\u0000\u0000V\u0289\u0001"+
		"\u0000\u0000\u0000X\u028b\u0001\u0000\u0000\u0000Z\u028d\u0001\u0000\u0000"+
		"\u0000\\\u028f\u0001\u0000\u0000\u0000^\u0291\u0001\u0000\u0000\u0000"+
		"`\u0293\u0001\u0000\u0000\u0000b\u0295\u0001\u0000\u0000\u0000d\u02a6"+
		"\u0001\u0000\u0000\u0000f\u02a8\u0001\u0000\u0000\u0000h\u02c0\u0001\u0000"+
		"\u0000\u0000j\u02ca\u0001\u0000\u0000\u0000l\u02d9\u0001\u0000\u0000\u0000"+
		"n\u02e9\u0001\u0000\u0000\u0000p\u02f5\u0001\u0000\u0000\u0000r\u0300"+
		"\u0001\u0000\u0000\u0000t\u0302\u0001\u0000\u0000\u0000v\u0306\u0001\u0000"+
		"\u0000\u0000xz\u0003\u0002\u0001\u0000yx\u0001\u0000\u0000\u0000z}\u0001"+
		"\u0000\u0000\u0000{y\u0001\u0000\u0000\u0000{|\u0001\u0000\u0000\u0000"+
		"|~\u0001\u0000\u0000\u0000}{\u0001\u0000\u0000\u0000~\u007f\u0005\u0000"+
		"\u0000\u0001\u007f\u0001\u0001\u0000\u0000\u0000\u0080\u008a\u0003\u0004"+
		"\u0002\u0000\u0081\u008a\u0003*\u0015\u0000\u0082\u008a\u0003B!\u0000"+
		"\u0083\u008a\u0003D\"\u0000\u0084\u008a\u0003@ \u0000\u0085\u008a\u0003"+
		"\u001a\r\u0000\u0086\u008a\u0003h4\u0000\u0087\u0088\u0005\u0001\u0000"+
		"\u0000\u0088\u008a\u0003\u0006\u0003\u0000\u0089\u0080\u0001\u0000\u0000"+
		"\u0000\u0089\u0081\u0001\u0000\u0000\u0000\u0089\u0082\u0001\u0000\u0000"+
		"\u0000\u0089\u0083\u0001\u0000\u0000\u0000\u0089\u0084\u0001\u0000\u0000"+
		"\u0000\u0089\u0085\u0001\u0000\u0000\u0000\u0089\u0086\u0001\u0000\u0000"+
		"\u0000\u0089\u0087\u0001\u0000\u0000\u0000\u008a\u0003\u0001\u0000\u0000"+
		"\u0000\u008b\u0090\u0003\u000e\u0007\u0000\u008c\u0091\u0003P(\u0000\u008d"+
		"\u0091\u0003Z-\u0000\u008e\u0091\u0003^/\u0000\u008f\u0091\u0003\\.\u0000"+
		"\u0090\u008c\u0001\u0000\u0000\u0000\u0090\u008d\u0001\u0000\u0000\u0000"+
		"\u0090\u008e\u0001\u0000\u0000\u0000\u0090\u008f\u0001\u0000\u0000\u0000"+
		"\u0091\u0092\u0001\u0000\u0000\u0000\u0092\u0093\u0003\u0006\u0003\u0000"+
		"\u0093\u0005\u0001\u0000\u0000\u0000\u0094\u0095\u0006\u0003\uffff\uffff"+
		"\u0000\u0095\u009b\u0003\b\u0004\u0000\u0096\u0097\u0003F#\u0000\u0097"+
		"\u0098\u0003\u0006\u0003\u0000\u0098\u009a\u0001\u0000\u0000\u0000\u0099"+
		"\u0096\u0001\u0000\u0000\u0000\u009a\u009d\u0001\u0000\u0000\u0000\u009b"+
		"\u0099\u0001\u0000\u0000\u0000\u009b\u009c\u0001\u0000\u0000\u0000\u009c"+
		"\u00b2\u0001\u0000\u0000\u0000\u009d\u009b\u0001\u0000\u0000\u0000\u009e"+
		"\u009f\u0003R)\u0000\u009f\u00a0\u0003\u0006\u0003\u0005\u00a0\u00b2\u0001"+
		"\u0000\u0000\u0000\u00a1\u00a2\u0005\u0002\u0000\u0000\u00a2\u00a3\u0003"+
		"\u0006\u0003\u0000\u00a3\u00a5\u0005\u0003\u0000\u0000\u00a4\u00a6\u0003"+
		"\u001c\u000e\u0000\u00a5\u00a4\u0001\u0000\u0000\u0000\u00a6\u00a7\u0001"+
		"\u0000\u0000\u0000\u00a7\u00a5\u0001\u0000\u0000\u0000\u00a7\u00a8\u0001"+
		"\u0000\u0000\u0000\u00a8\u00a9\u0001\u0000\u0000\u0000\u00a9\u00aa\u0005"+
		"\u0004\u0000\u0000\u00aa\u00b2\u0001\u0000\u0000\u0000\u00ab\u00ac\u0005"+
		"\u0005\u0000\u0000\u00ac\u00ad\u0003>\u001f\u0000\u00ad\u00ae\u0005\u0004"+
		"\u0000\u0000\u00ae\u00b2\u0001\u0000\u0000\u0000\u00af\u00b0\u0005\u0006"+
		"\u0000\u0000\u00b0\u00b2\u0003\u0006\u0003\u0002\u00b1\u0094\u0001\u0000"+
		"\u0000\u0000\u00b1\u009e\u0001\u0000\u0000\u0000\u00b1\u00a1\u0001\u0000"+
		"\u0000\u0000\u00b1\u00ab\u0001\u0000\u0000\u0000\u00b1\u00af\u0001\u0000"+
		"\u0000\u0000\u00b2\u00bf\u0001\u0000\u0000\u0000\u00b3\u00b4\n\u0007\u0000"+
		"\u0000\u00b4\u00b5\u0003F#\u0000\u00b5\u00b6\u0003\u0006\u0003\b\u00b6"+
		"\u00be\u0001\u0000\u0000\u0000\u00b7\u00b8\n\u0006\u0000\u0000\u00b8\u00b9"+
		"\u0003F#\u0000\u00b9\u00ba\u0003\u0006\u0003\u0006\u00ba\u00be\u0001\u0000"+
		"\u0000\u0000\u00bb\u00bc\n\u0001\u0000\u0000\u00bc\u00be\u0005\u0007\u0000"+
		"\u0000\u00bd\u00b3\u0001\u0000\u0000\u0000\u00bd\u00b7\u0001\u0000\u0000"+
		"\u0000\u00bd\u00bb\u0001\u0000\u0000\u0000\u00be\u00c1\u0001\u0000\u0000"+
		"\u0000\u00bf\u00bd\u0001\u0000\u0000\u0000\u00bf\u00c0\u0001\u0000\u0000"+
		"\u0000\u00c0\u0007\u0001\u0000\u0000\u0000\u00c1\u00bf\u0001\u0000\u0000"+
		"\u0000\u00c2\u00ce\u0003\f\u0006\u0000\u00c3\u00ce\u0003\u000e\u0007\u0000"+
		"\u00c4\u00ce\u0003\u0012\t\u0000\u00c5\u00ce\u0003(\u0014\u0000\u00c6"+
		"\u00ce\u0003\u0014\n\u0000\u00c7\u00ce\u0003f3\u0000\u00c8\u00c9\u0005"+
		"\b\u0000\u0000\u00c9\u00ca\u0003\u0006\u0003\u0000\u00ca\u00cb\u0005\t"+
		"\u0000\u0000\u00cb\u00ce\u0001\u0000\u0000\u0000\u00cc\u00ce\u0003j5\u0000"+
		"\u00cd\u00c2\u0001\u0000\u0000\u0000\u00cd\u00c3\u0001\u0000\u0000\u0000"+
		"\u00cd\u00c4\u0001\u0000\u0000\u0000\u00cd\u00c5\u0001\u0000\u0000\u0000"+
		"\u00cd\u00c6\u0001\u0000\u0000\u0000\u00cd\u00c7\u0001\u0000\u0000\u0000"+
		"\u00cd\u00c8\u0001\u0000\u0000\u0000\u00cd\u00cc\u0001\u0000\u0000\u0000"+
		"\u00ce\u00d2\u0001\u0000\u0000\u0000\u00cf\u00d1\u0003\n\u0005\u0000\u00d0"+
		"\u00cf\u0001\u0000\u0000\u0000\u00d1\u00d4\u0001\u0000\u0000\u0000\u00d2"+
		"\u00d0\u0001\u0000\u0000\u0000\u00d2\u00d3\u0001\u0000\u0000\u0000\u00d3"+
		"\t\u0001\u0000\u0000\u0000\u00d4\u00d2\u0001\u0000\u0000\u0000\u00d5\u00d6"+
		"\u0005\n\u0000\u0000\u00d6\u00d7\u0005u\u0000\u0000\u00d7\u00d9\u0005"+
		"\b\u0000\u0000\u00d8\u00da\u0003d2\u0000\u00d9\u00d8\u0001\u0000\u0000"+
		"\u0000\u00d9\u00da\u0001\u0000\u0000\u0000\u00da\u00db\u0001\u0000\u0000"+
		"\u0000\u00db\u00e3\u0005\t\u0000\u0000\u00dc\u00dd\u0005\u000b\u0000\u0000"+
		"\u00dd\u00e3\u0005u\u0000\u0000\u00de\u00df\u0005\f\u0000\u0000\u00df"+
		"\u00e0\u0003\u0006\u0003\u0000\u00e0\u00e1\u0005\r\u0000\u0000\u00e1\u00e3"+
		"\u0001\u0000\u0000\u0000\u00e2\u00d5\u0001\u0000\u0000\u0000\u00e2\u00dc"+
		"\u0001\u0000\u0000\u0000\u00e2\u00de\u0001\u0000\u0000\u0000\u00e3\u000b"+
		"\u0001\u0000\u0000\u0000\u00e4\u00e5\u0007\u0000\u0000\u0000\u00e5\r\u0001"+
		"\u0000\u0000\u0000\u00e6\u00e7\u0006\u0007\uffff\uffff\u0000\u00e7\u00e8"+
		"\u0005u\u0000\u0000\u00e8\u00f5\u0001\u0000\u0000\u0000\u00e9\u00ea\n"+
		"\u0003\u0000\u0000\u00ea\u00f4\u0003\u0010\b\u0000\u00eb\u00ec\n\u0002"+
		"\u0000\u0000\u00ec\u00ed\u0005\f\u0000\u0000\u00ed\u00ee\u0003\u0006\u0003"+
		"\u0000\u00ee\u00ef\u0005\r\u0000\u0000\u00ef\u00f4\u0001\u0000\u0000\u0000"+
		"\u00f0\u00f1\n\u0001\u0000\u0000\u00f1\u00f2\u0005\u000b\u0000\u0000\u00f2"+
		"\u00f4\u0005u\u0000\u0000\u00f3\u00e9\u0001\u0000\u0000\u0000\u00f3\u00eb"+
		"\u0001\u0000\u0000\u0000\u00f3\u00f0\u0001\u0000\u0000\u0000\u00f4\u00f7"+
		"\u0001\u0000\u0000\u0000\u00f5\u00f3\u0001\u0000\u0000\u0000\u00f5\u00f6"+
		"\u0001\u0000\u0000\u0000\u00f6\u000f\u0001\u0000\u0000\u0000\u00f7\u00f5"+
		"\u0001\u0000\u0000\u0000\u00f8\u00f9\u0005\u000e\u0000\u0000\u00f9\u00ff"+
		"\u0005u\u0000\u0000\u00fa\u00fb\u0005\u000f\u0000\u0000\u00fb\u00fc\u0003"+
		"\u0006\u0003\u0000\u00fc\u00fd\u0005\r\u0000\u0000\u00fd\u00ff\u0001\u0000"+
		"\u0000\u0000\u00fe\u00f8\u0001\u0000\u0000\u0000\u00fe\u00fa\u0001\u0000"+
		"\u0000\u0000\u00ff\u0011\u0001\u0000\u0000\u0000\u0100\u0101\u0003\u000e"+
		"\u0007\u0000\u0101\u010a\u0005\b\u0000\u0000\u0102\u0107\u0003\u0006\u0003"+
		"\u0000\u0103\u0104\u0005\u0010\u0000\u0000\u0104\u0106\u0003\u0006\u0003"+
		"\u0000\u0105\u0103\u0001\u0000\u0000\u0000\u0106\u0109\u0001\u0000\u0000"+
		"\u0000\u0107\u0105\u0001\u0000\u0000\u0000\u0107\u0108\u0001\u0000\u0000"+
		"\u0000\u0108\u010b\u0001\u0000\u0000\u0000\u0109\u0107\u0001\u0000\u0000"+
		"\u0000\u010a\u0102\u0001\u0000\u0000\u0000\u010a\u010b\u0001\u0000\u0000"+
		"\u0000\u010b\u010c\u0001\u0000\u0000\u0000\u010c\u010d\u0005\t\u0000\u0000"+
		"\u010d\u011f\u0001\u0000\u0000\u0000\u010e\u010f\u0003\u000e\u0007\u0000"+
		"\u010f\u0110\u0005\n\u0000\u0000\u0110\u0111\u0005u\u0000\u0000\u0111"+
		"\u011a\u0005\b\u0000\u0000\u0112\u0117\u0003\u0006\u0003\u0000\u0113\u0114"+
		"\u0005\u0010\u0000\u0000\u0114\u0116\u0003\u0006\u0003\u0000\u0115\u0113"+
		"\u0001\u0000\u0000\u0000\u0116\u0119\u0001\u0000\u0000\u0000\u0117\u0115"+
		"\u0001\u0000\u0000\u0000\u0117\u0118\u0001\u0000\u0000\u0000\u0118\u011b"+
		"\u0001\u0000\u0000\u0000\u0119\u0117\u0001\u0000\u0000\u0000\u011a\u0112"+
		"\u0001\u0000\u0000\u0000\u011a\u011b\u0001\u0000\u0000\u0000\u011b\u011c"+
		"\u0001\u0000\u0000\u0000\u011c\u011d\u0005\t\u0000\u0000\u011d\u011f\u0001"+
		"\u0000\u0000\u0000\u011e\u0100\u0001\u0000\u0000\u0000\u011e\u010e\u0001"+
		"\u0000\u0000\u0000\u011f\u0013\u0001\u0000\u0000\u0000\u0120\u0129\u0005"+
		"\u0011\u0000\u0000\u0121\u0126\u0003$\u0012\u0000\u0122\u0123\u0005\u0010"+
		"\u0000\u0000\u0123\u0125\u0003$\u0012\u0000\u0124\u0122\u0001\u0000\u0000"+
		"\u0000\u0125\u0128\u0001\u0000\u0000\u0000\u0126\u0124\u0001\u0000\u0000"+
		"\u0000\u0126\u0127\u0001\u0000\u0000\u0000\u0127\u012a\u0001\u0000\u0000"+
		"\u0000\u0128\u0126\u0001\u0000\u0000\u0000\u0129\u0121\u0001\u0000\u0000"+
		"\u0000\u0129\u012a\u0001\u0000\u0000\u0000\u012a\u012d\u0001\u0000\u0000"+
		"\u0000\u012b\u012c\u0005\u0010\u0000\u0000\u012c\u012e\u0003\u0016\u000b"+
		"\u0000\u012d\u012b\u0001\u0000\u0000\u0000\u012d\u012e\u0001\u0000\u0000"+
		"\u0000\u012e\u012f\u0001\u0000\u0000\u0000\u012f\u0130\u0005\u0012\u0000"+
		"\u0000\u0130\u0015\u0001\u0000\u0000\u0000\u0131\u0132\u0005\u0013\u0000"+
		"\u0000\u0132\u0133\u0005\u0014\u0000\u0000\u0133\u0139\u0003\u0006\u0003"+
		"\u0000\u0134\u0135\u0005\u0011\u0000\u0000\u0135\u0136\u0003\u0018\f\u0000"+
		"\u0136\u0137\u0005\u0012\u0000\u0000\u0137\u0139\u0001\u0000\u0000\u0000"+
		"\u0138\u0131\u0001\u0000\u0000\u0000\u0138\u0134\u0001\u0000\u0000\u0000"+
		"\u0139\u0017\u0001\u0000\u0000\u0000\u013a\u013f\u0003\"\u0011\u0000\u013b"+
		"\u013c\u0005\u0010\u0000\u0000\u013c\u013e\u0003\"\u0011\u0000\u013d\u013b"+
		"\u0001\u0000\u0000\u0000\u013e\u0141\u0001\u0000\u0000\u0000\u013f\u013d"+
		"\u0001\u0000\u0000\u0000\u013f\u0140\u0001\u0000\u0000\u0000\u0140\u0019"+
		"\u0001\u0000\u0000\u0000\u0141\u013f\u0001\u0000\u0000\u0000\u0142\u0143"+
		"\u0005\u0015\u0000\u0000\u0143\u0144\u0005u\u0000\u0000\u0144\u0145\u0005"+
		"\u0015\u0000\u0000\u0145\u001b\u0001\u0000\u0000\u0000\u0146\u0149\u0003"+
		"\u001e\u000f\u0000\u0147\u0148\u0005\u0016\u0000\u0000\u0148\u014a\u0003"+
		"\u0006\u0003\u0000\u0149\u0147\u0001\u0000\u0000\u0000\u0149\u014a\u0001"+
		"\u0000\u0000\u0000\u014a\u014b\u0001\u0000\u0000\u0000\u014b\u014c\u0005"+
		"\u0017\u0000\u0000\u014c\u014d\u0003\u0006\u0003\u0000\u014d\u001d\u0001"+
		"\u0000\u0000\u0000\u014e\u014f\u0006\u000f\uffff\uffff\u0000\u014f\u015e"+
		"\u0003\f\u0006\u0000\u0150\u015e\u0003\u0014\n\u0000\u0151\u015e\u0005"+
		"u\u0000\u0000\u0152\u0153\u0005\u0011\u0000\u0000\u0153\u0158\u0003 \u0010"+
		"\u0000\u0154\u0155\u0005\u0010\u0000\u0000\u0155\u0157\u0003 \u0010\u0000"+
		"\u0156\u0154\u0001\u0000\u0000\u0000\u0157\u015a\u0001\u0000\u0000\u0000"+
		"\u0158\u0156\u0001\u0000\u0000\u0000\u0158\u0159\u0001\u0000\u0000\u0000"+
		"\u0159\u015b\u0001\u0000\u0000\u0000\u015a\u0158\u0001\u0000\u0000\u0000"+
		"\u015b\u015c\u0005\u0012\u0000\u0000\u015c\u015e\u0001\u0000\u0000\u0000"+
		"\u015d\u014e\u0001\u0000\u0000\u0000\u015d\u0150\u0001\u0000\u0000\u0000"+
		"\u015d\u0151\u0001\u0000\u0000\u0000\u015d\u0152\u0001\u0000\u0000\u0000"+
		"\u015e\u0164\u0001\u0000\u0000\u0000\u015f\u0160\n\u0002\u0000\u0000\u0160"+
		"\u0161\u0005\u0018\u0000\u0000\u0161\u0163\u0003\u001e\u000f\u0003\u0162"+
		"\u015f\u0001\u0000\u0000\u0000\u0163\u0166\u0001\u0000\u0000\u0000\u0164"+
		"\u0162\u0001\u0000\u0000\u0000\u0164\u0165\u0001\u0000\u0000\u0000\u0165"+
		"\u001f\u0001\u0000\u0000\u0000\u0166\u0164\u0001\u0000\u0000\u0000\u0167"+
		"\u0168\u0005u\u0000\u0000\u0168\u0169\u0005\u0014\u0000\u0000\u0169\u0172"+
		"\u0003\u001e\u000f\u0000\u016a\u016b\u0005\f\u0000\u0000\u016b\u016c\u0003"+
		"\u0006\u0003\u0000\u016c\u016d\u0005\r\u0000\u0000\u016d\u016e\u0005\u0014"+
		"\u0000\u0000\u016e\u016f\u0003\u001e\u000f\u0000\u016f\u0172\u0001\u0000"+
		"\u0000\u0000\u0170\u0172\u0003\u001e\u000f\u0000\u0171\u0167\u0001\u0000"+
		"\u0000\u0000\u0171\u016a\u0001\u0000\u0000\u0000\u0171\u0170\u0001\u0000"+
		"\u0000\u0000\u0172!\u0001\u0000\u0000\u0000\u0173\u0174\u0007\u0001\u0000"+
		"\u0000\u0174#\u0001\u0000\u0000\u0000\u0175\u0176\u0005u\u0000\u0000\u0176"+
		"\u0177\u0005\u0014\u0000\u0000\u0177\u0183\u0003\u0006\u0003\u0000\u0178"+
		"\u0179\u0005\f\u0000\u0000\u0179\u017a\u0003\u0006\u0003\u0000\u017a\u017b"+
		"\u0005\r\u0000\u0000\u017b\u017c\u0005\u0014\u0000\u0000\u017c\u017d\u0003"+
		"\u0006\u0003\u0000\u017d\u0183\u0001\u0000\u0000\u0000\u017e\u0183\u0003"+
		"\u0006\u0003\u0000\u017f\u0180\u0005u\u0000\u0000\u0180\u0181\u0005\n"+
		"\u0000\u0000\u0181\u0183\u0003f3\u0000\u0182\u0175\u0001\u0000\u0000\u0000"+
		"\u0182\u0178\u0001\u0000\u0000\u0000\u0182\u017e\u0001\u0000\u0000\u0000"+
		"\u0182\u017f\u0001\u0000\u0000\u0000\u0183%\u0001\u0000\u0000\u0000\u0184"+
		"\u0185\u0003\u0006\u0003\u0000\u0185\u0186\u0003L&\u0000\u0186\u0187\u0003"+
		"\u0006\u0003\u0000\u0187\u01a9\u0001\u0000\u0000\u0000\u0188\u0189\u0003"+
		"\u0006\u0003\u0000\u0189\u018a\u0003N\'\u0000\u018a\u018b\u0003\u0006"+
		"\u0003\u0000\u018b\u01a9\u0001\u0000\u0000\u0000\u018c\u018d\u0003\u0006"+
		"\u0003\u0000\u018d\u018e\u0003J%\u0000\u018e\u018f\u0003\u0006\u0003\u0000"+
		"\u018f\u01a9\u0001\u0000\u0000\u0000\u0190\u0191\u0003\u0006\u0003\u0000"+
		"\u0191\u0192\u0003H$\u0000\u0192\u0193\u0003\u0006\u0003\u0000\u0193\u01a9"+
		"\u0001\u0000\u0000\u0000\u0194\u0195\u0003\u0006\u0003\u0000\u0195\u0196"+
		"\u0003T*\u0000\u0196\u0197\u0003\u0006\u0003\u0000\u0197\u01a9\u0001\u0000"+
		"\u0000\u0000\u0198\u0199\u0003\u0006\u0003\u0000\u0199\u019a\u0003X,\u0000"+
		"\u019a\u019b\u0003\u0006\u0003\u0000\u019b\u01a9\u0001\u0000\u0000\u0000"+
		"\u019c\u019d\u0003\u0006\u0003\u0000\u019d\u019e\u0003\\.\u0000\u019e"+
		"\u019f\u0003\u0006\u0003\u0000\u019f\u01a9\u0001\u0000\u0000\u0000\u01a0"+
		"\u01a1\u0003\u0006\u0003\u0000\u01a1\u01a2\u0003^/\u0000\u01a2\u01a3\u0003"+
		"\u0006\u0003\u0000\u01a3\u01a9\u0001\u0000\u0000\u0000\u01a4\u01a5\u0003"+
		"\u0006\u0003\u0000\u01a5\u01a6\u0003Z-\u0000\u01a6\u01a7\u0003\u0006\u0003"+
		"\u0000\u01a7\u01a9\u0001\u0000\u0000\u0000\u01a8\u0184\u0001\u0000\u0000"+
		"\u0000\u01a8\u0188\u0001\u0000\u0000\u0000\u01a8\u018c\u0001\u0000\u0000"+
		"\u0000\u01a8\u0190\u0001\u0000\u0000\u0000\u01a8\u0194\u0001\u0000\u0000"+
		"\u0000\u01a8\u0198\u0001\u0000\u0000\u0000\u01a8\u019c\u0001\u0000\u0000"+
		"\u0000\u01a8\u01a0\u0001\u0000\u0000\u0000\u01a8\u01a4\u0001\u0000\u0000"+
		"\u0000\u01a9\'\u0001\u0000\u0000\u0000\u01aa\u01ab\u0003R)\u0000\u01ab"+
		"\u01ac\u0003\u0006\u0003\u0000\u01ac\u01b0\u0001\u0000\u0000\u0000\u01ad"+
		"\u01ae\u0005)\u0000\u0000\u01ae\u01b0\u0003\u0006\u0003\u0000\u01af\u01aa"+
		"\u0001\u0000\u0000\u0000\u01af\u01ad\u0001\u0000\u0000\u0000\u01b0)\u0001"+
		"\u0000\u0000\u0000\u01b1\u01ba\u0003,\u0016\u0000\u01b2\u01ba\u0003.\u0017"+
		"\u0000\u01b3\u01ba\u00030\u0018\u0000\u01b4\u01ba\u00032\u0019\u0000\u01b5"+
		"\u01ba\u00034\u001a\u0000\u01b6\u01ba\u00036\u001b\u0000\u01b7\u01ba\u0003"+
		"8\u001c\u0000\u01b8\u01ba\u0003:\u001d\u0000\u01b9\u01b1\u0001\u0000\u0000"+
		"\u0000\u01b9\u01b2\u0001\u0000\u0000\u0000\u01b9\u01b3\u0001\u0000\u0000"+
		"\u0000\u01b9\u01b4\u0001\u0000\u0000\u0000\u01b9\u01b5\u0001\u0000\u0000"+
		"\u0000\u01b9\u01b6\u0001\u0000\u0000\u0000\u01b9\u01b7\u0001\u0000\u0000"+
		"\u0000\u01b9\u01b8\u0001\u0000\u0000\u0000\u01ba+\u0001\u0000\u0000\u0000"+
		"\u01bb\u01bc\u0005*\u0000\u0000\u01bc\u01bd\u0003\u0006\u0003\u0000\u01bd"+
		"\u01be\u0005+\u0000\u0000\u01be\u01c6\u0003>\u001f\u0000\u01bf\u01c0\u0005"+
		",\u0000\u0000\u01c0\u01c1\u0003\u0006\u0003\u0000\u01c1\u01c2\u0005+\u0000"+
		"\u0000\u01c2\u01c3\u0003>\u001f\u0000\u01c3\u01c5\u0001\u0000\u0000\u0000"+
		"\u01c4\u01bf\u0001\u0000\u0000\u0000\u01c5\u01c8\u0001\u0000\u0000\u0000"+
		"\u01c6\u01c4\u0001\u0000\u0000\u0000\u01c6\u01c7\u0001\u0000\u0000\u0000"+
		"\u01c7\u01cb\u0001\u0000\u0000\u0000\u01c8\u01c6\u0001\u0000\u0000\u0000"+
		"\u01c9\u01ca\u0005-\u0000\u0000\u01ca\u01cc\u0003>\u001f\u0000\u01cb\u01c9"+
		"\u0001\u0000\u0000\u0000\u01cb\u01cc\u0001\u0000\u0000\u0000\u01cc\u01cd"+
		"\u0001\u0000\u0000\u0000\u01cd\u01ce\u0005\u0004\u0000\u0000\u01ce-\u0001"+
		"\u0000\u0000\u0000\u01cf\u01d0\u0005.\u0000\u0000\u01d0\u01d1\u0003\u0006"+
		"\u0003\u0000\u01d1\u01d2\u0005/\u0000\u0000\u01d2\u01d3\u0003>\u001f\u0000"+
		"\u01d3\u01d4\u0005\u0004\u0000\u0000\u01d4/\u0001\u0000\u0000\u0000\u01d5"+
		"\u01d6\u00050\u0000\u0000\u01d6\u01d7\u0003>\u001f\u0000\u01d7\u01d8\u0005"+
		"1\u0000\u0000\u01d8\u01d9\u0003\u0006\u0003\u0000\u01d91\u0001\u0000\u0000"+
		"\u0000\u01da\u01db\u00052\u0000\u0000\u01db\u01dc\u0005u\u0000\u0000\u01dc"+
		"\u01dd\u0005\u0014\u0000\u0000\u01dd\u01de\u0003\u0006\u0003\u0000\u01de"+
		"\u01df\u0005\u0010\u0000\u0000\u01df\u01e2\u0003\u0006\u0003\u0000\u01e0"+
		"\u01e1\u0005\u0010\u0000\u0000\u01e1\u01e3\u0003\u0006\u0003\u0000\u01e2"+
		"\u01e0\u0001\u0000\u0000\u0000\u01e2\u01e3\u0001\u0000\u0000\u0000\u01e3"+
		"\u01e4\u0001\u0000\u0000\u0000\u01e4\u01e5\u0005/\u0000\u0000\u01e5\u01e6"+
		"\u0003>\u001f\u0000\u01e6\u01e7\u0005\u0004\u0000\u0000\u01e7\u01f1\u0001"+
		"\u0000\u0000\u0000\u01e8\u01e9\u00052\u0000\u0000\u01e9\u01ea\u0003b1"+
		"\u0000\u01ea\u01eb\u00053\u0000\u0000\u01eb\u01ec\u0003d2\u0000\u01ec"+
		"\u01ed\u0005/\u0000\u0000\u01ed\u01ee\u0003>\u001f\u0000\u01ee\u01ef\u0005"+
		"\u0004\u0000\u0000\u01ef\u01f1\u0001\u0000\u0000\u0000\u01f0\u01da\u0001"+
		"\u0000\u0000\u0000\u01f0\u01e8\u0001\u0000\u0000\u0000\u01f13\u0001\u0000"+
		"\u0000\u0000\u01f2\u01f3\u00054\u0000\u0000\u01f35\u0001\u0000\u0000\u0000"+
		"\u01f4\u01f5\u00055\u0000\u0000\u01f5\u01f6\u0005u\u0000\u0000\u01f67"+
		"\u0001\u0000\u0000\u0000\u01f7\u01f8\u00056\u0000\u0000\u01f8\u01f9\u0005"+
		"\u000b\u0000\u0000\u01f9\u01fa\u0007\u0002\u0000\u0000\u01fa9\u0001\u0000"+
		"\u0000\u0000\u01fb\u01fd\u0007\u0003\u0000\u0000\u01fc\u01fe\u0007\u0004"+
		"\u0000\u0000\u01fd\u01fc\u0001\u0000\u0000\u0000\u01fd\u01fe\u0001\u0000"+
		"\u0000\u0000\u01fe\u0200\u0001\u0000\u0000\u0000\u01ff\u0201\u0005u\u0000"+
		"\u0000\u0200\u01ff\u0001\u0000\u0000\u0000\u0200\u0201\u0001\u0000\u0000"+
		"\u0000\u0201\u0202\u0001\u0000\u0000\u0000\u0202\u0205\u0005\b\u0000\u0000"+
		"\u0203\u0206\u0003d2\u0000\u0204\u0206\u0003<\u001e\u0000\u0205\u0203"+
		"\u0001\u0000\u0000\u0000\u0205\u0204\u0001\u0000\u0000\u0000\u0206\u0207"+
		"\u0001\u0000\u0000\u0000\u0207\u0208\u0005\t\u0000\u0000\u0208;\u0001"+
		"\u0000\u0000\u0000\u0209\u020a\u0005u\u0000\u0000\u020a\u020b\u0005\u0014"+
		"\u0000\u0000\u020b\u0212\u0003\u0006\u0003\u0000\u020c\u020d\u0005\u0010"+
		"\u0000\u0000\u020d\u020e\u0005u\u0000\u0000\u020e\u020f\u0005\u0014\u0000"+
		"\u0000\u020f\u0211\u0003\u0006\u0003\u0000\u0210\u020c\u0001\u0000\u0000"+
		"\u0000\u0211\u0214\u0001\u0000\u0000\u0000\u0212\u0210\u0001\u0000\u0000"+
		"\u0000\u0212\u0213\u0001\u0000\u0000\u0000\u0213=\u0001\u0000\u0000\u0000"+
		"\u0214\u0212\u0001\u0000\u0000\u0000\u0215\u0217\u0003\u0002\u0001\u0000"+
		"\u0216\u0215\u0001\u0000\u0000\u0000\u0217\u021a\u0001\u0000\u0000\u0000"+
		"\u0218\u0216\u0001\u0000\u0000\u0000\u0218\u0219\u0001\u0000\u0000\u0000"+
		"\u0219?\u0001\u0000\u0000\u0000\u021a\u0218\u0001\u0000\u0000\u0000\u021b"+
		"\u021d\u0005@\u0000\u0000\u021c\u021e\u0003l6\u0000\u021d\u021c\u0001"+
		"\u0000\u0000\u0000\u021d\u021e\u0001\u0000\u0000\u0000\u021e\u021f\u0001"+
		"\u0000\u0000\u0000\u021f\u0222\u0005u\u0000\u0000\u0220\u0221\u0005\u0014"+
		"\u0000\u0000\u0221\u0223\u0003\u0006\u0003\u0000\u0222\u0220\u0001\u0000"+
		"\u0000\u0000\u0222\u0223\u0001\u0000\u0000\u0000\u0223\u0242\u0001\u0000"+
		"\u0000\u0000\u0224\u0225\u0005@\u0000\u0000\u0225\u022a\u0005u\u0000\u0000"+
		"\u0226\u0227\u0005\u0010\u0000\u0000\u0227\u0229\u0005u\u0000\u0000\u0228"+
		"\u0226\u0001\u0000\u0000\u0000\u0229\u022c\u0001\u0000\u0000\u0000\u022a"+
		"\u0228\u0001\u0000\u0000\u0000\u022a\u022b\u0001\u0000\u0000\u0000\u022b"+
		"\u022d\u0001\u0000\u0000\u0000\u022c\u022a\u0001\u0000\u0000\u0000\u022d"+
		"\u022e\u0005\u0014\u0000\u0000\u022e\u0242\u0003d2\u0000\u022f\u0230\u0005"+
		"@\u0000\u0000\u0230\u0231\u0005A\u0000\u0000\u0231\u0232\u0005u\u0000"+
		"\u0000\u0232\u023b\u0005\b\u0000\u0000\u0233\u0238\u0005u\u0000\u0000"+
		"\u0234\u0235\u0005\u0010\u0000\u0000\u0235\u0237\u0005u\u0000\u0000\u0236"+
		"\u0234\u0001\u0000\u0000\u0000\u0237\u023a\u0001\u0000\u0000\u0000\u0238"+
		"\u0236\u0001\u0000\u0000\u0000\u0238\u0239\u0001\u0000\u0000\u0000\u0239"+
		"\u023c\u0001\u0000\u0000\u0000\u023a\u0238\u0001\u0000\u0000\u0000\u023b"+
		"\u0233\u0001\u0000\u0000\u0000\u023b\u023c\u0001\u0000\u0000\u0000\u023c"+
		"\u023d\u0001\u0000\u0000\u0000\u023d\u023e\u0005\t\u0000\u0000\u023e\u023f"+
		"\u0003>\u001f\u0000\u023f\u0240\u0005\u0004\u0000\u0000\u0240\u0242\u0001"+
		"\u0000\u0000\u0000\u0241\u021b\u0001\u0000\u0000\u0000\u0241\u0224\u0001"+
		"\u0000\u0000\u0000\u0241\u022f\u0001\u0000\u0000\u0000\u0242A\u0001\u0000"+
		"\u0000\u0000\u0243\u0248\u0005A\u0000\u0000\u0244\u0245\u0005u\u0000\u0000"+
		"\u0245\u0249\u0005\u000b\u0000\u0000\u0246\u0247\u0005u\u0000\u0000\u0247"+
		"\u0249\u0005\n\u0000\u0000\u0248\u0244\u0001\u0000\u0000\u0000\u0248\u0246"+
		"\u0001\u0000\u0000\u0000\u0248\u0249\u0001\u0000\u0000\u0000\u0249\u024a"+
		"\u0001\u0000\u0000\u0000\u024a\u024b\u0005u\u0000\u0000\u024b\u0259\u0005"+
		"\b\u0000\u0000\u024c\u0251\u0005u\u0000\u0000\u024d\u024e\u0005\u0010"+
		"\u0000\u0000\u024e\u0250\u0005u\u0000\u0000\u024f\u024d\u0001\u0000\u0000"+
		"\u0000\u0250\u0253\u0001\u0000\u0000\u0000\u0251\u024f\u0001\u0000\u0000"+
		"\u0000\u0251\u0252\u0001\u0000\u0000\u0000\u0252\u0256\u0001\u0000\u0000"+
		"\u0000\u0253\u0251\u0001\u0000\u0000\u0000\u0254\u0255\u0005\u0010\u0000"+
		"\u0000\u0255\u0257\u0003V+\u0000\u0256\u0254\u0001\u0000\u0000\u0000\u0256"+
		"\u0257\u0001\u0000\u0000\u0000\u0257\u025a\u0001\u0000\u0000\u0000\u0258"+
		"\u025a\u0003V+\u0000\u0259\u024c\u0001\u0000\u0000\u0000\u0259\u0258\u0001"+
		"\u0000\u0000\u0000\u0259\u025a\u0001\u0000\u0000\u0000\u025a\u025b\u0001"+
		"\u0000\u0000\u0000\u025b\u025c\u0005\t\u0000\u0000\u025c\u025d\u0003>"+
		"\u001f\u0000\u025d\u025e\u0005\u0004\u0000\u0000\u025eC\u0001\u0000\u0000"+
		"\u0000\u025f\u0269\u0005B\u0000\u0000\u0260\u0265\u0003\u0006\u0003\u0000"+
		"\u0261\u0262\u0005\u0010\u0000\u0000\u0262\u0264\u0003\u0006\u0003\u0000"+
		"\u0263\u0261\u0001\u0000\u0000\u0000\u0264\u0267\u0001\u0000\u0000\u0000"+
		"\u0265\u0263\u0001\u0000\u0000\u0000\u0265\u0266\u0001\u0000\u0000\u0000"+
		"\u0266\u026a\u0001\u0000\u0000\u0000\u0267\u0265\u0001\u0000\u0000\u0000"+
		"\u0268\u026a\u0003\u0012\t\u0000\u0269\u0260\u0001\u0000\u0000\u0000\u0269"+
		"\u0268\u0001\u0000\u0000\u0000\u0269\u026a\u0001\u0000\u0000\u0000\u026a"+
		"E\u0001\u0000\u0000\u0000\u026b\u027a\u0003H$\u0000\u026c\u027a\u0003"+
		"J%\u0000\u026d\u027a\u0003L&\u0000\u026e\u027a\u0003N\'\u0000\u026f\u027a"+
		"\u0003P(\u0000\u0270\u027a\u0003R)\u0000\u0271\u027a\u0003T*\u0000\u0272"+
		"\u027a\u0003X,\u0000\u0273\u027a\u0003Z-\u0000\u0274\u027a\u0003\\.\u0000"+
		"\u0275\u027a\u0003^/\u0000\u0276\u027a\u0005\u0017\u0000\u0000\u0277\u027a"+
		"\u0003\u0010\b\u0000\u0278\u027a\u0003`0\u0000\u0279\u026b\u0001\u0000"+
		"\u0000\u0000\u0279\u026c\u0001\u0000\u0000\u0000\u0279\u026d\u0001\u0000"+
		"\u0000\u0000\u0279\u026e\u0001\u0000\u0000\u0000\u0279\u026f\u0001\u0000"+
		"\u0000\u0000\u0279\u0270\u0001\u0000\u0000\u0000\u0279\u0271\u0001\u0000"+
		"\u0000\u0000\u0279\u0272\u0001\u0000\u0000\u0000\u0279\u0273\u0001\u0000"+
		"\u0000\u0000\u0279\u0274\u0001\u0000\u0000\u0000\u0279\u0275\u0001\u0000"+
		"\u0000\u0000\u0279\u0276\u0001\u0000\u0000\u0000\u0279\u0277\u0001\u0000"+
		"\u0000\u0000\u0279\u0278\u0001\u0000\u0000\u0000\u027aG\u0001\u0000\u0000"+
		"\u0000\u027b\u027c\u0007\u0005\u0000\u0000\u027cI\u0001\u0000\u0000\u0000"+
		"\u027d\u027e\u0007\u0006\u0000\u0000\u027eK\u0001\u0000\u0000\u0000\u027f"+
		"\u0280\u0007\u0007\u0000\u0000\u0280M\u0001\u0000\u0000\u0000\u0281\u0282"+
		"\u0007\b\u0000\u0000\u0282O\u0001\u0000\u0000\u0000\u0283\u0284\u0007"+
		"\t\u0000\u0000\u0284Q\u0001\u0000\u0000\u0000\u0285\u0286\u0007\n\u0000"+
		"\u0000\u0286S\u0001\u0000\u0000\u0000\u0287\u0288\u0005`\u0000\u0000\u0288"+
		"U\u0001\u0000\u0000\u0000\u0289\u028a\u0005a\u0000\u0000\u028aW\u0001"+
		"\u0000\u0000\u0000\u028b\u028c\u0007\u000b\u0000\u0000\u028cY\u0001\u0000"+
		"\u0000\u0000\u028d\u028e\u0007\f\u0000\u0000\u028e[\u0001\u0000\u0000"+
		"\u0000\u028f\u0290\u0005h\u0000\u0000\u0290]\u0001\u0000\u0000\u0000\u0291"+
		"\u0292\u0007\r\u0000\u0000\u0292_\u0001\u0000\u0000\u0000\u0293\u0294"+
		"\u0005\u0007\u0000\u0000\u0294a\u0001\u0000\u0000\u0000\u0295\u029a\u0005"+
		"u\u0000\u0000\u0296\u0297\u0005\u0010\u0000\u0000\u0297\u0299\u0005u\u0000"+
		"\u0000\u0298\u0296\u0001\u0000\u0000\u0000\u0299\u029c\u0001\u0000\u0000"+
		"\u0000\u029a\u0298\u0001\u0000\u0000\u0000\u029a\u029b\u0001\u0000\u0000"+
		"\u0000\u029bc\u0001\u0000\u0000\u0000\u029c\u029a\u0001\u0000\u0000\u0000"+
		"\u029d\u02a2\u0003\u0006\u0003\u0000\u029e\u029f\u0005\u0010\u0000\u0000"+
		"\u029f\u02a1\u0003\u0006\u0003\u0000\u02a0\u029e\u0001\u0000\u0000\u0000"+
		"\u02a1\u02a4\u0001\u0000\u0000\u0000\u02a2\u02a0\u0001\u0000\u0000\u0000"+
		"\u02a2\u02a3\u0001\u0000\u0000\u0000\u02a3\u02a7\u0001\u0000\u0000\u0000"+
		"\u02a4\u02a2\u0001\u0000\u0000\u0000\u02a5\u02a7\u0003V+\u0000\u02a6\u029d"+
		"\u0001\u0000\u0000\u0000\u02a6\u02a5\u0001\u0000\u0000\u0000\u02a7e\u0001"+
		"\u0000\u0000\u0000\u02a8\u02ab\u0005A\u0000\u0000\u02a9\u02aa\u0005u\u0000"+
		"\u0000\u02aa\u02ac\u0005\n\u0000\u0000\u02ab\u02a9\u0001\u0000\u0000\u0000"+
		"\u02ab\u02ac\u0001\u0000\u0000\u0000\u02ac\u02ad\u0001\u0000\u0000\u0000"+
		"\u02ad\u02b6\u0005\b\u0000\u0000\u02ae\u02b3\u0005u\u0000\u0000\u02af"+
		"\u02b0\u0005\u0010\u0000\u0000\u02b0\u02b2\u0005u\u0000\u0000\u02b1\u02af"+
		"\u0001\u0000\u0000\u0000\u02b2\u02b5\u0001\u0000\u0000\u0000\u02b3\u02b1"+
		"\u0001\u0000\u0000\u0000\u02b3\u02b4\u0001\u0000\u0000\u0000\u02b4\u02b7"+
		"\u0001\u0000\u0000\u0000\u02b5\u02b3\u0001\u0000\u0000\u0000\u02b6\u02ae"+
		"\u0001\u0000\u0000\u0000\u02b6\u02b7\u0001\u0000\u0000\u0000\u02b7\u02b8"+
		"\u0001\u0000\u0000\u0000\u02b8\u02be\u0005\t\u0000\u0000\u02b9\u02ba\u0003"+
		">\u001f\u0000\u02ba\u02bb\u0005\u0004\u0000\u0000\u02bb\u02bf\u0001\u0000"+
		"\u0000\u0000\u02bc\u02bd\u0005\u0017\u0000\u0000\u02bd\u02bf\u0003\u0006"+
		"\u0003\u0000\u02be\u02b9\u0001\u0000\u0000\u0000\u02be\u02bc\u0001\u0000"+
		"\u0000\u0000\u02bfg\u0001\u0000\u0000\u0000\u02c0\u02c1\u0005k\u0000\u0000"+
		"\u02c1\u02c4\u0005\b\u0000\u0000\u02c2\u02c5\u0005)\u0000\u0000\u02c3"+
		"\u02c5\u0003\u0006\u0003\u0000\u02c4\u02c2\u0001\u0000\u0000\u0000\u02c4"+
		"\u02c3\u0001\u0000\u0000\u0000\u02c5\u02c6\u0001\u0000\u0000\u0000\u02c6"+
		"\u02c7\u0005\u0010\u0000\u0000\u02c7\u02c8\u0003\u0006\u0003\u0000\u02c8"+
		"\u02c9\u0005\t\u0000\u0000\u02c9i\u0001\u0000\u0000\u0000\u02ca\u02d3"+
		"\u0005\b\u0000\u0000\u02cb\u02d0\u0005u\u0000\u0000\u02cc\u02cd\u0005"+
		"\u0010\u0000\u0000\u02cd\u02cf\u0005u\u0000\u0000\u02ce\u02cc\u0001\u0000"+
		"\u0000\u0000\u02cf\u02d2\u0001\u0000\u0000\u0000\u02d0\u02ce\u0001\u0000"+
		"\u0000\u0000\u02d0\u02d1\u0001\u0000\u0000\u0000\u02d1\u02d4\u0001\u0000"+
		"\u0000\u0000\u02d2\u02d0\u0001\u0000\u0000\u0000\u02d3\u02cb\u0001\u0000"+
		"\u0000\u0000\u02d3\u02d4\u0001\u0000\u0000\u0000\u02d4\u02d5\u0001\u0000"+
		"\u0000\u0000\u02d5\u02d6\u0005\t\u0000\u0000\u02d6\u02d7\u0005\u0017\u0000"+
		"\u0000\u02d7\u02d8\u0003\u0006\u0003\u0000\u02d8k\u0001\u0000\u0000\u0000"+
		"\u02d9\u02da\u0005\n\u0000\u0000\u02da\u02db\u0003n7\u0000\u02dbm\u0001"+
		"\u0000\u0000\u0000\u02dc\u02dd\u00067\uffff\uffff\u0000\u02dd\u02ea\u0005"+
		"l\u0000\u0000\u02de\u02ea\u0005m\u0000\u0000\u02df\u02ea\u0005n\u0000"+
		"\u0000\u02e0\u02ea\u0005o\u0000\u0000\u02e1\u02ea\u0005A\u0000\u0000\u02e2"+
		"\u02ea\u0005p\u0000\u0000\u02e3\u02e4\u0005r\u0000\u0000\u02e4\u02e5\u0003"+
		"n7\u0000\u02e5\u02e6\u0005\u0010\u0000\u0000\u02e6\u02e7\u0003n7\u0000"+
		"\u02e7\u02e8\u0005I\u0000\u0000\u02e8\u02ea\u0001\u0000\u0000\u0000\u02e9"+
		"\u02dc\u0001\u0000\u0000\u0000\u02e9\u02de\u0001\u0000\u0000\u0000\u02e9"+
		"\u02df\u0001\u0000\u0000\u0000\u02e9\u02e0\u0001\u0000\u0000\u0000\u02e9"+
		"\u02e1\u0001\u0000\u0000\u0000\u02e9\u02e2\u0001\u0000\u0000\u0000\u02e9"+
		"\u02e3\u0001\u0000\u0000\u0000\u02ea\u02ef\u0001\u0000\u0000\u0000\u02eb"+
		"\u02ec\n\u0002\u0000\u0000\u02ec\u02ee\u0005q\u0000\u0000\u02ed\u02eb"+
		"\u0001\u0000\u0000\u0000\u02ee\u02f1\u0001\u0000\u0000\u0000\u02ef\u02ed"+
		"\u0001\u0000\u0000\u0000\u02ef\u02f0\u0001\u0000\u0000\u0000\u02f0o\u0001"+
		"\u0000\u0000\u0000\u02f1\u02ef\u0001\u0000\u0000\u0000\u02f2\u02f6\u0003"+
		"r9\u0000\u02f3\u02f6\u0003t:\u0000\u02f4\u02f6\u0003v;\u0000\u02f5\u02f2"+
		"\u0001\u0000\u0000\u0000\u02f5\u02f3\u0001\u0000\u0000\u0000\u02f5\u02f4"+
		"\u0001\u0000\u0000\u0000\u02f6q\u0001\u0000\u0000\u0000\u02f7\u02f8\u0003"+
		"\u0006\u0003\u0000\u02f8\u02f9\u0005\u000e\u0000\u0000\u02f9\u02fa\u0005"+
		"u\u0000\u0000\u02fa\u0301\u0001\u0000\u0000\u0000\u02fb\u02fc\u0003\u0006"+
		"\u0003\u0000\u02fc\u02fd\u0005\u000f\u0000\u0000\u02fd\u02fe\u0003\u0006"+
		"\u0003\u0000\u02fe\u02ff\u0005\r\u0000\u0000\u02ff\u0301\u0001\u0000\u0000"+
		"\u0000\u0300\u02f7\u0001\u0000\u0000\u0000\u0300\u02fb\u0001\u0000\u0000"+
		"\u0000\u0301s\u0001\u0000\u0000\u0000\u0302\u0303\u0003\u0006\u0003\u0000"+
		"\u0303\u0304\u0005s\u0000\u0000\u0304\u0305\u0003\u0006\u0003\u0000\u0305"+
		"u\u0001\u0000\u0000\u0000\u0306\u0307\u0005t\u0000\u0000\u0307\u030d\u0005"+
		"u\u0000\u0000\u0308\u030a\u0005\b\u0000\u0000\u0309\u030b\u0003d2\u0000"+
		"\u030a\u0309\u0001\u0000\u0000\u0000\u030a\u030b\u0001\u0000\u0000\u0000"+
		"\u030b\u030c\u0001\u0000\u0000\u0000\u030c\u030e\u0005\t\u0000\u0000\u030d"+
		"\u0308\u0001\u0000\u0000\u0000\u030d\u030e\u0001\u0000\u0000\u0000\u030e"+
		"w\u0001\u0000\u0000\u0000H{\u0089\u0090\u009b\u00a7\u00b1\u00bd\u00bf"+
		"\u00cd\u00d2\u00d9\u00e2\u00f3\u00f5\u00fe\u0107\u010a\u0117\u011a\u011e"+
		"\u0126\u0129\u012d\u0138\u013f\u0149\u0158\u015d\u0164\u0171\u0182\u01a8"+
		"\u01af\u01b9\u01c6\u01cb\u01e2\u01f0\u01fd\u0200\u0205\u0212\u0218\u021d"+
		"\u0222\u022a\u0238\u023b\u0241\u0248\u0251\u0256\u0259\u0265\u0269\u0279"+
		"\u029a\u02a2\u02a6\u02ab\u02b3\u02b6\u02be\u02c4\u02d0\u02d3\u02e9\u02ef"+
		"\u02f5\u0300\u030a\u030d";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}