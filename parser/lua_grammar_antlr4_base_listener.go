// Code generated from Lua_grammar_antlr4.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Lua_grammar_antlr4

import "github.com/antlr4-go/antlr/v4"

// BaseLua_grammar_antlr4Listener is a complete listener for a parse tree produced by Lua_grammar_antlr4Parser.
type BaseLua_grammar_antlr4Listener struct{}

var _ Lua_grammar_antlr4Listener = &BaseLua_grammar_antlr4Listener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseLua_grammar_antlr4Listener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseLua_grammar_antlr4Listener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseLua_grammar_antlr4Listener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseLua_grammar_antlr4Listener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseLua_grammar_antlr4Listener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseLua_grammar_antlr4Listener) ExitProgram(ctx *ProgramContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseLua_grammar_antlr4Listener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseLua_grammar_antlr4Listener) ExitStatement(ctx *StatementContext) {}

// EnterAssignStatement is called when production assignStatement is entered.
func (s *BaseLua_grammar_antlr4Listener) EnterAssignStatement(ctx *AssignStatementContext) {}

// ExitAssignStatement is called when production assignStatement is exited.
func (s *BaseLua_grammar_antlr4Listener) ExitAssignStatement(ctx *AssignStatementContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseLua_grammar_antlr4Listener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseLua_grammar_antlr4Listener) ExitExpression(ctx *ExpressionContext) {}

// EnterPrimaryExpression is called when production primaryExpression is entered.
func (s *BaseLua_grammar_antlr4Listener) EnterPrimaryExpression(ctx *PrimaryExpressionContext) {}

// ExitPrimaryExpression is called when production primaryExpression is exited.
func (s *BaseLua_grammar_antlr4Listener) ExitPrimaryExpression(ctx *PrimaryExpressionContext) {}

// EnterMethodChain is called when production methodChain is entered.
func (s *BaseLua_grammar_antlr4Listener) EnterMethodChain(ctx *MethodChainContext) {}

// ExitMethodChain is called when production methodChain is exited.
func (s *BaseLua_grammar_antlr4Listener) ExitMethodChain(ctx *MethodChainContext) {}

// EnterPropertyChain is called when production propertyChain is entered.
func (s *BaseLua_grammar_antlr4Listener) EnterPropertyChain(ctx *PropertyChainContext) {}

// ExitPropertyChain is called when production propertyChain is exited.
func (s *BaseLua_grammar_antlr4Listener) ExitPropertyChain(ctx *PropertyChainContext) {}

// EnterIndexChain is called when production indexChain is entered.
func (s *BaseLua_grammar_antlr4Listener) EnterIndexChain(ctx *IndexChainContext) {}

// ExitIndexChain is called when production indexChain is exited.
func (s *BaseLua_grammar_antlr4Listener) ExitIndexChain(ctx *IndexChainContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BaseLua_grammar_antlr4Listener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BaseLua_grammar_antlr4Listener) ExitLiteral(ctx *LiteralContext) {}

// EnterVariable is called when production variable is entered.
func (s *BaseLua_grammar_antlr4Listener) EnterVariable(ctx *VariableContext) {}

// ExitVariable is called when production variable is exited.
func (s *BaseLua_grammar_antlr4Listener) ExitVariable(ctx *VariableContext) {}

// EnterSafeAccess is called when production safeAccess is entered.
func (s *BaseLua_grammar_antlr4Listener) EnterSafeAccess(ctx *SafeAccessContext) {}

// ExitSafeAccess is called when production safeAccess is exited.
func (s *BaseLua_grammar_antlr4Listener) ExitSafeAccess(ctx *SafeAccessContext) {}

// EnterFunctionCall is called when production functionCall is entered.
func (s *BaseLua_grammar_antlr4Listener) EnterFunctionCall(ctx *FunctionCallContext) {}

// ExitFunctionCall is called when production functionCall is exited.
func (s *BaseLua_grammar_antlr4Listener) ExitFunctionCall(ctx *FunctionCallContext) {}

// EnterTableConstructor is called when production tableConstructor is entered.
func (s *BaseLua_grammar_antlr4Listener) EnterTableConstructor(ctx *TableConstructorContext) {}

// ExitTableConstructor is called when production tableConstructor is exited.
func (s *BaseLua_grammar_antlr4Listener) ExitTableConstructor(ctx *TableConstructorContext) {}

// EnterMetatable is called when production metatable is entered.
func (s *BaseLua_grammar_antlr4Listener) EnterMetatable(ctx *MetatableContext) {}

