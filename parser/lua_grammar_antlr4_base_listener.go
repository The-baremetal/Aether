// Code generated from lua_grammar_antlr4.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // lua_grammar_antlr4

import "github.com/antlr4-go/antlr/v4"

// Baselua_grammar_antlr4Listener is a complete listener for a parse tree produced by lua_grammar_antlr4Parser.
type Baselua_grammar_antlr4Listener struct{}

var _ lua_grammar_antlr4Listener = &Baselua_grammar_antlr4Listener{}

// VisitTerminal is called when a terminal node is visited.
func (s *Baselua_grammar_antlr4Listener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *Baselua_grammar_antlr4Listener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *Baselua_grammar_antlr4Listener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *Baselua_grammar_antlr4Listener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *Baselua_grammar_antlr4Listener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *Baselua_grammar_antlr4Listener) ExitProgram(ctx *ProgramContext) {}

// EnterStatement is called when production statement is entered.
func (s *Baselua_grammar_antlr4Listener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *Baselua_grammar_antlr4Listener) ExitStatement(ctx *StatementContext) {}

// EnterControl_statement is called when production control_statement is entered.
func (s *Baselua_grammar_antlr4Listener) EnterControl_statement(ctx *Control_statementContext) {}

// ExitControl_statement is called when production control_statement is exited.
func (s *Baselua_grammar_antlr4Listener) ExitControl_statement(ctx *Control_statementContext) {}

// EnterStatement_terminator is called when production statement_terminator is entered.
func (s *Baselua_grammar_antlr4Listener) EnterStatement_terminator(ctx *Statement_terminatorContext) {
}

// ExitStatement_terminator is called when production statement_terminator is exited.
func (s *Baselua_grammar_antlr4Listener) ExitStatement_terminator(ctx *Statement_terminatorContext) {}

// EnterAssignment is called when production assignment is entered.
func (s *Baselua_grammar_antlr4Listener) EnterAssignment(ctx *AssignmentContext) {}

// ExitAssignment is called when production assignment is exited.
func (s *Baselua_grammar_antlr4Listener) ExitAssignment(ctx *AssignmentContext) {}

// EnterExpression is called when production expression is entered.
func (s *Baselua_grammar_antlr4Listener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *Baselua_grammar_antlr4Listener) ExitExpression(ctx *ExpressionContext) {}

// EnterPrefix_expression is called when production prefix_expression is entered.
func (s *Baselua_grammar_antlr4Listener) EnterPrefix_expression(ctx *Prefix_expressionContext) {}

// ExitPrefix_expression is called when production prefix_expression is exited.
func (s *Baselua_grammar_antlr4Listener) ExitPrefix_expression(ctx *Prefix_expressionContext) {}

// EnterPrimary_expression is called when production primary_expression is entered.
func (s *Baselua_grammar_antlr4Listener) EnterPrimary_expression(ctx *Primary_expressionContext) {}

// ExitPrimary_expression is called when production primary_expression is exited.
func (s *Baselua_grammar_antlr4Listener) ExitPrimary_expression(ctx *Primary_expressionContext) {}

// EnterOperators is called when production operators is entered.
func (s *Baselua_grammar_antlr4Listener) EnterOperators(ctx *OperatorsContext) {}

// ExitOperators is called when production operators is exited.
func (s *Baselua_grammar_antlr4Listener) ExitOperators(ctx *OperatorsContext) {}

// EnterComparison_operator is called when production comparison_operator is entered.
func (s *Baselua_grammar_antlr4Listener) EnterComparison_operator(ctx *Comparison_operatorContext) {}

// ExitComparison_operator is called when production comparison_operator is exited.
func (s *Baselua_grammar_antlr4Listener) ExitComparison_operator(ctx *Comparison_operatorContext) {}

// EnterArith_operator is called when production arith_operator is entered.
func (s *Baselua_grammar_antlr4Listener) EnterArith_operator(ctx *Arith_operatorContext) {}

// ExitArith_operator is called when production arith_operator is exited.
func (s *Baselua_grammar_antlr4Listener) ExitArith_operator(ctx *Arith_operatorContext) {}

// EnterLogical_operator is called when production logical_operator is entered.
func (s *Baselua_grammar_antlr4Listener) EnterLogical_operator(ctx *Logical_operatorContext) {}

