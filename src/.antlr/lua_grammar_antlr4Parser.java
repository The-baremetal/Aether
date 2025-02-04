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
		T__113=114, T__114=115, IDENTIFIER=116, BOOL=117, NIL=118, NUMBER=119, 
		STRING=120, WS=121, COMMENT=122;
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
		RULE_incrOp = 45, RULE_coalesceOp = 46, RULE_shiftAssignOp = 47, RULE_identifierList = 48, 
		RULE_expressionList = 49, RULE_functionExpression = 50, RULE_selectStatement = 51, 
		RULE_lambdaExpression = 52, RULE_typeAnnotation = 53, RULE_typeSpec = 54, 
		RULE_experimentalExpression = 55, RULE_safeNavigation = 56, RULE_pipeOperator = 57, 
		RULE_decoratorSyntax = 58;
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
			"incrOp", "coalesceOp", "shiftAssignOp", "identifierList", "expressionList", 
			"functionExpression", "selectStatement", "lambdaExpression", "typeAnnotation", 
			"typeSpec", "experimentalExpression", "safeNavigation", "pipeOperator", 
			"decoratorSyntax"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'spawn'", "'match'", "'with'", "'end'", "'async'", "'await'", 
			"'('", "')'", "':'", "'.'", "'['", "']'", "'?.'", "'?['", "','", "'{'", 
			"'}'", "'__metatable'", "'='", "'::'", "'when'", "'=>'", "'|'", "'__add'", 
			"'__sub'", "'__mul'", "'__div'", "'__mod'", "'__pow'", "'__unm'", "'__concat'", 
			"'__len'", "'__eq'", "'__lt'", "'__le'", "'__tostring'", "'__pairs'", 
			"'__ipairs'", "'__call'", "'#'", "'if'", "'then'", "'elseif'", "'else'", 
			"'while'", "'do'", "'repeat'", "'until'", "'for'", "'in'", "'break'", 
			"'goto'", "'coroutine'", "'create'", "'resume'", "'yield'", "'status'", 
			"'running'", "'wrap'", "'isyieldable'", "'pcall'", "'xpcall'", "'local'", 
			"'function'", "'return'", "'and'", "'or'", "'=='", "'>='", "'<='", "'~='", 
			"'>'", "'<'", "'+'", "'-'", "'*'", "'/'", "'//'", "'%'", "'^'", "'&'", 
			"'~'", "'<<'", "'>>'", "'+='", "'-='", "'*='", "'/='", "'//='", "'^='", 
			"'&='", "'|='", "'not'", "'typeof'", "'..'", "'...'", "'..='", "'??='", 
			"'+_='", "'-_='", "'+_'", "'-_'", "'??'", "'<<='", "'>>='", "'select'", 
			"'number'", "'string'", "'boolean'", "'table'", "'any'", "'[]'", "'table<'", 
			"'|>'", "'@'", null, null, "'nil'"
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
			null, null, null, null, null, null, null, null, "IDENTIFIER", "BOOL", 
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
			setState(121);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & -2289339339679727614L) != 0) || ((((_la - 64)) & ~0x3f) == 0 && ((1L << (_la - 64)) & 4507997673881603L) != 0)) {
				{
				{
				setState(118);
				statement();
				}
				}
				setState(123);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(124);
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
			setState(135);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case IDENTIFIER:
				enterOuterAlt(_localctx, 1);
				{
				setState(126);
				assignStatement();
				}
				break;
			case T__40:
			case T__44:
			case T__46:
			case T__48:
			case T__50:
			case T__51:
			case T__52:
			case T__60:
			case T__61:
				enterOuterAlt(_localctx, 2);
				{
				setState(127);
				controlFlowStatement();
				}
				break;
			case T__63:
				enterOuterAlt(_localctx, 3);
				{
				setState(128);
				functionDeclaration();
				}
				break;
			case T__64:
				enterOuterAlt(_localctx, 4);
				{
				setState(129);
				returnStatement();
				}
				break;
			case T__62:
				enterOuterAlt(_localctx, 5);
				{
				setState(130);
				localDeclaration();
				}
				break;
			case T__19:
				enterOuterAlt(_localctx, 6);
				{
				setState(131);
				labelStatement();
				}
				break;
			case T__105:
				enterOuterAlt(_localctx, 7);
				{
				setState(132);
				selectStatement();
				}
				break;
			case T__0:
				enterOuterAlt(_localctx, 8);
				{
				setState(133);
				match(T__0);
				setState(134);
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
			setState(137);
			variable(0);
			setState(142);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__18:
			case T__84:
			case T__85:
			case T__86:
			case T__87:
			case T__88:
			case T__89:
			case T__90:
			case T__91:
				{
				setState(138);
				assignOp();
				}
				break;
			case T__100:
			case T__101:
				{
				setState(139);
				incrOp();
				}
				break;
			case T__103:
			case T__104:
				{
				setState(140);
				shiftAssignOp();
				}
				break;
			case T__102:
				{
				setState(141);
				coalesceOp();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			setState(144);
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
			setState(175);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,5,_ctx) ) {
			case 1:
				{
				setState(147);
				primaryExpression();
				setState(153);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,3,_ctx);
				while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
					if ( _alt==1 ) {
						{
						{
						setState(148);
						operatorGroup();
						setState(149);
						expression(0);
						}
						} 
					}
					setState(155);
					_errHandler.sync(this);
					_alt = getInterpreter().adaptivePredict(_input,3,_ctx);
				}
				}
				break;
			case 2:
				{
				setState(156);
				unaryOp();
				setState(157);
				expression(4);
				}
				break;
			case 3:
				{
				setState(159);
				match(T__1);
				setState(160);
				expression(0);
				setState(161);
				match(T__2);
				setState(163); 
				_errHandler.sync(this);
				_la = _input.LA(1);
				do {
					{
					{
					setState(162);
					matchArm();
					}
					}
					setState(165); 
					_errHandler.sync(this);
					_la = _input.LA(1);
				} while ( _la==T__15 || ((((_la - 116)) & ~0x3f) == 0 && ((1L << (_la - 116)) & 31L) != 0) );
				setState(167);
				match(T__3);
				}
				break;
			case 4:
				{
				setState(169);
				match(T__4);
				setState(170);
				block();
				setState(171);
				match(T__3);
				}
				break;
			case 5:
				{
				setState(173);
				match(T__5);
				setState(174);
				expression(1);
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(187);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,7,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(185);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,6,_ctx) ) {
					case 1:
						{
						_localctx = new ExpressionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(177);
						if (!(precpred(_ctx, 6))) throw new FailedPredicateException(this, "precpred(_ctx, 6)");
						setState(178);
						operatorGroup();
						setState(179);
						expression(7);
						}
						break;
					case 2:
						{
						_localctx = new ExpressionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(181);
						if (!(precpred(_ctx, 5))) throw new FailedPredicateException(this, "precpred(_ctx, 5)");
						setState(182);
						operatorGroup();
						setState(183);
						expression(5);
						}
						break;
					}
					} 
				}
				setState(189);
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
			setState(201);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,8,_ctx) ) {
			case 1:
				{
				setState(190);
				literal();
				}
				break;
			case 2:
				{
				setState(191);
				variable(0);
				}
				break;
			case 3:
				{
				setState(192);
				functionCall();
				}
				break;
			case 4:
				{
				setState(193);
				unaryOperation();
				}
				break;
			case 5:
				{
				setState(194);
				tableConstructor();
				}
				break;
			case 6:
				{
				setState(195);
				functionExpression();
				}
				break;
			case 7:
				{
				setState(196);
				match(T__6);
				setState(197);
				expression(0);
				setState(198);
				match(T__7);
				}
				break;
			case 8:
				{
				setState(200);
				lambdaExpression();
				}
				break;
			}
			setState(206);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,9,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(203);
					callChain();
					}
					} 
				}
				setState(208);
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
			setState(222);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__8:
				_localctx = new MethodChainContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(209);
				match(T__8);
				setState(210);
				match(IDENTIFIER);
				setState(211);
				match(T__6);
				setState(213);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 1099511693540L) != 0) || ((((_la - 64)) & ~0x3f) == 0 && ((1L << (_la - 64)) & 139611594354329601L) != 0)) {
					{
					setState(212);
					expressionList();
					}
				}

				setState(215);
				match(T__7);
				}
				break;
			case T__9:
				_localctx = new PropertyChainContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(216);
				match(T__9);
				setState(217);
				match(IDENTIFIER);
				}
				break;
			case T__10:
				_localctx = new IndexChainContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(218);
				match(T__10);
				setState(219);
				expression(0);
				setState(220);
				match(T__11);
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
			setState(224);
			_la = _input.LA(1);
			if ( !(((((_la - 117)) & ~0x3f) == 0 && ((1L << (_la - 117)) & 15L) != 0)) ) {
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
			setState(227);
			match(IDENTIFIER);
			}
			_ctx.stop = _input.LT(-1);
			setState(241);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,13,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(239);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,12,_ctx) ) {
					case 1:
						{
						_localctx = new VariableContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_variable);
						setState(229);
						if (!(precpred(_ctx, 3))) throw new FailedPredicateException(this, "precpred(_ctx, 3)");
						setState(230);
						safeAccess();
						}
						break;
					case 2:
						{
						_localctx = new VariableContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_variable);
						setState(231);
						if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
						setState(232);
						match(T__10);
						setState(233);
						expression(0);
						setState(234);
						match(T__11);
						}
						break;
					case 3:
						{
						_localctx = new VariableContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_variable);
						setState(236);
						if (!(precpred(_ctx, 1))) throw new FailedPredicateException(this, "precpred(_ctx, 1)");
						setState(237);
						match(T__9);
						setState(238);
						match(IDENTIFIER);
						}
						break;
					}
					} 
				}
				setState(243);
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
			setState(250);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__12:
				enterOuterAlt(_localctx, 1);
				{
				setState(244);
				match(T__12);
				setState(245);
				match(IDENTIFIER);
				}
				break;
			case T__13:
				enterOuterAlt(_localctx, 2);
				{
				setState(246);
				match(T__13);
				setState(247);
				expression(0);
				setState(248);
				match(T__11);
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
			setState(282);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,19,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(252);
				variable(0);
				setState(253);
				match(T__6);
				setState(262);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 1099511693540L) != 0) || ((((_la - 64)) & ~0x3f) == 0 && ((1L << (_la - 64)) & 139611590059362305L) != 0)) {
					{
					setState(254);
					expression(0);
					setState(259);
					_errHandler.sync(this);
					_la = _input.LA(1);
					while (_la==T__14) {
						{
						{
						setState(255);
						match(T__14);
						setState(256);
						expression(0);
						}
						}
						setState(261);
						_errHandler.sync(this);
						_la = _input.LA(1);
					}
					}
				}

				setState(264);
				match(T__7);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(266);
				variable(0);
				setState(267);
				match(T__8);
				setState(268);
				match(IDENTIFIER);
				setState(269);
				match(T__6);
				setState(278);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 1099511693540L) != 0) || ((((_la - 64)) & ~0x3f) == 0 && ((1L << (_la - 64)) & 139611590059362305L) != 0)) {
					{
					setState(270);
					expression(0);
					setState(275);
					_errHandler.sync(this);
					_la = _input.LA(1);
					while (_la==T__14) {
						{
						{
						setState(271);
						match(T__14);
						setState(272);
						expression(0);
						}
						}
						setState(277);
						_errHandler.sync(this);
						_la = _input.LA(1);
					}
					}
				}

				setState(280);
				match(T__7);
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
			setState(284);
			match(T__15);
			setState(293);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 1099511695588L) != 0) || ((((_la - 64)) & ~0x3f) == 0 && ((1L << (_la - 64)) & 139611590059362305L) != 0)) {
				{
				setState(285);
				field();
				setState(290);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,20,_ctx);
				while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
					if ( _alt==1 ) {
						{
						{
						setState(286);
						match(T__14);
						setState(287);
						field();
						}
						} 
					}
					setState(292);
					_errHandler.sync(this);
					_alt = getInterpreter().adaptivePredict(_input,20,_ctx);
				}
				}
			}

			setState(297);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__14) {
				{
				setState(295);
				match(T__14);
				setState(296);
				metatable();
				}
			}

			setState(299);
			match(T__16);
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
			setState(308);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__17:
				enterOuterAlt(_localctx, 1);
				{
				setState(301);
				match(T__17);
				setState(302);
				match(T__18);
				setState(303);
				expression(0);
				}
				break;
			case T__15:
				enterOuterAlt(_localctx, 2);
				{
				setState(304);
				match(T__15);
				setState(305);
				metamethods();
				setState(306);
				match(T__16);
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
			setState(310);
			metamethod();
			setState(315);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__14) {
				{
				{
				setState(311);
				match(T__14);
				setState(312);
				metamethod();
				}
				}
				setState(317);
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
			setState(318);
			match(T__19);
			setState(319);
			match(IDENTIFIER);
			setState(320);
			match(T__19);
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
			setState(322);
			pattern(0);
			setState(325);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__20) {
				{
				setState(323);
				match(T__20);
				setState(324);
				expression(0);
				}
			}

			setState(327);
			match(T__21);
			setState(328);
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
			setState(345);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,27,_ctx) ) {
			case 1:
				{
				setState(331);
				literal();
				}
				break;
			case 2:
				{
				setState(332);
				tableConstructor();
				}
				break;
			case 3:
				{
				setState(333);
				match(IDENTIFIER);
				}
				break;
			case 4:
				{
				setState(334);
				match(T__15);
				setState(335);
				fieldPattern();
				setState(340);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==T__14) {
					{
					{
					setState(336);
					match(T__14);
					setState(337);
					fieldPattern();
					}
					}
					setState(342);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(343);
				match(T__16);
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(352);
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
					setState(347);
					if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
					setState(348);
					match(T__22);
					setState(349);
					pattern(3);
					}
					} 
				}
				setState(354);
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
			setState(365);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,29,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(355);
				match(IDENTIFIER);
				setState(356);
				match(T__18);
				setState(357);
				pattern(0);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(358);
				match(T__10);
				setState(359);
				expression(0);
				setState(360);
				match(T__11);
				setState(361);
				match(T__18);
				setState(362);
				pattern(0);
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(364);
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
			setState(367);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 1099494850560L) != 0)) ) {
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
			setState(382);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,30,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(369);
				match(IDENTIFIER);
				setState(370);
				match(T__18);
				setState(371);
				expression(0);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(372);
				match(T__10);
				setState(373);
				expression(0);
				setState(374);
				match(T__11);
				setState(375);
				match(T__18);
				setState(376);
				expression(0);
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(378);
				expression(0);
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(379);
				match(IDENTIFIER);
				setState(380);
				match(T__8);
				setState(381);
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
			setState(420);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,31,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(384);
				expression(0);
				setState(385);
				arithOp();
				setState(386);
				expression(0);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(388);
				expression(0);
				setState(389);
				bitwiseOp();
				setState(390);
				expression(0);
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(392);
				expression(0);
				setState(393);
				comparisonOp();
				setState(394);
				expression(0);
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(396);
				expression(0);
				setState(397);
				logicalOp();
				setState(398);
				expression(0);
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(400);
				expression(0);
				setState(401);
				concatOp();
				setState(402);
				expression(0);
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(404);
				expression(0);
				setState(405);
				compoundAssignOp();
				setState(406);
				expression(0);
				}
				break;
			case 7:
				enterOuterAlt(_localctx, 7);
				{
				setState(408);
				expression(0);
				setState(409);
				coalesceOp();
				setState(410);
				expression(0);
				}
				break;
			case 8:
				enterOuterAlt(_localctx, 8);
				{
				setState(412);
				expression(0);
				setState(413);
				shiftAssignOp();
				setState(414);
				expression(0);
				}
				break;
			case 9:
				enterOuterAlt(_localctx, 9);
				{
				setState(416);
				expression(0);
				setState(417);
				incrOp();
				setState(418);
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
			setState(427);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,32,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(422);
				unaryOp();
				setState(423);
				expression(0);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(425);
				match(T__39);
				setState(426);
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
			setState(437);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__40:
				enterOuterAlt(_localctx, 1);
				{
				setState(429);
				ifStatement();
				}
				break;
			case T__44:
				enterOuterAlt(_localctx, 2);
				{
				setState(430);
				whileStatement();
				}
				break;
			case T__46:
				enterOuterAlt(_localctx, 3);
				{
				setState(431);
				repeatStatement();
				}
				break;
			case T__48:
				enterOuterAlt(_localctx, 4);
				{
				setState(432);
				forStatement();
				}
				break;
			case T__50:
				enterOuterAlt(_localctx, 5);
				{
				setState(433);
				breakStatement();
				}
				break;
			case T__51:
				enterOuterAlt(_localctx, 6);
				{
				setState(434);
				gotoStatement();
				}
				break;
			case T__52:
				enterOuterAlt(_localctx, 7);
				{
				setState(435);
				coroutineStatement();
				}
				break;
			case T__60:
			case T__61:
				enterOuterAlt(_localctx, 8);
				{
				setState(436);
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
			setState(439);
			match(T__40);
			setState(440);
			expression(0);
			setState(441);
			match(T__41);
			setState(442);
			block();
			setState(450);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__42) {
				{
				{
				setState(443);
				match(T__42);
				setState(444);
				expression(0);
				setState(445);
				match(T__41);
				setState(446);
				block();
				}
				}
				setState(452);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(455);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__43) {
				{
				setState(453);
				match(T__43);
				setState(454);
				block();
				}
			}

			setState(457);
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
			setState(459);
			match(T__44);
			setState(460);
			expression(0);
			setState(461);
			match(T__45);
			setState(462);
			block();
			setState(463);
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
			setState(465);
			match(T__46);
			setState(466);
			block();
			setState(467);
			match(T__47);
			setState(468);
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
			setState(492);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,37,_ctx) ) {
			case 1:
				_localctx = new NumericForContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(470);
				match(T__48);
				setState(471);
				match(IDENTIFIER);
				setState(472);
				match(T__18);
				setState(473);
				expression(0);
				setState(474);
				match(T__14);
				setState(475);
				expression(0);
				setState(478);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==T__14) {
					{
					setState(476);
					match(T__14);
					setState(477);
					expression(0);
					}
				}

				setState(480);
				match(T__45);
				setState(481);
				block();
				setState(482);
				match(T__3);
				}
				break;
			case 2:
				_localctx = new GenericForContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(484);
				match(T__48);
				setState(485);
				identifierList();
				setState(486);
				match(T__49);
				setState(487);
				expressionList();
				setState(488);
				match(T__45);
				setState(489);
				block();
				setState(490);
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
			setState(494);
			match(T__50);
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
			setState(496);
			match(T__51);
			setState(497);
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
			setState(499);
			match(T__52);
			setState(500);
			match(T__9);
			setState(501);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 2287828610704211968L) != 0)) ) {
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
			setState(503);
			_la = _input.LA(1);
			if ( !(_la==T__60 || _la==T__61) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			setState(505);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__8 || _la==T__9) {
				{
				setState(504);
				_la = _input.LA(1);
				if ( !(_la==T__8 || _la==T__9) ) {
				_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				}
			}

			setState(508);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==IDENTIFIER) {
				{
				setState(507);
				match(IDENTIFIER);
				}
			}

			}
			setState(510);
			match(T__6);
			setState(513);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,40,_ctx) ) {
			case 1:
				{
				setState(511);
				expressionList();
				}
				break;
			case 2:
				{
				setState(512);
				namedArgs();
				}
				break;
			}
			setState(515);
			match(T__7);
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
			setState(517);
			match(IDENTIFIER);
			setState(518);
			match(T__18);
			setState(519);
			expression(0);
			setState(526);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__14) {
				{
				{
				setState(520);
				match(T__14);
				setState(521);
				match(IDENTIFIER);
				setState(522);
				match(T__18);
				setState(523);
				expression(0);
				}
				}
				setState(528);
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
			setState(532);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & -2289339339679727614L) != 0) || ((((_la - 64)) & ~0x3f) == 0 && ((1L << (_la - 64)) & 4507997673881603L) != 0)) {
				{
				{
				setState(529);
				statement();
				}
				}
				setState(534);
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
			setState(573);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,48,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(535);
				match(T__62);
				setState(537);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==T__8) {
					{
					setState(536);
					typeAnnotation();
					}
				}

				setState(539);
				match(IDENTIFIER);
				setState(542);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==T__18) {
					{
					setState(540);
					match(T__18);
					setState(541);
					expression(0);
					}
				}

				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(544);
				match(T__62);
				setState(545);
				match(IDENTIFIER);
				setState(550);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==T__14) {
					{
					{
					setState(546);
					match(T__14);
					setState(547);
					match(IDENTIFIER);
					}
					}
					setState(552);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(553);
				match(T__18);
				setState(554);
				expressionList();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(555);
				match(T__62);
				setState(556);
				match(T__63);
				setState(557);
				match(IDENTIFIER);
				setState(558);
				match(T__6);
				setState(567);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==IDENTIFIER) {
					{
					setState(559);
					match(IDENTIFIER);
					setState(564);
					_errHandler.sync(this);
					_la = _input.LA(1);
					while (_la==T__14) {
						{
						{
						setState(560);
						match(T__14);
						setState(561);
						match(IDENTIFIER);
						}
						}
						setState(566);
						_errHandler.sync(this);
						_la = _input.LA(1);
					}
					}
				}

				setState(569);
				match(T__7);
				setState(570);
				block();
				setState(571);
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
			setState(575);
			match(T__63);
			setState(580);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,49,_ctx) ) {
			case 1:
				{
				setState(576);
				match(IDENTIFIER);
				setState(577);
				match(T__9);
				}
				break;
			case 2:
				{
				setState(578);
				match(IDENTIFIER);
				setState(579);
				match(T__8);
				}
				break;
			}
			setState(582);
			match(IDENTIFIER);
			setState(583);
			match(T__6);
			setState(597);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case IDENTIFIER:
				{
				setState(584);
				match(IDENTIFIER);
				setState(589);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,50,_ctx);
				while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
					if ( _alt==1 ) {
						{
						{
						setState(585);
						match(T__14);
						setState(586);
						match(IDENTIFIER);
						}
						} 
					}
					setState(591);
					_errHandler.sync(this);
					_alt = getInterpreter().adaptivePredict(_input,50,_ctx);
				}
				setState(594);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==T__14) {
					{
					setState(592);
					match(T__14);
					setState(593);
					varargOp();
					}
				}

				}
				break;
			case T__95:
				{
				setState(596);
				varargOp();
				}
				break;
			case T__7:
				break;
			default:
				break;
			}
			setState(599);
			match(T__7);
			setState(600);
			block();
			setState(601);
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
			setState(603);
			match(T__64);
			setState(613);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,54,_ctx) ) {
			case 1:
				{
				setState(604);
				expression(0);
				setState(609);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==T__14) {
					{
					{
					setState(605);
					match(T__14);
					setState(606);
					expression(0);
					}
					}
					setState(611);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				}
				break;
			case 2:
				{
				setState(612);
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
		public OperatorGroupContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_operatorGroup; }
	}

	public final OperatorGroupContext operatorGroup() throws RecognitionException {
		OperatorGroupContext _localctx = new OperatorGroupContext(_ctx, getState());
		enterRule(_localctx, 70, RULE_operatorGroup);
		try {
			setState(628);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,55,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(615);
				logicalOp();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(616);
				comparisonOp();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(617);
				arithOp();
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(618);
				bitwiseOp();
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(619);
				assignOp();
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(620);
				unaryOp();
				}
				break;
			case 7:
				enterOuterAlt(_localctx, 7);
				{
				setState(621);
				concatOp();
				}
				break;
			case 8:
				enterOuterAlt(_localctx, 8);
				{
				setState(622);
				compoundAssignOp();
				}
				break;
			case 9:
				enterOuterAlt(_localctx, 9);
				{
				setState(623);
				incrOp();
				}
				break;
			case 10:
				enterOuterAlt(_localctx, 10);
				{
				setState(624);
				coalesceOp();
				}
				break;
			case 11:
				enterOuterAlt(_localctx, 11);
				{
				setState(625);
				shiftAssignOp();
				}
				break;
			case 12:
				enterOuterAlt(_localctx, 12);
				{
				setState(626);
				match(T__21);
				}
				break;
			case 13:
				enterOuterAlt(_localctx, 13);
				{
				setState(627);
				safeAccess();
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
			setState(630);
			_la = _input.LA(1);
			if ( !(_la==T__65 || _la==T__66) ) {
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
			setState(632);
			_la = _input.LA(1);
			if ( !(((((_la - 68)) & ~0x3f) == 0 && ((1L << (_la - 68)) & 63L) != 0)) ) {
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
			setState(634);
			_la = _input.LA(1);
			if ( !(((((_la - 74)) & ~0x3f) == 0 && ((1L << (_la - 74)) & 127L) != 0)) ) {
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
			setState(636);
			_la = _input.LA(1);
			if ( !(((((_la - 23)) & ~0x3f) == 0 && ((1L << (_la - 23)) & 4323455642275676161L) != 0)) ) {
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
			setState(638);
			_la = _input.LA(1);
			if ( !(_la==T__18 || ((((_la - 85)) & ~0x3f) == 0 && ((1L << (_la - 85)) & 255L) != 0)) ) {
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
			setState(640);
			_la = _input.LA(1);
			if ( !(((((_la - 40)) & ~0x3f) == 0 && ((1L << (_la - 40)) & 27026030170472449L) != 0)) ) {
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
			setState(642);
			match(T__94);
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
			setState(644);
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
			setState(646);
			_la = _input.LA(1);
			if ( !(((((_la - 85)) & ~0x3f) == 0 && ((1L << (_la - 85)) & 61503L) != 0)) ) {
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
			setState(648);
			_la = _input.LA(1);
			if ( !(_la==T__100 || _la==T__101) ) {
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
			setState(650);
			match(T__102);
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
			setState(652);
			_la = _input.LA(1);
			if ( !(_la==T__103 || _la==T__104) ) {
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
		enterRule(_localctx, 96, RULE_identifierList);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(654);
			match(IDENTIFIER);
			setState(659);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__14) {
				{
				{
				setState(655);
				match(T__14);
				setState(656);
				match(IDENTIFIER);
				}
				}
				setState(661);
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
		enterRule(_localctx, 98, RULE_expressionList);
		int _la;
		try {
			setState(671);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__1:
			case T__4:
			case T__5:
			case T__6:
			case T__15:
			case T__39:
			case T__63:
			case T__74:
			case T__81:
			case T__92:
			case T__93:
			case IDENTIFIER:
			case BOOL:
			case NIL:
			case NUMBER:
			case STRING:
				enterOuterAlt(_localctx, 1);
				{
				setState(662);
				expression(0);
				setState(667);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==T__14) {
					{
					{
					setState(663);
					match(T__14);
					setState(664);
					expression(0);
					}
					}
					setState(669);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				}
				break;
			case T__95:
				enterOuterAlt(_localctx, 2);
				{
				setState(670);
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
		enterRule(_localctx, 100, RULE_functionExpression);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(673);
			match(T__63);
			setState(676);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==IDENTIFIER) {
				{
				setState(674);
				match(IDENTIFIER);
				setState(675);
				match(T__8);
				}
			}

			setState(678);
			match(T__6);
			setState(687);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==IDENTIFIER) {
				{
				setState(679);
				match(IDENTIFIER);
				setState(684);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==T__14) {
					{
					{
					setState(680);
					match(T__14);
					setState(681);
					match(IDENTIFIER);
					}
					}
					setState(686);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				}
			}

			setState(689);
			match(T__7);
			setState(695);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__0:
			case T__3:
			case T__19:
			case T__40:
			case T__44:
			case T__46:
			case T__48:
			case T__50:
			case T__51:
			case T__52:
			case T__60:
			case T__61:
			case T__62:
			case T__63:
			case T__64:
			case T__105:
			case IDENTIFIER:
				{
				setState(690);
				block();
				setState(691);
				match(T__3);
				}
				break;
			case T__21:
				{
				setState(693);
				match(T__21);
				setState(694);
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
		enterRule(_localctx, 102, RULE_selectStatement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(697);
			match(T__105);
			setState(698);
			match(T__6);
			setState(701);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,63,_ctx) ) {
			case 1:
				{
				setState(699);
				match(T__39);
				}
				break;
			case 2:
				{
				setState(700);
				expression(0);
				}
				break;
			}
			setState(703);
			match(T__14);
			setState(704);
			expression(0);
			setState(705);
			match(T__7);
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
		enterRule(_localctx, 104, RULE_lambdaExpression);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(707);
			match(T__6);
			setState(716);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==IDENTIFIER) {
				{
				setState(708);
				match(IDENTIFIER);
				setState(713);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==T__14) {
					{
					{
					setState(709);
					match(T__14);
					setState(710);
					match(IDENTIFIER);
					}
					}
					setState(715);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				}
			}

			setState(718);
			match(T__7);
			setState(719);
			match(T__21);
			setState(720);
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
		enterRule(_localctx, 106, RULE_typeAnnotation);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(722);
			match(T__8);
			setState(723);
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
		int _startState = 108;
		enterRecursionRule(_localctx, 108, RULE_typeSpec, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(738);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__106:
				{
				setState(726);
				match(T__106);
				}
				break;
			case T__107:
				{
				setState(727);
				match(T__107);
				}
				break;
			case T__108:
				{
				setState(728);
				match(T__108);
				}
				break;
			case T__109:
				{
				setState(729);
				match(T__109);
				}
				break;
			case T__63:
				{
				setState(730);
				match(T__63);
				}
				break;
			case T__110:
				{
				setState(731);
				match(T__110);
				}
				break;
			case T__112:
				{
				setState(732);
				match(T__112);
				setState(733);
				typeSpec(0);
				setState(734);
				match(T__14);
				setState(735);
				typeSpec(0);
				setState(736);
				match(T__71);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			_ctx.stop = _input.LT(-1);
			setState(744);
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
					setState(740);
					if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
					setState(741);
					match(T__111);
					}
					} 
				}
				setState(746);
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
		enterRule(_localctx, 110, RULE_experimentalExpression);
		try {
			setState(753);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,68,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(747);
				if (!(experiments.isEnabled("safe_nav"))) throw new FailedPredicateException(this, "experiments.isEnabled(\"safe_nav\")");
				setState(748);
				safeNavigation();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(749);
				if (!(experiments.isEnabled("pipe"))) throw new FailedPredicateException(this, "experiments.isEnabled(\"pipe\")");
				setState(750);
				pipeOperator();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(751);
				if (!(experiments.isEnabled("decorators"))) throw new FailedPredicateException(this, "experiments.isEnabled(\"decorators\")");
				setState(752);
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
		enterRule(_localctx, 112, RULE_safeNavigation);
		try {
			setState(764);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,69,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(755);
				expression(0);
				setState(756);
				match(T__12);
				setState(757);
				match(IDENTIFIER);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(759);
				expression(0);
				setState(760);
				match(T__13);
				setState(761);
				expression(0);
				setState(762);
				match(T__11);
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
		enterRule(_localctx, 114, RULE_pipeOperator);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(766);
			expression(0);
			setState(767);
			match(T__113);
			setState(768);
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
		enterRule(_localctx, 116, RULE_decoratorSyntax);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(770);
			match(T__114);
			setState(771);
			match(IDENTIFIER);
			setState(777);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__6) {
				{
				setState(772);
				match(T__6);
				setState(774);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 1099511693540L) != 0) || ((((_la - 64)) & ~0x3f) == 0 && ((1L << (_la - 64)) & 139611594354329601L) != 0)) {
					{
					setState(773);
					expressionList();
					}
				}

				setState(776);
				match(T__7);
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
		case 54:
			return typeSpec_sempred((TypeSpecContext)_localctx, predIndex);
		case 55:
			return experimentalExpression_sempred((ExperimentalExpressionContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean expression_sempred(ExpressionContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 6);
		case 1:
			return precpred(_ctx, 5);
		}
		return true;
	}
	private boolean variable_sempred(VariableContext _localctx, int predIndex) {
		switch (predIndex) {
		case 2:
			return precpred(_ctx, 3);
		case 3:
			return precpred(_ctx, 2);
		case 4:
			return precpred(_ctx, 1);
		}
		return true;
	}
	private boolean pattern_sempred(PatternContext _localctx, int predIndex) {
		switch (predIndex) {
		case 5:
			return precpred(_ctx, 2);
		}
		return true;
	}
	private boolean typeSpec_sempred(TypeSpecContext _localctx, int predIndex) {
		switch (predIndex) {
		case 6:
			return precpred(_ctx, 2);
		}
		return true;
	}
	private boolean experimentalExpression_sempred(ExperimentalExpressionContext _localctx, int predIndex) {
		switch (predIndex) {
		case 7:
			return experiments.isEnabled("safe_nav");
		case 8:
			return experiments.isEnabled("pipe");
		case 9:
			return experiments.isEnabled("decorators");
		}
		return true;
	}

	public static final String _serializedATN =
		"\u0004\u0001z\u030c\u0002\u0000\u0007\u0000\u0002\u0001\u0007\u0001\u0002"+
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
		"7\u00077\u00028\u00078\u00029\u00079\u0002:\u0007:\u0001\u0000\u0005\u0000"+
		"x\b\u0000\n\u0000\f\u0000{\t\u0000\u0001\u0000\u0001\u0000\u0001\u0001"+
		"\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001"+
		"\u0001\u0001\u0001\u0001\u0003\u0001\u0088\b\u0001\u0001\u0002\u0001\u0002"+
		"\u0001\u0002\u0001\u0002\u0001\u0002\u0003\u0002\u008f\b\u0002\u0001\u0002"+
		"\u0001\u0002\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003"+
		"\u0005\u0003\u0098\b\u0003\n\u0003\f\u0003\u009b\t\u0003\u0001\u0003\u0001"+
		"\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0004"+
		"\u0003\u00a4\b\u0003\u000b\u0003\f\u0003\u00a5\u0001\u0003\u0001\u0003"+
		"\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003"+
		"\u0003\u0003\u00b0\b\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003"+
		"\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0005\u0003\u00ba\b\u0003"+
		"\n\u0003\f\u0003\u00bd\t\u0003\u0001\u0004\u0001\u0004\u0001\u0004\u0001"+
		"\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001"+
		"\u0004\u0001\u0004\u0003\u0004\u00ca\b\u0004\u0001\u0004\u0005\u0004\u00cd"+
		"\b\u0004\n\u0004\f\u0004\u00d0\t\u0004\u0001\u0005\u0001\u0005\u0001\u0005"+
		"\u0001\u0005\u0003\u0005\u00d6\b\u0005\u0001\u0005\u0001\u0005\u0001\u0005"+
		"\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0003\u0005\u00df\b\u0005"+
		"\u0001\u0006\u0001\u0006\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007"+
		"\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007"+
		"\u0001\u0007\u0001\u0007\u0001\u0007\u0005\u0007\u00f0\b\u0007\n\u0007"+
		"\f\u0007\u00f3\t\u0007\u0001\b\u0001\b\u0001\b\u0001\b\u0001\b\u0001\b"+
		"\u0003\b\u00fb\b\b\u0001\t\u0001\t\u0001\t\u0001\t\u0001\t\u0005\t\u0102"+
		"\b\t\n\t\f\t\u0105\t\t\u0003\t\u0107\b\t\u0001\t\u0001\t\u0001\t\u0001"+
		"\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001\t\u0005\t\u0112\b\t\n\t\f\t\u0115"+
		"\t\t\u0003\t\u0117\b\t\u0001\t\u0001\t\u0003\t\u011b\b\t\u0001\n\u0001"+
		"\n\u0001\n\u0001\n\u0005\n\u0121\b\n\n\n\f\n\u0124\t\n\u0003\n\u0126\b"+
		"\n\u0001\n\u0001\n\u0003\n\u012a\b\n\u0001\n\u0001\n\u0001\u000b\u0001"+
		"\u000b\u0001\u000b\u0001\u000b\u0001\u000b\u0001\u000b\u0001\u000b\u0003"+
		"\u000b\u0135\b\u000b\u0001\f\u0001\f\u0001\f\u0005\f\u013a\b\f\n\f\f\f"+
		"\u013d\t\f\u0001\r\u0001\r\u0001\r\u0001\r\u0001\u000e\u0001\u000e\u0001"+
		"\u000e\u0003\u000e\u0146\b\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001"+
		"\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0001"+
		"\u000f\u0001\u000f\u0005\u000f\u0153\b\u000f\n\u000f\f\u000f\u0156\t\u000f"+
		"\u0001\u000f\u0001\u000f\u0003\u000f\u015a\b\u000f\u0001\u000f\u0001\u000f"+
		"\u0001\u000f\u0005\u000f\u015f\b\u000f\n\u000f\f\u000f\u0162\t\u000f\u0001"+
		"\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001"+
		"\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0003\u0010\u016e\b\u0010\u0001"+
		"\u0011\u0001\u0011\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001"+
		"\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001"+
		"\u0012\u0001\u0012\u0001\u0012\u0003\u0012\u017f\b\u0012\u0001\u0013\u0001"+
		"\u0013\u0001\u0013\u0001\u0013\u0001\u0013\u0001\u0013\u0001\u0013\u0001"+
		"\u0013\u0001\u0013\u0001\u0013\u0001\u0013\u0001\u0013\u0001\u0013\u0001"+
		"\u0013\u0001\u0013\u0001\u0013\u0001\u0013\u0001\u0013\u0001\u0013\u0001"+
		"\u0013\u0001\u0013\u0001\u0013\u0001\u0013\u0001\u0013\u0001\u0013\u0001"+
		"\u0013\u0001\u0013\u0001\u0013\u0001\u0013\u0001\u0013\u0001\u0013\u0001"+
		"\u0013\u0001\u0013\u0001\u0013\u0001\u0013\u0001\u0013\u0003\u0013\u01a5"+
		"\b\u0013\u0001\u0014\u0001\u0014\u0001\u0014\u0001\u0014\u0001\u0014\u0003"+
		"\u0014\u01ac\b\u0014\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015\u0001"+
		"\u0015\u0001\u0015\u0001\u0015\u0001\u0015\u0003\u0015\u01b6\b\u0015\u0001"+
		"\u0016\u0001\u0016\u0001\u0016\u0001\u0016\u0001\u0016\u0001\u0016\u0001"+
		"\u0016\u0001\u0016\u0001\u0016\u0005\u0016\u01c1\b\u0016\n\u0016\f\u0016"+
		"\u01c4\t\u0016\u0001\u0016\u0001\u0016\u0003\u0016\u01c8\b\u0016\u0001"+
		"\u0016\u0001\u0016\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017\u0001"+
		"\u0017\u0001\u0017\u0001\u0018\u0001\u0018\u0001\u0018\u0001\u0018\u0001"+
		"\u0018\u0001\u0019\u0001\u0019\u0001\u0019\u0001\u0019\u0001\u0019\u0001"+
		"\u0019\u0001\u0019\u0001\u0019\u0003\u0019\u01df\b\u0019\u0001\u0019\u0001"+
		"\u0019\u0001\u0019\u0001\u0019\u0001\u0019\u0001\u0019\u0001\u0019\u0001"+
		"\u0019\u0001\u0019\u0001\u0019\u0001\u0019\u0001\u0019\u0003\u0019\u01ed"+
		"\b\u0019\u0001\u001a\u0001\u001a\u0001\u001b\u0001\u001b\u0001\u001b\u0001"+
		"\u001c\u0001\u001c\u0001\u001c\u0001\u001c\u0001\u001d\u0001\u001d\u0003"+
		"\u001d\u01fa\b\u001d\u0001\u001d\u0003\u001d\u01fd\b\u001d\u0001\u001d"+
		"\u0001\u001d\u0001\u001d\u0003\u001d\u0202\b\u001d\u0001\u001d\u0001\u001d"+
		"\u0001\u001e\u0001\u001e\u0001\u001e\u0001\u001e\u0001\u001e\u0001\u001e"+
		"\u0001\u001e\u0005\u001e\u020d\b\u001e\n\u001e\f\u001e\u0210\t\u001e\u0001"+
		"\u001f\u0005\u001f\u0213\b\u001f\n\u001f\f\u001f\u0216\t\u001f\u0001 "+
		"\u0001 \u0003 \u021a\b \u0001 \u0001 \u0001 \u0003 \u021f\b \u0001 \u0001"+
		" \u0001 \u0001 \u0005 \u0225\b \n \f \u0228\t \u0001 \u0001 \u0001 \u0001"+
		" \u0001 \u0001 \u0001 \u0001 \u0001 \u0005 \u0233\b \n \f \u0236\t \u0003"+
		" \u0238\b \u0001 \u0001 \u0001 \u0001 \u0003 \u023e\b \u0001!\u0001!\u0001"+
		"!\u0001!\u0001!\u0003!\u0245\b!\u0001!\u0001!\u0001!\u0001!\u0001!\u0005"+
		"!\u024c\b!\n!\f!\u024f\t!\u0001!\u0001!\u0003!\u0253\b!\u0001!\u0003!"+
		"\u0256\b!\u0001!\u0001!\u0001!\u0001!\u0001\"\u0001\"\u0001\"\u0001\""+
		"\u0005\"\u0260\b\"\n\"\f\"\u0263\t\"\u0001\"\u0003\"\u0266\b\"\u0001#"+
		"\u0001#\u0001#\u0001#\u0001#\u0001#\u0001#\u0001#\u0001#\u0001#\u0001"+
		"#\u0001#\u0001#\u0003#\u0275\b#\u0001$\u0001$\u0001%\u0001%\u0001&\u0001"+
		"&\u0001\'\u0001\'\u0001(\u0001(\u0001)\u0001)\u0001*\u0001*\u0001+\u0001"+
		"+\u0001,\u0001,\u0001-\u0001-\u0001.\u0001.\u0001/\u0001/\u00010\u0001"+
		"0\u00010\u00050\u0292\b0\n0\f0\u0295\t0\u00011\u00011\u00011\u00051\u029a"+
		"\b1\n1\f1\u029d\t1\u00011\u00031\u02a0\b1\u00012\u00012\u00012\u00032"+
		"\u02a5\b2\u00012\u00012\u00012\u00012\u00052\u02ab\b2\n2\f2\u02ae\t2\u0003"+
		"2\u02b0\b2\u00012\u00012\u00012\u00012\u00012\u00012\u00032\u02b8\b2\u0001"+
		"3\u00013\u00013\u00013\u00033\u02be\b3\u00013\u00013\u00013\u00013\u0001"+
		"4\u00014\u00014\u00014\u00054\u02c8\b4\n4\f4\u02cb\t4\u00034\u02cd\b4"+
		"\u00014\u00014\u00014\u00014\u00015\u00015\u00015\u00016\u00016\u0001"+
		"6\u00016\u00016\u00016\u00016\u00016\u00016\u00016\u00016\u00016\u0001"+
		"6\u00036\u02e3\b6\u00016\u00016\u00056\u02e7\b6\n6\f6\u02ea\t6\u00017"+
		"\u00017\u00017\u00017\u00017\u00017\u00037\u02f2\b7\u00018\u00018\u0001"+
		"8\u00018\u00018\u00018\u00018\u00018\u00018\u00038\u02fd\b8\u00019\u0001"+
		"9\u00019\u00019\u0001:\u0001:\u0001:\u0001:\u0003:\u0307\b:\u0001:\u0003"+
		":\u030a\b:\u0001:\u0000\u0004\u0006\u000e\u001el;\u0000\u0002\u0004\u0006"+
		"\b\n\f\u000e\u0010\u0012\u0014\u0016\u0018\u001a\u001c\u001e \"$&(*,."+
		"02468:<>@BDFHJLNPRTVXZ\\^`bdfhjlnprt\u0000\u000e\u0001\u0000ux\u0001\u0000"+
		"\u0018\'\u0001\u00006<\u0001\u0000=>\u0001\u0000\t\n\u0001\u0000BC\u0001"+
		"\u0000DI\u0001\u0000JP\u0002\u0000\u0017\u0017QT\u0002\u0000\u0013\u0013"+
		"U\\\u0004\u0000((KKRR]^\u0002\u0000UZad\u0001\u0000ef\u0001\u0000hi\u0352"+
		"\u0000y\u0001\u0000\u0000\u0000\u0002\u0087\u0001\u0000\u0000\u0000\u0004"+
		"\u0089\u0001\u0000\u0000\u0000\u0006\u00af\u0001\u0000\u0000\u0000\b\u00c9"+
		"\u0001\u0000\u0000\u0000\n\u00de\u0001\u0000\u0000\u0000\f\u00e0\u0001"+
		"\u0000\u0000\u0000\u000e\u00e2\u0001\u0000\u0000\u0000\u0010\u00fa\u0001"+
		"\u0000\u0000\u0000\u0012\u011a\u0001\u0000\u0000\u0000\u0014\u011c\u0001"+
		"\u0000\u0000\u0000\u0016\u0134\u0001\u0000\u0000\u0000\u0018\u0136\u0001"+
		"\u0000\u0000\u0000\u001a\u013e\u0001\u0000\u0000\u0000\u001c\u0142\u0001"+
		"\u0000\u0000\u0000\u001e\u0159\u0001\u0000\u0000\u0000 \u016d\u0001\u0000"+
		"\u0000\u0000\"\u016f\u0001\u0000\u0000\u0000$\u017e\u0001\u0000\u0000"+
		"\u0000&\u01a4\u0001\u0000\u0000\u0000(\u01ab\u0001\u0000\u0000\u0000*"+
		"\u01b5\u0001\u0000\u0000\u0000,\u01b7\u0001\u0000\u0000\u0000.\u01cb\u0001"+
		"\u0000\u0000\u00000\u01d1\u0001\u0000\u0000\u00002\u01ec\u0001\u0000\u0000"+
		"\u00004\u01ee\u0001\u0000\u0000\u00006\u01f0\u0001\u0000\u0000\u00008"+
		"\u01f3\u0001\u0000\u0000\u0000:\u01f7\u0001\u0000\u0000\u0000<\u0205\u0001"+
		"\u0000\u0000\u0000>\u0214\u0001\u0000\u0000\u0000@\u023d\u0001\u0000\u0000"+
		"\u0000B\u023f\u0001\u0000\u0000\u0000D\u025b\u0001\u0000\u0000\u0000F"+
		"\u0274\u0001\u0000\u0000\u0000H\u0276\u0001\u0000\u0000\u0000J\u0278\u0001"+
		"\u0000\u0000\u0000L\u027a\u0001\u0000\u0000\u0000N\u027c\u0001\u0000\u0000"+
		"\u0000P\u027e\u0001\u0000\u0000\u0000R\u0280\u0001\u0000\u0000\u0000T"+
		"\u0282\u0001\u0000\u0000\u0000V\u0284\u0001\u0000\u0000\u0000X\u0286\u0001"+
		"\u0000\u0000\u0000Z\u0288\u0001\u0000\u0000\u0000\\\u028a\u0001\u0000"+
		"\u0000\u0000^\u028c\u0001\u0000\u0000\u0000`\u028e\u0001\u0000\u0000\u0000"+
		"b\u029f\u0001\u0000\u0000\u0000d\u02a1\u0001\u0000\u0000\u0000f\u02b9"+
		"\u0001\u0000\u0000\u0000h\u02c3\u0001\u0000\u0000\u0000j\u02d2\u0001\u0000"+
		"\u0000\u0000l\u02e2\u0001\u0000\u0000\u0000n\u02f1\u0001\u0000\u0000\u0000"+
		"p\u02fc\u0001\u0000\u0000\u0000r\u02fe\u0001\u0000\u0000\u0000t\u0302"+
		"\u0001\u0000\u0000\u0000vx\u0003\u0002\u0001\u0000wv\u0001\u0000\u0000"+
		"\u0000x{\u0001\u0000\u0000\u0000yw\u0001\u0000\u0000\u0000yz\u0001\u0000"+
		"\u0000\u0000z|\u0001\u0000\u0000\u0000{y\u0001\u0000\u0000\u0000|}\u0005"+
		"\u0000\u0000\u0001}\u0001\u0001\u0000\u0000\u0000~\u0088\u0003\u0004\u0002"+
		"\u0000\u007f\u0088\u0003*\u0015\u0000\u0080\u0088\u0003B!\u0000\u0081"+
		"\u0088\u0003D\"\u0000\u0082\u0088\u0003@ \u0000\u0083\u0088\u0003\u001a"+
		"\r\u0000\u0084\u0088\u0003f3\u0000\u0085\u0086\u0005\u0001\u0000\u0000"+
		"\u0086\u0088\u0003\u0006\u0003\u0000\u0087~\u0001\u0000\u0000\u0000\u0087"+
		"\u007f\u0001\u0000\u0000\u0000\u0087\u0080\u0001\u0000\u0000\u0000\u0087"+
		"\u0081\u0001\u0000\u0000\u0000\u0087\u0082\u0001\u0000\u0000\u0000\u0087"+
		"\u0083\u0001\u0000\u0000\u0000\u0087\u0084\u0001\u0000\u0000\u0000\u0087"+
		"\u0085\u0001\u0000\u0000\u0000\u0088\u0003\u0001\u0000\u0000\u0000\u0089"+
		"\u008e\u0003\u000e\u0007\u0000\u008a\u008f\u0003P(\u0000\u008b\u008f\u0003"+
		"Z-\u0000\u008c\u008f\u0003^/\u0000\u008d\u008f\u0003\\.\u0000\u008e\u008a"+
		"\u0001\u0000\u0000\u0000\u008e\u008b\u0001\u0000\u0000\u0000\u008e\u008c"+
		"\u0001\u0000\u0000\u0000\u008e\u008d\u0001\u0000\u0000\u0000\u008f\u0090"+
		"\u0001\u0000\u0000\u0000\u0090\u0091\u0003\u0006\u0003\u0000\u0091\u0005"+
		"\u0001\u0000\u0000\u0000\u0092\u0093\u0006\u0003\uffff\uffff\u0000\u0093"+
		"\u0099\u0003\b\u0004\u0000\u0094\u0095\u0003F#\u0000\u0095\u0096\u0003"+
		"\u0006\u0003\u0000\u0096\u0098\u0001\u0000\u0000\u0000\u0097\u0094\u0001"+
		"\u0000\u0000\u0000\u0098\u009b\u0001\u0000\u0000\u0000\u0099\u0097\u0001"+
		"\u0000\u0000\u0000\u0099\u009a\u0001\u0000\u0000\u0000\u009a\u00b0\u0001"+
		"\u0000\u0000\u0000\u009b\u0099\u0001\u0000\u0000\u0000\u009c\u009d\u0003"+
		"R)\u0000\u009d\u009e\u0003\u0006\u0003\u0004\u009e\u00b0\u0001\u0000\u0000"+
		"\u0000\u009f\u00a0\u0005\u0002\u0000\u0000\u00a0\u00a1\u0003\u0006\u0003"+
		"\u0000\u00a1\u00a3\u0005\u0003\u0000\u0000\u00a2\u00a4\u0003\u001c\u000e"+
		"\u0000\u00a3\u00a2\u0001\u0000\u0000\u0000\u00a4\u00a5\u0001\u0000\u0000"+
		"\u0000\u00a5\u00a3\u0001\u0000\u0000\u0000\u00a5\u00a6\u0001\u0000\u0000"+
		"\u0000\u00a6\u00a7\u0001\u0000\u0000\u0000\u00a7\u00a8\u0005\u0004\u0000"+
		"\u0000\u00a8\u00b0\u0001\u0000\u0000\u0000\u00a9\u00aa\u0005\u0005\u0000"+
		"\u0000\u00aa\u00ab\u0003>\u001f\u0000\u00ab\u00ac\u0005\u0004\u0000\u0000"+
		"\u00ac\u00b0\u0001\u0000\u0000\u0000\u00ad\u00ae\u0005\u0006\u0000\u0000"+
		"\u00ae\u00b0\u0003\u0006\u0003\u0001\u00af\u0092\u0001\u0000\u0000\u0000"+
		"\u00af\u009c\u0001\u0000\u0000\u0000\u00af\u009f\u0001\u0000\u0000\u0000"+
		"\u00af\u00a9\u0001\u0000\u0000\u0000\u00af\u00ad\u0001\u0000\u0000\u0000"+
		"\u00b0\u00bb\u0001\u0000\u0000\u0000\u00b1\u00b2\n\u0006\u0000\u0000\u00b2"+
		"\u00b3\u0003F#\u0000\u00b3\u00b4\u0003\u0006\u0003\u0007\u00b4\u00ba\u0001"+
		"\u0000\u0000\u0000\u00b5\u00b6\n\u0005\u0000\u0000\u00b6\u00b7\u0003F"+
		"#\u0000\u00b7\u00b8\u0003\u0006\u0003\u0005\u00b8\u00ba\u0001\u0000\u0000"+
		"\u0000\u00b9\u00b1\u0001\u0000\u0000\u0000\u00b9\u00b5\u0001\u0000\u0000"+
		"\u0000\u00ba\u00bd\u0001\u0000\u0000\u0000\u00bb\u00b9\u0001\u0000\u0000"+
		"\u0000\u00bb\u00bc\u0001\u0000\u0000\u0000\u00bc\u0007\u0001\u0000\u0000"+
		"\u0000\u00bd\u00bb\u0001\u0000\u0000\u0000\u00be\u00ca\u0003\f\u0006\u0000"+
		"\u00bf\u00ca\u0003\u000e\u0007\u0000\u00c0\u00ca\u0003\u0012\t\u0000\u00c1"+
		"\u00ca\u0003(\u0014\u0000\u00c2\u00ca\u0003\u0014\n\u0000\u00c3\u00ca"+
		"\u0003d2\u0000\u00c4\u00c5\u0005\u0007\u0000\u0000\u00c5\u00c6\u0003\u0006"+
		"\u0003\u0000\u00c6\u00c7\u0005\b\u0000\u0000\u00c7\u00ca\u0001\u0000\u0000"+
		"\u0000\u00c8\u00ca\u0003h4\u0000\u00c9\u00be\u0001\u0000\u0000\u0000\u00c9"+
		"\u00bf\u0001\u0000\u0000\u0000\u00c9\u00c0\u0001\u0000\u0000\u0000\u00c9"+
		"\u00c1\u0001\u0000\u0000\u0000\u00c9\u00c2\u0001\u0000\u0000\u0000\u00c9"+
		"\u00c3\u0001\u0000\u0000\u0000\u00c9\u00c4\u0001\u0000\u0000\u0000\u00c9"+
		"\u00c8\u0001\u0000\u0000\u0000\u00ca\u00ce\u0001\u0000\u0000\u0000\u00cb"+
		"\u00cd\u0003\n\u0005\u0000\u00cc\u00cb\u0001\u0000\u0000\u0000\u00cd\u00d0"+
		"\u0001\u0000\u0000\u0000\u00ce\u00cc\u0001\u0000\u0000\u0000\u00ce\u00cf"+
		"\u0001\u0000\u0000\u0000\u00cf\t\u0001\u0000\u0000\u0000\u00d0\u00ce\u0001"+
		"\u0000\u0000\u0000\u00d1\u00d2\u0005\t\u0000\u0000\u00d2\u00d3\u0005t"+
		"\u0000\u0000\u00d3\u00d5\u0005\u0007\u0000\u0000\u00d4\u00d6\u0003b1\u0000"+
		"\u00d5\u00d4\u0001\u0000\u0000\u0000\u00d5\u00d6\u0001\u0000\u0000\u0000"+
		"\u00d6\u00d7\u0001\u0000\u0000\u0000\u00d7\u00df\u0005\b\u0000\u0000\u00d8"+
		"\u00d9\u0005\n\u0000\u0000\u00d9\u00df\u0005t\u0000\u0000\u00da\u00db"+
		"\u0005\u000b\u0000\u0000\u00db\u00dc\u0003\u0006\u0003\u0000\u00dc\u00dd"+
		"\u0005\f\u0000\u0000\u00dd\u00df\u0001\u0000\u0000\u0000\u00de\u00d1\u0001"+
		"\u0000\u0000\u0000\u00de\u00d8\u0001\u0000\u0000\u0000\u00de\u00da\u0001"+
		"\u0000\u0000\u0000\u00df\u000b\u0001\u0000\u0000\u0000\u00e0\u00e1\u0007"+
		"\u0000\u0000\u0000\u00e1\r\u0001\u0000\u0000\u0000\u00e2\u00e3\u0006\u0007"+
		"\uffff\uffff\u0000\u00e3\u00e4\u0005t\u0000\u0000\u00e4\u00f1\u0001\u0000"+
		"\u0000\u0000\u00e5\u00e6\n\u0003\u0000\u0000\u00e6\u00f0\u0003\u0010\b"+
		"\u0000\u00e7\u00e8\n\u0002\u0000\u0000\u00e8\u00e9\u0005\u000b\u0000\u0000"+
		"\u00e9\u00ea\u0003\u0006\u0003\u0000\u00ea\u00eb\u0005\f\u0000\u0000\u00eb"+
		"\u00f0\u0001\u0000\u0000\u0000\u00ec\u00ed\n\u0001\u0000\u0000\u00ed\u00ee"+
		"\u0005\n\u0000\u0000\u00ee\u00f0\u0005t\u0000\u0000\u00ef\u00e5\u0001"+
		"\u0000\u0000\u0000\u00ef\u00e7\u0001\u0000\u0000\u0000\u00ef\u00ec\u0001"+
		"\u0000\u0000\u0000\u00f0\u00f3\u0001\u0000\u0000\u0000\u00f1\u00ef\u0001"+
		"\u0000\u0000\u0000\u00f1\u00f2\u0001\u0000\u0000\u0000\u00f2\u000f\u0001"+
		"\u0000\u0000\u0000\u00f3\u00f1\u0001\u0000\u0000\u0000\u00f4\u00f5\u0005"+
		"\r\u0000\u0000\u00f5\u00fb\u0005t\u0000\u0000\u00f6\u00f7\u0005\u000e"+
		"\u0000\u0000\u00f7\u00f8\u0003\u0006\u0003\u0000\u00f8\u00f9\u0005\f\u0000"+
		"\u0000\u00f9\u00fb\u0001\u0000\u0000\u0000\u00fa\u00f4\u0001\u0000\u0000"+
		"\u0000\u00fa\u00f6\u0001\u0000\u0000\u0000\u00fb\u0011\u0001\u0000\u0000"+
		"\u0000\u00fc\u00fd\u0003\u000e\u0007\u0000\u00fd\u0106\u0005\u0007\u0000"+
		"\u0000\u00fe\u0103\u0003\u0006\u0003\u0000\u00ff\u0100\u0005\u000f\u0000"+
		"\u0000\u0100\u0102\u0003\u0006\u0003\u0000\u0101\u00ff\u0001\u0000\u0000"+
		"\u0000\u0102\u0105\u0001\u0000\u0000\u0000\u0103\u0101\u0001\u0000\u0000"+
		"\u0000\u0103\u0104\u0001\u0000\u0000\u0000\u0104\u0107\u0001\u0000\u0000"+
		"\u0000\u0105\u0103\u0001\u0000\u0000\u0000\u0106\u00fe\u0001\u0000\u0000"+
		"\u0000\u0106\u0107\u0001\u0000\u0000\u0000\u0107\u0108\u0001\u0000\u0000"+
		"\u0000\u0108\u0109\u0005\b\u0000\u0000\u0109\u011b\u0001\u0000\u0000\u0000"+
		"\u010a\u010b\u0003\u000e\u0007\u0000\u010b\u010c\u0005\t\u0000\u0000\u010c"+
		"\u010d\u0005t\u0000\u0000\u010d\u0116\u0005\u0007\u0000\u0000\u010e\u0113"+
		"\u0003\u0006\u0003\u0000\u010f\u0110\u0005\u000f\u0000\u0000\u0110\u0112"+
		"\u0003\u0006\u0003\u0000\u0111\u010f\u0001\u0000\u0000\u0000\u0112\u0115"+
		"\u0001\u0000\u0000\u0000\u0113\u0111\u0001\u0000\u0000\u0000\u0113\u0114"+
		"\u0001\u0000\u0000\u0000\u0114\u0117\u0001\u0000\u0000\u0000\u0115\u0113"+
		"\u0001\u0000\u0000\u0000\u0116\u010e\u0001\u0000\u0000\u0000\u0116\u0117"+
		"\u0001\u0000\u0000\u0000\u0117\u0118\u0001\u0000\u0000\u0000\u0118\u0119"+
		"\u0005\b\u0000\u0000\u0119\u011b\u0001\u0000\u0000\u0000\u011a\u00fc\u0001"+
		"\u0000\u0000\u0000\u011a\u010a\u0001\u0000\u0000\u0000\u011b\u0013\u0001"+
		"\u0000\u0000\u0000\u011c\u0125\u0005\u0010\u0000\u0000\u011d\u0122\u0003"+
		"$\u0012\u0000\u011e\u011f\u0005\u000f\u0000\u0000\u011f\u0121\u0003$\u0012"+
		"\u0000\u0120\u011e\u0001\u0000\u0000\u0000\u0121\u0124\u0001\u0000\u0000"+
		"\u0000\u0122\u0120\u0001\u0000\u0000\u0000\u0122\u0123\u0001\u0000\u0000"+
		"\u0000\u0123\u0126\u0001\u0000\u0000\u0000\u0124\u0122\u0001\u0000\u0000"+
		"\u0000\u0125\u011d\u0001\u0000\u0000\u0000\u0125\u0126\u0001\u0000\u0000"+
		"\u0000\u0126\u0129\u0001\u0000\u0000\u0000\u0127\u0128\u0005\u000f\u0000"+
		"\u0000\u0128\u012a\u0003\u0016\u000b\u0000\u0129\u0127\u0001\u0000\u0000"+
		"\u0000\u0129\u012a\u0001\u0000\u0000\u0000\u012a\u012b\u0001\u0000\u0000"+
		"\u0000\u012b\u012c\u0005\u0011\u0000\u0000\u012c\u0015\u0001\u0000\u0000"+
		"\u0000\u012d\u012e\u0005\u0012\u0000\u0000\u012e\u012f\u0005\u0013\u0000"+
		"\u0000\u012f\u0135\u0003\u0006\u0003\u0000\u0130\u0131\u0005\u0010\u0000"+
		"\u0000\u0131\u0132\u0003\u0018\f\u0000\u0132\u0133\u0005\u0011\u0000\u0000"+
		"\u0133\u0135\u0001\u0000\u0000\u0000\u0134\u012d\u0001\u0000\u0000\u0000"+
		"\u0134\u0130\u0001\u0000\u0000\u0000\u0135\u0017\u0001\u0000\u0000\u0000"+
		"\u0136\u013b\u0003\"\u0011\u0000\u0137\u0138\u0005\u000f\u0000\u0000\u0138"+
		"\u013a\u0003\"\u0011\u0000\u0139\u0137\u0001\u0000\u0000\u0000\u013a\u013d"+
		"\u0001\u0000\u0000\u0000\u013b\u0139\u0001\u0000\u0000\u0000\u013b\u013c"+
		"\u0001\u0000\u0000\u0000\u013c\u0019\u0001\u0000\u0000\u0000\u013d\u013b"+
		"\u0001\u0000\u0000\u0000\u013e\u013f\u0005\u0014\u0000\u0000\u013f\u0140"+
		"\u0005t\u0000\u0000\u0140\u0141\u0005\u0014\u0000\u0000\u0141\u001b\u0001"+
		"\u0000\u0000\u0000\u0142\u0145\u0003\u001e\u000f\u0000\u0143\u0144\u0005"+
		"\u0015\u0000\u0000\u0144\u0146\u0003\u0006\u0003\u0000\u0145\u0143\u0001"+
		"\u0000\u0000\u0000\u0145\u0146\u0001\u0000\u0000\u0000\u0146\u0147\u0001"+
		"\u0000\u0000\u0000\u0147\u0148\u0005\u0016\u0000\u0000\u0148\u0149\u0003"+
		"\u0006\u0003\u0000\u0149\u001d\u0001\u0000\u0000\u0000\u014a\u014b\u0006"+
		"\u000f\uffff\uffff\u0000\u014b\u015a\u0003\f\u0006\u0000\u014c\u015a\u0003"+
		"\u0014\n\u0000\u014d\u015a\u0005t\u0000\u0000\u014e\u014f\u0005\u0010"+
		"\u0000\u0000\u014f\u0154\u0003 \u0010\u0000\u0150\u0151\u0005\u000f\u0000"+
		"\u0000\u0151\u0153\u0003 \u0010\u0000\u0152\u0150\u0001\u0000\u0000\u0000"+
		"\u0153\u0156\u0001\u0000\u0000\u0000\u0154\u0152\u0001\u0000\u0000\u0000"+
		"\u0154\u0155\u0001\u0000\u0000\u0000\u0155\u0157\u0001\u0000\u0000\u0000"+
		"\u0156\u0154\u0001\u0000\u0000\u0000\u0157\u0158\u0005\u0011\u0000\u0000"+
		"\u0158\u015a\u0001\u0000\u0000\u0000\u0159\u014a\u0001\u0000\u0000\u0000"+
		"\u0159\u014c\u0001\u0000\u0000\u0000\u0159\u014d\u0001\u0000\u0000\u0000"+
		"\u0159\u014e\u0001\u0000\u0000\u0000\u015a\u0160\u0001\u0000\u0000\u0000"+
		"\u015b\u015c\n\u0002\u0000\u0000\u015c\u015d\u0005\u0017\u0000\u0000\u015d"+
		"\u015f\u0003\u001e\u000f\u0003\u015e\u015b\u0001\u0000\u0000\u0000\u015f"+
		"\u0162\u0001\u0000\u0000\u0000\u0160\u015e\u0001\u0000\u0000\u0000\u0160"+
		"\u0161\u0001\u0000\u0000\u0000\u0161\u001f\u0001\u0000\u0000\u0000\u0162"+
		"\u0160\u0001\u0000\u0000\u0000\u0163\u0164\u0005t\u0000\u0000\u0164\u0165"+
		"\u0005\u0013\u0000\u0000\u0165\u016e\u0003\u001e\u000f\u0000\u0166\u0167"+
		"\u0005\u000b\u0000\u0000\u0167\u0168\u0003\u0006\u0003\u0000\u0168\u0169"+
		"\u0005\f\u0000\u0000\u0169\u016a\u0005\u0013\u0000\u0000\u016a\u016b\u0003"+
		"\u001e\u000f\u0000\u016b\u016e\u0001\u0000\u0000\u0000\u016c\u016e\u0003"+
		"\u001e\u000f\u0000\u016d\u0163\u0001\u0000\u0000\u0000\u016d\u0166\u0001"+
		"\u0000\u0000\u0000\u016d\u016c\u0001\u0000\u0000\u0000\u016e!\u0001\u0000"+
		"\u0000\u0000\u016f\u0170\u0007\u0001\u0000\u0000\u0170#\u0001\u0000\u0000"+
		"\u0000\u0171\u0172\u0005t\u0000\u0000\u0172\u0173\u0005\u0013\u0000\u0000"+
		"\u0173\u017f\u0003\u0006\u0003\u0000\u0174\u0175\u0005\u000b\u0000\u0000"+
		"\u0175\u0176\u0003\u0006\u0003\u0000\u0176\u0177\u0005\f\u0000\u0000\u0177"+
		"\u0178\u0005\u0013\u0000\u0000\u0178\u0179\u0003\u0006\u0003\u0000\u0179"+
		"\u017f\u0001\u0000\u0000\u0000\u017a\u017f\u0003\u0006\u0003\u0000\u017b"+
		"\u017c\u0005t\u0000\u0000\u017c\u017d\u0005\t\u0000\u0000\u017d\u017f"+
		"\u0003d2\u0000\u017e\u0171\u0001\u0000\u0000\u0000\u017e\u0174\u0001\u0000"+
		"\u0000\u0000\u017e\u017a\u0001\u0000\u0000\u0000\u017e\u017b\u0001\u0000"+
		"\u0000\u0000\u017f%\u0001\u0000\u0000\u0000\u0180\u0181\u0003\u0006\u0003"+
		"\u0000\u0181\u0182\u0003L&\u0000\u0182\u0183\u0003\u0006\u0003\u0000\u0183"+
		"\u01a5\u0001\u0000\u0000\u0000\u0184\u0185\u0003\u0006\u0003\u0000\u0185"+
		"\u0186\u0003N\'\u0000\u0186\u0187\u0003\u0006\u0003\u0000\u0187\u01a5"+
		"\u0001\u0000\u0000\u0000\u0188\u0189\u0003\u0006\u0003\u0000\u0189\u018a"+
		"\u0003J%\u0000\u018a\u018b\u0003\u0006\u0003\u0000\u018b\u01a5\u0001\u0000"+
		"\u0000\u0000\u018c\u018d\u0003\u0006\u0003\u0000\u018d\u018e\u0003H$\u0000"+
		"\u018e\u018f\u0003\u0006\u0003\u0000\u018f\u01a5\u0001\u0000\u0000\u0000"+
		"\u0190\u0191\u0003\u0006\u0003\u0000\u0191\u0192\u0003T*\u0000\u0192\u0193"+
		"\u0003\u0006\u0003\u0000\u0193\u01a5\u0001\u0000\u0000\u0000\u0194\u0195"+
		"\u0003\u0006\u0003\u0000\u0195\u0196\u0003X,\u0000\u0196\u0197\u0003\u0006"+
		"\u0003\u0000\u0197\u01a5\u0001\u0000\u0000\u0000\u0198\u0199\u0003\u0006"+
		"\u0003\u0000\u0199\u019a\u0003\\.\u0000\u019a\u019b\u0003\u0006\u0003"+
		"\u0000\u019b\u01a5\u0001\u0000\u0000\u0000\u019c\u019d\u0003\u0006\u0003"+
		"\u0000\u019d\u019e\u0003^/\u0000\u019e\u019f\u0003\u0006\u0003\u0000\u019f"+
		"\u01a5\u0001\u0000\u0000\u0000\u01a0\u01a1\u0003\u0006\u0003\u0000\u01a1"+
		"\u01a2\u0003Z-\u0000\u01a2\u01a3\u0003\u0006\u0003\u0000\u01a3\u01a5\u0001"+
		"\u0000\u0000\u0000\u01a4\u0180\u0001\u0000\u0000\u0000\u01a4\u0184\u0001"+
		"\u0000\u0000\u0000\u01a4\u0188\u0001\u0000\u0000\u0000\u01a4\u018c\u0001"+
		"\u0000\u0000\u0000\u01a4\u0190\u0001\u0000\u0000\u0000\u01a4\u0194\u0001"+
		"\u0000\u0000\u0000\u01a4\u0198\u0001\u0000\u0000\u0000\u01a4\u019c\u0001"+
		"\u0000\u0000\u0000\u01a4\u01a0\u0001\u0000\u0000\u0000\u01a5\'\u0001\u0000"+
		"\u0000\u0000\u01a6\u01a7\u0003R)\u0000\u01a7\u01a8\u0003\u0006\u0003\u0000"+
		"\u01a8\u01ac\u0001\u0000\u0000\u0000\u01a9\u01aa\u0005(\u0000\u0000\u01aa"+
		"\u01ac\u0003\u0006\u0003\u0000\u01ab\u01a6\u0001\u0000\u0000\u0000\u01ab"+
		"\u01a9\u0001\u0000\u0000\u0000\u01ac)\u0001\u0000\u0000\u0000\u01ad\u01b6"+
		"\u0003,\u0016\u0000\u01ae\u01b6\u0003.\u0017\u0000\u01af\u01b6\u00030"+
		"\u0018\u0000\u01b0\u01b6\u00032\u0019\u0000\u01b1\u01b6\u00034\u001a\u0000"+
		"\u01b2\u01b6\u00036\u001b\u0000\u01b3\u01b6\u00038\u001c\u0000\u01b4\u01b6"+
		"\u0003:\u001d\u0000\u01b5\u01ad\u0001\u0000\u0000\u0000\u01b5\u01ae\u0001"+
		"\u0000\u0000\u0000\u01b5\u01af\u0001\u0000\u0000\u0000\u01b5\u01b0\u0001"+
		"\u0000\u0000\u0000\u01b5\u01b1\u0001\u0000\u0000\u0000\u01b5\u01b2\u0001"+
		"\u0000\u0000\u0000\u01b5\u01b3\u0001\u0000\u0000\u0000\u01b5\u01b4\u0001"+
		"\u0000\u0000\u0000\u01b6+\u0001\u0000\u0000\u0000\u01b7\u01b8\u0005)\u0000"+
		"\u0000\u01b8\u01b9\u0003\u0006\u0003\u0000\u01b9\u01ba\u0005*\u0000\u0000"+
		"\u01ba\u01c2\u0003>\u001f\u0000\u01bb\u01bc\u0005+\u0000\u0000\u01bc\u01bd"+
		"\u0003\u0006\u0003\u0000\u01bd\u01be\u0005*\u0000\u0000\u01be\u01bf\u0003"+
		">\u001f\u0000\u01bf\u01c1\u0001\u0000\u0000\u0000\u01c0\u01bb\u0001\u0000"+
		"\u0000\u0000\u01c1\u01c4\u0001\u0000\u0000\u0000\u01c2\u01c0\u0001\u0000"+
		"\u0000\u0000\u01c2\u01c3\u0001\u0000\u0000\u0000\u01c3\u01c7\u0001\u0000"+
		"\u0000\u0000\u01c4\u01c2\u0001\u0000\u0000\u0000\u01c5\u01c6\u0005,\u0000"+
		"\u0000\u01c6\u01c8\u0003>\u001f\u0000\u01c7\u01c5\u0001\u0000\u0000\u0000"+
		"\u01c7\u01c8\u0001\u0000\u0000\u0000\u01c8\u01c9\u0001\u0000\u0000\u0000"+
		"\u01c9\u01ca\u0005\u0004\u0000\u0000\u01ca-\u0001\u0000\u0000\u0000\u01cb"+
		"\u01cc\u0005-\u0000\u0000\u01cc\u01cd\u0003\u0006\u0003\u0000\u01cd\u01ce"+
		"\u0005.\u0000\u0000\u01ce\u01cf\u0003>\u001f\u0000\u01cf\u01d0\u0005\u0004"+
		"\u0000\u0000\u01d0/\u0001\u0000\u0000\u0000\u01d1\u01d2\u0005/\u0000\u0000"+
		"\u01d2\u01d3\u0003>\u001f\u0000\u01d3\u01d4\u00050\u0000\u0000\u01d4\u01d5"+
		"\u0003\u0006\u0003\u0000\u01d51\u0001\u0000\u0000\u0000\u01d6\u01d7\u0005"+
		"1\u0000\u0000\u01d7\u01d8\u0005t\u0000\u0000\u01d8\u01d9\u0005\u0013\u0000"+
		"\u0000\u01d9\u01da\u0003\u0006\u0003\u0000\u01da\u01db\u0005\u000f\u0000"+
		"\u0000\u01db\u01de\u0003\u0006\u0003\u0000\u01dc\u01dd\u0005\u000f\u0000"+
		"\u0000\u01dd\u01df\u0003\u0006\u0003\u0000\u01de\u01dc\u0001\u0000\u0000"+
		"\u0000\u01de\u01df\u0001\u0000\u0000\u0000\u01df\u01e0\u0001\u0000\u0000"+
		"\u0000\u01e0\u01e1\u0005.\u0000\u0000\u01e1\u01e2\u0003>\u001f\u0000\u01e2"+
		"\u01e3\u0005\u0004\u0000\u0000\u01e3\u01ed\u0001\u0000\u0000\u0000\u01e4"+
		"\u01e5\u00051\u0000\u0000\u01e5\u01e6\u0003`0\u0000\u01e6\u01e7\u0005"+
		"2\u0000\u0000\u01e7\u01e8\u0003b1\u0000\u01e8\u01e9\u0005.\u0000\u0000"+
		"\u01e9\u01ea\u0003>\u001f\u0000\u01ea\u01eb\u0005\u0004\u0000\u0000\u01eb"+
		"\u01ed\u0001\u0000\u0000\u0000\u01ec\u01d6\u0001\u0000\u0000\u0000\u01ec"+
		"\u01e4\u0001\u0000\u0000\u0000\u01ed3\u0001\u0000\u0000\u0000\u01ee\u01ef"+
		"\u00053\u0000\u0000\u01ef5\u0001\u0000\u0000\u0000\u01f0\u01f1\u00054"+
		"\u0000\u0000\u01f1\u01f2\u0005t\u0000\u0000\u01f27\u0001\u0000\u0000\u0000"+
		"\u01f3\u01f4\u00055\u0000\u0000\u01f4\u01f5\u0005\n\u0000\u0000\u01f5"+
		"\u01f6\u0007\u0002\u0000\u0000\u01f69\u0001\u0000\u0000\u0000\u01f7\u01f9"+
		"\u0007\u0003\u0000\u0000\u01f8\u01fa\u0007\u0004\u0000\u0000\u01f9\u01f8"+
		"\u0001\u0000\u0000\u0000\u01f9\u01fa\u0001\u0000\u0000\u0000\u01fa\u01fc"+
		"\u0001\u0000\u0000\u0000\u01fb\u01fd\u0005t\u0000\u0000\u01fc\u01fb\u0001"+
		"\u0000\u0000\u0000\u01fc\u01fd\u0001\u0000\u0000\u0000\u01fd\u01fe\u0001"+
		"\u0000\u0000\u0000\u01fe\u0201\u0005\u0007\u0000\u0000\u01ff\u0202\u0003"+
		"b1\u0000\u0200\u0202\u0003<\u001e\u0000\u0201\u01ff\u0001\u0000\u0000"+
		"\u0000\u0201\u0200\u0001\u0000\u0000\u0000\u0202\u0203\u0001\u0000\u0000"+
		"\u0000\u0203\u0204\u0005\b\u0000\u0000\u0204;\u0001\u0000\u0000\u0000"+
		"\u0205\u0206\u0005t\u0000\u0000\u0206\u0207\u0005\u0013\u0000\u0000\u0207"+
		"\u020e\u0003\u0006\u0003\u0000\u0208\u0209\u0005\u000f\u0000\u0000\u0209"+
		"\u020a\u0005t\u0000\u0000\u020a\u020b\u0005\u0013\u0000\u0000\u020b\u020d"+
		"\u0003\u0006\u0003\u0000\u020c\u0208\u0001\u0000\u0000\u0000\u020d\u0210"+
		"\u0001\u0000\u0000\u0000\u020e\u020c\u0001\u0000\u0000\u0000\u020e\u020f"+
		"\u0001\u0000\u0000\u0000\u020f=\u0001\u0000\u0000\u0000\u0210\u020e\u0001"+
		"\u0000\u0000\u0000\u0211\u0213\u0003\u0002\u0001\u0000\u0212\u0211\u0001"+
		"\u0000\u0000\u0000\u0213\u0216\u0001\u0000\u0000\u0000\u0214\u0212\u0001"+
		"\u0000\u0000\u0000\u0214\u0215\u0001\u0000\u0000\u0000\u0215?\u0001\u0000"+
		"\u0000\u0000\u0216\u0214\u0001\u0000\u0000\u0000\u0217\u0219\u0005?\u0000"+
		"\u0000\u0218\u021a\u0003j5\u0000\u0219\u0218\u0001\u0000\u0000\u0000\u0219"+
		"\u021a\u0001\u0000\u0000\u0000\u021a\u021b\u0001\u0000\u0000\u0000\u021b"+
		"\u021e\u0005t\u0000\u0000\u021c\u021d\u0005\u0013\u0000\u0000\u021d\u021f"+
		"\u0003\u0006\u0003\u0000\u021e\u021c\u0001\u0000\u0000\u0000\u021e\u021f"+
		"\u0001\u0000\u0000\u0000\u021f\u023e\u0001\u0000\u0000\u0000\u0220\u0221"+
		"\u0005?\u0000\u0000\u0221\u0226\u0005t\u0000\u0000\u0222\u0223\u0005\u000f"+
		"\u0000\u0000\u0223\u0225\u0005t\u0000\u0000\u0224\u0222\u0001\u0000\u0000"+
		"\u0000\u0225\u0228\u0001\u0000\u0000\u0000\u0226\u0224\u0001\u0000\u0000"+
		"\u0000\u0226\u0227\u0001\u0000\u0000\u0000\u0227\u0229\u0001\u0000\u0000"+
		"\u0000\u0228\u0226\u0001\u0000\u0000\u0000\u0229\u022a\u0005\u0013\u0000"+
		"\u0000\u022a\u023e\u0003b1\u0000\u022b\u022c\u0005?\u0000\u0000\u022c"+
		"\u022d\u0005@\u0000\u0000\u022d\u022e\u0005t\u0000\u0000\u022e\u0237\u0005"+
		"\u0007\u0000\u0000\u022f\u0234\u0005t\u0000\u0000\u0230\u0231\u0005\u000f"+
		"\u0000\u0000\u0231\u0233\u0005t\u0000\u0000\u0232\u0230\u0001\u0000\u0000"+
		"\u0000\u0233\u0236\u0001\u0000\u0000\u0000\u0234\u0232\u0001\u0000\u0000"+
		"\u0000\u0234\u0235\u0001\u0000\u0000\u0000\u0235\u0238\u0001\u0000\u0000"+
		"\u0000\u0236\u0234\u0001\u0000\u0000\u0000\u0237\u022f\u0001\u0000\u0000"+
		"\u0000\u0237\u0238\u0001\u0000\u0000\u0000\u0238\u0239\u0001\u0000\u0000"+
		"\u0000\u0239\u023a\u0005\b\u0000\u0000\u023a\u023b\u0003>\u001f\u0000"+
		"\u023b\u023c\u0005\u0004\u0000\u0000\u023c\u023e\u0001\u0000\u0000\u0000"+
		"\u023d\u0217\u0001\u0000\u0000\u0000\u023d\u0220\u0001\u0000\u0000\u0000"+
		"\u023d\u022b\u0001\u0000\u0000\u0000\u023eA\u0001\u0000\u0000\u0000\u023f"+
		"\u0244\u0005@\u0000\u0000\u0240\u0241\u0005t\u0000\u0000\u0241\u0245\u0005"+
		"\n\u0000\u0000\u0242\u0243\u0005t\u0000\u0000\u0243\u0245\u0005\t\u0000"+
		"\u0000\u0244\u0240\u0001\u0000\u0000\u0000\u0244\u0242\u0001\u0000\u0000"+
		"\u0000\u0244\u0245\u0001\u0000\u0000\u0000\u0245\u0246\u0001\u0000\u0000"+
		"\u0000\u0246\u0247\u0005t\u0000\u0000\u0247\u0255\u0005\u0007\u0000\u0000"+
		"\u0248\u024d\u0005t\u0000\u0000\u0249\u024a\u0005\u000f\u0000\u0000\u024a"+
		"\u024c\u0005t\u0000\u0000\u024b\u0249\u0001\u0000\u0000\u0000\u024c\u024f"+
		"\u0001\u0000\u0000\u0000\u024d\u024b\u0001\u0000\u0000\u0000\u024d\u024e"+
		"\u0001\u0000\u0000\u0000\u024e\u0252\u0001\u0000\u0000\u0000\u024f\u024d"+
		"\u0001\u0000\u0000\u0000\u0250\u0251\u0005\u000f\u0000\u0000\u0251\u0253"+
		"\u0003V+\u0000\u0252\u0250\u0001\u0000\u0000\u0000\u0252\u0253\u0001\u0000"+
		"\u0000\u0000\u0253\u0256\u0001\u0000\u0000\u0000\u0254\u0256\u0003V+\u0000"+
		"\u0255\u0248\u0001\u0000\u0000\u0000\u0255\u0254\u0001\u0000\u0000\u0000"+
		"\u0255\u0256\u0001\u0000\u0000\u0000\u0256\u0257\u0001\u0000\u0000\u0000"+
		"\u0257\u0258\u0005\b\u0000\u0000\u0258\u0259\u0003>\u001f\u0000\u0259"+
		"\u025a\u0005\u0004\u0000\u0000\u025aC\u0001\u0000\u0000\u0000\u025b\u0265"+
		"\u0005A\u0000\u0000\u025c\u0261\u0003\u0006\u0003\u0000\u025d\u025e\u0005"+
		"\u000f\u0000\u0000\u025e\u0260\u0003\u0006\u0003\u0000\u025f\u025d\u0001"+
		"\u0000\u0000\u0000\u0260\u0263\u0001\u0000\u0000\u0000\u0261\u025f\u0001"+
		"\u0000\u0000\u0000\u0261\u0262\u0001\u0000\u0000\u0000\u0262\u0266\u0001"+
		"\u0000\u0000\u0000\u0263\u0261\u0001\u0000\u0000\u0000\u0264\u0266\u0003"+
		"\u0012\t\u0000\u0265\u025c\u0001\u0000\u0000\u0000\u0265\u0264\u0001\u0000"+
		"\u0000\u0000\u0265\u0266\u0001\u0000\u0000\u0000\u0266E\u0001\u0000\u0000"+
		"\u0000\u0267\u0275\u0003H$\u0000\u0268\u0275\u0003J%\u0000\u0269\u0275"+
		"\u0003L&\u0000\u026a\u0275\u0003N\'\u0000\u026b\u0275\u0003P(\u0000\u026c"+
		"\u0275\u0003R)\u0000\u026d\u0275\u0003T*\u0000\u026e\u0275\u0003X,\u0000"+
		"\u026f\u0275\u0003Z-\u0000\u0270\u0275\u0003\\.\u0000\u0271\u0275\u0003"+
		"^/\u0000\u0272\u0275\u0005\u0016\u0000\u0000\u0273\u0275\u0003\u0010\b"+
		"\u0000\u0274\u0267\u0001\u0000\u0000\u0000\u0274\u0268\u0001\u0000\u0000"+
		"\u0000\u0274\u0269\u0001\u0000\u0000\u0000\u0274\u026a\u0001\u0000\u0000"+
		"\u0000\u0274\u026b\u0001\u0000\u0000\u0000\u0274\u026c\u0001\u0000\u0000"+
		"\u0000\u0274\u026d\u0001\u0000\u0000\u0000\u0274\u026e\u0001\u0000\u0000"+
		"\u0000\u0274\u026f\u0001\u0000\u0000\u0000\u0274\u0270\u0001\u0000\u0000"+
		"\u0000\u0274\u0271\u0001\u0000\u0000\u0000\u0274\u0272\u0001\u0000\u0000"+
		"\u0000\u0274\u0273\u0001\u0000\u0000\u0000\u0275G\u0001\u0000\u0000\u0000"+
		"\u0276\u0277\u0007\u0005\u0000\u0000\u0277I\u0001\u0000\u0000\u0000\u0278"+
		"\u0279\u0007\u0006\u0000\u0000\u0279K\u0001\u0000\u0000\u0000\u027a\u027b"+
		"\u0007\u0007\u0000\u0000\u027bM\u0001\u0000\u0000\u0000\u027c\u027d\u0007"+
		"\b\u0000\u0000\u027dO\u0001\u0000\u0000\u0000\u027e\u027f\u0007\t\u0000"+
		"\u0000\u027fQ\u0001\u0000\u0000\u0000\u0280\u0281\u0007\n\u0000\u0000"+
		"\u0281S\u0001\u0000\u0000\u0000\u0282\u0283\u0005_\u0000\u0000\u0283U"+
		"\u0001\u0000\u0000\u0000\u0284\u0285\u0005`\u0000\u0000\u0285W\u0001\u0000"+
		"\u0000\u0000\u0286\u0287\u0007\u000b\u0000\u0000\u0287Y\u0001\u0000\u0000"+
		"\u0000\u0288\u0289\u0007\f\u0000\u0000\u0289[\u0001\u0000\u0000\u0000"+
		"\u028a\u028b\u0005g\u0000\u0000\u028b]\u0001\u0000\u0000\u0000\u028c\u028d"+
		"\u0007\r\u0000\u0000\u028d_\u0001\u0000\u0000\u0000\u028e\u0293\u0005"+
		"t\u0000\u0000\u028f\u0290\u0005\u000f\u0000\u0000\u0290\u0292\u0005t\u0000"+
		"\u0000\u0291\u028f\u0001\u0000\u0000\u0000\u0292\u0295\u0001\u0000\u0000"+
		"\u0000\u0293\u0291\u0001\u0000\u0000\u0000\u0293\u0294\u0001\u0000\u0000"+
		"\u0000\u0294a\u0001\u0000\u0000\u0000\u0295\u0293\u0001\u0000\u0000\u0000"+
		"\u0296\u029b\u0003\u0006\u0003\u0000\u0297\u0298\u0005\u000f\u0000\u0000"+
		"\u0298\u029a\u0003\u0006\u0003\u0000\u0299\u0297\u0001\u0000\u0000\u0000"+
		"\u029a\u029d\u0001\u0000\u0000\u0000\u029b\u0299\u0001\u0000\u0000\u0000"+
		"\u029b\u029c\u0001\u0000\u0000\u0000\u029c\u02a0\u0001\u0000\u0000\u0000"+
		"\u029d\u029b\u0001\u0000\u0000\u0000\u029e\u02a0\u0003V+\u0000\u029f\u0296"+
		"\u0001\u0000\u0000\u0000\u029f\u029e\u0001\u0000\u0000\u0000\u02a0c\u0001"+
		"\u0000\u0000\u0000\u02a1\u02a4\u0005@\u0000\u0000\u02a2\u02a3\u0005t\u0000"+
		"\u0000\u02a3\u02a5\u0005\t\u0000\u0000\u02a4\u02a2\u0001\u0000\u0000\u0000"+
		"\u02a4\u02a5\u0001\u0000\u0000\u0000\u02a5\u02a6\u0001\u0000\u0000\u0000"+
		"\u02a6\u02af\u0005\u0007\u0000\u0000\u02a7\u02ac\u0005t\u0000\u0000\u02a8"+
		"\u02a9\u0005\u000f\u0000\u0000\u02a9\u02ab\u0005t\u0000\u0000\u02aa\u02a8"+
		"\u0001\u0000\u0000\u0000\u02ab\u02ae\u0001\u0000\u0000\u0000\u02ac\u02aa"+
		"\u0001\u0000\u0000\u0000\u02ac\u02ad\u0001\u0000\u0000\u0000\u02ad\u02b0"+
		"\u0001\u0000\u0000\u0000\u02ae\u02ac\u0001\u0000\u0000\u0000\u02af\u02a7"+
		"\u0001\u0000\u0000\u0000\u02af\u02b0\u0001\u0000\u0000\u0000\u02b0\u02b1"+
		"\u0001\u0000\u0000\u0000\u02b1\u02b7\u0005\b\u0000\u0000\u02b2\u02b3\u0003"+
		">\u001f\u0000\u02b3\u02b4\u0005\u0004\u0000\u0000\u02b4\u02b8\u0001\u0000"+
		"\u0000\u0000\u02b5\u02b6\u0005\u0016\u0000\u0000\u02b6\u02b8\u0003\u0006"+
		"\u0003\u0000\u02b7\u02b2\u0001\u0000\u0000\u0000\u02b7\u02b5\u0001\u0000"+
		"\u0000\u0000\u02b8e\u0001\u0000\u0000\u0000\u02b9\u02ba\u0005j\u0000\u0000"+
		"\u02ba\u02bd\u0005\u0007\u0000\u0000\u02bb\u02be\u0005(\u0000\u0000\u02bc"+
		"\u02be\u0003\u0006\u0003\u0000\u02bd\u02bb\u0001\u0000\u0000\u0000\u02bd"+
		"\u02bc\u0001\u0000\u0000\u0000\u02be\u02bf\u0001\u0000\u0000\u0000\u02bf"+
		"\u02c0\u0005\u000f\u0000\u0000\u02c0\u02c1\u0003\u0006\u0003\u0000\u02c1"+
		"\u02c2\u0005\b\u0000\u0000\u02c2g\u0001\u0000\u0000\u0000\u02c3\u02cc"+
		"\u0005\u0007\u0000\u0000\u02c4\u02c9\u0005t\u0000\u0000\u02c5\u02c6\u0005"+
		"\u000f\u0000\u0000\u02c6\u02c8\u0005t\u0000\u0000\u02c7\u02c5\u0001\u0000"+
		"\u0000\u0000\u02c8\u02cb\u0001\u0000\u0000\u0000\u02c9\u02c7\u0001\u0000"+
		"\u0000\u0000\u02c9\u02ca\u0001\u0000\u0000\u0000\u02ca\u02cd\u0001\u0000"+
		"\u0000\u0000\u02cb\u02c9\u0001\u0000\u0000\u0000\u02cc\u02c4\u0001\u0000"+
		"\u0000\u0000\u02cc\u02cd\u0001\u0000\u0000\u0000\u02cd\u02ce\u0001\u0000"+
		"\u0000\u0000\u02ce\u02cf\u0005\b\u0000\u0000\u02cf\u02d0\u0005\u0016\u0000"+
		"\u0000\u02d0\u02d1\u0003\u0006\u0003\u0000\u02d1i\u0001\u0000\u0000\u0000"+
		"\u02d2\u02d3\u0005\t\u0000\u0000\u02d3\u02d4\u0003l6\u0000\u02d4k\u0001"+
		"\u0000\u0000\u0000\u02d5\u02d6\u00066\uffff\uffff\u0000\u02d6\u02e3\u0005"+
		"k\u0000\u0000\u02d7\u02e3\u0005l\u0000\u0000\u02d8\u02e3\u0005m\u0000"+
		"\u0000\u02d9\u02e3\u0005n\u0000\u0000\u02da\u02e3\u0005@\u0000\u0000\u02db"+
		"\u02e3\u0005o\u0000\u0000\u02dc\u02dd\u0005q\u0000\u0000\u02dd\u02de\u0003"+
		"l6\u0000\u02de\u02df\u0005\u000f\u0000\u0000\u02df\u02e0\u0003l6\u0000"+
		"\u02e0\u02e1\u0005H\u0000\u0000\u02e1\u02e3\u0001\u0000\u0000\u0000\u02e2"+
		"\u02d5\u0001\u0000\u0000\u0000\u02e2\u02d7\u0001\u0000\u0000\u0000\u02e2"+
		"\u02d8\u0001\u0000\u0000\u0000\u02e2\u02d9\u0001\u0000\u0000\u0000\u02e2"+
		"\u02da\u0001\u0000\u0000\u0000\u02e2\u02db\u0001\u0000\u0000\u0000\u02e2"+
		"\u02dc\u0001\u0000\u0000\u0000\u02e3\u02e8\u0001\u0000\u0000\u0000\u02e4"+
		"\u02e5\n\u0002\u0000\u0000\u02e5\u02e7\u0005p\u0000\u0000\u02e6\u02e4"+
		"\u0001\u0000\u0000\u0000\u02e7\u02ea\u0001\u0000\u0000\u0000\u02e8\u02e6"+
		"\u0001\u0000\u0000\u0000\u02e8\u02e9\u0001\u0000\u0000\u0000\u02e9m\u0001"+
		"\u0000\u0000\u0000\u02ea\u02e8\u0001\u0000\u0000\u0000\u02eb\u02ec\u0004"+
		"7\u0007\u0000\u02ec\u02f2\u0003p8\u0000\u02ed\u02ee\u00047\b\u0000\u02ee"+
		"\u02f2\u0003r9\u0000\u02ef\u02f0\u00047\t\u0000\u02f0\u02f2\u0003t:\u0000"+
		"\u02f1\u02eb\u0001\u0000\u0000\u0000\u02f1\u02ed\u0001\u0000\u0000\u0000"+
		"\u02f1\u02ef\u0001\u0000\u0000\u0000\u02f2o\u0001\u0000\u0000\u0000\u02f3"+
		"\u02f4\u0003\u0006\u0003\u0000\u02f4\u02f5\u0005\r\u0000\u0000\u02f5\u02f6"+
		"\u0005t\u0000\u0000\u02f6\u02fd\u0001\u0000\u0000\u0000\u02f7\u02f8\u0003"+
		"\u0006\u0003\u0000\u02f8\u02f9\u0005\u000e\u0000\u0000\u02f9\u02fa\u0003"+
		"\u0006\u0003\u0000\u02fa\u02fb\u0005\f\u0000\u0000\u02fb\u02fd\u0001\u0000"+
		"\u0000\u0000\u02fc\u02f3\u0001\u0000\u0000\u0000\u02fc\u02f7\u0001\u0000"+
		"\u0000\u0000\u02fdq\u0001\u0000\u0000\u0000\u02fe\u02ff\u0003\u0006\u0003"+
		"\u0000\u02ff\u0300\u0005r\u0000\u0000\u0300\u0301\u0003\u0006\u0003\u0000"+
		"\u0301s\u0001\u0000\u0000\u0000\u0302\u0303\u0005s\u0000\u0000\u0303\u0309"+
		"\u0005t\u0000\u0000\u0304\u0306\u0005\u0007\u0000\u0000\u0305\u0307\u0003"+
		"b1\u0000\u0306\u0305\u0001\u0000\u0000\u0000\u0306\u0307\u0001\u0000\u0000"+
		"\u0000\u0307\u0308\u0001\u0000\u0000\u0000\u0308\u030a\u0005\b\u0000\u0000"+
		"\u0309\u0304\u0001\u0000\u0000\u0000\u0309\u030a\u0001\u0000\u0000\u0000"+
		"\u030au\u0001\u0000\u0000\u0000Hy\u0087\u008e\u0099\u00a5\u00af\u00b9"+
		"\u00bb\u00c9\u00ce\u00d5\u00de\u00ef\u00f1\u00fa\u0103\u0106\u0113\u0116"+
		"\u011a\u0122\u0125\u0129\u0134\u013b\u0145\u0154\u0159\u0160\u016d\u017e"+
		"\u01a4\u01ab\u01b5\u01c2\u01c7\u01de\u01ec\u01f9\u01fc\u0201\u020e\u0214"+
		"\u0219\u021e\u0226\u0234\u0237\u023d\u0244\u024d\u0252\u0255\u0261\u0265"+
		"\u0274\u0293\u029b\u029f\u02a4\u02ac\u02af\u02b7\u02bd\u02c9\u02cc\u02e2"+
		"\u02e8\u02f1\u02fc\u0306\u0309";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}