// ExitMetatable is called when production metatable is exited.
func (s *BaseLua_grammar_antlr4Listener) ExitMetatable(ctx *MetatableContext) {}

// EnterMetamethods is called when production metamethods is entered.
func (s *BaseLua_grammar_antlr4Listener) EnterMetamethods(ctx *MetamethodsContext) {}

// ExitMetamethods is called when production metamethods is exited.
func (s *BaseLua_grammar_antlr4Listener) ExitMetamethods(ctx *MetamethodsContext) {}

// EnterLabelStatement is called when production labelStatement is entered.
func (s *BaseLua_grammar_antlr4Listener) EnterLabelStatement(ctx *LabelStatementContext) {}

// ExitLabelStatement is called when production labelStatement is exited.
func (s *BaseLua_grammar_antlr4Listener) ExitLabelStatement(ctx *LabelStatementContext) {}

// EnterMatchArm is called when production matchArm is entered.
func (s *BaseLua_grammar_antlr4Listener) EnterMatchArm(ctx *MatchArmContext) {}

// ExitMatchArm is called when production matchArm is exited.
func (s *BaseLua_grammar_antlr4Listener) ExitMatchArm(ctx *MatchArmContext) {}

// EnterPattern is called when production pattern is entered.
func (s *BaseLua_grammar_antlr4Listener) EnterPattern(ctx *PatternContext) {}

// ExitPattern is called when production pattern is exited.
func (s *BaseLua_grammar_antlr4Listener) ExitPattern(ctx *PatternContext) {}

// EnterFieldPattern is called when production fieldPattern is entered.
func (s *BaseLua_grammar_antlr4Listener) EnterFieldPattern(ctx *FieldPatternContext) {}

// ExitFieldPattern is called when production fieldPattern is exited.
func (s *BaseLua_grammar_antlr4Listener) ExitFieldPattern(ctx *FieldPatternContext) {}

// EnterMetamethod is called when production metamethod is entered.
func (s *BaseLua_grammar_antlr4Listener) EnterMetamethod(ctx *MetamethodContext) {}

// ExitMetamethod is called when production metamethod is exited.
func (s *BaseLua_grammar_antlr4Listener) ExitMetamethod(ctx *MetamethodContext) {}

// EnterField is called when production field is entered.
func (s *BaseLua_grammar_antlr4Listener) EnterField(ctx *FieldContext) {}

// ExitField is called when production field is exited.
func (s *BaseLua_grammar_antlr4Listener) ExitField(ctx *FieldContext) {}

// EnterBinaryOperation is called when production binaryOperation is entered.
func (s *BaseLua_grammar_antlr4Listener) EnterBinaryOperation(ctx *BinaryOperationContext) {}

// ExitBinaryOperation is called when production binaryOperation is exited.
func (s *BaseLua_grammar_antlr4Listener) ExitBinaryOperation(ctx *BinaryOperationContext) {}

// EnterUnaryOperation is called when production unaryOperation is entered.
func (s *BaseLua_grammar_antlr4Listener) EnterUnaryOperation(ctx *UnaryOperationContext) {}

// ExitUnaryOperation is called when production unaryOperation is exited.
func (s *BaseLua_grammar_antlr4Listener) ExitUnaryOperation(ctx *UnaryOperationContext) {}

// EnterControlFlowStatement is called when production controlFlowStatement is entered.
func (s *BaseLua_grammar_antlr4Listener) EnterControlFlowStatement(ctx *ControlFlowStatementContext) {
}

// ExitControlFlowStatement is called when production controlFlowStatement is exited.
func (s *BaseLua_grammar_antlr4Listener) ExitControlFlowStatement(ctx *ControlFlowStatementContext) {}

// EnterIfStatement is called when production ifStatement is entered.
func (s *BaseLua_grammar_antlr4Listener) EnterIfStatement(ctx *IfStatementContext) {}

// ExitIfStatement is called when production ifStatement is exited.
func (s *BaseLua_grammar_antlr4Listener) ExitIfStatement(ctx *IfStatementContext) {}

// EnterWhileStatement is called when production whileStatement is entered.
func (s *BaseLua_grammar_antlr4Listener) EnterWhileStatement(ctx *WhileStatementContext) {}

// ExitWhileStatement is called when production whileStatement is exited.
func (s *BaseLua_grammar_antlr4Listener) ExitWhileStatement(ctx *WhileStatementContext) {}

// EnterRepeatStatement is called when production repeatStatement is entered.
func (s *BaseLua_grammar_antlr4Listener) EnterRepeatStatement(ctx *RepeatStatementContext) {}