// ExitLogical_operator is called when production logical_operator is exited.
func (s *Baselua_grammar_antlr4Listener) ExitLogical_operator(ctx *Logical_operatorContext) {}

// EnterBitwise_operator is called when production bitwise_operator is entered.
func (s *Baselua_grammar_antlr4Listener) EnterBitwise_operator(ctx *Bitwise_operatorContext) {}

// ExitBitwise_operator is called when production bitwise_operator is exited.
func (s *Baselua_grammar_antlr4Listener) ExitBitwise_operator(ctx *Bitwise_operatorContext) {}

// EnterLiteral is called when production literal is entered.
func (s *Baselua_grammar_antlr4Listener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *Baselua_grammar_antlr4Listener) ExitLiteral(ctx *LiteralContext) {}

// EnterFunction_call is called when production function_call is entered.
func (s *Baselua_grammar_antlr4Listener) EnterFunction_call(ctx *Function_callContext) {}

// ExitFunction_call is called when production function_call is exited.
func (s *Baselua_grammar_antlr4Listener) ExitFunction_call(ctx *Function_callContext) {}

// EnterTable_insert is called when production table_insert is entered.
func (s *Baselua_grammar_antlr4Listener) EnterTable_insert(ctx *Table_insertContext) {}

// ExitTable_insert is called when production table_insert is exited.
func (s *Baselua_grammar_antlr4Listener) ExitTable_insert(ctx *Table_insertContext) {}

// EnterFunction_declaration is called when production function_declaration is entered.
func (s *Baselua_grammar_antlr4Listener) EnterFunction_declaration(ctx *Function_declarationContext) {
}

// ExitFunction_declaration is called when production function_declaration is exited.
func (s *Baselua_grammar_antlr4Listener) ExitFunction_declaration(ctx *Function_declarationContext) {}

// EnterBlock is called when production block is entered.
func (s *Baselua_grammar_antlr4Listener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *Baselua_grammar_antlr4Listener) ExitBlock(ctx *BlockContext) {}

// EnterIf_statement is called when production if_statement is entered.
func (s *Baselua_grammar_antlr4Listener) EnterIf_statement(ctx *If_statementContext) {}

// ExitIf_statement is called when production if_statement is exited.
func (s *Baselua_grammar_antlr4Listener) ExitIf_statement(ctx *If_statementContext) {}

// EnterFor_statement is called when production for_statement is entered.
func (s *Baselua_grammar_antlr4Listener) EnterFor_statement(ctx *For_statementContext) {}

// ExitFor_statement is called when production for_statement is exited.
func (s *Baselua_grammar_antlr4Listener) ExitFor_statement(ctx *For_statementContext) {}

// EnterWhile_statement is called when production while_statement is entered.
func (s *Baselua_grammar_antlr4Listener) EnterWhile_statement(ctx *While_statementContext) {}

// ExitWhile_statement is called when production while_statement is exited.
func (s *Baselua_grammar_antlr4Listener) ExitWhile_statement(ctx *While_statementContext) {}

// EnterTable is called when production table is entered.
func (s *Baselua_grammar_antlr4Listener) EnterTable(ctx *TableContext) {}

// ExitTable is called when production table is exited.
func (s *Baselua_grammar_antlr4Listener) ExitTable(ctx *TableContext) {}

// EnterField is called when production field is entered.
func (s *Baselua_grammar_antlr4Listener) EnterField(ctx *FieldContext) {}

// ExitField is called when production field is exited.
func (s *Baselua_grammar_antlr4Listener) ExitField(ctx *FieldContext) {}

// EnterTable_access is called when production table_access is entered.
func (s *Baselua_grammar_antlr4Listener) EnterTable_access(ctx *Table_accessContext) {}

// ExitTable_access is called when production table_access is exited.
func (s *Baselua_grammar_antlr4Listener) ExitTable_access(ctx *Table_accessContext) {}

// EnterSingle_line_comment is called when production single_line_comment is entered.
func (s *Baselua_grammar_antlr4Listener) EnterSingle_line_comment(ctx *Single_line_commentContext) {}

// ExitSingle_line_comment is called when production single_line_comment is exited.
func (s *Baselua_grammar_antlr4Listener) ExitSingle_line_comment(ctx *Single_line_commentContext) {}

