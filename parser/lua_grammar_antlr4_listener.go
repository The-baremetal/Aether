// Code generated from Lua_grammar_antlr4.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Lua_grammar_antlr4

import "github.com/antlr4-go/antlr/v4"

// Lua_grammar_antlr4Listener is a complete listener for a parse tree produced by Lua_grammar_antlr4Parser.
type Lua_grammar_antlr4Listener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterAssignStatement is called when entering the assignStatement production.
	EnterAssignStatement(c *AssignStatementContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterPrimaryExpression is called when entering the primaryExpression production.
	EnterPrimaryExpression(c *PrimaryExpressionContext)

	// EnterMethodChain is called when entering the methodChain production.
	EnterMethodChain(c *MethodChainContext)

	// EnterPropertyChain is called when entering the propertyChain production.
	EnterPropertyChain(c *PropertyChainContext)

	// EnterIndexChain is called when entering the indexChain production.
	EnterIndexChain(c *IndexChainContext)

	// EnterLiteral is called when entering the literal production.
	EnterLiteral(c *LiteralContext)

	// EnterVariable is called when entering the variable production.
	EnterVariable(c *VariableContext)

	// EnterSafeAccess is called when entering the safeAccess production.
	EnterSafeAccess(c *SafeAccessContext)

	// EnterFunctionCall is called when entering the functionCall production.
	EnterFunctionCall(c *FunctionCallContext)

	// EnterTableConstructor is called when entering the tableConstructor production.
	EnterTableConstructor(c *TableConstructorContext)

	// EnterMetatable is called when entering the metatable production.
	EnterMetatable(c *MetatableContext)

	// EnterMetamethods is called when entering the metamethods production.
	EnterMetamethods(c *MetamethodsContext)

	// EnterLabelStatement is called when entering the labelStatement production.
	EnterLabelStatement(c *LabelStatementContext)

	// EnterMatchArm is called when entering the matchArm production.
	EnterMatchArm(c *MatchArmContext)

	// EnterPattern is called when entering the pattern production.
	EnterPattern(c *PatternContext)

	// EnterFieldPattern is called when entering the fieldPattern production.
	EnterFieldPattern(c *FieldPatternContext)

	// EnterMetamethod is called when entering the metamethod production.
	EnterMetamethod(c *MetamethodContext)

	// EnterField is called when entering the field production.
	EnterField(c *FieldContext)

	// EnterBinaryOperation is called when entering the binaryOperation production.
	EnterBinaryOperation(c *BinaryOperationContext)

	// EnterUnaryOperation is called when entering the unaryOperation production.
	EnterUnaryOperation(c *UnaryOperationContext)

	// EnterControlFlowStatement is called when entering the controlFlowStatement production.
	EnterControlFlowStatement(c *ControlFlowStatementContext)

	// EnterIfStatement is called when entering the ifStatement production.
	EnterIfStatement(c *IfStatementContext)

	// EnterWhileStatement is called when entering the whileStatement production.
	EnterWhileStatement(c *WhileStatementContext)

	// EnterRepeatStatement is called when entering the repeatStatement production.
	EnterRepeatStatement(c *RepeatStatementContext)

	// EnterNumericFor is called when entering the numericFor production.
	EnterNumericFor(c *NumericForContext)

	// EnterGenericFor is called when entering the genericFor production.
	EnterGenericFor(c *GenericForContext)

	// EnterBreakStatement is called when entering the breakStatement production.
	EnterBreakStatement(c *BreakStatementContext)

	// EnterGotoStatement is called when entering the gotoStatement production.
	EnterGotoStatement(c *GotoStatementContext)

	// EnterProtectedCallStatement is called when entering the protectedCallStatement production.
	EnterProtectedCallStatement(c *ProtectedCallStatementContext)

	// EnterNamedArgs is called when entering the namedArgs production.
	EnterNamedArgs(c *NamedArgsContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterLocalDeclaration is called when entering the localDeclaration production.
	EnterLocalDeclaration(c *LocalDeclarationContext)

	// EnterFunctionDeclaration is called when entering the functionDeclaration production.
	EnterFunctionDeclaration(c *FunctionDeclarationContext)

	// EnterReturnStatement is called when entering the returnStatement production.
	EnterReturnStatement(c *ReturnStatementContext)

	// EnterOperatorGroup is called when entering the operatorGroup production.
	EnterOperatorGroup(c *OperatorGroupContext)

	// EnterLogicalOp is called when entering the logicalOp production.
	EnterLogicalOp(c *LogicalOpContext)

	// EnterComparisonOp is called when entering the comparisonOp production.
	EnterComparisonOp(c *ComparisonOpContext)

	// EnterArithOp is called when entering the arithOp production.
	EnterArithOp(c *ArithOpContext)

	// EnterBitwiseOp is called when entering the bitwiseOp production.
	EnterBitwiseOp(c *BitwiseOpContext)

	// EnterAssignOp is called when entering the assignOp production.
	EnterAssignOp(c *AssignOpContext)

	// EnterUnaryOp is called when entering the unaryOp production.
	EnterUnaryOp(c *UnaryOpContext)

	// EnterConcatOp is called when entering the concatOp production.
	EnterConcatOp(c *ConcatOpContext)

	// EnterVarargOp is called when entering the varargOp production.
	EnterVarargOp(c *VarargOpContext)

	// EnterCompoundAssignOp is called when entering the compoundAssignOp production.
	EnterCompoundAssignOp(c *CompoundAssignOpContext)

	// EnterIncrOp is called when entering the incrOp production.
	EnterIncrOp(c *IncrOpContext)

	// EnterCoalesceOp is called when entering the coalesceOp production.
	EnterCoalesceOp(c *CoalesceOpContext)

	// EnterShiftAssignOp is called when entering the shiftAssignOp production.
	EnterShiftAssignOp(c *ShiftAssignOpContext)

	// EnterNonNullAssertOp is called when entering the nonNullAssertOp production.
	EnterNonNullAssertOp(c *NonNullAssertOpContext)

	// EnterIdentifierList is called when entering the identifierList production.
	EnterIdentifierList(c *IdentifierListContext)

	// EnterExpressionList is called when entering the expressionList production.
	EnterExpressionList(c *ExpressionListContext)

	// EnterFunctionExpression is called when entering the functionExpression production.
	EnterFunctionExpression(c *FunctionExpressionContext)

	// EnterParameterList is called when entering the parameterList production.
	EnterParameterList(c *ParameterListContext)

	// EnterSelectStatement is called when entering the selectStatement production.
	EnterSelectStatement(c *SelectStatementContext)

	// EnterLambdaExpression is called when entering the lambdaExpression production.
	EnterLambdaExpression(c *LambdaExpressionContext)

	// EnterTypeAnnotation is called when entering the typeAnnotation production.
	EnterTypeAnnotation(c *TypeAnnotationContext)

	// EnterTypeSpec is called when entering the typeSpec production.
	EnterTypeSpec(c *TypeSpecContext)

	// EnterExperimentalExpression is called when entering the experimentalExpression production.
	EnterExperimentalExpression(c *ExperimentalExpressionContext)

	// EnterSafeNavigation is called when entering the safeNavigation production.
	EnterSafeNavigation(c *SafeNavigationContext)

	// EnterPipeOperator is called when entering the pipeOperator production.
	EnterPipeOperator(c *PipeOperatorContext)

	// EnterDecoratorSyntax is called when entering the decoratorSyntax production.
	EnterDecoratorSyntax(c *DecoratorSyntaxContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitAssignStatement is called when exiting the assignStatement production.
	ExitAssignStatement(c *AssignStatementContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitPrimaryExpression is called when exiting the primaryExpression production.
	ExitPrimaryExpression(c *PrimaryExpressionContext)

	// ExitMethodChain is called when exiting the methodChain production.
	ExitMethodChain(c *MethodChainContext)

	// ExitPropertyChain is called when exiting the propertyChain production.
	ExitPropertyChain(c *PropertyChainContext)

	// ExitIndexChain is called when exiting the indexChain production.
	ExitIndexChain(c *IndexChainContext)

	// ExitLiteral is called when exiting the literal production.
	ExitLiteral(c *LiteralContext)

	// ExitVariable is called when exiting the variable production.
	ExitVariable(c *VariableContext)

	// ExitSafeAccess is called when exiting the safeAccess production.
	ExitSafeAccess(c *SafeAccessContext)

	// ExitFunctionCall is called when exiting the functionCall production.
	ExitFunctionCall(c *FunctionCallContext)

	// ExitTableConstructor is called when exiting the tableConstructor production.
	ExitTableConstructor(c *TableConstructorContext)

	// ExitMetatable is called when exiting the metatable production.
	ExitMetatable(c *MetatableContext)

	// ExitMetamethods is called when exiting the metamethods production.
	ExitMetamethods(c *MetamethodsContext)

	// ExitLabelStatement is called when exiting the labelStatement production.
	ExitLabelStatement(c *LabelStatementContext)

	// ExitMatchArm is called when exiting the matchArm production.
	ExitMatchArm(c *MatchArmContext)

	// ExitPattern is called when exiting the pattern production.
	ExitPattern(c *PatternContext)

	// ExitFieldPattern is called when exiting the fieldPattern production.
	ExitFieldPattern(c *FieldPatternContext)

	// ExitMetamethod is called when exiting the metamethod production.
	ExitMetamethod(c *MetamethodContext)

	// ExitField is called when exiting the field production.
	ExitField(c *FieldContext)

	// ExitBinaryOperation is called when exiting the binaryOperation production.
	ExitBinaryOperation(c *BinaryOperationContext)

	// ExitUnaryOperation is called when exiting the unaryOperation production.
	ExitUnaryOperation(c *UnaryOperationContext)

	// ExitControlFlowStatement is called when exiting the controlFlowStatement production.
	ExitControlFlowStatement(c *ControlFlowStatementContext)

	// ExitIfStatement is called when exiting the ifStatement production.
	ExitIfStatement(c *IfStatementContext)

	// ExitWhileStatement is called when exiting the whileStatement production.
	ExitWhileStatement(c *WhileStatementContext)

	// ExitRepeatStatement is called when exiting the repeatStatement production.
	ExitRepeatStatement(c *RepeatStatementContext)

	// ExitNumericFor is called when exiting the numericFor production.
	ExitNumericFor(c *NumericForContext)

	// ExitGenericFor is called when exiting the genericFor production.
	ExitGenericFor(c *GenericForContext)

	// ExitBreakStatement is called when exiting the breakStatement production.
	ExitBreakStatement(c *BreakStatementContext)

	// ExitGotoStatement is called when exiting the gotoStatement production.
	ExitGotoStatement(c *GotoStatementContext)

	// ExitProtectedCallStatement is called when exiting the protectedCallStatement production.
	ExitProtectedCallStatement(c *ProtectedCallStatementContext)

	// ExitNamedArgs is called when exiting the namedArgs production.
	ExitNamedArgs(c *NamedArgsContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitLocalDeclaration is called when exiting the localDeclaration production.
	ExitLocalDeclaration(c *LocalDeclarationContext)

	// ExitFunctionDeclaration is called when exiting the functionDeclaration production.
	ExitFunctionDeclaration(c *FunctionDeclarationContext)

	// ExitReturnStatement is called when exiting the returnStatement production.
	ExitReturnStatement(c *ReturnStatementContext)

	// ExitOperatorGroup is called when exiting the operatorGroup production.
	ExitOperatorGroup(c *OperatorGroupContext)

	// ExitLogicalOp is called when exiting the logicalOp production.
	ExitLogicalOp(c *LogicalOpContext)

	// ExitComparisonOp is called when exiting the comparisonOp production.
	ExitComparisonOp(c *ComparisonOpContext)

	// ExitArithOp is called when exiting the arithOp production.
	ExitArithOp(c *ArithOpContext)

	// ExitBitwiseOp is called when exiting the bitwiseOp production.
	ExitBitwiseOp(c *BitwiseOpContext)

	// ExitAssignOp is called when exiting the assignOp production.
	ExitAssignOp(c *AssignOpContext)

	// ExitUnaryOp is called when exiting the unaryOp production.
	ExitUnaryOp(c *UnaryOpContext)

	// ExitConcatOp is called when exiting the concatOp production.
	ExitConcatOp(c *ConcatOpContext)

	// ExitVarargOp is called when exiting the varargOp production.
	ExitVarargOp(c *VarargOpContext)

	// ExitCompoundAssignOp is called when exiting the compoundAssignOp production.
	ExitCompoundAssignOp(c *CompoundAssignOpContext)

	// ExitIncrOp is called when exiting the incrOp production.
	ExitIncrOp(c *IncrOpContext)

	// ExitCoalesceOp is called when exiting the coalesceOp production.
	ExitCoalesceOp(c *CoalesceOpContext)

	// ExitShiftAssignOp is called when exiting the shiftAssignOp production.
	ExitShiftAssignOp(c *ShiftAssignOpContext)

	// ExitNonNullAssertOp is called when exiting the nonNullAssertOp production.
	ExitNonNullAssertOp(c *NonNullAssertOpContext)

	// ExitIdentifierList is called when exiting the identifierList production.
	ExitIdentifierList(c *IdentifierListContext)

	// ExitExpressionList is called when exiting the expressionList production.
	ExitExpressionList(c *ExpressionListContext)

	// ExitFunctionExpression is called when exiting the functionExpression production.
	ExitFunctionExpression(c *FunctionExpressionContext)

	// ExitParameterList is called when exiting the parameterList production.
	ExitParameterList(c *ParameterListContext)

	// ExitSelectStatement is called when exiting the selectStatement production.
	ExitSelectStatement(c *SelectStatementContext)

	// ExitLambdaExpression is called when exiting the lambdaExpression production.
	ExitLambdaExpression(c *LambdaExpressionContext)

	// ExitTypeAnnotation is called when exiting the typeAnnotation production.
	ExitTypeAnnotation(c *TypeAnnotationContext)

	// ExitTypeSpec is called when exiting the typeSpec production.
	ExitTypeSpec(c *TypeSpecContext)

	// ExitExperimentalExpression is called when exiting the experimentalExpression production.
	ExitExperimentalExpression(c *ExperimentalExpressionContext)

	// ExitSafeNavigation is called when exiting the safeNavigation production.
	ExitSafeNavigation(c *SafeNavigationContext)

	// ExitPipeOperator is called when exiting the pipeOperator production.
	ExitPipeOperator(c *PipeOperatorContext)

	// ExitDecoratorSyntax is called when exiting the decoratorSyntax production.
	ExitDecoratorSyntax(c *DecoratorSyntaxContext)
}