// ExitRepeatStatement is called when production repeatStatement is exited.
func (s *BaseLua_grammar_antlr4Listener) ExitRepeatStatement(ctx *RepeatStatementContext) {}

// EnterNumericFor is called when production numericFor is entered.
func (s *BaseLua_grammar_antlr4Listener) EnterNumericFor(ctx *NumericForContext) {}

// ExitNumericFor is called when production numericFor is exited.
func (s *BaseLua_grammar_antlr4Listener) ExitNumericFor(ctx *NumericForContext) {}

// EnterGenericFor is called when production genericFor is entered.
func (s *BaseLua_grammar_antlr4Listener) EnterGenericFor(ctx *GenericForContext) {}

// ExitGenericFor is called when production genericFor is exited.
func (s *BaseLua_grammar_antlr4Listener) ExitGenericFor(ctx *GenericForContext) {}

// EnterBreakStatement is called when production breakStatement is entered.
func (s *BaseLua_grammar_antlr4Listener) EnterBreakStatement(ctx *BreakStatementContext) {}

// ExitBreakStatement is called when production breakStatement is exited.
func (s *BaseLua_grammar_antlr4Listener) ExitBreakStatement(ctx *BreakStatementContext) {}

// EnterGotoStatement is called when production gotoStatement is entered.
func (s *BaseLua_grammar_antlr4Listener) EnterGotoStatement(ctx *GotoStatementContext) {}

// ExitGotoStatement is called when production gotoStatement is exited.
func (s *BaseLua_grammar_antlr4Listener) ExitGotoStatement(ctx *GotoStatementContext) {}

// EnterProtectedCallStatement is called when production protectedCallStatement is entered.
func (s *BaseLua_grammar_antlr4Listener) EnterProtectedCallStatement(ctx *ProtectedCallStatementContext) {
}

// ExitProtectedCallStatement is called when production protectedCallStatement is exited.
func (s *BaseLua_grammar_antlr4Listener) ExitProtectedCallStatement(ctx *ProtectedCallStatementContext) {
}

// EnterNamedArgs is called when production namedArgs is entered.
func (s *BaseLua_grammar_antlr4Listener) EnterNamedArgs(ctx *NamedArgsContext) {}

// ExitNamedArgs is called when production namedArgs is exited.
func (s *BaseLua_grammar_antlr4Listener) ExitNamedArgs(ctx *NamedArgsContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseLua_grammar_antlr4Listener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseLua_grammar_antlr4Listener) ExitBlock(ctx *BlockContext) {}

// EnterLocalDeclaration is called when production localDeclaration is entered.
func (s *BaseLua_grammar_antlr4Listener) EnterLocalDeclaration(ctx *LocalDeclarationContext) {}

// ExitLocalDeclaration is called when production localDeclaration is exited.
func (s *BaseLua_grammar_antlr4Listener) ExitLocalDeclaration(ctx *LocalDeclarationContext) {}

// EnterFunctionDeclaration is called when production functionDeclaration is entered.
func (s *BaseLua_grammar_antlr4Listener) EnterFunctionDeclaration(ctx *FunctionDeclarationContext) {}

// ExitFunctionDeclaration is called when production functionDeclaration is exited.
func (s *BaseLua_grammar_antlr4Listener) ExitFunctionDeclaration(ctx *FunctionDeclarationContext) {}

// EnterReturnStatement is called when production returnStatement is entered.
func (s *BaseLua_grammar_antlr4Listener) EnterReturnStatement(ctx *ReturnStatementContext) {}

// ExitReturnStatement is called when production returnStatement is exited.
func (s *BaseLua_grammar_antlr4Listener) ExitReturnStatement(ctx *ReturnStatementContext) {}

// EnterOperatorGroup is called when production operatorGroup is entered.
func (s *BaseLua_grammar_antlr4Listener) EnterOperatorGroup(ctx *OperatorGroupContext) {}

// ExitOperatorGroup is called when production operatorGroup is exited.
func (s *BaseLua_grammar_antlr4Listener) ExitOperatorGroup(ctx *OperatorGroupContext) {}

// EnterLogicalOp is called when production logicalOp is entered.
func (s *BaseLua_grammar_antlr4Listener) EnterLogicalOp(ctx *LogicalOpContext) {}

// ExitLogicalOp is called when production logicalOp is exited.
func (s *BaseLua_grammar_antlr4Listener) ExitLogicalOp(ctx *LogicalOpContext) {}

