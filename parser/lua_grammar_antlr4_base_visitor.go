// Code generated from Lua_grammar_antlr4.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Lua_grammar_antlr4

import "github.com/antlr4-go/antlr/v4"

type BaseLua_grammar_antlr4Visitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseLua_grammar_antlr4Visitor) VisitProgram(ctx *ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLua_grammar_antlr4Visitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLua_grammar_antlr4Visitor) VisitAssignStatement(ctx *AssignStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLua_grammar_antlr4Visitor) VisitExpression(ctx *ExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLua_grammar_antlr4Visitor) VisitPrimaryExpression(ctx *PrimaryExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLua_grammar_antlr4Visitor) VisitMethodChain(ctx *MethodChainContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLua_grammar_antlr4Visitor) VisitPropertyChain(ctx *PropertyChainContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLua_grammar_antlr4Visitor) VisitIndexChain(ctx *IndexChainContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLua_grammar_antlr4Visitor) VisitLiteral(ctx *LiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLua_grammar_antlr4Visitor) VisitVariable(ctx *VariableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLua_grammar_antlr4Visitor) VisitSafeAccess(ctx *SafeAccessContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLua_grammar_antlr4Visitor) VisitFunctionCall(ctx *FunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLua_grammar_antlr4Visitor) VisitTableConstructor(ctx *TableConstructorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLua_grammar_antlr4Visitor) VisitMetatable(ctx *MetatableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLua_grammar_antlr4Visitor) VisitMetamethods(ctx *MetamethodsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLua_grammar_antlr4Visitor) VisitLabelStatement(ctx *LabelStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLua_grammar_antlr4Visitor) VisitMatchArm(ctx *MatchArmContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLua_grammar_antlr4Visitor) VisitPattern(ctx *PatternContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLua_grammar_antlr4Visitor) VisitFieldPattern(ctx *FieldPatternContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLua_grammar_antlr4Visitor) VisitMetamethod(ctx *MetamethodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLua_grammar_antlr4Visitor) VisitField(ctx *FieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLua_grammar_antlr4Visitor) VisitBinaryOperation(ctx *BinaryOperationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLua_grammar_antlr4Visitor) VisitUnaryOperation(ctx *UnaryOperationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLua_grammar_antlr4Visitor) VisitControlFlowStatement(ctx *ControlFlowStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLua_grammar_antlr4Visitor) VisitIfStatement(ctx *IfStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLua_grammar_antlr4Visitor) VisitWhileStatement(ctx *WhileStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLua_grammar_antlr4Visitor) VisitRepeatStatement(ctx *RepeatStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLua_grammar_antlr4Visitor) VisitNumericFor(ctx *NumericForContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLua_grammar_antlr4Visitor) VisitGenericFor(ctx *GenericForContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLua_grammar_antlr4Visitor) VisitBreakStatement(ctx *BreakStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLua_grammar_antlr4Visitor) VisitGotoStatement(ctx *GotoStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLua_grammar_antlr4Visitor) VisitCoroutineStatement(ctx *CoroutineStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLua_grammar_antlr4Visitor) VisitProtectedCallStatement(ctx *ProtectedCallStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLua_grammar_antlr4Visitor) VisitNamedArgs(ctx *NamedArgsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLua_grammar_antlr4Visitor) VisitBlock(ctx *BlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLua_grammar_antlr4Visitor) VisitLocalDeclaration(ctx *LocalDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLua_grammar_antlr4Visitor) VisitFunctionDeclaration(ctx *FunctionDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLua_grammar_antlr4Visitor) VisitReturnStatement(ctx *ReturnStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLua_grammar_antlr4Visitor) VisitOperatorGroup(ctx *OperatorGroupContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLua_grammar_antlr4Visitor) VisitLogicalOp(ctx *LogicalOpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLua_grammar_antlr4Visitor) VisitComparisonOp(ctx *ComparisonOpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLua_grammar_antlr4Visitor) VisitArithOp(ctx *ArithOpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLua_grammar_antlr4Visitor) VisitBitwiseOp(ctx *BitwiseOpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLua_grammar_antlr4Visitor) VisitAssignOp(ctx *AssignOpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLua_grammar_antlr4Visitor) VisitUnaryOp(ctx *UnaryOpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLua_grammar_antlr4Visitor) VisitConcatOp(ctx *ConcatOpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLua_grammar_antlr4Visitor) VisitVarargOp(ctx *VarargOpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLua_grammar_antlr4Visitor) VisitCompoundAssignOp(ctx *CompoundAssignOpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLua_grammar_antlr4Visitor) VisitIncrOp(ctx *IncrOpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLua_grammar_antlr4Visitor) VisitCoalesceOp(ctx *CoalesceOpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLua_grammar_antlr4Visitor) VisitShiftAssignOp(ctx *ShiftAssignOpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLua_grammar_antlr4Visitor) VisitNonNullAssertOp(ctx *NonNullAssertOpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLua_grammar_antlr4Visitor) VisitIdentifierList(ctx *IdentifierListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLua_grammar_antlr4Visitor) VisitExpressionList(ctx *ExpressionListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLua_grammar_antlr4Visitor) VisitFunctionExpression(ctx *FunctionExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLua_grammar_antlr4Visitor) VisitSelectStatement(ctx *SelectStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLua_grammar_antlr4Visitor) VisitLambdaExpression(ctx *LambdaExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLua_grammar_antlr4Visitor) VisitTypeAnnotation(ctx *TypeAnnotationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLua_grammar_antlr4Visitor) VisitTypeSpec(ctx *TypeSpecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLua_grammar_antlr4Visitor) VisitExperimentalExpression(ctx *ExperimentalExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLua_grammar_antlr4Visitor) VisitSafeNavigation(ctx *SafeNavigationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLua_grammar_antlr4Visitor) VisitPipeOperator(ctx *PipeOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLua_grammar_antlr4Visitor) VisitDecoratorSyntax(ctx *DecoratorSyntaxContext) interface{} {
	return v.VisitChildren(ctx)
}
