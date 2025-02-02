// Code generated from lua_grammar_antlr4.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // lua_grammar_antlr4

import "github.com/antlr4-go/antlr/v4"

type Baselua_grammar_antlr4Visitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *Baselua_grammar_antlr4Visitor) VisitProgram(ctx *ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baselua_grammar_antlr4Visitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baselua_grammar_antlr4Visitor) VisitAssignStatement(ctx *AssignStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baselua_grammar_antlr4Visitor) VisitExpression(ctx *ExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baselua_grammar_antlr4Visitor) VisitPrimaryExpression(ctx *PrimaryExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baselua_grammar_antlr4Visitor) VisitLiteral(ctx *LiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baselua_grammar_antlr4Visitor) VisitVariable(ctx *VariableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baselua_grammar_antlr4Visitor) VisitFunctionCall(ctx *FunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baselua_grammar_antlr4Visitor) VisitTableConstructor(ctx *TableConstructorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baselua_grammar_antlr4Visitor) VisitMetatable(ctx *MetatableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baselua_grammar_antlr4Visitor) VisitMetamethods(ctx *MetamethodsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baselua_grammar_antlr4Visitor) VisitLabelStatement(ctx *LabelStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baselua_grammar_antlr4Visitor) VisitMetamethod(ctx *MetamethodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baselua_grammar_antlr4Visitor) VisitField(ctx *FieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baselua_grammar_antlr4Visitor) VisitBinaryOperation(ctx *BinaryOperationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baselua_grammar_antlr4Visitor) VisitUnaryOperation(ctx *UnaryOperationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baselua_grammar_antlr4Visitor) VisitControlFlowStatement(ctx *ControlFlowStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baselua_grammar_antlr4Visitor) VisitIfStatement(ctx *IfStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baselua_grammar_antlr4Visitor) VisitWhileStatement(ctx *WhileStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baselua_grammar_antlr4Visitor) VisitRepeatStatement(ctx *RepeatStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baselua_grammar_antlr4Visitor) VisitForStatement(ctx *ForStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baselua_grammar_antlr4Visitor) VisitBreakStatement(ctx *BreakStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baselua_grammar_antlr4Visitor) VisitGotoStatement(ctx *GotoStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baselua_grammar_antlr4Visitor) VisitCoroutineStatement(ctx *CoroutineStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baselua_grammar_antlr4Visitor) VisitBlock(ctx *BlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baselua_grammar_antlr4Visitor) VisitLocalDeclaration(ctx *LocalDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baselua_grammar_antlr4Visitor) VisitFunctionDeclaration(ctx *FunctionDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baselua_grammar_antlr4Visitor) VisitReturnStatement(ctx *ReturnStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baselua_grammar_antlr4Visitor) VisitOperatorGroup(ctx *OperatorGroupContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baselua_grammar_antlr4Visitor) VisitLogicalOp(ctx *LogicalOpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baselua_grammar_antlr4Visitor) VisitComparisonOp(ctx *ComparisonOpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baselua_grammar_antlr4Visitor) VisitArithOp(ctx *ArithOpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baselua_grammar_antlr4Visitor) VisitBitwiseOp(ctx *BitwiseOpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baselua_grammar_antlr4Visitor) VisitAssignOp(ctx *AssignOpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baselua_grammar_antlr4Visitor) VisitUnaryOp(ctx *UnaryOpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baselua_grammar_antlr4Visitor) VisitConcatOp(ctx *ConcatOpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baselua_grammar_antlr4Visitor) VisitVarargOp(ctx *VarargOpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baselua_grammar_antlr4Visitor) VisitCompoundAssignOp(ctx *CompoundAssignOpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baselua_grammar_antlr4Visitor) VisitIncrOp(ctx *IncrOpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baselua_grammar_antlr4Visitor) VisitCoalesceOp(ctx *CoalesceOpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baselua_grammar_antlr4Visitor) VisitShiftAssignOp(ctx *ShiftAssignOpContext) interface{} {
	return v.VisitChildren(ctx)
}