// EnterComparisonOp is called when production comparisonOp is entered.
func (s *BaseLua_grammar_antlr4Listener) EnterComparisonOp(ctx *ComparisonOpContext) {}

// ExitComparisonOp is called when production comparisonOp is exited.
func (s *BaseLua_grammar_antlr4Listener) ExitComparisonOp(ctx *ComparisonOpContext) {}

// EnterArithOp is called when production arithOp is entered.
func (s *BaseLua_grammar_antlr4Listener) EnterArithOp(ctx *ArithOpContext) {}

// ExitArithOp is called when production arithOp is exited.
func (s *BaseLua_grammar_antlr4Listener) ExitArithOp(ctx *ArithOpContext) {}

// EnterBitwiseOp is called when production bitwiseOp is entered.
func (s *BaseLua_grammar_antlr4Listener) EnterBitwiseOp(ctx *BitwiseOpContext) {}

// ExitBitwiseOp is called when production bitwiseOp is exited.
func (s *BaseLua_grammar_antlr4Listener) ExitBitwiseOp(ctx *BitwiseOpContext) {}

// EnterAssignOp is called when production assignOp is entered.
func (s *BaseLua_grammar_antlr4Listener) EnterAssignOp(ctx *AssignOpContext) {}

// ExitAssignOp is called when production assignOp is exited.
func (s *BaseLua_grammar_antlr4Listener) ExitAssignOp(ctx *AssignOpContext) {}

// EnterUnaryOp is called when production unaryOp is entered.
func (s *BaseLua_grammar_antlr4Listener) EnterUnaryOp(ctx *UnaryOpContext) {}

// ExitUnaryOp is called when production unaryOp is exited.
func (s *BaseLua_grammar_antlr4Listener) ExitUnaryOp(ctx *UnaryOpContext) {}

// EnterConcatOp is called when production concatOp is entered.
func (s *BaseLua_grammar_antlr4Listener) EnterConcatOp(ctx *ConcatOpContext) {}

// ExitConcatOp is called when production concatOp is exited.
func (s *BaseLua_grammar_antlr4Listener) ExitConcatOp(ctx *ConcatOpContext) {}

// EnterVarargOp is called when production varargOp is entered.
func (s *BaseLua_grammar_antlr4Listener) EnterVarargOp(ctx *VarargOpContext) {}

// ExitVarargOp is called when production varargOp is exited.
func (s *BaseLua_grammar_antlr4Listener) ExitVarargOp(ctx *VarargOpContext) {}

// EnterCompoundAssignOp is called when production compoundAssignOp is entered.
func (s *BaseLua_grammar_antlr4Listener) EnterCompoundAssignOp(ctx *CompoundAssignOpContext) {}

// ExitCompoundAssignOp is called when production compoundAssignOp is exited.
func (s *BaseLua_grammar_antlr4Listener) ExitCompoundAssignOp(ctx *CompoundAssignOpContext) {}

// EnterIncrOp is called when production incrOp is entered.
func (s *BaseLua_grammar_antlr4Listener) EnterIncrOp(ctx *IncrOpContext) {}

// ExitIncrOp is called when production incrOp is exited.
func (s *BaseLua_grammar_antlr4Listener) ExitIncrOp(ctx *IncrOpContext) {}

// EnterCoalesceOp is called when production coalesceOp is entered.
func (s *BaseLua_grammar_antlr4Listener) EnterCoalesceOp(ctx *CoalesceOpContext) {}

// ExitCoalesceOp is called when production coalesceOp is exited.
func (s *BaseLua_grammar_antlr4Listener) ExitCoalesceOp(ctx *CoalesceOpContext) {}

// EnterShiftAssignOp is called when production shiftAssignOp is entered.
func (s *BaseLua_grammar_antlr4Listener) EnterShiftAssignOp(ctx *ShiftAssignOpContext) {}

// ExitShiftAssignOp is called when production shiftAssignOp is exited.
func (s *BaseLua_grammar_antlr4Listener) ExitShiftAssignOp(ctx *ShiftAssignOpContext) {}

// EnterNonNullAssertOp is called when production nonNullAssertOp is entered.
func (s *BaseLua_grammar_antlr4Listener) EnterNonNullAssertOp(ctx *NonNullAssertOpContext) {}

// ExitNonNullAssertOp is called when production nonNullAssertOp is exited.
func (s *BaseLua_grammar_antlr4Listener) ExitNonNullAssertOp(ctx *NonNullAssertOpContext) {}

