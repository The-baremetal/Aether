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

func (v *Baselua_grammar_antlr4Visitor) VisitControl_statement(ctx *Control_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baselua_grammar_antlr4Visitor) VisitStatement_terminator(ctx *Statement_terminatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baselua_grammar_antlr4Visitor) VisitAssignment(ctx *AssignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baselua_grammar_antlr4Visitor) VisitExpression(ctx *ExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baselua_grammar_antlr4Visitor) VisitPrefix_expression(ctx *Prefix_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baselua_grammar_antlr4Visitor) VisitPrimary_expression(ctx *Primary_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baselua_grammar_antlr4Visitor) VisitOperators(ctx *OperatorsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baselua_grammar_antlr4Visitor) VisitComparison_operator(ctx *Comparison_operatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baselua_grammar_antlr4Visitor) VisitArith_operator(ctx *Arith_operatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baselua_grammar_antlr4Visitor) VisitLogical_operator(ctx *Logical_operatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baselua_grammar_antlr4Visitor) VisitBitwise_operator(ctx *Bitwise_operatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baselua_grammar_antlr4Visitor) VisitConcat_operator(ctx *Concat_operatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baselua_grammar_antlr4Visitor) VisitLiteral(ctx *LiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baselua_grammar_antlr4Visitor) VisitFunction_call(ctx *Function_callContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baselua_grammar_antlr4Visitor) VisitTable_insert(ctx *Table_insertContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baselua_grammar_antlr4Visitor) VisitFunction_declaration(ctx *Function_declarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baselua_grammar_antlr4Visitor) VisitBlock(ctx *BlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baselua_grammar_antlr4Visitor) VisitIf_statement(ctx *If_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baselua_grammar_antlr4Visitor) VisitFor_statement(ctx *For_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baselua_grammar_antlr4Visitor) VisitWhile_statement(ctx *While_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baselua_grammar_antlr4Visitor) VisitTable(ctx *TableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baselua_grammar_antlr4Visitor) VisitField(ctx *FieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baselua_grammar_antlr4Visitor) VisitTable_access(ctx *Table_accessContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baselua_grammar_antlr4Visitor) VisitSingle_line_comment(ctx *Single_line_commentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baselua_grammar_antlr4Visitor) VisitPrint_statement(ctx *Print_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baselua_grammar_antlr4Visitor) VisitIdentifier(ctx *IdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baselua_grammar_antlr4Visitor) VisitRepeat_statement(ctx *Repeat_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baselua_grammar_antlr4Visitor) VisitIdentifier_list(ctx *Identifier_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baselua_grammar_antlr4Visitor) VisitExpression_list(ctx *Expression_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baselua_grammar_antlr4Visitor) VisitReturn_statement(ctx *Return_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baselua_grammar_antlr4Visitor) VisitBreak_statement(ctx *Break_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baselua_grammar_antlr4Visitor) VisitGoto_statement(ctx *Goto_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baselua_grammar_antlr4Visitor) VisitLabel_statement(ctx *Label_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baselua_grammar_antlr4Visitor) VisitFunction_expression(ctx *Function_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baselua_grammar_antlr4Visitor) VisitMethod_call(ctx *Method_callContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baselua_grammar_antlr4Visitor) VisitMetatable_field(ctx *Metatable_fieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baselua_grammar_antlr4Visitor) VisitMetamethod(ctx *MetamethodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Baselua_grammar_antlr4Visitor) VisitCoroutine_statement(ctx *Coroutine_statementContext) interface{} {
	return v.VisitChildren(ctx)
}
