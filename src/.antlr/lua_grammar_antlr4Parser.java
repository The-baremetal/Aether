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
		T__31=32, SEPARATOR=33, KW_IN=34, KW_PRINT=35, KW_AND=36, KW_BREAK=37, 
		KW_DO=38, KW_ELSE=39, KW_ELSEIF=40, KW_END=41, KW_FALSE=42, KW_FOR=43, 
		KW_GOTO=44, KW_IF=45, KW_NIL=46, KW_NOT=47, KW_OR=48, KW_REPEAT=49, KW_RETURN=50, 
		KW_THEN=51, KW_TRUE=52, KW_UNTIL=53, KW_WHILE=54, KW_LOCAL=55, KW_FUNCTION=56, 
		KW_INDEX=57, KW_NEWINDEX=58, KW_MODE=59, NUMBER=60, STRING=61, LETTER=62, 
		DIGIT=63, WS=64, SINGLE_LINE_COMMENT=65, MULTI_LINE_COMMENT=66;
	public static final int
		RULE_program = 0, RULE_statement = 1, RULE_control_statement = 2, RULE_statement_terminator = 3, 
		RULE_assignment = 4, RULE_expression = 5, RULE_prefix_expression = 6, 
		RULE_primary_expression = 7, RULE_operators = 8, RULE_comparison_operator = 9, 
		RULE_arith_operator = 10, RULE_logical_operator = 11, RULE_bitwise_operator = 12, 
		RULE_literal = 13, RULE_function_call = 14, RULE_table_insert = 15, RULE_function_declaration = 16, 
		RULE_block = 17, RULE_if_statement = 18, RULE_for_statement = 19, RULE_while_statement = 20, 
		RULE_table = 21, RULE_field = 22, RULE_table_access = 23, RULE_single_line_comment = 24, 
		RULE_print_statement = 25, RULE_identifier = 26, RULE_repeat_statement = 27, 
		RULE_identifier_list = 28, RULE_expression_list = 29, RULE_return_statement = 30, 
		RULE_break_statement = 31, RULE_goto_statement = 32, RULE_label_statement = 33, 
		RULE_function_expression = 34, RULE_method_call = 35, RULE_metatable_field = 36;
	private static String[] makeRuleNames() {
		return new String[] {
			"program", "statement", "control_statement", "statement_terminator", 
			"assignment", "expression", "prefix_expression", "primary_expression", 
			"operators", "comparison_operator", "arith_operator", "logical_operator", 
			"bitwise_operator", "literal", "function_call", "table_insert", "function_declaration", 
			"block", "if_statement", "for_statement", "while_statement", "table", 
			"field", "table_access", "single_line_comment", "print_statement", "identifier", 
			"repeat_statement", "identifier_list", "expression_list", "return_statement", 
			"break_statement", "goto_statement", "label_statement", "function_expression", 
			"method_call", "metatable_field"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'='", "'('", "')'", "'#'", "'>'", "'<'", "'>='", "'=='", "'<='", 
			"'~='", "'*'", "'/'", "'+'", "'-'", "'//'", "'&'", "'|'", "'~'", "'<<'", 
			"'>>'", "':'", "'table.insert'", "','", "'{'", "'}'", "'['", "']'", "'.'", 
			"'--'", "'_'", "'::'", "'__'", null, "'in'", "'print'", "'and'", "'break'", 
			"'do'", "'else'", "'elseif'", "'end'", "'false'", "'for'", "'goto'", 
			"'if'", "'nil'", "'not'", "'or'", "'repeat'", "'return'", "'then'", "'true'", 
			"'until'", "'while'", "'local'", "'function'", "'index'", "'newindex'", 
			"'mode'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, null, "SEPARATOR", "KW_IN", 
			"KW_PRINT", "KW_AND", "KW_BREAK", "KW_DO", "KW_ELSE", "KW_ELSEIF", "KW_END", 
			"KW_FALSE", "KW_FOR", "KW_GOTO", "KW_IF", "KW_NIL", "KW_NOT", "KW_OR", 
			"KW_REPEAT", "KW_RETURN", "KW_THEN", "KW_TRUE", "KW_UNTIL", "KW_WHILE", 
			"KW_LOCAL", "KW_FUNCTION", "KW_INDEX", "KW_NEWINDEX", "KW_MODE", "NUMBER", 
			"STRING", "LETTER", "DIGIT", "WS", "SINGLE_LINE_COMMENT", "MULTI_LINE_COMMENT"
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
			setState(77);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 8202880284710928388L) != 0)) {
				{
				{
				setState(74);
				statement();
				}
				}
				setState(79);
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
			setState(88);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,1,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(80);
				assignment();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(81);
				expression(0);
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(82);
				control_statement();
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(83);
				function_declaration();
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(84);
				function_call();
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(85);
				return_statement();
				}
				break;
			case 7:
				enterOuterAlt(_localctx, 7);
				{
				setState(86);
				break_statement();
				}
				break;
			case 8:
				enterOuterAlt(_localctx, 8);
				{
				setState(87);
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
		public Control_statementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_control_statement; }
	}

	public final Control_statementContext control_statement() throws RecognitionException {
		Control_statementContext _localctx = new Control_statementContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_control_statement);
		try {
			setState(95);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case KW_IF:
				enterOuterAlt(_localctx, 1);
				{
				setState(90);
				if_statement();
				}
				break;
			case KW_FOR:
				enterOuterAlt(_localctx, 2);
				{
				setState(91);
				for_statement();
				}
				break;
			case KW_WHILE:
				enterOuterAlt(_localctx, 3);
				{
				setState(92);
				while_statement();
				}
				break;
			case KW_REPEAT:
				enterOuterAlt(_localctx, 4);
				{
				setState(93);
				repeat_statement();
				}
				break;
			case KW_GOTO:
				enterOuterAlt(_localctx, 5);
				{
				setState(94);
				goto_statement();
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
			setState(97);
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
			setState(100);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==KW_LOCAL) {
				{
				setState(99);
				match(KW_LOCAL);
				}
			}

			setState(102);
			identifier_list();
			setState(103);
			match(T__0);
			setState(104);
			expression_list();
			setState(105);
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
			setState(118);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,4,_ctx) ) {
			case 1:
				{
				setState(108);
				literal();
				}
				break;
			case 2:
				{
				setState(109);
				identifier();
				}
				break;
			case 3:
				{
				setState(110);
				match(T__1);
				setState(111);
				expression(0);
				setState(112);
				match(T__2);
				}
				break;
			case 4:
				{
				setState(114);
				function_call();
				}
				break;
			case 5:
				{
				setState(115);
				table();
				}
				break;
			case 6:
				{
				setState(116);
				table_access();
				}
				break;
			case 7:
				{
				setState(117);
				function_expression();
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(126);
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
					setState(120);
					if (!(precpred(_ctx, 5))) throw new FailedPredicateException(this, "precpred(_ctx, 5)");
					setState(121);
					operators();
					setState(122);
					expression(6);
					}
					} 
				}
				setState(128);
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
			setState(134);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__1:
			case T__21:
			case T__23:
			case KW_PRINT:
			case KW_FALSE:
			case KW_NIL:
			case KW_TRUE:
			case NUMBER:
			case STRING:
			case LETTER:
				enterOuterAlt(_localctx, 1);
				{
				setState(129);
				primary_expression();
				}
				break;
			case KW_NOT:
				enterOuterAlt(_localctx, 2);
				{
				setState(130);
				match(KW_NOT);
				setState(131);
				prefix_expression();
				}
				break;
			case T__3:
				enterOuterAlt(_localctx, 3);
				{
				setState(132);
				match(T__3);
				setState(133);
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
			setState(145);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,7,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(136);
				literal();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(137);
				identifier();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(138);
				match(T__1);
				setState(139);
				expression(0);
				setState(140);
				match(T__2);
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(142);
				function_call();
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(143);
				table();
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(144);
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
		public OperatorsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_operators; }
	}

	public final OperatorsContext operators() throws RecognitionException {
		OperatorsContext _localctx = new OperatorsContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_operators);
		try {
			setState(151);
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
				setState(147);
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
				setState(148);
				arith_operator();
				}
				break;
			case KW_AND:
			case KW_OR:
				enterOuterAlt(_localctx, 3);
				{
				setState(149);
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
				setState(150);
				bitwise_operator();
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
			setState(153);
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
			setState(155);
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
			setState(157);
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
			setState(159);
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
	public static class LiteralContext extends ParserRuleContext {
		public TerminalNode NUMBER() { return getToken(lua_grammar_antlr4Parser.NUMBER, 0); }
		public TerminalNode STRING() { return getToken(lua_grammar_antlr4Parser.STRING, 0); }
		public TerminalNode KW_TRUE() { return getToken(lua_grammar_antlr4Parser.KW_TRUE, 0); }
		public TerminalNode KW_FALSE() { return getToken(lua_grammar_antlr4Parser.KW_FALSE, 0); }
		public TerminalNode KW_NIL() { return getToken(lua_grammar_antlr4Parser.KW_NIL, 0); }
		public LiteralContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_literal; }
	}

	public final LiteralContext literal() throws RecognitionException {
		LiteralContext _localctx = new LiteralContext(_ctx, getState());
		enterRule(_localctx, 26, RULE_literal);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(161);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 3463342880238600192L) != 0)) ) {
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
		public Expression_listContext expression_list() {
			return getRuleContext(Expression_listContext.class,0);
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
		enterRule(_localctx, 28, RULE_function_call);
		int _la;
		try {
			setState(188);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,12,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(169);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,9,_ctx) ) {
				case 1:
					{
					setState(163);
					identifier();
					}
					break;
				case 2:
					{
					setState(164);
					table_access();
					}
					break;
				case 3:
					{
					setState(165);
					match(T__1);
					setState(166);
					expression(0);
					setState(167);
					match(T__2);
					}
					break;
				}
				setState(173);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==T__20) {
					{
					setState(171);
					match(T__20);
					setState(172);
					identifier();
					}
				}

				setState(175);
				match(T__1);
				setState(177);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 8147086527084625924L) != 0)) {
					{
					setState(176);
					expression_list();
					}
				}

				setState(179);
				match(T__2);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(181);
				table_insert();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(182);
				match(KW_PRINT);
				setState(183);
				match(T__1);
				setState(184);
				expression_list();
				setState(185);
				match(T__2);
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(187);
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
		enterRule(_localctx, 30, RULE_table_insert);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(190);
			match(T__21);
			setState(191);
			match(T__1);
			setState(192);
			identifier();
			setState(193);
			match(T__22);
			setState(194);
			expression(0);
			setState(195);
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
		public Function_declarationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_function_declaration; }
	}

	public final Function_declarationContext function_declaration() throws RecognitionException {
		Function_declarationContext _localctx = new Function_declarationContext(_ctx, getState());
		enterRule(_localctx, 32, RULE_function_declaration);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(197);
			match(KW_FUNCTION);
			setState(198);
			identifier();
			setState(199);
			match(T__1);
			setState(208);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==LETTER) {
				{
				setState(200);
				identifier();
				setState(205);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==T__22) {
					{
					{
					setState(201);
					match(T__22);
					setState(202);
					identifier();
					}
					}
					setState(207);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				}
			}

			setState(210);
			match(T__2);
			setState(211);
			block();
			setState(212);
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
		enterRule(_localctx, 34, RULE_block);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(218); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(214);
				statement();
				setState(216);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==SEPARATOR) {
					{
					setState(215);
					statement_terminator();
					}
				}

				}
				}
				setState(220); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & 8202880284710928388L) != 0) );
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
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
		enterRule(_localctx, 36, RULE_if_statement);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(222);
			match(KW_IF);
			setState(223);
			expression(0);
			setState(224);
			match(KW_THEN);
			setState(225);
			block();
			setState(233);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==KW_ELSEIF) {
				{
				{
				setState(226);
				match(KW_ELSEIF);
				setState(227);
				expression(0);
				setState(228);
				match(KW_THEN);
				setState(229);
				block();
				}
				}
				setState(235);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(238);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==KW_ELSE) {
				{
				setState(236);
				match(KW_ELSE);
				setState(237);
				block();
				}
			}

			setState(240);
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
		enterRule(_localctx, 38, RULE_for_statement);
		int _la;
		try {
			setState(264);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,20,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(242);
				match(KW_FOR);
				setState(243);
				identifier();
				setState(244);
				match(KW_IN);
				setState(245);
				expression(0);
				setState(246);
				match(KW_DO);
				setState(247);
				block();
				setState(248);
				match(KW_END);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(250);
				match(KW_FOR);
				setState(251);
				identifier();
				setState(252);
				match(T__0);
				setState(253);
				expression(0);
				setState(254);
				match(T__22);
				setState(255);
				expression(0);
				setState(258);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==T__22) {
					{
					setState(256);
					match(T__22);
					setState(257);
					expression(0);
					}
				}

				setState(260);
				match(KW_DO);
				setState(261);
				block();
				setState(262);
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
		enterRule(_localctx, 40, RULE_while_statement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(266);
			match(KW_WHILE);
			setState(267);
			expression(0);
			setState(268);
			match(KW_DO);
			setState(269);
			block();
			setState(270);
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
		enterRule(_localctx, 42, RULE_table);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(272);
			match(T__23);
			setState(286);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__1:
			case T__21:
			case T__23:
			case KW_PRINT:
			case KW_FALSE:
			case KW_NIL:
			case KW_TRUE:
			case KW_FUNCTION:
			case NUMBER:
			case STRING:
			case LETTER:
				{
				setState(273);
				field();
				setState(278);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,21,_ctx);
				while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
					if ( _alt==1 ) {
						{
						{
						setState(274);
						match(T__22);
						setState(275);
						field();
						}
						} 
					}
					setState(280);
					_errHandler.sync(this);
					_alt = getInterpreter().adaptivePredict(_input,21,_ctx);
				}
				setState(283);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==T__22) {
					{
					setState(281);
					match(T__22);
					setState(282);
					metatable_field();
					}
				}

				}
				break;
			case T__31:
				{
				setState(285);
				metatable_field();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			setState(288);
			match(T__24);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
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
		enterRule(_localctx, 44, RULE_field);
		try {
			setState(295);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,24,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(290);
				identifier();
				setState(291);
				match(T__0);
				setState(292);
				expression(0);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(294);
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
		enterRule(_localctx, 46, RULE_table_access);
		try {
			setState(306);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,25,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(297);
				identifier();
				setState(298);
				match(T__25);
				setState(299);
				expression(0);
				setState(300);
				match(T__26);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(302);
				identifier();
				setState(303);
				match(T__27);
				setState(304);
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
		enterRule(_localctx, 48, RULE_single_line_comment);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(308);
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
		enterRule(_localctx, 50, RULE_print_statement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(310);
			match(KW_PRINT);
			setState(311);
			match(T__1);
			setState(312);
			expression(0);
			setState(313);
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
		enterRule(_localctx, 52, RULE_identifier);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(315);
			match(LETTER);
			setState(319);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,26,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(316);
					_la = _input.LA(1);
					if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & -4611686017353646080L) != 0)) ) {
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
				setState(321);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,26,_ctx);
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
		public Repeat_statementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_repeat_statement; }
	}

	public final Repeat_statementContext repeat_statement() throws RecognitionException {
		Repeat_statementContext _localctx = new Repeat_statementContext(_ctx, getState());
		enterRule(_localctx, 54, RULE_repeat_statement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(322);
			match(KW_REPEAT);
			setState(323);
			block();
			setState(324);
			match(KW_UNTIL);
			setState(325);
			expression(0);
			setState(327);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,27,_ctx) ) {
			case 1:
				{
				setState(326);
				statement_terminator();
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
		enterRule(_localctx, 56, RULE_identifier_list);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(329);
			identifier();
			setState(334);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__22) {
				{
				{
				setState(330);
				match(T__22);
				setState(331);
				identifier();
				}
				}
				setState(336);
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
		public Expression_listContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expression_list; }
	}

	public final Expression_listContext expression_list() throws RecognitionException {
		Expression_listContext _localctx = new Expression_listContext(_ctx, getState());
		enterRule(_localctx, 58, RULE_expression_list);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(337);
			expression(0);
			setState(342);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__22) {
				{
				{
				setState(338);
				match(T__22);
				setState(339);
				expression(0);
				}
				}
				setState(344);
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
		enterRule(_localctx, 60, RULE_return_statement);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(345);
			match(KW_RETURN);
			setState(354);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 8147086527084625924L) != 0)) {
				{
				setState(346);
				expression(0);
				setState(351);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==T__22) {
					{
					{
					setState(347);
					match(T__22);
					setState(348);
					expression(0);
					}
					}
					setState(353);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				}
			}

			setState(356);
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
		enterRule(_localctx, 62, RULE_break_statement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(358);
			match(KW_BREAK);
			setState(359);
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
		enterRule(_localctx, 64, RULE_goto_statement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(361);
			match(KW_GOTO);
			setState(362);
			identifier();
			setState(363);
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
		enterRule(_localctx, 66, RULE_label_statement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(365);
			match(T__30);
			setState(366);
			identifier();
			setState(367);
			match(T__30);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
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
		enterRule(_localctx, 68, RULE_function_expression);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(369);
			match(KW_FUNCTION);
			setState(370);
			match(T__1);
			setState(379);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==LETTER) {
				{
				setState(371);
				identifier();
				setState(376);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==T__22) {
					{
					{
					setState(372);
					match(T__22);
					setState(373);
					identifier();
					}
					}
					setState(378);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				}
			}

			setState(381);
			match(T__2);
			setState(382);
			block();
			setState(383);
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
		enterRule(_localctx, 70, RULE_method_call);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(385);
			identifier();
			setState(386);
			match(T__20);
			setState(387);
			identifier();
			setState(388);
			match(T__1);
			setState(397);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 8147086527084625924L) != 0)) {
				{
				setState(389);
				expression(0);
				setState(394);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==T__22) {
					{
					{
					setState(390);
					match(T__22);
					setState(391);
					expression(0);
					}
					}
					setState(396);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				}
			}

			setState(399);
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
		enterRule(_localctx, 72, RULE_metatable_field);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(401);
			match(T__31);
			setState(402);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 1008806316530991104L) != 0)) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			setState(403);
			match(T__0);
			setState(404);
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
		"\u0004\u0001B\u0197\u0002\u0000\u0007\u0000\u0002\u0001\u0007\u0001\u0002"+
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
		"#\u0007#\u0002$\u0007$\u0001\u0000\u0005\u0000L\b\u0000\n\u0000\f\u0000"+
		"O\t\u0000\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001"+
		"\u0001\u0001\u0001\u0001\u0001\u0001\u0003\u0001Y\b\u0001\u0001\u0002"+
		"\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0003\u0002`\b\u0002"+
		"\u0001\u0003\u0001\u0003\u0001\u0004\u0003\u0004e\b\u0004\u0001\u0004"+
		"\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0005\u0001\u0005"+
		"\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005"+
		"\u0001\u0005\u0001\u0005\u0001\u0005\u0003\u0005w\b\u0005\u0001\u0005"+
		"\u0001\u0005\u0001\u0005\u0001\u0005\u0005\u0005}\b\u0005\n\u0005\f\u0005"+
		"\u0080\t\u0005\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006"+
		"\u0003\u0006\u0087\b\u0006\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007"+
		"\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0003\u0007"+
		"\u0092\b\u0007\u0001\b\u0001\b\u0001\b\u0001\b\u0003\b\u0098\b\b\u0001"+
		"\t\u0001\t\u0001\n\u0001\n\u0001\u000b\u0001\u000b\u0001\f\u0001\f\u0001"+
		"\r\u0001\r\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e"+
		"\u0001\u000e\u0003\u000e\u00aa\b\u000e\u0001\u000e\u0001\u000e\u0003\u000e"+
		"\u00ae\b\u000e\u0001\u000e\u0001\u000e\u0003\u000e\u00b2\b\u000e\u0001"+
		"\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001"+
		"\u000e\u0001\u000e\u0001\u000e\u0003\u000e\u00bd\b\u000e\u0001\u000f\u0001"+
		"\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0001"+
		"\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0005"+
		"\u0010\u00cc\b\u0010\n\u0010\f\u0010\u00cf\t\u0010\u0003\u0010\u00d1\b"+
		"\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0011\u0001"+
		"\u0011\u0003\u0011\u00d9\b\u0011\u0004\u0011\u00db\b\u0011\u000b\u0011"+
		"\f\u0011\u00dc\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012"+
		"\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0005\u0012\u00e8\b\u0012"+
		"\n\u0012\f\u0012\u00eb\t\u0012\u0001\u0012\u0001\u0012\u0003\u0012\u00ef"+
		"\b\u0012\u0001\u0012\u0001\u0012\u0001\u0013\u0001\u0013\u0001\u0013\u0001"+
		"\u0013\u0001\u0013\u0001\u0013\u0001\u0013\u0001\u0013\u0001\u0013\u0001"+
		"\u0013\u0001\u0013\u0001\u0013\u0001\u0013\u0001\u0013\u0001\u0013\u0001"+
		"\u0013\u0003\u0013\u0103\b\u0013\u0001\u0013\u0001\u0013\u0001\u0013\u0001"+
		"\u0013\u0003\u0013\u0109\b\u0013\u0001\u0014\u0001\u0014\u0001\u0014\u0001"+
		"\u0014\u0001\u0014\u0001\u0014\u0001\u0015\u0001\u0015\u0001\u0015\u0001"+
		"\u0015\u0005\u0015\u0115\b\u0015\n\u0015\f\u0015\u0118\t\u0015\u0001\u0015"+
		"\u0001\u0015\u0003\u0015\u011c\b\u0015\u0001\u0015\u0003\u0015\u011f\b"+
		"\u0015\u0001\u0015\u0001\u0015\u0001\u0016\u0001\u0016\u0001\u0016\u0001"+
		"\u0016\u0001\u0016\u0003\u0016\u0128\b\u0016\u0001\u0017\u0001\u0017\u0001"+
		"\u0017\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017\u0001"+
		"\u0017\u0003\u0017\u0133\b\u0017\u0001\u0018\u0001\u0018\u0001\u0019\u0001"+
		"\u0019\u0001\u0019\u0001\u0019\u0001\u0019\u0001\u001a\u0001\u001a\u0005"+
		"\u001a\u013e\b\u001a\n\u001a\f\u001a\u0141\t\u001a\u0001\u001b\u0001\u001b"+
		"\u0001\u001b\u0001\u001b\u0001\u001b\u0003\u001b\u0148\b\u001b\u0001\u001c"+
		"\u0001\u001c\u0001\u001c\u0005\u001c\u014d\b\u001c\n\u001c\f\u001c\u0150"+
		"\t\u001c\u0001\u001d\u0001\u001d\u0001\u001d\u0005\u001d\u0155\b\u001d"+
		"\n\u001d\f\u001d\u0158\t\u001d\u0001\u001e\u0001\u001e\u0001\u001e\u0001"+
		"\u001e\u0005\u001e\u015e\b\u001e\n\u001e\f\u001e\u0161\t\u001e\u0003\u001e"+
		"\u0163\b\u001e\u0001\u001e\u0001\u001e\u0001\u001f\u0001\u001f\u0001\u001f"+
		"\u0001 \u0001 \u0001 \u0001 \u0001!\u0001!\u0001!\u0001!\u0001\"\u0001"+
		"\"\u0001\"\u0001\"\u0001\"\u0005\"\u0177\b\"\n\"\f\"\u017a\t\"\u0003\""+
		"\u017c\b\"\u0001\"\u0001\"\u0001\"\u0001\"\u0001#\u0001#\u0001#\u0001"+
		"#\u0001#\u0001#\u0001#\u0005#\u0189\b#\n#\f#\u018c\t#\u0003#\u018e\b#"+
		"\u0001#\u0001#\u0001$\u0001$\u0001$\u0001$\u0001$\u0001$\u0000\u0001\n"+
		"%\u0000\u0002\u0004\u0006\b\n\f\u000e\u0010\u0012\u0014\u0016\u0018\u001a"+
		"\u001c\u001e \"$&(*,.02468:<>@BDFH\u0000\u0007\u0001\u0000\u0005\n\u0001"+
		"\u0000\u000b\u000f\u0002\u0000$$00\u0001\u0000\u0010\u0014\u0004\u0000"+
		"**..44<=\u0002\u0000\u001e\u001e>?\u0001\u00009;\u01ad\u0000M\u0001\u0000"+
		"\u0000\u0000\u0002X\u0001\u0000\u0000\u0000\u0004_\u0001\u0000\u0000\u0000"+
		"\u0006a\u0001\u0000\u0000\u0000\bd\u0001\u0000\u0000\u0000\nv\u0001\u0000"+
		"\u0000\u0000\f\u0086\u0001\u0000\u0000\u0000\u000e\u0091\u0001\u0000\u0000"+
		"\u0000\u0010\u0097\u0001\u0000\u0000\u0000\u0012\u0099\u0001\u0000\u0000"+
		"\u0000\u0014\u009b\u0001\u0000\u0000\u0000\u0016\u009d\u0001\u0000\u0000"+
		"\u0000\u0018\u009f\u0001\u0000\u0000\u0000\u001a\u00a1\u0001\u0000\u0000"+
		"\u0000\u001c\u00bc\u0001\u0000\u0000\u0000\u001e\u00be\u0001\u0000\u0000"+
		"\u0000 \u00c5\u0001\u0000\u0000\u0000\"\u00da\u0001\u0000\u0000\u0000"+
		"$\u00de\u0001\u0000\u0000\u0000&\u0108\u0001\u0000\u0000\u0000(\u010a"+
		"\u0001\u0000\u0000\u0000*\u0110\u0001\u0000\u0000\u0000,\u0127\u0001\u0000"+
		"\u0000\u0000.\u0132\u0001\u0000\u0000\u00000\u0134\u0001\u0000\u0000\u0000"+
		"2\u0136\u0001\u0000\u0000\u00004\u013b\u0001\u0000\u0000\u00006\u0142"+
		"\u0001\u0000\u0000\u00008\u0149\u0001\u0000\u0000\u0000:\u0151\u0001\u0000"+
		"\u0000\u0000<\u0159\u0001\u0000\u0000\u0000>\u0166\u0001\u0000\u0000\u0000"+
		"@\u0169\u0001\u0000\u0000\u0000B\u016d\u0001\u0000\u0000\u0000D\u0171"+
		"\u0001\u0000\u0000\u0000F\u0181\u0001\u0000\u0000\u0000H\u0191\u0001\u0000"+
		"\u0000\u0000JL\u0003\u0002\u0001\u0000KJ\u0001\u0000\u0000\u0000LO\u0001"+
		"\u0000\u0000\u0000MK\u0001\u0000\u0000\u0000MN\u0001\u0000\u0000\u0000"+
		"N\u0001\u0001\u0000\u0000\u0000OM\u0001\u0000\u0000\u0000PY\u0003\b\u0004"+
		"\u0000QY\u0003\n\u0005\u0000RY\u0003\u0004\u0002\u0000SY\u0003 \u0010"+
		"\u0000TY\u0003\u001c\u000e\u0000UY\u0003<\u001e\u0000VY\u0003>\u001f\u0000"+
		"WY\u0003B!\u0000XP\u0001\u0000\u0000\u0000XQ\u0001\u0000\u0000\u0000X"+
		"R\u0001\u0000\u0000\u0000XS\u0001\u0000\u0000\u0000XT\u0001\u0000\u0000"+
		"\u0000XU\u0001\u0000\u0000\u0000XV\u0001\u0000\u0000\u0000XW\u0001\u0000"+
		"\u0000\u0000Y\u0003\u0001\u0000\u0000\u0000Z`\u0003$\u0012\u0000[`\u0003"+
		"&\u0013\u0000\\`\u0003(\u0014\u0000]`\u00036\u001b\u0000^`\u0003@ \u0000"+
		"_Z\u0001\u0000\u0000\u0000_[\u0001\u0000\u0000\u0000_\\\u0001\u0000\u0000"+
		"\u0000_]\u0001\u0000\u0000\u0000_^\u0001\u0000\u0000\u0000`\u0005\u0001"+
		"\u0000\u0000\u0000ab\u0005!\u0000\u0000b\u0007\u0001\u0000\u0000\u0000"+
		"ce\u00057\u0000\u0000dc\u0001\u0000\u0000\u0000de\u0001\u0000\u0000\u0000"+
		"ef\u0001\u0000\u0000\u0000fg\u00038\u001c\u0000gh\u0005\u0001\u0000\u0000"+
		"hi\u0003:\u001d\u0000ij\u0003\u0006\u0003\u0000j\t\u0001\u0000\u0000\u0000"+
		"kl\u0006\u0005\uffff\uffff\u0000lw\u0003\u001a\r\u0000mw\u00034\u001a"+
		"\u0000no\u0005\u0002\u0000\u0000op\u0003\n\u0005\u0000pq\u0005\u0003\u0000"+
		"\u0000qw\u0001\u0000\u0000\u0000rw\u0003\u001c\u000e\u0000sw\u0003*\u0015"+
		"\u0000tw\u0003.\u0017\u0000uw\u0003D\"\u0000vk\u0001\u0000\u0000\u0000"+
		"vm\u0001\u0000\u0000\u0000vn\u0001\u0000\u0000\u0000vr\u0001\u0000\u0000"+
		"\u0000vs\u0001\u0000\u0000\u0000vt\u0001\u0000\u0000\u0000vu\u0001\u0000"+
		"\u0000\u0000w~\u0001\u0000\u0000\u0000xy\n\u0005\u0000\u0000yz\u0003\u0010"+
		"\b\u0000z{\u0003\n\u0005\u0006{}\u0001\u0000\u0000\u0000|x\u0001\u0000"+
		"\u0000\u0000}\u0080\u0001\u0000\u0000\u0000~|\u0001\u0000\u0000\u0000"+
		"~\u007f\u0001\u0000\u0000\u0000\u007f\u000b\u0001\u0000\u0000\u0000\u0080"+
		"~\u0001\u0000\u0000\u0000\u0081\u0087\u0003\u000e\u0007\u0000\u0082\u0083"+
		"\u0005/\u0000\u0000\u0083\u0087\u0003\f\u0006\u0000\u0084\u0085\u0005"+
		"\u0004\u0000\u0000\u0085\u0087\u0003\f\u0006\u0000\u0086\u0081\u0001\u0000"+
		"\u0000\u0000\u0086\u0082\u0001\u0000\u0000\u0000\u0086\u0084\u0001\u0000"+
		"\u0000\u0000\u0087\r\u0001\u0000\u0000\u0000\u0088\u0092\u0003\u001a\r"+
		"\u0000\u0089\u0092\u00034\u001a\u0000\u008a\u008b\u0005\u0002\u0000\u0000"+
		"\u008b\u008c\u0003\n\u0005\u0000\u008c\u008d\u0005\u0003\u0000\u0000\u008d"+
		"\u0092\u0001\u0000\u0000\u0000\u008e\u0092\u0003\u001c\u000e\u0000\u008f"+
		"\u0092\u0003*\u0015\u0000\u0090\u0092\u0003.\u0017\u0000\u0091\u0088\u0001"+
		"\u0000\u0000\u0000\u0091\u0089\u0001\u0000\u0000\u0000\u0091\u008a\u0001"+
		"\u0000\u0000\u0000\u0091\u008e\u0001\u0000\u0000\u0000\u0091\u008f\u0001"+
		"\u0000\u0000\u0000\u0091\u0090\u0001\u0000\u0000\u0000\u0092\u000f\u0001"+
		"\u0000\u0000\u0000\u0093\u0098\u0003\u0012\t\u0000\u0094\u0098\u0003\u0014"+
		"\n\u0000\u0095\u0098\u0003\u0016\u000b\u0000\u0096\u0098\u0003\u0018\f"+
		"\u0000\u0097\u0093\u0001\u0000\u0000\u0000\u0097\u0094\u0001\u0000\u0000"+
		"\u0000\u0097\u0095\u0001\u0000\u0000\u0000\u0097\u0096\u0001\u0000\u0000"+
		"\u0000\u0098\u0011\u0001\u0000\u0000\u0000\u0099\u009a\u0007\u0000\u0000"+
		"\u0000\u009a\u0013\u0001\u0000\u0000\u0000\u009b\u009c\u0007\u0001\u0000"+
		"\u0000\u009c\u0015\u0001\u0000\u0000\u0000\u009d\u009e\u0007\u0002\u0000"+
		"\u0000\u009e\u0017\u0001\u0000\u0000\u0000\u009f\u00a0\u0007\u0003\u0000"+
		"\u0000\u00a0\u0019\u0001\u0000\u0000\u0000\u00a1\u00a2\u0007\u0004\u0000"+
		"\u0000\u00a2\u001b\u0001\u0000\u0000\u0000\u00a3\u00aa\u00034\u001a\u0000"+
		"\u00a4\u00aa\u0003.\u0017\u0000\u00a5\u00a6\u0005\u0002\u0000\u0000\u00a6"+
		"\u00a7\u0003\n\u0005\u0000\u00a7\u00a8\u0005\u0003\u0000\u0000\u00a8\u00aa"+
		"\u0001\u0000\u0000\u0000\u00a9\u00a3\u0001\u0000\u0000\u0000\u00a9\u00a4"+
		"\u0001\u0000\u0000\u0000\u00a9\u00a5\u0001\u0000\u0000\u0000\u00aa\u00ad"+
		"\u0001\u0000\u0000\u0000\u00ab\u00ac\u0005\u0015\u0000\u0000\u00ac\u00ae"+
		"\u00034\u001a\u0000\u00ad\u00ab\u0001\u0000\u0000\u0000\u00ad\u00ae\u0001"+
		"\u0000\u0000\u0000\u00ae\u00af\u0001\u0000\u0000\u0000\u00af\u00b1\u0005"+
		"\u0002\u0000\u0000\u00b0\u00b2\u0003:\u001d\u0000\u00b1\u00b0\u0001\u0000"+
		"\u0000\u0000\u00b1\u00b2\u0001\u0000\u0000\u0000\u00b2\u00b3\u0001\u0000"+
		"\u0000\u0000\u00b3\u00b4\u0005\u0003\u0000\u0000\u00b4\u00bd\u0001\u0000"+
		"\u0000\u0000\u00b5\u00bd\u0003\u001e\u000f\u0000\u00b6\u00b7\u0005#\u0000"+
		"\u0000\u00b7\u00b8\u0005\u0002\u0000\u0000\u00b8\u00b9\u0003:\u001d\u0000"+
		"\u00b9\u00ba\u0005\u0003\u0000\u0000\u00ba\u00bd\u0001\u0000\u0000\u0000"+
		"\u00bb\u00bd\u0003F#\u0000\u00bc\u00a9\u0001\u0000\u0000\u0000\u00bc\u00b5"+
		"\u0001\u0000\u0000\u0000\u00bc\u00b6\u0001\u0000\u0000\u0000\u00bc\u00bb"+
		"\u0001\u0000\u0000\u0000\u00bd\u001d\u0001\u0000\u0000\u0000\u00be\u00bf"+
		"\u0005\u0016\u0000\u0000\u00bf\u00c0\u0005\u0002\u0000\u0000\u00c0\u00c1"+
		"\u00034\u001a\u0000\u00c1\u00c2\u0005\u0017\u0000\u0000\u00c2\u00c3\u0003"+
		"\n\u0005\u0000\u00c3\u00c4\u0005\u0003\u0000\u0000\u00c4\u001f\u0001\u0000"+
		"\u0000\u0000\u00c5\u00c6\u00058\u0000\u0000\u00c6\u00c7\u00034\u001a\u0000"+
		"\u00c7\u00d0\u0005\u0002\u0000\u0000\u00c8\u00cd\u00034\u001a\u0000\u00c9"+
		"\u00ca\u0005\u0017\u0000\u0000\u00ca\u00cc\u00034\u001a\u0000\u00cb\u00c9"+
		"\u0001\u0000\u0000\u0000\u00cc\u00cf\u0001\u0000\u0000\u0000\u00cd\u00cb"+
		"\u0001\u0000\u0000\u0000\u00cd\u00ce\u0001\u0000\u0000\u0000\u00ce\u00d1"+
		"\u0001\u0000\u0000\u0000\u00cf\u00cd\u0001\u0000\u0000\u0000\u00d0\u00c8"+
		"\u0001\u0000\u0000\u0000\u00d0\u00d1\u0001\u0000\u0000\u0000\u00d1\u00d2"+
		"\u0001\u0000\u0000\u0000\u00d2\u00d3\u0005\u0003\u0000\u0000\u00d3\u00d4"+
		"\u0003\"\u0011\u0000\u00d4\u00d5\u0005)\u0000\u0000\u00d5!\u0001\u0000"+
		"\u0000\u0000\u00d6\u00d8\u0003\u0002\u0001\u0000\u00d7\u00d9\u0003\u0006"+
		"\u0003\u0000\u00d8\u00d7\u0001\u0000\u0000\u0000\u00d8\u00d9\u0001\u0000"+
		"\u0000\u0000\u00d9\u00db\u0001\u0000\u0000\u0000\u00da\u00d6\u0001\u0000"+
		"\u0000\u0000\u00db\u00dc\u0001\u0000\u0000\u0000\u00dc\u00da\u0001\u0000"+
		"\u0000\u0000\u00dc\u00dd\u0001\u0000\u0000\u0000\u00dd#\u0001\u0000\u0000"+
		"\u0000\u00de\u00df\u0005-\u0000\u0000\u00df\u00e0\u0003\n\u0005\u0000"+
		"\u00e0\u00e1\u00053\u0000\u0000\u00e1\u00e9\u0003\"\u0011\u0000\u00e2"+
		"\u00e3\u0005(\u0000\u0000\u00e3\u00e4\u0003\n\u0005\u0000\u00e4\u00e5"+
		"\u00053\u0000\u0000\u00e5\u00e6\u0003\"\u0011\u0000\u00e6\u00e8\u0001"+
		"\u0000\u0000\u0000\u00e7\u00e2\u0001\u0000\u0000\u0000\u00e8\u00eb\u0001"+
		"\u0000\u0000\u0000\u00e9\u00e7\u0001\u0000\u0000\u0000\u00e9\u00ea\u0001"+
		"\u0000\u0000\u0000\u00ea\u00ee\u0001\u0000\u0000\u0000\u00eb\u00e9\u0001"+
		"\u0000\u0000\u0000\u00ec\u00ed\u0005\'\u0000\u0000\u00ed\u00ef\u0003\""+
		"\u0011\u0000\u00ee\u00ec\u0001\u0000\u0000\u0000\u00ee\u00ef\u0001\u0000"+
		"\u0000\u0000\u00ef\u00f0\u0001\u0000\u0000\u0000\u00f0\u00f1\u0005)\u0000"+
		"\u0000\u00f1%\u0001\u0000\u0000\u0000\u00f2\u00f3\u0005+\u0000\u0000\u00f3"+
		"\u00f4\u00034\u001a\u0000\u00f4\u00f5\u0005\"\u0000\u0000\u00f5\u00f6"+
		"\u0003\n\u0005\u0000\u00f6\u00f7\u0005&\u0000\u0000\u00f7\u00f8\u0003"+
		"\"\u0011\u0000\u00f8\u00f9\u0005)\u0000\u0000\u00f9\u0109\u0001\u0000"+
		"\u0000\u0000\u00fa\u00fb\u0005+\u0000\u0000\u00fb\u00fc\u00034\u001a\u0000"+
		"\u00fc\u00fd\u0005\u0001\u0000\u0000\u00fd\u00fe\u0003\n\u0005\u0000\u00fe"+
		"\u00ff\u0005\u0017\u0000\u0000\u00ff\u0102\u0003\n\u0005\u0000\u0100\u0101"+
		"\u0005\u0017\u0000\u0000\u0101\u0103\u0003\n\u0005\u0000\u0102\u0100\u0001"+
		"\u0000\u0000\u0000\u0102\u0103\u0001\u0000\u0000\u0000\u0103\u0104\u0001"+
		"\u0000\u0000\u0000\u0104\u0105\u0005&\u0000\u0000\u0105\u0106\u0003\""+
		"\u0011\u0000\u0106\u0107\u0005)\u0000\u0000\u0107\u0109\u0001\u0000\u0000"+
		"\u0000\u0108\u00f2\u0001\u0000\u0000\u0000\u0108\u00fa\u0001\u0000\u0000"+
		"\u0000\u0109\'\u0001\u0000\u0000\u0000\u010a\u010b\u00056\u0000\u0000"+
		"\u010b\u010c\u0003\n\u0005\u0000\u010c\u010d\u0005&\u0000\u0000\u010d"+
		"\u010e\u0003\"\u0011\u0000\u010e\u010f\u0005)\u0000\u0000\u010f)\u0001"+
		"\u0000\u0000\u0000\u0110\u011e\u0005\u0018\u0000\u0000\u0111\u0116\u0003"+
		",\u0016\u0000\u0112\u0113\u0005\u0017\u0000\u0000\u0113\u0115\u0003,\u0016"+
		"\u0000\u0114\u0112\u0001\u0000\u0000\u0000\u0115\u0118\u0001\u0000\u0000"+
		"\u0000\u0116\u0114\u0001\u0000\u0000\u0000\u0116\u0117\u0001\u0000\u0000"+
		"\u0000\u0117\u011b\u0001\u0000\u0000\u0000\u0118\u0116\u0001\u0000\u0000"+
		"\u0000\u0119\u011a\u0005\u0017\u0000\u0000\u011a\u011c\u0003H$\u0000\u011b"+
		"\u0119\u0001\u0000\u0000\u0000\u011b\u011c\u0001\u0000\u0000\u0000\u011c"+
		"\u011f\u0001\u0000\u0000\u0000\u011d\u011f\u0003H$\u0000\u011e\u0111\u0001"+
		"\u0000\u0000\u0000\u011e\u011d\u0001\u0000\u0000\u0000\u011f\u0120\u0001"+
		"\u0000\u0000\u0000\u0120\u0121\u0005\u0019\u0000\u0000\u0121+\u0001\u0000"+
		"\u0000\u0000\u0122\u0123\u00034\u001a\u0000\u0123\u0124\u0005\u0001\u0000"+
		"\u0000\u0124\u0125\u0003\n\u0005\u0000\u0125\u0128\u0001\u0000\u0000\u0000"+
		"\u0126\u0128\u0003\n\u0005\u0000\u0127\u0122\u0001\u0000\u0000\u0000\u0127"+
		"\u0126\u0001\u0000\u0000\u0000\u0128-\u0001\u0000\u0000\u0000\u0129\u012a"+
		"\u00034\u001a\u0000\u012a\u012b\u0005\u001a\u0000\u0000\u012b\u012c\u0003"+
		"\n\u0005\u0000\u012c\u012d\u0005\u001b\u0000\u0000\u012d\u0133\u0001\u0000"+
		"\u0000\u0000\u012e\u012f\u00034\u001a\u0000\u012f\u0130\u0005\u001c\u0000"+
		"\u0000\u0130\u0131\u00034\u001a\u0000\u0131\u0133\u0001\u0000\u0000\u0000"+
		"\u0132\u0129\u0001\u0000\u0000\u0000\u0132\u012e\u0001\u0000\u0000\u0000"+
		"\u0133/\u0001\u0000\u0000\u0000\u0134\u0135\u0005\u001d\u0000\u0000\u0135"+
		"1\u0001\u0000\u0000\u0000\u0136\u0137\u0005#\u0000\u0000\u0137\u0138\u0005"+
		"\u0002\u0000\u0000\u0138\u0139\u0003\n\u0005\u0000\u0139\u013a\u0005\u0003"+
		"\u0000\u0000\u013a3\u0001\u0000\u0000\u0000\u013b\u013f\u0005>\u0000\u0000"+
		"\u013c\u013e\u0007\u0005\u0000\u0000\u013d\u013c\u0001\u0000\u0000\u0000"+
		"\u013e\u0141\u0001\u0000\u0000\u0000\u013f\u013d\u0001\u0000\u0000\u0000"+
		"\u013f\u0140\u0001\u0000\u0000\u0000\u01405\u0001\u0000\u0000\u0000\u0141"+
		"\u013f\u0001\u0000\u0000\u0000\u0142\u0143\u00051\u0000\u0000\u0143\u0144"+
		"\u0003\"\u0011\u0000\u0144\u0145\u00055\u0000\u0000\u0145\u0147\u0003"+
		"\n\u0005\u0000\u0146\u0148\u0003\u0006\u0003\u0000\u0147\u0146\u0001\u0000"+
		"\u0000\u0000\u0147\u0148\u0001\u0000\u0000\u0000\u01487\u0001\u0000\u0000"+
		"\u0000\u0149\u014e\u00034\u001a\u0000\u014a\u014b\u0005\u0017\u0000\u0000"+
		"\u014b\u014d\u00034\u001a\u0000\u014c\u014a\u0001\u0000\u0000\u0000\u014d"+
		"\u0150\u0001\u0000\u0000\u0000\u014e\u014c\u0001\u0000\u0000\u0000\u014e"+
		"\u014f\u0001\u0000\u0000\u0000\u014f9\u0001\u0000\u0000\u0000\u0150\u014e"+
		"\u0001\u0000\u0000\u0000\u0151\u0156\u0003\n\u0005\u0000\u0152\u0153\u0005"+
		"\u0017\u0000\u0000\u0153\u0155\u0003\n\u0005\u0000\u0154\u0152\u0001\u0000"+
		"\u0000\u0000\u0155\u0158\u0001\u0000\u0000\u0000\u0156\u0154\u0001\u0000"+
		"\u0000\u0000\u0156\u0157\u0001\u0000\u0000\u0000\u0157;\u0001\u0000\u0000"+
		"\u0000\u0158\u0156\u0001\u0000\u0000\u0000\u0159\u0162\u00052\u0000\u0000"+
		"\u015a\u015f\u0003\n\u0005\u0000\u015b\u015c\u0005\u0017\u0000\u0000\u015c"+
		"\u015e\u0003\n\u0005\u0000\u015d\u015b\u0001\u0000\u0000\u0000\u015e\u0161"+
		"\u0001\u0000\u0000\u0000\u015f\u015d\u0001\u0000\u0000\u0000\u015f\u0160"+
		"\u0001\u0000\u0000\u0000\u0160\u0163\u0001\u0000\u0000\u0000\u0161\u015f"+
		"\u0001\u0000\u0000\u0000\u0162\u015a\u0001\u0000\u0000\u0000\u0162\u0163"+
		"\u0001\u0000\u0000\u0000\u0163\u0164\u0001\u0000\u0000\u0000\u0164\u0165"+
		"\u0003\u0006\u0003\u0000\u0165=\u0001\u0000\u0000\u0000\u0166\u0167\u0005"+
		"%\u0000\u0000\u0167\u0168\u0003\u0006\u0003\u0000\u0168?\u0001\u0000\u0000"+
		"\u0000\u0169\u016a\u0005,\u0000\u0000\u016a\u016b\u00034\u001a\u0000\u016b"+
		"\u016c\u0003\u0006\u0003\u0000\u016cA\u0001\u0000\u0000\u0000\u016d\u016e"+
		"\u0005\u001f\u0000\u0000\u016e\u016f\u00034\u001a\u0000\u016f\u0170\u0005"+
		"\u001f\u0000\u0000\u0170C\u0001\u0000\u0000\u0000\u0171\u0172\u00058\u0000"+
		"\u0000\u0172\u017b\u0005\u0002\u0000\u0000\u0173\u0178\u00034\u001a\u0000"+
		"\u0174\u0175\u0005\u0017\u0000\u0000\u0175\u0177\u00034\u001a\u0000\u0176"+
		"\u0174\u0001\u0000\u0000\u0000\u0177\u017a\u0001\u0000\u0000\u0000\u0178"+
		"\u0176\u0001\u0000\u0000\u0000\u0178\u0179\u0001\u0000\u0000\u0000\u0179"+
		"\u017c\u0001\u0000\u0000\u0000\u017a\u0178\u0001\u0000\u0000\u0000\u017b"+
		"\u0173\u0001\u0000\u0000\u0000\u017b\u017c\u0001\u0000\u0000\u0000\u017c"+
		"\u017d\u0001\u0000\u0000\u0000\u017d\u017e\u0005\u0003\u0000\u0000\u017e"+
		"\u017f\u0003\"\u0011\u0000\u017f\u0180\u0005)\u0000\u0000\u0180E\u0001"+
		"\u0000\u0000\u0000\u0181\u0182\u00034\u001a\u0000\u0182\u0183\u0005\u0015"+
		"\u0000\u0000\u0183\u0184\u00034\u001a\u0000\u0184\u018d\u0005\u0002\u0000"+
		"\u0000\u0185\u018a\u0003\n\u0005\u0000\u0186\u0187\u0005\u0017\u0000\u0000"+
		"\u0187\u0189\u0003\n\u0005\u0000\u0188\u0186\u0001\u0000\u0000\u0000\u0189"+
		"\u018c\u0001\u0000\u0000\u0000\u018a\u0188\u0001\u0000\u0000\u0000\u018a"+
		"\u018b\u0001\u0000\u0000\u0000\u018b\u018e\u0001\u0000\u0000\u0000\u018c"+
		"\u018a\u0001\u0000\u0000\u0000\u018d\u0185\u0001\u0000\u0000\u0000\u018d"+
		"\u018e\u0001\u0000\u0000\u0000\u018e\u018f\u0001\u0000\u0000\u0000\u018f"+
		"\u0190\u0005\u0003\u0000\u0000\u0190G\u0001\u0000\u0000\u0000\u0191\u0192"+
		"\u0005 \u0000\u0000\u0192\u0193\u0007\u0006\u0000\u0000\u0193\u0194\u0005"+
		"\u0001\u0000\u0000\u0194\u0195\u0003\n\u0005\u0000\u0195I\u0001\u0000"+
		"\u0000\u0000$MX_dv~\u0086\u0091\u0097\u00a9\u00ad\u00b1\u00bc\u00cd\u00d0"+
		"\u00d8\u00dc\u00e9\u00ee\u0102\u0108\u0116\u011b\u011e\u0127\u0132\u013f"+
		"\u0147\u014e\u0156\u015f\u0162\u0178\u017b\u018a\u018d";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}