// EnterIdentifierList is called when production identifierList is entered.
func (s *BaseLua_grammar_antlr4Listener) EnterIdentifierList(ctx *IdentifierListContext) {}

// ExitIdentifierList is called when production identifierList is exited.
func (s *BaseLua_grammar_antlr4Listener) ExitIdentifierList(ctx *IdentifierListContext) {}

// EnterExpressionList is called when production expressionList is entered.
func (s *BaseLua_grammar_antlr4Listener) EnterExpressionList(ctx *ExpressionListContext) {}

// ExitExpressionList is called when production expressionList is exited.
func (s *BaseLua_grammar_antlr4Listener) ExitExpressionList(ctx *ExpressionListContext) {}

// EnterFunctionExpression is called when production functionExpression is entered.
func (s *BaseLua_grammar_antlr4Listener) EnterFunctionExpression(ctx *FunctionExpressionContext) {}

// ExitFunctionExpression is called when production functionExpression is exited.
func (s *BaseLua_grammar_antlr4Listener) ExitFunctionExpression(ctx *FunctionExpressionContext) {}

// EnterParameterList is called when production parameterList is entered.
func (s *BaseLua_grammar_antlr4Listener) EnterParameterList(ctx *ParameterListContext) {}

// ExitParameterList is called when production parameterList is exited.
func (s *BaseLua_grammar_antlr4Listener) ExitParameterList(ctx *ParameterListContext) {}

// EnterSelectStatement is called when production selectStatement is entered.
func (s *BaseLua_grammar_antlr4Listener) EnterSelectStatement(ctx *SelectStatementContext) {}

// ExitSelectStatement is called when production selectStatement is exited.
func (s *BaseLua_grammar_antlr4Listener) ExitSelectStatement(ctx *SelectStatementContext) {}

// EnterLambdaExpression is called when production lambdaExpression is entered.
func (s *BaseLua_grammar_antlr4Listener) EnterLambdaExpression(ctx *LambdaExpressionContext) {}

// ExitLambdaExpression is called when production lambdaExpression is exited.
func (s *BaseLua_grammar_antlr4Listener) ExitLambdaExpression(ctx *LambdaExpressionContext) {}

// EnterTypeAnnotation is called when production typeAnnotation is entered.
func (s *BaseLua_grammar_antlr4Listener) EnterTypeAnnotation(ctx *TypeAnnotationContext) {}

// ExitTypeAnnotation is called when production typeAnnotation is exited.
func (s *BaseLua_grammar_antlr4Listener) ExitTypeAnnotation(ctx *TypeAnnotationContext) {}

// EnterTypeSpec is called when production typeSpec is entered.
func (s *BaseLua_grammar_antlr4Listener) EnterTypeSpec(ctx *TypeSpecContext) {}

// ExitTypeSpec is called when production typeSpec is exited.
func (s *BaseLua_grammar_antlr4Listener) ExitTypeSpec(ctx *TypeSpecContext) {}

// EnterExperimentalExpression is called when production experimentalExpression is entered.
func (s *BaseLua_grammar_antlr4Listener) EnterExperimentalExpression(ctx *ExperimentalExpressionContext) {
}

// ExitExperimentalExpression is called when production experimentalExpression is exited.
func (s *BaseLua_grammar_antlr4Listener) ExitExperimentalExpression(ctx *ExperimentalExpressionContext) {
}

// EnterSafeNavigation is called when production safeNavigation is entered.
func (s *BaseLua_grammar_antlr4Listener) EnterSafeNavigation(ctx *SafeNavigationContext) {}

// ExitSafeNavigation is called when production safeNavigation is exited.
func (s *BaseLua_grammar_antlr4Listener) ExitSafeNavigation(ctx *SafeNavigationContext) {}

// EnterPipeOperator is called when production pipeOperator is entered.
func (s *BaseLua_grammar_antlr4Listener) EnterPipeOperator(ctx *PipeOperatorContext) {}

// ExitPipeOperator is called when production pipeOperator is exited.
func (s *BaseLua_grammar_antlr4Listener) ExitPipeOperator(ctx *PipeOperatorContext) {}

// EnterDecoratorSyntax is called when production decoratorSyntax is entered.
func (s *BaseLua_grammar_antlr4Listener) EnterDecoratorSyntax(ctx *DecoratorSyntaxContext) {}

// ExitDecoratorSyntax is called when production decoratorSyntax is exited.
func (s *BaseLua_grammar_antlr4Listener) ExitDecoratorSyntax(ctx *DecoratorSyntaxContext) {}
