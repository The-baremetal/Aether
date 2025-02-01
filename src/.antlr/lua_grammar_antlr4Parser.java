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
		T__45=46, T__46=47, T__47=48, T__48=49, T__49=50, T__50=51, SEPARATOR=52, 
		KW_IN=53, KW_PRINT=54, KW_AND=55, KW_BREAK=56, KW_DO=57, KW_ELSE=58, KW_ELSEIF=59, 
		KW_END=60, KW_FALSE=61, KW_FOR=62, KW_GOTO=63, KW_IF=64, KW_NIL=65, KW_NOT=66, 
		KW_OR=67, KW_REPEAT=68, KW_RETURN=69, KW_THEN=70, KW_TRUE=71, KW_UNTIL=72, 
		KW_WHILE=73, KW_LOCAL=74, KW_FUNCTION=75, KW_INDEX=76, KW_NEWINDEX=77, 
		KW_MODE=78, KW_PCALL=79, KW_XPCALL=80, KW_COROUTINE=81, KW_CREATE=82, 
		KW_RESUME=83, KW_YIELD=84, KW_STATUS=85, KW_NAN=86, KW_INF=87, KW_ERROR=88, 
		KW_ASSERT=89, KW_PAIRS=90, KW_IPAIRS=91, NUMBER=92, STRING=93, LETTER=94, 
		DIGIT=95, WS=96, SINGLE_LINE_COMMENT=97, MULTI_LINE_COMMENT=98, VARARG=99, 
		INCREMENT=100, DECREMENT=101, PLUS_ASSIGN=102, MINUS_ASSIGN=103, MULT_ASSIGN=104, 
		DIV_ASSIGN=105, FLOOR_DIV_ASSIGN=106, EXP_ASSIGN=107, CONCAT_ASSIGN=108, 
		NULL_COALESCE=109;
	public static final int
		RULE_program = 0, RULE_statement = 1, RULE_control_statement = 2, RULE_statement_terminator = 3, 
		RULE_assignment = 4, RULE_expression = 5, RULE_prefix_expression = 6, 
		RULE_primary_expression = 7, RULE_operators = 8, RULE_comparison_operator = 9, 
		RULE_arith_operator = 10, RULE_logical_operator = 11, RULE_bitwise_operator = 12, 
		RULE_concat_operator = 13, RULE_literal = 14, RULE_function_call = 15, 
		RULE_table_insert = 16, RULE_function_declaration = 17, RULE_parameter = 18, 
		RULE_block = 19, RULE_if_statement = 20, RULE_for_statement = 21, RULE_while_statement = 22, 
		RULE_do_statement = 23, RULE_table = 24, RULE_field_separator = 25, RULE_field = 26, 
		RULE_table_access = 27, RULE_single_line_comment = 28, RULE_print_statement = 29, 
		RULE_identifier = 30, RULE_repeat_statement = 31, RULE_identifier_list = 32, 
		RULE_expression_list = 33, RULE_return_statement = 34, RULE_break_statement = 35, 
		RULE_goto_statement = 36, RULE_label_statement = 37, RULE_function_expression = 38, 
		RULE_method_call = 39, RULE_metatable_field = 40, RULE_metamethod = 41, 
		RULE_coroutine_statement = 42;
	private static String[] makeRuleNames() {
		return new String[] {
			"program", "statement", "control_statement", "statement_terminator", 
			"assignment", "expression", "prefix_expression", "primary_expression", 
			"operators", "comparison_operator", "arith_operator", "logical_operator", 
			"bitwise_operator", "concat_operator", "literal", "function_call", "table_insert", 
			"function_declaration", "parameter", "block", "if_statement", "for_statement", 
			"while_statement", "do_statement", "table", "field_separator", "field", 
			"table_access", "single_line_comment", "print_statement", "identifier", 
			"repeat_statement", "identifier_list", "expression_list", "return_statement", 
			"break_statement", "goto_statement", "label_statement", "function_expression", 
			"method_call", "metatable_field", "metamethod", "coroutine_statement"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'='", "'^'", "'('", "')'", "'#'", "'-'", "'~'", "'>'", "'<'", 
			"'>='", "'=='", "'<='", "'~='", "'*'", "'/'", "'+'", "'//'", "'&'", "'|'", 
			"'<<'", "'>>'", "'>>>'", "'..'", "':'", "'table.insert'", "','", "'.'", 
			"'{'", "'}'", "';'", "'['", "']'", "'_'", "'::'", "'__'", "'__add'", 
			"'__sub'", "'__mul'", "'__div'", "'__mod'", "'__pow'", "'__unm'", "'__concat'", 
			"'__len'", "'__eq'", "'__lt'", "'__le'", "'__tostring'", "'__pairs'", 
			"'__ipairs'", "'__call'", null, "'in'", "'print'", "'and'", "'break'", 
			"'do'", "'else'", "'elseif'", "'end'", "'false'", "'for'", "'goto'", 
			"'if'", "'nil'", "'not'", "'or'", "'repeat'", "'return'", "'then'", "'true'", 
			"'until'", "'while'", "'local'", "'function'", "'index'", "'newindex'", 
			"'mode'", "'pcall'", "'xpcall'", "'coroutine'", "'create'", "'resume'", 
			"'yield'", "'status'", "'nan'", "'inf'", "'error'", "'assert'", "'pairs'", 
			"'ipairs'", null, null, null, null, null, null, null, "'...'", "'++'", 
			"'--'", "'+='", "'-='", "'*='", "'/='", "'//='", "'^='", "'..='", "'??'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, "SEPARATOR", "KW_IN", "KW_PRINT", "KW_AND", "KW_BREAK", 
			"KW_DO", "KW_ELSE", "KW_ELSEIF", "KW_END", "KW_FALSE", "KW_FOR", "KW_GOTO", 
			"KW_IF", "KW_NIL", "KW_NOT", "KW_OR", "KW_REPEAT", "KW_RETURN", "KW_THEN", 
			"KW_TRUE", "KW_UNTIL", "KW_WHILE", "KW_LOCAL", "KW_FUNCTION", "KW_INDEX", 
			"KW_NEWINDEX", "KW_MODE", "KW_PCALL", "KW_XPCALL", "KW_COROUTINE", "KW_CREATE", 
			"KW_RESUME", "KW_YIELD", "KW_STATUS", "KW_NAN", "KW_INF", "KW_ERROR", 
			"KW_ASSERT", "KW_PAIRS", "KW_IPAIRS", "NUMBER", "STRING", "LETTER", "DIGIT", 
			"WS", "SINGLE_LINE_COMMENT", "MULTI_LINE_COMMENT", "VARARG", "INCREMENT", 
			"DECREMENT", "PLUS_ASSIGN", "MINUS_ASSIGN", "MULT_ASSIGN", "DIV_ASSIGN", 
			"FLOOR_DIV_ASSIGN", "EXP_ASSIGN", "CONCAT_ASSIGN", "NULL_COALESCE"
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
			setState(89);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & -2071655811108568856L) != 0) || ((((_la - 64)) & ~0x3f) == 0 && ((1L << (_la - 64)) & 208100495031L) != 0)) {
				{
				{
				setState(86);
				statement();
				}
				}
				setState(91);
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
			setState(100);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,1,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(92);
				assignment();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(93);
				expression(0);
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(94);
				control_statement();
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(95);
				function_declaration();
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(96);
				function_call();
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(97);
				return_statement();
				}
				break;
			case 7:
				enterOuterAlt(_localctx, 7);
				{
				setState(98);
				break_statement();
				}
				break;
			case 8:
				enterOuterAlt(_localctx, 8);
				{
				setState(99);
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
		public Do_statementContext do_statement() {
			return getRuleContext(Do_statementContext.class,0);
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
			setState(108);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case KW_IF:
				enterOuterAlt(_localctx, 1);
				{
				setState(102);
				if_statement();
				}
				break;
			case KW_FOR:
				enterOuterAlt(_localctx, 2);
				{
				setState(103);
				for_statement();
				}
				break;
			case KW_WHILE:
				enterOuterAlt(_localctx, 3);
				{
				setState(104);
				while_statement();
				}
				break;
			case KW_REPEAT:
				enterOuterAlt(_localctx, 4);
				{
				setState(105);
				repeat_statement();
				}
				break;
			case KW_GOTO:
				enterOuterAlt(_localctx, 5);
				{
				setState(106);
				goto_statement();
				}
				break;
			case KW_DO:
				enterOuterAlt(_localctx, 6);
				{
				setState(107);
				do_statement();
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
			setState(110);
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
		public TerminalNode PLUS_ASSIGN() { return getToken(lua_grammar_antlr4Parser.PLUS_ASSIGN, 0); }
		public TerminalNode MINUS_ASSIGN() { return getToken(lua_grammar_antlr4Parser.MINUS_ASSIGN, 0); }
		public TerminalNode MULT_ASSIGN() { return getToken(lua_grammar_antlr4Parser.MULT_ASSIGN, 0); }
		public TerminalNode DIV_ASSIGN() { return getToken(lua_grammar_antlr4Parser.DIV_ASSIGN, 0); }
		public TerminalNode FLOOR_DIV_ASSIGN() { return getToken(lua_grammar_antlr4Parser.FLOOR_DIV_ASSIGN, 0); }
		public TerminalNode EXP_ASSIGN() { return getToken(lua_grammar_antlr4Parser.EXP_ASSIGN, 0); }
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
			setState(113);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==KW_LOCAL) {
				{
				setState(112);
				match(KW_LOCAL);
				}
			}

			setState(115);
			identifier_list();
			setState(116);
			_la = _input.LA(1);
			if ( !(_la==T__0 || ((((_la - 102)) & ~0x3f) == 0 && ((1L << (_la - 102)) & 63L) != 0)) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			setState(117);
			expression_list();
			setState(118);
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
		public Prefix_expressionContext prefix_expression() {
			return getRuleContext(Prefix_expressionContext.class,0);
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
		public TerminalNode INCREMENT() { return getToken(lua_grammar_antlr4Parser.INCREMENT, 0); }
		public TerminalNode DECREMENT() { return getToken(lua_grammar_antlr4Parser.DECREMENT, 0); }
		public OperatorsContext operators() {
			return getRuleContext(OperatorsContext.class,0);
		}
		public TerminalNode NULL_COALESCE() { return getToken(lua_grammar_antlr4Parser.NULL_COALESCE, 0); }
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
			setState(136);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,4,_ctx) ) {
			case 1:
				{
				setState(121);
				literal();
				}
				break;
			case 2:
				{
				setState(122);
				identifier();
				}
				break;
			case 3:
				{
				setState(123);
				prefix_expression();
				}
				break;
			case 4:
				{
				setState(124);
				match(T__2);
				setState(125);
				expression(0);
				setState(126);
				match(T__3);
				}
				break;
			case 5:
				{
				setState(128);
				function_call();
				}
				break;
			case 6:
				{
				setState(129);
				table();
				}
				break;
			case 7:
				{
				setState(130);
				table_access();
				}
				break;
			case 8:
				{
				setState(131);
				function_expression();
				}
				break;
			case 9:
				{
				setState(132);
				match(INCREMENT);
				setState(133);
				expression(3);
				}
				break;
			case 10:
				{
				setState(134);
				match(DECREMENT);
				setState(135);
				expression(1);
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(154);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,6,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(152);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,5,_ctx) ) {
					case 1:
						{
						_localctx = new ExpressionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(138);
						if (!(precpred(_ctx, 12))) throw new FailedPredicateException(this, "precpred(_ctx, 12)");
						setState(139);
						match(T__1);
						setState(140);
						expression(13);
						}
						break;
					case 2:
						{
						_localctx = new ExpressionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(141);
						if (!(precpred(_ctx, 10))) throw new FailedPredicateException(this, "precpred(_ctx, 10)");
						setState(142);
						operators();
						setState(143);
						expression(11);
						}
						break;
					case 3:
						{
						_localctx = new ExpressionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(145);
						if (!(precpred(_ctx, 5))) throw new FailedPredicateException(this, "precpred(_ctx, 5)");
						setState(146);
						match(NULL_COALESCE);
						setState(147);
						expression(6);
						}
						break;
					case 4:
						{
						_localctx = new ExpressionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(148);
						if (!(precpred(_ctx, 4))) throw new FailedPredicateException(this, "precpred(_ctx, 4)");
						setState(149);
						match(INCREMENT);
						}
						break;
					case 5:
						{
						_localctx = new ExpressionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(150);
						if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
						setState(151);
						match(DECREMENT);
						}
						break;
					}
					} 
				}
				setState(156);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,6,_ctx);
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
		public TerminalNode INCREMENT() { return getToken(lua_grammar_antlr4Parser.INCREMENT, 0); }
		public TerminalNode DECREMENT() { return getToken(lua_grammar_antlr4Parser.DECREMENT, 0); }
		public Prefix_expressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_prefix_expression; }
	}

	public final Prefix_expressionContext prefix_expression() throws RecognitionException {
		Prefix_expressionContext _localctx = new Prefix_expressionContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_prefix_expression);
		try {
			setState(170);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__2:
			case T__24:
			case T__27:
			case KW_PRINT:
			case KW_FALSE:
			case KW_NIL:
			case KW_TRUE:
			case KW_PCALL:
			case KW_XPCALL:
			case KW_NAN:
			case KW_INF:
			case KW_ERROR:
			case KW_ASSERT:
			case NUMBER:
			case STRING:
			case LETTER:
				enterOuterAlt(_localctx, 1);
				{
				setState(157);
				primary_expression(0);
				}
				break;
			case KW_NOT:
				enterOuterAlt(_localctx, 2);
				{
				setState(158);
				match(KW_NOT);
				setState(159);
				prefix_expression();
				}
				break;
			case T__4:
				enterOuterAlt(_localctx, 3);
				{
				setState(160);
				match(T__4);
				setState(161);
				prefix_expression();
				}
				break;
			case T__5:
				enterOuterAlt(_localctx, 4);
				{
				setState(162);
				match(T__5);
				setState(163);
				prefix_expression();
				}
				break;
			case T__6:
				enterOuterAlt(_localctx, 5);
				{
				setState(164);
				match(T__6);
				setState(165);
				prefix_expression();
				}
				break;
			case INCREMENT:
				enterOuterAlt(_localctx, 6);
				{
				setState(166);
				match(INCREMENT);
				setState(167);
				prefix_expression();
				}
				break;
			case DECREMENT:
				enterOuterAlt(_localctx, 7);
				{
				setState(168);
				match(DECREMENT);
				setState(169);
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
		public Primary_expressionContext primary_expression() {
			return getRuleContext(Primary_expressionContext.class,0);
		}
		public TerminalNode INCREMENT() { return getToken(lua_grammar_antlr4Parser.INCREMENT, 0); }
		public TerminalNode DECREMENT() { return getToken(lua_grammar_antlr4Parser.DECREMENT, 0); }
		public Primary_expressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_primary_expression; }
	}

	public final Primary_expressionContext primary_expression() throws RecognitionException {
		return primary_expression(0);
	}

	private Primary_expressionContext primary_expression(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		Primary_expressionContext _localctx = new Primary_expressionContext(_ctx, _parentState);
		Primary_expressionContext _prevctx = _localctx;
		int _startState = 14;
		enterRecursionRule(_localctx, 14, RULE_primary_expression, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(182);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,8,_ctx) ) {
			case 1:
				{
				setState(173);
				literal();
				}
				break;
			case 2:
				{
				setState(174);
				identifier();
				}
				break;
			case 3:
				{
				setState(175);
				match(T__2);
				setState(176);
				expression(0);
				setState(177);
				match(T__3);
				}
				break;
			case 4:
				{
				setState(179);
				function_call();
				}
				break;
			case 5:
				{
				setState(180);
				table();
				}
				break;
			case 6:
				{
				setState(181);
				table_access();
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(190);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,10,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(188);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,9,_ctx) ) {
					case 1:
						{
						_localctx = new Primary_expressionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_primary_expression);
						setState(184);
						if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
						setState(185);
						match(INCREMENT);
						}
						break;
					case 2:
						{
						_localctx = new Primary_expressionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_primary_expression);
						setState(186);
						if (!(precpred(_ctx, 1))) throw new FailedPredicateException(this, "precpred(_ctx, 1)");
						setState(187);
						match(DECREMENT);
						}
						break;
					}
					} 
				}
				setState(192);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,10,_ctx);
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
			setState(198);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__7:
			case T__8:
			case T__9:
			case T__10:
			case T__11:
			case T__12:
				enterOuterAlt(_localctx, 1);
				{
				setState(193);
				comparison_operator();
				}
				break;
			case T__5:
			case T__13:
			case T__14:
			case T__15:
			case T__16:
				enterOuterAlt(_localctx, 2);
				{
				setState(194);
				arith_operator();
				}
				break;
			case KW_AND:
			case KW_OR:
				enterOuterAlt(_localctx, 3);
				{
				setState(195);
				logical_operator();
				}
				break;
			case T__6:
			case T__17:
			case T__18:
			case T__19:
			case T__20:
			case T__21:
				enterOuterAlt(_localctx, 4);
				{
				setState(196);
				bitwise_operator();
				}
				break;
			case T__22:
				enterOuterAlt(_localctx, 5);
				{
				setState(197);
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
			setState(200);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 16128L) != 0)) ) {
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
			setState(202);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 245824L) != 0)) ) {
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
			setState(204);
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
			setState(206);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 8126592L) != 0)) ) {
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
			setState(208);
			match(T__22);
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
			setState(210);
			_la = _input.LA(1);
			if ( !(((((_la - 61)) & ~0x3f) == 0 && ((1L << (_la - 61)) & 6543115281L) != 0)) ) {
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
		public TerminalNode KW_ERROR() { return getToken(lua_grammar_antlr4Parser.KW_ERROR, 0); }
		public TerminalNode KW_ASSERT() { return getToken(lua_grammar_antlr4Parser.KW_ASSERT, 0); }
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
			setState(247);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,15,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(212);
				_la = _input.LA(1);
				if ( !(_la==KW_PCALL || _la==KW_XPCALL) ) {
				_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				setState(213);
				match(T__2);
				setState(214);
				expression_list();
				setState(215);
				match(T__3);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(217);
				_la = _input.LA(1);
				if ( !(_la==KW_ERROR || _la==KW_ASSERT) ) {
				_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				setState(218);
				match(T__2);
				setState(219);
				expression_list();
				setState(220);
				match(T__3);
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(228);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,12,_ctx) ) {
				case 1:
					{
					setState(222);
					identifier();
					}
					break;
				case 2:
					{
					setState(223);
					table_access();
					}
					break;
				case 3:
					{
					setState(224);
					match(T__2);
					setState(225);
					expression(0);
					setState(226);
					match(T__3);
					}
					break;
				}
				setState(232);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==T__23) {
					{
					setState(230);
					match(T__23);
					setState(231);
					identifier();
					}
				}

				setState(234);
				match(T__2);
				setState(236);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 2323857408025166056L) != 0) || ((((_la - 65)) & ~0x3f) == 0 && ((1L << (_la - 65)) & 121230115907L) != 0)) {
					{
					setState(235);
					expression_list();
					}
				}

				setState(238);
				match(T__3);
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(240);
				table_insert();
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(241);
				match(KW_PRINT);
				setState(242);
				match(T__2);
				setState(243);
				expression_list();
				setState(244);
				match(T__3);
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(246);
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
			setState(249);
			match(T__24);
			setState(250);
			match(T__2);
			setState(251);
			identifier();
			setState(252);
			match(T__25);
			setState(253);
			expression(0);
			setState(254);
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
		public List<ParameterContext> parameter() {
			return getRuleContexts(ParameterContext.class);
		}
		public ParameterContext parameter(int i) {
			return getRuleContext(ParameterContext.class,i);
		}
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
			setState(257);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==KW_LOCAL) {
				{
				setState(256);
				match(KW_LOCAL);
				}
			}

			setState(259);
			match(KW_FUNCTION);
			setState(264);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,17,_ctx) ) {
			case 1:
				{
				setState(260);
				identifier();
				setState(261);
				match(T__26);
				setState(262);
				identifier();
				}
				break;
			}
			setState(266);
			identifier();
			setState(267);
			match(T__2);
			setState(277);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case LETTER:
				{
				setState(268);
				parameter();
				setState(273);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==T__25) {
					{
					{
					setState(269);
					match(T__25);
					setState(270);
					parameter();
					}
					}
					setState(275);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				}
				break;
			case VARARG:
				{
				setState(276);
				match(VARARG);
				}
				break;
			case T__3:
				break;
			default:
				break;
			}
			setState(279);
			match(T__3);
			setState(280);
			block();
			setState(281);
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
	public static class ParameterContext extends ParserRuleContext {
		public IdentifierContext identifier() {
			return getRuleContext(IdentifierContext.class,0);
		}
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public ParameterContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_parameter; }
	}

	public final ParameterContext parameter() throws RecognitionException {
		ParameterContext _localctx = new ParameterContext(_ctx, getState());
		enterRule(_localctx, 36, RULE_parameter);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(283);
			identifier();
			setState(286);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__0) {
				{
				setState(284);
				match(T__0);
				setState(285);
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
		enterRule(_localctx, 38, RULE_block);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(298); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(291);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==SEPARATOR) {
					{
					{
					setState(288);
					statement_terminator();
					}
					}
					setState(293);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(294);
				statement();
				setState(296);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,22,_ctx) ) {
				case 1:
					{
					setState(295);
					statement_terminator();
					}
					break;
				}
				}
				}
				setState(300); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & -2067152211481198360L) != 0) || ((((_la - 64)) & ~0x3f) == 0 && ((1L << (_la - 64)) & 208100495031L) != 0) );
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
		enterRule(_localctx, 40, RULE_if_statement);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(302);
			match(KW_IF);
			setState(303);
			expression(0);
			setState(304);
			match(KW_THEN);
			setState(305);
			block();
			setState(313);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==KW_ELSEIF) {
				{
				{
				setState(306);
				match(KW_ELSEIF);
				setState(307);
				expression(0);
				setState(308);
				match(KW_THEN);
				setState(309);
				block();
				}
				}
				setState(315);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(318);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==KW_ELSE) {
				{
				setState(316);
				match(KW_ELSE);
				setState(317);
				block();
				}
			}

			setState(320);
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
		enterRule(_localctx, 42, RULE_for_statement);
		int _la;
		try {
			setState(344);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,27,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(322);
				match(KW_FOR);
				setState(323);
				identifier();
				setState(324);
				match(KW_IN);
				setState(325);
				expression(0);
				setState(326);
				match(KW_DO);
				setState(327);
				block();
				setState(328);
				match(KW_END);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(330);
				match(KW_FOR);
				setState(331);
				identifier();
				setState(332);
				match(T__0);
				setState(333);
				expression(0);
				setState(334);
				match(T__25);
				setState(335);
				expression(0);
				setState(338);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==T__25) {
					{
					setState(336);
					match(T__25);
					setState(337);
					expression(0);
					}
				}

				setState(340);
				match(KW_DO);
				setState(341);
				block();
				setState(342);
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
		enterRule(_localctx, 44, RULE_while_statement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(346);
			match(KW_WHILE);
			setState(347);
			expression(0);
			setState(348);
			match(KW_DO);
			setState(349);
			block();
			setState(350);
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
	public static class Do_statementContext extends ParserRuleContext {
		public TerminalNode KW_DO() { return getToken(lua_grammar_antlr4Parser.KW_DO, 0); }
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public TerminalNode KW_END() { return getToken(lua_grammar_antlr4Parser.KW_END, 0); }
		public Do_statementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_do_statement; }
	}

	public final Do_statementContext do_statement() throws RecognitionException {
		Do_statementContext _localctx = new Do_statementContext(_ctx, getState());
		enterRule(_localctx, 46, RULE_do_statement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(352);
			match(KW_DO);
			setState(353);
			block();
			setState(354);
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
		public List<Field_separatorContext> field_separator() {
			return getRuleContexts(Field_separatorContext.class);
		}
		public Field_separatorContext field_separator(int i) {
			return getRuleContext(Field_separatorContext.class,i);
		}
		public TableContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_table; }
	}

	public final TableContext table() throws RecognitionException {
		TableContext _localctx = new TableContext(_ctx, getState());
		enterRule(_localctx, 48, RULE_table);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(356);
			match(T__27);
			setState(369);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 2323857410172649704L) != 0) || ((((_la - 65)) & ~0x3f) == 0 && ((1L << (_la - 65)) & 104050246723L) != 0)) {
				{
				setState(357);
				field();
				setState(363);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,28,_ctx);
				while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
					if ( _alt==1 ) {
						{
						{
						setState(358);
						field_separator();
						setState(359);
						field();
						}
						} 
					}
					setState(365);
					_errHandler.sync(this);
					_alt = getInterpreter().adaptivePredict(_input,28,_ctx);
				}
				setState(367);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==T__25 || _la==T__29) {
					{
					setState(366);
					field_separator();
					}
				}

				}
			}

			setState(371);
			match(T__28);
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
	public static class Field_separatorContext extends ParserRuleContext {
		public Field_separatorContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_field_separator; }
	}

	public final Field_separatorContext field_separator() throws RecognitionException {
		Field_separatorContext _localctx = new Field_separatorContext(_ctx, getState());
		enterRule(_localctx, 50, RULE_field_separator);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(373);
			_la = _input.LA(1);
			if ( !(_la==T__25 || _la==T__29) ) {
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
		public IdentifierContext identifier() {
			return getRuleContext(IdentifierContext.class,0);
		}
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public FieldContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_field; }
	}

	public final FieldContext field() throws RecognitionException {
		FieldContext _localctx = new FieldContext(_ctx, getState());
		enterRule(_localctx, 52, RULE_field);
		try {
			setState(386);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,31,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(375);
				identifier();
				setState(376);
				match(T__0);
				setState(377);
				expression(0);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(379);
				match(T__30);
				setState(380);
				expression(0);
				setState(381);
				match(T__31);
				setState(382);
				match(T__0);
				setState(383);
				expression(0);
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(385);
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
		enterRule(_localctx, 54, RULE_table_access);
		try {
			setState(397);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,32,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(388);
				identifier();
				setState(389);
				match(T__30);
				setState(390);
				expression(0);
				setState(391);
				match(T__31);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(393);
				identifier();
				setState(394);
				match(T__26);
				setState(395);
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
		public TerminalNode DECREMENT() { return getToken(lua_grammar_antlr4Parser.DECREMENT, 0); }
		public Single_line_commentContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_single_line_comment; }
	}

	public final Single_line_commentContext single_line_comment() throws RecognitionException {
		Single_line_commentContext _localctx = new Single_line_commentContext(_ctx, getState());
		enterRule(_localctx, 56, RULE_single_line_comment);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(399);
			match(DECREMENT);
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
		enterRule(_localctx, 58, RULE_print_statement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(401);
			match(KW_PRINT);
			setState(402);
			match(T__2);
			setState(403);
			expression(0);
			setState(404);
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
		enterRule(_localctx, 60, RULE_identifier);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(406);
			match(LETTER);
			setState(410);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,33,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(407);
					_la = _input.LA(1);
					if ( !(((((_la - 33)) & ~0x3f) == 0 && ((1L << (_la - 33)) & 6917529027641081857L) != 0)) ) {
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
				setState(412);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,33,_ctx);
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
		enterRule(_localctx, 62, RULE_repeat_statement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(413);
			match(KW_REPEAT);
			setState(414);
			block();
			setState(415);
			match(KW_UNTIL);
			setState(416);
			expression(0);
			setState(419);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case SEPARATOR:
				{
				setState(417);
				statement_terminator();
				}
				break;
			case EOF:
				{
				setState(418);
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
		enterRule(_localctx, 64, RULE_identifier_list);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(421);
			identifier();
			setState(426);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__25) {
				{
				{
				setState(422);
				match(T__25);
				setState(423);
				identifier();
				}
				}
				setState(428);
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
		enterRule(_localctx, 66, RULE_expression_list);
		int _la;
		try {
			setState(438);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__2:
			case T__4:
			case T__5:
			case T__6:
			case T__24:
			case T__27:
			case KW_PRINT:
			case KW_FALSE:
			case KW_NIL:
			case KW_NOT:
			case KW_TRUE:
			case KW_FUNCTION:
			case KW_PCALL:
			case KW_XPCALL:
			case KW_NAN:
			case KW_INF:
			case KW_ERROR:
			case KW_ASSERT:
			case NUMBER:
			case STRING:
			case LETTER:
			case INCREMENT:
			case DECREMENT:
				enterOuterAlt(_localctx, 1);
				{
				setState(429);
				expression(0);
				setState(434);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==T__25) {
					{
					{
					setState(430);
					match(T__25);
					setState(431);
					expression(0);
					}
					}
					setState(436);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				}
				break;
			case VARARG:
				enterOuterAlt(_localctx, 2);
				{
				setState(437);
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
		enterRule(_localctx, 68, RULE_return_statement);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(440);
			match(KW_RETURN);
			setState(449);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 2323857408025166056L) != 0) || ((((_la - 65)) & ~0x3f) == 0 && ((1L << (_la - 65)) & 104050246723L) != 0)) {
				{
				setState(441);
				expression(0);
				setState(446);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==T__25) {
					{
					{
					setState(442);
					match(T__25);
					setState(443);
					expression(0);
					}
					}
					setState(448);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				}
			}

			setState(451);
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
		enterRule(_localctx, 70, RULE_break_statement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(453);
			match(KW_BREAK);
			setState(454);
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
		enterRule(_localctx, 72, RULE_goto_statement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(456);
			match(KW_GOTO);
			setState(457);
			identifier();
			setState(458);
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
		enterRule(_localctx, 74, RULE_label_statement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(460);
			match(T__33);
			setState(461);
			identifier();
			setState(462);
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
		enterRule(_localctx, 76, RULE_function_expression);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(464);
			match(KW_FUNCTION);
			setState(465);
			match(T__2);
			setState(474);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==LETTER) {
				{
				setState(466);
				identifier();
				setState(471);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==T__25) {
					{
					{
					setState(467);
					match(T__25);
					setState(468);
					identifier();
					}
					}
					setState(473);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				}
			}

			setState(476);
			match(T__3);
			setState(477);
			block();
			setState(478);
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
		enterRule(_localctx, 78, RULE_method_call);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(480);
			identifier();
			setState(481);
			match(T__23);
			setState(482);
			identifier();
			setState(483);
			match(T__2);
			setState(492);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 2323857408025166056L) != 0) || ((((_la - 65)) & ~0x3f) == 0 && ((1L << (_la - 65)) & 104050246723L) != 0)) {
				{
				setState(484);
				expression(0);
				setState(489);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==T__25) {
					{
					{
					setState(485);
					match(T__25);
					setState(486);
					expression(0);
					}
					}
					setState(491);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				}
			}

			setState(494);
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
	public static class Metatable_fieldContext extends ParserRuleContext {
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public MetamethodContext metamethod() {
			return getRuleContext(MetamethodContext.class,0);
		}
		public IdentifierContext identifier() {
			return getRuleContext(IdentifierContext.class,0);
		}
		public Metatable_fieldContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_metatable_field; }
	}

	public final Metatable_fieldContext metatable_field() throws RecognitionException {
		Metatable_fieldContext _localctx = new Metatable_fieldContext(_ctx, getState());
		enterRule(_localctx, 80, RULE_metatable_field);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(496);
			match(T__34);
			setState(499);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__35:
			case T__36:
			case T__37:
			case T__38:
			case T__39:
			case T__40:
			case T__41:
			case T__42:
			case T__43:
			case T__44:
			case T__45:
			case T__46:
			case T__47:
			case T__48:
			case T__49:
			case T__50:
				{
				setState(497);
				metamethod();
				}
				break;
			case LETTER:
				{
				setState(498);
				identifier();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			setState(501);
			match(T__0);
			setState(502);
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
		enterRule(_localctx, 82, RULE_metamethod);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(504);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 4503530907893760L) != 0)) ) {
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
		enterRule(_localctx, 84, RULE_coroutine_statement);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(506);
			match(KW_COROUTINE);
			setState(507);
			match(T__26);
			setState(508);
			_la = _input.LA(1);
			if ( !(((((_la - 82)) & ~0x3f) == 0 && ((1L << (_la - 82)) & 15L) != 0)) ) {
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
		case 7:
			return primary_expression_sempred((Primary_expressionContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean expression_sempred(ExpressionContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 12);
		case 1:
			return precpred(_ctx, 10);
		case 2:
			return precpred(_ctx, 5);
		case 3:
			return precpred(_ctx, 4);
		case 4:
			return precpred(_ctx, 2);
		}
		return true;
	}
	private boolean primary_expression_sempred(Primary_expressionContext _localctx, int predIndex) {
		switch (predIndex) {
		case 5:
			return precpred(_ctx, 2);
		case 6:
			return precpred(_ctx, 1);
		}
		return true;
	}

	public static final String _serializedATN =
		"\u0004\u0001m\u01ff\u0002\u0000\u0007\u0000\u0002\u0001\u0007\u0001\u0002"+
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
		"(\u0007(\u0002)\u0007)\u0002*\u0007*\u0001\u0000\u0005\u0000X\b\u0000"+
		"\n\u0000\f\u0000[\t\u0000\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001"+
		"\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0003\u0001e\b\u0001"+
		"\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002"+
		"\u0003\u0002m\b\u0002\u0001\u0003\u0001\u0003\u0001\u0004\u0003\u0004"+
		"r\b\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004"+
		"\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005"+
		"\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005"+
		"\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0003\u0005\u0089\b\u0005"+
		"\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005"+
		"\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005"+
		"\u0001\u0005\u0001\u0005\u0005\u0005\u0099\b\u0005\n\u0005\f\u0005\u009c"+
		"\t\u0005\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001"+
		"\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001"+
		"\u0006\u0001\u0006\u0003\u0006\u00ab\b\u0006\u0001\u0007\u0001\u0007\u0001"+
		"\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001"+
		"\u0007\u0001\u0007\u0003\u0007\u00b7\b\u0007\u0001\u0007\u0001\u0007\u0001"+
		"\u0007\u0001\u0007\u0005\u0007\u00bd\b\u0007\n\u0007\f\u0007\u00c0\t\u0007"+
		"\u0001\b\u0001\b\u0001\b\u0001\b\u0001\b\u0003\b\u00c7\b\b\u0001\t\u0001"+
		"\t\u0001\n\u0001\n\u0001\u000b\u0001\u000b\u0001\f\u0001\f\u0001\r\u0001"+
		"\r\u0001\u000e\u0001\u000e\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u000f"+
		"\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u000f"+
		"\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u000f"+
		"\u0003\u000f\u00e5\b\u000f\u0001\u000f\u0001\u000f\u0003\u000f\u00e9\b"+
		"\u000f\u0001\u000f\u0001\u000f\u0003\u000f\u00ed\b\u000f\u0001\u000f\u0001"+
		"\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0001"+
		"\u000f\u0001\u000f\u0003\u000f\u00f8\b\u000f\u0001\u0010\u0001\u0010\u0001"+
		"\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0011\u0003"+
		"\u0011\u0102\b\u0011\u0001\u0011\u0001\u0011\u0001\u0011\u0001\u0011\u0001"+
		"\u0011\u0003\u0011\u0109\b\u0011\u0001\u0011\u0001\u0011\u0001\u0011\u0001"+
		"\u0011\u0001\u0011\u0005\u0011\u0110\b\u0011\n\u0011\f\u0011\u0113\t\u0011"+
		"\u0001\u0011\u0003\u0011\u0116\b\u0011\u0001\u0011\u0001\u0011\u0001\u0011"+
		"\u0001\u0011\u0001\u0012\u0001\u0012\u0001\u0012\u0003\u0012\u011f\b\u0012"+
		"\u0001\u0013\u0005\u0013\u0122\b\u0013\n\u0013\f\u0013\u0125\t\u0013\u0001"+
		"\u0013\u0001\u0013\u0003\u0013\u0129\b\u0013\u0004\u0013\u012b\b\u0013"+
		"\u000b\u0013\f\u0013\u012c\u0001\u0014\u0001\u0014\u0001\u0014\u0001\u0014"+
		"\u0001\u0014\u0001\u0014\u0001\u0014\u0001\u0014\u0001\u0014\u0005\u0014"+
		"\u0138\b\u0014\n\u0014\f\u0014\u013b\t\u0014\u0001\u0014\u0001\u0014\u0003"+
		"\u0014\u013f\b\u0014\u0001\u0014\u0001\u0014\u0001\u0015\u0001\u0015\u0001"+
		"\u0015\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015\u0001"+
		"\u0015\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015\u0001"+
		"\u0015\u0001\u0015\u0003\u0015\u0153\b\u0015\u0001\u0015\u0001\u0015\u0001"+
		"\u0015\u0001\u0015\u0003\u0015\u0159\b\u0015\u0001\u0016\u0001\u0016\u0001"+
		"\u0016\u0001\u0016\u0001\u0016\u0001\u0016\u0001\u0017\u0001\u0017\u0001"+
		"\u0017\u0001\u0017\u0001\u0018\u0001\u0018\u0001\u0018\u0001\u0018\u0001"+
		"\u0018\u0005\u0018\u016a\b\u0018\n\u0018\f\u0018\u016d\t\u0018\u0001\u0018"+
		"\u0003\u0018\u0170\b\u0018\u0003\u0018\u0172\b\u0018\u0001\u0018\u0001"+
		"\u0018\u0001\u0019\u0001\u0019\u0001\u001a\u0001\u001a\u0001\u001a\u0001"+
		"\u001a\u0001\u001a\u0001\u001a\u0001\u001a\u0001\u001a\u0001\u001a\u0001"+
		"\u001a\u0001\u001a\u0003\u001a\u0183\b\u001a\u0001\u001b\u0001\u001b\u0001"+
		"\u001b\u0001\u001b\u0001\u001b\u0001\u001b\u0001\u001b\u0001\u001b\u0001"+
		"\u001b\u0003\u001b\u018e\b\u001b\u0001\u001c\u0001\u001c\u0001\u001d\u0001"+
		"\u001d\u0001\u001d\u0001\u001d\u0001\u001d\u0001\u001e\u0001\u001e\u0005"+
		"\u001e\u0199\b\u001e\n\u001e\f\u001e\u019c\t\u001e\u0001\u001f\u0001\u001f"+
		"\u0001\u001f\u0001\u001f\u0001\u001f\u0001\u001f\u0003\u001f\u01a4\b\u001f"+
		"\u0001 \u0001 \u0001 \u0005 \u01a9\b \n \f \u01ac\t \u0001!\u0001!\u0001"+
		"!\u0005!\u01b1\b!\n!\f!\u01b4\t!\u0001!\u0003!\u01b7\b!\u0001\"\u0001"+
		"\"\u0001\"\u0001\"\u0005\"\u01bd\b\"\n\"\f\"\u01c0\t\"\u0003\"\u01c2\b"+
		"\"\u0001\"\u0001\"\u0001#\u0001#\u0001#\u0001$\u0001$\u0001$\u0001$\u0001"+
		"%\u0001%\u0001%\u0001%\u0001&\u0001&\u0001&\u0001&\u0001&\u0005&\u01d6"+
		"\b&\n&\f&\u01d9\t&\u0003&\u01db\b&\u0001&\u0001&\u0001&\u0001&\u0001\'"+
		"\u0001\'\u0001\'\u0001\'\u0001\'\u0001\'\u0001\'\u0005\'\u01e8\b\'\n\'"+
		"\f\'\u01eb\t\'\u0003\'\u01ed\b\'\u0001\'\u0001\'\u0001(\u0001(\u0001("+
		"\u0003(\u01f4\b(\u0001(\u0001(\u0001(\u0001)\u0001)\u0001*\u0001*\u0001"+
		"*\u0001*\u0001*\u0000\u0002\n\u000e+\u0000\u0002\u0004\u0006\b\n\f\u000e"+
		"\u0010\u0012\u0014\u0016\u0018\u001a\u001c\u001e \"$&(*,.02468:<>@BDF"+
		"HJLNPRT\u0000\f\u0002\u0000\u0001\u0001fk\u0001\u0000\b\r\u0002\u0000"+
		"\u0006\u0006\u000e\u0011\u0002\u000077CC\u0002\u0000\u0007\u0007\u0012"+
		"\u0016\u0005\u0000==AAGGVW\\]\u0001\u0000OP\u0001\u0000XY\u0002\u0000"+
		"\u001a\u001a\u001e\u001e\u0002\u0000!!^_\u0001\u0000$3\u0001\u0000RU\u0228"+
		"\u0000Y\u0001\u0000\u0000\u0000\u0002d\u0001\u0000\u0000\u0000\u0004l"+
		"\u0001\u0000\u0000\u0000\u0006n\u0001\u0000\u0000\u0000\bq\u0001\u0000"+
		"\u0000\u0000\n\u0088\u0001\u0000\u0000\u0000\f\u00aa\u0001\u0000\u0000"+
		"\u0000\u000e\u00b6\u0001\u0000\u0000\u0000\u0010\u00c6\u0001\u0000\u0000"+
		"\u0000\u0012\u00c8\u0001\u0000\u0000\u0000\u0014\u00ca\u0001\u0000\u0000"+
		"\u0000\u0016\u00cc\u0001\u0000\u0000\u0000\u0018\u00ce\u0001\u0000\u0000"+
		"\u0000\u001a\u00d0\u0001\u0000\u0000\u0000\u001c\u00d2\u0001\u0000\u0000"+
		"\u0000\u001e\u00f7\u0001\u0000\u0000\u0000 \u00f9\u0001\u0000\u0000\u0000"+
		"\"\u0101\u0001\u0000\u0000\u0000$\u011b\u0001\u0000\u0000\u0000&\u012a"+
		"\u0001\u0000\u0000\u0000(\u012e\u0001\u0000\u0000\u0000*\u0158\u0001\u0000"+
		"\u0000\u0000,\u015a\u0001\u0000\u0000\u0000.\u0160\u0001\u0000\u0000\u0000"+
		"0\u0164\u0001\u0000\u0000\u00002\u0175\u0001\u0000\u0000\u00004\u0182"+
		"\u0001\u0000\u0000\u00006\u018d\u0001\u0000\u0000\u00008\u018f\u0001\u0000"+
		"\u0000\u0000:\u0191\u0001\u0000\u0000\u0000<\u0196\u0001\u0000\u0000\u0000"+
		">\u019d\u0001\u0000\u0000\u0000@\u01a5\u0001\u0000\u0000\u0000B\u01b6"+
		"\u0001\u0000\u0000\u0000D\u01b8\u0001\u0000\u0000\u0000F\u01c5\u0001\u0000"+
		"\u0000\u0000H\u01c8\u0001\u0000\u0000\u0000J\u01cc\u0001\u0000\u0000\u0000"+
		"L\u01d0\u0001\u0000\u0000\u0000N\u01e0\u0001\u0000\u0000\u0000P\u01f0"+
		"\u0001\u0000\u0000\u0000R\u01f8\u0001\u0000\u0000\u0000T\u01fa\u0001\u0000"+
		"\u0000\u0000VX\u0003\u0002\u0001\u0000WV\u0001\u0000\u0000\u0000X[\u0001"+
		"\u0000\u0000\u0000YW\u0001\u0000\u0000\u0000YZ\u0001\u0000\u0000\u0000"+
		"Z\u0001\u0001\u0000\u0000\u0000[Y\u0001\u0000\u0000\u0000\\e\u0003\b\u0004"+
		"\u0000]e\u0003\n\u0005\u0000^e\u0003\u0004\u0002\u0000_e\u0003\"\u0011"+
		"\u0000`e\u0003\u001e\u000f\u0000ae\u0003D\"\u0000be\u0003F#\u0000ce\u0003"+
		"J%\u0000d\\\u0001\u0000\u0000\u0000d]\u0001\u0000\u0000\u0000d^\u0001"+
		"\u0000\u0000\u0000d_\u0001\u0000\u0000\u0000d`\u0001\u0000\u0000\u0000"+
		"da\u0001\u0000\u0000\u0000db\u0001\u0000\u0000\u0000dc\u0001\u0000\u0000"+
		"\u0000e\u0003\u0001\u0000\u0000\u0000fm\u0003(\u0014\u0000gm\u0003*\u0015"+
		"\u0000hm\u0003,\u0016\u0000im\u0003>\u001f\u0000jm\u0003H$\u0000km\u0003"+
		".\u0017\u0000lf\u0001\u0000\u0000\u0000lg\u0001\u0000\u0000\u0000lh\u0001"+
		"\u0000\u0000\u0000li\u0001\u0000\u0000\u0000lj\u0001\u0000\u0000\u0000"+
		"lk\u0001\u0000\u0000\u0000m\u0005\u0001\u0000\u0000\u0000no\u00054\u0000"+
		"\u0000o\u0007\u0001\u0000\u0000\u0000pr\u0005J\u0000\u0000qp\u0001\u0000"+
		"\u0000\u0000qr\u0001\u0000\u0000\u0000rs\u0001\u0000\u0000\u0000st\u0003"+
		"@ \u0000tu\u0007\u0000\u0000\u0000uv\u0003B!\u0000vw\u0003\u0006\u0003"+
		"\u0000w\t\u0001\u0000\u0000\u0000xy\u0006\u0005\uffff\uffff\u0000y\u0089"+
		"\u0003\u001c\u000e\u0000z\u0089\u0003<\u001e\u0000{\u0089\u0003\f\u0006"+
		"\u0000|}\u0005\u0003\u0000\u0000}~\u0003\n\u0005\u0000~\u007f\u0005\u0004"+
		"\u0000\u0000\u007f\u0089\u0001\u0000\u0000\u0000\u0080\u0089\u0003\u001e"+
		"\u000f\u0000\u0081\u0089\u00030\u0018\u0000\u0082\u0089\u00036\u001b\u0000"+
		"\u0083\u0089\u0003L&\u0000\u0084\u0085\u0005d\u0000\u0000\u0085\u0089"+
		"\u0003\n\u0005\u0003\u0086\u0087\u0005e\u0000\u0000\u0087\u0089\u0003"+
		"\n\u0005\u0001\u0088x\u0001\u0000\u0000\u0000\u0088z\u0001\u0000\u0000"+
		"\u0000\u0088{\u0001\u0000\u0000\u0000\u0088|\u0001\u0000\u0000\u0000\u0088"+
		"\u0080\u0001\u0000\u0000\u0000\u0088\u0081\u0001\u0000\u0000\u0000\u0088"+
		"\u0082\u0001\u0000\u0000\u0000\u0088\u0083\u0001\u0000\u0000\u0000\u0088"+
		"\u0084\u0001\u0000\u0000\u0000\u0088\u0086\u0001\u0000\u0000\u0000\u0089"+
		"\u009a\u0001\u0000\u0000\u0000\u008a\u008b\n\f\u0000\u0000\u008b\u008c"+
		"\u0005\u0002\u0000\u0000\u008c\u0099\u0003\n\u0005\r\u008d\u008e\n\n\u0000"+
		"\u0000\u008e\u008f\u0003\u0010\b\u0000\u008f\u0090\u0003\n\u0005\u000b"+
		"\u0090\u0099\u0001\u0000\u0000\u0000\u0091\u0092\n\u0005\u0000\u0000\u0092"+
		"\u0093\u0005m\u0000\u0000\u0093\u0099\u0003\n\u0005\u0006\u0094\u0095"+
		"\n\u0004\u0000\u0000\u0095\u0099\u0005d\u0000\u0000\u0096\u0097\n\u0002"+
		"\u0000\u0000\u0097\u0099\u0005e\u0000\u0000\u0098\u008a\u0001\u0000\u0000"+
		"\u0000\u0098\u008d\u0001\u0000\u0000\u0000\u0098\u0091\u0001\u0000\u0000"+
		"\u0000\u0098\u0094\u0001\u0000\u0000\u0000\u0098\u0096\u0001\u0000\u0000"+
		"\u0000\u0099\u009c\u0001\u0000\u0000\u0000\u009a\u0098\u0001\u0000\u0000"+
		"\u0000\u009a\u009b\u0001\u0000\u0000\u0000\u009b\u000b\u0001\u0000\u0000"+
		"\u0000\u009c\u009a\u0001\u0000\u0000\u0000\u009d\u00ab\u0003\u000e\u0007"+
		"\u0000\u009e\u009f\u0005B\u0000\u0000\u009f\u00ab\u0003\f\u0006\u0000"+
		"\u00a0\u00a1\u0005\u0005\u0000\u0000\u00a1\u00ab\u0003\f\u0006\u0000\u00a2"+
		"\u00a3\u0005\u0006\u0000\u0000\u00a3\u00ab\u0003\f\u0006\u0000\u00a4\u00a5"+
		"\u0005\u0007\u0000\u0000\u00a5\u00ab\u0003\f\u0006\u0000\u00a6\u00a7\u0005"+
		"d\u0000\u0000\u00a7\u00ab\u0003\f\u0006\u0000\u00a8\u00a9\u0005e\u0000"+
		"\u0000\u00a9\u00ab\u0003\f\u0006\u0000\u00aa\u009d\u0001\u0000\u0000\u0000"+
		"\u00aa\u009e\u0001\u0000\u0000\u0000\u00aa\u00a0\u0001\u0000\u0000\u0000"+
		"\u00aa\u00a2\u0001\u0000\u0000\u0000\u00aa\u00a4\u0001\u0000\u0000\u0000"+
		"\u00aa\u00a6\u0001\u0000\u0000\u0000\u00aa\u00a8\u0001\u0000\u0000\u0000"+
		"\u00ab\r\u0001\u0000\u0000\u0000\u00ac\u00ad\u0006\u0007\uffff\uffff\u0000"+
		"\u00ad\u00b7\u0003\u001c\u000e\u0000\u00ae\u00b7\u0003<\u001e\u0000\u00af"+
		"\u00b0\u0005\u0003\u0000\u0000\u00b0\u00b1\u0003\n\u0005\u0000\u00b1\u00b2"+
		"\u0005\u0004\u0000\u0000\u00b2\u00b7\u0001\u0000\u0000\u0000\u00b3\u00b7"+
		"\u0003\u001e\u000f\u0000\u00b4\u00b7\u00030\u0018\u0000\u00b5\u00b7\u0003"+
		"6\u001b\u0000\u00b6\u00ac\u0001\u0000\u0000\u0000\u00b6\u00ae\u0001\u0000"+
		"\u0000\u0000\u00b6\u00af\u0001\u0000\u0000\u0000\u00b6\u00b3\u0001\u0000"+
		"\u0000\u0000\u00b6\u00b4\u0001\u0000\u0000\u0000\u00b6\u00b5\u0001\u0000"+
		"\u0000\u0000\u00b7\u00be\u0001\u0000\u0000\u0000\u00b8\u00b9\n\u0002\u0000"+
		"\u0000\u00b9\u00bd\u0005d\u0000\u0000\u00ba\u00bb\n\u0001\u0000\u0000"+
		"\u00bb\u00bd\u0005e\u0000\u0000\u00bc\u00b8\u0001\u0000\u0000\u0000\u00bc"+
		"\u00ba\u0001\u0000\u0000\u0000\u00bd\u00c0\u0001\u0000\u0000\u0000\u00be"+
		"\u00bc\u0001\u0000\u0000\u0000\u00be\u00bf\u0001\u0000\u0000\u0000\u00bf"+
		"\u000f\u0001\u0000\u0000\u0000\u00c0\u00be\u0001\u0000\u0000\u0000\u00c1"+
		"\u00c7\u0003\u0012\t\u0000\u00c2\u00c7\u0003\u0014\n\u0000\u00c3\u00c7"+
		"\u0003\u0016\u000b\u0000\u00c4\u00c7\u0003\u0018\f\u0000\u00c5\u00c7\u0003"+
		"\u001a\r\u0000\u00c6\u00c1\u0001\u0000\u0000\u0000\u00c6\u00c2\u0001\u0000"+
		"\u0000\u0000\u00c6\u00c3\u0001\u0000\u0000\u0000\u00c6\u00c4\u0001\u0000"+
		"\u0000\u0000\u00c6\u00c5\u0001\u0000\u0000\u0000\u00c7\u0011\u0001\u0000"+
		"\u0000\u0000\u00c8\u00c9\u0007\u0001\u0000\u0000\u00c9\u0013\u0001\u0000"+
		"\u0000\u0000\u00ca\u00cb\u0007\u0002\u0000\u0000\u00cb\u0015\u0001\u0000"+
		"\u0000\u0000\u00cc\u00cd\u0007\u0003\u0000\u0000\u00cd\u0017\u0001\u0000"+
		"\u0000\u0000\u00ce\u00cf\u0007\u0004\u0000\u0000\u00cf\u0019\u0001\u0000"+
		"\u0000\u0000\u00d0\u00d1\u0005\u0017\u0000\u0000\u00d1\u001b\u0001\u0000"+
		"\u0000\u0000\u00d2\u00d3\u0007\u0005\u0000\u0000\u00d3\u001d\u0001\u0000"+
		"\u0000\u0000\u00d4\u00d5\u0007\u0006\u0000\u0000\u00d5\u00d6\u0005\u0003"+
		"\u0000\u0000\u00d6\u00d7\u0003B!\u0000\u00d7\u00d8\u0005\u0004\u0000\u0000"+
		"\u00d8\u00f8\u0001\u0000\u0000\u0000\u00d9\u00da\u0007\u0007\u0000\u0000"+
		"\u00da\u00db\u0005\u0003\u0000\u0000\u00db\u00dc\u0003B!\u0000\u00dc\u00dd"+
		"\u0005\u0004\u0000\u0000\u00dd\u00f8\u0001\u0000\u0000\u0000\u00de\u00e5"+
		"\u0003<\u001e\u0000\u00df\u00e5\u00036\u001b\u0000\u00e0\u00e1\u0005\u0003"+
		"\u0000\u0000\u00e1\u00e2\u0003\n\u0005\u0000\u00e2\u00e3\u0005\u0004\u0000"+
		"\u0000\u00e3\u00e5\u0001\u0000\u0000\u0000\u00e4\u00de\u0001\u0000\u0000"+
		"\u0000\u00e4\u00df\u0001\u0000\u0000\u0000\u00e4\u00e0\u0001\u0000\u0000"+
		"\u0000\u00e5\u00e8\u0001\u0000\u0000\u0000\u00e6\u00e7\u0005\u0018\u0000"+
		"\u0000\u00e7\u00e9\u0003<\u001e\u0000\u00e8\u00e6\u0001\u0000\u0000\u0000"+
		"\u00e8\u00e9\u0001\u0000\u0000\u0000\u00e9\u00ea\u0001\u0000\u0000\u0000"+
		"\u00ea\u00ec\u0005\u0003\u0000\u0000\u00eb\u00ed\u0003B!\u0000\u00ec\u00eb"+
		"\u0001\u0000\u0000\u0000\u00ec\u00ed\u0001\u0000\u0000\u0000\u00ed\u00ee"+
		"\u0001\u0000\u0000\u0000\u00ee\u00ef\u0005\u0004\u0000\u0000\u00ef\u00f8"+
		"\u0001\u0000\u0000\u0000\u00f0\u00f8\u0003 \u0010\u0000\u00f1\u00f2\u0005"+
		"6\u0000\u0000\u00f2\u00f3\u0005\u0003\u0000\u0000\u00f3\u00f4\u0003B!"+
		"\u0000\u00f4\u00f5\u0005\u0004\u0000\u0000\u00f5\u00f8\u0001\u0000\u0000"+
		"\u0000\u00f6\u00f8\u0003N\'\u0000\u00f7\u00d4\u0001\u0000\u0000\u0000"+
		"\u00f7\u00d9\u0001\u0000\u0000\u0000\u00f7\u00e4\u0001\u0000\u0000\u0000"+
		"\u00f7\u00f0\u0001\u0000\u0000\u0000\u00f7\u00f1\u0001\u0000\u0000\u0000"+
		"\u00f7\u00f6\u0001\u0000\u0000\u0000\u00f8\u001f\u0001\u0000\u0000\u0000"+
		"\u00f9\u00fa\u0005\u0019\u0000\u0000\u00fa\u00fb\u0005\u0003\u0000\u0000"+
		"\u00fb\u00fc\u0003<\u001e\u0000\u00fc\u00fd\u0005\u001a\u0000\u0000\u00fd"+
		"\u00fe\u0003\n\u0005\u0000\u00fe\u00ff\u0005\u0004\u0000\u0000\u00ff!"+
		"\u0001\u0000\u0000\u0000\u0100\u0102\u0005J\u0000\u0000\u0101\u0100\u0001"+
		"\u0000\u0000\u0000\u0101\u0102\u0001\u0000\u0000\u0000\u0102\u0103\u0001"+
		"\u0000\u0000\u0000\u0103\u0108\u0005K\u0000\u0000\u0104\u0105\u0003<\u001e"+
		"\u0000\u0105\u0106\u0005\u001b\u0000\u0000\u0106\u0107\u0003<\u001e\u0000"+
		"\u0107\u0109\u0001\u0000\u0000\u0000\u0108\u0104\u0001\u0000\u0000\u0000"+
		"\u0108\u0109\u0001\u0000\u0000\u0000\u0109\u010a\u0001\u0000\u0000\u0000"+
		"\u010a\u010b\u0003<\u001e\u0000\u010b\u0115\u0005\u0003\u0000\u0000\u010c"+
		"\u0111\u0003$\u0012\u0000\u010d\u010e\u0005\u001a\u0000\u0000\u010e\u0110"+
		"\u0003$\u0012\u0000\u010f\u010d\u0001\u0000\u0000\u0000\u0110\u0113\u0001"+
		"\u0000\u0000\u0000\u0111\u010f\u0001\u0000\u0000\u0000\u0111\u0112\u0001"+
		"\u0000\u0000\u0000\u0112\u0116\u0001\u0000\u0000\u0000\u0113\u0111\u0001"+
		"\u0000\u0000\u0000\u0114\u0116\u0005c\u0000\u0000\u0115\u010c\u0001\u0000"+
		"\u0000\u0000\u0115\u0114\u0001\u0000\u0000\u0000\u0115\u0116\u0001\u0000"+
		"\u0000\u0000\u0116\u0117\u0001\u0000\u0000\u0000\u0117\u0118\u0005\u0004"+
		"\u0000\u0000\u0118\u0119\u0003&\u0013\u0000\u0119\u011a\u0005<\u0000\u0000"+
		"\u011a#\u0001\u0000\u0000\u0000\u011b\u011e\u0003<\u001e\u0000\u011c\u011d"+
		"\u0005\u0001\u0000\u0000\u011d\u011f\u0003\n\u0005\u0000\u011e\u011c\u0001"+
		"\u0000\u0000\u0000\u011e\u011f\u0001\u0000\u0000\u0000\u011f%\u0001\u0000"+
		"\u0000\u0000\u0120\u0122\u0003\u0006\u0003\u0000\u0121\u0120\u0001\u0000"+
		"\u0000\u0000\u0122\u0125\u0001\u0000\u0000\u0000\u0123\u0121\u0001\u0000"+
		"\u0000\u0000\u0123\u0124\u0001\u0000\u0000\u0000\u0124\u0126\u0001\u0000"+
		"\u0000\u0000\u0125\u0123\u0001\u0000\u0000\u0000\u0126\u0128\u0003\u0002"+
		"\u0001\u0000\u0127\u0129\u0003\u0006\u0003\u0000\u0128\u0127\u0001\u0000"+
		"\u0000\u0000\u0128\u0129\u0001\u0000\u0000\u0000\u0129\u012b\u0001\u0000"+
		"\u0000\u0000\u012a\u0123\u0001\u0000\u0000\u0000\u012b\u012c\u0001\u0000"+
		"\u0000\u0000\u012c\u012a\u0001\u0000\u0000\u0000\u012c\u012d\u0001\u0000"+
		"\u0000\u0000\u012d\'\u0001\u0000\u0000\u0000\u012e\u012f\u0005@\u0000"+
		"\u0000\u012f\u0130\u0003\n\u0005\u0000\u0130\u0131\u0005F\u0000\u0000"+
		"\u0131\u0139\u0003&\u0013\u0000\u0132\u0133\u0005;\u0000\u0000\u0133\u0134"+
		"\u0003\n\u0005\u0000\u0134\u0135\u0005F\u0000\u0000\u0135\u0136\u0003"+
		"&\u0013\u0000\u0136\u0138\u0001\u0000\u0000\u0000\u0137\u0132\u0001\u0000"+
		"\u0000\u0000\u0138\u013b\u0001\u0000\u0000\u0000\u0139\u0137\u0001\u0000"+
		"\u0000\u0000\u0139\u013a\u0001\u0000\u0000\u0000\u013a\u013e\u0001\u0000"+
		"\u0000\u0000\u013b\u0139\u0001\u0000\u0000\u0000\u013c\u013d\u0005:\u0000"+
		"\u0000\u013d\u013f\u0003&\u0013\u0000\u013e\u013c\u0001\u0000\u0000\u0000"+
		"\u013e\u013f\u0001\u0000\u0000\u0000\u013f\u0140\u0001\u0000\u0000\u0000"+
		"\u0140\u0141\u0005<\u0000\u0000\u0141)\u0001\u0000\u0000\u0000\u0142\u0143"+
		"\u0005>\u0000\u0000\u0143\u0144\u0003<\u001e\u0000\u0144\u0145\u00055"+
		"\u0000\u0000\u0145\u0146\u0003\n\u0005\u0000\u0146\u0147\u00059\u0000"+
		"\u0000\u0147\u0148\u0003&\u0013\u0000\u0148\u0149\u0005<\u0000\u0000\u0149"+
		"\u0159\u0001\u0000\u0000\u0000\u014a\u014b\u0005>\u0000\u0000\u014b\u014c"+
		"\u0003<\u001e\u0000\u014c\u014d\u0005\u0001\u0000\u0000\u014d\u014e\u0003"+
		"\n\u0005\u0000\u014e\u014f\u0005\u001a\u0000\u0000\u014f\u0152\u0003\n"+
		"\u0005\u0000\u0150\u0151\u0005\u001a\u0000\u0000\u0151\u0153\u0003\n\u0005"+
		"\u0000\u0152\u0150\u0001\u0000\u0000\u0000\u0152\u0153\u0001\u0000\u0000"+
		"\u0000\u0153\u0154\u0001\u0000\u0000\u0000\u0154\u0155\u00059\u0000\u0000"+
		"\u0155\u0156\u0003&\u0013\u0000\u0156\u0157\u0005<\u0000\u0000\u0157\u0159"+
		"\u0001\u0000\u0000\u0000\u0158\u0142\u0001\u0000\u0000\u0000\u0158\u014a"+
		"\u0001\u0000\u0000\u0000\u0159+\u0001\u0000\u0000\u0000\u015a\u015b\u0005"+
		"I\u0000\u0000\u015b\u015c\u0003\n\u0005\u0000\u015c\u015d\u00059\u0000"+
		"\u0000\u015d\u015e\u0003&\u0013\u0000\u015e\u015f\u0005<\u0000\u0000\u015f"+
		"-\u0001\u0000\u0000\u0000\u0160\u0161\u00059\u0000\u0000\u0161\u0162\u0003"+
		"&\u0013\u0000\u0162\u0163\u0005<\u0000\u0000\u0163/\u0001\u0000\u0000"+
		"\u0000\u0164\u0171\u0005\u001c\u0000\u0000\u0165\u016b\u00034\u001a\u0000"+
		"\u0166\u0167\u00032\u0019\u0000\u0167\u0168\u00034\u001a\u0000\u0168\u016a"+
		"\u0001\u0000\u0000\u0000\u0169\u0166\u0001\u0000\u0000\u0000\u016a\u016d"+
		"\u0001\u0000\u0000\u0000\u016b\u0169\u0001\u0000\u0000\u0000\u016b\u016c"+
		"\u0001\u0000\u0000\u0000\u016c\u016f\u0001\u0000\u0000\u0000\u016d\u016b"+
		"\u0001\u0000\u0000\u0000\u016e\u0170\u00032\u0019\u0000\u016f\u016e\u0001"+
		"\u0000\u0000\u0000\u016f\u0170\u0001\u0000\u0000\u0000\u0170\u0172\u0001"+
		"\u0000\u0000\u0000\u0171\u0165\u0001\u0000\u0000\u0000\u0171\u0172\u0001"+
		"\u0000\u0000\u0000\u0172\u0173\u0001\u0000\u0000\u0000\u0173\u0174\u0005"+
		"\u001d\u0000\u0000\u01741\u0001\u0000\u0000\u0000\u0175\u0176\u0007\b"+
		"\u0000\u0000\u01763\u0001\u0000\u0000\u0000\u0177\u0178\u0003<\u001e\u0000"+
		"\u0178\u0179\u0005\u0001\u0000\u0000\u0179\u017a\u0003\n\u0005\u0000\u017a"+
		"\u0183\u0001\u0000\u0000\u0000\u017b\u017c\u0005\u001f\u0000\u0000\u017c"+
		"\u017d\u0003\n\u0005\u0000\u017d\u017e\u0005 \u0000\u0000\u017e\u017f"+
		"\u0005\u0001\u0000\u0000\u017f\u0180\u0003\n\u0005\u0000\u0180\u0183\u0001"+
		"\u0000\u0000\u0000\u0181\u0183\u0003\n\u0005\u0000\u0182\u0177\u0001\u0000"+
		"\u0000\u0000\u0182\u017b\u0001\u0000\u0000\u0000\u0182\u0181\u0001\u0000"+
		"\u0000\u0000\u01835\u0001\u0000\u0000\u0000\u0184\u0185\u0003<\u001e\u0000"+
		"\u0185\u0186\u0005\u001f\u0000\u0000\u0186\u0187\u0003\n\u0005\u0000\u0187"+
		"\u0188\u0005 \u0000\u0000\u0188\u018e\u0001\u0000\u0000\u0000\u0189\u018a"+
		"\u0003<\u001e\u0000\u018a\u018b\u0005\u001b\u0000\u0000\u018b\u018c\u0003"+
		"<\u001e\u0000\u018c\u018e\u0001\u0000\u0000\u0000\u018d\u0184\u0001\u0000"+
		"\u0000\u0000\u018d\u0189\u0001\u0000\u0000\u0000\u018e7\u0001\u0000\u0000"+
		"\u0000\u018f\u0190\u0005e\u0000\u0000\u01909\u0001\u0000\u0000\u0000\u0191"+
		"\u0192\u00056\u0000\u0000\u0192\u0193\u0005\u0003\u0000\u0000\u0193\u0194"+
		"\u0003\n\u0005\u0000\u0194\u0195\u0005\u0004\u0000\u0000\u0195;\u0001"+
		"\u0000\u0000\u0000\u0196\u019a\u0005^\u0000\u0000\u0197\u0199\u0007\t"+
		"\u0000\u0000\u0198\u0197\u0001\u0000\u0000\u0000\u0199\u019c\u0001\u0000"+
		"\u0000\u0000\u019a\u0198\u0001\u0000\u0000\u0000\u019a\u019b\u0001\u0000"+
		"\u0000\u0000\u019b=\u0001\u0000\u0000\u0000\u019c\u019a\u0001\u0000\u0000"+
		"\u0000\u019d\u019e\u0005D\u0000\u0000\u019e\u019f\u0003&\u0013\u0000\u019f"+
		"\u01a0\u0005H\u0000\u0000\u01a0\u01a3\u0003\n\u0005\u0000\u01a1\u01a4"+
		"\u0003\u0006\u0003\u0000\u01a2\u01a4\u0005\u0000\u0000\u0001\u01a3\u01a1"+
		"\u0001\u0000\u0000\u0000\u01a3\u01a2\u0001\u0000\u0000\u0000\u01a4?\u0001"+
		"\u0000\u0000\u0000\u01a5\u01aa\u0003<\u001e\u0000\u01a6\u01a7\u0005\u001a"+
		"\u0000\u0000\u01a7\u01a9\u0003<\u001e\u0000\u01a8\u01a6\u0001\u0000\u0000"+
		"\u0000\u01a9\u01ac\u0001\u0000\u0000\u0000\u01aa\u01a8\u0001\u0000\u0000"+
		"\u0000\u01aa\u01ab\u0001\u0000\u0000\u0000\u01abA\u0001\u0000\u0000\u0000"+
		"\u01ac\u01aa\u0001\u0000\u0000\u0000\u01ad\u01b2\u0003\n\u0005\u0000\u01ae"+
		"\u01af\u0005\u001a\u0000\u0000\u01af\u01b1\u0003\n\u0005\u0000\u01b0\u01ae"+
		"\u0001\u0000\u0000\u0000\u01b1\u01b4\u0001\u0000\u0000\u0000\u01b2\u01b0"+
		"\u0001\u0000\u0000\u0000\u01b2\u01b3\u0001\u0000\u0000\u0000\u01b3\u01b7"+
		"\u0001\u0000\u0000\u0000\u01b4\u01b2\u0001\u0000\u0000\u0000\u01b5\u01b7"+
		"\u0005c\u0000\u0000\u01b6\u01ad\u0001\u0000\u0000\u0000\u01b6\u01b5\u0001"+
		"\u0000\u0000\u0000\u01b7C\u0001\u0000\u0000\u0000\u01b8\u01c1\u0005E\u0000"+
		"\u0000\u01b9\u01be\u0003\n\u0005\u0000\u01ba\u01bb\u0005\u001a\u0000\u0000"+
		"\u01bb\u01bd\u0003\n\u0005\u0000\u01bc\u01ba\u0001\u0000\u0000\u0000\u01bd"+
		"\u01c0\u0001\u0000\u0000\u0000\u01be\u01bc\u0001\u0000\u0000\u0000\u01be"+
		"\u01bf\u0001\u0000\u0000\u0000\u01bf\u01c2\u0001\u0000\u0000\u0000\u01c0"+
		"\u01be\u0001\u0000\u0000\u0000\u01c1\u01b9\u0001\u0000\u0000\u0000\u01c1"+
		"\u01c2\u0001\u0000\u0000\u0000\u01c2\u01c3\u0001\u0000\u0000\u0000\u01c3"+
		"\u01c4\u0003\u0006\u0003\u0000\u01c4E\u0001\u0000\u0000\u0000\u01c5\u01c6"+
		"\u00058\u0000\u0000\u01c6\u01c7\u0003\u0006\u0003\u0000\u01c7G\u0001\u0000"+
		"\u0000\u0000\u01c8\u01c9\u0005?\u0000\u0000\u01c9\u01ca\u0003<\u001e\u0000"+
		"\u01ca\u01cb\u0003\u0006\u0003\u0000\u01cbI\u0001\u0000\u0000\u0000\u01cc"+
		"\u01cd\u0005\"\u0000\u0000\u01cd\u01ce\u0003<\u001e\u0000\u01ce\u01cf"+
		"\u0005\"\u0000\u0000\u01cfK\u0001\u0000\u0000\u0000\u01d0\u01d1\u0005"+
		"K\u0000\u0000\u01d1\u01da\u0005\u0003\u0000\u0000\u01d2\u01d7\u0003<\u001e"+
		"\u0000\u01d3\u01d4\u0005\u001a\u0000\u0000\u01d4\u01d6\u0003<\u001e\u0000"+
		"\u01d5\u01d3\u0001\u0000\u0000\u0000\u01d6\u01d9\u0001\u0000\u0000\u0000"+
		"\u01d7\u01d5\u0001\u0000\u0000\u0000\u01d7\u01d8\u0001\u0000\u0000\u0000"+
		"\u01d8\u01db\u0001\u0000\u0000\u0000\u01d9\u01d7\u0001\u0000\u0000\u0000"+
		"\u01da\u01d2\u0001\u0000\u0000\u0000\u01da\u01db\u0001\u0000\u0000\u0000"+
		"\u01db\u01dc\u0001\u0000\u0000\u0000\u01dc\u01dd\u0005\u0004\u0000\u0000"+
		"\u01dd\u01de\u0003&\u0013\u0000\u01de\u01df\u0005<\u0000\u0000\u01dfM"+
		"\u0001\u0000\u0000\u0000\u01e0\u01e1\u0003<\u001e\u0000\u01e1\u01e2\u0005"+
		"\u0018\u0000\u0000\u01e2\u01e3\u0003<\u001e\u0000\u01e3\u01ec\u0005\u0003"+
		"\u0000\u0000\u01e4\u01e9\u0003\n\u0005\u0000\u01e5\u01e6\u0005\u001a\u0000"+
		"\u0000\u01e6\u01e8\u0003\n\u0005\u0000\u01e7\u01e5\u0001\u0000\u0000\u0000"+
		"\u01e8\u01eb\u0001\u0000\u0000\u0000\u01e9\u01e7\u0001\u0000\u0000\u0000"+
		"\u01e9\u01ea\u0001\u0000\u0000\u0000\u01ea\u01ed\u0001\u0000\u0000\u0000"+
		"\u01eb\u01e9\u0001\u0000\u0000\u0000\u01ec\u01e4\u0001\u0000\u0000\u0000"+
		"\u01ec\u01ed\u0001\u0000\u0000\u0000\u01ed\u01ee\u0001\u0000\u0000\u0000"+
		"\u01ee\u01ef\u0005\u0004\u0000\u0000\u01efO\u0001\u0000\u0000\u0000\u01f0"+
		"\u01f3\u0005#\u0000\u0000\u01f1\u01f4\u0003R)\u0000\u01f2\u01f4\u0003"+
		"<\u001e\u0000\u01f3\u01f1\u0001\u0000\u0000\u0000\u01f3\u01f2\u0001\u0000"+
		"\u0000\u0000\u01f4\u01f5\u0001\u0000\u0000\u0000\u01f5\u01f6\u0005\u0001"+
		"\u0000\u0000\u01f6\u01f7\u0003\n\u0005\u0000\u01f7Q\u0001\u0000\u0000"+
		"\u0000\u01f8\u01f9\u0007\n\u0000\u0000\u01f9S\u0001\u0000\u0000\u0000"+
		"\u01fa\u01fb\u0005Q\u0000\u0000\u01fb\u01fc\u0005\u001b\u0000\u0000\u01fc"+
		"\u01fd\u0007\u000b\u0000\u0000\u01fdU\u0001\u0000\u0000\u0000-Ydlq\u0088"+
		"\u0098\u009a\u00aa\u00b6\u00bc\u00be\u00c6\u00e4\u00e8\u00ec\u00f7\u0101"+
		"\u0108\u0111\u0115\u011e\u0123\u0128\u012c\u0139\u013e\u0152\u0158\u016b"+
		"\u016f\u0171\u0182\u018d\u019a\u01a3\u01aa\u01b2\u01b6\u01be\u01c1\u01d7"+
		"\u01da\u01e9\u01ec\u01f3";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}