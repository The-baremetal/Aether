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
		SEPARATOR=46, KW_IN=47, KW_PRINT=48, KW_AND=49, KW_BREAK=50, KW_DO=51, 
		KW_ELSE=52, KW_ELSEIF=53, KW_END=54, KW_FALSE=55, KW_FOR=56, KW_GOTO=57, 
		KW_IF=58, KW_NIL=59, KW_NOT=60, KW_OR=61, KW_REPEAT=62, KW_RETURN=63, 
		KW_THEN=64, KW_TRUE=65, KW_UNTIL=66, KW_WHILE=67, KW_LOCAL=68, KW_FUNCTION=69, 
		KW_INDEX=70, KW_NEWINDEX=71, KW_MODE=72, KW_PCALL=73, KW_XPCALL=74, KW_COROUTINE=75, 
		KW_CREATE=76, KW_RESUME=77, KW_YIELD=78, KW_STATUS=79, KW_NAN=80, KW_INF=81, 
		NUMBER=82, STRING=83, LETTER=84, DIGIT=85, WS=86, SINGLE_LINE_COMMENT=87, 
		MULTI_LINE_COMMENT=88, VARARG=89;
	public static final int
		RULE_program = 0, RULE_statement = 1, RULE_control_statement = 2, RULE_statement_terminator = 3, 
		RULE_assignment = 4, RULE_expression = 5, RULE_prefix_expression = 6, 
		RULE_primary_expression = 7, RULE_operators = 8, RULE_comparison_operator = 9, 
		RULE_arith_operator = 10, RULE_logical_operator = 11, RULE_bitwise_operator = 12, 
		RULE_concat_operator = 13, RULE_literal = 14, RULE_function_call = 15, 
		RULE_table_insert = 16, RULE_function_declaration = 17, RULE_block = 18, 
		RULE_if_statement = 19, RULE_for_statement = 20, RULE_while_statement = 21, 
		RULE_table = 22, RULE_field = 23, RULE_table_access = 24, RULE_single_line_comment = 25, 
		RULE_print_statement = 26, RULE_identifier = 27, RULE_repeat_statement = 28, 
		RULE_identifier_list = 29, RULE_expression_list = 30, RULE_return_statement = 31, 
		RULE_break_statement = 32, RULE_goto_statement = 33, RULE_label_statement = 34, 
		RULE_function_expression = 35, RULE_method_call = 36, RULE_metatable_field = 37, 
		RULE_metamethod = 38, RULE_coroutine_statement = 39;
	private static String[] makeRuleNames() {
		return new String[] {
			"program", "statement", "control_statement", "statement_terminator", 
			"assignment", "expression", "prefix_expression", "primary_expression", 
			"operators", "comparison_operator", "arith_operator", "logical_operator", 
			"bitwise_operator", "concat_operator", "literal", "function_call", "table_insert", 
			"function_declaration", "block", "if_statement", "for_statement", "while_statement", 
			"table", "field", "table_access", "single_line_comment", "print_statement", 
			"identifier", "repeat_statement", "identifier_list", "expression_list", 
			"return_statement", "break_statement", "goto_statement", "label_statement", 
			"function_expression", "method_call", "metatable_field", "metamethod", 
			"coroutine_statement"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'='", "'('", "')'", "'#'", "'>'", "'<'", "'>='", "'=='", "'<='", 
			"'~='", "'*'", "'/'", "'+'", "'-'", "'//'", "'&'", "'|'", "'~'", "'<<'", 
			"'>>'", "'..'", "':'", "'table.insert'", "','", "'{'", "'}'", "'['", 
			"']'", "'.'", "'--'", "'_'", "'::'", "'__'", "'__add'", "'__sub'", "'__mul'", 
			"'__div'", "'__mod'", "'__pow'", "'__unm'", "'__concat'", "'__len'", 
			"'__eq'", "'__lt'", "'__le'", null, "'in'", "'print'", "'and'", "'break'", 
			"'do'", "'else'", "'elseif'", "'end'", "'false'", "'for'", "'goto'", 
			"'if'", "'nil'", "'not'", "'or'", "'repeat'", "'return'", "'then'", "'true'", 
			"'until'", "'while'", "'local'", "'function'", "'index'", "'newindex'", 
			"'mode'", "'pcall'", "'xpcall'", "'coroutine'", "'create'", "'resume'", 
			"'yield'", "'status'", "'nan'", "'inf'", null, null, null, null, null, 
			null, null, "'...'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, null, null, "SEPARATOR", 
			"KW_IN", "KW_PRINT", "KW_AND", "KW_BREAK", "KW_DO", "KW_ELSE", "KW_ELSEIF", 
			"KW_END", "KW_FALSE", "KW_FOR", "KW_GOTO", "KW_IF", "KW_NIL", "KW_NOT", 
			"KW_OR", "KW_REPEAT", "KW_RETURN", "KW_THEN", "KW_TRUE", "KW_UNTIL", 
			"KW_WHILE", "KW_LOCAL", "KW_FUNCTION", "KW_INDEX", "KW_NEWINDEX", "KW_MODE", 
			"KW_PCALL", "KW_XPCALL", "KW_COROUTINE", "KW_CREATE", "KW_RESUME", "KW_YIELD", 
			"KW_STATUS", "KW_NAN", "KW_INF", "NUMBER", "STRING", "LETTER", "DIGIT", 
			"WS", "SINGLE_LINE_COMMENT", "MULTI_LINE_COMMENT", "VARARG"
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
			setState(83);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & -3493385931619041276L) != 0) || ((((_la - 65)) & ~0x3f) == 0 && ((1L << (_la - 65)) & 1017629L) != 0)) {
				{
				{
				setState(80);
				statement();
				}
				}
				setState(85);
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
	public static class StatementContext extends ParserRuleContext {
		public AssignmentContext assignment() {
			return getRuleContext(AssignmentContext.class,0);
		}
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public Control_statementContext control_statement() {
			return getRuleContext(Control_statementContext.class,0);
		}
		public Function_declarationContext function_declaration() {
			return getRuleContext(Function_declarationContext.class,0);
		}
		public Function_callContext function_call() {
			return getRuleContext(Function_callContext.class,0);
		}
		public Return_statementContext return_statement() {
			return getRuleContext(Return_statementContext.class,0);
		}
		public Break_statementContext break_statement() {
			return getRuleContext(Break_statementContext.class,0);
		}
		public Label_statementContext label_statement() {
			return getRuleContext(Label_statementContext.class,0);
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
			setState(94);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,1,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(86);
				assignment();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(87);
				expression(0);
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(88);
				control_statement();
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(89);
				function_declaration();
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(90);
				function_call();
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(91);
				return_statement();
				}
				break;
			case 7:
				enterOuterAlt(_localctx, 7);
				{
				setState(92);
				break_statement();
				}
				break;
			case 8:
				enterOuterAlt(_localctx, 8);
				{
				setState(93);
				label_statement();
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
	public static class Control_statementContext extends ParserRuleContext {
		public If_statementContext if_statement() {
			return getRuleContext(If_statementContext.class,0);
		}
		public For_statementContext for_statement() {
			return getRuleContext(For_statementContext.class,0);
		}
		public While_statementContext while_statement() {
			return getRuleContext(While_statementContext.class,0);
		}
		public Repeat_statementContext repeat_statement() {
			return getRuleContext(Repeat_statementContext.class,0);
		}
		public Goto_statementContext goto_statement() {
			return getRuleContext(Goto_statementContext.class,0);
		}
		public Coroutine_statementContext coroutine_statement() {
			return getRuleContext(Coroutine_statementContext.class,0);
		}
		public Control_statementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_control_statement; }
	}

	public final Control_statementContext control_statement() throws RecognitionException {
		Control_statementContext _localctx = new Control_statementContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_control_statement);
		try {
			setState(102);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case KW_IF:
				enterOuterAlt(_localctx, 1);
				{
				setState(96);
				if_statement();
				}
				break;
			case KW_FOR:
				enterOuterAlt(_localctx, 2);
				{
				setState(97);
				for_statement();
				}
				break;
			case KW_WHILE:
				enterOuterAlt(_localctx, 3);
				{
				setState(98);
				while_statement();
				}
				break;
			case KW_REPEAT:
				enterOuterAlt(_localctx, 4);
				{
				setState(99);
				repeat_statement();
				}
				break;
			case KW_GOTO:
				enterOuterAlt(_localctx, 5);
				{
				setState(100);
				goto_statement();
				}
				break;
			case KW_COROUTINE:
				enterOuterAlt(_localctx, 6);
				{
				setState(101);
				coroutine_statement();
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
	public static class Statement_terminatorContext extends ParserRuleContext {
		public TerminalNode SEPARATOR() { return getToken(lua_grammar_antlr4Parser.SEPARATOR, 0); }
		public Statement_terminatorContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_statement_terminator; }
	}

	public final Statement_terminatorContext statement_terminator() throws RecognitionException {
		Statement_terminatorContext _localctx = new Statement_terminatorContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_statement_terminator);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(104);
			match(SEPARATOR);
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
	public static class AssignmentContext extends ParserRuleContext {
		public Identifier_listContext identifier_list() {
			return getRuleContext(Identifier_listContext.class,0);
		}
		public Expression_listContext expression_list() {
			return getRuleContext(Expression_listContext.class,0);
		}
		public Statement_terminatorContext statement_terminator() {
			return getRuleContext(Statement_terminatorContext.class,0);
		}
		public TerminalNode KW_LOCAL() { return getToken(lua_grammar_antlr4Parser.KW_LOCAL, 0); }
		public AssignmentContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_assignment; }
	}

	public final AssignmentContext assignment() throws RecognitionException {
		AssignmentContext _localctx = new AssignmentContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_assignment);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(107);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==KW_LOCAL) {
				{
				setState(106);
				match(KW_LOCAL);
				}
			}

			setState(109);
			identifier_list();
			setState(110);
			match(T__0);
			setState(111);
			expression_list();
			setState(112);
			statement_terminator();
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
		public LiteralContext literal() {
			return getRuleContext(LiteralContext.class,0);
		}
		public IdentifierContext identifier() {
			return getRuleContext(IdentifierContext.class,0);
		}
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public Function_callContext function_call() {
			return getRuleContext(Function_callContext.class,0);
		}
		public TableContext table() {
			return getRuleContext(TableContext.class,0);
		}
		public Table_accessContext table_access() {
			return getRuleContext(Table_accessContext.class,0);
		}
		public Function_expressionContext function_expression() {
			return getRuleContext(Function_expressionContext.class,0);
		}
		public OperatorsContext operators() {
			return getRuleContext(OperatorsContext.class,0);
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
		int _startState = 10;
		enterRecursionRule(_localctx, 10, RULE_expression, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(125);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,4,_ctx) ) {
			case 1:
				{
				setState(115);
				literal();
				}
				break;
			case 2:
				{
				setState(116);
				identifier();
				}
				break;
			case 3:
				{
				setState(117);
				match(T__1);
				setState(118);
				expression(0);
				setState(119);
				match(T__2);
				}
				break;
			case 4:
				{
				setState(121);
				function_call();
				}
				break;
			case 5:
				{
				setState(122);
				table();
				}
				break;
			case 6:
				{
				setState(123);
				table_access();
				}
				break;
			case 7:
				{
				setState(124);
				function_expression();
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(133);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,5,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new ExpressionContext(_parentctx, _parentState);
					pushNewRecursionContext(_localctx, _startState, RULE_expression);
					setState(127);
					if (!(precpred(_ctx, 5))) throw new FailedPredicateException(this, "precpred(_ctx, 5)");
					setState(128);
					operators();
					setState(129);
					expression(6);
					}
					} 
				}
				setState(135);
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

	@SuppressWarnings("CheckReturnValue")
	public static class Prefix_expressionContext extends ParserRuleContext {
		public Primary_expressionContext primary_expression() {
			return getRuleContext(Primary_expressionContext.class,0);
		}
		public TerminalNode KW_NOT() { return getToken(lua_grammar_antlr4Parser.KW_NOT, 0); }
		public Prefix_expressionContext prefix_expression() {
			return getRuleContext(Prefix_expressionContext.class,0);
		}
		public Prefix_expressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_prefix_expression; }
	}

	public final Prefix_expressionContext prefix_expression() throws RecognitionException {
		Prefix_expressionContext _localctx = new Prefix_expressionContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_prefix_expression);
		try {
			setState(141);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__1:
			case T__22:
			case T__24:
			case KW_PRINT:
			case KW_FALSE:
			case KW_NIL:
			case KW_TRUE:
			case KW_PCALL:
			case KW_XPCALL:
			case KW_NAN:
			case KW_INF:
			case NUMBER:
			case STRING:
			case LETTER:
				enterOuterAlt(_localctx, 1);
				{
				setState(136);
				primary_expression();
				}
				break;
			case KW_NOT:
				enterOuterAlt(_localctx, 2);
				{
				setState(137);
				match(KW_NOT);
				setState(138);
				prefix_expression();
				}
				break;
			case T__3:
				enterOuterAlt(_localctx, 3);
				{
				setState(139);
				match(T__3);
				setState(140);
				prefix_expression();
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
	public static class Primary_expressionContext extends ParserRuleContext {
		public LiteralContext literal() {
			return getRuleContext(LiteralContext.class,0);
		}
		public IdentifierContext identifier() {
			return getRuleContext(IdentifierContext.class,0);
		}
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public Function_callContext function_call() {
			return getRuleContext(Function_callContext.class,0);
		}
		public TableContext table() {
			return getRuleContext(TableContext.class,0);
		}
		public Table_accessContext table_access() {
			return getRuleContext(Table_accessContext.class,0);
		}
		public Primary_expressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_primary_expression; }
	}

	public final Primary_expressionContext primary_expression() throws RecognitionException {
		Primary_expressionContext _localctx = new Primary_expressionContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_primary_expression);
		try {
			setState(152);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,7,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(143);
				literal();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(144);
				identifier();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(145);
				match(T__1);
				setState(146);
				expression(0);
				setState(147);
				match(T__2);
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(149);
				function_call();
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(150);
				table();
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(151);
				table_access();
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
	public static class OperatorsContext extends ParserRuleContext {
		public Comparison_operatorContext comparison_operator() {
			return getRuleContext(Comparison_operatorContext.class,0);
		}
		public Arith_operatorContext arith_operator() {
			return getRuleContext(Arith_operatorContext.class,0);
		}
		public Logical_operatorContext logical_operator() {
			return getRuleContext(Logical_operatorContext.class,0);
		}
		public Bitwise_operatorContext bitwise_operator() {
			return getRuleContext(Bitwise_operatorContext.class,0);
		}
		public Concat_operatorContext concat_operator() {
			return getRuleContext(Concat_operatorContext.class,0);
		}
		public OperatorsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_operators; }
	}

	public final OperatorsContext operators() throws RecognitionException {
		OperatorsContext _localctx = new OperatorsContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_operators);
		try {
			setState(159);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__4:
			case T__5:
			case T__6:
			case T__7:
			case T__8:
			case T__9:
				enterOuterAlt(_localctx, 1);
				{
				setState(154);
				comparison_operator();
				}
				break;
			case T__10:
			case T__11:
			case T__12:
			case T__13:
			case T__14:
				enterOuterAlt(_localctx, 2);
				{
				setState(155);
				arith_operator();
				}
				break;
			case KW_AND:
			case KW_OR:
				enterOuterAlt(_localctx, 3);
				{
				setState(156);
				logical_operator();
				}
				break;
			case T__15:
			case T__16:
			case T__17:
			case T__18:
			case T__19:
				enterOuterAlt(_localctx, 4);
				{
				setState(157);
				bitwise_operator();
				}
				break;
			case T__20:
				enterOuterAlt(_localctx, 5);
				{
				setState(158);
				concat_operator();
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
	public static class Comparison_operatorContext extends ParserRuleContext {
		public Comparison_operatorContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_comparison_operator; }
	}

	public final Comparison_operatorContext comparison_operator() throws RecognitionException {
		Comparison_operatorContext _localctx = new Comparison_operatorContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_comparison_operator);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(161);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 2016L) != 0)) ) {
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
	public static class Arith_operatorContext extends ParserRuleContext {
		public Arith_operatorContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_arith_operator; }
	}

	public final Arith_operatorContext arith_operator() throws RecognitionException {
		Arith_operatorContext _localctx = new Arith_operatorContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_arith_operator);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(163);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 63488L) != 0)) ) {
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
	public static class Logical_operatorContext extends ParserRuleContext {
		public TerminalNode KW_AND() { return getToken(lua_grammar_antlr4Parser.KW_AND, 0); }
		public TerminalNode KW_OR() { return getToken(lua_grammar_antlr4Parser.KW_OR, 0); }
		public Logical_operatorContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_logical_operator; }
	}

	public final Logical_operatorContext logical_operator() throws RecognitionException {
		Logical_operatorContext _localctx = new Logical_operatorContext(_ctx, getState());
		enterRule(_localctx, 22, RULE_logical_operator);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(165);
			_la = _input.LA(1);
			if ( !(_la==KW_AND || _la==KW_OR) ) {
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
	public static class Bitwise_operatorContext extends ParserRuleContext {
		public Bitwise_operatorContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_bitwise_operator; }
	}

	public final Bitwise_operatorContext bitwise_operator() throws RecognitionException {
		Bitwise_operatorContext _localctx = new Bitwise_operatorContext(_ctx, getState());
		enterRule(_localctx, 24, RULE_bitwise_operator);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(167);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 2031616L) != 0)) ) {
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
	public static class Concat_operatorContext extends ParserRuleContext {
		public Concat_operatorContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_concat_operator; }
	}

	public final Concat_operatorContext concat_operator() throws RecognitionException {
		Concat_operatorContext _localctx = new Concat_operatorContext(_ctx, getState());
		enterRule(_localctx, 26, RULE_concat_operator);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(169);
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
	public static class LiteralContext extends ParserRuleContext {
		public TerminalNode NUMBER() { return getToken(lua_grammar_antlr4Parser.NUMBER, 0); }
		public TerminalNode STRING() { return getToken(lua_grammar_antlr4Parser.STRING, 0); }
		public TerminalNode KW_TRUE() { return getToken(lua_grammar_antlr4Parser.KW_TRUE, 0); }
		public TerminalNode KW_FALSE() { return getToken(lua_grammar_antlr4Parser.KW_FALSE, 0); }
		public TerminalNode KW_NIL() { return getToken(lua_grammar_antlr4Parser.KW_NIL, 0); }
		public TerminalNode KW_NAN() { return getToken(lua_grammar_antlr4Parser.KW_NAN, 0); }
		public TerminalNode KW_INF() { return getToken(lua_grammar_antlr4Parser.KW_INF, 0); }
		public LiteralContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_literal; }
	}

	public final LiteralContext literal() throws RecognitionException {
		LiteralContext _localctx = new LiteralContext(_ctx, getState());
		enterRule(_localctx, 28, RULE_literal);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(171);
			_la = _input.LA(1);
			if ( !(((((_la - 55)) & ~0x3f) == 0 && ((1L << (_la - 55)) & 503317521L) != 0)) ) {
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
	public static class Function_callContext extends ParserRuleContext {
		public Expression_listContext expression_list() {
			return getRuleContext(Expression_listContext.class,0);
		}
		public TerminalNode KW_PCALL() { return getToken(lua_grammar_antlr4Parser.KW_PCALL, 0); }
		public TerminalNode KW_XPCALL() { return getToken(lua_grammar_antlr4Parser.KW_XPCALL, 0); }
		public List<IdentifierContext> identifier() {
			return getRuleContexts(IdentifierContext.class);
		}
		public IdentifierContext identifier(int i) {
			return getRuleContext(IdentifierContext.class,i);
		}
		public Table_accessContext table_access() {
			return getRuleContext(Table_accessContext.class,0);
		}
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public Table_insertContext table_insert() {
			return getRuleContext(Table_insertContext.class,0);
		}
		public TerminalNode KW_PRINT() { return getToken(lua_grammar_antlr4Parser.KW_PRINT, 0); }
		public Method_callContext method_call() {
			return getRuleContext(Method_callContext.class,0);
		}
		public Function_callContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_function_call; }
	}

	public final Function_callContext function_call() throws RecognitionException {
		Function_callContext _localctx = new Function_callContext(_ctx, getState());
		enterRule(_localctx, 30, RULE_function_call);
		int _la;
		try {
			setState(203);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,12,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(173);
				_la = _input.LA(1);
				if ( !(_la==KW_PCALL || _la==KW_XPCALL) ) {
				_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				setState(174);
				match(T__1);
				setState(175);
				expression_list();
				setState(176);
				match(T__2);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(184);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,9,_ctx) ) {
				case 1:
					{
					setState(178);
					identifier();
					}
					break;
				case 2:
					{
					setState(179);
					table_access();
					}
					break;
				case 3:
					{
					setState(180);
					match(T__1);
					setState(181);
					expression(0);
					setState(182);
					match(T__2);
					}
					break;
				}
				setState(188);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==T__21) {
					{
					setState(186);
					match(T__21);
					setState(187);
					identifier();
					}
				}

				setState(190);
				match(T__1);
				setState(192);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 612771024341041156L) != 0) || ((((_la - 65)) & ~0x3f) == 0 && ((1L << (_la - 65)) & 17793809L) != 0)) {
					{
					setState(191);
					expression_list();
					}
				}

				setState(194);
				match(T__2);
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(196);
				table_insert();
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(197);
				match(KW_PRINT);
				setState(198);
				match(T__1);
				setState(199);
				expression_list();
				setState(200);
				match(T__2);
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(202);
				method_call();
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
	public static class Table_insertContext extends ParserRuleContext {
		public IdentifierContext identifier() {
			return getRuleContext(IdentifierContext.class,0);
		}
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public Table_insertContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_table_insert; }
	}

	public final Table_insertContext table_insert() throws RecognitionException {
		Table_insertContext _localctx = new Table_insertContext(_ctx, getState());
		enterRule(_localctx, 32, RULE_table_insert);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(205);
			match(T__22);
			setState(206);
			match(T__1);
			setState(207);
			identifier();
			setState(208);
			match(T__23);
			setState(209);
			expression(0);
			setState(210);
			match(T__2);
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
	public static class Function_declarationContext extends ParserRuleContext {
		public TerminalNode KW_FUNCTION() { return getToken(lua_grammar_antlr4Parser.KW_FUNCTION, 0); }
		public List<IdentifierContext> identifier() {
			return getRuleContexts(IdentifierContext.class);
		}
		public IdentifierContext identifier(int i) {
			return getRuleContext(IdentifierContext.class,i);
		}
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public TerminalNode KW_END() { return getToken(lua_grammar_antlr4Parser.KW_END, 0); }
		public TerminalNode KW_LOCAL() { return getToken(lua_grammar_antlr4Parser.KW_LOCAL, 0); }
		public TerminalNode VARARG() { return getToken(lua_grammar_antlr4Parser.VARARG, 0); }
		public Function_declarationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_function_declaration; }
	}

	public final Function_declarationContext function_declaration() throws RecognitionException {
		Function_declarationContext _localctx = new Function_declarationContext(_ctx, getState());
		enterRule(_localctx, 34, RULE_function_declaration);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(213);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==KW_LOCAL) {
				{
				setState(212);
				match(KW_LOCAL);
				}
			}

			setState(215);
			match(KW_FUNCTION);
			setState(216);
			identifier();
			setState(217);
			match(T__1);
			setState(227);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case LETTER:
				{
				setState(218);
				identifier();
				setState(223);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==T__23) {
					{
					{
					setState(219);
					match(T__23);
					setState(220);
					identifier();
					}
					}
					setState(225);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				}
				break;
			case VARARG:
				{
				setState(226);
				match(VARARG);
				}
				break;
			case T__2:
				break;
			default:
				break;
			}
			setState(229);
			match(T__2);
			setState(230);
			block();
			setState(231);
			match(KW_END);
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
		public List<Statement_terminatorContext> statement_terminator() {
			return getRuleContexts(Statement_terminatorContext.class);
		}
		public Statement_terminatorContext statement_terminator(int i) {
			return getRuleContext(Statement_terminatorContext.class,i);
		}
		public BlockContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_block; }
	}

	public final BlockContext block() throws RecognitionException {
		BlockContext _localctx = new BlockContext(_ctx, getState());
		enterRule(_localctx, 36, RULE_block);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(243); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(236);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==SEPARATOR) {
					{
					{
					setState(233);
					statement_terminator();
					}
					}
					setState(238);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(239);
				statement();
				setState(241);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,17,_ctx) ) {
				case 1:
					{
					setState(240);
					statement_terminator();
					}
					break;
				}
				}
				}
				setState(245); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & -3493315562874863612L) != 0) || ((((_la - 65)) & ~0x3f) == 0 && ((1L << (_la - 65)) & 1017629L) != 0) );
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
	public static class If_statementContext extends ParserRuleContext {
		public TerminalNode KW_IF() { return getToken(lua_grammar_antlr4Parser.KW_IF, 0); }
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public List<TerminalNode> KW_THEN() { return getTokens(lua_grammar_antlr4Parser.KW_THEN); }
		public TerminalNode KW_THEN(int i) {
			return getToken(lua_grammar_antlr4Parser.KW_THEN, i);
		}
		public List<BlockContext> block() {
			return getRuleContexts(BlockContext.class);
		}
		public BlockContext block(int i) {
			return getRuleContext(BlockContext.class,i);
		}
		public TerminalNode KW_END() { return getToken(lua_grammar_antlr4Parser.KW_END, 0); }
		public List<TerminalNode> KW_ELSEIF() { return getTokens(lua_grammar_antlr4Parser.KW_ELSEIF); }
		public TerminalNode KW_ELSEIF(int i) {
			return getToken(lua_grammar_antlr4Parser.KW_ELSEIF, i);
		}
		public TerminalNode KW_ELSE() { return getToken(lua_grammar_antlr4Parser.KW_ELSE, 0); }
		public If_statementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_if_statement; }
	}

	public final If_statementContext if_statement() throws RecognitionException {
		If_statementContext _localctx = new If_statementContext(_ctx, getState());
		enterRule(_localctx, 38, RULE_if_statement);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(247);
			match(KW_IF);
			setState(248);
			expression(0);
			setState(249);
			match(KW_THEN);
			setState(250);
			block();
			setState(258);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==KW_ELSEIF) {
				{
				{
				setState(251);
				match(KW_ELSEIF);
				setState(252);
				expression(0);
				setState(253);
				match(KW_THEN);
				setState(254);
				block();
				}
				}
				setState(260);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(263);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==KW_ELSE) {
				{
				setState(261);
				match(KW_ELSE);
				setState(262);
				block();
				}
			}

			setState(265);
			match(KW_END);
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
	public static class For_statementContext extends ParserRuleContext {
		public TerminalNode KW_FOR() { return getToken(lua_grammar_antlr4Parser.KW_FOR, 0); }
		public IdentifierContext identifier() {
			return getRuleContext(IdentifierContext.class,0);
		}
		public TerminalNode KW_IN() { return getToken(lua_grammar_antlr4Parser.KW_IN, 0); }
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public TerminalNode KW_DO() { return getToken(lua_grammar_antlr4Parser.KW_DO, 0); }
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public TerminalNode KW_END() { return getToken(lua_grammar_antlr4Parser.KW_END, 0); }
		public For_statementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_for_statement; }
	}

	public final For_statementContext for_statement() throws RecognitionException {
		For_statementContext _localctx = new For_statementContext(_ctx, getState());
		enterRule(_localctx, 40, RULE_for_statement);
		int _la;
		try {
			setState(289);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,22,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(267);
				match(KW_FOR);
				setState(268);
				identifier();
				setState(269);
				match(KW_IN);
				setState(270);
				expression(0);
				setState(271);
				match(KW_DO);
				setState(272);
				block();
				setState(273);
				match(KW_END);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(275);
				match(KW_FOR);
				setState(276);
				identifier();
				setState(277);
				match(T__0);
				setState(278);
				expression(0);
				setState(279);
				match(T__23);
				setState(280);
				expression(0);
				setState(283);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==T__23) {
					{
					setState(281);
					match(T__23);
					setState(282);
					expression(0);
					}
				}

				setState(285);
				match(KW_DO);
				setState(286);
				block();
				setState(287);
				match(KW_END);
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
	public static class While_statementContext extends ParserRuleContext {
		public TerminalNode KW_WHILE() { return getToken(lua_grammar_antlr4Parser.KW_WHILE, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode KW_DO() { return getToken(lua_grammar_antlr4Parser.KW_DO, 0); }
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public TerminalNode KW_END() { return getToken(lua_grammar_antlr4Parser.KW_END, 0); }
		public While_statementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_while_statement; }
	}

	public final While_statementContext while_statement() throws RecognitionException {
		While_statementContext _localctx = new While_statementContext(_ctx, getState());
		enterRule(_localctx, 42, RULE_while_statement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(291);
			match(KW_WHILE);
			setState(292);
			expression(0);
			setState(293);
			match(KW_DO);
			setState(294);
			block();
			setState(295);
			match(KW_END);
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
	public static class TableContext extends ParserRuleContext {
		public List<FieldContext> field() {
			return getRuleContexts(FieldContext.class);
		}
		public FieldContext field(int i) {
			return getRuleContext(FieldContext.class,i);
		}
		public Metatable_fieldContext metatable_field() {
			return getRuleContext(Metatable_fieldContext.class,0);
		}
		public TableContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_table; }
	}

	public final TableContext table() throws RecognitionException {
		TableContext _localctx = new TableContext(_ctx, getState());
		enterRule(_localctx, 44, RULE_table);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(297);
			match(T__24);
			setState(311);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__1:
			case T__22:
			case T__24:
			case KW_PRINT:
			case KW_FALSE:
			case KW_NIL:
			case KW_TRUE:
			case KW_FUNCTION:
			case KW_PCALL:
			case KW_XPCALL:
			case KW_NAN:
			case KW_INF:
			case NUMBER:
			case STRING:
			case LETTER:
				{
				setState(298);
				field();
				setState(303);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,23,_ctx);
				while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
					if ( _alt==1 ) {
						{
						{
						setState(299);
						match(T__23);
						setState(300);
						field();
						}
						} 
					}
					setState(305);
					_errHandler.sync(this);
					_alt = getInterpreter().adaptivePredict(_input,23,_ctx);
				}
				setState(308);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==T__23) {
					{
					setState(306);
					match(T__23);
					setState(307);
					metatable_field();
					}
				}

				}
				break;
			case T__32:
				{
				setState(310);
				metatable_field();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			setState(313);
			match(T__25);
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
		public IdentifierContext identifier() {
			return getRuleContext(IdentifierContext.class,0);
		}
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
		enterRule(_localctx, 46, RULE_field);
		try {
			setState(320);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,26,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(315);
				identifier();
				setState(316);
				match(T__0);
				setState(317);
				expression(0);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(319);
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
	public static class Table_accessContext extends ParserRuleContext {
		public List<IdentifierContext> identifier() {
			return getRuleContexts(IdentifierContext.class);
		}
		public IdentifierContext identifier(int i) {
			return getRuleContext(IdentifierContext.class,i);
		}
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public Table_accessContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_table_access; }
	}

	public final Table_accessContext table_access() throws RecognitionException {
		Table_accessContext _localctx = new Table_accessContext(_ctx, getState());
		enterRule(_localctx, 48, RULE_table_access);
		try {
			setState(331);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,27,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(322);
				identifier();
				setState(323);
				match(T__26);
				setState(324);
				expression(0);
				setState(325);
				match(T__27);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(327);
				identifier();
				setState(328);
				match(T__28);
				setState(329);
				identifier();
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
	public static class Single_line_commentContext extends ParserRuleContext {
		public Single_line_commentContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_single_line_comment; }
	}

	public final Single_line_commentContext single_line_comment() throws RecognitionException {
		Single_line_commentContext _localctx = new Single_line_commentContext(_ctx, getState());
		enterRule(_localctx, 50, RULE_single_line_comment);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(333);
			match(T__29);
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
	public static class Print_statementContext extends ParserRuleContext {
		public TerminalNode KW_PRINT() { return getToken(lua_grammar_antlr4Parser.KW_PRINT, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public Print_statementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_print_statement; }
	}

	public final Print_statementContext print_statement() throws RecognitionException {
		Print_statementContext _localctx = new Print_statementContext(_ctx, getState());
		enterRule(_localctx, 52, RULE_print_statement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(335);
			match(KW_PRINT);
			setState(336);
			match(T__1);
			setState(337);
			expression(0);
			setState(338);
			match(T__2);
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
	public static class IdentifierContext extends ParserRuleContext {
		public List<TerminalNode> LETTER() { return getTokens(lua_grammar_antlr4Parser.LETTER); }
		public TerminalNode LETTER(int i) {
			return getToken(lua_grammar_antlr4Parser.LETTER, i);
		}
		public List<TerminalNode> DIGIT() { return getTokens(lua_grammar_antlr4Parser.DIGIT); }
		public TerminalNode DIGIT(int i) {
			return getToken(lua_grammar_antlr4Parser.DIGIT, i);
		}
		public IdentifierContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_identifier; }
	}

	public final IdentifierContext identifier() throws RecognitionException {
		IdentifierContext _localctx = new IdentifierContext(_ctx, getState());
		enterRule(_localctx, 54, RULE_identifier);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(340);
			match(LETTER);
			setState(344);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,28,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(341);
					_la = _input.LA(1);
					if ( !(((((_la - 31)) & ~0x3f) == 0 && ((1L << (_la - 31)) & 27021597764222977L) != 0)) ) {
					_errHandler.recoverInline(this);
					}
					else {
						if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
						_errHandler.reportMatch(this);
						consume();
					}
					}
					} 
				}
				setState(346);
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
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Repeat_statementContext extends ParserRuleContext {
		public TerminalNode KW_REPEAT() { return getToken(lua_grammar_antlr4Parser.KW_REPEAT, 0); }
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public TerminalNode KW_UNTIL() { return getToken(lua_grammar_antlr4Parser.KW_UNTIL, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public Statement_terminatorContext statement_terminator() {
			return getRuleContext(Statement_terminatorContext.class,0);
		}
		public TerminalNode EOF() { return getToken(lua_grammar_antlr4Parser.EOF, 0); }
		public Repeat_statementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_repeat_statement; }
	}

	public final Repeat_statementContext repeat_statement() throws RecognitionException {
		Repeat_statementContext _localctx = new Repeat_statementContext(_ctx, getState());
		enterRule(_localctx, 56, RULE_repeat_statement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(347);
			match(KW_REPEAT);
			setState(348);
			block();
			setState(349);
			match(KW_UNTIL);
			setState(350);
			expression(0);
			setState(353);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case SEPARATOR:
				{
				setState(351);
				statement_terminator();
				}
				break;
			case EOF:
				{
				setState(352);
				match(EOF);
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
	public static class Identifier_listContext extends ParserRuleContext {
		public List<IdentifierContext> identifier() {
			return getRuleContexts(IdentifierContext.class);
		}
		public IdentifierContext identifier(int i) {
			return getRuleContext(IdentifierContext.class,i);
		}
		public Identifier_listContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_identifier_list; }
	}

	public final Identifier_listContext identifier_list() throws RecognitionException {
		Identifier_listContext _localctx = new Identifier_listContext(_ctx, getState());
		enterRule(_localctx, 58, RULE_identifier_list);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(355);
			identifier();
			setState(360);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__23) {
				{
				{
				setState(356);
				match(T__23);
				setState(357);
				identifier();
				}
				}
				setState(362);
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
	public static class Expression_listContext extends ParserRuleContext {
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public TerminalNode VARARG() { return getToken(lua_grammar_antlr4Parser.VARARG, 0); }
		public Expression_listContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expression_list; }
	}

	public final Expression_listContext expression_list() throws RecognitionException {
		Expression_listContext _localctx = new Expression_listContext(_ctx, getState());
		enterRule(_localctx, 60, RULE_expression_list);
		int _la;
		try {
			setState(372);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__1:
			case T__22:
			case T__24:
			case KW_PRINT:
			case KW_FALSE:
			case KW_NIL:
			case KW_TRUE:
			case KW_FUNCTION:
			case KW_PCALL:
			case KW_XPCALL:
			case KW_NAN:
			case KW_INF:
			case NUMBER:
			case STRING:
			case LETTER:
				enterOuterAlt(_localctx, 1);
				{
				setState(363);
				expression(0);
				setState(368);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==T__23) {
					{
					{
					setState(364);
					match(T__23);
					setState(365);
					expression(0);
					}
					}
					setState(370);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				}
				break;
			case VARARG:
				enterOuterAlt(_localctx, 2);
				{
				setState(371);
				match(VARARG);
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
	public static class Return_statementContext extends ParserRuleContext {
		public TerminalNode KW_RETURN() { return getToken(lua_grammar_antlr4Parser.KW_RETURN, 0); }
		public Statement_terminatorContext statement_terminator() {
			return getRuleContext(Statement_terminatorContext.class,0);
		}
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public Return_statementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_return_statement; }
	}

	public final Return_statementContext return_statement() throws RecognitionException {
		Return_statementContext _localctx = new Return_statementContext(_ctx, getState());
		enterRule(_localctx, 62, RULE_return_statement);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(374);
			match(KW_RETURN);
			setState(383);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 612771024341041156L) != 0) || ((((_la - 65)) & ~0x3f) == 0 && ((1L << (_la - 65)) & 1016593L) != 0)) {
				{
				setState(375);
				expression(0);
				setState(380);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==T__23) {
					{
					{
					setState(376);
					match(T__23);
					setState(377);
					expression(0);
					}
					}
					setState(382);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				}
			}

			setState(385);
			statement_terminator();
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
	public static class Break_statementContext extends ParserRuleContext {
		public TerminalNode KW_BREAK() { return getToken(lua_grammar_antlr4Parser.KW_BREAK, 0); }
		public Statement_terminatorContext statement_terminator() {
			return getRuleContext(Statement_terminatorContext.class,0);
		}
		public Break_statementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_break_statement; }
	}

	public final Break_statementContext break_statement() throws RecognitionException {
		Break_statementContext _localctx = new Break_statementContext(_ctx, getState());
		enterRule(_localctx, 64, RULE_break_statement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(387);
			match(KW_BREAK);
			setState(388);
			statement_terminator();
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
	public static class Goto_statementContext extends ParserRuleContext {
		public TerminalNode KW_GOTO() { return getToken(lua_grammar_antlr4Parser.KW_GOTO, 0); }
		public IdentifierContext identifier() {
			return getRuleContext(IdentifierContext.class,0);
		}
		public Statement_terminatorContext statement_terminator() {
			return getRuleContext(Statement_terminatorContext.class,0);
		}
		public Goto_statementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_goto_statement; }
	}

	public final Goto_statementContext goto_statement() throws RecognitionException {
		Goto_statementContext _localctx = new Goto_statementContext(_ctx, getState());
		enterRule(_localctx, 66, RULE_goto_statement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(390);
			match(KW_GOTO);
			setState(391);
			identifier();
			setState(392);
			statement_terminator();
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
	public static class Label_statementContext extends ParserRuleContext {
		public IdentifierContext identifier() {
			return getRuleContext(IdentifierContext.class,0);
		}
		public Label_statementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_label_statement; }
	}

	public final Label_statementContext label_statement() throws RecognitionException {
		Label_statementContext _localctx = new Label_statementContext(_ctx, getState());
		enterRule(_localctx, 68, RULE_label_statement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(394);
			match(T__31);
			setState(395);
			identifier();
			setState(396);
			match(T__31);
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
	public static class Function_expressionContext extends ParserRuleContext {
		public TerminalNode KW_FUNCTION() { return getToken(lua_grammar_antlr4Parser.KW_FUNCTION, 0); }
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public TerminalNode KW_END() { return getToken(lua_grammar_antlr4Parser.KW_END, 0); }
		public List<IdentifierContext> identifier() {
			return getRuleContexts(IdentifierContext.class);
		}
		public IdentifierContext identifier(int i) {
			return getRuleContext(IdentifierContext.class,i);
		}
		public Function_expressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_function_expression; }
	}

	public final Function_expressionContext function_expression() throws RecognitionException {
		Function_expressionContext _localctx = new Function_expressionContext(_ctx, getState());
		enterRule(_localctx, 70, RULE_function_expression);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(398);
			match(KW_FUNCTION);
			setState(399);
			match(T__1);
			setState(408);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==LETTER) {
				{
				setState(400);
				identifier();
				setState(405);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==T__23) {
					{
					{
					setState(401);
					match(T__23);
					setState(402);
					identifier();
					}
					}
					setState(407);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				}
			}

			setState(410);
			match(T__2);
			setState(411);
			block();
			setState(412);
			match(KW_END);
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
	public static class Method_callContext extends ParserRuleContext {
		public List<IdentifierContext> identifier() {
			return getRuleContexts(IdentifierContext.class);
		}
		public IdentifierContext identifier(int i) {
			return getRuleContext(IdentifierContext.class,i);
		}
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public Method_callContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_method_call; }
	}

	public final Method_callContext method_call() throws RecognitionException {
		Method_callContext _localctx = new Method_callContext(_ctx, getState());
		enterRule(_localctx, 72, RULE_method_call);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(414);
			identifier();
			setState(415);
			match(T__21);
			setState(416);
			identifier();
			setState(417);
			match(T__1);
			setState(426);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 612771024341041156L) != 0) || ((((_la - 65)) & ~0x3f) == 0 && ((1L << (_la - 65)) & 1016593L) != 0)) {
				{
				setState(418);
				expression(0);
				setState(423);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==T__23) {
					{
					{
					setState(419);
					match(T__23);
					setState(420);
					expression(0);
					}
					}
					setState(425);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				}
			}

			setState(428);
			match(T__2);
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
	public static class Metatable_fieldContext extends ParserRuleContext {
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode KW_INDEX() { return getToken(lua_grammar_antlr4Parser.KW_INDEX, 0); }
		public TerminalNode KW_NEWINDEX() { return getToken(lua_grammar_antlr4Parser.KW_NEWINDEX, 0); }
		public TerminalNode KW_MODE() { return getToken(lua_grammar_antlr4Parser.KW_MODE, 0); }
		public Metatable_fieldContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_metatable_field; }
	}

	public final Metatable_fieldContext metatable_field() throws RecognitionException {
		Metatable_fieldContext _localctx = new Metatable_fieldContext(_ctx, getState());
		enterRule(_localctx, 74, RULE_metatable_field);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(430);
			match(T__32);
			setState(431);
			_la = _input.LA(1);
			if ( !(((((_la - 70)) & ~0x3f) == 0 && ((1L << (_la - 70)) & 7L) != 0)) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			setState(432);
			match(T__0);
			setState(433);
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
	public static class MetamethodContext extends ParserRuleContext {
		public MetamethodContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_metamethod; }
	}

	public final MetamethodContext metamethod() throws RecognitionException {
		MetamethodContext _localctx = new MetamethodContext(_ctx, getState());
		enterRule(_localctx, 76, RULE_metamethod);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(435);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 70351564308480L) != 0)) ) {
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
	public static class Coroutine_statementContext extends ParserRuleContext {
		public TerminalNode KW_COROUTINE() { return getToken(lua_grammar_antlr4Parser.KW_COROUTINE, 0); }
		public TerminalNode KW_CREATE() { return getToken(lua_grammar_antlr4Parser.KW_CREATE, 0); }
		public TerminalNode KW_RESUME() { return getToken(lua_grammar_antlr4Parser.KW_RESUME, 0); }
		public TerminalNode KW_YIELD() { return getToken(lua_grammar_antlr4Parser.KW_YIELD, 0); }
		public TerminalNode KW_STATUS() { return getToken(lua_grammar_antlr4Parser.KW_STATUS, 0); }
		public Coroutine_statementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_coroutine_statement; }
	}

	public final Coroutine_statementContext coroutine_statement() throws RecognitionException {
		Coroutine_statementContext _localctx = new Coroutine_statementContext(_ctx, getState());
		enterRule(_localctx, 78, RULE_coroutine_statement);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(437);
			match(KW_COROUTINE);
			setState(438);
			match(T__28);
			setState(439);
			_la = _input.LA(1);
			if ( !(((((_la - 76)) & ~0x3f) == 0 && ((1L << (_la - 76)) & 15L) != 0)) ) {
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
		case 5:
			return expression_sempred((ExpressionContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean expression_sempred(ExpressionContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 5);
		}
		return true;
	}

	public static final String _serializedATN =
		"\u0004\u0001Y\u01ba\u0002\u0000\u0007\u0000\u0002\u0001\u0007\u0001\u0002"+
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
		"#\u0007#\u0002$\u0007$\u0002%\u0007%\u0002&\u0007&\u0002\'\u0007\'\u0001"+
		"\u0000\u0005\u0000R\b\u0000\n\u0000\f\u0000U\t\u0000\u0001\u0001\u0001"+
		"\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001"+
		"\u0001\u0003\u0001_\b\u0001\u0001\u0002\u0001\u0002\u0001\u0002\u0001"+
		"\u0002\u0001\u0002\u0001\u0002\u0003\u0002g\b\u0002\u0001\u0003\u0001"+
		"\u0003\u0001\u0004\u0003\u0004l\b\u0004\u0001\u0004\u0001\u0004\u0001"+
		"\u0004\u0001\u0004\u0001\u0004\u0001\u0005\u0001\u0005\u0001\u0005\u0001"+
		"\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001"+
		"\u0005\u0001\u0005\u0003\u0005~\b\u0005\u0001\u0005\u0001\u0005\u0001"+
		"\u0005\u0001\u0005\u0005\u0005\u0084\b\u0005\n\u0005\f\u0005\u0087\t\u0005"+
		"\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0003\u0006"+
		"\u008e\b\u0006\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007"+
		"\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0003\u0007\u0099\b\u0007"+
		"\u0001\b\u0001\b\u0001\b\u0001\b\u0001\b\u0003\b\u00a0\b\b\u0001\t\u0001"+
		"\t\u0001\n\u0001\n\u0001\u000b\u0001\u000b\u0001\f\u0001\f\u0001\r\u0001"+
		"\r\u0001\u000e\u0001\u000e\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u000f"+
		"\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u000f"+
		"\u0001\u000f\u0003\u000f\u00b9\b\u000f\u0001\u000f\u0001\u000f\u0003\u000f"+
		"\u00bd\b\u000f\u0001\u000f\u0001\u000f\u0003\u000f\u00c1\b\u000f\u0001"+
		"\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0001"+
		"\u000f\u0001\u000f\u0001\u000f\u0003\u000f\u00cc\b\u000f\u0001\u0010\u0001"+
		"\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001"+
		"\u0011\u0003\u0011\u00d6\b\u0011\u0001\u0011\u0001\u0011\u0001\u0011\u0001"+
		"\u0011\u0001\u0011\u0001\u0011\u0005\u0011\u00de\b\u0011\n\u0011\f\u0011"+
		"\u00e1\t\u0011\u0001\u0011\u0003\u0011\u00e4\b\u0011\u0001\u0011\u0001"+
		"\u0011\u0001\u0011\u0001\u0011\u0001\u0012\u0005\u0012\u00eb\b\u0012\n"+
		"\u0012\f\u0012\u00ee\t\u0012\u0001\u0012\u0001\u0012\u0003\u0012\u00f2"+
		"\b\u0012\u0004\u0012\u00f4\b\u0012\u000b\u0012\f\u0012\u00f5\u0001\u0013"+
		"\u0001\u0013\u0001\u0013\u0001\u0013\u0001\u0013\u0001\u0013\u0001\u0013"+
		"\u0001\u0013\u0001\u0013\u0005\u0013\u0101\b\u0013\n\u0013\f\u0013\u0104"+
		"\t\u0013\u0001\u0013\u0001\u0013\u0003\u0013\u0108\b\u0013\u0001\u0013"+
		"\u0001\u0013\u0001\u0014\u0001\u0014\u0001\u0014\u0001\u0014\u0001\u0014"+
		"\u0001\u0014\u0001\u0014\u0001\u0014\u0001\u0014\u0001\u0014\u0001\u0014"+
		"\u0001\u0014\u0001\u0014\u0001\u0014\u0001\u0014\u0001\u0014\u0003\u0014"+
		"\u011c\b\u0014\u0001\u0014\u0001\u0014\u0001\u0014\u0001\u0014\u0003\u0014"+
		"\u0122\b\u0014\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015"+
		"\u0001\u0015\u0001\u0016\u0001\u0016\u0001\u0016\u0001\u0016\u0005\u0016"+
		"\u012e\b\u0016\n\u0016\f\u0016\u0131\t\u0016\u0001\u0016\u0001\u0016\u0003"+
		"\u0016\u0135\b\u0016\u0001\u0016\u0003\u0016\u0138\b\u0016\u0001\u0016"+
		"\u0001\u0016\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017"+
		"\u0003\u0017\u0141\b\u0017\u0001\u0018\u0001\u0018\u0001\u0018\u0001\u0018"+
		"\u0001\u0018\u0001\u0018\u0001\u0018\u0001\u0018\u0001\u0018\u0003\u0018"+
		"\u014c\b\u0018\u0001\u0019\u0001\u0019\u0001\u001a\u0001\u001a\u0001\u001a"+
		"\u0001\u001a\u0001\u001a\u0001\u001b\u0001\u001b\u0005\u001b\u0157\b\u001b"+
		"\n\u001b\f\u001b\u015a\t\u001b\u0001\u001c\u0001\u001c\u0001\u001c\u0001"+
		"\u001c\u0001\u001c\u0001\u001c\u0003\u001c\u0162\b\u001c\u0001\u001d\u0001"+
		"\u001d\u0001\u001d\u0005\u001d\u0167\b\u001d\n\u001d\f\u001d\u016a\t\u001d"+
		"\u0001\u001e\u0001\u001e\u0001\u001e\u0005\u001e\u016f\b\u001e\n\u001e"+
		"\f\u001e\u0172\t\u001e\u0001\u001e\u0003\u001e\u0175\b\u001e\u0001\u001f"+
		"\u0001\u001f\u0001\u001f\u0001\u001f\u0005\u001f\u017b\b\u001f\n\u001f"+
		"\f\u001f\u017e\t\u001f\u0003\u001f\u0180\b\u001f\u0001\u001f\u0001\u001f"+
		"\u0001 \u0001 \u0001 \u0001!\u0001!\u0001!\u0001!\u0001\"\u0001\"\u0001"+
		"\"\u0001\"\u0001#\u0001#\u0001#\u0001#\u0001#\u0005#\u0194\b#\n#\f#\u0197"+
		"\t#\u0003#\u0199\b#\u0001#\u0001#\u0001#\u0001#\u0001$\u0001$\u0001$\u0001"+
		"$\u0001$\u0001$\u0001$\u0005$\u01a6\b$\n$\f$\u01a9\t$\u0003$\u01ab\b$"+
		"\u0001$\u0001$\u0001%\u0001%\u0001%\u0001%\u0001%\u0001&\u0001&\u0001"+
		"\'\u0001\'\u0001\'\u0001\'\u0001\'\u0000\u0001\n(\u0000\u0002\u0004\u0006"+
		"\b\n\f\u000e\u0010\u0012\u0014\u0016\u0018\u001a\u001c\u001e \"$&(*,."+
		"02468:<>@BDFHJLN\u0000\n\u0001\u0000\u0005\n\u0001\u0000\u000b\u000f\u0002"+
		"\u000011==\u0001\u0000\u0010\u0014\u0004\u000077;;AAPS\u0001\u0000IJ\u0002"+
		"\u0000\u001f\u001fTU\u0001\u0000FH\u0001\u0000\"-\u0001\u0000LO\u01d4"+
		"\u0000S\u0001\u0000\u0000\u0000\u0002^\u0001\u0000\u0000\u0000\u0004f"+
		"\u0001\u0000\u0000\u0000\u0006h\u0001\u0000\u0000\u0000\bk\u0001\u0000"+
		"\u0000\u0000\n}\u0001\u0000\u0000\u0000\f\u008d\u0001\u0000\u0000\u0000"+
		"\u000e\u0098\u0001\u0000\u0000\u0000\u0010\u009f\u0001\u0000\u0000\u0000"+
		"\u0012\u00a1\u0001\u0000\u0000\u0000\u0014\u00a3\u0001\u0000\u0000\u0000"+
		"\u0016\u00a5\u0001\u0000\u0000\u0000\u0018\u00a7\u0001\u0000\u0000\u0000"+
		"\u001a\u00a9\u0001\u0000\u0000\u0000\u001c\u00ab\u0001\u0000\u0000\u0000"+
		"\u001e\u00cb\u0001\u0000\u0000\u0000 \u00cd\u0001\u0000\u0000\u0000\""+
		"\u00d5\u0001\u0000\u0000\u0000$\u00f3\u0001\u0000\u0000\u0000&\u00f7\u0001"+
		"\u0000\u0000\u0000(\u0121\u0001\u0000\u0000\u0000*\u0123\u0001\u0000\u0000"+
		"\u0000,\u0129\u0001\u0000\u0000\u0000.\u0140\u0001\u0000\u0000\u00000"+
		"\u014b\u0001\u0000\u0000\u00002\u014d\u0001\u0000\u0000\u00004\u014f\u0001"+
		"\u0000\u0000\u00006\u0154\u0001\u0000\u0000\u00008\u015b\u0001\u0000\u0000"+
		"\u0000:\u0163\u0001\u0000\u0000\u0000<\u0174\u0001\u0000\u0000\u0000>"+
		"\u0176\u0001\u0000\u0000\u0000@\u0183\u0001\u0000\u0000\u0000B\u0186\u0001"+
		"\u0000\u0000\u0000D\u018a\u0001\u0000\u0000\u0000F\u018e\u0001\u0000\u0000"+
		"\u0000H\u019e\u0001\u0000\u0000\u0000J\u01ae\u0001\u0000\u0000\u0000L"+
		"\u01b3\u0001\u0000\u0000\u0000N\u01b5\u0001\u0000\u0000\u0000PR\u0003"+
		"\u0002\u0001\u0000QP\u0001\u0000\u0000\u0000RU\u0001\u0000\u0000\u0000"+
		"SQ\u0001\u0000\u0000\u0000ST\u0001\u0000\u0000\u0000T\u0001\u0001\u0000"+
		"\u0000\u0000US\u0001\u0000\u0000\u0000V_\u0003\b\u0004\u0000W_\u0003\n"+
		"\u0005\u0000X_\u0003\u0004\u0002\u0000Y_\u0003\"\u0011\u0000Z_\u0003\u001e"+
		"\u000f\u0000[_\u0003>\u001f\u0000\\_\u0003@ \u0000]_\u0003D\"\u0000^V"+
		"\u0001\u0000\u0000\u0000^W\u0001\u0000\u0000\u0000^X\u0001\u0000\u0000"+
		"\u0000^Y\u0001\u0000\u0000\u0000^Z\u0001\u0000\u0000\u0000^[\u0001\u0000"+
		"\u0000\u0000^\\\u0001\u0000\u0000\u0000^]\u0001\u0000\u0000\u0000_\u0003"+
		"\u0001\u0000\u0000\u0000`g\u0003&\u0013\u0000ag\u0003(\u0014\u0000bg\u0003"+
		"*\u0015\u0000cg\u00038\u001c\u0000dg\u0003B!\u0000eg\u0003N\'\u0000f`"+
		"\u0001\u0000\u0000\u0000fa\u0001\u0000\u0000\u0000fb\u0001\u0000\u0000"+
		"\u0000fc\u0001\u0000\u0000\u0000fd\u0001\u0000\u0000\u0000fe\u0001\u0000"+
		"\u0000\u0000g\u0005\u0001\u0000\u0000\u0000hi\u0005.\u0000\u0000i\u0007"+
		"\u0001\u0000\u0000\u0000jl\u0005D\u0000\u0000kj\u0001\u0000\u0000\u0000"+
		"kl\u0001\u0000\u0000\u0000lm\u0001\u0000\u0000\u0000mn\u0003:\u001d\u0000"+
		"no\u0005\u0001\u0000\u0000op\u0003<\u001e\u0000pq\u0003\u0006\u0003\u0000"+
		"q\t\u0001\u0000\u0000\u0000rs\u0006\u0005\uffff\uffff\u0000s~\u0003\u001c"+
		"\u000e\u0000t~\u00036\u001b\u0000uv\u0005\u0002\u0000\u0000vw\u0003\n"+
		"\u0005\u0000wx\u0005\u0003\u0000\u0000x~\u0001\u0000\u0000\u0000y~\u0003"+
		"\u001e\u000f\u0000z~\u0003,\u0016\u0000{~\u00030\u0018\u0000|~\u0003F"+
		"#\u0000}r\u0001\u0000\u0000\u0000}t\u0001\u0000\u0000\u0000}u\u0001\u0000"+
		"\u0000\u0000}y\u0001\u0000\u0000\u0000}z\u0001\u0000\u0000\u0000}{\u0001"+
		"\u0000\u0000\u0000}|\u0001\u0000\u0000\u0000~\u0085\u0001\u0000\u0000"+
		"\u0000\u007f\u0080\n\u0005\u0000\u0000\u0080\u0081\u0003\u0010\b\u0000"+
		"\u0081\u0082\u0003\n\u0005\u0006\u0082\u0084\u0001\u0000\u0000\u0000\u0083"+
		"\u007f\u0001\u0000\u0000\u0000\u0084\u0087\u0001\u0000\u0000\u0000\u0085"+
		"\u0083\u0001\u0000\u0000\u0000\u0085\u0086\u0001\u0000\u0000\u0000\u0086"+
		"\u000b\u0001\u0000\u0000\u0000\u0087\u0085\u0001\u0000\u0000\u0000\u0088"+
		"\u008e\u0003\u000e\u0007\u0000\u0089\u008a\u0005<\u0000\u0000\u008a\u008e"+
		"\u0003\f\u0006\u0000\u008b\u008c\u0005\u0004\u0000\u0000\u008c\u008e\u0003"+
		"\f\u0006\u0000\u008d\u0088\u0001\u0000\u0000\u0000\u008d\u0089\u0001\u0000"+
		"\u0000\u0000\u008d\u008b\u0001\u0000\u0000\u0000\u008e\r\u0001\u0000\u0000"+
		"\u0000\u008f\u0099\u0003\u001c\u000e\u0000\u0090\u0099\u00036\u001b\u0000"+
		"\u0091\u0092\u0005\u0002\u0000\u0000\u0092\u0093\u0003\n\u0005\u0000\u0093"+
		"\u0094\u0005\u0003\u0000\u0000\u0094\u0099\u0001\u0000\u0000\u0000\u0095"+
		"\u0099\u0003\u001e\u000f\u0000\u0096\u0099\u0003,\u0016\u0000\u0097\u0099"+
		"\u00030\u0018\u0000\u0098\u008f\u0001\u0000\u0000\u0000\u0098\u0090\u0001"+
		"\u0000\u0000\u0000\u0098\u0091\u0001\u0000\u0000\u0000\u0098\u0095\u0001"+
		"\u0000\u0000\u0000\u0098\u0096\u0001\u0000\u0000\u0000\u0098\u0097\u0001"+
		"\u0000\u0000\u0000\u0099\u000f\u0001\u0000\u0000\u0000\u009a\u00a0\u0003"+
		"\u0012\t\u0000\u009b\u00a0\u0003\u0014\n\u0000\u009c\u00a0\u0003\u0016"+
		"\u000b\u0000\u009d\u00a0\u0003\u0018\f\u0000\u009e\u00a0\u0003\u001a\r"+
		"\u0000\u009f\u009a\u0001\u0000\u0000\u0000\u009f\u009b\u0001\u0000\u0000"+
		"\u0000\u009f\u009c\u0001\u0000\u0000\u0000\u009f\u009d\u0001\u0000\u0000"+
		"\u0000\u009f\u009e\u0001\u0000\u0000\u0000\u00a0\u0011\u0001\u0000\u0000"+
		"\u0000\u00a1\u00a2\u0007\u0000\u0000\u0000\u00a2\u0013\u0001\u0000\u0000"+
		"\u0000\u00a3\u00a4\u0007\u0001\u0000\u0000\u00a4\u0015\u0001\u0000\u0000"+
		"\u0000\u00a5\u00a6\u0007\u0002\u0000\u0000\u00a6\u0017\u0001\u0000\u0000"+
		"\u0000\u00a7\u00a8\u0007\u0003\u0000\u0000\u00a8\u0019\u0001\u0000\u0000"+
		"\u0000\u00a9\u00aa\u0005\u0015\u0000\u0000\u00aa\u001b\u0001\u0000\u0000"+
		"\u0000\u00ab\u00ac\u0007\u0004\u0000\u0000\u00ac\u001d\u0001\u0000\u0000"+
		"\u0000\u00ad\u00ae\u0007\u0005\u0000\u0000\u00ae\u00af\u0005\u0002\u0000"+
		"\u0000\u00af\u00b0\u0003<\u001e\u0000\u00b0\u00b1\u0005\u0003\u0000\u0000"+
		"\u00b1\u00cc\u0001\u0000\u0000\u0000\u00b2\u00b9\u00036\u001b\u0000\u00b3"+
		"\u00b9\u00030\u0018\u0000\u00b4\u00b5\u0005\u0002\u0000\u0000\u00b5\u00b6"+
		"\u0003\n\u0005\u0000\u00b6\u00b7\u0005\u0003\u0000\u0000\u00b7\u00b9\u0001"+
		"\u0000\u0000\u0000\u00b8\u00b2\u0001\u0000\u0000\u0000\u00b8\u00b3\u0001"+
		"\u0000\u0000\u0000\u00b8\u00b4\u0001\u0000\u0000\u0000\u00b9\u00bc\u0001"+
		"\u0000\u0000\u0000\u00ba\u00bb\u0005\u0016\u0000\u0000\u00bb\u00bd\u0003"+
		"6\u001b\u0000\u00bc\u00ba\u0001\u0000\u0000\u0000\u00bc\u00bd\u0001\u0000"+
		"\u0000\u0000\u00bd\u00be\u0001\u0000\u0000\u0000\u00be\u00c0\u0005\u0002"+
		"\u0000\u0000\u00bf\u00c1\u0003<\u001e\u0000\u00c0\u00bf\u0001\u0000\u0000"+
		"\u0000\u00c0\u00c1\u0001\u0000\u0000\u0000\u00c1\u00c2\u0001\u0000\u0000"+
		"\u0000\u00c2\u00c3\u0005\u0003\u0000\u0000\u00c3\u00cc\u0001\u0000\u0000"+
		"\u0000\u00c4\u00cc\u0003 \u0010\u0000\u00c5\u00c6\u00050\u0000\u0000\u00c6"+
		"\u00c7\u0005\u0002\u0000\u0000\u00c7\u00c8\u0003<\u001e\u0000\u00c8\u00c9"+
		"\u0005\u0003\u0000\u0000\u00c9\u00cc\u0001\u0000\u0000\u0000\u00ca\u00cc"+
		"\u0003H$\u0000\u00cb\u00ad\u0001\u0000\u0000\u0000\u00cb\u00b8\u0001\u0000"+
		"\u0000\u0000\u00cb\u00c4\u0001\u0000\u0000\u0000\u00cb\u00c5\u0001\u0000"+
		"\u0000\u0000\u00cb\u00ca\u0001\u0000\u0000\u0000\u00cc\u001f\u0001\u0000"+
		"\u0000\u0000\u00cd\u00ce\u0005\u0017\u0000\u0000\u00ce\u00cf\u0005\u0002"+
		"\u0000\u0000\u00cf\u00d0\u00036\u001b\u0000\u00d0\u00d1\u0005\u0018\u0000"+
		"\u0000\u00d1\u00d2\u0003\n\u0005\u0000\u00d2\u00d3\u0005\u0003\u0000\u0000"+
		"\u00d3!\u0001\u0000\u0000\u0000\u00d4\u00d6\u0005D\u0000\u0000\u00d5\u00d4"+
		"\u0001\u0000\u0000\u0000\u00d5\u00d6\u0001\u0000\u0000\u0000\u00d6\u00d7"+
		"\u0001\u0000\u0000\u0000\u00d7\u00d8\u0005E\u0000\u0000\u00d8\u00d9\u0003"+
		"6\u001b\u0000\u00d9\u00e3\u0005\u0002\u0000\u0000\u00da\u00df\u00036\u001b"+
		"\u0000\u00db\u00dc\u0005\u0018\u0000\u0000\u00dc\u00de\u00036\u001b\u0000"+
		"\u00dd\u00db\u0001\u0000\u0000\u0000\u00de\u00e1\u0001\u0000\u0000\u0000"+
		"\u00df\u00dd\u0001\u0000\u0000\u0000\u00df\u00e0\u0001\u0000\u0000\u0000"+
		"\u00e0\u00e4\u0001\u0000\u0000\u0000\u00e1\u00df\u0001\u0000\u0000\u0000"+
		"\u00e2\u00e4\u0005Y\u0000\u0000\u00e3\u00da\u0001\u0000\u0000\u0000\u00e3"+
		"\u00e2\u0001\u0000\u0000\u0000\u00e3\u00e4\u0001\u0000\u0000\u0000\u00e4"+
		"\u00e5\u0001\u0000\u0000\u0000\u00e5\u00e6\u0005\u0003\u0000\u0000\u00e6"+
		"\u00e7\u0003$\u0012\u0000\u00e7\u00e8\u00056\u0000\u0000\u00e8#\u0001"+
		"\u0000\u0000\u0000\u00e9\u00eb\u0003\u0006\u0003\u0000\u00ea\u00e9\u0001"+
		"\u0000\u0000\u0000\u00eb\u00ee\u0001\u0000\u0000\u0000\u00ec\u00ea\u0001"+
		"\u0000\u0000\u0000\u00ec\u00ed\u0001\u0000\u0000\u0000\u00ed\u00ef\u0001"+
		"\u0000\u0000\u0000\u00ee\u00ec\u0001\u0000\u0000\u0000\u00ef\u00f1\u0003"+
		"\u0002\u0001\u0000\u00f0\u00f2\u0003\u0006\u0003\u0000\u00f1\u00f0\u0001"+
		"\u0000\u0000\u0000\u00f1\u00f2\u0001\u0000\u0000\u0000\u00f2\u00f4\u0001"+
		"\u0000\u0000\u0000\u00f3\u00ec\u0001\u0000\u0000\u0000\u00f4\u00f5\u0001"+
		"\u0000\u0000\u0000\u00f5\u00f3\u0001\u0000\u0000\u0000\u00f5\u00f6\u0001"+
		"\u0000\u0000\u0000\u00f6%\u0001\u0000\u0000\u0000\u00f7\u00f8\u0005:\u0000"+
		"\u0000\u00f8\u00f9\u0003\n\u0005\u0000\u00f9\u00fa\u0005@\u0000\u0000"+
		"\u00fa\u0102\u0003$\u0012\u0000\u00fb\u00fc\u00055\u0000\u0000\u00fc\u00fd"+
		"\u0003\n\u0005\u0000\u00fd\u00fe\u0005@\u0000\u0000\u00fe\u00ff\u0003"+
		"$\u0012\u0000\u00ff\u0101\u0001\u0000\u0000\u0000\u0100\u00fb\u0001\u0000"+
		"\u0000\u0000\u0101\u0104\u0001\u0000\u0000\u0000\u0102\u0100\u0001\u0000"+
		"\u0000\u0000\u0102\u0103\u0001\u0000\u0000\u0000\u0103\u0107\u0001\u0000"+
		"\u0000\u0000\u0104\u0102\u0001\u0000\u0000\u0000\u0105\u0106\u00054\u0000"+
		"\u0000\u0106\u0108\u0003$\u0012\u0000\u0107\u0105\u0001\u0000\u0000\u0000"+
		"\u0107\u0108\u0001\u0000\u0000\u0000\u0108\u0109\u0001\u0000\u0000\u0000"+
		"\u0109\u010a\u00056\u0000\u0000\u010a\'\u0001\u0000\u0000\u0000\u010b"+
		"\u010c\u00058\u0000\u0000\u010c\u010d\u00036\u001b\u0000\u010d\u010e\u0005"+
		"/\u0000\u0000\u010e\u010f\u0003\n\u0005\u0000\u010f\u0110\u00053\u0000"+
		"\u0000\u0110\u0111\u0003$\u0012\u0000\u0111\u0112\u00056\u0000\u0000\u0112"+
		"\u0122\u0001\u0000\u0000\u0000\u0113\u0114\u00058\u0000\u0000\u0114\u0115"+
		"\u00036\u001b\u0000\u0115\u0116\u0005\u0001\u0000\u0000\u0116\u0117\u0003"+
		"\n\u0005\u0000\u0117\u0118\u0005\u0018\u0000\u0000\u0118\u011b\u0003\n"+
		"\u0005\u0000\u0119\u011a\u0005\u0018\u0000\u0000\u011a\u011c\u0003\n\u0005"+
		"\u0000\u011b\u0119\u0001\u0000\u0000\u0000\u011b\u011c\u0001\u0000\u0000"+
		"\u0000\u011c\u011d\u0001\u0000\u0000\u0000\u011d\u011e\u00053\u0000\u0000"+
		"\u011e\u011f\u0003$\u0012\u0000\u011f\u0120\u00056\u0000\u0000\u0120\u0122"+
		"\u0001\u0000\u0000\u0000\u0121\u010b\u0001\u0000\u0000\u0000\u0121\u0113"+
		"\u0001\u0000\u0000\u0000\u0122)\u0001\u0000\u0000\u0000\u0123\u0124\u0005"+
		"C\u0000\u0000\u0124\u0125\u0003\n\u0005\u0000\u0125\u0126\u00053\u0000"+
		"\u0000\u0126\u0127\u0003$\u0012\u0000\u0127\u0128\u00056\u0000\u0000\u0128"+
		"+\u0001\u0000\u0000\u0000\u0129\u0137\u0005\u0019\u0000\u0000\u012a\u012f"+
		"\u0003.\u0017\u0000\u012b\u012c\u0005\u0018\u0000\u0000\u012c\u012e\u0003"+
		".\u0017\u0000\u012d\u012b\u0001\u0000\u0000\u0000\u012e\u0131\u0001\u0000"+
		"\u0000\u0000\u012f\u012d\u0001\u0000\u0000\u0000\u012f\u0130\u0001\u0000"+
		"\u0000\u0000\u0130\u0134\u0001\u0000\u0000\u0000\u0131\u012f\u0001\u0000"+
		"\u0000\u0000\u0132\u0133\u0005\u0018\u0000\u0000\u0133\u0135\u0003J%\u0000"+
		"\u0134\u0132\u0001\u0000\u0000\u0000\u0134\u0135\u0001\u0000\u0000\u0000"+
		"\u0135\u0138\u0001\u0000\u0000\u0000\u0136\u0138\u0003J%\u0000\u0137\u012a"+
		"\u0001\u0000\u0000\u0000\u0137\u0136\u0001\u0000\u0000\u0000\u0138\u0139"+
		"\u0001\u0000\u0000\u0000\u0139\u013a\u0005\u001a\u0000\u0000\u013a-\u0001"+
		"\u0000\u0000\u0000\u013b\u013c\u00036\u001b\u0000\u013c\u013d\u0005\u0001"+
		"\u0000\u0000\u013d\u013e\u0003\n\u0005\u0000\u013e\u0141\u0001\u0000\u0000"+
		"\u0000\u013f\u0141\u0003\n\u0005\u0000\u0140\u013b\u0001\u0000\u0000\u0000"+
		"\u0140\u013f\u0001\u0000\u0000\u0000\u0141/\u0001\u0000\u0000\u0000\u0142"+
		"\u0143\u00036\u001b\u0000\u0143\u0144\u0005\u001b\u0000\u0000\u0144\u0145"+
		"\u0003\n\u0005\u0000\u0145\u0146\u0005\u001c\u0000\u0000\u0146\u014c\u0001"+
		"\u0000\u0000\u0000\u0147\u0148\u00036\u001b\u0000\u0148\u0149\u0005\u001d"+
		"\u0000\u0000\u0149\u014a\u00036\u001b\u0000\u014a\u014c\u0001\u0000\u0000"+
		"\u0000\u014b\u0142\u0001\u0000\u0000\u0000\u014b\u0147\u0001\u0000\u0000"+
		"\u0000\u014c1\u0001\u0000\u0000\u0000\u014d\u014e\u0005\u001e\u0000\u0000"+
		"\u014e3\u0001\u0000\u0000\u0000\u014f\u0150\u00050\u0000\u0000\u0150\u0151"+
		"\u0005\u0002\u0000\u0000\u0151\u0152\u0003\n\u0005\u0000\u0152\u0153\u0005"+
		"\u0003\u0000\u0000\u01535\u0001\u0000\u0000\u0000\u0154\u0158\u0005T\u0000"+
		"\u0000\u0155\u0157\u0007\u0006\u0000\u0000\u0156\u0155\u0001\u0000\u0000"+
		"\u0000\u0157\u015a\u0001\u0000\u0000\u0000\u0158\u0156\u0001\u0000\u0000"+
		"\u0000\u0158\u0159\u0001\u0000\u0000\u0000\u01597\u0001\u0000\u0000\u0000"+
		"\u015a\u0158\u0001\u0000\u0000\u0000\u015b\u015c\u0005>\u0000\u0000\u015c"+
		"\u015d\u0003$\u0012\u0000\u015d\u015e\u0005B\u0000\u0000\u015e\u0161\u0003"+
		"\n\u0005\u0000\u015f\u0162\u0003\u0006\u0003\u0000\u0160\u0162\u0005\u0000"+
		"\u0000\u0001\u0161\u015f\u0001\u0000\u0000\u0000\u0161\u0160\u0001\u0000"+
		"\u0000\u0000\u01629\u0001\u0000\u0000\u0000\u0163\u0168\u00036\u001b\u0000"+
		"\u0164\u0165\u0005\u0018\u0000\u0000\u0165\u0167\u00036\u001b\u0000\u0166"+
		"\u0164\u0001\u0000\u0000\u0000\u0167\u016a\u0001\u0000\u0000\u0000\u0168"+
		"\u0166\u0001\u0000\u0000\u0000\u0168\u0169\u0001\u0000\u0000\u0000\u0169"+
		";\u0001\u0000\u0000\u0000\u016a\u0168\u0001\u0000\u0000\u0000\u016b\u0170"+
		"\u0003\n\u0005\u0000\u016c\u016d\u0005\u0018\u0000\u0000\u016d\u016f\u0003"+
		"\n\u0005\u0000\u016e\u016c\u0001\u0000\u0000\u0000\u016f\u0172\u0001\u0000"+
		"\u0000\u0000\u0170\u016e\u0001\u0000\u0000\u0000\u0170\u0171\u0001\u0000"+
		"\u0000\u0000\u0171\u0175\u0001\u0000\u0000\u0000\u0172\u0170\u0001\u0000"+
		"\u0000\u0000\u0173\u0175\u0005Y\u0000\u0000\u0174\u016b\u0001\u0000\u0000"+
		"\u0000\u0174\u0173\u0001\u0000\u0000\u0000\u0175=\u0001\u0000\u0000\u0000"+
		"\u0176\u017f\u0005?\u0000\u0000\u0177\u017c\u0003\n\u0005\u0000\u0178"+
		"\u0179\u0005\u0018\u0000\u0000\u0179\u017b\u0003\n\u0005\u0000\u017a\u0178"+
		"\u0001\u0000\u0000\u0000\u017b\u017e\u0001\u0000\u0000\u0000\u017c\u017a"+
		"\u0001\u0000\u0000\u0000\u017c\u017d\u0001\u0000\u0000\u0000\u017d\u0180"+
		"\u0001\u0000\u0000\u0000\u017e\u017c\u0001\u0000\u0000\u0000\u017f\u0177"+
		"\u0001\u0000\u0000\u0000\u017f\u0180\u0001\u0000\u0000\u0000\u0180\u0181"+
		"\u0001\u0000\u0000\u0000\u0181\u0182\u0003\u0006\u0003\u0000\u0182?\u0001"+
		"\u0000\u0000\u0000\u0183\u0184\u00052\u0000\u0000\u0184\u0185\u0003\u0006"+
		"\u0003\u0000\u0185A\u0001\u0000\u0000\u0000\u0186\u0187\u00059\u0000\u0000"+
		"\u0187\u0188\u00036\u001b\u0000\u0188\u0189\u0003\u0006\u0003\u0000\u0189"+
		"C\u0001\u0000\u0000\u0000\u018a\u018b\u0005 \u0000\u0000\u018b\u018c\u0003"+
		"6\u001b\u0000\u018c\u018d\u0005 \u0000\u0000\u018dE\u0001\u0000\u0000"+
		"\u0000\u018e\u018f\u0005E\u0000\u0000\u018f\u0198\u0005\u0002\u0000\u0000"+
		"\u0190\u0195\u00036\u001b\u0000\u0191\u0192\u0005\u0018\u0000\u0000\u0192"+
		"\u0194\u00036\u001b\u0000\u0193\u0191\u0001\u0000\u0000\u0000\u0194\u0197"+
		"\u0001\u0000\u0000\u0000\u0195\u0193\u0001\u0000\u0000\u0000\u0195\u0196"+
		"\u0001\u0000\u0000\u0000\u0196\u0199\u0001\u0000\u0000\u0000\u0197\u0195"+
		"\u0001\u0000\u0000\u0000\u0198\u0190\u0001\u0000\u0000\u0000\u0198\u0199"+
		"\u0001\u0000\u0000\u0000\u0199\u019a\u0001\u0000\u0000\u0000\u019a\u019b"+
		"\u0005\u0003\u0000\u0000\u019b\u019c\u0003$\u0012\u0000\u019c\u019d\u0005"+
		"6\u0000\u0000\u019dG\u0001\u0000\u0000\u0000\u019e\u019f\u00036\u001b"+
		"\u0000\u019f\u01a0\u0005\u0016\u0000\u0000\u01a0\u01a1\u00036\u001b\u0000"+
		"\u01a1\u01aa\u0005\u0002\u0000\u0000\u01a2\u01a7\u0003\n\u0005\u0000\u01a3"+
		"\u01a4\u0005\u0018\u0000\u0000\u01a4\u01a6\u0003\n\u0005\u0000\u01a5\u01a3"+
		"\u0001\u0000\u0000\u0000\u01a6\u01a9\u0001\u0000\u0000\u0000\u01a7\u01a5"+
		"\u0001\u0000\u0000\u0000\u01a7\u01a8\u0001\u0000\u0000\u0000\u01a8\u01ab"+
		"\u0001\u0000\u0000\u0000\u01a9\u01a7\u0001\u0000\u0000\u0000\u01aa\u01a2"+
		"\u0001\u0000\u0000\u0000\u01aa\u01ab\u0001\u0000\u0000\u0000\u01ab\u01ac"+
		"\u0001\u0000\u0000\u0000\u01ac\u01ad\u0005\u0003\u0000\u0000\u01adI\u0001"+
		"\u0000\u0000\u0000\u01ae\u01af\u0005!\u0000\u0000\u01af\u01b0\u0007\u0007"+
		"\u0000\u0000\u01b0\u01b1\u0005\u0001\u0000\u0000\u01b1\u01b2\u0003\n\u0005"+
		"\u0000\u01b2K\u0001\u0000\u0000\u0000\u01b3\u01b4\u0007\b\u0000\u0000"+
		"\u01b4M\u0001\u0000\u0000\u0000\u01b5\u01b6\u0005K\u0000\u0000\u01b6\u01b7"+
		"\u0005\u001d\u0000\u0000\u01b7\u01b8\u0007\t\u0000\u0000\u01b8O\u0001"+
		"\u0000\u0000\u0000\'S^fk}\u0085\u008d\u0098\u009f\u00b8\u00bc\u00c0\u00cb"+
		"\u00d5\u00df\u00e3\u00ec\u00f1\u00f5\u0102\u0107\u011b\u0121\u012f\u0134"+
		"\u0137\u0140\u014b\u0158\u0161\u0168\u0170\u0174\u017c\u017f\u0195\u0198"+
		"\u01a7\u01aa";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}