// EnterPrint_statement is called when production print_statement is entered.
func (s *Baselua_grammar_antlr4Listener) EnterPrint_statement(ctx *Print_statementContext) {}

// ExitPrint_statement is called when production print_statement is exited.
func (s *Baselua_grammar_antlr4Listener) ExitPrint_statement(ctx *Print_statementContext) {}

// EnterIdentifier is called when production identifier is entered.
func (s *Baselua_grammar_antlr4Listener) EnterIdentifier(ctx *IdentifierContext) {}

// ExitIdentifier is called when production identifier is exited.
func (s *Baselua_grammar_antlr4Listener) ExitIdentifier(ctx *IdentifierContext) {}

// EnterRepeat_statement is called when production repeat_statement is entered.
func (s *Baselua_grammar_antlr4Listener) EnterRepeat_statement(ctx *Repeat_statementContext) {}

// ExitRepeat_statement is called when production repeat_statement is exited.
func (s *Baselua_grammar_antlr4Listener) ExitRepeat_statement(ctx *Repeat_statementContext) {}

// EnterIdentifier_list is called when production identifier_list is entered.
func (s *Baselua_grammar_antlr4Listener) EnterIdentifier_list(ctx *Identifier_listContext) {}

// ExitIdentifier_list is called when production identifier_list is exited.
func (s *Baselua_grammar_antlr4Listener) ExitIdentifier_list(ctx *Identifier_listContext) {}

// EnterExpression_list is called when production expression_list is entered.
func (s *Baselua_grammar_antlr4Listener) EnterExpression_list(ctx *Expression_listContext) {}

// ExitExpression_list is called when production expression_list is exited.
func (s *Baselua_grammar_antlr4Listener) ExitExpression_list(ctx *Expression_listContext) {}

// EnterReturn_statement is called when production return_statement is entered.
func (s *Baselua_grammar_antlr4Listener) EnterReturn_statement(ctx *Return_statementContext) {}

// ExitReturn_statement is called when production return_statement is exited.
func (s *Baselua_grammar_antlr4Listener) ExitReturn_statement(ctx *Return_statementContext) {}

// EnterBreak_statement is called when production break_statement is entered.
func (s *Baselua_grammar_antlr4Listener) EnterBreak_statement(ctx *Break_statementContext) {}

// ExitBreak_statement is called when production break_statement is exited.
func (s *Baselua_grammar_antlr4Listener) ExitBreak_statement(ctx *Break_statementContext) {}

// EnterGoto_statement is called when production goto_statement is entered.
func (s *Baselua_grammar_antlr4Listener) EnterGoto_statement(ctx *Goto_statementContext) {}

// ExitGoto_statement is called when production goto_statement is exited.
func (s *Baselua_grammar_antlr4Listener) ExitGoto_statement(ctx *Goto_statementContext) {}

// EnterLabel_statement is called when production label_statement is entered.
func (s *Baselua_grammar_antlr4Listener) EnterLabel_statement(ctx *Label_statementContext) {}

// ExitLabel_statement is called when production label_statement is exited.
func (s *Baselua_grammar_antlr4Listener) ExitLabel_statement(ctx *Label_statementContext) {}

// EnterFunction_expression is called when production function_expression is entered.
func (s *Baselua_grammar_antlr4Listener) EnterFunction_expression(ctx *Function_expressionContext) {}

// ExitFunction_expression is called when production function_expression is exited.
func (s *Baselua_grammar_antlr4Listener) ExitFunction_expression(ctx *Function_expressionContext) {}

// EnterMethod_call is called when production method_call is entered.
func (s *Baselua_grammar_antlr4Listener) EnterMethod_call(ctx *Method_callContext) {}

// ExitMethod_call is called when production method_call is exited.
func (s *Baselua_grammar_antlr4Listener) ExitMethod_call(ctx *Method_callContext) {}

// EnterMetatable_field is called when production metatable_field is entered.
func (s *Baselua_grammar_antlr4Listener) EnterMetatable_field(ctx *Metatable_fieldContext) {}

// ExitMetatable_field is called when production metatable_field is exited.
func (s *Baselua_grammar_antlr4Listener) ExitMetatable_field(ctx *Metatable_fieldContext) {}
