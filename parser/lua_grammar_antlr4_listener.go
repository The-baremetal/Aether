// Code generated from lua_grammar_antlr4.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // lua_grammar_antlr4

import "github.com/antlr4-go/antlr/v4"

// lua_grammar_antlr4Listener is a complete listener for a parse tree produced by lua_grammar_antlr4Parser.
type lua_grammar_antlr4Listener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterControl_statement is called when entering the control_statement production.
	EnterControl_statement(c *Control_statementContext)

	// EnterStatement_terminator is called when entering the statement_terminator production.
	EnterStatement_terminator(c *Statement_terminatorContext)

	// EnterAssignment is called when entering the assignment production.
	EnterAssignment(c *AssignmentContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterPrefix_expression is called when entering the prefix_expression production.
	EnterPrefix_expression(c *Prefix_expressionContext)

	// EnterPrimary_expression is called when entering the primary_expression production.
	EnterPrimary_expression(c *Primary_expressionContext)

	// EnterOperators is called when entering the operators production.
	EnterOperators(c *OperatorsContext)

	// EnterComparison_operator is called when entering the comparison_operator production.
	EnterComparison_operator(c *Comparison_operatorContext)

	// EnterArith_operator is called when entering the arith_operator production.
	EnterArith_operator(c *Arith_operatorContext)

	// EnterLogical_operator is called when entering the logical_operator production.
	EnterLogical_operator(c *Logical_operatorContext)

	// EnterBitwise_operator is called when entering the bitwise_operator production.
	EnterBitwise_operator(c *Bitwise_operatorContext)

	// EnterLiteral is called when entering the literal production.
	EnterLiteral(c *LiteralContext)

	// EnterFunction_call is called when entering the function_call production.
	EnterFunction_call(c *Function_callContext)

	// EnterTable_insert is called when entering the table_insert production.
	EnterTable_insert(c *Table_insertContext)

	// EnterFunction_declaration is called when entering the function_declaration production.
	EnterFunction_declaration(c *Function_declarationContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterIf_statement is called when entering the if_statement production.
	EnterIf_statement(c *If_statementContext)

	// EnterFor_statement is called when entering the for_statement production.
	EnterFor_statement(c *For_statementContext)

	// EnterWhile_statement is called when entering the while_statement production.
	EnterWhile_statement(c *While_statementContext)

	// EnterTable is called when entering the table production.
	EnterTable(c *TableContext)

	// EnterField is called when entering the field production.
	EnterField(c *FieldContext)

	// EnterTable_access is called when entering the table_access production.
	EnterTable_access(c *Table_accessContext)

	// EnterSingle_line_comment is called when entering the single_line_comment production.
	EnterSingle_line_comment(c *Single_line_commentContext)

	// EnterPrint_statement is called when entering the print_statement production.
	EnterPrint_statement(c *Print_statementContext)

	// EnterIdentifier is called when entering the identifier production.
	EnterIdentifier(c *IdentifierContext)

	// EnterRepeat_statement is called when entering the repeat_statement production.
	EnterRepeat_statement(c *Repeat_statementContext)

	// EnterIdentifier_list is called when entering the identifier_list production.
	EnterIdentifier_list(c *Identifier_listContext)

	// EnterExpression_list is called when entering the expression_list production.
	EnterExpression_list(c *Expression_listContext)

	// EnterReturn_statement is called when entering the return_statement production.
	EnterReturn_statement(c *Return_statementContext)

	// EnterBreak_statement is called when entering the break_statement production.
	EnterBreak_statement(c *Break_statementContext)

	// EnterGoto_statement is called when entering the goto_statement production.
	EnterGoto_statement(c *Goto_statementContext)

	// EnterLabel_statement is called when entering the label_statement production.
	EnterLabel_statement(c *Label_statementContext)

	// EnterFunction_expression is called when entering the function_expression production.
	EnterFunction_expression(c *Function_expressionContext)

	// EnterMethod_call is called when entering the method_call production.
	EnterMethod_call(c *Method_callContext)

	// EnterMetatable_field is called when entering the metatable_field production.
	EnterMetatable_field(c *Metatable_fieldContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitControl_statement is called when exiting the control_statement production.
	ExitControl_statement(c *Control_statementContext)

	// ExitStatement_terminator is called when exiting the statement_terminator production.
	ExitStatement_terminator(c *Statement_terminatorContext)

	// ExitAssignment is called when exiting the assignment production.
	ExitAssignment(c *AssignmentContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitPrefix_expression is called when exiting the prefix_expression production.
	ExitPrefix_expression(c *Prefix_expressionContext)

	// ExitPrimary_expression is called when exiting the primary_expression production.
	ExitPrimary_expression(c *Primary_expressionContext)

	// ExitOperators is called when exiting the operators production.
	ExitOperators(c *OperatorsContext)

	// ExitComparison_operator is called when exiting the comparison_operator production.
	ExitComparison_operator(c *Comparison_operatorContext)

	// ExitArith_operator is called when exiting the arith_operator production.
	ExitArith_operator(c *Arith_operatorContext)

	// ExitLogical_operator is called when exiting the logical_operator production.
	ExitLogical_operator(c *Logical_operatorContext)

	// ExitBitwise_operator is called when exiting the bitwise_operator production.
	ExitBitwise_operator(c *Bitwise_operatorContext)

	// ExitLiteral is called when exiting the literal production.
	ExitLiteral(c *LiteralContext)

	// ExitFunction_call is called when exiting the function_call production.
	ExitFunction_call(c *Function_callContext)

	// ExitTable_insert is called when exiting the table_insert production.
	ExitTable_insert(c *Table_insertContext)

	// ExitFunction_declaration is called when exiting the function_declaration production.
	ExitFunction_declaration(c *Function_declarationContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitIf_statement is called when exiting the if_statement production.
	ExitIf_statement(c *If_statementContext)

	// ExitFor_statement is called when exiting the for_statement production.
	ExitFor_statement(c *For_statementContext)

	// ExitWhile_statement is called when exiting the while_statement production.
	ExitWhile_statement(c *While_statementContext)

	// ExitTable is called when exiting the table production.
	ExitTable(c *TableContext)

	// ExitField is called when exiting the field production.
	ExitField(c *FieldContext)

	// ExitTable_access is called when exiting the table_access production.
	ExitTable_access(c *Table_accessContext)

	// ExitSingle_line_comment is called when exiting the single_line_comment production.
	ExitSingle_line_comment(c *Single_line_commentContext)

	// ExitPrint_statement is called when exiting the print_statement production.
	ExitPrint_statement(c *Print_statementContext)

	// ExitIdentifier is called when exiting the identifier production.
	ExitIdentifier(c *IdentifierContext)

	// ExitRepeat_statement is called when exiting the repeat_statement production.
	ExitRepeat_statement(c *Repeat_statementContext)

	// ExitIdentifier_list is called when exiting the identifier_list production.
	ExitIdentifier_list(c *Identifier_listContext)

	// ExitExpression_list is called when exiting the expression_list production.
	ExitExpression_list(c *Expression_listContext)

	// ExitReturn_statement is called when exiting the return_statement production.
	ExitReturn_statement(c *Return_statementContext)

	// ExitBreak_statement is called when exiting the break_statement production.
	ExitBreak_statement(c *Break_statementContext)

	// ExitGoto_statement is called when exiting the goto_statement production.
	ExitGoto_statement(c *Goto_statementContext)

	// ExitLabel_statement is called when exiting the label_statement production.
	ExitLabel_statement(c *Label_statementContext)

	// ExitFunction_expression is called when exiting the function_expression production.
	ExitFunction_expression(c *Function_expressionContext)

	// ExitMethod_call is called when exiting the method_call production.
	ExitMethod_call(c *Method_callContext)

	// ExitMetatable_field is called when exiting the metatable_field production.
	ExitMetatable_field(c *Metatable_fieldContext)
}
