// Generated from d:/Programs/Programmed/fluxlang/antlr4/FLUXASSEMBLY/src/lua_grammar_antlr4.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.tree.ParseTreeListener;

/**
 * This interface defines a complete listener for a parse tree produced by
 * {@link lua_grammar_antlr4Parser}.
 */
public interface lua_grammar_antlr4Listener extends ParseTreeListener {
	/**
	 * Enter a parse tree produced by {@link lua_grammar_antlr4Parser#program}.
	 * @param ctx the parse tree
	 */
	void enterProgram(lua_grammar_antlr4Parser.ProgramContext ctx);
	/**
	 * Exit a parse tree produced by {@link lua_grammar_antlr4Parser#program}.
	 * @param ctx the parse tree
	 */
	void exitProgram(lua_grammar_antlr4Parser.ProgramContext ctx);
	/**
	 * Enter a parse tree produced by {@link lua_grammar_antlr4Parser#statement}.
	 * @param ctx the parse tree
	 */
	void enterStatement(lua_grammar_antlr4Parser.StatementContext ctx);
	/**
	 * Exit a parse tree produced by {@link lua_grammar_antlr4Parser#statement}.
	 * @param ctx the parse tree
	 */
	void exitStatement(lua_grammar_antlr4Parser.StatementContext ctx);
	/**
	 * Enter a parse tree produced by {@link lua_grammar_antlr4Parser#assignStatement}.
	 * @param ctx the parse tree
	 */
	void enterAssignStatement(lua_grammar_antlr4Parser.AssignStatementContext ctx);
	/**
	 * Exit a parse tree produced by {@link lua_grammar_antlr4Parser#assignStatement}.
	 * @param ctx the parse tree
	 */
	void exitAssignStatement(lua_grammar_antlr4Parser.AssignStatementContext ctx);
	/**
	 * Enter a parse tree produced by {@link lua_grammar_antlr4Parser#expression}.
	 * @param ctx the parse tree
	 */
	void enterExpression(lua_grammar_antlr4Parser.ExpressionContext ctx);
	/**
	 * Exit a parse tree produced by {@link lua_grammar_antlr4Parser#expression}.
	 * @param ctx the parse tree
	 */
	void exitExpression(lua_grammar_antlr4Parser.ExpressionContext ctx);
	/**
	 * Enter a parse tree produced by {@link lua_grammar_antlr4Parser#primaryExpression}.
	 * @param ctx the parse tree
	 */
	void enterPrimaryExpression(lua_grammar_antlr4Parser.PrimaryExpressionContext ctx);
	/**
	 * Exit a parse tree produced by {@link lua_grammar_antlr4Parser#primaryExpression}.
	 * @param ctx the parse tree
	 */
	void exitPrimaryExpression(lua_grammar_antlr4Parser.PrimaryExpressionContext ctx);
	/**
	 * Enter a parse tree produced by {@link lua_grammar_antlr4Parser#literal}.
	 * @param ctx the parse tree
	 */
	void enterLiteral(lua_grammar_antlr4Parser.LiteralContext ctx);
	/**
	 * Exit a parse tree produced by {@link lua_grammar_antlr4Parser#literal}.
	 * @param ctx the parse tree
	 */
	void exitLiteral(lua_grammar_antlr4Parser.LiteralContext ctx);
	/**
	 * Enter a parse tree produced by {@link lua_grammar_antlr4Parser#variable}.
	 * @param ctx the parse tree
	 */
	void enterVariable(lua_grammar_antlr4Parser.VariableContext ctx);
	/**
	 * Exit a parse tree produced by {@link lua_grammar_antlr4Parser#variable}.
	 * @param ctx the parse tree
	 */
	void exitVariable(lua_grammar_antlr4Parser.VariableContext ctx);
	/**
	 * Enter a parse tree produced by {@link lua_grammar_antlr4Parser#functionCall}.
	 * @param ctx the parse tree
	 */
	void enterFunctionCall(lua_grammar_antlr4Parser.FunctionCallContext ctx);
	/**
	 * Exit a parse tree produced by {@link lua_grammar_antlr4Parser#functionCall}.
	 * @param ctx the parse tree
	 */
	void exitFunctionCall(lua_grammar_antlr4Parser.FunctionCallContext ctx);
	/**
	 * Enter a parse tree produced by {@link lua_grammar_antlr4Parser#tableConstructor}.
	 * @param ctx the parse tree
	 */
	void enterTableConstructor(lua_grammar_antlr4Parser.TableConstructorContext ctx);
	/**
	 * Exit a parse tree produced by {@link lua_grammar_antlr4Parser#tableConstructor}.
	 * @param ctx the parse tree
	 */
	void exitTableConstructor(lua_grammar_antlr4Parser.TableConstructorContext ctx);
	/**
	 * Enter a parse tree produced by {@link lua_grammar_antlr4Parser#metatable}.
	 * @param ctx the parse tree
	 */
	void enterMetatable(lua_grammar_antlr4Parser.MetatableContext ctx);
	/**
	 * Exit a parse tree produced by {@link lua_grammar_antlr4Parser#metatable}.
	 * @param ctx the parse tree
	 */
	void exitMetatable(lua_grammar_antlr4Parser.MetatableContext ctx);
	/**
	 * Enter a parse tree produced by {@link lua_grammar_antlr4Parser#metamethods}.
	 * @param ctx the parse tree
	 */
	void enterMetamethods(lua_grammar_antlr4Parser.MetamethodsContext ctx);
	/**
	 * Exit a parse tree produced by {@link lua_grammar_antlr4Parser#metamethods}.
	 * @param ctx the parse tree
	 */
	void exitMetamethods(lua_grammar_antlr4Parser.MetamethodsContext ctx);
	/**
	 * Enter a parse tree produced by {@link lua_grammar_antlr4Parser#labelStatement}.
	 * @param ctx the parse tree
	 */
	void enterLabelStatement(lua_grammar_antlr4Parser.LabelStatementContext ctx);
	/**
	 * Exit a parse tree produced by {@link lua_grammar_antlr4Parser#labelStatement}.
	 * @param ctx the parse tree
	 */
	void exitLabelStatement(lua_grammar_antlr4Parser.LabelStatementContext ctx);
	/**
	 * Enter a parse tree produced by {@link lua_grammar_antlr4Parser#metamethod}.
	 * @param ctx the parse tree
	 */
	void enterMetamethod(lua_grammar_antlr4Parser.MetamethodContext ctx);
	/**
	 * Exit a parse tree produced by {@link lua_grammar_antlr4Parser#metamethod}.
	 * @param ctx the parse tree
	 */
	void exitMetamethod(lua_grammar_antlr4Parser.MetamethodContext ctx);
	/**
	 * Enter a parse tree produced by {@link lua_grammar_antlr4Parser#field}.
	 * @param ctx the parse tree
	 */
	void enterField(lua_grammar_antlr4Parser.FieldContext ctx);
	/**
	 * Exit a parse tree produced by {@link lua_grammar_antlr4Parser#field}.
	 * @param ctx the parse tree
	 */
	void exitField(lua_grammar_antlr4Parser.FieldContext ctx);
	/**
	 * Enter a parse tree produced by {@link lua_grammar_antlr4Parser#binaryOperation}.
	 * @param ctx the parse tree
	 */
	void enterBinaryOperation(lua_grammar_antlr4Parser.BinaryOperationContext ctx);
	/**
	 * Exit a parse tree produced by {@link lua_grammar_antlr4Parser#binaryOperation}.
	 * @param ctx the parse tree
	 */
	void exitBinaryOperation(lua_grammar_antlr4Parser.BinaryOperationContext ctx);
	/**
	 * Enter a parse tree produced by {@link lua_grammar_antlr4Parser#unaryOperation}.
	 * @param ctx the parse tree
	 */
	void enterUnaryOperation(lua_grammar_antlr4Parser.UnaryOperationContext ctx);
	/**
	 * Exit a parse tree produced by {@link lua_grammar_antlr4Parser#unaryOperation}.
	 * @param ctx the parse tree
	 */
	void exitUnaryOperation(lua_grammar_antlr4Parser.UnaryOperationContext ctx);
	/**
	 * Enter a parse tree produced by {@link lua_grammar_antlr4Parser#controlFlowStatement}.
	 * @param ctx the parse tree
	 */
	void enterControlFlowStatement(lua_grammar_antlr4Parser.ControlFlowStatementContext ctx);
	/**
	 * Exit a parse tree produced by {@link lua_grammar_antlr4Parser#controlFlowStatement}.
	 * @param ctx the parse tree
	 */
	void exitControlFlowStatement(lua_grammar_antlr4Parser.ControlFlowStatementContext ctx);
	/**
	 * Enter a parse tree produced by {@link lua_grammar_antlr4Parser#ifStatement}.
	 * @param ctx the parse tree
	 */
	void enterIfStatement(lua_grammar_antlr4Parser.IfStatementContext ctx);
	/**
	 * Exit a parse tree produced by {@link lua_grammar_antlr4Parser#ifStatement}.
	 * @param ctx the parse tree
	 */
	void exitIfStatement(lua_grammar_antlr4Parser.IfStatementContext ctx);
	/**
	 * Enter a parse tree produced by {@link lua_grammar_antlr4Parser#whileStatement}.
	 * @param ctx the parse tree
	 */
	void enterWhileStatement(lua_grammar_antlr4Parser.WhileStatementContext ctx);
	/**
	 * Exit a parse tree produced by {@link lua_grammar_antlr4Parser#whileStatement}.
	 * @param ctx the parse tree
	 */
	void exitWhileStatement(lua_grammar_antlr4Parser.WhileStatementContext ctx);
	/**
	 * Enter a parse tree produced by {@link lua_grammar_antlr4Parser#repeatStatement}.
	 * @param ctx the parse tree
	 */
	void enterRepeatStatement(lua_grammar_antlr4Parser.RepeatStatementContext ctx);
	/**
	 * Exit a parse tree produced by {@link lua_grammar_antlr4Parser#repeatStatement}.
	 * @param ctx the parse tree
	 */
	void exitRepeatStatement(lua_grammar_antlr4Parser.RepeatStatementContext ctx);
	/**
	 * Enter a parse tree produced by {@link lua_grammar_antlr4Parser#forStatement}.
	 * @param ctx the parse tree
	 */
	void enterForStatement(lua_grammar_antlr4Parser.ForStatementContext ctx);
	/**
	 * Exit a parse tree produced by {@link lua_grammar_antlr4Parser#forStatement}.
	 * @param ctx the parse tree
	 */
	void exitForStatement(lua_grammar_antlr4Parser.ForStatementContext ctx);
	/**
	 * Enter a parse tree produced by {@link lua_grammar_antlr4Parser#breakStatement}.
	 * @param ctx the parse tree
	 */
	void enterBreakStatement(lua_grammar_antlr4Parser.BreakStatementContext ctx);
	/**
	 * Exit a parse tree produced by {@link lua_grammar_antlr4Parser#breakStatement}.
	 * @param ctx the parse tree
	 */
	void exitBreakStatement(lua_grammar_antlr4Parser.BreakStatementContext ctx);
	/**
	 * Enter a parse tree produced by {@link lua_grammar_antlr4Parser#gotoStatement}.
	 * @param ctx the parse tree
	 */
	void enterGotoStatement(lua_grammar_antlr4Parser.GotoStatementContext ctx);
	/**
	 * Exit a parse tree produced by {@link lua_grammar_antlr4Parser#gotoStatement}.
	 * @param ctx the parse tree
	 */
	void exitGotoStatement(lua_grammar_antlr4Parser.GotoStatementContext ctx);
	/**
	 * Enter a parse tree produced by {@link lua_grammar_antlr4Parser#coroutineStatement}.
	 * @param ctx the parse tree
	 */
	void enterCoroutineStatement(lua_grammar_antlr4Parser.CoroutineStatementContext ctx);
	/**
	 * Exit a parse tree produced by {@link lua_grammar_antlr4Parser#coroutineStatement}.
	 * @param ctx the parse tree
	 */
	void exitCoroutineStatement(lua_grammar_antlr4Parser.CoroutineStatementContext ctx);
	/**
	 * Enter a parse tree produced by {@link lua_grammar_antlr4Parser#block}.
	 * @param ctx the parse tree
	 */
	void enterBlock(lua_grammar_antlr4Parser.BlockContext ctx);
	/**
	 * Exit a parse tree produced by {@link lua_grammar_antlr4Parser#block}.
	 * @param ctx the parse tree
	 */
	void exitBlock(lua_grammar_antlr4Parser.BlockContext ctx);
	/**
	 * Enter a parse tree produced by {@link lua_grammar_antlr4Parser#localDeclaration}.
	 * @param ctx the parse tree
	 */
	void enterLocalDeclaration(lua_grammar_antlr4Parser.LocalDeclarationContext ctx);
	/**
	 * Exit a parse tree produced by {@link lua_grammar_antlr4Parser#localDeclaration}.
	 * @param ctx the parse tree
	 */
	void exitLocalDeclaration(lua_grammar_antlr4Parser.LocalDeclarationContext ctx);
	/**
	 * Enter a parse tree produced by {@link lua_grammar_antlr4Parser#functionDeclaration}.
	 * @param ctx the parse tree
	 */
	void enterFunctionDeclaration(lua_grammar_antlr4Parser.FunctionDeclarationContext ctx);
	/**
	 * Exit a parse tree produced by {@link lua_grammar_antlr4Parser#functionDeclaration}.
	 * @param ctx the parse tree
	 */
	void exitFunctionDeclaration(lua_grammar_antlr4Parser.FunctionDeclarationContext ctx);
	/**
	 * Enter a parse tree produced by {@link lua_grammar_antlr4Parser#returnStatement}.
	 * @param ctx the parse tree
	 */
	void enterReturnStatement(lua_grammar_antlr4Parser.ReturnStatementContext ctx);
	/**
	 * Exit a parse tree produced by {@link lua_grammar_antlr4Parser#returnStatement}.
	 * @param ctx the parse tree
	 */
	void exitReturnStatement(lua_grammar_antlr4Parser.ReturnStatementContext ctx);
	/**
	 * Enter a parse tree produced by {@link lua_grammar_antlr4Parser#operatorGroup}.
	 * @param ctx the parse tree
	 */
	void enterOperatorGroup(lua_grammar_antlr4Parser.OperatorGroupContext ctx);
	/**
	 * Exit a parse tree produced by {@link lua_grammar_antlr4Parser#operatorGroup}.
	 * @param ctx the parse tree
	 */
	void exitOperatorGroup(lua_grammar_antlr4Parser.OperatorGroupContext ctx);
	/**
	 * Enter a parse tree produced by {@link lua_grammar_antlr4Parser#logicalOp}.
	 * @param ctx the parse tree
	 */
	void enterLogicalOp(lua_grammar_antlr4Parser.LogicalOpContext ctx);
	/**
	 * Exit a parse tree produced by {@link lua_grammar_antlr4Parser#logicalOp}.
	 * @param ctx the parse tree
	 */
	void exitLogicalOp(lua_grammar_antlr4Parser.LogicalOpContext ctx);
	/**
	 * Enter a parse tree produced by {@link lua_grammar_antlr4Parser#comparisonOp}.
	 * @param ctx the parse tree
	 */
	void enterComparisonOp(lua_grammar_antlr4Parser.ComparisonOpContext ctx);
	/**
	 * Exit a parse tree produced by {@link lua_grammar_antlr4Parser#comparisonOp}.
	 * @param ctx the parse tree
	 */
	void exitComparisonOp(lua_grammar_antlr4Parser.ComparisonOpContext ctx);
	/**
	 * Enter a parse tree produced by {@link lua_grammar_antlr4Parser#arithOp}.
	 * @param ctx the parse tree
	 */
	void enterArithOp(lua_grammar_antlr4Parser.ArithOpContext ctx);
	/**
	 * Exit a parse tree produced by {@link lua_grammar_antlr4Parser#arithOp}.
	 * @param ctx the parse tree
	 */
	void exitArithOp(lua_grammar_antlr4Parser.ArithOpContext ctx);
	/**
	 * Enter a parse tree produced by {@link lua_grammar_antlr4Parser#bitwiseOp}.
	 * @param ctx the parse tree
	 */
	void enterBitwiseOp(lua_grammar_antlr4Parser.BitwiseOpContext ctx);
	/**
	 * Exit a parse tree produced by {@link lua_grammar_antlr4Parser#bitwiseOp}.
	 * @param ctx the parse tree
	 */
	void exitBitwiseOp(lua_grammar_antlr4Parser.BitwiseOpContext ctx);
	/**
	 * Enter a parse tree produced by {@link lua_grammar_antlr4Parser#assignOp}.
	 * @param ctx the parse tree
	 */
	void enterAssignOp(lua_grammar_antlr4Parser.AssignOpContext ctx);
	/**
	 * Exit a parse tree produced by {@link lua_grammar_antlr4Parser#assignOp}.
	 * @param ctx the parse tree
	 */
	void exitAssignOp(lua_grammar_antlr4Parser.AssignOpContext ctx);
	/**
	 * Enter a parse tree produced by {@link lua_grammar_antlr4Parser#unaryOp}.
	 * @param ctx the parse tree
	 */
	void enterUnaryOp(lua_grammar_antlr4Parser.UnaryOpContext ctx);
	/**
	 * Exit a parse tree produced by {@link lua_grammar_antlr4Parser#unaryOp}.
	 * @param ctx the parse tree
	 */
	void exitUnaryOp(lua_grammar_antlr4Parser.UnaryOpContext ctx);
	/**
	 * Enter a parse tree produced by {@link lua_grammar_antlr4Parser#concatOp}.
	 * @param ctx the parse tree
	 */
	void enterConcatOp(lua_grammar_antlr4Parser.ConcatOpContext ctx);
	/**
	 * Exit a parse tree produced by {@link lua_grammar_antlr4Parser#concatOp}.
	 * @param ctx the parse tree
	 */
	void exitConcatOp(lua_grammar_antlr4Parser.ConcatOpContext ctx);
	/**
	 * Enter a parse tree produced by {@link lua_grammar_antlr4Parser#varargOp}.
	 * @param ctx the parse tree
	 */
	void enterVarargOp(lua_grammar_antlr4Parser.VarargOpContext ctx);
	/**
	 * Exit a parse tree produced by {@link lua_grammar_antlr4Parser#varargOp}.
	 * @param ctx the parse tree
	 */
	void exitVarargOp(lua_grammar_antlr4Parser.VarargOpContext ctx);
	/**
	 * Enter a parse tree produced by {@link lua_grammar_antlr4Parser#compoundAssignOp}.
	 * @param ctx the parse tree
	 */
	void enterCompoundAssignOp(lua_grammar_antlr4Parser.CompoundAssignOpContext ctx);
	/**
	 * Exit a parse tree produced by {@link lua_grammar_antlr4Parser#compoundAssignOp}.
	 * @param ctx the parse tree
	 */
	void exitCompoundAssignOp(lua_grammar_antlr4Parser.CompoundAssignOpContext ctx);
	/**
	 * Enter a parse tree produced by {@link lua_grammar_antlr4Parser#incrOp}.
	 * @param ctx the parse tree
	 */
	void enterIncrOp(lua_grammar_antlr4Parser.IncrOpContext ctx);
	/**
	 * Exit a parse tree produced by {@link lua_grammar_antlr4Parser#incrOp}.
	 * @param ctx the parse tree
	 */
	void exitIncrOp(lua_grammar_antlr4Parser.IncrOpContext ctx);
	/**
	 * Enter a parse tree produced by {@link lua_grammar_antlr4Parser#coalesceOp}.
	 * @param ctx the parse tree
	 */
	void enterCoalesceOp(lua_grammar_antlr4Parser.CoalesceOpContext ctx);
	/**
	 * Exit a parse tree produced by {@link lua_grammar_antlr4Parser#coalesceOp}.
	 * @param ctx the parse tree
	 */
	void exitCoalesceOp(lua_grammar_antlr4Parser.CoalesceOpContext ctx);
	/**
	 * Enter a parse tree produced by {@link lua_grammar_antlr4Parser#shiftAssignOp}.
	 * @param ctx the parse tree
	 */
	void enterShiftAssignOp(lua_grammar_antlr4Parser.ShiftAssignOpContext ctx);
	/**
	 * Exit a parse tree produced by {@link lua_grammar_antlr4Parser#shiftAssignOp}.
	 * @param ctx the parse tree
	 */
	void exitShiftAssignOp(lua_grammar_antlr4Parser.ShiftAssignOpContext ctx);
}