// Code generated from Lua_grammar_antlr4.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Lua_grammar_antlr4

import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by Lua_grammar_antlr4Parser.
type Lua_grammar_antlr4Visitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by Lua_grammar_antlr4Parser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by Lua_grammar_antlr4Parser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by Lua_grammar_antlr4Parser#assignStatement.
	VisitAssignStatement(ctx *AssignStatementContext) interface{}

	// Visit a parse tree produced by Lua_grammar_antlr4Parser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by Lua_grammar_antlr4Parser#primaryExpression.
	VisitPrimaryExpression(ctx *PrimaryExpressionContext) interface{}

	// Visit a parse tree produced by Lua_grammar_antlr4Parser#methodChain.
	VisitMethodChain(ctx *MethodChainContext) interface{}

	// Visit a parse tree produced by Lua_grammar_antlr4Parser#propertyChain.
	VisitPropertyChain(ctx *PropertyChainContext) interface{}

	// Visit a parse tree produced by Lua_grammar_antlr4Parser#indexChain.
	VisitIndexChain(ctx *IndexChainContext) interface{}

	// Visit a parse tree produced by Lua_grammar_antlr4Parser#literal.
	VisitLiteral(ctx *LiteralContext) interface{}

	// Visit a parse tree produced by Lua_grammar_antlr4Parser#variable.
	VisitVariable(ctx *VariableContext) interface{}

	// Visit a parse tree produced by Lua_grammar_antlr4Parser#safeAccess.
	VisitSafeAccess(ctx *SafeAccessContext) interface{}

	// Visit a parse tree produced by Lua_grammar_antlr4Parser#functionCall.
	VisitFunctionCall(ctx *FunctionCallContext) interface{}

	// Visit a parse tree produced by Lua_grammar_antlr4Parser#tableConstructor.
	VisitTableConstructor(ctx *TableConstructorContext) interface{}

	// Visit a parse tree produced by Lua_grammar_antlr4Parser#metatable.
	VisitMetatable(ctx *MetatableContext) interface{}

	// Visit a parse tree produced by Lua_grammar_antlr4Parser#metamethods.
	VisitMetamethods(ctx *MetamethodsContext) interface{}

	// Visit a parse tree produced by Lua_grammar_antlr4Parser#labelStatement.
	VisitLabelStatement(ctx *LabelStatementContext) interface{}

	// Visit a parse tree produced by Lua_grammar_antlr4Parser#matchArm.
	VisitMatchArm(ctx *MatchArmContext) interface{}

	// Visit a parse tree produced by Lua_grammar_antlr4Parser#pattern.
	VisitPattern(ctx *PatternContext) interface{}

	// Visit a parse tree produced by Lua_grammar_antlr4Parser#fieldPattern.
	VisitFieldPattern(ctx *FieldPatternContext) interface{}

	// Visit a parse tree produced by Lua_grammar_antlr4Parser#metamethod.
	VisitMetamethod(ctx *MetamethodContext) interface{}

	// Visit a parse tree produced by Lua_grammar_antlr4Parser#field.
	VisitField(ctx *FieldContext) interface{}

	// Visit a parse tree produced by Lua_grammar_antlr4Parser#binaryOperation.
	VisitBinaryOperation(ctx *BinaryOperationContext) interface{}

	// Visit a parse tree produced by Lua_grammar_antlr4Parser#unaryOperation.
	VisitUnaryOperation(ctx *UnaryOperationContext) interface{}

	// Visit a parse tree produced by Lua_grammar_antlr4Parser#controlFlowStatement.
	VisitControlFlowStatement(ctx *ControlFlowStatementContext) interface{}

	// Visit a parse tree produced by Lua_grammar_antlr4Parser#ifStatement.
	VisitIfStatement(ctx *IfStatementContext) interface{}

	// Visit a parse tree produced by Lua_grammar_antlr4Parser#whileStatement.
	VisitWhileStatement(ctx *WhileStatementContext) interface{}

	// Visit a parse tree produced by Lua_grammar_antlr4Parser#repeatStatement.
	VisitRepeatStatement(ctx *RepeatStatementContext) interface{}

	// Visit a parse tree produced by Lua_grammar_antlr4Parser#numericFor.
	VisitNumericFor(ctx *NumericForContext) interface{}

	// Visit a parse tree produced by Lua_grammar_antlr4Parser#genericFor.
	VisitGenericFor(ctx *GenericForContext) interface{}

	// Visit a parse tree produced by Lua_grammar_antlr4Parser#breakStatement.
	VisitBreakStatement(ctx *BreakStatementContext) interface{}

	// Visit a parse tree produced by Lua_grammar_antlr4Parser#gotoStatement.
	VisitGotoStatement(ctx *GotoStatementContext) interface{}

	// Visit a parse tree produced by Lua_grammar_antlr4Parser#protectedCallStatement.
	VisitProtectedCallStatement(ctx *ProtectedCallStatementContext) interface{}

	// Visit a parse tree produced by Lua_grammar_antlr4Parser#namedArgs.
	VisitNamedArgs(ctx *NamedArgsContext) interface{}

	// Visit a parse tree produced by Lua_grammar_antlr4Parser#block.
	VisitBlock(ctx *BlockContext) interface{}

	// Visit a parse tree produced by Lua_grammar_antlr4Parser#localDeclaration.
	VisitLocalDeclaration(ctx *LocalDeclarationContext) interface{}

	// Visit a parse tree produced by Lua_grammar_antlr4Parser#functionDeclaration.
	VisitFunctionDeclaration(ctx *FunctionDeclarationContext) interface{}

	// Visit a parse tree produced by Lua_grammar_antlr4Parser#returnStatement.
	VisitReturnStatement(ctx *ReturnStatementContext) interface{}

	// Visit a parse tree produced by Lua_grammar_antlr4Parser#operatorGroup.
	VisitOperatorGroup(ctx *OperatorGroupContext) interface{}

	// Visit a parse tree produced by Lua_grammar_antlr4Parser#logicalOp.
	VisitLogicalOp(ctx *LogicalOpContext) interface{}

	// Visit a parse tree produced by Lua_grammar_antlr4Parser#comparisonOp.
	VisitComparisonOp(ctx *ComparisonOpContext) interface{}

	// Visit a parse tree produced by Lua_grammar_antlr4Parser#arithOp.
	VisitArithOp(ctx *ArithOpContext) interface{}

	// Visit a parse tree produced by Lua_grammar_antlr4Parser#bitwiseOp.
	VisitBitwiseOp(ctx *BitwiseOpContext) interface{}

	// Visit a parse tree produced by Lua_grammar_antlr4Parser#assignOp.
	VisitAssignOp(ctx *AssignOpContext) interface{}

	// Visit a parse tree produced by Lua_grammar_antlr4Parser#unaryOp.
	VisitUnaryOp(ctx *UnaryOpContext) interface{}

	// Visit a parse tree produced by Lua_grammar_antlr4Parser#concatOp.
	VisitConcatOp(ctx *ConcatOpContext) interface{}

	// Visit a parse tree produced by Lua_grammar_antlr4Parser#varargOp.
	VisitVarargOp(ctx *VarargOpContext) interface{}

	// Visit a parse tree produced by Lua_grammar_antlr4Parser#compoundAssignOp.
	VisitCompoundAssignOp(ctx *CompoundAssignOpContext) interface{}

	// Visit a parse tree produced by Lua_grammar_antlr4Parser#incrOp.
	VisitIncrOp(ctx *IncrOpContext) interface{}

	// Visit a parse tree produced by Lua_grammar_antlr4Parser#coalesceOp.
	VisitCoalesceOp(ctx *CoalesceOpContext) interface{}

	// Visit a parse tree produced by Lua_grammar_antlr4Parser#shiftAssignOp.
	VisitShiftAssignOp(ctx *ShiftAssignOpContext) interface{}

	// Visit a parse tree produced by Lua_grammar_antlr4Parser#nonNullAssertOp.
	VisitNonNullAssertOp(ctx *NonNullAssertOpContext) interface{}

	// Visit a parse tree produced by Lua_grammar_antlr4Parser#identifierList.
	VisitIdentifierList(ctx *IdentifierListContext) interface{}

	// Visit a parse tree produced by Lua_grammar_antlr4Parser#expressionList.
	VisitExpressionList(ctx *ExpressionListContext) interface{}

	// Visit a parse tree produced by Lua_grammar_antlr4Parser#functionExpression.
	VisitFunctionExpression(ctx *FunctionExpressionContext) interface{}

	// Visit a parse tree produced by Lua_grammar_antlr4Parser#parameterList.
	VisitParameterList(ctx *ParameterListContext) interface{}

	// Visit a parse tree produced by Lua_grammar_antlr4Parser#selectStatement.
	VisitSelectStatement(ctx *SelectStatementContext) interface{}

	// Visit a parse tree produced by Lua_grammar_antlr4Parser#lambdaExpression.
	VisitLambdaExpression(ctx *LambdaExpressionContext) interface{}

	// Visit a parse tree produced by Lua_grammar_antlr4Parser#typeAnnotation.
	VisitTypeAnnotation(ctx *TypeAnnotationContext) interface{}

	// Visit a parse tree produced by Lua_grammar_antlr4Parser#typeSpec.
	VisitTypeSpec(ctx *TypeSpecContext) interface{}

	// Visit a parse tree produced by Lua_grammar_antlr4Parser#experimentalExpression.
	VisitExperimentalExpression(ctx *ExperimentalExpressionContext) interface{}

	// Visit a parse tree produced by Lua_grammar_antlr4Parser#safeNavigation.
	VisitSafeNavigation(ctx *SafeNavigationContext) interface{}

	// Visit a parse tree produced by Lua_grammar_antlr4Parser#pipeOperator.
	VisitPipeOperator(ctx *PipeOperatorContext) interface{}

	// Visit a parse tree produced by Lua_grammar_antlr4Parser#decoratorSyntax.
	VisitDecoratorSyntax(ctx *DecoratorSyntaxContext) interface{}
}
